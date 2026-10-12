package obp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

type Client struct {
	baseURL     string
	bankID      string
	username    string
	password    string
	consumerKey string
	http        *http.Client

	mu    sync.Mutex
	token string
}

// NewClient returns a client for the OGCR dynamic entities of one bank. The
// registry keeps them at bank level, so every path is scoped to bankID.
func NewClient(baseURL, bankID, username, password, consumerKey string) *Client {
	return &Client{
		baseURL:     strings.TrimRight(baseURL, "/"),
		bankID:      bankID,
		username:    username,
		password:    password,
		consumerKey: consumerKey,
		http:        &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) authenticate() error {
	req, err := http.NewRequest("POST", c.baseURL+"/my/logins/direct", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf(
		`DirectLogin username="%s",password="%s",consumer_key="%s"`,
		c.username, c.password, c.consumerKey,
	))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("directlogin request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("directlogin failed (%d): %s", resp.StatusCode, body)
	}

	var result struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode token: %w", err)
	}
	if result.Token == "" {
		return fmt.Errorf("empty token in directlogin response")
	}

	c.mu.Lock()
	c.token = result.Token
	c.mu.Unlock()
	return nil
}

func (c *Client) get(path string, params url.Values, out interface{}) error {
	body, err := c.getRaw(path, params)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, out)
}

// getRaw performs an authenticated GET and returns the raw response body.
// It re-authenticates once on a 401 and retries.
func (c *Client) getRaw(path string, params url.Values) ([]byte, error) {
	fullURL := c.baseURL + path
	if len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	token := c.getToken()
	if token == "" {
		if err := c.authenticate(); err != nil {
			return nil, err
		}
		token = c.getToken()
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "DirectLogin token="+token)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s failed: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		// Token expired — re-authenticate once and retry
		if err := c.authenticate(); err != nil {
			return nil, err
		}
		return c.getRaw(path, params)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GET %s: read body: %w", path, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned %d: %s", path, resp.StatusCode, body)
	}
	return body, nil
}

func (c *Client) getToken() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.token
}

// Dynamic entity names read by the tokenizer.
const (
	entityParcel                  = "parcel"
	entityParcelOwnerVerification = "parcel_owner_verification"
	entityParcelActivity          = "parcel_activity"
	entityActivity                = "activity"
	entityActivityVerification    = "activity_verification"
	entityMonitoringVerification  = "activity_monitoring_period_verification"
	entityCertificate             = "certificate_of_compliance"
	entityOperator                = "operator"
	entityCertificationScheme     = "certification_scheme"
	entityCertificationBody       = "certification_body"
	entityAuditReport             = "audit_report"
)

// entityPath is the OBP path of a bank-level dynamic entity collection.
func (c *Client) entityPath(entity string) string {
	return "/obp/dynamic-entity/banks/" + url.PathEscape(c.bankID) + "/" + entity
}

// recordPath is the OBP get-entity-by-id path of one record.
func (c *Client) recordPath(entity, id string) string {
	return c.entityPath(entity) + "/" + url.PathEscape(id)
}

// recordURL is the absolute get-entity-by-id link stored on-chain as the
// record's url/uri.
func (c *Client) recordURL(entity, id string) string {
	return c.baseURL + c.recordPath(entity, id)
}

// list fetches a collection and decodes its "<entity>_list" array into out.
func (c *Client) list(entity string, params url.Values, out interface{}) error {
	body, err := c.getRaw(c.entityPath(entity), params)
	if err != nil {
		return err
	}
	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return fmt.Errorf("decode %s list: %w", entity, err)
	}
	raw, ok := wrapper[entity+"_list"]
	if !ok {
		return fmt.Errorf("%s list: response has no %s_list", entity, entity)
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return fmt.Errorf("decode %s list: %w", entity, err)
	}
	return nil
}

