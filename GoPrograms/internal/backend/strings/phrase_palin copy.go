package strings

import (
	"regexp"
	"strings"
	"unicode"
)

// Use rune to handle the indv chars of the string
func IsPhrasePalindrome1(s string) bool {
	var cleaned []rune

	for _, v := range s {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			cleaned = append(cleaned, unicode.ToLower(v))
		}
	}

	for i, j := 0, len(cleaned)-1; i < len(cleaned); i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}

	return true
}

// reverse(s) returns the reverse of the string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindromicPhrase finds if the phrase is palindrome
// using the regular expressions. This is used to remove the spl chars
func IsPhrasePalindrome2(s string) bool {
	re := regexp.MustCompile("[a-zA-Z0-9]+")
	cleaned := re.ReplaceAllString(s, "")
	cleaned = strings.ToLower(cleaned)

	// Reverse the cleaned string and compare
	reverse := ReverseString(cleaned)
	return cleaned == reverse
}

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
