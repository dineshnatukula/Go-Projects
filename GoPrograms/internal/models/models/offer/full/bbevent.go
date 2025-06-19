package offer

import "time"

type BBEvent struct {
	ID                  EventID                             `json:"id"`      // id of event
	FeedID              FeedID                              `json:"feedId"`  //unique id of the feed
	Version             int                                 `json:"version"` // version of event
	SportID             `json:"sportId"`                    // id of the sport to which event belogs to
	CompetitionID       `json:"competitionId"`              // id of the competition to which event belongs to
	Name                string                              `json:"name"`                                   // full name of the event
	NameTid             string                              `json:"nameTid"`                                // event name translation id
	ShortName           string                              `json:"shortName"`                              // event short name
	Status              map[BrandID]Status                  `json:"status"`                                 // status describes whether the event should be offered/shown to customers or not status of event
	OStatus             map[BrandID]BrandStatus             `json:"statusOverrides"`                        // status overrides
	TradeStatus         map[BrandID]TradeStatus             `json:"tradeStatus"`                            // trade status describes whether taking bets should be allowed on any of the Selections of all markets of the event
	OTradeStatus        map[BrandID]BrandTradeStatus        `json:"tradeStatusOverrides"`                   // trade status overrides
	Participants        [][]Participant                     `json:"participants,omitempty"`                 // list of participants involved in the event
	EventType           EventType                           `json:"type" db:"type"`                         // type of event (match/outright)
	Properties          EventProperties                     `json:"properties,omitempty" db:"properties"`   // colection of event properties
	Anticipated         TimeInterval                        `json:"anticipated,omitempty"`                  // originally planned time
	Real                TimeInterval                        `json:"actual,omitempty"`                       // when event really started/ended (feed also updates this)
	Betting             map[BrandID]TimeInterval            `json:"betting,omitempty"`                      // when bets are allowed
	OBetting            map[BrandID]OverridableTimeInterval `json:"bettingOverrides,omitempty"`             // overrides for betting interval
	BetDelay            int                                 `json:"betDelay"`                               // how long to delay the bet
	IsBetDelayValid     bool                                `json:"isBetDelayValid"`                        // validity of betdelay parameter (if 0 is real zero or non-initialized value)
	HasBetDelay         bool                                `json:"hasBetDelay"`                            // do we apply bet delay
	Result              EventResult                         `json:"result,omitempty"`                       // final result of the event
	ResultedTime        map[BrandID]time.Time               `json:"resultedTime"`                           // event result time - based on market result times for this event
	Contingency         Contingency                         `json:"contingency"`                            // relations to other entities
	NonBrandedOverrides EventNonBrandedOverrides            `json:"nonBrandedOverrides" db:",json"`         // overrides for non branded attributes
	InPlayBetDelay      map[BrandID]int                     `json:"inPlayBetDelay,omitempty"`               // bet delay for in play betting
	COdisabled          map[BrandID]COdisabled              `json:"coDisabled,omitempty"`                   // status of cashout option on event
	Overridden          []BrandID                           `json:"overridden,omitempty"`                   // list of overriden markets by Brand, this is populated automatically when object is saved to database. It does not reflect current changes to override of the object
	ModifiedBy          string                              `json:"modifiedBy"`                             // who made last modification on event
	Updated             time.Time                           `json:"updatedAt,omitempty" db:"modified,-ins"` // timestamp of last modification
	Flag                bool                                `json:"flag,omitempty"`
}

type BBProcessorEvent struct {
	EventId     EventID          `json:"eventId"`
	BBEventFlag map[BrandID]bool `json:"bbEventFlag"`
	Provider    string           `json:"provider"`
	License     string           `json:"license"`
}

type BBProcessorMarket struct {
	MarketId     MarketID         `json:"marketId"`
	BBMarketFlag map[BrandID]bool `json:"bbMarketFlag"`
	License      string           `json:"license"`
}
