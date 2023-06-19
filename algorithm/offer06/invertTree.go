package offer06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil { // 如果当前节点为空，则直接返回
		return nil
	}
	// 交换左右子节点
	root.Left, root.Right = root.Right, root.Left
	// 递归地翻转左右子树
	invertTree(root.Left)
	invertTree(root.Right)
	return root // 返回翻转后的根节点
}
