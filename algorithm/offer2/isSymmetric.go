package offer02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// 如果根节点为空，则返回 true
	if root == nil {
		return true
	}
	// 递归地比较左右子树是否对称
	return checkSymmetry(root.Left, root.Right)
}
func checkSymmetry(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	// 递归比较 p 的左子树和 q 的右子树是否对称，以及 p 的右子树和 q 的左子树是否对称
	return checkSymmetry(A.Left, B.Right) && checkSymmetry(A.Right, B.Left)
}
