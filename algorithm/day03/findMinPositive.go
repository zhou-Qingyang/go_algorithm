package day03

//41
func FindMinPositive(nums []int) int {
	n := len(nums)

	// 将小于等于数组长度的正整数放入对应位置上
	for i := 0; i < n; i++ {
		// 只处理小于等于数组长度的数
		for 1 <= nums[i] && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// 交换 nums[i] 和 nums[nums[i]-1] 的位置
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 找到第一个没有被放在正确位置上的数字的下标
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 如果数组中所有小于等于数组长度的数都在正确位置上，那么未出现的最小正整数为 n+1
	return n + 1
}
