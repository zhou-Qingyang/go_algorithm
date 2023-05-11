package offer02

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	res, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if sum > res {
			res = sum
		}
	}
	return res
}
