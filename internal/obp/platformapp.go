package obp

// The tokenizer is a Platform App. Setting it up:
//
//  1. Create an OIDC client for it in OBP-OIDC (OIDC_CLIENT_ID, OIDC_CLIENT_SECRET).
//  2. Start it. OBP creates a Consumer for the client on its first call. The
//     declaration below is refused until step 3, and the log names the
//     Consumer's CONSUMER_ID.
//  3. An administrator marks that Consumer as a Platform App:
//     POST /obp/v7.0.0/management/platform-apps.
//  4. The tokenizer declares the Scopes it needs
//     (PUT /obp/v7.0.0/consumers/current/platform-app). The administrator sees
//     them in GET /obp/v7.0.0/management/platform-apps and grants them with
//     POST /obp/v7.0.0/consumers/CONSUMER_ID/scopes.

import (
	"encoding/json"
	"fmt"
)

// Scope is one Role the app needs, at a bank id.
type Scope struct {
	RoleName  string `json:"role_name"`
	BankID    string `json:"bank_id"`
	NeededFor string `json:"needed_for"`
	Optional  bool   `json:"optional"`
}

// DeclaredScope is a declared Scope as OBP reports it back.
type DeclaredScope struct {
	Scope
	Held bool `json:"held"`
}

// PlatformApp is the app as OBP sees it after a declaration.
type PlatformApp struct {
	ConsumerID     string          `json:"consumer_id"`
	Label          string          `json:"label"`
	RequiredScopes []DeclaredScope `json:"required_scopes"`
}

// Missing returns the required Scopes the app's Consumer does not hold.
func (a *PlatformApp) Missing() []Scope {
	var missing []Scope
	for _, s := range a.RequiredScopes {
		if !s.Held && !s.Optional {
			missing = append(missing, s.Scope)
		}
	}
	return missing
}

// readEntities is every entity the tokenizer reads, with what it reads it for.
var readEntities = []struct{ entity, neededFor string }{
	{entityParcelOwnerVerification, "Finds verified parcels, the trigger for minting a ParcelNFT."},
	{entityParcel, "Reads the parcel whose response body is hashed into the ParcelNFT."},
	{entityParcelActivity, "Links a parcel to its activity, and so to the operator's wallet."},
	{entityActivityVerification, "Finds verified activities, the trigger for minting an ActivityNFT."},
	{entityActivity, "Reads the activity whose response body is hashed into the ActivityNFT."},
	{entityCertificate, "Finds certificates of compliance, the trigger for minting a CertificationNFT."},
	{entityMonitoringVerification, "Finds verified monitoring periods, the trigger for minting a CarbonCreditBatchNFT."},
	{entityOperator, "Reads the operator, whose ogcr_wallet_address receives every token."},
	{entityCertificationScheme, "Reads the certification scheme hashed into the CertificationNFT."},
	{entityCertificationBody, "Reads the certification body hashed into the CertificationNFT."},
	{entityAuditReport, "Reads the audit report hashed into the CertificationNFT."},
}

// RequiredScopes lists the Scopes the tokenizer needs: a read Role on each
// entity it reads, at its bank. It never writes to OBP, so it needs nothing
// else.
func (c *Client) RequiredScopes() []Scope {
	scopes := make([]Scope, 0, len(readEntities))
	for _, e := range readEntities {
		scopes = append(scopes, Scope{
			RoleName:  "CanGetDynamicEntityRecord_" + e.entity,
			BankID:    c.bankID,
			NeededFor: e.neededFor,
		})
	}
	return scopes
}

// DeclarePlatformApp declares the Scopes the tokenizer needs and returns which
// it holds. OBP refuses it (404) until an administrator has marked the app's
// Consumer as a Platform App, and refuses the whole declaration (400) if any
// Role does not exist, which a record Role does not until its entity has been
// created.
func (c *Client) DeclarePlatformApp() (*PlatformApp, error) {
	payload, err := json.Marshal(struct {
		RequiredScopes []Scope `json:"required_scopes"`
	}{c.RequiredScopes()})
	if err != nil {
		return nil, err
	}
	const path = "/obp/v7.0.0/consumers/current/platform-app"
	body, status, err := c.send("PUT", c.baseURL+path, payload)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("PUT %s returned %d: %s", path, status, body)
	}
	var app PlatformApp
	if err := json.Unmarshal(body, &app); err != nil {
		return nil, fmt.Errorf("decode platform app: %w", err)
	}
	return &app, nil
}

// ScopeNames formats Scopes as Role@bank for log messages.
func ScopeNames(scopes []Scope) []string {
	names := make([]string, len(scopes))
	for i, s := range scopes {
		names[i] = s.RoleName + "@" + s.BankID
	}
	return names
}
