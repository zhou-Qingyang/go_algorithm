package offer4

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}
