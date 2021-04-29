package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	connectTreeFunc(root.Left, root.Right)
	return root
}

func connectTreeFunc(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	connectTreeFunc(left.Left, left.Right)
	connectTreeFunc(left.Right, right.Left)
	connectTreeFunc(right.Left, right.Right)
}
