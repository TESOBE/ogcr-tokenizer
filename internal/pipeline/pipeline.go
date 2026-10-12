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

	"ogcr-tokenizer/internal/contract"
	"ogcr-tokenizer/internal/eth"
	"ogcr-tokenizer/internal/health"
	"ogcr-tokenizer/internal/obp"
)

type Pipeline struct {
	obp    *obp.Client
	minter *eth.Minter

	// Optional: records what the pipeline is doing for the health endpoint. Nil
	// when no health server is configured, so every call site guards on it.
	health *health.State

	// Fast-path in-memory dedup per token type. Authoritative check is on-chain.
	seenParcels      sync.Map
	seenActivities   sync.Map
	seenCertificates sync.Map
	seenMonitoring   sync.Map
}

func New(obpClient *obp.Client, minter *eth.Minter) *Pipeline {
	return &Pipeline{obp: obpClient, minter: minter}
}

// WithHealth attaches a health state for the pipeline to report into.
func (p *Pipeline) WithHealth(h *health.State) *Pipeline {
	p.health = h
	return p
}

func (p *Pipeline) recordPoll(name string, found int, err error) {
	if p.health != nil {
		p.health.RecordPoll(name, found, err)
	}
}

func (p *Pipeline) recordMinted() {
	if p.health != nil {
		p.health.Minted()
	}
}

func (p *Pipeline) recordFailed() {
	if p.health != nil {
		p.health.Failed()
	}
}

