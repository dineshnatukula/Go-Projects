package strings

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	tests := ReadTests("C:/Users/dinesh.natukula/Documents/Testing/internal/backend/strings/data.csv")
	for idx, test := range tests {
		t.Run(fmt.Sprintf("RestRegexMatch %d", (idx+1)), func(t *testing.T) {
			output, err := strconv.ParseBool(test[2])
			fmt.Println(test, test[0], test[1], test[2], output, err, IsMatch(test[0], test[1]))
			if IsMatch(test[0], test[1]) != output {
				t.Errorf("Expected %t, but got %t", output, IsMatch("aaa", "a*"))
			}
		})
	}
}
