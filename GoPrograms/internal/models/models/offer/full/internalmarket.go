package offer

type InternalMarket struct {
	Status       map[BrandID]Status           `json:"status"`               // a status designator of the market.Describes whether the market should be offered to customers or not
	OStatus      map[BrandID]BrandStatus      `json:"statusOverrides"`      // status overrides
	TradeStatus  map[BrandID]TradeStatus      `json:"tradeStatus"`          // a trade status designator of the Market. Used to suspend or allow placing bets.status of the market
	OTradeStatus map[BrandID]BrandTradeStatus `json:"tradeStatusOverrides"` // trade status overrides
	ID           MarketID                     `json:"id"`                   // unique identifier of the market//
	Version      int                          `json:"version"`              // version of the market
	SportID      `json:"sportId"`             // id of the sport to which market belongs to
	EventID      `json:"eventId"`             // id of the event to which market belongs to
	Name         string                       `json:"name"`    // full name of the market
	NameTid      string                       `json:"nameTid"` // market name translation id
	MarketType   `json:"marketType"`          // a structure that describes the type of a given Market from the perspective of data relevant for placing and settling bets, i.e., what data should be preserved on the bet
	Taxonomy     MarketTaxonomy               `json:"taxonomy"` // taxonomy describes the type of the given market from the perspective of grouping, categorizing and searching similar markets on the frontend.
	// Properties          map[string]string                                       `json:"properties"`                 // various market properties
	// Selections          [][]Selection                                           `json:"selections,omitempty"`       // holds all possible outcomes of market (example home/away/draw)
	// Betting             map[BrandID]TimeInterval                                `json:"betting,omitempty"`          // when bets area allowed
	// OBetting            map[BrandID]OverridableTimeInterval                     `json:"bettingOverrides,omitempty"` // overrides for betting interval
	// BetDelay            int                                                     `json:"betDelay"`                   // how long to delay a bet
	// IsBetDelayValid     bool                                                    `json:"isBetDelayValid"`            // validity of betdelay parameter (if 0 is real zero or non-initialized value)
	// HasBetDelay         bool                                                    `json:"hasBetDelay"`                // should bet delay be applied
	// AvailableBetTypes   []BetType                                               `json:"availableBetTypes"`          // available bet types for this market
	// ResultedTime        map[BrandID]time.Time                                   `json:"resultedTime"`               // time of last result per brand
	// EffectiveResult     map[BrandID][]MarketResult                              `json:"effectiveResult"`            // effective market results
	// TraderResult        map[BrandID][]MarketResult                              `json:"traderResult"`               // market results trader override
	// Result              map[BrandID][]MarketResult                              `json:"feedResult"`                 // market results received by feed
	// BetVoidPeriods      `json:"betVoidPeriods"`                                 // bets placed in these intervals should be voided
	// NonBrandedOverrides MarketNonBrandedOverrides                               `json:"nonBrandedOverrides"`             // overrides for non branded attributes
	// Adjustment          map[PricingDomainID]map[BrandID]Adjustment              `json:"adjustment,omitempty"`            // margin adjustment factor
	// COdisabled          map[BrandID]COdisabled                                  `json:"coDisabled,omitempty"`            // status of cashout option on event
	// Overridden          []BrandID                                               `json:"overridden,omitempty"`            // List of overriden markets by Brand, this is populated automatically when object is saved to database. It does not reflect current changes to override of the object
	// ModifiedBy          string                                                  `json:"modifiedBy"`                      // who made last modification on market
	// StartTime           time.Time                                               `json:"startTime"`                       // start time of event this market belongs to
	// MarginConfig        map[PricingDomainID]map[BrandID]MarginConfig            `json:"marginConfig,omitempty"`          // margin adjustment configuration
	// OMarginConfig       map[PricingDomainID]map[BrandID]OverridableMarginConfig `json:"marginConfigOverrides,omitempty"` // margin config overrides
	// UpdatedAt           time.Time                                               `json:"updatedAt,omitempty"`             // timestamp of last modification
	// Deletable           bool                                                    `json:"deletable,omitempty"`
	// BetRestrictions     BetRestrictions                                         `json:"betRestrictions"` //resetriction on market for allowed numer of legs
	License string `json:"license"`
}
