package strings

import "fmt"

func IsMatch(s, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* at the beginning
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' && dp[0][j-2] {
			dp[0][j] = true
		}
	}
	fmt.Println("After Initialization here")
	fmt.Println(dp)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] ||
					((p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j])
			}
			fmt.Println("i =", i, "j =", j, "s", s, "p", p)
			fmt.Println(dp)
		}
	}

	return dp[m][n]
}
