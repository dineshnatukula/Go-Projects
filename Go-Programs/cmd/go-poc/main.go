package main

import (
	"fmt"
	"sort"
)

func main() {
	type playerStat struct {
		PlayerID string `json:"playerID"`
		BrandID  string `json:"brandID"`
		IsLive   bool   `json:"isLive"`
	}

	playerStatsKeys := []playerStat{}
	// playerStatsKeys1 := []playerStat{}
	p1 := playerStat{
		PlayerID: "abc", BrandID: "pp", IsLive: false,
	}
	p2 := playerStat{
		PlayerID: "bac", BrandID: "pp", IsLive: true,
	}
	p3 := playerStat{
		PlayerID: "bad", BrandID: "pp", IsLive: false,
	}
	p4 := playerStat{
		PlayerID: "bac", BrandID: "pp1", IsLive: false,
	}
	p5 := playerStat{
		PlayerID: "bac", BrandID: "pp", IsLive: false,
	}

	playerStatsKeys = append(playerStatsKeys, p1)
	playerStatsKeys = append(playerStatsKeys, p5)
	playerStatsKeys = append(playerStatsKeys, p3)
	playerStatsKeys = append(playerStatsKeys, p2)
	playerStatsKeys = append(playerStatsKeys, p4)

	// playerStatsKeys1 = playerStatsKeys

	sort.Slice(playerStatsKeys, func(i, j int) bool {
		return playerStatsKeys[i].PlayerID < playerStatsKeys[j].PlayerID &&
			// playerStatsKeys[i].BrandID < playerStatsKeys[j].BrandID &&
			playerStatsKeys[i].IsLive
	})
	fmt.Println("Old Sorting")
	fmt.Println(playerStatsKeys)

	sort.Slice(playerStatsKeys, func(i, j int) bool {
		if playerStatsKeys[i].PlayerID != playerStatsKeys[j].PlayerID {
			return playerStatsKeys[i].PlayerID < playerStatsKeys[j].PlayerID
		}
		if playerStatsKeys[i].BrandID != playerStatsKeys[j].BrandID {
			return playerStatsKeys[i].BrandID < playerStatsKeys[j].BrandID
		}
		return playerStatsKeys[i].IsLive && !playerStatsKeys[j].IsLive
	})

	fmt.Println("New Sorting")
	fmt.Println(playerStatsKeys)
}
