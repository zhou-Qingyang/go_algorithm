package offer

//03
func findRepeatNumber(nums []int) int {
	// [2, 3, 1, 0, 2, 5, 3]
	// 输出：2 或 3
	// 定义哈希表 map
	m := make(map[int]bool)

	// 遍历数组 nums 中的每个元素
	for _, num := range nums {
		// 检查当前元素是否已经存在于 map 中，如果是则返回该数字
		if m[num] {
			return num
		}
		// 将当前元素加入到哈希表 map 中
		m[num] = true
	}

	// 如果遍历完整个数组 nums 后仍未找到重复数字，则返回 -1
	return -1
}
