package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"ogcr-tokenizer/internal/contract"
)

// Minter signs and submits mints for the OGCR token family from a single
// account. All contracts share one signing key, so nonce management is shared
// (sequential, mutex-guarded) across every token type.
type Minter struct {
	client      *ethclient.Client
	parcel      *contract.ParcelNFT
	activity    *contract.ActivityNFT
	certificate *contract.CertificationNFT
	batch       *contract.CarbonCreditBatchNFT
	credit      *contract.CarbonCredit
	auth        *bind.TransactOpts
	chainID     *big.Int
	nonceMu     sync.Mutex
	nonce       uint64

	// contracts lists the configured addresses for the status page.
	contracts []namedAddress
}

type namedAddress struct {
	name    string
	address common.Address
}

// ChainID is the chain the minter is bound to. Exposed for the health report,
// which names the chain it is minting to.
func (m *Minter) ChainID() uint64 {
	if m.chainID == nil {
		return 0
	}
	return m.chainID.Uint64()
}

func NewMinter(rpcURL, privateKeyHex, parcelAddr, activityAddr, certificateAddr, batchAddr, creditAddr string) (*Minter, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	parcel, err := contract.NewParcelNFT(common.HexToAddress(parcelAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind ParcelNFT: %w", err)
	}
	activity, err := contract.NewActivityNFT(common.HexToAddress(activityAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind ActivityNFT: %w", err)
	}
	certificate, err := contract.NewCertificationNFT(common.HexToAddress(certificateAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CertificationNFT: %w", err)
	}
	batch, err := contract.NewCarbonCreditBatchNFT(common.HexToAddress(batchAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CarbonCreditBatchNFT: %w", err)
	}
	credit, err := contract.NewCarbonCredit(common.HexToAddress(creditAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CarbonCredit: %w", err)
	}

	pubKey := privKey.Public().(*ecdsa.PublicKey)
	fromAddr := crypto.PubkeyToAddress(*pubKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	return &Minter{
		client:      client,
		parcel:      parcel,
		activity:    activity,
		certificate: certificate,
		batch:       batch,
		credit:      credit,
		auth:        auth,
		chainID:     chainID,
		nonce:       nonce,
		contracts: []namedAddress{
			{"ParcelNFT", common.HexToAddress(parcelAddr)},
			{"ActivityNFT", common.HexToAddress(activityAddr)},
			{"CertificationNFT", common.HexToAddress(certificateAddr)},
			{"CarbonCreditBatchNFT", common.HexToAddress(batchAddr)},
			{"CarbonCredit", common.HexToAddress(creditAddr)},
		},
	}, nil
}

// ---- ParcelNFT ----

type ParcelParams struct {
	To         common.Address
	ParcelID   string
	ParcelURI  string
	ParcelHash string
}

// ParcelAlreadyMinted reports whether a ParcelNFT exists for this parcel_id.
// The contract mapping returns 0 for unknown ids (no revert).
func (m *Minter) ParcelAlreadyMinted(ctx context.Context, parcelID string) (bool, error) {
	tokenID, err := m.parcel.TokenIdByParcelId(&bind.CallOpts{Context: ctx}, parcelID)
	if err != nil {
		return false, fmt.Errorf("tokenIdByParcelId(%s): %w", parcelID, err)
	}
	return tokenID.Sign() != 0, nil
}

func (m *Minter) MintParcel(ctx context.Context, p ParcelParams) error {
	return m.send(ctx, "ParcelNFT parcel_id="+p.ParcelID, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.parcel.Mint(opts, p.To, p.ParcelID, p.ParcelURI, p.ParcelHash)
	})
}

// ---- ActivityNFT ----

// ActivityParams is the ActivityNFT mint input: the recipient plus the on-chain
// ActivityData.
type ActivityParams struct {
	To   common.Address
	Data contract.ActivityNFTActivityData
}

// ActivityAlreadyMinted reports whether an ActivityNFT exists for this activity_id.
func (m *Minter) ActivityAlreadyMinted(ctx context.Context, activityID string) (bool, error) {
	tokenID, err := m.ActivityTokenID(ctx, activityID)
	if err != nil {
		return false, err
	}
	return tokenID.Sign() != 0, nil
}

// ActivityTokenID returns the ActivityNFT token id for an activity_id (0 if not
// yet minted). Dependent tokens (Certification, Batch) reference it as activity_nft_id.
func (m *Minter) ActivityTokenID(ctx context.Context, activityID string) (*big.Int, error) {
	tokenID, err := m.activity.TokenIdByActivityId(&bind.CallOpts{Context: ctx}, activityID)
	if err != nil {
		return nil, fmt.Errorf("tokenIdByActivityId(%s): %w", activityID, err)
	}
	return tokenID, nil
}

func (m *Minter) MintActivity(ctx context.Context, p ActivityParams) error {
	return m.send(ctx, "ActivityNFT activity_id="+p.Data.ActivityId, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.activity.Mint(opts, p.To, p.Data)
	})
}

// ---- CertificationNFT ----

// CertificationParams is the CertificationNFT mint input: the recipient plus
// the on-chain CertificationData.
type CertificationParams struct {
	To   common.Address
	Data contract.CertificationNFTCertificationData
}

// CertificationAlreadyMinted reports whether a CertificationNFT exists for this
// certificate_of_compliance_id.
func (m *Minter) CertificationAlreadyMinted(ctx context.Context, complianceID string) (bool, error) {
	tokenID, err := m.certificate.TokenIdByComplianceId(&bind.CallOpts{Context: ctx}, complianceID)
	if err != nil {
		return false, fmt.Errorf("tokenIdByComplianceId(%s): %w", complianceID, err)
	}
	return tokenID.Sign() != 0, nil
}

func (m *Minter) MintCertification(ctx context.Context, p CertificationParams) error {
	return m.send(ctx, "CertificationNFT compliance_id="+p.Data.CertificateOfComplianceId, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.certificate.Mint(opts, p.To, p.Data)
	})
}

// ---- CarbonCreditBatchNFT (ERC-6551) + CarbonCredit (ERC-20) ----

// BatchParams is the CarbonCreditBatchNFT mint input: the recipient plus the
// on-chain BatchData.
type BatchParams struct {
	To   common.Address
	Data contract.CarbonCreditBatchNFTBatchData
}

// FindBatch returns the token id, token-bound account and issued amount of the
// batch minted for a carbon_credit_batch_id, or found=false.
func (m *Minter) FindBatch(ctx context.Context, batchID string) (tokenID *big.Int, tba common.Address, amount *big.Int, found bool, err error) {
	opts := &bind.CallOpts{Context: ctx}
	tokenID, err = m.batch.TokenIdByBatchId(opts, batchID)
	if err != nil {
		return nil, common.Address{}, nil, false, fmt.Errorf("tokenIdByBatchId(%s): %w", batchID, err)
	}
	if tokenID.Sign() == 0 {
		return nil, common.Address{}, nil, false, nil
	}
	b, err := m.batch.GetBatch(opts, tokenID)
	if err != nil {
		return nil, common.Address{}, nil, false, fmt.Errorf("getBatch(%s): %w", tokenID, err)
	}
	tba, err = m.batch.GetTBA(opts, tokenID)
	if err != nil {
		return nil, common.Address{}, nil, false, fmt.Errorf("getTBA(%s): %w", tokenID, err)
	}
	return tokenID, tba, b.Amount, true, nil
}

// CreditMintedTo returns the cumulative CarbonCredit amount ever minted to an
// address (a batch TBA). Unlike the balance it does not drop when credits are
// moved out or retired, so it is safe to compare against the batch amount.
func (m *Minter) CreditMintedTo(ctx context.Context, addr common.Address) (*big.Int, error) {
	minted, err := m.credit.MintedTo(&bind.CallOpts{Context: ctx}, addr)
	if err != nil {
		return nil, fmt.Errorf("credit mintedTo(%s): %w", addr.Hex(), err)
	}
	return minted, nil
}

// MintBatch mints a CarbonCreditBatchNFT and returns the created token-bound
// account (TBA) parsed from the BatchMinted event.
func (m *Minter) MintBatch(ctx context.Context, p BatchParams) (tokenID *big.Int, tba common.Address, err error) {
	receipt, err := m.sendTx(ctx, "CarbonCreditBatchNFT batch_id="+p.Data.CarbonCreditBatchId, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.batch.Mint(opts, p.To, p.Data)
	})
	if err != nil {
		return nil, common.Address{}, err
	}
	for _, lg := range receipt.Logs {
		ev, perr := m.batch.ParseBatchMinted(*lg)
		if perr == nil {
			return ev.TokenId, ev.Tba, nil
		}
	}
	return nil, common.Address{}, fmt.Errorf("BatchMinted event not found in receipt for tx %s", receipt.TxHash.Hex())
}

// MintCredit mints `amount` CarbonCredit ERC-20 tokens into the batch's TBA.
func (m *Minter) MintCredit(ctx context.Context, tba common.Address, amount *big.Int) error {
	return m.send(ctx, "CarbonCredit to="+tba.Hex(), func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return m.credit.Mint(opts, tba, amount)
	})
}

// ---- shared submit ----

// send submits a tx and waits for a successful receipt, discarding it.
func (m *Minter) send(ctx context.Context, label string, fn func(*bind.TransactOpts) (*types.Transaction, error)) error {
	_, err := m.sendTx(ctx, label, fn)
	return err
}

// sendTx builds nonce-managed transact opts, invokes fn to submit the tx, waits
// for the receipt, and returns it (needed to parse emitted events, e.g. the
// batch's token-bound account). label identifies the operation in logs/errors.
func (m *Minter) sendTx(ctx context.Context, label string, fn func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.nonceMu.Lock()
	opts := &bind.TransactOpts{
		From:    m.auth.From,
		Signer:  m.auth.Signer,
		Nonce:   big.NewInt(int64(m.nonce)),
		Context: ctx,
	}
	m.nonce++
	m.nonceMu.Unlock()

	tx, err := fn(opts)
	if err != nil {
		m.syncNonce(ctx)
		return nil, fmt.Errorf("%s: tx failed: %w", label, err)
	}

	slog.Info("mint tx submitted", "op", label, "tx", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, m.client, tx)
	if err != nil {
		return nil, fmt.Errorf("%s: waiting for tx %s: %w", label, tx.Hash().Hex(), err)
	}
	if receipt.Status == 0 {
		return nil, fmt.Errorf("%s: tx %s reverted", label, tx.Hash().Hex())
	}

	slog.Info("mint confirmed", "op", label, "tx", tx.Hash().Hex(), "block", receipt.BlockNumber)
	return receipt, nil
}

func (m *Minter) syncNonce(ctx context.Context) {
	nonce, err := m.client.PendingNonceAt(ctx, m.auth.From)
	if err != nil {
		slog.Error("failed to sync nonce", "err", err)
		return
	}
	m.nonceMu.Lock()
	m.nonce = nonce
	m.nonceMu.Unlock()
}
