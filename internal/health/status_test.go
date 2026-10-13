package health

import (
	"errors"
	"log/slog"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"ogcr-tokenizer/internal/eth"
)

func page(t *testing.T, s *State) string {
	t.Helper()
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status page: %d", rec.Code)
	}
	return rec.Body.String()
}

func TestStatusPageBeforeFirstPoll(t *testing.T) {
	body := page(t, NewState(30*time.Second, 20253))
	for _, want := range []string{"OGCR Tokenizer", "Starting", "not yet", "None since the process started"} {
		if !strings.Contains(body, want) {
			t.Errorf("page lacks %q", want)
		}
	}
}

func TestStatusPageShowsWhatHappened(t *testing.T) {
	s := NewState(30*time.Second, 20253)
	s.SetAbout(About{OBPURL: "https://obp.example/secret?k=1", OBPBankID: "ogcr", RPCURL: "https://rpc.example/key/abc"})
	s.BeginCycle()
	s.RecordChain(eth.ChainInfo{At: time.Now(), ChainID: 20253, ClientVersion: "besu/v24", HeadBlock: 42, HeadTime: time.Now(),
		SignerBalance: big.NewInt(0),
		Contracts:     []eth.ContractInfo{{Name: "ParcelNFT", Address: common.HexToAddress("0x01"), CodeBytes: 100}}})
	s.RecordPoll("parcels", 3, nil)
	s.RecordMint("ParcelNFT", "parcel-1", "")
	s.EndCycle(30 * time.Second)

	body := page(t, s)
	for _, want := range []string{"parcel-1", "besu", "✓ deployed", "no gas", "https://rpc.example", "https://obp.example"} {
		if !strings.Contains(body, want) {
			t.Errorf("page lacks %q", want)
		}
	}
	for _, leak := range []string{"secret", "/key/abc"} {
		if strings.Contains(body, leak) {
			t.Errorf("page shows %q from a configured URL", leak)
		}
	}
}

func TestLogHandlerCollectsDistinctErrors(t *testing.T) {
	s := NewState(30*time.Second, 20253)
	log := slog.New(LogHandler(slog.DiscardHandler, s)).With("parcel_id", "p1")
	for range 3 {
		log.Error("no recipient", "err", errors.New("operator o1 has no valid ogcr_wallet_address"))
	}
	log.Info("not an error")

	if len(s.problems) != 1 || s.problems[0].Count != 3 {
		t.Fatalf("problems = %+v, want one seen 3 times", s.problems)
	}
	if got := s.problems[0].Text; !strings.Contains(got, "parcel_id=p1") || !strings.Contains(got, "ogcr_wallet_address") {
		t.Errorf("problem text %q lacks its attributes", got)
	}
	if body := page(t, s); !strings.Contains(body, "×3") {
		t.Error("page does not show the repeat count")
	}
}

func TestRunNowButton(t *testing.T) {
	s := NewState(30*time.Second, 20253)
	called := false
	s.SetRunNow(func() { called = true })
	rec := httptest.NewRecorder()
	s.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/run", nil))
	if !called || rec.Code != http.StatusSeeOther {
		t.Fatalf("called=%v code=%d, want true / 303", called, rec.Code)
	}
}
