package offer

// 10
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	f0, f1 := 0, 1
	for i := 2; i <= n; i++ {
		f2 := (f0 + f1) % 1000000007
		f0 = f1
		f1 = f2
	}
	return f1

}
