package offer

import (
	"time"

	"github.com/shopspring/decimal"
)

type MarketID string
type SportID string
type EventID string
type BrandID string
type Status int
type OverrideType int
type TraderID string
type ChannelID string
type TradeStatus int
type PricingDomainID string
type BetType int
type SelectionID string
type RelativeTimePeriod int
type BetVoidPeriods []VoidPeriod
type MarginConfig map[RelativeTimePeriod]Adjustment
type BetRestrictions map[BrandID]BetRestriction
type SelectionResult map[SelectionID]SelectionOutcome

// SelectionOutcome is enumeration describing outcome of selections!
type SelectionOutcome int
type SelectionResultTime map[SelectionID]time.Time

// Side is enumeration describing position of selection in line oriented markets!
type Side int

type BrandStatus struct {
	Original     Status       `json:"original"`
	Override     Status       `json:"override"`
	OverrideType OverrideType `json:"overrideType"`
	Who          TraderID     `json:"who"`
	When         time.Time    `json:"when"`
}

type Price struct {
	Up      int             `json:"up"`   // “up” side of the price in UK format.
	Down    int             `json:"down"` // “down” side of the price in UK format.
	Decimal decimal.Decimal `json:"dec"`  // decimal price
}
type OverridablePrice struct {
	Original     Price        `json:"original"`
	Override     Price        `json:"override"`
	OverrideType OverrideType `json:"overrideType"`
	Who          TraderID     `json:"who"`
	When         time.Time    `json:"when"`
}

type MarketType struct {
	LegType    string `json:"legType"`    // defines the type of the selections within the market (e.g., whether selections consider a simple or line type of betting).This should be changed to slice of acceptable leg types, but that can also be implicit information
	ResultType string `json:"resultType"` // this is implicit information - should be removed from struct
	Picks      int    `json:"picks"`      // number of selections that should be recorded in the bet leg.  Used for, e.g., forecast or tri-cast horse racing betting.  Defaults to 1 which means that only a single Selection is chosen and recorded per bet leg.  Not in use.

	OrderImportant bool `json:"orderImportant"` // flag that designates whether the order in which the Selections are recorded in bet legs is relevant or not.  Not in use.

	EachWay `json:"eachWay"` // additional properties for each way type of markets
}

type EachWay struct {
	IsEachWay      bool            `json:"isEachWay"`
	PlaceTerm      string          `json:"placeTerm,omitempty"`
	PlaceReduction decimal.Decimal `json:"placeReduction,omitempty"`
}
type ParticipantID string

type Selection struct {
	ID             SelectionID                                `json:"id"` // unique identifier of the selection
	ParticipantID  `json:"participantId"`                     // id of the participant to which given selection relates
	Name           string                                     `json:"name"`                 // full name of a given selection
	NameTid        string                                     `json:"nameTid"`              // translation id
	Status         map[BrandID]Status                         `json:"status"`               // a status designator of the selection.  Describes whether the Selection should be offered to customers or not (e.g., shown on betting site).
	OStatus        map[BrandID]BrandStatus                    `json:"statusOverrides"`      // status overrides
	TradeStatus    map[BrandID]TradeStatus                    `json:"tradeStatus"`          // trade status designator of the selection.  used to suspend or allow placing bets.
	OTradeStatus   map[BrandID]BrandTradeStatus               `json:"tradeStatusOverrides"` // trade status overrides
	Price          map[PricingDomainID]Price                  `json:"price"`                // price of the selection
	OPrice         map[BrandID]map[ChannelID]OverridablePrice `json:"priceOverrides"`       // price overrides
	ClosingPMPrice map[BrandID]map[ChannelID]Price            `json:"closingPMPrice"`       // price of the selection at the moment when it was closed for trading
	Side           Side                                       `json:"side"`                 // describes the position of the Selection in line type Markets (e.g., home, draw, away, over, under or unknown).
	Line           float64                                    `json:"line"`                 // defines the “line” in line type markets (e.g., for over/under markets the “line” can be defined as 3.5 which would mean that the participant should be over or under 3.5 point difference from the competitor.
	PushHonored    bool                                       `json:"pushHonored"`          // a flag describing whether to void a bet-leg or not when the actual outcome was not offered amongst the available selections.
	HomeScore      float64                                    `json:"homeScore"`            // defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team). Defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team).
	AwayScore      float64                                    `json:"awayScore"`            // defines the result of a home team (a number of “units” (goals, points, etc.) scored by the home team).
	Band           Band                                       `json:"band"`                 // designates the set of results that are considered as winning (e.g., number of goals scored from 4 to 6).
	Probability    string                                     `json:"probability"`          //
	SelTemplate    string                                     `json:"selTemplate"`          // describes the template of the Selection with possible values: HOME, DRAW, AWAY, HOME||DRAW, HOME||AWAY, DRAW||AWAY, HOME&&HOME, HOME&&DRAW, HOME&&AWAY, DRAW&&HOME, DRAW&&DRAW, DRAW&&AWAY, AWAY&&HOME, AWAY&&DRAW, AWAY&&AWAY, OVER, UNDER, SCORE, NONE, ODD, EVEN, EXACT, UNDEFINED, OTHERS'
	MovingLine     bool                                       `json:"movingLine"`
}

