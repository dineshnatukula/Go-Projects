package strings

import "fmt"

func GetSingleDigit(n int) int {
	sum := 0
	var getSingleDigitRecur func(n int) int
	getSingleDigitRecur = func(n int) int {
		if n <= 9 {
			return n
		}
		for ; n > 0; n = n / 10 {
			t := n % 10
			sum = sum + t
		}
		n = sum
		sum = 0
		return getSingleDigitRecur(n)
	}

	sum = getSingleDigitRecur(n)
	fmt.Println("Dgit sum < 0 is ", sum)
	return sum
}

func ExcelColumnTitle(n int) (s string) {
	result := ""
	for n > 0 {
		n-- // Shift to 0-based index
		// remainder := n % 26
		// result = string(65+remainder) + result
		n /= 26
	}
	return result
}
