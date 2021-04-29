package main


func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	//find 中序index
	index := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			index = i
			break
		}
	}

	root.Left = buildTree2(inorder[0:index], postorder[0:index])
	root.Right = buildTree2(inorder[index+1:len(inorder)], postorder[index:len(inorder)-1])

	return root
}


func main() {
	root := buildTree2([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	PrePrintTree(root)
}
