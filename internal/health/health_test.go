package health

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func get(t *testing.T, s *State) (report, int) {
	t.Helper()
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	var rep report
	if err := json.NewDecoder(rec.Body).Decode(&rep); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return rep, rec.Code
}

func TestStartingBeforeFirstPoll(t *testing.T) {
	// A process that is up but has never completed a poll is not healthy: it may
	// be unable to reach OBP at all.
	rep, code := get(t, NewState(30*time.Second, 2025))
	if rep.Status != "starting" || code != http.StatusServiceUnavailable {
		t.Fatalf("got %s / %d, want starting / 503", rep.Status, code)
	}
}

func TestHealthyAfterSuccessfulPoll(t *testing.T) {
	s := NewState(30*time.Second, 2025)
	s.RecordPoll("parcels", 3, nil)

	rep, code := get(t, s)
	if rep.Status != "ok" || code != http.StatusOK {
		t.Fatalf("got %s / %d, want ok / 200", rep.Status, code)
	}
	if rep.ChainID != 2025 || rep.EligibleLastPoll["parcels"] != 3 {
		t.Fatalf("report did not carry the poll detail: %+v", rep)
	}
}

func TestDegradedWhenPollFails(t *testing.T) {
	// The case a plain "is the process up" check misses entirely: running, but
	// unable to read its own input.
	s := NewState(30*time.Second, 2025)
	s.RecordPoll("parcels", 0, errors.New("OBP unreachable"))

	rep, code := get(t, s)
	if rep.Status != "degraded" || code != http.StatusServiceUnavailable {
		t.Fatalf("got %s / %d, want degraded / 503", rep.Status, code)
	}
	if rep.LastPollError == "" {
		t.Fatal("degraded report should say why")
	}
}

func TestStalledWhenPollLoopStops(t *testing.T) {
	s := NewState(1*time.Second, 2025)
	s.RecordPoll("parcels", 1, nil)

	// Wind the clock back past three intervals, as if the loop had wedged.
	s.mu.Lock()
	s.lastPollAt = time.Now().UTC().Add(-10 * time.Second)
	s.mu.Unlock()

	rep, code := get(t, s)
	if rep.Status != "stalled" || code != http.StatusServiceUnavailable {
		t.Fatalf("got %s / %d, want stalled / 503", rep.Status, code)
	}
}

func TestRecoversAfterAGoodPoll(t *testing.T) {
	s := NewState(30*time.Second, 2025)
	s.RecordPoll("parcels", 0, errors.New("boom"))
	s.RecordPoll("parcels", 2, nil)

	rep, _ := get(t, s)
	if rep.Status != "ok" || rep.LastPollError != "" {
		t.Fatalf("stale error not cleared: %+v", rep)
	}
}

func TestCountersAreCumulative(t *testing.T) {
	s := NewState(30*time.Second, 2025)
	s.Minted()
	s.Minted()
	s.Failed()
	s.Skipped()
	s.RecordPoll("parcels", 1, nil)

	rep, _ := get(t, s)
	if rep.MintedTotal != 2 || rep.FailedTotal != 1 || rep.SkippedTotal != 1 {
		t.Fatalf("counters wrong: %+v", rep)
	}
}

func TestConcurrentUse(t *testing.T) {
	// Minting runs in goroutines while the poll loop records cycles, so the
	// state has to survive that without racing. Run with -race.
	s := NewState(30*time.Second, 2025)
	done := make(chan struct{})
	for i := 0; i < 8; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				s.Minted()
				s.RecordPoll("parcels", j, nil)
				get(t, s)
			}
			done <- struct{}{}
		}()
	}
	for i := 0; i < 8; i++ {
		<-done
	}
}
