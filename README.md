# OGCR Tokenizer

The OGCR Tokenizer turns verified OGCR records in OBP into tokens on chain. It polls OBP's Dynamic Entities at one bank, and when a verification or certificate is ready it mints the matching token: a ParcelNFT, ActivityNFT, CertificationNFT, or a CarbonCreditBatchNFT with its CarbonCredit ERC-20. Each token is minted to the operator's `ogcr_wallet_address`.

It only reads from OBP and never writes to it. Every token is deduplicated on chain, so polling again never mints twice. A token that depends on another one waits until that one exists, so records created out of order are picked up on a later poll.

[docs/obp-quickstart.md](docs/obp-quickstart.md) says which records trigger which token, in what order to create them, and how to verify a mint.

## Settings

The tokenizer reads its settings from the environment, or from a `.env` file in the folder it runs in. [.env.example](.env.example) lists them all.

| Setting | Required | What it is |
|---------|----------|------------|
| `OBP_URL` | Yes | The OBP-API address, such as `https://dcr.ogcr.tesobe.com`. |
| `OBP_BANK_ID` | No | The bank the OGCR Dynamic Entities live under. Default `ogcr`. |
| `OIDC_ISSUER` | Yes | The OBP-OIDC issuer, see below. |
| `OIDC_CLIENT_ID`, `OIDC_CLIENT_SECRET` | Yes | The OIDC client created for the tokenizer in OBP-OIDC. |
| `RPC_URL` | Yes | The chain's JSON-RPC address. |
| `PRIVATE_KEY` | Yes | The signing key. Its address must be the `minter` on each contract. |
| `PARCEL_CONTRACT_ADDRESS`, `ACTIVITY_CONTRACT_ADDRESS`, `CERTIFICATION_CONTRACT_ADDRESS`, `BATCH_CONTRACT_ADDRESS`, `CREDIT_CONTRACT_ADDRESS` | Yes | The deployed OGCR token contracts. |
| `POLL_INTERVAL_SECONDS` | No | How often it polls OBP. Default 30. |
| `HEALTH_ADDR` | No | Where to serve the status page (`GET /`) and the health endpoint (`GET /healthz`), such as `:8081`. Unset, neither is served. See [docs/health.md](docs/health.md). |

### OIDC_ISSUER

`OIDC_ISSUER` is the `issuer` value OBP-OIDC publishes at `<issuer>/.well-known/openid-configuration`, without a trailing slash. OBP-OIDC serves everything under `/obp-oidc`, so the issuer ends in `/obp-oidc`, not the bare host:

| Where | `OIDC_ISSUER` | `OBP_URL` |
|-------|---------------|-----------|
| Locally | `http://localhost:9000/obp-oidc` | `http://localhost:8080` |
| dcr | `https://oidc.dcr.ogcr.tesobe.com/obp-oidc` | `https://dcr.ogcr.tesobe.com` |

The tokenizer finds the token endpoint from that document. To check an issuer, open `<issuer>/.well-known/openid-configuration`: it should answer with JSON whose `issuer` is exactly the value you set.

## Access to OBP

The tokenizer calls OBP as a Platform App, with its own OAuth2 client credentials, not as a User. Its first call creates its Consumer in OBP. An administrator then marks the Consumer as a Platform App and grants the read Scopes it declares. The status page shows each Scope as held or missing. [Platform App](docs/obp-quickstart.md#platform-app) in the quickstart gives the steps and the Scopes.

## Running it

```sh
cp .env.example .env    # then fill it in
go run ./cmd/main.go
```

Or build the container with the [Dockerfile](Dockerfile). [k8s/](k8s) has a ConfigMap, a Secret and a Deployment for Kubernetes; the Deployment uses the health endpoint for its probes.

To see the status page without OBP, a chain or a signing key, run `go run ./cmd/demo` and open http://127.0.0.1:8767/. It shows made-up data and never mints.

## Tests

```sh
go vet ./...
go test ./...
```

[obp-api-contract/](obp-api-contract) lists every OBP-API call the tokenizer makes. `internal/obp/contract_test.go` checks that list against the source and against `snapshot.json`, which `obp-api-contract/snapshot.py` writes from an OBP-API's resource docs, so a call to an endpoint OBP-API no longer documents fails the tests. [docs/ci-vet-build-test.md](docs/ci-vet-build-test.md) explains the Go checks.

## Licence

AGPL-3.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).
