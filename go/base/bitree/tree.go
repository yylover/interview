package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrePrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val , " ")
	PrePrintTree(root.Left)
	PrePrintTree(root.Right)
}

func InorderPrintTree(root * TreeNode) {
	if root == nil {
		return
	}

	InorderPrintTree(root.Left)
	fmt.Print(root.Val , " ")
	InorderPrintTree(root.Right)
}