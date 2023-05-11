package offer02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	root := postorder[len(postorder)-1]
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] > root {
			break
		}
	}
	j := i
	for ; j < len(postorder)-1; j++ {
		if postorder[j] < root {
			return false
		}
	}
	left := true
	if i > 0 {
		left = verifyPostorder(postorder[:i])
	}
	right := true
	if j < len(postorder)-1 {
		right = verifyPostorder(postorder[i : len(postorder)-1])
	}
	return left && right

}
