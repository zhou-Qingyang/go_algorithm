package offer06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 337. 打家劫舍 III
func rob(root *TreeNode) int {
	cache := make(map[*TreeNode]int)
	return dp(root, cache)
}
func dp(root *TreeNode, cache map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := cache[root]; ok {
		return v
	}
	val1 := root.Val
	if root.Left != nil {
		val1 += dp(root.Left.Left, cache) + dp(root.Left.Right, cache)
	}
	if root.Right != nil {
		val1 += dp(root.Right.Left, cache) + dp(root.Right.Right, cache)
	}
	val2 := dp(root.Left, cache) + dp(root.Right, cache)
	res := max(val1, val2)
	cache[root] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
