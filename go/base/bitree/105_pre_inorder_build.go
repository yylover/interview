package main



func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}
	//find 中序index
	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	root.Left = buildTree(preorder[1:index+1], inorder[0:index])
	root.Right = buildTree(preorder[index+1:len(preorder)], inorder[index+1:len(preorder)])

	return root
}


func main() {
	root := buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	PrePrintTree(root)
}
