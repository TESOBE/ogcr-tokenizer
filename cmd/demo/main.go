// Command demo serves the tokenizer's status page with made-up data, to see
// the page without OBP, a chain or a signing key. It never connects to
// anything and never mints.
//
//	go run ./cmd/demo            # http://127.0.0.1:8767/
//	go run ./cmd/demo -addr :9000
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"math/big"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"ogcr-tokenizer/internal/eth"
	"ogcr-tokenizer/internal/health"
)

const chainID = 20253

func main() {
	addr := flag.String("addr", "127.0.0.1:8767", "listen address")
	interval := flag.Duration("interval", 15*time.Second, "simulated poll interval")
	flag.Parse()

	s := health.NewState(*interval, chainID)
	s.SetAbout(health.About{OBPURL: "https://obp.demo.invalid", OBPBankID: "ogcr", RPCURL: "https://rpc.demo.invalid"})
	log := slog.New(health.LogHandler(slog.DiscardHandler, s))

	wake := make(chan struct{}, 1)
	s.SetRunNow(func() {
		select {
		case wake <- struct{}{}:
		default:
		}
	})

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go simulate(ctx, s, log, *interval, wake)

	server := &http.Server{Addr: *addr, Handler: s.Handler(), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()
	fmt.Printf("demo status page: http://%s/  (made-up data, nothing is minted)\n", *addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var contracts = []eth.ContractInfo{
	{Name: "ParcelNFT", Address: common.HexToAddress("0x62227531b82259561CC9ad4413188F08E536598a"), CodeBytes: 14211},
	{Name: "ActivityNFT", Address: common.HexToAddress("0x5034F49b27353CeDc562b49eA91C7438Ea351d36"), CodeBytes: 19870},
	{Name: "CertificationNFT", Address: common.HexToAddress("0x2e335F247E91caa168c64b63104C4475b2af3942"), CodeBytes: 21002},
	{Name: "CarbonCreditBatchNFT", Address: common.HexToAddress("0xe4e47451AAd6C89a6D9E4aD104A7b77FfE1D3b36"), CodeBytes: 23544},
	{Name: "CarbonCredit", Address: common.HexToAddress("0x3f6c622D32dA3BC70730C9E677ec343cb5acFe68"), CodeBytes: 9120},
}

// simulate runs fake polls so the page moves: the head block advances, a
// record is minted now and then, and one parcel keeps failing for want of an
// operator wallet, which is what a stuck record looks like.
func simulate(ctx context.Context, s *health.State, log *slog.Logger, interval time.Duration, wake <-chan struct{}) {
	head := uint64(184000)
	balance, _ := new(big.Int).SetString("4210000000000000000", 10)
	gasPerMint := big.NewInt(3_100_000_000_000_000)
	counts := map[string]int{"parcels": 11, "activities": 4, "certificates": 5, "monitoring_periods": 3}
	serial := 0

	for {
		s.BeginCycle()
		head += uint64(interval/(2*time.Second)) + 1
		s.RecordChain(eth.ChainInfo{
			At: time.Now(), ChainID: chainID, ClientVersion: "besu/v24.12.2/linux-x86_64/openjdk-java-21",
			HeadBlock: head, HeadTime: time.Now().Add(-time.Duration(rand.IntN(4)) * time.Second),
			Signer:        common.HexToAddress("0x9f2a1c0e5b7d4a3e8c6f1b2d3e4f5a6b7c8d9e0f"),
			SignerBalance: new(big.Int).Set(balance), Contracts: contracts,
		})
		time.Sleep(time.Duration(300+rand.IntN(900)) * time.Millisecond) // a poll takes a moment

		if rand.IntN(3) == 0 {
			serial++
			switch serial % 3 {
			case 1:
				counts["parcels"]++
				s.RecordMint("ParcelNFT", fmt.Sprintf("parcel-demo-%02d", 11+serial), "")
			case 2:
				s.RecordMint("CertificationNFT", fmt.Sprintf("coc-demo-%02d", serial), "on ActivityNFT #4")
			default:
				id := fmt.Sprintf("act-04-batch-%02d", serial)
				s.RecordMint("CarbonCreditBatchNFT", id, fmt.Sprintf("removal, token #%d", serial))
				s.RecordMint("CarbonCredit", id, "125000000000000000000 base units")
			}
			s.Minted()
			balance.Sub(balance, gasPerMint)
		} else {
			s.Skipped()
		}
		for name, n := range counts {
			s.RecordPoll(name, n, nil)
		}
		log.With("parcel_id", "parcel-demo-07").Error("no recipient", "err", errors.New("operator op-demo-3 has no valid ogcr_wallet_address"))
		s.EndCycle(interval)

		select {
		case <-ctx.Done():
			return
		case <-wake:
		case <-time.After(interval):
		}
	}
}
