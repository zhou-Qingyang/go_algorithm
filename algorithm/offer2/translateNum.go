package offer02

import "strconv"

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i > 1 {
			pre := s[i-2 : i]
			if pre >= "10" && pre <= "25" {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n]
}
