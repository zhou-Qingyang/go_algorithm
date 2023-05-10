package offer

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := 1.0
	for n != 0 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
		n /= 2
	}
	return res
}
