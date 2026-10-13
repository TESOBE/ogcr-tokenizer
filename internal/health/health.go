// Package health exposes the tokenizer's liveness over HTTP.
//
// The tokenizer reads OBP and mints on chain; it deliberately never writes to
// OBP (see docs/obp-quickstart.md). That rule leaves nothing on the OBP side to
// show it is running: the chain cache runs in the opposite direction, so a
// healthy cache says nothing about this service, and a tokenization backlog of
// zero on a dead tokenizer with no new work looks exactly like success.
//
// An HTTP endpoint answers that without touching the read-only rule, and gives
// the Kubernetes deployment the probes it currently lacks.
package health

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"ogcr-tokenizer/internal/eth"
)

// State is the tokenizer's view of itself. Safe for concurrent use: the poll
// loop records cycles while minting goroutines record outcomes.
type State struct {
	startedAt    time.Time
	pollInterval time.Duration
	chainID      uint64

	mu          sync.RWMutex
	lastPollAt  time.Time
	lastPollOK  bool
	lastPollErr string
	pollCounts  map[string]int

	// What the status page shows beyond the health report; see status.go.
	about       About
	running     bool
	cycleStart  time.Time
	cycleErrors int
	nextPollAt  time.Time
	cycles      []cycle // newest first
	mints       []mint  // newest first
	problems    []*problem
	chain       *eth.ChainInfo
	runNow      func()

	minted  atomic.Int64
	failed  atomic.Int64
	skipped atomic.Int64
}

func NewState(pollInterval time.Duration, chainID uint64) *State {
	return &State{
		startedAt:    time.Now().UTC(),
		pollInterval: pollInterval,
		chainID:      chainID,
		pollCounts:   map[string]int{},
	}
}

// RecordPoll notes the end of one poll of one trigger entity. `found` is how
// many records were eligible, not how many were minted: most are already on
// chain and skipped, and minting happens asynchronously anyway.
func (s *State) RecordPoll(name string, found int, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastPollAt = time.Now().UTC()
	s.pollCounts[name] = found
	if err != nil {
		s.lastPollOK = false
		s.lastPollErr = err.Error()
		return
	}
	s.lastPollOK = true
	s.lastPollErr = ""
}

func (s *State) Minted()  { s.minted.Add(1) }
func (s *State) Failed()  { s.failed.Add(1) }
func (s *State) Skipped() { s.skipped.Add(1) }

type report struct {
	Status              string         `json:"status"`
	StartedAt           string         `json:"started_at"`
	UptimeSeconds       int64          `json:"uptime_seconds"`
	ChainID             uint64         `json:"chain_id"`
	PollIntervalSeconds int            `json:"poll_interval_seconds"`
	LastPollAt          string         `json:"last_poll_at,omitempty"`
	LastPollAgeSeconds  *int64         `json:"last_poll_age_seconds,omitempty"`
	LastPollError       string         `json:"last_poll_error,omitempty"`
	EligibleLastPoll    map[string]int `json:"eligible_last_poll"`
	MintedTotal         int64          `json:"minted_total"`
	FailedTotal         int64          `json:"failed_total"`
	SkippedTotal        int64          `json:"skipped_total"`
}

// snapshot builds the response. Status is "starting" until the first poll
// completes, so a process that comes up but never manages a poll is not
// reported as healthy; "degraded" when the last poll failed, which is the case
// a plain "is the process up" check would miss entirely.
func (s *State) snapshot() (report, int) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	now := time.Now().UTC()
	rep := report{
		StartedAt:           s.startedAt.Format(time.RFC3339),
		UptimeSeconds:       int64(now.Sub(s.startedAt).Seconds()),
		ChainID:             s.chainID,
		PollIntervalSeconds: int(s.pollInterval.Seconds()),
		EligibleLastPoll:    map[string]int{},
		MintedTotal:         s.minted.Load(),
		FailedTotal:         s.failed.Load(),
		SkippedTotal:        s.skipped.Load(),
	}
	for k, v := range s.pollCounts {
		rep.EligibleLastPoll[k] = v
	}

	if s.lastPollAt.IsZero() {
		rep.Status = "starting"
		return rep, http.StatusServiceUnavailable
	}

	age := int64(now.Sub(s.lastPollAt).Seconds())
	rep.LastPollAt = s.lastPollAt.Format(time.RFC3339)
	rep.LastPollAgeSeconds = &age
	rep.LastPollError = s.lastPollErr

	if !s.lastPollOK {
		rep.Status = "degraded"
		return rep, http.StatusServiceUnavailable
	}

	// A poll loop that has stopped ticking is not healthy even though the
	// process is alive, which is the whole point of reporting the age.
	if s.pollInterval > 0 && now.Sub(s.lastPollAt) > 3*s.pollInterval {
		rep.Status = "stalled"
		return rep, http.StatusServiceUnavailable
	}

	rep.Status = "ok"
	return rep, http.StatusOK
}

// Handler serves the health report. Readiness and liveness share it: a
// tokenizer that cannot poll OBP is not doing anything useful, so there is no
// meaningful state where one passes and the other fails.
func (s *State) Handler() http.Handler {
	mux := http.NewServeMux()
	serve := func(w http.ResponseWriter, _ *http.Request) {
		rep, code := s.snapshot()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(rep)
	}
	mux.HandleFunc("/healthz", serve)
	mux.HandleFunc("/health", serve)
	mux.HandleFunc("GET /alive", handleAlive)
	mux.HandleFunc("GET /{$}", s.handleStatus)
	mux.HandleFunc("POST /run", s.handleRun)
	return mux
}

// Serve starts the health server. It is deliberately non-fatal: the tokenizer's
// job is minting, and losing the health endpoint is not a reason to stop doing
// it.
func (s *State) Serve(addr string) {
	if addr == "" {
		slog.Info("health endpoint disabled (no HEALTH_ADDR)")
		return
	}
	go func() {
		slog.Info("health endpoint listening", "addr", addr)
		server := &http.Server{
			Addr:              addr,
			Handler:           s.Handler(),
			ReadHeaderTimeout: 5 * time.Second,
		}
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("health endpoint stopped", "err", err)
		}
	}()
}