type Limit struct {
	Value           float32 `json:"value"`           // actual value of the boundary
	IsValueIncluded bool    `json:"isValueIncluded"` // hlag designating whether the boundary value is included by the Selection or not (< > vs. ≤ ≥).
	IsZero          bool    `json:"isZero"`          // flag designating whether the boundary is to be used or not.
}

// Band describes a set of results that are considered as winning.
type Band struct {
	Lower Limit `json:"lower"` // lower boundary of the band.
	Upper Limit `json:"upper"` // upper boundary of the band.
}

type BrandTradeStatus struct {
	Original     TradeStatus  `json:"original"`
	Override     TradeStatus  `json:"override"`
	OverrideType OverrideType `json:"overrideType"`
	Who          TraderID     `json:"who"`
	When         time.Time    `json:"when"`
}

type MarketTaxonomy struct {
	Type      string `json:"type,omitempty"`
	Period    string `json:"period,omitempty"`
	ScoreType string `json:"scoreType,omitempty"`
	IsLive    bool   `json:"isLive"`
}

type Market struct {
	ID                  MarketID                                                `json:"id"`      // unique identifier of the market//
	Version             int                                                     `json:"version"` // version of the market
	SportID             `json:"sportId"`                                        // id of the sport to which market belongs to
	EventID             `json:"eventId"`                                        // id of the event to which market belongs to
	Name                string                                                  `json:"name"`                 // full name of the market
	NameTid             string                                                  `json:"nameTid"`              // market name translation id
	Status              map[BrandID]Status                                      `json:"status"`               // a status designator of the market.Describes whether the market should be offered to customers or not
	OStatus             map[BrandID]BrandStatus                                 `json:"statusOverrides"`      // status overrides
	TradeStatus         map[BrandID]TradeStatus                                 `json:"tradeStatus"`          // a trade status designator of the Market. Used to suspend or allow placing bets.status of the market
	OTradeStatus        map[BrandID]BrandTradeStatus                            `json:"tradeStatusOverrides"` // trade status overrides
	MarketType          `json:"marketType"`                                     // a structure that describes the type of a given Market from the perspective of data relevant for placing and settling bets, i.e., what data should be preserved on the bet
	Taxonomy            MarketTaxonomy                                          `json:"taxonomy"`                   // taxonomy describes the type of the given market from the perspective of grouping, categorizing and searching similar markets on the frontend.
	Properties          map[string]string                                       `json:"properties"`                 // various market properties
	Selections          [][]Selection                                           `json:"selections,omitempty"`       // holds all possible outcomes of market (example home/away/draw)
	Betting             map[BrandID]TimeInterval                                `json:"betting,omitempty"`          // when bets area allowed
	OBetting            map[BrandID]OverridableTimeInterval                     `json:"bettingOverrides,omitempty"` // overrides for betting interval
	BetDelay            int                                                     `json:"betDelay"`                   // how long to delay a bet
	IsBetDelayValid     bool                                                    `json:"isBetDelayValid"`            // validity of betdelay parameter (if 0 is real zero or non-initialized value)
	HasBetDelay         bool                                                    `json:"hasBetDelay"`                // should bet delay be applied
	AvailableBetTypes   []BetType                                               `json:"availableBetTypes"`          // available bet types for this market
	ResultedTime        map[BrandID]time.Time                                   `json:"resultedTime"`               // time of last result per brand
	EffectiveResult     map[BrandID][]MarketResult                              `json:"effectiveResult"`            // effective market results
	TraderResult        map[BrandID][]MarketResult                              `json:"traderResult"`               // market results trader override
	Result              map[BrandID][]MarketResult                              `json:"feedResult"`                 // market results received by feed
	BetVoidPeriods      `json:"betVoidPeriods"`                                 // bets placed in these intervals should be voided
	NonBrandedOverrides MarketNonBrandedOverrides                               `json:"nonBrandedOverrides"`             // overrides for non branded attributes
	Adjustment          map[PricingDomainID]map[BrandID]Adjustment              `json:"adjustment,omitempty"`            // margin adjustment factor
	COdisabled          map[BrandID]COdisabled                                  `json:"coDisabled,omitempty"`            // status of cashout option on event
	Overridden          []BrandID                                               `json:"overridden,omitempty"`            // List of overriden markets by Brand, this is populated automatically when object is saved to database. It does not reflect current changes to override of the object
	ModifiedBy          string                                                  `json:"modifiedBy"`                      // who made last modification on market
	StartTime           time.Time                                               `json:"startTime"`                       // start time of event this market belongs to
	MarginConfig        map[PricingDomainID]map[BrandID]MarginConfig            `json:"marginConfig,omitempty"`          // margin adjustment configuration
	OMarginConfig       map[PricingDomainID]map[BrandID]OverridableMarginConfig `json:"marginConfigOverrides,omitempty"` // margin config overrides
	UpdatedAt           time.Time                                               `json:"updatedAt,omitempty"`             // timestamp of last modification
	Deletable           bool                                                    `json:"deletable,omitempty"`
	BetRestrictions     BetRestrictions                                         `json:"betRestrictions"` //resetriction on market for allowed numer of legs
	License             string                                                  `json:"license"`
}

