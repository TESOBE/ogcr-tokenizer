package obp

import "testing"

func TestEntityPathsAreAtTheBank(t *testing.T) {
	c := NewClient("http://obp.test/", "ogcr", "http://oidc.test", "id", "secret")
	if got, want := c.ParcelURL("p1"), "http://obp.test/obp/dynamic-entity/banks/ogcr/parcel/p1"; got != want {
		t.Errorf("ParcelURL = %q, want %q", got, want)
	}
	if got, want := c.entityPath("activity_verification"), "/obp/dynamic-entity/banks/ogcr/activity_verification"; got != want {
		t.Errorf("entityPath = %q, want %q", got, want)
	}
}
