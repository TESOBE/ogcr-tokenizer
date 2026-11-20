package obp

// The tokenizer calls OBP as a Platform App: as its own application, not as a
// User, the way OGCR-chain-cache and OBP-Sentinel do. It gets a token from
// OBP-OIDC with the OAuth2 client credentials grant and sends it as
// `Authorization: Bearer`. The Roles it needs are Scopes granted to its
// Consumer; see platformapp.go.

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// tokenMargin is how long before its expiry a token is replaced, so a request
// never goes out with one about to lapse.
const tokenMargin = 30 * time.Second

// discoverTokenEndpoint reads the issuer's token_endpoint from its OpenID
// configuration, once. The caller holds c.mu.
func (c *Client) discoverTokenEndpoint() (string, error) {
	if c.tokenEndpoint != "" {
		return c.tokenEndpoint, nil
	}
	resp, err := c.http.Get(c.issuer + "/.well-known/openid-configuration")
	if err != nil {
		return "", fmt.Errorf("oidc discovery failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("oidc discovery failed (%d): %s", resp.StatusCode, body)
	}
	var conf struct {
		TokenEndpoint string `json:"token_endpoint"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&conf); err != nil {
		return "", fmt.Errorf("failed to decode openid configuration: %w", err)
	}
	if conf.TokenEndpoint == "" {
		return "", fmt.Errorf("no token_endpoint in %s/.well-known/openid-configuration", c.issuer)
	}
	c.tokenEndpoint = conf.TokenEndpoint
	return c.tokenEndpoint, nil
}

// fetchToken gets a new access token with the client credentials grant.
// The caller holds c.mu.
func (c *Client) fetchToken() error {
	endpoint, err := c.discoverTokenEndpoint()
	if err != nil {
		return err
	}
	form := url.Values{"grant_type": {"client_credentials"}, "scope": {"openid"}}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	// Unescaped, as OBP-Sentinel and the chain cache send them: OBP-OIDC reads
	// them as given.
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("token request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("token request failed (%d): %s", resp.StatusCode, body)
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode token: %w", err)
	}
	if result.AccessToken == "" {
		return fmt.Errorf("empty access_token in token response")
	}
	expiresIn := time.Duration(result.ExpiresIn) * time.Second
	if expiresIn <= 0 {
		expiresIn = 300 * time.Second
	}
	c.token = result.AccessToken
	c.expiresAt = time.Now().Add(expiresIn - tokenMargin)
	return nil
}

// getToken returns a token that is still valid, fetching a new one if needed.
func (c *Client) getToken() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.token == "" || !time.Now().Before(c.expiresAt) {
		if err := c.fetchToken(); err != nil {
			return "", err
		}
	}
	return c.token, nil
}

// dropToken forgets the cached token, so the next request fetches a new one.
func (c *Client) dropToken() {
	c.mu.Lock()
	c.token = ""
	c.mu.Unlock()
}
