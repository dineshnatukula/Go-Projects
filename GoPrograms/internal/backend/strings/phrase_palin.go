package strings

// import (
// 	"regexp"
// 	"strings"
// 	"unicode"
// )

// func IsPhrasePalindrome1(s string) bool {
// 	var cleaned []rune

// 	// Clean the input: remove non-letter/digit characters and convert to lowercase
// 	for _, r := range s {
// 		if unicode.IsLetter(r) || unicode.IsDigit(r) {
// 			cleaned = append(cleaned, unicode.ToLower(r))
// 		}
// 	}

// 	// Compare characters from both ends
// 	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
// 		if cleaned[i] != cleaned[j] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func IsPhrasePalindrome2(s string) bool {
// 	// Keep only letters and digits
// 	re := regexp.MustCompile("[^a-zA-Z0-9]+")
// 	cleaned := re.ReplaceAllString(s, "")
// 	cleaned = strings.ToLower(cleaned)

// 	// Reverse and compare
// 	reversed := reverseString(cleaned)
// 	return cleaned == reversed
// }

// func IsPhrasePalindrome3(s string) bool {
// 	runes := []rune(s)
// 	left, right := 0, len(runes)-1

// 	for left < right {
// 		// Skip non-alphanumeric characters
// 		for left < right && !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
// 			left++
// 		}
// 		for left < right && !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
// 			right--
// 		}

// 		// Compare characters case-insensitively
// 		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
// 			return false
// 		}

// 		left++
// 		right--
// 	}

// 	return true
// }
