package offer06

// 55. 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	maxPosi := 0
	for i := 0; i < n; i++ {
		if maxPosi < i {
			return false
		}
		maxPosi = max(maxPosi, i+nums[i])
	}
	return maxPosi >= n-1
}