type TimeInterval struct {
	StartTime time.Time `json:"startTime,omitempty"` // timestamp of interval start
	EndTime   time.Time `json:"endTime,omitempty"`   // timestamp of interval end
}

type OverridableTimeInterval struct {
	Original     TimeInterval `json:"original"`
	Override     TimeInterval `json:"override"`
	OverrideType OverrideType `json:"overrideType"`
	Who          TraderID     `json:"who"`
	When         time.Time    `json:"when"`
}

type MarketResult struct {
	SelectionResult       SelectionResult       `json:"selectionResult"`                 // map containing results per Selections; key is SelectionId and value is WON|LOST|VOID|HALFWON|HALFLOST|NONRUNNER|OPEN;
	SelectionResultTime   SelectionResultTime   `json:"selectionResultTime"`             // timestamp of the selection result
	LoosersImplied        bool                  `json:"loosersImplied"`                  // flag whether results for all Selections are explicitly given, or the Selections which aren’t stated in the map are deemed as LOST
	IsPartial             bool                  `json:"isPartial"`                       // flag whether the results are partial (not all Selections resulted) or full (all Selections resulted)
	ResultedTime          time.Time             `json:"resultedTime"`                    // timestamp of when the resulting was done;
	ResultType            ResultType            `json:"resultType"`                      // designation whether the type of the result set is PROVISIONAL or CONFIRMED;
	SourceType            TradeSource           `json:"sourceType"`                      // designation whether the result was provided by feed or trader
	MakeUpValue           string                `json:"makeUpValue"`                     // a textual description given by trader when adding the market result. Usually the score (e.g. “3:1”) based on which the result is defined.
	TraderID              string                `json:"traderId"`                        //Trader or feed ID
	OperationIndicator    OperationIndicator    `json:"operationIndicator"`              // indicator of the last operation made on result
	Who                   string                `json:"who"`                             // id of the trader
	Class                 ResultClass           `json:"class"`                           // result class describing the “format” or “type” in which the result is given.
	ResultValue           string                `json:"resultValue"`                     // result value for results of HOME-AWAY, AWAY-HOME and HOME+AWAY result classes. Number of score units (in above examples “2”) is carried here.
	PushHonored           bool                  `json:"pushHonored"`                     // a flag describing whether to void a bet-leg or not when the actual outcome was not offered amongst the available selections
	CorrectScore          CorrectScore          `json:"correctScore,omitempty"`          // result value for results of CORRECT-SCORE result class. Number of units scored by the home and away team (in above example “3:1”) is carried here. Consists of the “homeScore” and “awayScore” properties
	ExtRefID              string                `json:"extRefId,omitempty"`              // unique identifier of the external reference on which the entered score was based upon.
	SelectionOrderResult  SelectionOrderResult  `json:"selectionOrderResult,omitempty"`  // result value for results of FINISHING-ORDER result class. Carries the information on the placement of each of the available selection (e.g. selection x placed y-th or was a non-runner meaning it did not participate in the race or didn’t finish it).
	SelectionDeadHeatRate SelectionDeadHeatRate `json:"selectionDeadHeatRate,omitempty"` //  used in races where multiple participants can finish at the same place. Factor used for calculating the return in the cases
}

type VoidPeriod struct {
	Cancelled bool `json:"cancelled"` // Set if interval is no longer valid
	TimeInterval
}
type MarketNonBrandedOverrides struct {
	BetVoidPeriods *OverridableTimeIntervalSlice `json:"betVoidPeriods,omitempty"`
}

type OverridableTimeIntervalSlice struct {
	Original, Override []VoidPeriod
	OverrideType       OverrideType
	Who                TraderID
	When               time.Time
}

type Adjustment struct {
	Value decimal.Decimal `json:"value"`
}

type COdisabled struct {
	Disabled bool      `json:"disabled"` // status of cashout option
	Who      TraderID  `json:"who"`      // id of the trader who made the last change
	When     time.Time `json:"when"`     // timestamp of the last change
}

// OverridableMarginConfig holds original and overriden values for margin config
type OverridableMarginConfig struct {
	Original, Override MarginConfig
	OverrideType       OverrideType
	Who                TraderID
	When               time.Time
}

type BetRestriction struct {
	Override     LegsRestriction `json:"override"`
	OverrideType OverrideType    `json:"overrideType"`
	Who          TraderID        `json:"who"`
	When         time.Time       `json:"when"`
}

type LegsRestriction struct {
	MinLegs int
	MaxLegs int
}

type ResultType int
type TradeSource int
type OperationIndicator int
type CorrectScore struct {
	Home string `json:"home,omitempty"` // score of the home team
	Away string `json:"away,omitempty"` // score of the away team
}

type SelectionOrderResult map[SelectionID]OrderResult
type SelectionDeadHeatRate map[SelectionID]decimal.Decimal

type OrderResult struct {
	Position int `json:"position"`
	Tied     int `json:"tied"`
}
type ResultClass int
