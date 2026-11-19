package obp

import "testing"

// Every path is scoped to the bank the client was made for: the registry keeps the OGCR entities at bank level.
func TestEntityPathsFollowTheBank(t *testing.T) {
	c := NewClient("http://obp.test/", "ogcr", "u", "p", "k")
	if got, want := c.ParcelURL("p1"), "http://obp.test/obp/dynamic-entity/banks/ogcr/parcel/p1"; got != want {
		t.Errorf("ParcelURL = %q, want %q", got, want)
	}
	if got, want := c.entityPath("activity_verification"), "/obp/dynamic-entity/banks/ogcr/activity_verification"; got != want {
		t.Errorf("entityPath = %q, want %q", got, want)
	}
	if got, want := c.ActivityURL("a1"), "http://obp.test/obp/dynamic-entity/banks/ogcr/activity/a1"; got != want {
		t.Errorf("ActivityURL = %q, want %q", got, want)
	}
}

func TestTheBankIdIsEscapedInThePath(t *testing.T) {
	c := NewClient("http://obp.test", "a b/c", "u", "p", "k")
	if got, want := c.entityPath("parcel"), "/obp/dynamic-entity/banks/a%20b%2Fc/parcel"; got != want {
		t.Errorf("entityPath = %q, want %q", got, want)
	}
}
