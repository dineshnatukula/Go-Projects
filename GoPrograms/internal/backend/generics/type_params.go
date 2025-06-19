package backend

import (
	"fmt"
	models "myproject/internal/models"
)

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func TypeParams() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	if idx := linearSearch(si, 15); idx {
		fmt.Println("Element Found in Integers array")
	} else {
		fmt.Println("Element Not Foun d in Integers Array")
	}

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	if idx := linearSearch(ss, "Hello"); idx {
		fmt.Println("Element Found in Strings array")
	} else {
		fmt.Println("Element Not Found in Strings Array")
	}

	b1 := models.Bet{
		BetID:    123,
		Amount:   2.2,
		Currency: "EUR",
	}

	b2 := models.Bet{
		BetID:    123,
		Amount:   2.2,
		Currency: "INR",
	}

	b3 := models.Bet{
		BetID:    123,
		Amount:   2.2,
		Currency: "INR",
	}

	bets := []models.Bet{b1, b2}
	if idx := linearSearch(bets, b3); idx {
		fmt.Println("Element Found in Bets array")
	} else {
		fmt.Println("Element Not Foun d in Bets Array")
	}
}

func linearSearch[Type comparable](arr []Type, val Type) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
