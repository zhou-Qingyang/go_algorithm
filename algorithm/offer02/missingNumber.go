package offer02

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}
	res := make([]int, len(nums))
	// 	输入: [0,1,2,3,4,5,6,7,9]
	// 输出: 8
	for i := 0; i < len(nums); i++ {
		res[i] = i
	}
	for j, v := range res {
		if nums[j] != v {
			return v
		}
	}
	return nums[len(nums)-1] + 1
}
