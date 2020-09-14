package main

import "fmt"

const BucketNum = 10

type ListNode struct {
	Val       int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func insert(head *ListNode, val int) {
	node := NewListNode(val)
	var pre, cur *ListNode = head, head.Next
	for cur != nil && cur.Val < val {
		pre = cur
		cur = cur.Next
	}

	pre.Next =  node
	node.Next = cur
}

func merge(head1 *ListNode, head2 *ListNode) *ListNode {
	var dummyNode *ListNode = NewListNode(0)
	var node = dummyNode
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			node.Next = head1
			node = head1
			head1 = head1.Next
		} else {
			node.Next = head2
			node = head2
			head2 = head2.Next
		}
	}

	if head1 != nil {
		node.Next = head1
	}

	if head2 != nil {
		node.Next = head2
	}
	return dummyNode.Next
}

func bucketSort(arr []int) []int {
	bucket := [BucketNum]*ListNode{}

	for i := 0; i < BucketNum; i++ {
		bucket[i] = NewListNode(0)
	}

	for _, item := range arr {
		bucketIndex := item % BucketNum
		insert(bucket[bucketIndex], item)
		//插入排序
	}

	//merge
	head := bucket[0].Next
	for i :=1 ; i < BucketNum; i++ {
		head = merge(head, bucket[i].Next)
		// fmt.Printf("%d ", i)
	}

	arr = []int{}
	for node := head; node != nil; node = node.Next {
		// fmt.Printf("%d ", node.Val)
		arr = append(arr, node.Val)
	}
	return arr
}

func main() {
	arr := []int{3,2,1,4,8,7,5,6, 15, 24, 14, 21}
	arr = bucketSort(arr)
	fmt.Println("sort : ", arr)
}
