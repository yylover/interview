package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func ListSort(head *ListNode) {
	var node *ListNode = nil
	var minNode *ListNode = nil
	for head != nil && head.Next != nil {
		minNode = head
		node = head.Next

	}
	print(node)
	print(minNode)
}

func PrintList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return l[i] < l[j]
}

// Swap swaps the elements with indexes i and j.
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func levelPrint(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TempNode{}
	preOrderTreeNode(&queue, root)

	sort.Sort(TempNodeList(queue))

	res := []int{}
	for _, item := range queue {
		fmt.Printf("node :%d %d ", item.Node.Value, item.Height)
		res = append(res, item.Node.Value)
	}

	return res
}

type TempNode struct {
	Node   *TreeNode
	Height int
}

type TempNodeList []*TempNode

func (l TempNodeList) Len() int {
	return len(l)
}

func (l TempNodeList) Less(i, j int) bool {
	return l[i].Height < l[j].Height
}

// Swap swaps the elements with indexes i and j.
func (l TempNodeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func preOrderTreeNode(list *[]*TempNode, node *TreeNode) int {
	if node == nil {
		return 0
	}

	n := &TempNode{Node: node, Height: 0}
	*list = append(*list, n)
	left := preOrderTreeNode(list, node.Left)
	right := preOrderTreeNode(list, node.Right)
	height := 0
	if left > right {
		height = left + 1
	} else {
		height = right + 1
	}

	n.Height = height
	return height
}

func getNearlestSequence(datas []int, target int) []int {
	res := []int{}
	if len(datas) < 3 {
		return res
	}

	// sort.Sort(datas)

	return res
}

func main() {
	fmt.Println("hello")
	// node := &ListNode{Value: 2, Next: nil}
	// node2 := &ListNode{Value: 2, Next: node}
	// node3 := &ListNode{Value: 3, Next: node2}
	// head := &ListNode{Value: 4, Next: node3}
	// PrintList(head)

	// 		    1
	// 		 2   5
	// 	  3
	// 4
	n1 := TreeNode{1, nil, nil}
	n2 := TreeNode{2, nil, nil}
	n3 := TreeNode{3, nil, nil}
	n4 := TreeNode{4, nil, nil}
	n5 := TreeNode{5, nil, nil}
	n1.Left = &n2
	n2.Left = &n3
	n3.Left = &n4
	n1.Right = &n5

	fmt.Println(levelPrint(&n1))
	// list := []*TempNode{}
	// preOrderTreeNode(list, &n1)
}
