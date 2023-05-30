package offer02

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				res[0][0] = grid[0][0]
			} else if i == 0 {
				res[i][j] = res[i][j-1] + grid[i][j]
			} else if j == 0 {
				res[i][j] = res[i-1][j] + grid[i][j]
			} else {
				res[i][j] = max(res[i-1][j], res[i][j-1]) + grid[i][j]
			}
		}

	}
	return res[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
