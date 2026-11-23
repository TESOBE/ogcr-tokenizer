package obp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

// TestDeclarationAndTokenRefresh checks the tokenizer declares one read Scope
// per entity at its bank, sends a Bearer token, and gets a new token after a
// 401.
func TestDeclarationAndTokenRefresh(t *testing.T) {
	var tokensIssued, refused atomic.Int32
	var declared []Scope
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/.well-known/openid-configuration":
			_, _ = w.Write([]byte(`{"token_endpoint":"` + srv.URL + `/token"}`))
		case "/token":
			n := tokensIssued.Add(1)
			_, _ = w.Write([]byte(`{"access_token":"t` + string(rune('0'+n)) + `","expires_in":300}`))
		case "/obp/v7.0.0/consumers/current/platform-app":
			// The first token is treated as revoked, so the client must fetch another.
			if r.Header.Get("Authorization") == "Bearer t1" {
				refused.Add(1)
				http.Error(w, `{"code":401}`, http.StatusUnauthorized)
				return
			}
			if r.Method != "PUT" || r.Header.Get("Authorization") != "Bearer t2" {
				t.Errorf("unexpected %s with %q", r.Method, r.Header.Get("Authorization"))
			}
			body, _ := io.ReadAll(r.Body)
			var in struct {
				RequiredScopes []Scope `json:"required_scopes"`
			}
			if err := json.Unmarshal(body, &in); err != nil {
				t.Fatal(err)
			}
			declared = in.RequiredScopes
			out := PlatformApp{ConsumerID: "consumer-1"}
			for i, s := range in.RequiredScopes {
				out.RequiredScopes = append(out.RequiredScopes, DeclaredScope{Scope: s, Held: i > 0})
			}
			_ = json.NewEncoder(w).Encode(out)
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	c := NewClient(srv.URL, "ogcr", srv.URL, "id", "secret")
	app, err := c.DeclarePlatformApp()
	if err != nil {
		t.Fatal(err)
	}
	if refused.Load() != 1 || tokensIssued.Load() != 2 {
		t.Errorf("refused %d, tokens issued %d; want 1 and 2", refused.Load(), tokensIssued.Load())
	}
	if len(declared) != len(readEntities) {
		t.Fatalf("declared %d Scopes, want %d", len(declared), len(readEntities))
	}
	for _, s := range declared {
		if s.BankID != "ogcr" || s.NeededFor == "" || s.Optional {
			t.Errorf("unexpected Scope %+v", s)
		}
	}
	if declared[0].RoleName != "CanGetDynamicEntityRecord_parcel_owner_verification" {
		t.Errorf("first Scope = %s", declared[0].RoleName)
	}
	if missing := app.Missing(); len(missing) != 1 || app.ConsumerID != "consumer-1" {
		t.Errorf("missing = %v, consumer = %s", ScopeNames(missing), app.ConsumerID)
	}
}