func (p *Pipeline) recordSkipped() {
	if p.health != nil {
		p.health.Skipped()
	}
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

// ---- shared lookups ----

// hashOf is the on-chain integrity hash of a registry record: keccak256 of the
// exact response body it was read from.
func hashOf(body []byte) string { return crypto.Keccak256Hash(body).Hex() }

// recipient returns the operator's wallet, the address every token of an
// activity is minted to. It refuses anything that is not a real address rather
// than minting to whatever a malformed value happens to decode to.
func recipient(operator *obp.Operator) (common.Address, error) {
	addr := operator.OGCRWalletAddress
	if !common.IsHexAddress(addr) || common.HexToAddress(addr) == (common.Address{}) {
		return common.Address{}, fmt.Errorf("operator %s has no valid ogcr_wallet_address", operator.OperatorID)
	}
	return common.HexToAddress(addr), nil
}

// activityRefs is an activity with its operator and their url/hash references.
type activityRefs struct {
	activity                  *obp.Activity
	activityURL, activityHash string
	operator                  *obp.Operator
	operatorURL, operatorHash string
}

func (p *Pipeline) gatherActivityRefs(activityID string) (activityRefs, error) {
	var r activityRefs

	activity, activityBody, err := p.obp.GetActivityWithRaw(activityID)
	if err != nil {
		return r, fmt.Errorf("activity: %w", err)
	}
	r.activity = activity
	r.activityURL = p.obp.ActivityURL(activityID)
	r.activityHash = hashOf(activityBody)

	operator, operatorBody, err := p.obp.GetOperatorWithRaw(activity.OperatorID)
	if err != nil {
		return r, fmt.Errorf("operator %s: %w", activity.OperatorID, err)
	}
	r.operator = operator
	r.operatorURL = p.obp.OperatorURL(activity.OperatorID)
	r.operatorHash = hashOf(operatorBody)

	return r, nil
}

// certificateRefs is a certificate of compliance with its scheme and body and
// their url/hash references. Scheme and body stay empty when the certificate
// does not name one.
type certificateRefs struct {
	cert                              *obp.Certificate
	certURL, certHash                 string
	schemeName, schemeURL, schemeHash string
	bodyName, bodyURL, bodyHash       string
}

func (p *Pipeline) gatherCertificateRefs(certificateID string) (certificateRefs, error) {
	var r certificateRefs

	cert, certBody, err := p.obp.GetCertificateWithRaw(certificateID)
	if err != nil {
		return r, fmt.Errorf("certificate: %w", err)
	}
	r.cert = cert
	r.certURL = p.obp.CertificateURL(certificateID)
	r.certHash = hashOf(certBody)

	if id := cert.CertificationSchemeID; id != "" {
		scheme, schemeBody, err := p.obp.GetCertificationSchemeWithRaw(id)
		if err != nil {
			return r, fmt.Errorf("certification_scheme %s: %w", id, err)
		}
		r.schemeName = scheme.Name
		r.schemeURL = p.obp.CertificationSchemeURL(id)
		r.schemeHash = hashOf(schemeBody)
	}

	if id := cert.CertificationBodyID; id != "" {
		body, bodyBody, err := p.obp.GetCertificationBodyWithRaw(id)
		if err != nil {
			return r, fmt.Errorf("certification_body %s: %w", id, err)
		}
		r.bodyName = body.LegalName
		r.bodyURL = p.obp.CertificationBodyURL(id)
		r.bodyHash = hashOf(bodyBody)
	}

	return r, nil
}

// ---- ParcelNFT ----

func (p *Pipeline) pollParcels(ctx context.Context) {
	parcels, err := p.obp.GetVerifiedParcels()
	p.recordPoll("parcels", len(parcels), err)
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
		p.recordSkipped()
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
	parcelHash := hashOf(parcelBody)

	// parcel_activity → activity → operator (for recipient wallet)
	link, err := p.obp.GetParcelActivityByParcelID(parcelID)
	if err != nil {
		log.Error("GetParcelActivity failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	activity, err := p.obp.GetActivity(link.ActivityID)
	if err != nil {
		log.Error("GetActivity failed", "activity_id", link.ActivityID, "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	operator, err := p.obp.GetOperator(activity.OperatorID)
	if err != nil {
		log.Error("GetOperator failed", "operator_id", activity.OperatorID, "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}
	to, err := recipient(operator)
	if err != nil {
		log.Error("no recipient", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}

	log.Info("minting ParcelNFT", "to", to.Hex(), "parcel_uri", parcelURI, "parcel_hash", parcelHash)

	if err := p.minter.MintParcel(ctx, eth.ParcelParams{
		To:         to,
		ParcelID:   parcelID,
		ParcelURI:  parcelURI,
		ParcelHash: parcelHash,
	}); err != nil {
		p.recordFailed()
		log.Error("mint failed", "err", err)
		p.seenParcels.Delete(parcelID)
		return
	}

	p.recordMinted()
	log.Info("ParcelNFT minted successfully")
}

// ---- ActivityNFT ----

// pollActivities finds activities that are ready to mint: a verified
// activity_verification AND at least one certificate_of_compliance.
func (p *Pipeline) pollActivities(ctx context.Context) {
	avs, err := p.obp.GetVerifiedActivityVerifications()
	if err != nil {
		p.recordPoll("activities", 0, err)
		slog.Error("GetVerifiedActivityVerifications failed", "err", err)
		return
	}
	certs, err := p.obp.GetCertificates()
	if err != nil {
		p.recordPoll("activities", 0, err)
		slog.Error("GetCertificates failed", "err", err)
		return
	}
	certified := map[string]bool{}
	for _, cert := range certs {
		certified[cert.ActivityID] = true
	}

	// An activity can carry several verified verifications; mint it once.
	ready := map[string]bool{}
	for _, av := range avs {
		if av.ActivityID != "" && certified[av.ActivityID] {
			ready[av.ActivityID] = true
		}
	}
	p.recordPoll("activities", len(ready), nil)
	slog.Info("activity poll complete", "verified_activities", len(avs), "verified_and_certified", len(ready))

	for activityID := range ready {
		if _, loaded := p.seenActivities.LoadOrStore(activityID, struct{}{}); loaded {
			continue
		}
		go p.processActivity(ctx, activityID)
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
		p.recordSkipped()
		log.Info("ActivityNFT already minted, skipping")
		return
	}

	refs, err := p.gatherActivityRefs(activityID)
	if err != nil {
		log.Error("gathering activity references failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}
	to, err := recipient(refs.operator)
	if err != nil {
		log.Error("no recipient", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}

	parcelIDs, err := p.obp.GetParcelIDsByActivityID(activityID)
	if err != nil {
		log.Error("GetParcelIDsByActivityID failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}

	activity := refs.activity
	log.Info("minting ActivityNFT", "to", to.Hex(), "activity_url", refs.activityURL,
		"activity_hash", refs.activityHash, "parcels", len(parcelIDs))

	if err := p.minter.MintActivity(ctx, eth.ActivityParams{
		To: to,
		Data: contract.ActivityNFTActivityData{
			ActivityId:            activityID,
			OperatorId:            activity.OperatorID,
			OperatorName:          refs.operator.LegalName,
			Name:                  activity.Name,
			Summary:               activity.Summary,
			Website:               activity.Website,
			ActivityType:          activity.ActivityType,
			UnitTypes:             activity.UnitTypes,
			City:                  activity.City,
			CountryId:             activity.CountryID,
			CertificationSchemeId: activity.CertificationSchemeID,
			ParcelIds:             parcelIDs,
			StartDate:             activity.StartDate,
			EndDate:               activity.EndDate,
			ActivityUrl:           refs.activityURL,
			ActivityHash:          refs.activityHash,
		},
	}); err != nil {
		p.recordFailed()
		log.Error("mint failed", "err", err)
		p.seenActivities.Delete(activityID)
		return
	}

	p.recordMinted()
	log.Info("ActivityNFT minted successfully")
}

// ---- CertificationNFT ----

// pollCertificates mints one CertificationNFT per certificate_of_compliance. A
// re-certification is a new certificate record for the same activity, so it
// arrives here as just another certificate.
func (p *Pipeline) pollCertificates(ctx context.Context) {
	certs, err := p.obp.GetCertificates()
	p.recordPoll("certificates", len(certs), err)
	if err != nil {
		slog.Error("GetCertificates failed", "err", err)
		return
	}
	slog.Info("certificate poll complete", "certificates", len(certs))

	for _, cert := range certs {
		if _, loaded := p.seenCertificates.LoadOrStore(cert.CertificateOfComplianceID, struct{}{}); loaded {
			continue
		}
		go p.processCertificate(ctx, cert.CertificateOfComplianceID, cert.ActivityID)
	}
}

func (p *Pipeline) processCertificate(ctx context.Context, complianceID, activityID string) {
	log := slog.With("certificate_of_compliance_id", complianceID, "activity_id", activityID)

	minted, err := p.minter.CertificationAlreadyMinted(ctx, complianceID)
	if err != nil {
		log.Error("on-chain dedup check failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	if minted {
		p.recordSkipped()
		log.Info("CertificationNFT already minted, skipping")
		return
	}

	// The certificate references an activity; the CertificationNFT stores that
	// ActivityNFT's token id. Defer until the ActivityNFT has been minted.
	activityTokenID, err := p.minter.ActivityTokenID(ctx, activityID)
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

	aRefs, err := p.gatherActivityRefs(activityID)
	if err != nil {
		log.Error("gathering activity references failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	to, err := recipient(aRefs.operator)
	if err != nil {
		log.Error("no recipient", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}

	cRefs, err := p.gatherCertificateRefs(complianceID)
	if err != nil {
		log.Error("gathering certificate references failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}
	cert := cRefs.cert

	// Audit report → audit_report_uri + hash, when the certificate names one.
	var auditURI, auditHash string
	if cert.AuditReportID != "" {
		auditBody, err := p.obp.GetAuditReportRaw(cert.AuditReportID)
		if err != nil {
			log.Error("GetAuditReportRaw failed", "audit_report_id", cert.AuditReportID, "err", err)
			p.seenCertificates.Delete(complianceID)
			return
		}
		auditURI = p.obp.AuditReportURL(cert.AuditReportID)
		auditHash = hashOf(auditBody)
	}

	log.Info("minting CertificationNFT",
		"to", to.Hex(),
		"activity_nft_id", activityTokenID.String(),
		"certificate_of_compliance_uri", cRefs.certURL,
	)

	if err := p.minter.MintCertification(ctx, eth.CertificationParams{
		To: to,
		Data: contract.CertificationNFTCertificationData{
			CertificateOfComplianceId:   complianceID,
			CertificateOfComplianceUri:  cRefs.certURL,
			CertificateOfComplianceHash: cRefs.certHash,
			ActivityId:                  activityID,
			ActivityNftId:               activityTokenID,
			ActivityUri:                 aRefs.activityURL,
			ActivityHash:                aRefs.activityHash,
			ActivityName:                aRefs.activity.Name,
			CertificationSchemeName:     cRefs.schemeName,
			CertificationSchemeUri:      cRefs.schemeURL,
			CertificationSchemeHash:     cRefs.schemeHash,
			CertificationBodyName:       cRefs.bodyName,
			CertificationBodyUri:        cRefs.bodyURL,
			CertificationBodyHash:       cRefs.bodyHash,
			CountryId:                   cert.CountryID,
			IssueDate:                   cert.IssueDate,
			ExpiryDate:                  cert.ExpiryDate,
			AuditReportUri:              auditURI,
			AuditReportHash:             auditHash,
			// Quantities go on-chain exactly as recorded; an absent one stays empty.
			CarbonRemovalsUnderBaseline:                   cert.CarbonRemovalsUnderBaseline.String(),
			SoilEmissionsUnderBaseline:                    cert.SoilEmissionsUnderBaseline.String(),
			PermanentNetCarbonRemovalBenefit:              cert.PermanentNetCarbonRemovalBenefit.String(),
			CarbonFarmingTemporaryNetCarbonRemovalBenefit: cert.CarbonFarmingTemporaryNetCarbonRemovalBenefit.String(),
			CarbonFarmingNetSoilEmissionReductionBenefit:  cert.CarbonFarmingNetSoilEmissionReductionBenefit.String(),
			CarbonStorageTemporaryNetCarbonRemovalBenefit: cert.CarbonStorageTemporaryNetCarbonRemovalBenefit.String(),
			UnitTypes:           cert.UnitTypes,
			CertificationStatus: cert.CertificationStatus,
		},
	}); err != nil {
		p.recordFailed()
		log.Error("mint failed", "err", err)
		p.seenCertificates.Delete(complianceID)
		return
	}

	p.recordMinted()
	log.Info("CertificationNFT minted successfully")
}

// ---- CarbonCreditBatchNFT (ERC-6551) + CarbonCredit (ERC-20) ----

func (p *Pipeline) pollMonitoring(ctx context.Context) {
	mvs, err := p.obp.GetVerifiedMonitoringVerifications()
	p.recordPoll("monitoring_periods", len(mvs), err)
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

// processMonitoring mints one CarbonCreditBatchNFT per benefit field that
// carries a value on the verification, then mints that amount of CarbonCredit
// ERC-20 into the batch's TBA.
func (p *Pipeline) processMonitoring(ctx context.Context, mv obp.MonitoringVerification) {
	vID := mv.ActivityMonitoringPeriodVerificationID
	log := slog.With("monitoring_verification_id", vID, "activity_id", mv.ActivityID)

	batches, err := mv.CreditBatches()
	if err != nil {
		// A malformed amount will not fix itself on retry within this run; it
		// is picked up again after a restart once the record is corrected.
		p.recordFailed()
		log.Error("invalid benefit amount, not minting any batch for this verification", "err", err)
		return
	}
	if len(batches) == 0 {
		log.Info("no benefit amounts, skipping")
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

	// Assemble the reference set shared by every batch of this verification
	// (they differ only by unit type + amount). Any fetch failure defers the
	// whole verification for retry.
	aRefs, cRefs, err := p.gatherBatchRefs(mv)
	if err != nil {
		log.Error("gathering batch references failed", "err", err)
		p.seenMonitoring.Delete(vID)
		return
	}
	to, err := recipient(aRefs.operator)
	if err != nil {
		log.Error("no recipient", "err", err)
		p.seenMonitoring.Delete(vID)
		return
	}

	failed := false
	for _, b := range batches {
		clog := log.With("carbon_credit_batch_id", b.BatchID, "unit_type", b.UnitType, "amount", b.Amount.String())

		// Mint the batch if it doesn't exist yet.
		tokenID, tba, issued, found, err := p.minter.FindBatch(ctx, b.BatchID)
		if err != nil {
			clog.Error("FindBatch failed", "err", err)
			failed = true
			continue
		}
		if !found {
			issued, err = obp.ToBaseUnits(b.Amount, 18)
			if err != nil {
				clog.Error("invalid amount", "err", err)
				failed = true
				continue
			}
			clog.Info("minting CarbonCreditBatchNFT")
			tokenID, tba, err = p.minter.MintBatch(ctx, batchParams(to, activityTokenID, b, issued, aRefs, cRefs))
			if err != nil {
				p.recordFailed()
				clog.Error("batch mint failed", "err", err)
				failed = true
				continue
			}
			p.recordMinted()
		}

		// Fund the batch's TBA with the amount fixed on the batch at mint. Batch
		// and credit are two txs, so a batch can exist without its credits.
		// What is still owed is measured against the cumulative amount minted to
		// the TBA, not its balance: the balance drops when the owner moves or
		// retires credits, and topping that up would issue them twice.
		minted, err := p.minter.CreditMintedTo(ctx, tba)
		if err != nil {
			clog.Error("CreditMintedTo check failed", "err", err)
			failed = true
			continue
		}
		shortfall := new(big.Int).Sub(issued, minted)
		if shortfall.Sign() <= 0 {
			p.recordSkipped()
			clog.Info("batch + credits already present, skipping", "batch_token", tokenID.String(), "credits_minted", minted.String())
			continue
		}
		if err := p.minter.MintCredit(ctx, tba, shortfall); err != nil {
			p.recordFailed()
			clog.Error("credit mint failed", "batch_token", tokenID.String(), "tba", tba.Hex(), "err", err)
			failed = true
			continue
		}
		clog.Info("CarbonCreditBatch + credits minted", "batch_token", tokenID.String(), "tba", tba.Hex(), "minted", shortfall.String())
	}

	if failed {
		// Retry next poll; batches and credits already minted are skipped.
		p.seenMonitoring.Delete(vID)
	}
}

// gatherBatchRefs resolves the activity, operator, certificate, scheme and body
// a monitoring verification's batches reference. The certificate is the one
// the verification names; an older record without one falls back to the
// activity's most recently issued certificate.
func (p *Pipeline) gatherBatchRefs(mv obp.MonitoringVerification) (activityRefs, certificateRefs, error) {
	aRefs, err := p.gatherActivityRefs(mv.ActivityID)
	if err != nil {
		return aRefs, certificateRefs{}, err
	}

	certificateID := mv.CertificateOfComplianceID
	if certificateID == "" {
		cert, err := p.obp.GetLatestCertificateByActivityID(mv.ActivityID)
		if err != nil {
			return aRefs, certificateRefs{}, fmt.Errorf("certificate: %w", err)
		}
		certificateID = cert.CertificateOfComplianceID
	}

	cRefs, err := p.gatherCertificateRefs(certificateID)
	if err != nil {
		return aRefs, cRefs, err
	}
	return aRefs, cRefs, nil
}

// batchStatusIssued is the status_code a batch carries at mint.
const batchStatusIssued = "issued"

func batchParams(
	to common.Address,
	activityNftID *big.Int,
	b obp.CreditBatch,
	amount *big.Int,
	a activityRefs,
	c certificateRefs,
) eth.BatchParams {
	return eth.BatchParams{
		To: to,
		Data: contract.CarbonCreditBatchNFTBatchData{
			CarbonCreditBatchId:         b.BatchID,
			BatchName:                   b.BatchName,
			ActivityId:                  a.activity.ActivityID,
			ActivityNftId:               activityNftID,
			ActivityName:                a.activity.Name,
			ActivityUrl:                 a.activityURL,
			ActivityHash:                a.activityHash,
			CountryId:                   a.activity.CountryID,
			OperatorUrl:                 a.operatorURL,
			OperatorHash:                a.operatorHash,
			CertificateOfComplianceUri:  c.certURL,
			CertificateOfComplianceHash: c.certHash,
			UnitType:                    b.UnitType,
			MonitoringPeriodStartDate:   a.activity.MonitoringPeriodStartDate,
			MonitoringPeriodEndDate:     a.activity.MonitoringPeriodEndDate,
			Amount:                      amount,
			StatusCode:                  batchStatusIssued,
			IssueDate:                   c.cert.IssueDate,
			ExpiryDate:                  c.cert.ExpiryDate,
			CertificationSchemeUrl:      c.schemeURL,
			CertificationSchemeHash:     c.schemeHash,
			CertificationBodyUrl:        c.bodyURL,
			CertificationBodyHash:       c.bodyHash,
		},
	}
}
