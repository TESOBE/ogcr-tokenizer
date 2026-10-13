package health

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"ogcr-tokenizer/internal/eth"
)

// The status page at / is the human view of the same State: what the
// tokenizer polls, what it has minted and what keeps failing. It mirrors the
// OGCR-chain-cache status page, the service running the other way.

//go:embed status.html
var statusHTML string

var statusTemplate = template.Must(template.New("status").Parse(statusHTML))

const (
	historySize = 20
	mintsSize   = 50
	// problemsSize bounds the distinct errors kept; problemsMaxAge drops one
	// not seen again for that long, so a fixed problem leaves the page.
	problemsSize   = 100
	problemsMaxAge = time.Hour
)

// About is the static configuration the page names. URLs are reduced to
// scheme and host before they are shown, since an RPC URL can carry an API key.
type About struct {
	OBPURL, OBPBankID, RPCURL string
}

func (s *State) SetAbout(a About) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.about = a
}

// SetRunNow wires the page's "Poll now" button to the poll loop.
func (s *State) SetRunNow(fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.runNow = fn
}

type cycle struct {
	Finished time.Time
	Took     time.Duration
	OK       bool
	Eligible int
	Errors   int
}

type mint struct {
	At     time.Time
	Token  string
	ID     string
	Detail string
}

// problem is one distinct error log line, counted rather than repeated: a
// record that cannot be minted is retried, and so logged, every poll.
type problem struct {
	Text        string
	Count       int
	First, Last time.Time
}

// BeginCycle and EndCycle bracket one pass over every trigger entity.
func (s *State) BeginCycle() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.running, s.cycleStart, s.cycleErrors = true, time.Now(), 0
}

func (s *State) EndCycle(next time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	c := cycle{Finished: now, Took: now.Sub(s.cycleStart), OK: s.lastPollOK, Errors: s.cycleErrors}
	for _, n := range s.pollCounts {
		c.Eligible += n
	}
	s.running, s.nextPollAt = false, now.Add(next)
	s.cycles = append([]cycle{c}, s.cycles...)
	if len(s.cycles) > historySize {
		s.cycles = s.cycles[:historySize]
	}
}

func (s *State) RecordChain(info eth.ChainInfo) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.chain = &info
}

// RecordMint notes a confirmed mint for the page's list of recent mints.
func (s *State) RecordMint(token, id, detail string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mints = append([]mint{{At: time.Now(), Token: token, ID: id, Detail: detail}}, s.mints...)
	if len(s.mints) > mintsSize {
		s.mints = s.mints[:mintsSize]
	}
}

func (s *State) recordProblem(text string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	if s.running {
		s.cycleErrors++
	}
	for _, p := range s.problems {
		if p.Text == text {
			p.Count++
			p.Last = now
			return
		}
	}
	s.problems = append(s.problems, &problem{Text: text, Count: 1, First: now, Last: now})
	if len(s.problems) > problemsSize {
		sort.Slice(s.problems, func(i, j int) bool { return s.problems[i].Last.After(s.problems[j].Last) })
		s.problems = s.problems[:problemsSize]
	}
}

// LogHandler wraps next and also feeds every Error record into s, so that each
// failure the pipeline logs shows on the page without a second reporting path.
func LogHandler(next slog.Handler, s *State) slog.Handler {
	return &logHandler{next: next, state: s}
}

type logHandler struct {
	next  slog.Handler
	state *State
	attrs []slog.Attr
}

// Enabled always takes errors, so the page sees them even where the log
// output itself is filtered.
func (h *logHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= slog.LevelError || h.next.Enabled(ctx, l)
}

func (h *logHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level >= slog.LevelError {
		var b strings.Builder
		b.WriteString(r.Message)
		write := func(a slog.Attr) bool {
			fmt.Fprintf(&b, " %s=%v", a.Key, a.Value)
			return true
		}
		for _, a := range h.attrs {
			write(a)
		}
		r.Attrs(write)
		h.state.recordProblem(b.String())
	}
	if !h.next.Enabled(ctx, r.Level) {
		return nil
	}
	return h.next.Handle(ctx, r)
}

