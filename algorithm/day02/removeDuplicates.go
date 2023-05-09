package day02

//26
func RemoveDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 10)
	slow, fast := 0, 1

	for fast < len(nums) {
		if nums[fast-1] != nums[fast] {
			result[slow] = nums[fast-1]
			slow++
		}
		fast++
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		result[slow] = nums[len(nums)-1]
	}
	return result
}
