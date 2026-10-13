package obp

import "testing"

func TestEntityPathsFollowTheSpace(t *testing.T) {
	bank := NewClient("http://obp.test/", "u", "p", "k", "ogcr")
	if got, want := bank.ParcelURL("p1"), "http://obp.test/obp/dynamic-entity/banks/ogcr/parcel/p1"; got != want {
		t.Errorf("bank ParcelURL = %q, want %q", got, want)
	}
	if got, want := bank.entityPath("activity_verification"), "/obp/dynamic-entity/banks/ogcr/activity_verification"; got != want {
		t.Errorf("bank entityPath = %q, want %q", got, want)
	}

	sys := NewClient("http://obp.test", "u", "p", "k", "")
	if got, want := sys.ActivityURL("a1"), "http://obp.test/obp/dynamic-entity/activity/a1"; got != want {
		t.Errorf("system ActivityURL = %q, want %q", got, want)
	}
}
