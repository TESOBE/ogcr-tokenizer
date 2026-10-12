package obp

import (
	"encoding/json"
	"testing"
)

func TestToBaseUnits(t *testing.T) {
	cases := []struct {
		in      string
		want    string
		wantErr bool
	}{
		{in: "1", want: "1000000000000000000"},
		{in: "6.5", want: "6500000000000000000"},
		{in: "120.50", want: "120500000000000000000"},
		{in: "0.000000000000000001", want: "1"},
		{in: "0", want: "0"},
		{in: "0.0", want: "0"},
		{in: ".5", want: "500000000000000000"},
		{in: "0.0000000000000000001", wantErr: true}, // 19 decimal places
		{in: "-1", wantErr: true},
		{in: "1e3", wantErr: true},
		{in: "abc", wantErr: true},
		{in: "", wantErr: true},
	}
	for _, c := range cases {
		got, err := ToBaseUnits(json.Number(c.in), 18)
		if c.wantErr {
			if err == nil {
				t.Errorf("ToBaseUnits(%q) = %s, want error", c.in, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("ToBaseUnits(%q): %v", c.in, err)
			continue
		}
		if got.String() != c.want {
			t.Errorf("ToBaseUnits(%q) = %s, want %s", c.in, got, c.want)
		}
	}
}

func TestCreditBatches(t *testing.T) {
	// Decoded from JSON, as the client does, so absent and zero fields are
	// exercised the way the registry delivers them.
	var mv MonitoringVerification
	body := `{
		"activity_monitoring_period_verification_id": "mv-1",
		"activity_id": "a-1",
		"permanent_net_carbon_removal_benefit": 120.5,
		"carbon_farming_temporary_net_carbon_removal_benefit": 0,
		"carbon_storage_temporary_net_carbon_removal_benefit": 7
	}`
	if err := json.Unmarshal([]byte(body), &mv); err != nil {
		t.Fatal(err)
	}

	got, err := mv.CreditBatches()
	if err != nil {
		t.Fatal(err)
	}
	want := []CreditBatch{
		{"mv-1:permanent_removal", "Permanent Removals Carbon Credit Batch", UnitTypePermanentRemoval, "120.5"},
		{"mv-1:carbon_storage_in_product", "Carbon Storage in Product Carbon Credit Batch", UnitTypeCarbonStorageProduct, "7"},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d batches, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("batch %d = %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestCreditBatchesRejectsMalformedAmount(t *testing.T) {
	mv := MonitoringVerification{
		ActivityMonitoringPeriodVerificationID:       "mv-1",
		PermanentNetCarbonRemovalBenefit:             "10",
		CarbonFarmingNetSoilEmissionReductionBenefit: "-3",
	}
	if got, err := mv.CreditBatches(); err == nil {
		t.Fatalf("want error for negative amount, got %+v", got)
	}
}
