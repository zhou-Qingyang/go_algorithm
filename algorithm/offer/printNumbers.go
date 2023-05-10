package offer

func printNumbers(n int) []int {
	if n == 0 {
		return []int{}
	}
	lens := 1
	for n >= 1 {
		lens *= 10
		n--
	}
	res := make([]int, lens-1)
	for i := 0; i < lens-1; i++ {
		res[i] = i + 1
	}
	return res
}
