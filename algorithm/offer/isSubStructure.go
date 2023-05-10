package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {

	if A == nil || B == nil {
		return false
	}

	return checkSubTree(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)

}

func checkSubTree(A *TreeNode, B *TreeNode) bool {
	if A == nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Val != B.Val {
		return false
	}
	return checkSubTree(A.Left, B.Left) && checkSubTree(A.Right, B.Right)
}