func (h *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logHandler{next: h.next.WithAttrs(attrs), state: h.state, attrs: append(append([]slog.Attr{}, h.attrs...), attrs...)}
}

func (h *logHandler) WithGroup(name string) slog.Handler {
	return &logHandler{next: h.next.WithGroup(name), state: h.state, attrs: h.attrs}
}

// handleAlive is the bare liveness check, the same contract as OBP-API's and
// OGCR-chain-cache's GET /alive: 200 with the JSON literal true.
func handleAlive(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, _ = w.Write([]byte("true"))
}

func (s *State) handleRun(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	fn := s.runNow
	s.mu.RUnlock()
	if fn != nil {
		fn()
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// statusView is everything the template shows, already formatted.
type statusView struct {
	Status, StatusLabel string
	ChainID             uint64
	OBPURL, Bank        string
	Interval, Uptime    string
	Now                 string
	CanRun, Running     bool
	RunningFor          string
	LastPollAgo         string
	LastPollTook        string
	LastPollErr         string
	NextPollIn          string
	Minted, Failed      int64
	Skipped             int64
	Eligible            []eligibleRow
	Chain               *chainView
	Problems            []problemRow
	Mints               []mintRow
	History             []historyRow
}

type eligibleRow struct {
	Name  string
	Count int
}

type chainView struct {
	Client, ClientVersion string
	ChainID               uint64
	RPCURL                string
	HeadBlock             uint64
	HeadTime, HeadAge     string
	CheckedAgo            string
	Signer, Balance       string
	NoGas                 bool
	Err                   string
	Contracts             []eth.ContractInfo
	MissingContracts      int
}

type problemRow struct {
	Text, LastAgo, FirstAgo string
	Count                   int
	Stale                   bool
}

type mintRow struct {
	At, Ago, Token, ID, Detail string
}

type historyRow struct {
	At, Took         string
	OK               bool
	Eligible, Errors int
}

var statusLabels = map[string]string{
	"starting": "Starting: no poll has completed yet",
	"ok":       "OK",
	"degraded": "Degraded: the last poll of OBP failed",
	"stalled":  "Stalled: no poll for more than three intervals",
}

func (s *State) handleStatus(w http.ResponseWriter, _ *http.Request) {
	rep, _ := s.snapshot()

	s.mu.RLock()
	now := time.Now()
	v := statusView{
		Status:   rep.Status,
		ChainID:  s.chainID,
		OBPURL:   hostOnly(s.about.OBPURL),
		Bank:     s.about.OBPBankID,
		Interval: s.pollInterval.String(),
		Uptime:   roundDuration(now.Sub(s.startedAt)),
		Now:      now.Format("15:04:05"),
		CanRun:   s.runNow != nil,
		Running:  s.running,
		Minted:   rep.MintedTotal,
		Failed:   rep.FailedTotal,
		Skipped:  rep.SkippedTotal,
	}
	if s.running {
		v.RunningFor = roundDuration(now.Sub(s.cycleStart))
	} else if !s.nextPollAt.IsZero() {
		v.NextPollIn = roundDuration(s.nextPollAt.Sub(now))
	}
	if !s.lastPollAt.IsZero() {
		v.LastPollAgo = roundDuration(now.Sub(s.lastPollAt))
		v.LastPollErr = s.lastPollErr
	}
	if len(s.cycles) > 0 {
		v.LastPollTook = roundDuration(s.cycles[0].Took)
	}
	for _, name := range []string{"parcels", "activities", "certificates", "monitoring_periods"} {
		if n, ok := s.pollCounts[name]; ok {
			v.Eligible = append(v.Eligible, eligibleRow{name, n})
		}
	}
	if s.chain != nil {
		v.Chain = newChainView(*s.chain, hostOnly(s.about.RPCURL))
	}

	// A problem not seen for three intervals has most likely been fixed or
	// has stopped being retried; it stays, greyed out, until problemsMaxAge.
	staleAfter := 3 * s.pollInterval
	problems := append([]*problem{}, s.problems...)
	sort.Slice(problems, func(i, j int) bool { return problems[i].Last.After(problems[j].Last) })
	for _, p := range problems {
		age := now.Sub(p.Last)
		if age > problemsMaxAge {
			continue
		}
		v.Problems = append(v.Problems, problemRow{
			Text: p.Text, Count: p.Count,
			LastAgo: roundDuration(age), FirstAgo: roundDuration(now.Sub(p.First)),
			Stale: s.pollInterval > 0 && age > staleAfter,
		})
	}
	for _, m := range s.mints {
		v.Mints = append(v.Mints, mintRow{At: m.At.Format("15:04:05"), Ago: roundDuration(now.Sub(m.At)), Token: m.Token, ID: m.ID, Detail: m.Detail})
	}
	for _, c := range s.cycles {
		v.History = append(v.History, historyRow{At: c.Finished.Format("15:04:05"), Took: roundDuration(c.Took), OK: c.OK, Eligible: c.Eligible, Errors: c.Errors})
	}
	s.mu.RUnlock()

	v.StatusLabel = statusLabels[v.Status]
	if v.StatusLabel == "" {
		v.StatusLabel = v.Status
	}
	// Polling OBP fine while every mint is bound to fail should not read as OK.
	if v.Status == "ok" && v.Chain != nil {
		switch {
		case v.Chain.Err != "":
			v.Status, v.StatusLabel = "degraded", "Polling OBP, but the chain is unreachable"
		case v.Chain.MissingContracts > 0:
			v.Status, v.StatusLabel = "degraded", fmt.Sprintf("Polling, but %d configured contract(s) are not on chain %d", v.Chain.MissingContracts, v.Chain.ChainID)
		case v.Chain.NoGas:
			v.Status, v.StatusLabel = "degraded", "Polling, but the signing account has no gas"
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := statusTemplate.Execute(w, v); err != nil {
		slog.Warn("render status page", "err", err)
	}
}

func newChainView(info eth.ChainInfo, rpcURL string) *chainView {
	v := &chainView{
		Client:        "unknown",
		ClientVersion: info.ClientVersion,
		ChainID:       info.ChainID,
		RPCURL:        rpcURL,
		HeadBlock:     info.HeadBlock,
		CheckedAgo:    roundDuration(time.Since(info.At)),
		Signer:        info.Signer.Hex(),
		Balance:       "–",
		Err:           info.Err,
		Contracts:     info.Contracts,
	}
	if name, _, _ := strings.Cut(info.ClientVersion, "/"); name != "" {
		v.Client = name
	}
	if !info.HeadTime.IsZero() {
		v.HeadTime = info.HeadTime.Format("2006-01-02 15:04:05")
		v.HeadAge = roundDuration(time.Since(info.HeadTime))
	}
	if info.SignerBalance != nil {
		v.Balance = formatEther(info.SignerBalance)
		v.NoGas = info.SignerBalance.Sign() == 0
	}
	for _, c := range info.Contracts {
		if c.Err == "" && c.CodeBytes == 0 {
			v.MissingContracts++
		}
	}
	return v
}

// formatEther shows a wei amount in whole units to four places.
func formatEther(wei *big.Int) string {
	f := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return f.Text('f', 4)
}

// hostOnly drops everything after the host from a URL, since an RPC URL can
// carry an API key in its path or query.
func hostOnly(raw string) string {
	if raw == "" {
		return "–"
	}
	u, err := url.Parse(raw)
	if err != nil || u.Host == "" {
		return "(not a URL)"
	}
	return u.Scheme + "://" + u.Host
}

// roundDuration formats a duration for people: 4s, 2m10s, 3h5m0s.
func roundDuration(d time.Duration) string {
	switch {
	case d < 0:
		return "0s"
	case d < time.Minute:
		return fmt.Sprintf("%ds", int(d.Seconds()))
	case d < time.Hour:
		return d.Round(time.Second).String()
	default:
		return d.Round(time.Minute).String()
	}
}