// record fetches one record by id, decodes it into out (when non-nil) and
// returns the raw response body. The raw bytes are what gets hashed
// (keccak256) into the record's on-chain integrity hash.
func (c *Client) record(entity, id string, out interface{}) ([]byte, error) {
	body, err := c.getRaw(c.recordPath(entity, id), nil)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return body, nil
	}
	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return nil, fmt.Errorf("decode %s %s: %w", entity, id, err)
	}
	raw, ok := wrapper[entity]
	if !ok {
		return nil, fmt.Errorf("%s %s: response has no %s object", entity, id, entity)
	}
	if err := json.Unmarshal(raw, out); err != nil {
		return nil, fmt.Errorf("decode %s %s: %w", entity, id, err)
	}
	return body, nil
}

// ---- parcel ----

// GetVerifiedParcels lists parcel_owner_verification records with
// status_code=verified — the trigger for ParcelNFT minting.
func (c *Client) GetVerifiedParcels() ([]ParcelOwnerVerification, error) {
	var out []ParcelOwnerVerification
	err := c.list(entityParcelOwnerVerification, url.Values{"status_code": {"verified"}}, &out)
	return out, err
}

// GetParcelActivityByParcelID returns a parcel_activity link for a parcel. It
// leads from the parcel to its activity and on to the operator wallet.
func (c *Client) GetParcelActivityByParcelID(parcelID string) (*ParcelActivity, error) {
	var out []ParcelActivity
	if err := c.list(entityParcelActivity, url.Values{"parcel_id": {parcelID}}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no parcel_activity found for parcel_id %s", parcelID)
	}
	return &out[0], nil
}

// GetParcelIDsByActivityID lists the parcels linked to an activity through
// parcel_activity, in a stable order.
func (c *Client) GetParcelIDsByActivityID(activityID string) ([]string, error) {
	var links []ParcelActivity
	if err := c.list(entityParcelActivity, url.Values{"activity_id": {activityID}}, &links); err != nil {
		return nil, err
	}
	seen := map[string]bool{}
	ids := []string{}
	for _, l := range links {
		if l.ParcelID == "" || seen[l.ParcelID] {
			continue
		}
		seen[l.ParcelID] = true
		ids = append(ids, l.ParcelID)
	}
	sort.Strings(ids)
	return ids, nil
}

// ParcelURL is the link stored on-chain as parcel_uri.
func (c *Client) ParcelURL(parcelID string) string { return c.recordURL(entityParcel, parcelID) }

// GetParcelRaw returns the raw JSON response body of the parcel entity, used to
// compute the integrity hash (keccak256 of these exact bytes).
func (c *Client) GetParcelRaw(parcelID string) ([]byte, error) {
	return c.record(entityParcel, parcelID, nil)
}

// ---- activity ----

// GetVerifiedActivityVerifications lists activity_verification records with
// status_code=verified — one half of the trigger for ActivityNFT minting.
func (c *Client) GetVerifiedActivityVerifications() ([]ActivityVerification, error) {
	var out []ActivityVerification
	err := c.list(entityActivityVerification, url.Values{"status_code": {"verified"}}, &out)
	return out, err
}

// ActivityURL is the link stored on-chain as activity_url / activity_uri.
func (c *Client) ActivityURL(activityID string) string {
	return c.recordURL(entityActivity, activityID)
}

func (c *Client) GetActivity(activityID string) (*Activity, error) {
	activity, _, err := c.GetActivityWithRaw(activityID)
	return activity, err
}

// GetActivityWithRaw returns the decoded activity and the raw response body in a
// single fetch. The raw bytes are hashed (keccak256) into activity_hash.
func (c *Client) GetActivityWithRaw(activityID string) (*Activity, []byte, error) {
	var a Activity
	body, err := c.record(entityActivity, activityID, &a)
	if err != nil {
		return nil, nil, err
	}
	return &a, body, nil
}

// ---- certificate_of_compliance ----

// GetCertificates lists all certificate_of_compliance records. Trigger is
// creation (no status filter); dedup is on-chain by certificate_of_compliance_id.
func (c *Client) GetCertificates() ([]Certificate, error) {
	var out []Certificate
	err := c.list(entityCertificate, nil, &out)
	return out, err
}

