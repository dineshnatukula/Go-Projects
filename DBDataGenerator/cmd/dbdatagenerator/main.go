package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/dinesh.natukula/mylib"
	"github.com/dinesh.natukula/mylib/dblibs"
)

type Employee struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Aadhar string `json:"aadhar,omitempty"`
	PAN    string `json:"pan,omitempty"`
}

func generatePAN() string {
	// PAN format: ABCDE1234F
	// A to Z: 65 to 90 (ASCII values)
	// 0 to 9: 48 to 57 (ASCII values)

	// Older version of go for rand.seed...
	// rand.Seed(time.Now().UnixNano())

	// Newer version of go for rand.seed...
	rand.New(rand.NewSource(time.Now().UnixNano()))

	firstFiveChars := ""
	for i := 0; i < 5; i++ {
		randomChar := rune(rand.Intn(26) + 65) // ASCII values for A to Z
		firstFiveChars += string(randomChar)
	}

	// Generate the next four characters (digits)
	middleFourChars := fmt.Sprintf("%04d", rand.Intn(10000)) // Random 4-digit number

	// Generate the last character (alphabet)
	lastChar := rune(rand.Intn(26) + 65) // ASCII values for A to Z

	// Concatenate the parts to form the PAN number
	panNumber := fmt.Sprintf("%s%s%s", firstFiveChars, middleFourChars, string(lastChar))

	return panNumber
}

// generateRandomGender generates a random gender ("M" or "F")
func generateRandomGender() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	genders := []string{"M", "F"}
	return genders[rand.Intn(len(genders))]
}

func generateRandomAadharNumber() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	aadhaar := strconv.Itoa(rand.Intn(9) + 1)
	// Add 11 more random digits (0–9)
	n := 11
	for i := 1; i < n; i++ {
		aadhaar += strconv.Itoa(rand.Intn(10))
	}
	return aadhaar
}

func GenerateRandomEmployee() []Employee {
	n := 10000
	employees := []Employee{}
	for i := range n {
		emp := Employee{
			ID:     strconv.Itoa(i),
			PAN:    generatePAN(),
			Name:   gofakeit.Name(),
			Email:  gofakeit.Email(),
			Aadhar: generateRandomAadharNumber(),
			Gender: generateRandomGender(),
		}
		employees = append(employees, emp)
	}
	return employees
}

func main() {
	mylib.Hello("Dinesh")
	// Initialize the DB with the required params...
	dblibs.DBInit(dbHost, dbPort, dbDatabaseName, dbUser, dbPassword)

}

// 🔍 Summary Table that discusses the different between the General Pseudo-Random Algorithm and Crypto Based
// Random Algorithm
//
// Feature					math/rand (General)					crypto/rand (Secure)
// Predictable				✅ Yes (if seed is known)			❌ No (very hard to predict)
// Speed					🚀 Fast								🐢 Slower
// Suitable for games?		✅ Yes								❌ Overkill
// Suitable for passwords?	❌ No								✅ Yes
// Seed required?			✅ Yes (for math/rand)				❌ No (uses system entropy)
// Source					Pseudo-random algorithm				System-level entropy
