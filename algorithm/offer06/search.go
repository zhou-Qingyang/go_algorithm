package offer06

//33. 搜索旋转排序数组
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	//确定 最小位置
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	minIndex := left
	if target >= nums[minIndex] && target <= nums[n-1] {
		left, right = minIndex, n-1
	} else {
		left, right = 0, minIndex-1
	}
	for left <= right {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
