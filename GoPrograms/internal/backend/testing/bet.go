package my_test

// Bet struct represents a single bet with an odds boost field
type Bet struct {
	BetID        int
	Amount       float64
	HasOddsBoost bool
}

func (b Bet) GetOddsBoostStatus() bool {
	return b.HasOddsBoost
}

func (b Bet) GetBetID() int {
	return b.BetID
}

func (b Bet) GetAmount() float64 {
	return b.Amount
}
