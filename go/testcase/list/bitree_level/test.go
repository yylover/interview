package main

import "fmt"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

type queue []*TreeNode
func (q queue) insert(node *TreeNode) {
	q = append(q, node)
}
func (q queue) out(node *TreeNode) *TreeNode {
	if len(q) == 0 {
		return nil
	}
	node = q[0]
	q = q[1:]
	return node
}
func (q queue)print() {
	fmt.Println("len : ", len(q))
}

func levelPrint(node *TreeNode) []int {

	return []int{}
}

func main() {
	var q queue
	q.insert(&TreeNode{1, nil, nil})
	q.insert(&TreeNode{2, nil, nil})
	q.

}