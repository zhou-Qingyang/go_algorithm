package offer

func exchange(nums []int) []int {
	n, m := 0, len(nums)-1
	for n < m {
		for nums[n]%2 == 1 {
			n++
		}
		for nums[m]%2 == 0 {
			m--
		}
		nums[n], nums[m] = nums[m], nums[n]
	}
	return nums
}
