package day02

//31
func NextPermutation(nums []int) {
	n := len(nums)
	i, j := n-2, n-1

	//第一个不是升序的地方
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		//找到第一个 j大于i 的数字
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	left, right := i+1, n-1
	//从n-2 到i 都是升序
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
