package my_test

import (
	"testing"
)

func TestRunAllTests(t *testing.T) {
	t.Parallel()
	t.Run("TestBetHasOddsBoost", func(t *testing.T) {
		bet1 := Bet{
			BetID:        101,
			Amount:       50.00,
			HasOddsBoost: false,
		}

		if bet1.GetOddsBoostStatus() != false {
			t.Errorf("Expected no odds boost, but got %v", bet1.GetOddsBoostStatus())
		}

		// Creating a bet with odds boost
		bet2 := Bet{
			BetID:        102,
			Amount:       100.00,
			HasOddsBoost: true,
		}

		if bet2.GetOddsBoostStatus() != true {
			t.Errorf("Expected odds boost, but got %v", bet2.GetOddsBoostStatus())
		}
	})

	t.Run("TestBetID", func(t *testing.T) {
		t.Parallel()
		bet1 := Bet{
			BetID:        101,
			Amount:       50.00,
			HasOddsBoost: false,
		}

		if bet1.GetBetID() != 101 {
			t.Errorf("Expected odds boost, but got %v", bet1.GetBetID())
		}
	})

	t.Run("TestAmount", func(t *testing.T) {
		t.Parallel()
		bet1 := Bet{
			BetID:        101,
			Amount:       50.00,
			HasOddsBoost: false,
		}

		if bet1.GetAmount() != 50.00 {
			t.Errorf("Expected odds boost, but got %v", bet1.GetBetID())
		}
	})
}
