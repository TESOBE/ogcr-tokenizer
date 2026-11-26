package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ogcr-tokenizer/config"
	"ogcr-tokenizer/internal/eth"
	"ogcr-tokenizer/internal/health"
	"ogcr-tokenizer/internal/obp"
	"ogcr-tokenizer/internal/pipeline"
)

func main() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	slog.SetDefault(slog.New(logHandler))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}

	obpClient := obp.NewClient(cfg.OBPURL, cfg.OBPBankID, cfg.OIDCIssuer, cfg.OIDCClientID, cfg.OIDCClientSecret)

	minter, err := eth.NewMinter(cfg.RPCURL, cfg.PrivateKey,
		cfg.ParcelContractAddress, cfg.ActivityContractAddress, cfg.CertificationContractAddress,
		cfg.BatchContractAddress, cfg.CreditContractAddress)
	if err != nil {
		slog.Error("failed to initialise minter", "err", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Liveness over HTTP. The tokenizer never writes to OBP, so this is the only
	// place it can say it is alive; a chain-side backlog answers "is it keeping
	// up" but reads as healthy when a dead tokenizer simply has no new work.
	// The same server carries the status page at /, which lists the errors the
	// pipeline logs, so from here on every log line also reaches the page.
	healthState := health.NewState(cfg.PollInterval, minter.ChainID())
	healthState.SetAbout(health.About{OBPURL: cfg.OBPURL, OBPBankID: cfg.OBPBankID, RPCURL: cfg.RPCURL})
	slog.SetDefault(slog.New(health.LogHandler(logHandler, healthState)))

	go keepDeclaring(ctx, obpClient, healthState, cfg.PollInterval)

	p := pipeline.New(obpClient, minter).WithHealth(healthState)
	healthState.SetRunNow(p.RunNow)
	healthState.Serve(cfg.HealthAddr)
	p.Run(ctx, cfg.PollInterval)
}

// redeclareEvery is how often the Scopes are declared again once all are held,
// so a Scope an administrator takes away shows on the status page.
const redeclareEvery = time.Hour

// keepDeclaring declares the Platform App's Scopes at start, then every poll
// interval until all of them are held, then hourly. A grant shows on the status
// page by the next poll. A failure only warns: a read a missing Scope blocks
// fails on its own with OBP's message.
func keepDeclaring(ctx context.Context, client *obp.Client, state *health.State, pollInterval time.Duration) {
	for {
		app, err := declarePlatformApp(client)
		state.RecordPlatformApp(app, err)
		wait := redeclareEvery
		if err != nil || len(app.Missing()) > 0 {
			wait = pollInterval
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(wait):
		}
	}
}

// declarePlatformApp declares the Scopes the tokenizer needs and says which
// are not granted yet.
func declarePlatformApp(client *obp.Client) (*obp.PlatformApp, error) {
	app, err := client.DeclarePlatformApp()
	if err != nil {
		slog.Warn("could not declare the Platform App's Scopes; the error says why (OBP-35046 means an administrator must mark this Consumer as a Platform App)", "err", err)
		return nil, err
	}
	if missing := app.Missing(); len(missing) > 0 {
		slog.Warn("Platform App is missing Scopes; an administrator must grant them to its Consumer", "consumer_id", app.ConsumerID, "missing", obp.ScopeNames(missing))
	}
	return app, nil
}
