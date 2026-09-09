# OGCR Tokenizer — OBP Quickstart

**Goal:** create the right OBP dynamic entities, in the right order, so the
tokenizer mints each token in the OGCR family.

The tokenizer **polls OBP** (default every 30s), and mints on-chain when it sees
a trigger entity. It never writes to OBP — you (an OBP user with create rights)
create the entities; the tokenizer reacts.

Because of that rule, nothing on the OBP side reveals whether the tokenizer is
running: it only ever reads. It reports on itself over HTTP instead, at
`GET /healthz` on `HEALTH_ADDR` (see [health.md](./health.md)).

- **Exactly once:** every token is deduplicated on-chain, so re-creating or
  re-polling never double-mints.
- **Auto-defer:** tokens that depend on another token (Certification, Carbon
  Credit Batch depend on the Activity NFT) are retried automatically until their
  dependency exists — so if you create things "out of order", the tokenizer
  catches up on a later poll.
- **Recipient:** every token is minted to the **operator's `ogcr_wallet_address`**.

---

## What triggers what

| Token | Trigger entity (what you create) | Also required to exist |
|-------|----------------------------------|------------------------|
| **ParcelNFT** | `parcel_owner_verification` (`status_code=verified`) | parcel, activity_parcel_verification, activity, operator |
| **ActivityNFT** | `activity_verification` (`status_code=verified`) | activity, operator |
| **CertificationNFT** | `certificate_of_compliance` (created) | ActivityNFT already minted, activity, operator |
| **CarbonCreditBatchNFT** (+ bound **CarbonCredit** ERC-20) | `activity_monitoring_period_verification` (`status_code=verified`) | ActivityNFT minted, activity, operator, certificate_of_compliance, certification_scheme, certification_body |

---

## Dependency order (create top-to-bottom)

```
FOUNDATION (create once, reuse):
  1. operator                     ← has ogcr_wallet_address (mint recipient)
  2. activity                     ← activity.operator_id → operator
  3. parcel                       ← the geographic parcel
  4. certification_scheme         ← referenced by carbon-credit batches
  5. certification_body           ← referenced by carbon-credit batches

TRIGGERS (create to mint):
  6. activity_parcel_verification         (parcel ↔ activity link)
  7. parcel_owner_verification  verified  →  mints ParcelNFT
  8. activity_verification      verified  →  mints ActivityNFT   ← do before 9 & 10
  9. certificate_of_compliance            →  mints CertificationNFT
 10. activity_monitoring_period_verification verified → mints CarbonCreditBatch(es) + CarbonCredit
```

You don't have to be strict about 8-before-9/10 — the tokenizer defers 9 and 10
until the Activity NFT exists — but creating `activity_verification` first is the
clean path.

---

## Recipes (minimum to mint each token)

### 1. ParcelNFT
Create, in order:
1. `operator` — with a valid `ogcr_wallet_address`.
2. `activity` — with `operator_id` = your operator.
3. `parcel` — the geographic parcel (any valid `parcel_id`).
4. `activity_parcel_verification` — `{ parcel_id, activity_id }` (links the
   parcel to the activity so the recipient wallet can be resolved).
5. `parcel_owner_verification` — `{ parcel_id, status_code: "verified" }`.

→ Tokenizer mints **ParcelNFT** to the operator wallet, storing
`parcel_id`, `parcel_url`, `parcel_hash`.

### 2. ActivityNFT
1. `operator`, 2. `activity` (as above), then:
3. `activity_verification` — `{ activity_id, status_code: "verified" }`.

→ Tokenizer mints **ActivityNFT** with `activity_id, operator_id, name, type,
start_date, end_date, activity_url, activity_hash`.

### 3. CertificationNFT
Requires the **ActivityNFT** to be minted first (step 2).
1. `certification_scheme` and `certification_body` (referenced by id).
2. `certificate_of_compliance` — referencing `activity_id`,
   `certification_scheme_id`, `certification_body_id`.

