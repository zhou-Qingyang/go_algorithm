package offer3

func lastRemaining(n int, m int) int {
	var ans int
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}
