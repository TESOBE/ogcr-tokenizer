package obp

type ParcelOwnerVerification struct {
	ParcelOwnerVerificationID string `json:"parcel_owner_verification_id"`
	ParcelID                  string `json:"parcel_id"`
	StatusCode                string `json:"status_code"`
	StatusMessage             string `json:"status_message"`
}

type ParcelOwnerVerificationList struct {
	List []ParcelOwnerVerification `json:"parcel_owner_verification_list"`
}

type ActivityVerification struct {
	ActivityVerificationID string `json:"activity_verification_id"`
	ActivityID             string `json:"activity_id"`
	StatusCode             string `json:"status_code"`
	StatusMessage          string `json:"status_message"`
}

type ActivityVerificationList struct {
	List []ActivityVerification `json:"activity_verification_list"`
}

type ActivityParcelVerification struct {
	ActivityParcelVerificationID string `json:"activity_parcel_verification_id"`
	ParcelID                     string `json:"parcel_id"`
	ActivityID                   string `json:"activity_id"`
	Amount                       int64  `json:"amount"`
	StatusCode                   string `json:"status_code"`
	StatusMessage                string `json:"status_message"`
}

type ActivityParcelVerificationList struct {
	List []ActivityParcelVerification `json:"activity_parcel_verification_list"`
}

type Activity struct {
	ActivityID                string `json:"activity_id"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	CountryCode               string `json:"country_code"`
	City                      string `json:"city"`
	StartDate                 string `json:"start_date"`
	EndDate                   string `json:"end_date"`
	ActivityPlanID            string `json:"activity_plan_id"`
	Type                      string `json:"type"`
	Methodologies             string `json:"methodologies"`
	MonitoringPeriodYears     int    `json:"monitoring_period_years"`
	MonitoringPeriodStartDate string `json:"monitoring_period_start_date"`
	MonitoringPeriodEndDate   string `json:"monitoring_period_end_date"`
	OperatorID                string `json:"operator_id"`
}

type MonitoringVerification struct {
	ActivityMonitoringPeriodVerificationID string `json:"activity_monitoring_period_verification_id"`
	ActivityID                             string `json:"activity_id"`
	StatusCode                             string `json:"status_code"`

	PermanentNetCarbonRemovalBenefit              int64 `json:"permanent_net_carbon_removal_benefit"`
	CarbonFarmingTemporaryNetCarbonRemovalBenefit int64 `json:"carbon_farming_temporary_net_carbon_removal_benefit"`
	CarbonFarmingNetSoilEmissionReductionBenefit  int64 `json:"carbon_farming_net_soil_emission_reduction_benefit"`
	CarbonStorageTemporaryNetCarbonRemovalBenefit int64 `json:"carbon_storage_temporary_net_carbon_removal_benefit"`
}

type MonitoringVerificationList struct {
	List []MonitoringVerification `json:"activity_monitoring_period_verification_list"`
}

// CreditAmount pairs a credit_type (the benefit field name) with its amount.
type CreditAmount struct {
	CreditType string
	Amount     int64
}

// CreditAmounts returns the benefit fields with amount > 0 — one CarbonCreditBatch
// (and its ERC-20 credits) is minted per entry.
func (m MonitoringVerification) CreditAmounts() []CreditAmount {
	all := []CreditAmount{
		{"permanent_net_carbon_removal_benefit", m.PermanentNetCarbonRemovalBenefit},
		{"carbon_farming_temporary_net_carbon_removal_benefit", m.CarbonFarmingTemporaryNetCarbonRemovalBenefit},
		{"carbon_farming_net_soil_emission_reduction_benefit", m.CarbonFarmingNetSoilEmissionReductionBenefit},
		{"carbon_storage_temporary_net_carbon_removal_benefit", m.CarbonStorageTemporaryNetCarbonRemovalBenefit},
	}
	var out []CreditAmount
	for _, c := range all {
		if c.Amount > 0 {
			out = append(out, c)
		}
	}
	return out
}

type Certificate struct {
	CertificateOfComplianceID string `json:"certificate_of_compliance_id"`
	CertificationSchemeID     string `json:"certification_scheme_id"`
	CertificationBodyID       string `json:"certification_body_id"`
	IssueDate                 string `json:"issue_date"`
	ExpiryDate                string `json:"expiry_date"`
	CertificationStatus       string `json:"certification_status"`
	ActivityID                string `json:"activity_id"`
}

type CertificateList struct {
	List []Certificate `json:"certificate_of_compliance_list"`
}

type Operator struct {
	OperatorID        string `json:"operator_id"`
	LegalName         string `json:"legal_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	AddressLine1      string `json:"address_line_1"`
	AddressLine2      string `json:"address_line_2"`
	Postcode          string `json:"postcode"`
	CountryCode       string `json:"country_code"`
	OGCRWalletAddress string `json:"ogcr_wallet_address"`
}
