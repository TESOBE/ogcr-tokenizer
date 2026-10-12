package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"ogcr-tokenizer/config"
	"ogcr-tokenizer/internal/eth"
	"ogcr-tokenizer/internal/health"
	"ogcr-tokenizer/internal/obp"
	"ogcr-tokenizer/internal/pipeline"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config error", "err", err)
		os.Exit(1)
	}

	obpClient := obp.NewClient(cfg.OBPURL, cfg.OBPBankID, cfg.OBPUsername, cfg.OBPPassword, cfg.OBPConsumerKey)

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
	healthState := health.NewState(cfg.PollInterval, minter.ChainID())
	healthState.Serve(cfg.HealthAddr)

	p := pipeline.New(obpClient, minter).WithHealth(healthState)
	p.Run(ctx, cfg.PollInterval)
}
