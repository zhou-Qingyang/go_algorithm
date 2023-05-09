package day03

//45
func jump(nums []int) int {

	n := len(nums)
	if n == 1 {
		return 0
	}

	start, end, maxPos, step := 0, 0, 0, 0
	for start <= end {
		for i := start; i < end; i++ {
			if nums[i]+i > maxPos {
				maxPos = nums[i] + i
			}
		}
		start, end = end+1, maxPos
		step++
		if end >= n-1 {
			return step
		}
	}
	return -1

}
