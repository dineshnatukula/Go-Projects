package backend

import (
	"encoding/json"
	"fmt"
	offer "myproject/internal/models/models/offer/full"
)

func TestFullMarket() {
	var market *offer.Market
	err := json.Unmarshal([]byte(marketJson), &market)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	// fmt.Println("event.id", event)
}

func TestInternalMarket() {
	var market *offer.InternalMarket
	err := json.Unmarshal([]byte(marketJson), &market)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	// fmt.Println("event.id", event.ID)
}

var marketJson = (`{
	"id": "U-70983541",
	"version": 41,
	"sportId": "FBL",
	"eventId": "U-2365128",
	"name": "Who Will Score 1. Goal? (10 min interval)",
	"nameTid": "",
	"status": {
		"beebet": 2,
		"luckysplash": 2,
		"madridbet": 2,
		"mistycasino": 2,
		"onwin": 2,
		"pokerstarsgermany": 2,
		"quinnbet": 2,
		"slotin": 2,
		"tipobet": 2,
		"zlot": 2
	},
	"statusOverrides": {},
	"tradeStatus": {
		"beebet": 2,
		"luckysplash": 2,
		"madridbet": 2,
		"mistycasino": 2,
		"onwin": 2,
		"pokerstarsgermany": 2,
		"quinnbet": 2,
		"slotin": 2,
		"tipobet": 2,
		"zlot": 2
	},
	"tradeStatusOverrides": {},
	"marketType": {
		"legType": "SIMPLE",
		"resultType": "",
		"picks": 1,
		"orderImportant": false,
		"eachWay": {
			"isEachWay": false,
			"placeReduction": "0"
		}
	},
	"taxonomy": {
		"type": "WINNER2D-E-WS-10",
		"period": "CUSTOM",
		"scoreType": "GL",
		"isLive": false
	},
	"properties": {
		"FeedMsgTime": "1748022885000",
		"externalId": "sr:match:57310963#101#goalnr=1",
		"feedId": "BETRADAR",
		"messageTime": "2025-05-23 17:54:45.318 +0000 UTC",
		"participantIds": "U-39398,U-54164",
		"player1": "U-39398",
		"player2": "U-54164",
		"productID": "3",
		"regulatorBetTypeId": "",
		"regulatorSportEventId": "6500148df864a8d54c10ad0d",
		"scoreNumber": "1"
	},
	"selections": [
		[
			{
				"id": "G-1-10",
				"participantId": "",
				"name": "1-10",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "4.95"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 3,
							"down": 1,
							"dec": "4",
							"american": "300",
							"hongkong": "3",
							"malay": "-0.333",
							"indonesian": "3"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "400",
							"hongkong": "4",
							"malay": "-0.25",
							"indonesian": "4"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "-128",
							"hongkong": "0.78",
							"malay": "0.78",
							"indonesian": "-1.28"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.158334",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "598",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-11-20",
				"participantId": "",
				"name": "11-20",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "5.2"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 13,
							"down": 4,
							"dec": "4.25",
							"american": "325",
							"hongkong": "3.25",
							"malay": "-0.308",
							"indonesian": "3.25"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "425",
							"hongkong": "4.25",
							"malay": "-0.235",
							"indonesian": "4.25"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 17,
							"down": 4,
							"dec": "5.25",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.150431",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "600",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-21-30",
				"participantId": "",
				"name": "21-30",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "6.15"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 4,
							"down": 1,
							"dec": "5",
							"american": "400",
							"hongkong": "4",
							"malay": "-0.25",
							"indonesian": "4"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "525",
							"hongkong": "5.25",
							"malay": "-0.19",
							"indonesian": "5.25"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 21,
							"down": 4,
							"dec": "6.25",
							"american": "-114",
							"hongkong": "0.87",
							"malay": "0.87",
							"indonesian": "-1.14"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.125751",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "602",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-31-40",
				"participantId": "",
				"name": "31-40",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "7.3"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 19,
							"down": 4,
							"dec": "5.75",
							"american": "475",
							"hongkong": "4.75",
							"malay": "-0.211",
							"indonesian": "4.75"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "650",
							"hongkong": "6.5",
							"malay": "-0.154",
							"indonesian": "6.5"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.10474",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "604",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-41-50",
				"participantId": "",
				"name": "41-50",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "7.45"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 19,
							"down": 4,
							"dec": "5.75",
							"american": "475",
							"hongkong": "4.75",
							"malay": "-0.211",
							"indonesian": "4.75"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "650",
							"hongkong": "6.5",
							"malay": "-0.154",
							"indonesian": "6.5"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 13,
							"down": 2,
							"dec": "7.5",
							"american": "-121",
							"hongkong": "0.82",
							"malay": "0.82",
							"indonesian": "-1.21"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.102915",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "606",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-51-60",
				"participantId": "",
				"name": "51-60",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "10.3"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 7,
							"down": 1,
							"dec": "8",
							"american": "700",
							"hongkong": "7",
							"malay": "-0.143",
							"indonesian": "7"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "950",
							"hongkong": "9.5",
							"malay": "-0.105",
							"indonesian": "9.5"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.073742",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "608",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-61-70",
				"participantId": "",
				"name": "61-70",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "13"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 17,
							"down": 2,
							"dec": "9.5",
							"american": "850",
							"hongkong": "8.5",
							"malay": "-0.118",
							"indonesian": "8.5"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "1200",
							"hongkong": "12",
							"malay": "-0.083",
							"indonesian": "12"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 12,
							"down": 1,
							"dec": "13",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.0578655",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "610",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-71-80",
				"participantId": "",
				"name": "71-80",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "16.5"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 10,
							"down": 1,
							"dec": "11",
							"american": "1000",
							"hongkong": "10",
							"malay": "-0.1",
							"indonesian": "10"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "1600",
							"hongkong": "16",
							"malay": "-0.062",
							"indonesian": "16"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 16,
							"down": 1,
							"dec": "17",
							"american": "-125",
							"hongkong": "0.8",
							"malay": "0.8",
							"indonesian": "-1.25"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.0454128",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "612",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-81-90",
				"participantId": "",
				"name": "81-90",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "14.6"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 19,
							"down": 2,
							"dec": "10.5",
							"american": "950",
							"hongkong": "9.5",
							"malay": "-0.105",
							"indonesian": "9.5"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "1400",
							"hongkong": "14",
							"malay": "-0.071",
							"indonesian": "14"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 14,
							"down": 1,
							"dec": "15",
							"american": "-129",
							"hongkong": "0.77",
							"malay": "0.77",
							"indonesian": "-1.29"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.0516526",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "614",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			},
			{
				"id": "G-NONE",
				"participantId": "",
				"name": "NONE",
				"nameTid": "",
				"status": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"statusOverrides": {},
				"tradeStatus": {
					"beebet": 2,
					"luckysplash": 2,
					"madridbet": 2,
					"mistycasino": 2,
					"onwin": 2,
					"pokerstarsgermany": 2,
					"quinnbet": 2,
					"slotin": 2,
					"tipobet": 2,
					"zlot": 2
				},
				"tradeStatusOverrides": {},
				"price": {
					"0": {
						"up": 0,
						"down": 0,
						"dec": "6"
					}
				},
				"priceOverrides": {},
				"closingPMPrice": {
					"beebet": {
						"beebet-0": {
							"up": 15,
							"down": 4,
							"dec": "4.75",
							"american": "375",
							"hongkong": "3.75",
							"malay": "-0.267",
							"indonesian": "3.75"
						}
					},
					"luckysplash": {
						"luckysplash-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"madridbet": {
						"madridbet-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"mistycasino": {
						"mistycasino-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"onwin": {
						"onwin-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"pokerstarsgermany": {
						"pokerstarsgermany-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"quinnbet": {
						"quinnbet-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "500",
							"hongkong": "5",
							"malay": "-0.2",
							"indonesian": "5"
						}
					},
					"slotin": {
						"slotin-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"tipobet": {
						"tipobet-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					},
					"zlot": {
						"zlot-0": {
							"up": 5,
							"down": 1,
							"dec": "6",
							"american": "-117",
							"hongkong": "0.85",
							"malay": "0.85",
							"indonesian": "-1.17"
						}
					}
				},
				"side": 0,
				"line": 0,
				"pushHonored": false,
				"homeScore": 0,
				"awayScore": 0,
				"band": {
					"lower": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					},
					"upper": {
						"value": 0,
						"isValueIncluded": false,
						"isZero": true
					}
				},
				"probability": "0.129157",
				"selTemplate": "SCORE",
				"movingLine": false,
				"externalID": "616",
				"favouriteRank": 0,
				"isPriceAvailable": null,
				"isProxyPriceAvailable": null,
				"OisPriceAvailable": null,
				"emergency": false,
				"isScratched": false,
				"runnerType": "",
				"runnerNo": 0
			}
		]
	],
	"betting": {
		"beebet": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"luckysplash": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"madridbet": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"mistycasino": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"onwin": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"pokerstarsgermany": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"quinnbet": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"slotin": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"tipobet": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		},
		"zlot": {
			"startTime": "0001-01-01T00:00:00Z",
			"endTime": "2025-05-24T19:00:00Z"
		}
	},
	"betDelay": 0,
	"isBetDelayValid": false,
	"hasBetDelay": false,
	"availableBetTypes": [],
	"resultedTime": {},
	"effectiveResult": null,
	"traderResult": null,
	"feedResult": {},
	"betVoidPeriods": [],
	"traderVoidPeriods": {},
	"nonBrandedOverrides": {},
	"pcbAdjustment": {
		"0": {
			"beebet": {
				"p": "0.02",
				"c": "0.01",
				"b": "0.92",
				"lId": "",
				"usePL": false
			},
			"luckysplash": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			},
			"madridbet": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			},
			"mistycasino": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			},
			"onwin": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			},
			"pokerstarsgermany": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			},
			"quinnbet": {
				"p": "0",
				"c": "0",
				"b": "1",
				"lId": "",
				"usePL": false
			}
		}
	},
	"modifiedBy": "FEED",
	"startTime": "2025-05-24T19:00:00Z",
	"hasbets": [],
	"pcbConfig": {
		"0": {
			"beebet": {
				"usePL": false,
				"adjustments": {
					"0": {
						"p": "0.02",
						"c": "0.01",
						"b": "0.93",
						"lId": ""
					},
					"-100000": {
						"p": "0.02",
						"c": "0.01",
						"b": "0.92",
						"lId": ""
					},
					"-1440": {
						"p": "0.02",
						"c": "0.01",
						"b": "0.93",
						"lId": ""
					},
					"-60": {
						"p": "0.02",
						"c": "0.01",
						"b": "0.94",
						"lId": ""
					}
				}
			},
			"luckysplash": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"madridbet": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"mistycasino": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"onwin": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"pokerstarsgermany": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"quinnbet": {
				"usePL": false,
				"adjustments": {
					"-100000": {
						"p": "0",
						"c": "0",
						"b": "1",
						"lId": ""
					}
				}
			},
			"slotin": {
				"usePL": false,
				"adjustments": {}
			},
			"tipobet": {
				"usePL": false,
				"adjustments": {}
			},
			"zlot": {
				"usePL": false,
				"adjustments": {}
			}
		}
	},
	"betRestrictions": null,
	"license": "BETRADAR-L2",
	"isEachWay": false,
	"updatedAt": "2025-05-23T17:54:54.13081Z"
}`)
