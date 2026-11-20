package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	OBPURL string
	// OBPBankID is the bank the OGCR dynamic entities live under.
	OBPBankID string
	// OIDCIssuer is the OBP-OIDC issuer URL; its token endpoint is discovered
	// from <issuer>/.well-known/openid-configuration. OIDCClientID and
	// OIDCClientSecret are the tokenizer's OIDC client: it calls OBP as a
	// Platform App, with the Scopes granted to that client's Consumer.
	OIDCIssuer                   string
	OIDCClientID                 string
	OIDCClientSecret             string
	RPCURL                       string
	PrivateKey                   string
	ParcelContractAddress        string
	ActivityContractAddress      string
	CertificationContractAddress string
	BatchContractAddress         string
	CreditContractAddress        string
	PollInterval                 time.Duration

	// HealthAddr is the listen address for the health endpoint, e.g. ":8081".
	// Optional: empty disables it, so nothing changes for an existing deployment
	// that does not set it.
	HealthAddr string
}

func Load() (*Config, error) {
	// Load .env if present; already-set env vars take precedence.
	_ = godotenv.Load()

	healthAddr := os.Getenv("HEALTH_ADDR")

	bankID := os.Getenv("OBP_BANK_ID")
	if bankID == "" {
		bankID = "ogcr"
	}

	interval := 30 * time.Second
	if s := os.Getenv("POLL_INTERVAL_SECONDS"); s != "" {
		secs, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("invalid POLL_INTERVAL_SECONDS: %w", err)
		}
		interval = time.Duration(secs) * time.Second
	}

	required := map[string]*string{
		"OBP_URL":                        new(string),
		"OIDC_ISSUER":                    new(string),
		"OIDC_CLIENT_ID":                 new(string),
		"OIDC_CLIENT_SECRET":             new(string),
		"RPC_URL":                        new(string),
		"PRIVATE_KEY":                    new(string),
		"PARCEL_CONTRACT_ADDRESS":        new(string),
		"ACTIVITY_CONTRACT_ADDRESS":      new(string),
		"CERTIFICATION_CONTRACT_ADDRESS": new(string),
		"BATCH_CONTRACT_ADDRESS":         new(string),
		"CREDIT_CONTRACT_ADDRESS":        new(string),
	}

	var missing []string
	for key, ptr := range required {
		val := os.Getenv(key)
		if val == "" {
			missing = append(missing, key)
		}
		*ptr = val
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %v", missing)
	}

	return &Config{
		OBPURL:                       *required["OBP_URL"],
		OBPBankID:                    bankID,
		OIDCIssuer:                   *required["OIDC_ISSUER"],
		OIDCClientID:                 *required["OIDC_CLIENT_ID"],
		OIDCClientSecret:             *required["OIDC_CLIENT_SECRET"],
		RPCURL:                       *required["RPC_URL"],
		PrivateKey:                   *required["PRIVATE_KEY"],
		ParcelContractAddress:        *required["PARCEL_CONTRACT_ADDRESS"],
		ActivityContractAddress:      *required["ACTIVITY_CONTRACT_ADDRESS"],
		CertificationContractAddress: *required["CERTIFICATION_CONTRACT_ADDRESS"],
		BatchContractAddress:         *required["BATCH_CONTRACT_ADDRESS"],
		CreditContractAddress:        *required["CREDIT_CONTRACT_ADDRESS"],
		PollInterval:                 interval,
		HealthAddr:                   healthAddr,
	}, nil
}
