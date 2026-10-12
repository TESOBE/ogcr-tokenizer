package obp

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type ParcelOwnerVerification struct {
	ParcelOwnerVerificationID string `json:"parcel_owner_verification_id"`
	ParcelID                  string `json:"parcel_id"`
	StatusCode                string `json:"status_code"`
	StatusMessage             string `json:"status_message"`
}

// ParcelActivity is the intersecting entity linking parcels to activities.
type ParcelActivity struct {
	ParcelID   string `json:"parcel_id"`
	ActivityID string `json:"activity_id"`
}

type ActivityVerification struct {
	ActivityVerificationID string `json:"activity_verification_id"`
	ActivityID             string `json:"activity_id"`
	StatusCode             string `json:"status_code"`
	StatusMessage          string `json:"status_message"`
}

type Activity struct {
	ActivityID                string `json:"activity_id"`
	Name                      string `json:"name"`
	Summary                   string `json:"summary"`
	Description               string `json:"description"`
	Website                   string `json:"website"`
	CountryID                 string `json:"country_id"`
	City                      string `json:"city"`
	StartDate                 string `json:"start_date"`
	EndDate                   string `json:"end_date"`
	ActivityType              string `json:"activity_type"`
	UnitTypes                 string `json:"unit_types"`
	CertificationSchemeID     string `json:"certification_scheme_id"`
	MonitoringPeriodStartDate string `json:"monitoring_period_start_date"`
	MonitoringPeriodEndDate   string `json:"monitoring_period_end_date"`
	OperatorID                string `json:"operator_id"`
}

// Unit types as the registry names them (activity.unit_types,
// carbon_credit_batch.unit_type).
const (
	UnitTypePermanentRemoval      = "Permanent Removal"
	UnitTypeCarbonFarming         = "Carbon Farming Sequestration"
	UnitTypeSoilEmissionReduction = "Soil Emission Reduction"
	UnitTypeCarbonStorageProduct  = "Carbon Storage in Product"
)

// Quantities are tonnes CO2e and may be fractional. They are kept as
// json.Number so the exact decimal recorded in the registry reaches the chain,
// with no float rounding. An absent field decodes to the empty Number.
type MonitoringVerification struct {
	ActivityMonitoringPeriodVerificationID string `json:"activity_monitoring_period_verification_id"`
	ActivityID                             string `json:"activity_id"`
	CertificateOfComplianceID              string `json:"certificate_of_compliance_id"`
	StatusCode                             string `json:"status_code"`

	PermanentNetCarbonRemovalBenefit              json.Number `json:"permanent_net_carbon_removal_benefit"`
	CarbonFarmingTemporaryNetCarbonRemovalBenefit json.Number `json:"carbon_farming_temporary_net_carbon_removal_benefit"`
	CarbonFarmingNetSoilEmissionReductionBenefit  json.Number `json:"carbon_farming_net_soil_emission_reduction_benefit"`
	CarbonStorageTemporaryNetCarbonRemovalBenefit json.Number `json:"carbon_storage_temporary_net_carbon_removal_benefit"`
}

// CreditBatch describes one carbon credit batch to mint for a monitoring
// verification: one per benefit field that carries a value.
type CreditBatch struct {
	// BatchID is the carbon_credit_batch_id stored on-chain. It is derived from
	// the verification id and the unit type, so re-processing the same
	// verification always arrives at the same batch.
	BatchID   string
	BatchName string
	UnitType  string
	// Amount is the benefit value in tonnes CO2e, as recorded.
	Amount json.Number
}

