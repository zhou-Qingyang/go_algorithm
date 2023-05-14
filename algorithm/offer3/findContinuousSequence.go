package offer3

func findContinuousSequence(target int) [][]int {
	var res [][]int
	i, j, sum := 1, 1, 0
	for i <= target/2 {
		if sum < target {
			sum += j
			j++
		} else if sum > target {
			sum -= i
			i++
		} else {
			var arr []int
			for k := i; k < j; k++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
			sum -= i
			i++
		}
	}
	return res
}
