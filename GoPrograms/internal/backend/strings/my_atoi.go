package strings

import "math"

func MyAtoi(s string) (res int) {
	i := 0
	n := len(s)

	for i < n && s[i] == ' ' {
		i += 1
	}
	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	res = 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		// Check overflow before adding the digit
		if res > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + digit
		i++
	}
	return res * sign
}