// CreditBatches returns one CreditBatch per benefit field with a value > 0.
// A field that does not parse as a non-negative decimal is an error, so a
// malformed amount is never minted as something else.
func (m MonitoringVerification) CreditBatches() ([]CreditBatch, error) {
	all := []struct {
		field, suffix, name, unitType string
		amount                        json.Number
	}{
		{"permanent_net_carbon_removal_benefit", "permanent_removal",
			"Permanent Removals Carbon Credit Batch", UnitTypePermanentRemoval,
			m.PermanentNetCarbonRemovalBenefit},
		{"carbon_farming_temporary_net_carbon_removal_benefit", "carbon_farming_sequestration",
			"Carbon Farming Removals Carbon Credit Batch", UnitTypeCarbonFarming,
			m.CarbonFarmingTemporaryNetCarbonRemovalBenefit},
		{"carbon_farming_net_soil_emission_reduction_benefit", "soil_emission_reduction",
			"Soil Emissions Reductions Carbon Credit Batch", UnitTypeSoilEmissionReduction,
			m.CarbonFarmingNetSoilEmissionReductionBenefit},
		{"carbon_storage_temporary_net_carbon_removal_benefit", "carbon_storage_in_product",
			"Carbon Storage in Product Carbon Credit Batch", UnitTypeCarbonStorageProduct,
			m.CarbonStorageTemporaryNetCarbonRemovalBenefit},
	}
	var out []CreditBatch
	for _, c := range all {
		if c.amount == "" {
			continue
		}
		units, err := ToBaseUnits(c.amount, 18)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", c.field, err)
		}
		if units.Sign() == 0 {
			continue
		}
		out = append(out, CreditBatch{
			BatchID:   m.ActivityMonitoringPeriodVerificationID + ":" + c.suffix,
			BatchName: c.name,
			UnitType:  c.unitType,
			Amount:    c.amount,
		})
	}
	return out, nil
}

// ToBaseUnits converts a non-negative decimal amount (e.g. tonnes CO2e) into
// integer token base units with the given number of decimals, exactly. It
// rejects negative values, exponents and more fractional digits than fit.
func ToBaseUnits(amount json.Number, decimals int) (*big.Int, error) {
	s := strings.TrimSpace(amount.String())
	if s == "" {
		return nil, fmt.Errorf("empty amount")
	}
	whole, frac, _ := strings.Cut(s, ".")
	if whole == "" {
		whole = "0"
	}
	if !isDigits(whole) || (frac != "" && !isDigits(frac)) {
		return nil, fmt.Errorf("amount %q is not a non-negative decimal", s)
	}
	frac = strings.TrimRight(frac, "0")
	if len(frac) > decimals {
		return nil, fmt.Errorf("amount %q has more than %d decimal places", s, decimals)
	}
	frac += strings.Repeat("0", decimals-len(frac))
	units, ok := new(big.Int).SetString(whole+frac, 10)
	if !ok {
		return nil, fmt.Errorf("amount %q is not a non-negative decimal", s)
	}
	return units, nil
}

func isDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

type Certificate struct {
	CertificateOfComplianceID string `json:"certificate_of_compliance_id"`
	CertificationSchemeID     string `json:"certification_scheme_id"`
	CertificationBodyID       string `json:"certification_body_id"`
	AuditReportID             string `json:"audit_report_id"`
	CountryID                 string `json:"country_id"`
	IssueDate                 string `json:"issue_date"`
	ExpiryDate                string `json:"expiry_date"`
	UnitTypes                 string `json:"unit_types"`
	CertificationStatus       string `json:"certification_status"`
	ActivityID                string `json:"activity_id"`

	CarbonRemovalsUnderBaseline                   json.Number `json:"carbon_removals_under_baseline"`
	SoilEmissionsUnderBaseline                    json.Number `json:"soil_emissions_under_baseline"`
	PermanentNetCarbonRemovalBenefit              json.Number `json:"permanent_net_carbon_removal_benefit"`
	CarbonFarmingTemporaryNetCarbonRemovalBenefit json.Number `json:"carbon_farming_temporary_net_carbon_removal_benefit"`
	CarbonFarmingNetSoilEmissionReductionBenefit  json.Number `json:"carbon_farming_net_soil_emission_reduction_benefit"`
	CarbonStorageTemporaryNetCarbonRemovalBenefit json.Number `json:"carbon_storage_temporary_net_carbon_removal_benefit"`
}

type CertificationScheme struct {
	CertificationSchemeID string `json:"certification_scheme_id"`
	Name                  string `json:"name"`
}

type CertificationBody struct {
	CertificationBodyID string `json:"certification_body_id"`
	LegalName           string `json:"legal_name"`
}

type Operator struct {
	OperatorID        string `json:"operator_id"`
	LegalName         string `json:"legal_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	AddressLine1      string `json:"address_line_1"`
	AddressLine2      string `json:"address_line_2"`
	Postcode          string `json:"postcode"`
	CountryID         string `json:"country_id"`
	OGCRWalletAddress string `json:"ogcr_wallet_address"`
}
