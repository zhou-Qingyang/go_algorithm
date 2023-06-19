package offer06

//39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var track []int // 记录当前路径
	backtrack(candidates, target, 0, 0, &track, &res)
	return res
}

func backtrack(candidates []int, target int, start int, sum int, track *[]int, res *[][]int) {
	if sum == target { // 找到一组符合要求的组合
		temp := make([]int, len(*track))
		copy(temp, *track)
		*res = append(*res, temp)
		return
	}
	if sum > target || start >= len(candidates) { // 当前路径无法满足要求，剪枝
		return
	}
	// 选择当前元素
	*track = append(*track, candidates[start])
	backtrack(candidates, target, start, sum+candidates[start], track, res)
	*track = (*track)[:len(*track)-1]
	// 不选择当前元素
	backtrack(candidates, target, start+1, sum, track, res)
}
