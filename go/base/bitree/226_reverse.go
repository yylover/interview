package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	node := root.Left
	root.Left = root.Right
	root.Right = node
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}


