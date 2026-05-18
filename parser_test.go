package shared

import "testing"

func TestPlatformValues(t *testing.T) {
	platforms := []Platform{
		PlatformYouTube, PlatformX, PlatformTelegram, PlatformFacebook,
		PlatformInstagram, PlatformLinkedIn, PlatformHabr,
		PlatformStackOverflow, PlatformTBankPulse, PlatformSmartLab,
	}
	seen := map[Platform]bool{}
	for _, p := range platforms {
		if p == "" {
			t.Fatal("empty platform value")
		}
		if seen[p] {
			t.Fatalf("duplicate platform: %s", p)
		}
		seen[p] = true
	}
}
