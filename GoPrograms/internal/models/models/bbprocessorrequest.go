package models

// type BBProcessorRequest struct {
// 	EventId     string   `json:"eventid"`
// 	MarketId    string   `json:"marketid"`
// 	Derivatives []string `json:"derivatives"`
// }

// type Selection struct {
// 	ID    string
// 	Name  string
// 	Price decimal.Decimal
// }

// type DerivativeMarket struct {
// 	ID         string
// 	Selections []Selection
// }

// type Sport struct {
// 	ID   string
// 	Name string
// }

// type EventDetail struct {
// 	Id   string
// 	Name string
// }

// type BBProcessorResponse struct {
// 	DerivativeMarkets []DerivativeMarket
// 	EventDetail       EventDetail
// 	SportDetial       Sport
// 	Participants      [][]offer.Participant
// }

// type BBEvent struct {
// 	ID           offer.EventID `json:"id"`           // id of event
// 	Name         string        `json:"name"`         // full name of the event
// 	Participants string        `json:"participants"` // list of participants involved in the event
// 	SportID      offer.SportID `json:"sportid"`      // id of the sport to which event belogs to
// }

// type Price struct {
// 	Up      int    `json:"up"`   // “up” side of the price in UK format.
// 	Down    int    `json:"down"` // “down” side of the price in UK format.
// 	Decimal string `json:"dec"`  // decimal price
// }

// type TxSelection struct {
// 	Type       string
// 	Selections string
// }

// type ParticipantID string

// // The Price of the original selection has been changed to
// // interface type. So Use this for fetching the data from the DB.
// type BBSelection struct {
// 	ID             offer.SelectionID                                            `json:"id"` // unique identifier of the selection
// 	ParticipantID  `json:"participantId"`                                       // id of the participant to which given selection relates
// 	Name           string                                                       `json:"name"`                 // full name of a given selection
// 	NameTid        string                                                       `json:"nameTid"`              // translation id
// 	Status         map[offer.BrandID]offer.Status                               `json:"status"`               // a status designator of the selection.  Describes whether the Selection should be offered to customers or not (e.g., shown on betting site).
// 	OStatus        map[offer.BrandID]offer.BrandStatus                          `json:"statusOverrides"`      // status overrides
// 	TradeStatus    map[offer.BrandID]offer.TradeStatus                          `json:"tradeStatus"`          // trade status designator of the selection.  used to suspend or allow placing bets.
// 	OTradeStatus   map[offer.BrandID]offer.BrandTradeStatus                     `json:"tradeStatusOverrides"` // trade status overrides
// 	Price          map[offer.PricingDomainID]interface{}                        `json:"price"`                // price of the selection
// 	OPrice         map[offer.BrandID]map[offer.ChannelID]offer.OverridablePrice `json:"priceOverrides"`       // price overrides
// 	ClosingPMPrice map[offer.BrandID]map[offer.ChannelID]Price                  `json:"closingPMPrice"`       // price of the selection at the moment when it was closed for trading
// 	Side           offer.Side                                                   `json:"side"`                 // describes the position of the Selection in line type Markets (e.g., home, draw, away, over, under or unknown).
// 	Line           float64                                                      `json:"line"`                 // defines the “line” in line type markets (e.g., for over/under markets the “line” can be defined as 3.5 which would mean that the participant should be over or under 3.5 point difference from the competitor.
// 	PushHonored    bool                                                         `json:"pushHonored"`          // a flag describing whether to void a bet-leg or not when the actual outcome was not offered amongst the available selections.
// 	HomeScore      float64                                                      `json:"homeScore"`            // defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team). Defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team).
// 	AwayScore      float64                                                      `json:"awayScore"`            // defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team).
// 	Band           offer.Band                                                   `json:"band"`                 // designates the set of results that are considered as winning (e.g., number of goals scored from 4 to 6).
// 	Probability    string                                                       `json:"probability"`          //
// 	SelTemplate    string                                                       `json:"selTemplate"`          // describes the template of the Selection with possible values: HOME, DRAW, AWAY, HOME||DRAW, HOME||AWAY, DRAW||AWAY, HOME&&HOME, HOME&&DRAW, HOME&&AWAY, DRAW&&HOME, DRAW&&DRAW, DRAW&&AWAY, AWAY&&HOME, AWAY&&DRAW, AWAY&&AWAY, OVER, UNDER, SCORE, NONE, ODD, EVEN, EXACT, UNDEFINED, OTHERS'
// 	MovingLine     bool                                                         `json:"movingLine"`
// }
