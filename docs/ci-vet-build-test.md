# Go CI: `vet` / `build` / `test`

This document explains the Go continuous-integration workflow (`go vet`,
`go build`, `go test`) — what each step does, why it was chosen, and how it
relates to the OBP container pipeline (`build_container.yml`).

> **Status:** this workflow was removed from `.github/workflows/` on request
> (the old `build.yml` was a pre-skill relict). This doc preserves the rationale
> so it can be reinstated deliberately. The reference YAML is at the bottom.

## Why a Go CI job at all

The repository ships one thing: a Go binary that signs Ethereum transactions and
mints NFTs against a live chain. A bad build reaching production doesn't just
throw a 500 — it can submit malformed or failing transactions, burn gas on
reverts, or mis-map on-chain data. Cheap, fast checks on every push are the first
line of defence, and all three below run in seconds with zero external
dependencies (no chain, no OBP credentials).

The guiding principle: **catch a class of defect at the earliest, cheapest stage
that can catch it.** Compilation errors should never reach a Docker build;
suspicious constructs should never reach code review.

## The three steps

### 1. `go vet ./...` — static correctness

`go vet` is Go's built-in static analyzer. It does not check style; it flags
constructs that compile but are almost always bugs. For this codebase the
high-value vet checks are:

- **`printf` family mismatches** — wrong verbs or argument counts in
  `fmt.Errorf`/`slog` calls. This code logs and wraps errors heavily
  (`fmt.Errorf("mint tx failed: %w", err)`), so a `%w` with no error, or a `%s`
  fed a struct, is a real risk vet catches.
- **Lost `context.Context`** and struct-tag mistakes — the OBP types rely on
  correct `json:"..."` tags (`ogcr_wallet_address`, `activity_id`, …). A
  malformed tag silently breaks unmarshalling; `vet`'s `structtag` check catches
  the malformed-syntax cases.
- **Copying lock values** — the minter holds a `sync.Mutex` for nonce
  management; accidentally copying a struct containing it is a classic
  concurrency bug vet flags.

Vet runs first because it is the fastest and its failures are unambiguous.

### 2. `go build ./...` — it compiles, everywhere

`go build ./...` compiles every package in the module. This is the non-negotiable
gate: **if it doesn't compile, nothing else matters.**

Two reasons it earns its own step even though `go test` also compiles:

- `./...` covers **all** packages including ones with no tests (`cmd/`,
  `internal/contract` generated bindings). `go test` only compiles packages it
  runs, so a broken package with no test file can pass `go test` but fail
  `go build`.
- A clean `build` failure is a clearer signal than a compile error surfaced
  midway through a test run.

This step also guards the `abigen`-generated contract bindings
(`internal/contract/parcel_nft.go`). Regenerating a binding against a new ABI is
routine here; `go build` immediately catches a signature drift between the
binding and its callers in `internal/eth`.

### 3. `go test ./...` — behaviour

`go test ./...` runs every test in the module. Today coverage is thin, but the
step is kept because:

- It costs nothing when there are no tests (a package with no tests reports `ok
  ... [no test files]` and passes).
- It is the natural home for the checks this project most needs as it grows —
  e.g. the keccak256 hashing of parcel bodies, the url-builder helpers, and the
  dedup logic are all pure functions that deserve unit tests. Keeping the CI step
  present means adding a `_test.go` file is all that's required to get CI
  coverage; no workflow change needed.

## Why these three, in this order

They form a cost/latency ladder — each step is a strictly deeper (and slightly
slower) check than the one before, and failing early gives the clearest signal:

| Step | Catches | Cost | Needs |
|------|---------|------|-------|
| `go vet` | Likely-bug constructs that still compile | ~1s | nothing |
| `go build` | Won't-compile across **all** packages | ~seconds | nothing |
| `go test` | Wrong behaviour | seconds+ | nothing |

None require a chain connection, OBP credentials, or Docker — so they run on
every push and on forks, with no secrets. That is exactly the profile of a check
that should gate every commit.

## Why it's separate from the container pipeline

The OBP container workflow (`build_container.yml`) intentionally covers only the
container recipe: build the image, sign it, push it. Language build/test belongs
in its own workflow because:

- **Different failure meaning.** A `go test` failure means "the code is wrong";
  a container-push failure means "publishing is misconfigured". Mixing them
  muddies what a red check tells you.
- **Different gating.** Tests should run on every push and every fork PR
  (no secrets needed). Publishing is gated on Docker Hub secrets being present.
  One workflow can't cleanly have both policies.
- **The OBP skill scope.** The shared OBP pipeline pattern deliberately leaves
  language build/test steps out and untouched — so they live beside it, not
  inside it.

Note the container build is *not* a substitute for this job. `docker build` does
compile the binary (a broken build fails the image build), so it incidentally
covers `go build`. But it does **not** run `go vet` or `go test`, and it only
builds what the Dockerfile targets — so it is a weaker, slower, secret-adjacent
signal than a dedicated Go CI job.

## Workflow design choices

- **`actions/setup-go@v5` with `go-version: "1.25"`** — matches the `go 1.25.x`
  toolchain in `go.mod` and the `golang:1.25-alpine` builder in the Dockerfile,
  so CI and the shipped image compile with the same major version.
- **`cache: true`** — caches the Go build and module cache between runs; on a
  dependency-heavy module (go-ethereum) this is the difference between a
  multi-minute and a sub-minute run.
- **Triggers on push + PR to `main`/`master`, and `v*` tags** — every change and
  every release candidate is checked; PR runs gate merges.
- **No secrets, no services** — the job is hermetic, so it is safe on fork PRs
  and never flakes on external outages.

## Reference workflow

To reinstate, save as `.github/workflows/test.yml`:

```yaml
name: test

on:
  push:
    branches: [main, master]
    tags: ["v*"]
  pull_request:
    branches: [main, master]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: "1.25"
          cache: true
      - run: go vet ./...
      - run: go build ./...
      - run: go test ./...
```
