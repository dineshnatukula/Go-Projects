package offer

import "time"

type InternalEvent struct {
	ID            EventID                      `json:"id"`      // id of event
	FeedID        FeedID                       `json:"feedId"`  //unique id of the feed
	Version       int                          `json:"version"` // version of event
	SportID       `json:"sportId"`             // id of the sport to which event belogs to
	CompetitionID `json:"competitionId"`       // id of the competition to which event belongs to
	Name          string                       `json:"name"`                                   // full name of the event
	Status        map[BrandID]Status           `json:"status"`                                 // status describes whether the event should be offered/shown to customers or not status of event
	OStatus       map[BrandID]BrandStatus      `json:"statusOverrides"`                        // status overrides
	TradeStatus   map[BrandID]TradeStatus      `json:"tradeStatus"`                            // trade status describes whether taking bets should be allowed on any of the Selections of all markets of the event
	OTradeStatus  map[BrandID]BrandTradeStatus `json:"tradeStatusOverrides"`                   // trade status overrides
	EventType     EventType                    `json:"type" db:"type"`                         // type of event (match/outright)
	ModifiedBy    string                       `json:"modifiedBy"`                             // who made last modification on event
	Updated       time.Time                    `json:"updatedAt,omitempty" db:"modified,-ins"` // timestamp of last modification
	License       string                       `json:"license"`
}
