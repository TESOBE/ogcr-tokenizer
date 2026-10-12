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
- **Bank level:** the OGCR entities live under a bank (`OBP_BANK_ID`, default
  `ogcr`). The tokenizer reads `/obp/dynamic-entity/banks/<bank>/<entity>` only.

---

## What triggers what

| Token | Trigger entity (what you create) | Also required to exist |
|-------|----------------------------------|------------------------|
| **ParcelNFT** | `parcel_owner_verification` (`status_code=verified`) | parcel, parcel_activity, activity, operator |
| **ActivityNFT** | `activity_verification` (`status_code=verified`) **and** a `certificate_of_compliance` for the activity | activity, operator |
| **CertificationNFT** | `certificate_of_compliance` (created; a re-certification is a new record and mints a new token) | ActivityNFT already minted, activity, operator, certification_scheme, certification_body, audit_report if referenced |
| **CarbonCreditBatchNFT** (+ bound **CarbonCredit** ERC-20) | `activity_monitoring_period_verification` (`status_code=verified`) | ActivityNFT minted, activity, operator, certificate_of_compliance, certification_scheme, certification_body |

---

## Dependency order (create top-to-bottom)

```
FOUNDATION (create once, reuse):
  1. operator                     ← has ogcr_wallet_address (mint recipient)
  2. activity                     ← activity.operator_id → operator
  3. parcel                       ← the geographic parcel
  4. certification_scheme         ← referenced by certificates and batches
  5. certification_body           ← referenced by certificates and batches

TRIGGERS (create to mint):
  6. parcel_activity                      (parcel ↔ activity link)
  7. parcel_owner_verification  verified  →  mints ParcelNFT
  8. activity_verification      verified  ┐
  9. certificate_of_compliance            ┴→ mints ActivityNFT, then CertificationNFT
 10. activity_monitoring_period_verification verified → mints CarbonCreditBatch(es) + CarbonCredit
```

The ActivityNFT needs both 8 and 9. The CertificationNFT and the batches wait
for the ActivityNFT, so creating things in another order only delays them — the
tokenizer catches up on a later poll.

---

## Recipes (minimum to mint each token)

### 1. ParcelNFT
Create, in order:
1. `operator` — with a valid `ogcr_wallet_address`.
2. `activity` — with `operator_id` = your operator.
3. `parcel` — the geographic parcel (any valid `parcel_id`).
4. `parcel_activity` — `{ parcel_id, activity_id }` (links the
   parcel to the activity so the recipient wallet can be resolved).
5. `parcel_owner_verification` — `{ parcel_id, status_code: "verified" }`.

→ Tokenizer mints **ParcelNFT** to the operator wallet, storing
`parcel_id`, `parcel_uri`, `parcel_hash`.

### 2. ActivityNFT
1. `operator`, 2. `activity` (as above), then:
3. `activity_verification` — `{ activity_id, status_code: "verified" }`.
4. `certificate_of_compliance` — referencing `activity_id` (see recipe 3).

→ Tokenizer mints **ActivityNFT** with `activity_id, operator_id, operator_name`
(operator `legal_name`), `name, summary, website, activity_type, unit_types, city,
country_id, certification_scheme_id, parcel_ids` (every `parcel_activity` of the
activity at mint time), `start_date, end_date, activity_url, activity_hash`.

### 3. CertificationNFT
Requires the **ActivityNFT** to be minted first (step 2).
1. `certification_scheme` and `certification_body` (referenced by id).
2. `certificate_of_compliance` — referencing `activity_id`,
   `certification_scheme_id`, `certification_body_id`, optionally
   `audit_report_id`.

