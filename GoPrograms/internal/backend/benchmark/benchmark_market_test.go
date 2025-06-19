package backend

import (
	"encoding/json"
	offer "myproject/internal/models/models/offer/full"
	"testing"
)

func BenchmarkFulMarket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TestFullMarket()
		var market *offer.Market
		json.Unmarshal([]byte(marketJson), &market)
	}
}

func BenchmarkInternalMarket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestInternalMarket()
	}
}
