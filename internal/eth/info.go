package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// ContractInfo is one configured contract and whether this chain has code at
// its address. A minter pointed at addresses from another chain otherwise
// fails every mint with an opaque revert.
type ContractInfo struct {
	Name      string
	Address   common.Address
	CodeBytes int
	Err       string
}

// ChainInfo is what the status page shows about the chain: the node answering,
// how far it has got, the signing account and the contracts minted to.
type ChainInfo struct {
	At            time.Time
	ChainID       uint64
	ClientVersion string
	HeadBlock     uint64
	HeadTime      time.Time
	Signer        common.Address
	// SignerBalance is nil when it could not be read.
	SignerBalance *big.Int
	Contracts     []ContractInfo
	// Err is set when the node could not be reached at all.
	Err string
}

// Info reads ChainInfo. It never fails: whatever cannot be read is reported in
// the result, since the page is most useful exactly when the chain is unwell.
func (m *Minter) Info(ctx context.Context) ChainInfo {
	info := ChainInfo{At: time.Now(), ChainID: m.ChainID(), Signer: m.auth.From}

	head, err := m.client.HeaderByNumber(ctx, nil)
	if err != nil {
		info.Err = err.Error()
		return info
	}
	info.HeadBlock = head.Number.Uint64()
	info.HeadTime = time.Unix(int64(head.Time), 0)

	var version string
	if err := m.client.Client().CallContext(ctx, &version, "web3_clientVersion"); err == nil {
		info.ClientVersion = version
	}
	if bal, err := m.client.BalanceAt(ctx, m.auth.From, nil); err == nil {
		info.SignerBalance = bal
	}

	for _, c := range m.contracts {
		ci := ContractInfo{Name: c.name, Address: c.address}
		code, err := m.client.CodeAt(ctx, c.address, nil)
		if err != nil {
			ci.Err = err.Error()
		} else {
			ci.CodeBytes = len(code)
		}
		info.Contracts = append(info.Contracts, ci)
	}
	return info
}
