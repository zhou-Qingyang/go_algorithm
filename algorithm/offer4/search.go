package offer4

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			i, j := mid, mid
			for i >= 0 && nums[i] == target {
				i--
			}
			for j < len(nums) && nums[j] == target {
				j++
			}
			return j - i - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
