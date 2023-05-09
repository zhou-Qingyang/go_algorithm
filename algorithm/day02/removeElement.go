package day02

//27
func removeElement(nums []int, val int) []int {
	if len(nums) == 0 {
		return nil
	}

	slow := 0

	result := make([]int, 10)

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			result[slow] = nums[fast]
			slow++
		}
	}

	return result
}
