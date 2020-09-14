package main

import (
	"fmt"
)

// TreeNode二叉树结构


// 先序遍历
func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf(" %d ", root.Val)
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
		fmt.Printf(" %d ", node.Val)
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
	fmt.Printf(" %d ", root.Val)
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
		fmt.Printf(" %d ", node.Val)
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
	fmt.Printf(" %d ", root.Val)
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
		fmt.Printf(" %d ", node.Val)
	}
}

func levelTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	queue := NewQueueList()
	queue.PushBack(root)
	for !queue.Emtpy() {
		for i := 0; i < queue.Len; i++ {
			node := queue.Pop()
			fmt.Printf(" %d ", node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
// 二叉树层序遍历,slice 操作头部插入元素，slice动态遍历长度计算
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := NewQueueList()
	queue.PushBack(root)
	for !queue.Emtpy() {
		tmp := []int{}
		len := queue.Len
		for i := 0; i < len; i++ { // bug
			node := queue.Pop()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		//res =  append(res, tmp)

		res = append([][]int{tmp}, res...)
	}

	return res
}

// 层序输出
func printTree(root *TreeNode) {

}

func main() {
	// 			1
	// 		2
	// 	 3
	// 4     5
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
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
	levelTraverse(n1)
	fmt.Println()
	fmt.Printf("%v \n", levelOrderBottom(n1))
}
