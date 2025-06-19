package models

type Bet struct {
	BetID    int     `json:"betID"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
