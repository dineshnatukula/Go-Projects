package strings

import (
	"encoding/csv"
	"os"
)

func ReadTests(filePath string) (records [][]string) {
	// func ReadTests() (records []string) {
	// filePath := "C:/Users/dinesh.natukula/Documents/Testing/internal/backend/strings/data.csv"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return append(records, data...)
}
