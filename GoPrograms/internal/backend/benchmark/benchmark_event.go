package backend

import (
	"encoding/json"
	"fmt"
	offer "myproject/internal/models/models/offer/full"
)

func TestFullEvent() {
	var event *offer.Event
	err := json.Unmarshal([]byte(eventJson), &event)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	// fmt.Println("event.id", event)
}

func TestInternalEvent() {
	event := &offer.InternalEvent{}
	err := json.Unmarshal([]byte(eventJson), &event)
	if err != nil {
		fmt.Println("Error is :", err)
	}
	// fmt.Println("event.id", event.ID)
}

var eventJson = (`{
	"id": "U-2273751",
	"feedId": "BETRADAR",
	"version": 0,
	"sportId": "FBL",
	"competitionId": "U-3845",
	"name": "Shenzhen Juniors FC v Guandong GZ-Power FC",
	"nameTid": "",
	"shortName": "SHE vs GUA",
	"status": {
		"beebet": 1,
		"quinnbet": 1,
		"slotin": 1,
		"zlot": 1
	},
	"statusOverrides": {},
	"tradeStatus": {
		"beebet": 2,
		"quinnbet": 2,
		"slotin": 2,
		"zlot": 2
	},
	"tradeStatusOverrides": {},
	"participants": [
		[
			{
				"id": "U-763409",
				"name": "Shenzhen Juniors FC",
				"nameTid": "",
				"shortName": "SHE",
				"gender": "male",
				"participantType": 1,
				"participantTeamData": {
					"participantSide": 1
				}
			},
			{
				"id": "U-763401",
				"name": "Guandong GZ-Power FC",
				"nameTid": "",
				"shortName": "GUA",
				"gender": "male",
				"participantType": 1,
				"participantTeamData": {
					"participantSide": 2
				}
			}
		],
		[
			{
				"id": "U-244074",
				"name": "Yuelei, Cheng",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-763400",
				"name": "Yifan, Tian",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-45105",
				"name": "Yucheng, Shi",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-244136",
				"name": "Pengchao, Zu",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-1244768",
				"name": "Yingjian, Li",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-258662",
				"name": "Yang, Men",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-876762",
				"name": "Zhou, Weijun",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-35338",
				"name": "Marcic, Milan",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-701659",
				"name": "Nzuzi Mata, Kevin",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409"
				}
			},
			{
				"id": "U-258690",
				"name": "Guanghui, Han",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409",
					"jerseyNum": "12"
				}
			},
			{
				"id": "U-253800",
				"name": "Ming, Hu",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409",
					"jerseyNum": "17"
				}
			},
			{
				"id": "U-1259736",
				"name": "Mai, Sijing",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409",
					"jerseyNum": "26"
				}
			},
			{
				"id": "U-294235",
				"name": "Lin, Feiyang",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409",
					"jerseyNum": "36"
				}
			},
			{
				"id": "U-91068",
				"name": "Nouble, Jon",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763409",
					"jerseyNum": "38"
				}
			},
			{
				"id": "U-349392",
				"name": "Wang, Geon Myeong",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-244141",
				"name": "Rosa, Farley",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-84390",
				"name": "Joao Carlos",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-244114",
				"name": "Yin, Shang",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-1192431",
				"name": "Chen, Kanglin",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-1248530",
				"name": "Hao, Yang",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-258831",
				"name": "Xia, Dalong",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			},
			{
				"id": "U-244113",
				"name": "Xueming, Liang",
				"nameTid": "",
				"shortName": "",
				"gender": "male",
				"participantType": 2,
				"participantPlayerData": {
					"teamId": "U-763401"
				}
			}
		]
	],
	"type": 1,
	"properties": {
		"groupInformation": {
			"groupId": "China League 1"
		},
		"groupParentInformation": {},
		"externalId": "sr:match:58590911",
		"matchStructure": "HT-45",
		"competitionCountry": "CN",
		"isInternationalCompetition": "false",
		"competitionGender": "M"
	},
	"anticipated": {
		"startTime": "2025-05-26T11:30:00Z",
		"endTime": "0001-01-01T00:00:00Z"
	},
	"actual": {
		"startTime": "0001-01-01T00:00:00Z",
		"endTime": "0001-01-01T00:00:00Z"
	},
	"betting": {
		"beebet": {
			"startTime": "2025-04-14T19:30:00Z",
			"endTime": "2025-05-26T11:30:00Z"
		},
		"quinnbet": {
			"startTime": "2025-04-14T19:30:00Z",
			"endTime": "2025-05-26T11:30:00Z"
		},
		"slotin": {
			"startTime": "2025-04-14T19:30:00Z",
			"endTime": "2025-05-26T11:30:00Z"
		},
		"zlot": {
			"startTime": "2025-04-14T19:30:00Z",
			"endTime": "2025-05-26T11:30:00Z"
		}
	},
	"hasbets": [],
	"betDelay": 0,
	"isBetDelayValid": false,
	"hasBetDelay": false,
	"result": {
		"currentScore": [],
		"currentPeriod": {
			"id": "",
			"description": "",
			"descriptionTid": "",
			"additional": "",
			"isPeriodFinished": false,
			"shortName": ""
		},
		"nextPeriod": "",
		"clock": {
			"isCountdown": false,
			"isRunning": false,
			"displayClock": false,
			"current": 0,
			"referenceTime": "0001-01-01T00:00:00Z"
		},
		"timeModified": "0001-01-01T00:00:00Z"
	},
	"resultedTime": null,
	"contingency": {
		"relatedevents": []
	},
	"nonBrandedOverrides": {},
	"inPlayBetDelay": {
		"beebet": 10,
		"quinnbet": 5,
		"slotin": 5,
		"zlot": 5
	},
	"modifiedBy": "FEED",
	"racing": {
		"racingresults": null,
		"eventTime": "0001-01-01T00:00:00Z",
		"utcTime": "2025-05-26T11:30:00Z",
		"properties": {
			"racingConfig": null,
			"isHandicapped": false,
			"racingActual": {
				"startTime": "0001-01-01T00:00:00Z",
				"endTime": "0001-01-01T00:00:00Z"
			},
			"bog": null,
			"bogOverride": {},
			"bogTiming": {},
			"bogTimingOverride": {}
		},
		"activeRunnersCounts": {}
	},
	"license": "BETRADAR-L2",
	"updatedAt": "2025-04-21T07:51:36.726412Z"
}`)
