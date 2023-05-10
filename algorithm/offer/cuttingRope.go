package offer

//14
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	res := make([]int, n+1)
	res[2] = 2
	res[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i-1; j++ {
			res[i] = max(res[i], max(j, res[j])*max(i-j, res[i-j]))
		}
	}
	return res[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
