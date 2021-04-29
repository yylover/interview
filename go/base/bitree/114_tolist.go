package main

import "fmt"

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	tmp := &TreeNode{}
	connectTree(root, &tmp)
}

func connectTree(root *TreeNode, last **TreeNode)  {
 	if root == nil {
 		return
	}
	*last = root
	rightLef := root.Right
	leftLast := root
	if root.Left != nil {
		connectTree(root.Left, last)
		tmp := root.Left
		root.Left = nil
		root.Right = tmp
		leftLast = *last
	}

	if rightLef != nil {
		connectTree(rightLef, last)
		leftLast.Right = rightLef
	}
}

func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	flatten2(root.Left)
	flatten2(root.Right)
	right := root.Right
	root.Right = root.Left
	root.Left =  nil

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}


	PrePrintTree(root)
	fmt.Println()
	InorderPrintTree(root)
	fmt.Println()

	flatten2(root)

	PrePrintTree(root)
	fmt.Println()
	InorderPrintTree(root)
	fmt.Println()
}