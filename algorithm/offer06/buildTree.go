package offer06

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	indexMap := make(map[int]int)
	for i, v := range inorder {
		indexMap[v] = i
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, indexMap)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, indexMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]   //3
	rootIndex := indexMap[rootVal]  //1
	leftSize := rootIndex - inStart // 1
	rightSize := inEnd - rootIndex  // 3

	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}

	//神奇
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, rootIndex-1, indexMap)
	root.Right = build(preorder, preEnd-rightSize+1, preEnd, inorder, rootIndex+1, inEnd, indexMap)

	return root
}
