package all

/**
1. 需要有三个节点
2. 边界条件的处理
3.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	front := head.Next
	back := head
	var headRes *ListNode = nil
	var pre *ListNode = nil

	for front != nil && back != nil {
		if headRes == nil {
			headRes = front
		}

		if pre != nil {
			pre.Next = front
		}
		back.Next = front.Next
		front.Next = back
		if back.Next == nil {
			break
		}
		pre = back
		back = back.Next
		front = back.Next
	}
	return headRes
}
