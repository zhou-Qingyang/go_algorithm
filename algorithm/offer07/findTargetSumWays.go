package offer07

//494. 目标和
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 假设 sum 为所有元素的和，那么 S1 - S2 = target，
	// 并且 S1 + S2 = sum，则 S1 = (target + sum) / 2
	if (target+sum)%2 != 0 {
		return 0
	}

	// 目标值
	targetSum := (target + sum) / 2

	// 创建一个一维数组 dp，大小为 targetSum+1，并初始化为0
	dp := make([]int, targetSum+1)
	dp[0] = 1

	for _, num := range nums {
		// 注意这里需要从后向前遍历，确保每个元素只被计算一次
		for i := targetSum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	return dp[targetSum]
}
