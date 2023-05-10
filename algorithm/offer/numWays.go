package offer

//11
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 || n == 2 {
		return n
	}
	const mod = 1000000007
	f0, f1 := 1, 2
	for i := 3; i <= n; i++ {
		f2 := (f0 + f1) % mod
		f0 = f1
		f1 = f2
	}
	return f1

}
