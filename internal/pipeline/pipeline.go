package pipeline

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"ogcr-tokenizer/internal/eth"
	"ogcr-tokenizer/internal/obp"
)

type Pipeline struct {
	obp    *obp.Client
	minter *eth.Minter

	// Fast-path in-memory dedup per token type. Authoritative check is on-chain.
	seenParcels      sync.Map
	seenActivities   sync.Map
	seenCertificates sync.Map
	seenMonitoring   sync.Map
}

func New(obpClient *obp.Client, minter *eth.Minter) *Pipeline {
	return &Pipeline{obp: obpClient, minter: minter}
}

func (p *Pipeline) Run(ctx context.Context, interval time.Duration) {
	slog.Info("pipeline started", "poll_interval", interval)
	for {
		p.pollParcels(ctx)
		p.pollActivities(ctx)
		p.pollCertificates(ctx)
		p.pollMonitoring(ctx)
		select {
		case <-ctx.Done():
			slog.Info("pipeline stopped")
			return
		case <-time.After(interval):
		}
	}
}

// ---- ParcelNFT ----

func (p *Pipeline) pollParcels(ctx context.Context) {
	parcels, err := p.obp.GetVerifiedParcels()
	if err != nil {
		slog.Error("GetVerifiedParcels failed", "err", err)
		return
	}
	slog.Info("parcel poll complete", "verified_parcels", len(parcels))

	for _, pov := range parcels {
		if _, loaded := p.seenParcels.LoadOrStore(pov.ParcelID, struct{}{}); loaded {
			continue
		}
		go p.processParcel(ctx, pov.ParcelID)
	}
}

