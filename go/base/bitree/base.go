package main1

import (
	"fmt"
)

// TreeNode二叉树结构
type TreeNode struct {
	Value int
	Left  *TreeNode;
	Right *TreeNode;
}

// 先序遍历
func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf(" %d ", root.Value)
	preOrderTraverse(root.Left)
	preOrderTraverse(root.Right)
}

// 先序非递归
func preOrderNoRecursion(root *TreeNode) {
	if root == nil {
		return
	}

	stack := NewQueueList()
	stack.PushBack(root)
	for !stack.Emtpy() {
		node := stack.Pop()
		fmt.Printf(" %d ", node.Value)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}

	}
}

// 中序遍历
func inOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderTraverse(root.Left)
	fmt.Printf(" %d ", root.Value)
	inOrderTraverse(root.Right)
}

//中序非递归
func inOrderNoRecursion(root *TreeNode) {
	if root == nil {
		return
	}

	stack := NewQueueList()
	stack.PushBack(root)
	for !stack.Emtpy() {
		node := stack.Pop()

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		fmt.Printf(" %d ", node.Value)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}

	}
}

//
func postOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	postOrderTraverse(root.Left)
	postOrderTraverse(root.Right)
	fmt.Printf(" %d ", root.Value)
}

//中序非递归
func postOrderNoRecursion(root *TreeNode) {
	if root == nil {
		return
	}

	stack := NewQueueList()
	stack.PushBack(root)
	for !stack.Emtpy() {
		node := stack.Pop()

		if node.Left != nil {
			stack.PushBack(node.Left)
		}

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		fmt.Printf(" %d ", node.Value)
	}
}

func levelTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	queue := NewQueueList()
	queue.PushBack(root)
	for !queue.Emtpy() {
		node := queue
	}
}

// 层序输出
func printTree(root *TreeNode) {

}

func main() {
	// 			1
	// 		2
	// 	 3
	// 4     5
	n1 := &TreeNode{Value: 1}
	n2 := &TreeNode{Value: 2}
	n3 := &TreeNode{Value: 3}
	n4 := &TreeNode{Value: 4}
	n5 := &TreeNode{Value: 5}
	n1.Left = n2
	n2.Left = n3
	n3.Left = n4
	n3.Right = n5
	preOrderTraverse(n1)
	fmt.Println()
	preOrderNoRecursion(n1)
	fmt.Println()
	inOrderTraverse(n1)
	fmt.Println()
	postOrderTraverse(n1)
	fmt.Println()
	inOrderNoRecursion(n1)
	fmt.Println()
	postOrderNoRecursion(n1)
	fmt.Println()

}