// CertificateURL is the link stored on-chain as certificate_of_compliance_uri.
func (c *Client) CertificateURL(id string) string { return c.recordURL(entityCertificate, id) }

// GetCertificateWithRaw returns the decoded certificate and the raw body in one
// fetch (raw bytes → certificate_of_compliance_hash).
func (c *Client) GetCertificateWithRaw(id string) (*Certificate, []byte, error) {
	var cert Certificate
	body, err := c.record(entityCertificate, id, &cert)
	if err != nil {
		return nil, nil, err
	}
	return &cert, body, nil
}

// GetLatestCertificateByActivityID returns the most recently issued
// certificate_of_compliance of an activity. An activity gains a certificate per
// re-certification, so "the" certificate is the newest one.
func (c *Client) GetLatestCertificateByActivityID(activityID string) (*Certificate, error) {
	var out []Certificate
	if err := c.list(entityCertificate, url.Values{"activity_id": {activityID}}, &out); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no certificate_of_compliance found for activity_id %s", activityID)
	}
	latest := &out[0]
	for i := range out[1:] {
		// ISO dates compare correctly as strings.
		if out[i+1].IssueDate > latest.IssueDate {
			latest = &out[i+1]
		}
	}
	return latest, nil
}

// ---- activity_monitoring_period_verification ----

// GetVerifiedMonitoringVerifications lists activity_monitoring_period_verification
// records with status_code=verified — the trigger for CarbonCreditBatch minting.
func (c *Client) GetVerifiedMonitoringVerifications() ([]MonitoringVerification, error) {
	var out []MonitoringVerification
	err := c.list(entityMonitoringVerification, url.Values{"status_code": {"verified"}}, &out)
	return out, err
}

// ---- operator ----

// OperatorURL is the link stored on-chain as operator_url.
func (c *Client) OperatorURL(id string) string { return c.recordURL(entityOperator, id) }

func (c *Client) GetOperator(operatorID string) (*Operator, error) {
	operator, _, err := c.GetOperatorWithRaw(operatorID)
	return operator, err
}

// GetOperatorWithRaw returns the decoded operator and the raw body in one fetch
// (raw bytes → operator_hash).
func (c *Client) GetOperatorWithRaw(operatorID string) (*Operator, []byte, error) {
	var o Operator
	body, err := c.record(entityOperator, operatorID, &o)
	if err != nil {
		return nil, nil, err
	}
	return &o, body, nil
}

// ---- certification_scheme / certification_body / audit_report ----

// CertificationSchemeURL is the link stored on-chain for the certification scheme.
func (c *Client) CertificationSchemeURL(id string) string {
	return c.recordURL(entityCertificationScheme, id)
}

// GetCertificationSchemeWithRaw returns the decoded scheme and the raw body
// (raw bytes → certification_scheme_hash).
func (c *Client) GetCertificationSchemeWithRaw(id string) (*CertificationScheme, []byte, error) {
	var s CertificationScheme
	body, err := c.record(entityCertificationScheme, id, &s)
	if err != nil {
		return nil, nil, err
	}
	return &s, body, nil
}

// CertificationBodyURL is the link stored on-chain for the certification body.
func (c *Client) CertificationBodyURL(id string) string {
	return c.recordURL(entityCertificationBody, id)
}

// GetCertificationBodyWithRaw returns the decoded body and the raw response
// (raw bytes → certification_body_hash).
func (c *Client) GetCertificationBodyWithRaw(id string) (*CertificationBody, []byte, error) {
	var b CertificationBody
	body, err := c.record(entityCertificationBody, id, &b)
	if err != nil {
		return nil, nil, err
	}
	return &b, body, nil
}

// AuditReportURL is the link stored on-chain as audit_report_uri.
func (c *Client) AuditReportURL(id string) string { return c.recordURL(entityAuditReport, id) }

// GetAuditReportRaw returns the raw body (keccak256 → audit_report_hash).
func (c *Client) GetAuditReportRaw(id string) ([]byte, error) {
	return c.record(entityAuditReport, id, nil)
}