func (p *Pipeline) processParcel(ctx context.Context, parcelID string) {
	log := slog.With("parcel_id", parcelID)

	minted, err := p.minter.ParcelAlreadyMinted(ctx, parcelID)
	if err != nil {
		log.Error("on-chain dedup check failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	if minted {
		log.Info("ParcelNFT already minted, skipping")
		return
	}

	// Raw parcel entity → parcel_uri + integrity hash
	parcelBody, err := p.obp.GetParcelRaw(parcelID)
	if err != nil {
		log.Error("GetParcelRaw failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	parcelURI := p.obp.ParcelURL(parcelID)
	parcelHash := crypto.Keccak256Hash(parcelBody).Hex()

	// activity_parcel_verification → activity → operator (for recipient wallet)
	apv, err := p.obp.GetActivityParcelVerificationByParcelID(parcelID)
	if err != nil {
		log.Error("GetActivityParcelVerification failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	activity, err := p.obp.GetActivity(apv.ActivityID)
	if err != nil {
		log.Error("GetActivity failed", "activity_id", apv.ActivityID, "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	operator, err := p.obp.GetOperator(activity.OperatorID)
	if err != nil {
		log.Error("GetOperator failed", "operator_id", activity.OperatorID, "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	if operator.OGCRWalletAddress == "" {
		log.Error("operator has no ogcr_wallet_address")
		p.seenParcels.Delete(parcelID)
		return
	}

	log.Info("minting ParcelNFT", "to", operator.OGCRWalletAddress, "parcel_uri", parcelURI, "parcel_hash", parcelHash)

	if err := p.minter.MintParcel(ctx, eth.ParcelParams{
		To:         common.HexToAddress(operator.OGCRWalletAddress),
		ParcelID:   parcelID,
		ParcelURI:  parcelURI,
		ParcelHash: parcelHash,
	}); err != nil {
		log.Error("mint failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}

	log.Info("ParcelNFT minted successfully")
}

// ---- ActivityNFT ----

func (p *Pipeline) pollActivities(ctx context.Context) {
	avs, err := p.obp.GetVerifiedActivityVerifications()
	if err != nil {
		slog.Error("GetVerifiedActivityVerifications failed", "err", err)
		return
	}
	slog.Info("activity poll complete", "verified_activities", len(avs))

	for _, av := range avs {
		if _, loaded := p.seenActivities.LoadOrStore(av.ActivityID, struct{}{}); loaded {
			continue
		}
		go p.processActivity(ctx, av.ActivityID)
	}
}

func (p *Pipeline) processActivity(ctx context.Context, activityID string) {
	log := slog.With("activity_id", activityID)

	minted, err := p.minter.ActivityAlreadyMinted(ctx, activityID)
	if err != nil {
		log.Error("on-chain dedup check failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}
	if minted {
		log.Info("ActivityNFT already minted, skipping")
		return
	}

	// Raw activity entity → decoded fields + activity_uri + integrity hash
	activity, body, err := p.obp.GetActivityWithRaw(activityID)
	if err != nil {
		log.Error("GetActivityWithRaw failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}
	activityURI := p.obp.ActivityURL(activityID)
	activityHash := crypto.Keccak256Hash(body).Hex()

	operator, err := p.obp.GetOperator(activity.OperatorID)
	if err != nil {
		log.Error("GetOperator failed", "operator_id", activity.OperatorID, "err", err)
		p.seenActivities.Delete(activityID)
		return
	}
	if operator.OGCRWalletAddress == "" {
		log.Error("operator has no ogcr_wallet_address")
		p.seenActivities.Delete(activityID)
		return
	}

	log.Info("minting ActivityNFT", "to", operator.OGCRWalletAddress, "activity_uri", activityURI, "activity_hash", activityHash)

	if err := p.minter.MintActivity(ctx, eth.ActivityParams{
		To:           common.HexToAddress(operator.OGCRWalletAddress),
		ActivityID:   activityID,
		OperatorID:   activity.OperatorID,
		Name:         activity.Name,
		ActivityType: activity.Type,
		StartDate:    activity.StartDate,
		EndDate:      activity.EndDate,
		ActivityURI:  activityURI,
		ActivityHash: activityHash,
	}); err != nil {
		log.Error("mint failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}

	log.Info("ActivityNFT minted successfully")
}

// ---- CertificationNFT ----

func (p *Pipeline) pollCertificates(ctx context.Context) {
	certs, err := p.obp.GetCertificates()
	if err != nil {
		slog.Error("GetCertificates failed", "err", err)
		return
	}
	slog.Info("certificate poll complete", "certificates", len(certs))

	for _, cert := range certs {
		if _, loaded := p.seenCertificates.LoadOrStore(cert.CertificateOfComplianceID, struct{}{}); loaded {
			continue
		}
		go p.processCertificate(ctx, cert)
	}
}

func (p *Pipeline) processCertificate(ctx context.Context, cert obp.Certificate) {
	complianceID := cert.CertificateOfComplianceID
	log := slog.With("certificate_of_compliance_id", complianceID, "activity_id", cert.ActivityID)

	minted, err := p.minter.CertificationAlreadyMinted(ctx, complianceID)
	if err != nil {
		log.Error("on-chain dedup check failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	if minted {
		log.Info("CertificationNFT already minted, skipping")
		return
	}

	// The certificate references an activity; the CertificationNFT stores that
	// ActivityNFT's token id. Defer until the ActivityNFT has been minted.
	activityTokenID, err := p.minter.ActivityTokenID(ctx, cert.ActivityID)
	if err != nil {
		log.Error("ActivityTokenID lookup failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	if activityTokenID.Sign() == 0 {
		log.Info("ActivityNFT not yet minted for this certificate, deferring")
		p.seenCertificates.Delete(complianceID)
		return
	}

	// Linked activity → activity_url + hash (for the cert) and operator wallet (recipient).
	activity, activityBody, err := p.obp.GetActivityWithRaw(cert.ActivityID)
	if err != nil {
		log.Error("GetActivityWithRaw failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	activityURI := p.obp.ActivityURL(cert.ActivityID)
	activityHash := crypto.Keccak256Hash(activityBody).Hex()

	operator, err := p.obp.GetOperator(activity.OperatorID)
	if err != nil {
		log.Error("GetOperator failed", "operator_id", activity.OperatorID, "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	if operator.OGCRWalletAddress == "" {
		log.Error("operator has no ogcr_wallet_address")
		p.seenCertificates.Delete(complianceID)
		return
	}

	// Certificate entity → certification_url + hash.
	certBody, err := p.obp.GetCertificateRaw(complianceID)
	if err != nil {
		log.Error("GetCertificateRaw failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	certificationURI := p.obp.CertificateURL(complianceID)
	certificationHash := crypto.Keccak256Hash(certBody).Hex()

	log.Info("minting CertificationNFT",
		"to", operator.OGCRWalletAddress,
		"activity_nft_id", activityTokenID.String(),
		"certification_url", certificationURI,
	)

	if err := p.minter.MintCertification(ctx, eth.CertificationParams{
		To:                          common.HexToAddress(operator.OGCRWalletAddress),
		ActivityNftID:               activityTokenID,
		CertificationSchemeID:       cert.CertificationSchemeID,
		CertificationBodyID:         cert.CertificationBodyID,
		CertificationOfComplianceID: complianceID,
		IssueDate:                   cert.IssueDate,
		ExpiryDate:                  cert.ExpiryDate,
		CertificationStatus:         cert.CertificationStatus,
		ActivityURI:                 activityURI,
		ActivityHash:                activityHash,
		CertificationURI:            certificationURI,
		CertificationHash:           certificationHash,
	}); err != nil {
		log.Error("mint failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}

	log.Info("CertificationNFT minted successfully")
}

// ---- CarbonCreditBatchNFT (ERC-6551) + CarbonCredit (ERC-20) ----

func (p *Pipeline) pollMonitoring(ctx context.Context) {
	mvs, err := p.obp.GetVerifiedMonitoringVerifications()
	if err != nil {
		slog.Error("GetVerifiedMonitoringVerifications failed", "err", err)
		return
	}
	slog.Info("monitoring poll complete", "verified_monitoring", len(mvs))

	for _, mv := range mvs {
		if _, loaded := p.seenMonitoring.LoadOrStore(mv.ActivityMonitoringPeriodVerificationID, struct{}{}); loaded {
			continue
		}
		go p.processMonitoring(ctx, mv)
	}
}

// processMonitoring mints one CarbonCreditBatchNFT per positive benefit amount on
// the verification, then mints that amount of CarbonCredit ERC-20 into the batch's TBA.
func (p *Pipeline) processMonitoring(ctx context.Context, mv obp.MonitoringVerification) {
	vID := mv.ActivityMonitoringPeriodVerificationID
	log := slog.With("monitoring_verification_id", vID, "activity_id", mv.ActivityID)

	amounts := mv.CreditAmounts()
	if len(amounts) == 0 {
		log.Info("no positive benefit amounts, skipping")
		return
	}

	// The batch stores the ActivityNFT token id; defer until it's minted.
	activityTokenID, err := p.minter.ActivityTokenID(ctx, mv.ActivityID)
	if err != nil {
		log.Error("ActivityTokenID lookup failed", "err", err)
		p.seenMonitoring.Delete(vID)
		return
	}
	if activityTokenID.Sign() == 0 {
		log.Info("ActivityNFT not yet minted, deferring")
		p.seenMonitoring.Delete(vID)
		return
	}

	// Assemble the shared url/hash reference set (activity, operator, certificate,
	// scheme, body). Any fetch failure defers the whole verification for retry.
	refs, operator, err := p.gatherBatchRefs(ctx, mv.ActivityID)
	if err != nil {
		log.Error("gathering batch references failed", "err", err)
		p.seenMonitoring.Delete(vID)
		return
	}
	if operator.OGCRWalletAddress == "" {
		log.Error("operator has no ogcr_wallet_address")
		p.seenMonitoring.Delete(vID)
		return
	}
	to := common.HexToAddress(operator.OGCRWalletAddress)

	failed := false
	for _, ca := range amounts {
		clog := log.With("credit_type", ca.CreditType, "amount", ca.Amount)
		amount := big.NewInt(ca.Amount)

		// Mint the batch if it doesn't exist yet.
		tokenID, tba, found, err := p.minter.FindBatch(ctx, activityTokenID, ca.CreditType)
		if err != nil {
			clog.Error("FindBatch failed", "err", err)
			failed = true
			continue
		}
		if !found {
			clog.Info("minting CarbonCreditBatchNFT")
			tokenID, tba, err = p.minter.MintBatch(ctx, refs.batchParams(to, activityTokenID, ca.CreditType))
			if err != nil {
				clog.Error("batch mint failed", "err", err)
				failed = true
				continue
			}
		}

		// Ensure the batch's TBA holds exactly `amount` credits. Batch and credit
		// are two txs; minting the shortfall (not a flat amount) recovers a batch
		// minted without its credit and is idempotent on re-run — it converges to
		// `amount` and never over-mints.
		bal, err := p.minter.CreditBalance(ctx, tba)
		if err != nil {
			clog.Error("CreditBalance check failed", "err", err)
			failed = true
			continue
		}
		shortfall := new(big.Int).Sub(amount, bal)
		if shortfall.Sign() <= 0 {
			clog.Info("batch + credits already present, skipping", "batch_token", tokenID.String(), "balance", bal.String())
			continue
		}
		if err := p.minter.MintCredit(ctx, tba, shortfall); err != nil {
			clog.Error("credit mint failed", "batch_token", tokenID.String(), "tba", tba.Hex(), "err", err)
			failed = true
			continue
		}
		clog.Info("CarbonCreditBatch + credits minted", "batch_token", tokenID.String(), "tba", tba.Hex(), "minted", shortfall.String())
	}

	if failed {
		// Retry next poll; already-minted batches are skipped by BatchExists.
		p.seenMonitoring.Delete(vID)
	}
}

// batchRefs holds the resolved url/hash reference set shared by every batch of a
// monitoring verification (they differ only by credit_type + amount).
type batchRefs struct {
	activityURI, activityHash string
	operatorURI, operatorHash string
	certURI, certHash         string
	schemeURI, schemeHash     string
	bodyURI, bodyHash         string
}

func (r batchRefs) batchParams(to common.Address, activityNftID *big.Int, creditType string) eth.BatchParams {
	return eth.BatchParams{
		To:                      to,
		ActivityNftID:           activityNftID,
		CreditType:              creditType,
		ActivityURI:             r.activityURI,
		ActivityHash:            r.activityHash,
		OperatorURI:             r.operatorURI,
		OperatorHash:            r.operatorHash,
		CertificationURI:        r.certURI,
		CertificationHash:       r.certHash,
		CertificationSchemeURI:  r.schemeURI,
		CertificationSchemeHash: r.schemeHash,
		CertificationBodyURI:    r.bodyURI,
		CertificationBodyHash:   r.bodyHash,
	}
}

// gatherBatchRefs fetches every entity referenced by a batch and computes its
// url + keccak256 hash. Also returns the operator (for the recipient wallet).
func (p *Pipeline) gatherBatchRefs(ctx context.Context, activityID string) (batchRefs, *obp.Operator, error) {
	var r batchRefs

	activity, activityBody, err := p.obp.GetActivityWithRaw(activityID)
	if err != nil {
		return r, nil, fmt.Errorf("activity: %w", err)
	}
	r.activityURI = p.obp.ActivityURL(activityID)
	r.activityHash = crypto.Keccak256Hash(activityBody).Hex()

	operator, operatorBody, err := p.obp.GetOperatorWithRaw(activity.OperatorID)
	if err != nil {
		return r, nil, fmt.Errorf("operator: %w", err)
	}
	r.operatorURI = p.obp.OperatorURL(activity.OperatorID)
	r.operatorHash = crypto.Keccak256Hash(operatorBody).Hex()

	cert, err := p.obp.GetCertificateByActivityID(activityID)
	if err != nil {
		return r, nil, fmt.Errorf("certificate: %w", err)
	}
	certBody, err := p.obp.GetCertificateRaw(cert.CertificateOfComplianceID)
	if err != nil {
		return r, nil, fmt.Errorf("certificate raw: %w", err)
	}
	r.certURI = p.obp.CertificateURL(cert.CertificateOfComplianceID)
	r.certHash = crypto.Keccak256Hash(certBody).Hex()

	schemeBody, err := p.obp.GetCertificationSchemeRaw(cert.CertificationSchemeID)
	if err != nil {
		return r, nil, fmt.Errorf("certification_scheme: %w", err)
	}
	r.schemeURI = p.obp.CertificationSchemeURL(cert.CertificationSchemeID)
	r.schemeHash = crypto.Keccak256Hash(schemeBody).Hex()

	bodyBody, err := p.obp.GetCertificationBodyRaw(cert.CertificationBodyID)
	if err != nil {
		return r, nil, fmt.Errorf("certification_body: %w", err)
	}
	r.bodyURI = p.obp.CertificationBodyURL(cert.CertificationBodyID)
	r.bodyHash = crypto.Keccak256Hash(bodyBody).Hex()

	return r, operator, nil
}
