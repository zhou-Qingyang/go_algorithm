package offer02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {

	var res []int
	kthToArray(root, &res)
	return res[len(res)-k]
}

func kthToArray(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	kthToArray(root.Left, arr)
	*arr = append(*arr, root.Val)
	kthToArray(root.Right, arr)
}
