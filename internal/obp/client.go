package obp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Client struct {
	baseURL     string
	username    string
	password    string
	consumerKey string
	http        *http.Client

	mu    sync.Mutex
	token string
}

func NewClient(baseURL, username, password, consumerKey string) *Client {
	return &Client{
		baseURL:     strings.TrimRight(baseURL, "/"),
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

func (c *Client) GetVerifiedParcels() ([]ParcelOwnerVerification, error) {
	var result ParcelOwnerVerificationList
	err := c.get("/obp/dynamic-entity/parcel_owner_verification",
		url.Values{"status_code": {"verified"}}, &result)
	return result.List, err
}

func (c *Client) GetActivityParcelVerificationByParcelID(parcelID string) (*ActivityParcelVerification, error) {
	var result ActivityParcelVerificationList
	err := c.get("/obp/dynamic-entity/activity_parcel_verification",
		url.Values{"parcel_id": {parcelID}}, &result)
	if err != nil {
		return nil, err
	}
	if len(result.List) == 0 {
		return nil, fmt.Errorf("no activity_parcel_verification found for parcel_id %s", parcelID)
	}
	return &result.List[0], nil
}

// parcelPath is the OBP get-entity-by-id path for a parcel.
func parcelPath(parcelID string) string { return "/obp/dynamic-entity/parcel/" + parcelID }

// ParcelURL is the absolute get-entity-by-id link stored on-chain as parcel_uri.
func (c *Client) ParcelURL(parcelID string) string { return c.baseURL + parcelPath(parcelID) }

// GetParcelRaw returns the raw JSON response body of the parcel entity, used to
// compute the integrity hash (keccak256 of these exact bytes).
func (c *Client) GetParcelRaw(parcelID string) ([]byte, error) {
	return c.getRaw(parcelPath(parcelID), nil)
}

// GetVerifiedActivityVerifications lists activity_verification records with
// status_code=verified — the trigger for ActivityNFT minting.
func (c *Client) GetVerifiedActivityVerifications() ([]ActivityVerification, error) {
	var result ActivityVerificationList
	err := c.get("/obp/dynamic-entity/activity_verification",
		url.Values{"status_code": {"verified"}}, &result)
	return result.List, err
}

// activityPath is the OBP get-entity-by-id path for an activity.
func activityPath(activityID string) string { return "/obp/dynamic-entity/activity/" + activityID }

// ActivityURL is the absolute get-entity-by-id link stored on-chain as activity_url.
func (c *Client) ActivityURL(activityID string) string { return c.baseURL + activityPath(activityID) }

func (c *Client) GetActivity(activityID string) (*Activity, error) {
	activity, _, err := c.GetActivityWithRaw(activityID)
	return activity, err
}

// GetActivityWithRaw returns the decoded activity and the raw response body in a
// single fetch. The raw bytes are hashed (keccak256) into activity_hash.
func (c *Client) GetActivityWithRaw(activityID string) (*Activity, []byte, error) {
	body, err := c.getRaw(activityPath(activityID), nil)
	if err != nil {
		return nil, nil, err
	}
	var wrapper struct {
		Activity Activity `json:"activity"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return nil, nil, fmt.Errorf("decode activity %s: %w", activityID, err)
	}
	return &wrapper.Activity, body, nil
}

// GetCertificates lists all certificate_of_compliance objects. Trigger is
// creation (no status filter); dedup is on-chain by certificate_of_compliance_id.
func (c *Client) GetCertificates() ([]Certificate, error) {
	var result CertificateList
	err := c.get("/obp/dynamic-entity/certificate_of_compliance", nil, &result)
	return result.List, err
}

// certificatePath is the OBP get-entity-by-id path for a certificate_of_compliance.
func certificatePath(id string) string { return "/obp/dynamic-entity/certificate_of_compliance/" + id }

// CertificateURL is the absolute get-entity-by-id link stored as certification_url.
func (c *Client) CertificateURL(id string) string { return c.baseURL + certificatePath(id) }

// GetCertificateRaw returns the raw JSON body of the certificate entity, hashed
// (keccak256) into certification_hash.
func (c *Client) GetCertificateRaw(id string) ([]byte, error) {
	return c.getRaw(certificatePath(id), nil)
}

// GetCertificateByActivityID returns the certificate_of_compliance linked to an
// activity (used to source certification/scheme/body references for a batch).
func (c *Client) GetCertificateByActivityID(activityID string) (*Certificate, error) {
	var result CertificateList
	err := c.get("/obp/dynamic-entity/certificate_of_compliance",
		url.Values{"activity_id": {activityID}}, &result)
	if err != nil {
		return nil, err
	}
	if len(result.List) == 0 {
		return nil, fmt.Errorf("no certificate_of_compliance found for activity_id %s", activityID)
	}
	return &result.List[0], nil
}

// GetVerifiedMonitoringVerifications lists activity_monitoring_period_verification
// records with status_code=verified — the trigger for CarbonCreditBatch minting.
func (c *Client) GetVerifiedMonitoringVerifications() ([]MonitoringVerification, error) {
	var result MonitoringVerificationList
	err := c.get("/obp/dynamic-entity/activity_monitoring_period_verification",
		url.Values{"status_code": {"verified"}}, &result)
	return result.List, err
}

func operatorPath(id string) string { return "/obp/dynamic-entity/operator/" + id }

// OperatorURL is the absolute get-entity-by-id link stored as operator_url.
func (c *Client) OperatorURL(id string) string { return c.baseURL + operatorPath(id) }

// GetOperatorRaw returns the raw JSON body of the operator entity (keccak256 → operator_hash).
func (c *Client) GetOperatorRaw(id string) ([]byte, error) { return c.getRaw(operatorPath(id), nil) }

func (c *Client) GetOperator(operatorID string) (*Operator, error) {
	operator, _, err := c.GetOperatorWithRaw(operatorID)
	return operator, err
}

// GetOperatorWithRaw returns the decoded operator and the raw body in one fetch
// (raw bytes → operator_hash).
func (c *Client) GetOperatorWithRaw(operatorID string) (*Operator, []byte, error) {
	body, err := c.getRaw(operatorPath(operatorID), nil)
	if err != nil {
		return nil, nil, err
	}
	var wrapper struct {
		Operator Operator `json:"operator"`
	}
	if err := json.Unmarshal(body, &wrapper); err != nil {
		return nil, nil, fmt.Errorf("decode operator %s: %w", operatorID, err)
	}
	return &wrapper.Operator, body, nil
}

func schemePath(id string) string { return "/obp/dynamic-entity/certification_scheme/" + id }

// CertificationSchemeURL is the link stored as certification_scheme_url.
func (c *Client) CertificationSchemeURL(id string) string { return c.baseURL + schemePath(id) }

// GetCertificationSchemeRaw returns the raw body (keccak256 → certification_scheme_hash).
func (c *Client) GetCertificationSchemeRaw(id string) ([]byte, error) {
	return c.getRaw(schemePath(id), nil)
}

func bodyPath(id string) string { return "/obp/dynamic-entity/certification_body/" + id }

// CertificationBodyURL is the link stored as certification_body_url.
func (c *Client) CertificationBodyURL(id string) string { return c.baseURL + bodyPath(id) }

// GetCertificationBodyRaw returns the raw body (keccak256 → certification_body_hash).
func (c *Client) GetCertificationBodyRaw(id string) ([]byte, error) {
	return c.getRaw(bodyPath(id), nil)
}
