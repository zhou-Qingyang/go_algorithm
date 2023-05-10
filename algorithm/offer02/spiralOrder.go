package offer02

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	index := 0
	top, bottom, left, right := 0, m-1, 0, n-1
	for top <= bottom && left <= right {
		// 从左到右遍历上边界
		for i := left; i <= right; i++ {
			res[index] = matrix[top][i]
			index++
		}
		// 从上到下遍历右边界
		for i := top + 1; i <= bottom; i++ {
			res[index] = matrix[i][right]
			index++
		}
		// 如果上下边界重合，则已经遍历完所有行和列，退出循环
		if top < bottom && left < right {
			// 从右到左遍历下边界
			for i := right - 1; i > left; i-- {
				res[index] = matrix[bottom][i]
				index++
			}
			// 从下到上遍历左边界
			for i := bottom; i > top; i-- {
				res[index] = matrix[i][left]
				index++
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return res

}