→ On any `certificate_of_compliance`, tokenizer mints **CertificationNFT**
(linking the ActivityNFT's token id as `activity_nft_id`).

### 4. CarbonCreditBatchNFT + CarbonCredit
Requires the **ActivityNFT** minted (step 2) **and** a
`certificate_of_compliance` for the activity (step 3), plus
`certification_scheme` and `certification_body`.
1. `activity_monitoring_period_verification` — `{ activity_id,
   status_code: "verified", <benefit fields> }`.

The **benefit fields** (each `> 0` mints one batch, `credit_type` = the field name):
- `permanent_net_carbon_removal_benefit`
- `carbon_farming_temporary_net_carbon_removal_benefit`
- `carbon_farming_net_soil_emission_reduction_benefit`
- `carbon_storage_temporary_net_carbon_removal_benefit`

→ For each benefit field `> 0`, the tokenizer mints one **CarbonCreditBatchNFT**
(ERC-6551) — which auto-creates a token-bound account (TBA) — then mints that
amount of **CarbonCredit** ERC-20 into the TBA.

---

## Key fields the tokenizer reads

Entities may have more fields; only these matter for minting.

| Entity | Fields the tokenizer uses |
|--------|---------------------------|
| `operator` | `operator_id`, `ogcr_wallet_address` |
| `activity` | `activity_id`, `operator_id`, `name`, `type`, `start_date`, `end_date` |
| `parcel` | `parcel_id` (entity must exist; whole response is hashed) |
| `certification_scheme` | `certification_scheme_id` (entity must exist; hashed) |
| `certification_body` | `certification_body_id` (entity must exist; hashed) |
| `parcel_owner_verification` | `parcel_id`, `status_code` (=`verified`) |
| `activity_parcel_verification` | `parcel_id`, `activity_id` |
| `activity_verification` | `activity_id`, `status_code` (=`verified`) |
| `certificate_of_compliance` | `certificate_of_compliance_id`, `activity_id`, `certification_scheme_id`, `certification_body_id`, `issue_date`, `expiry_date`, `certification_status` |
| `activity_monitoring_period_verification` | `activity_id`, `status_code` (=`verified`), the 4 benefit amount fields |

> **`_url` / `_hash` attributes** on the tokens are derived automatically: `_url`
> is the OBP get-entity-by-id link; `_hash` is `keccak256` of that entity's JSON
> response body. You don't set these — the tokenizer computes them.

---

## Prerequisites

- Tokenizer running and pointed at the deployed contracts (its `*_CONTRACT_ADDRESS`
  env vars) with a signing key that is the `minter` on each contract.
- The **tokenizer user** has read roles on every entity above:
  `CanGetDynamicEntity_System<entity>` for `operator`, `activity`, `parcel`,
  `certification_scheme`, `certification_body`, `parcel_owner_verification`,
  `activity_parcel_verification`, `activity_verification`,
  `certificate_of_compliance`, `activity_monitoring_period_verification`.
- **You** (the data creator) have the matching
  `CanCreateDynamicEntity_System<entity>` roles.

`scripts/create_test_data.py` automates a full sample chain (operator → activity →
parcels → verifications) if you want a one-command starting point.

---

## Verifying a mint

Watch the tokenizer logs for `... minted successfully`, or read the chain:

```bash
# Was it minted? (non-zero token id = yes)
cast call <PARCEL_ADDR>  "tokenIdByParcelId(string)(uint256)"   "<parcel_id>"     --rpc-url <RPC>
cast call <ACT_ADDR>     "tokenIdByActivityId(string)(uint256)" "<activity_id>"   --rpc-url <RPC>
cast call <CERT_ADDR>    "tokenIdByComplianceId(string)(uint256)" "<compliance_id>" --rpc-url <RPC>

# See the on-chain metadata/attributes:
cast call <ADDR> "tokenURI(uint256)(string)" <tokenId> --rpc-url <RPC>
# → data:application/json;base64,... (base64-decode to read name + attributes)
```

## Troubleshooting

| Symptom (in logs) | Cause | Fix |
|-------------------|-------|-----|
| `GetActivityParcelVerification failed ... not found` | Parcel is verified but has no `activity_parcel_verification` | Create the parcel↔activity link |
| `ActivityNFT not yet minted ... deferring` | Certificate/monitoring created before the Activity NFT | Create `activity_verification` (verified); tokenizer retries automatically |
| `operator has no ogcr_wallet_address` | Operator missing the wallet field | Set `ogcr_wallet_address` on the operator |
| `403 ... missing one or more roles` | Tokenizer user lacks a read role | Grant `CanGetDynamicEntity_System<entity>` |
| Nothing mints | Trigger not `verified`, or already minted | Check `status_code=verified`; on-chain dedup skips existing tokens |
