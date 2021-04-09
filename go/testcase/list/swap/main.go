package main

type ListNode struct {
	Val  int
	Next *ListNode;
}

func PrintListNode(node *ListNode) {
	for node != nil {
		print(node.Val, " ")
		node = node.Next
	}
	println()
}

func GetListNode(n int) *ListNode {
	var head *ListNode = nil
	var node *ListNode = nil
	for i := 1; i <= n; i++ {
		tmp := &ListNode{i, nil}
		if head == nil {
			head = tmp
			node = tmp
		} else {
			node.Next = tmp
			node = node.Next
		}
	}
	return head
}

func ReverseListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var dummy *ListNode = &ListNode{0, nil}
	dummy.Next = head
	var after *ListNode = dummy
	var pre *ListNode = after.Next
	for pre != nil && pre.Next != nil {
		next := pre.Next.Next
		after.Next = pre.Next
		pre.Next.Next = pre
		pre.Next = next
		after = pre
		pre = next
	}

	return dummy.Next
}

func main() {
	n1 := GetListNode(2)
	PrintListNode(n1)
	n1 = ReverseListNode(n1)
	PrintListNode(n1)
}