→ On any `certificate_of_compliance`, tokenizer mints **CertificationNFT**
(linking the ActivityNFT's token id as `activity_nft_id`). It stores the
certificate, activity, scheme, body and audit report as uri + hash, the activity,
scheme and body names, `country_id`, `issue_date`, `expiry_date`, `unit_types`,
`certification_status`, and the baseline and benefit quantities the certificate
carries. A quantity the certificate does not carry stays empty on-chain.

A **re-certification** is a new `certificate_of_compliance` record for the same
activity. It mints a further CertificationNFT; the earlier one stays.

### 4. CarbonCreditBatchNFT + CarbonCredit
Requires the **ActivityNFT** minted (step 2) **and** a
`certificate_of_compliance` for the activity (step 3), plus
`certification_scheme` and `certification_body`.
1. `activity_monitoring_period_verification` — `{ activity_id,
   certificate_of_compliance_id, status_code: "verified", <benefit fields> }`.

Each **benefit field** with a value `> 0` mints one batch:

| Benefit field | Batch name | `unit_type` |
|---------------|-----------|-------------|
| `permanent_net_carbon_removal_benefit` | Permanent Removals Carbon Credit Batch | Permanent Removal |
| `carbon_farming_temporary_net_carbon_removal_benefit` | Carbon Farming Removals Carbon Credit Batch | Carbon Farming Sequestration |
| `carbon_farming_net_soil_emission_reduction_benefit` | Soil Emissions Reductions Carbon Credit Batch | Soil Emission Reduction |
| `carbon_storage_temporary_net_carbon_removal_benefit` | Carbon Storage in Product Carbon Credit Batch | Carbon Storage in Product |

→ For each, the tokenizer mints one **CarbonCreditBatchNFT** (ERC-6551) — which
auto-creates a token-bound account (TBA) — then mints the amount as
**CarbonCredit** ERC-20 into the TBA.

- **`amount`** is the benefit value in tonnes CO2e, which may be fractional. It
  is stored and minted in CarbonCredit base units (18 decimals): `6.5` tonnes is
  `6500000000000000000`.
- **`carbon_credit_batch_id`** is `<activity_monitoring_period_verification_id>:<unit>`,
  e.g. `…:permanent_removal`, so the same verification always maps to the same batch.
- **Certificate:** the one named by `certificate_of_compliance_id` on the
  verification; without it, the activity's most recently issued certificate.
  `issue_date` and `expiry_date` of the batch come from that certificate.
- **Monitoring period:** `monitoring_period_start_date` / `_end_date` of the activity.
- **`status_code`** is `issued` at mint.

---

## Key fields the tokenizer reads

Entities may have more fields; only these matter for minting.

| Entity | Fields the tokenizer uses |
|--------|---------------------------|
| `operator` | `operator_id`, `legal_name`, `ogcr_wallet_address` |
| `activity` | `activity_id`, `operator_id`, `name`, `summary`, `website`, `activity_type`, `unit_types`, `city`, `country_id`, `certification_scheme_id`, `start_date`, `end_date`, `monitoring_period_start_date`, `monitoring_period_end_date` |
| `parcel` | `parcel_id` (entity must exist; whole response is hashed) |
| `certification_scheme` | `certification_scheme_id`, `name` (whole response is hashed) |
| `certification_body` | `certification_body_id`, `legal_name` (whole response is hashed) |
| `audit_report` | entity must exist when a certificate references it; whole response is hashed |
| `parcel_owner_verification` | `parcel_id`, `status_code` (=`verified`) |
| `parcel_activity` | `parcel_id`, `activity_id` |
| `activity_verification` | `activity_id`, `status_code` (=`verified`) |
| `certificate_of_compliance` | `certificate_of_compliance_id`, `activity_id`, `certification_scheme_id`, `certification_body_id`, `audit_report_id`, `country_id`, `issue_date`, `expiry_date`, `unit_types`, `certification_status`, `carbon_removals_under_baseline`, `soil_emissions_under_baseline`, the 4 benefit amount fields |
| `activity_monitoring_period_verification` | `activity_id`, `certificate_of_compliance_id`, `status_code` (=`verified`), the 4 benefit amount fields |

> **`_url` / `_uri` / `_hash` attributes** on the tokens are derived automatically:
> `_url` / `_uri` is the OBP get-entity-by-id link; `_hash` is `keccak256` of that
> entity's JSON response body. You don't set these — the tokenizer computes them.

---

## Prerequisites

- Tokenizer running and pointed at the deployed contracts (its `*_CONTRACT_ADDRESS`
  env vars) with a signing key that is the `minter` on each contract.
- The **tokenizer user** has read roles on every entity above, at the bank
  (`OBP_BANK_ID`): `CanGetDynamicEntityRecord_<entity>` for `operator`,
  `activity`, `parcel`, `certification_scheme`, `certification_body`,
  `audit_report`, `parcel_owner_verification`, `parcel_activity`,
  `activity_verification`, `certificate_of_compliance`,
  `activity_monitoring_period_verification`.
- **You** (the data creator) have the matching
  `CanCreateDynamicEntityRecord_<entity>` roles at the same bank.

`scripts/create_test_data.py` predates the move to bank-level entities and
`parcel_activity`; it still writes the old system-level entities and needs
updating before it is useful again.

---

## Verifying a mint

Watch the tokenizer logs for `... minted successfully`, or read the chain:

```bash
# Was it minted? (non-zero token id = yes)
cast call <PARCEL_ADDR>  "tokenIdByParcelId(string)(uint256)"   "<parcel_id>"     --rpc-url <RPC>
cast call <ACT_ADDR>     "tokenIdByActivityId(string)(uint256)" "<activity_id>"   --rpc-url <RPC>
cast call <CERT_ADDR>    "tokenIdByComplianceId(string)(uint256)" "<compliance_id>" --rpc-url <RPC>
cast call <BATCH_ADDR>   "tokenIdByBatchId(string)(uint256)"    "<carbon_credit_batch_id>" --rpc-url <RPC>

# See the on-chain metadata/attributes:
cast call <ADDR> "tokenURI(uint256)(string)" <tokenId> --rpc-url <RPC>
# → data:application/json;base64,... (base64-decode to read name + attributes)
```

## Troubleshooting

| Symptom (in logs) | Cause | Fix |
|-------------------|-------|-----|
| `GetParcelActivity failed ... no parcel_activity found` | Parcel is verified but has no `parcel_activity` | Create the parcel↔activity link |
| `ActivityNFT not yet minted ... deferring` | Certificate/monitoring exists but the activity has no verified `activity_verification` yet | Create `activity_verification` (verified); tokenizer retries automatically |
| `operator ... has no valid ogcr_wallet_address` | Operator wallet field missing or not an address | Set `ogcr_wallet_address` on the operator |
| `invalid benefit amount` | A benefit field is negative or not a plain decimal | Correct the verification record, then restart the tokenizer |
| `403 ... missing one or more roles` | Tokenizer user lacks a read role | Grant `CanGetDynamicEntityRecord_<entity>` at the bank |
| Activity never mints | It is verified but has no `certificate_of_compliance` | Create the certificate |
| Nothing mints | Trigger not `verified`, or already minted | Check `status_code=verified`; on-chain dedup skips existing tokens |
