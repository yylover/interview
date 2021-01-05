package listdef

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func GetList(vals []int) *ListNode {
	var head *ListNode = nil
	var next *ListNode = nil
	for _, val := range vals {
		node := &ListNode{val, nil}
		if head == nil {
			head = node
			next = node
		} else {
			next.Next = node
			next = node
		}
	}
	return head
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}