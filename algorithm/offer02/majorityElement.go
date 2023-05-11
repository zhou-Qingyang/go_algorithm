package offer02

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	res := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if res == nums[i] {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				res = nums[i]
				cnt = 1
			}

		}
	}
	return res
}
