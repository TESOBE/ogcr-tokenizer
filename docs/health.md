# Health endpoint

`GET /healthz` (also `/health`) on `HEALTH_ADDR`, e.g. `:8081`. Unset the
variable and the endpoint is not started, so nothing changes for a deployment
that does not want it.

## Why it exists

The tokenizer reads OBP and mints on chain, and [deliberately never writes to
OBP](./obp-quickstart.md). That rule leaves nothing on the OBP side to show it
is running:

- **OGCR-chain-cache runs the opposite way**, chain to OBP. It can be perfectly
  healthy while this service is dead, because it is mirroring a chain that has
  simply stopped gaining tokens.
- **A tokenization backlog** answers "is it keeping up", but a backlog of zero
  on a dead tokenizer with no new work looks exactly like success.

So liveness is reported over HTTP, which also gives the Kubernetes deployment
the probes it previously lacked.

## What it reports

```json
{
  "status": "ok",
  "started_at": "2026-09-09T06:52:37Z",
  "uptime_seconds": 1,
  "chain_id": 2025,
  "poll_interval_seconds": 15,
  "last_poll_at": "2026-09-09T06:52:38Z",
  "last_poll_age_seconds": 1,
  "eligible_last_poll": {"parcels": 1, "activities": 1, "certificates": 1, "monitoring_periods": 1},
  "minted_total": 1,
  "failed_total": 0,
  "skipped_total": 3
}
```

| Status | HTTP | Meaning |
|---|---|---|
| `starting` | 503 | Up, but no poll has completed yet. Not healthy: it may be unable to reach OBP at all. |
| `ok` | 200 | Last poll succeeded and the loop is ticking. |
| `degraded` | 503 | Last poll failed. The process is running but cannot read its input — the case a bare process check misses. |
| `stalled` | 503 | No poll for more than three intervals. The loop has stopped even though the process is alive. |

Two things worth knowing about the numbers:

- **`eligible_last_poll` counts records the tokenizer *could* mint, not new
  work.** In steady state almost all are already on chain and skipped, so these
  stay constant and are not a sign of a backlog.
- **The totals are cumulative since the process started, not per poll.** Minting
  runs in goroutines that outlive the poll that started them, so a per-cycle
  count would be a guess. A rising `failed_total` alongside a fresh
  `last_poll_at` is the signature of a tokenizer that is running but not working.

Readiness and liveness share the endpoint: a tokenizer that cannot poll OBP is
not doing anything useful, so there is no state where one should pass and the
other fail. The liveness probe is deliberately slack, because a restart mid-mint
is worse than a slow recovery.

## Status page

`GET /` on the same address is a status page for people, modelled on the
OGCR-chain-cache one. It refreshes every 10 seconds and shows:

- the headline status, with the poll timings and the mint totals;
- the chain: node, head block and its age, the signing account and its balance,
  and whether each configured contract has code on this chain;
- how many records each trigger entity had eligible on the last poll;
- **errors in the last hour**: every Error log line, shown once with a count,
  since a record that cannot be minted is retried (and logged) every poll;
- recent mints and recent polls (kept in memory, so they reset on restart).

The headline also turns `degraded` when OBP polls fine but minting cannot work:
the chain is unreachable, a contract is missing, or the signer has no gas.
`/healthz` is unchanged by these, so the probes do not restart a pod over them.

`POST /run` (the "Poll now" button) starts the next poll without waiting out the
interval. `GET /alive` returns `true`, the same as OBP-API and the chain cache.

## On the marketplace status page

Set `TOKENIZER_HEALTH_URL` in OGCR-App to this endpoint and it appears on
`/status` alongside OBP, Redis and the OIDC providers. Left unset it is simply
not monitored, rather than permanently unhealthy.

## Demo mode

`go run ./cmd/demo` serves the status page with made-up data on
`http://127.0.0.1:8767/` (change it with `-addr`). It connects to nothing and
never mints, so it needs no `.env`; it is for seeing the page, not testing the
tokenizer.
