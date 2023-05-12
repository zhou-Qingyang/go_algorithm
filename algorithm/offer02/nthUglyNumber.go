package offer02

func nthUglyNumber(n int) int {
	res := make([]int, n)
	res[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 0; i < n; i++ {
		res[i] = min(res[p2]*2, min(res[p3]*3, res[p5]*5))
		if res[i] == res[p2]*2 {
			p2++
		}
		if res[i] == res[p3]*3 {
			p3++
		}
		if res[i] == res[p5]*5 {
			p5++
		}
	}
	return res[n-1]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
