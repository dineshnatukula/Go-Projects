package offer

import "time"

// FeedID represents unique id of the feed.
type FeedID string

// swagger:model CompetitionID
// CompetitionID represents unique identifier of a Competition instance.
type CompetitionID string

// EventType represents event type
type EventType int
type ExternalID string

// PeriodID id of period
type PeriodID string

// Clock defines event clock
type Clock struct {
	IsCountdown   bool      `json:"isCountdown"`   // a flag describing whether event clock is counting down or up.
	IsRunning     bool      `json:"isRunning"`     // a flag describing whether the event clock is currently running or is stopped.
	DisplayClock  bool      `json:"displayClock"`  // a flag designating whether the current event clock should be displayed or not.
	Current       int       `json:"current"`       // current event clock time in seconds
	ReferenceTime time.Time `json:"referenceTime"` // reference time in UTC when clock properties were as described by values of fields explained above
}

// EventResult contains all known event results
type EventResult struct {
	CurrentScore  []ParticipantScore     `json:"currentScore"` // current score is given per each of the participants of the event.
	CurrentPeriod `json:"currentPeriod"` // descriptor of the current event period
	NextPeriod    PeriodID               `json:"nextPeriod"` // a unique identifier of the event period that follows the current period.
	Clock         `json:"clock"`         // descriptor of the current event time and it’s properties
	TimeModified  time.Time              `json:"timeModified"`
	ResultStats   `json:"resultStat,omitempty"`
}

//ResultStats contains result statistics
type ResultStats map[PeriodID]map[ScType]map[ParticipantID]string

// ScType holds ScoreType
type ScType string

// CurrentPeriod defines current period of the game
type CurrentPeriod struct {
	ID               PeriodID `json:"id"`               // a unique identifier of the current event period.
	Description      string   `json:"description"`      // description of the current event period.
	DescriptionTid   string   `json:"descriptionTid"`   // translation id
	Additional       string   `json:"additional"`       // additional description of the current event period.
	IsPeriodFinished bool     `json:"isPeriodFinished"` // flag designating whether the current event period is finished or not.
	ShortName        string   `json:"shortName"`        // short name of the current event period
}

//EventProperties represents collection of event properties
type EventProperties struct {
	GroupInformation       GroupInformation `json:"groupInformation,omitempty"`
	GroupParentInformation GroupInformation `json:"groupParentInformation,omitempty"` //Parent group information
	ExternalID             ExternalID       `json:"externalId"`                       // event identifier from the feed
	Venue                  string           `json:"venue,omitempty"`
	VenueId                string           `json:"venueId,omitempty"`
	MatchStructure         string           `json:"matchStructure,omitempty"`
	ScoreHistoryData       string           `json:"scoreHistoryData,omitempty"`
}

//GroupInformation contains information of a group
type GroupInformation struct {
	GroupName string  `json:"groupName,omitempty"` // name of the group
	GroupId   GroupId `json:"groupId,omitempty"`   // id of the group
}

//GroupId represent unique identifier of a group
type GroupId string

// Contingency describes relations with other entities in the offer that are mutualy related
// Betting on two contingency related offers is not allowed
type Contingency struct {
	RelatedEvents []EventID `json:"relatedevents"` // array of event ids that are mutually related
}

type Event struct {
	ID                  string                              `json:"id"`      // id of event
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
	License             string                              `json:"license"`
}

// EventNonBrandedOverrides EventNonBrandedOverrides
type EventNonBrandedOverrides struct {
	Anticipated *OverridableTimeInterval `json:"anticipated,omitempty"`
	Real        *OverridableTimeInterval `json:"actual,omitempty"`
	Venue       *OverridableString       `json:"venue,omitempty"`
	Result      *OverridableEventResult  `json:"result,omitempty"`
}

// EventResultOP is a overridable subset of EventResult properties
type EventResultOP struct {
	CurrentScore []ParticipantScore `json:"currentScore,omitempty"`
	TimeModified time.Time          `json:"timeModified,omitempty"`
}

// ParticipantScore defines score for one participant
type ParticipantScore struct {
	ParticipantID `json:"participantId"` // a unique identifier of the actual participant of the event
	Score         []string               `json:"score"` // a complete score for the participant.  For example, for tennis matches this array has the score in sets, games and points in the current game.
}

// OverridableEventResult holds overridable event result
type OverridableEventResult struct {
	Original     EventResultOP `json:"original"`
	Override     EventResultOP `json:"override"`
	OverrideType OverrideType  `json:"overrideType"`
	Who          TraderID      `json:"who"`
	When         time.Time     `json:"when"`
}

// OverridableString holds overridble string
type OverridableString struct {
	Original     string       `json:"original"`
	Override     string       `json:"override"`
	OverrideType OverrideType `json:"overrideType"`
	Who          TraderID     `json:"who"`
	When         time.Time    `json:"when"`
}
