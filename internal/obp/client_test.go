package obp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// fakeOBP serves fixed bodies by request path (plus query). It is also the
// OIDC issuer: it serves discovery and a client credentials token, and refuses
// any OBP request without that token.
func fakeOBP(t *testing.T, routes map[string]string) *httptest.Server {
	t.Helper()
	var srv *httptest.Server
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/.well-known/openid-configuration":
			_, _ = w.Write([]byte(`{"token_endpoint":"` + srv.URL + `/token"}`))
			return
		case "/token":
			id, secret, ok := r.BasicAuth()
			if !ok || id != "tokenizer" || secret != "s3cret" || r.FormValue("grant_type") != "client_credentials" {
				http.Error(w, `{"error":"invalid_client"}`, http.StatusUnauthorized)
				return
			}
			_, _ = w.Write([]byte(`{"access_token":"t","token_type":"Bearer","expires_in":300}`))
			return
		}
		if r.Header.Get("Authorization") != "Bearer t" {
			http.Error(w, `{"code":401,"message":"OBP-20001: User not logged in."}`, http.StatusUnauthorized)
			return
		}
		key := r.URL.Path
		if r.URL.RawQuery != "" {
			key += "?" + r.URL.RawQuery
		}
		body, ok := routes[key]
		if !ok {
			http.Error(w, `{"code":404,"message":"not found: `+key+`"}`, http.StatusNotFound)
			return
		}
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	return srv
}

// newTestClient is a client of srv for bank ogcr, with the client fakeOBP knows.
func newTestClient(srv *httptest.Server) *Client {
	return NewClient(srv.URL, "ogcr", srv.URL, "tokenizer", "s3cret")
}

func TestBankLevelPathsAndDecoding(t *testing.T) {
	const activityBody = `{"bank_id":"ogcr","activity":{"activity_id":"a-1","name":"Alpha","activity_type":"CARBON_FARMING","country_id":"DK","operator_id":"op-1"}}`
	srv := fakeOBP(t, map[string]string{
		"/obp/dynamic-entity/banks/ogcr/activity/a-1": activityBody,
		"/obp/dynamic-entity/banks/ogcr/parcel_activity?activity_id=a-1": `{"bank_id":"ogcr","parcel_activity_list":[
			{"parcel_id":"p-2","activity_id":"a-1"},{"parcel_id":"p-1","activity_id":"a-1"},{"parcel_id":"p-2","activity_id":"a-1"}]}`,
		"/obp/dynamic-entity/banks/ogcr/certificate_of_compliance?activity_id=a-1": `{"bank_id":"ogcr","certificate_of_compliance_list":[
			{"certificate_of_compliance_id":"coc-old","activity_id":"a-1","issue_date":"2025-01-01"},
			{"certificate_of_compliance_id":"coc-new","activity_id":"a-1","issue_date":"2026-01-01","permanent_net_carbon_removal_benefit":6.5},
			{"certificate_of_compliance_id":"coc-mid","activity_id":"a-1","issue_date":"2025-06-01"}]}`,
	})
	c := newTestClient(srv)

	activity, raw, err := c.GetActivityWithRaw("a-1")
	if err != nil {
		t.Fatal(err)
	}
	if activity.Name != "Alpha" || activity.ActivityType != "CARBON_FARMING" || activity.CountryID != "DK" {
		t.Errorf("unexpected activity: %+v", activity)
	}
	// The hash input must be the exact response body.
	if string(raw) != activityBody {
		t.Errorf("raw body altered: %s", raw)
	}
	if got, want := c.ActivityURL("a-1"), srv.URL+"/obp/dynamic-entity/banks/ogcr/activity/a-1"; got != want {
		t.Errorf("ActivityURL = %s, want %s", got, want)
	}

	parcels, err := c.GetParcelIDsByActivityID("a-1")
	if err != nil {
		t.Fatal(err)
	}
	if len(parcels) != 2 || parcels[0] != "p-1" || parcels[1] != "p-2" {
		t.Errorf("parcels = %v, want [p-1 p-2]", parcels)
	}

	cert, err := c.GetLatestCertificateByActivityID("a-1")
	if err != nil {
		t.Fatal(err)
	}
	if cert.CertificateOfComplianceID != "coc-new" {
		t.Errorf("latest certificate = %s, want coc-new", cert.CertificateOfComplianceID)
	}
	if cert.PermanentNetCarbonRemovalBenefit.String() != "6.5" || cert.SoilEmissionsUnderBaseline.String() != "" {
		t.Errorf("quantities not kept verbatim: %+v", cert)
	}
}

func TestMissingRecordIsAnError(t *testing.T) {
	srv := fakeOBP(t, map[string]string{
		"/obp/dynamic-entity/banks/ogcr/parcel_activity?parcel_id=p-9": `{"bank_id":"ogcr","parcel_activity_list":[]}`,
	})
	c := newTestClient(srv)

	if _, _, err := c.GetActivityWithRaw("missing"); err == nil {
		t.Error("want error for a missing activity")
	}
	if _, err := c.GetParcelActivityByParcelID("p-9"); err == nil {
		t.Error("want error when a parcel has no parcel_activity")
	}
}
