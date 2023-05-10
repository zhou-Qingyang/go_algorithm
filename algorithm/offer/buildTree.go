package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 07
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	rootValue := preorder[0]
	root := &TreeNode{rootValue, nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootValue {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
