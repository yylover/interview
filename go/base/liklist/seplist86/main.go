package main

import (
	. "github.com/yylover/interview/go/base/liklist/listdef"
)

/**
1. 如何保证原来的顺序
2. 在原来的表基础上调整，不如直接构造新的表再组装
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 *
 */

func partition(head *ListNode, x int) *ListNode {
	//判断
	if head == nil {
		return nil
	}

	//寻找对应的节点
	var lessPtr *ListNode = nil
	var lessNextPtr *ListNode = nil
	var greaterPtr *ListNode = nil
	var greaterNextPtr *ListNode = nil


	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < x {
			if lessPtr == nil {
				lessPtr = head
				lessNextPtr = head
			} else {
				lessNextPtr.Next = head
				lessNextPtr = head
			}
		} else {
			if greaterPtr == nil {
				greaterPtr = head
				greaterNextPtr = head
			} else {
				greaterNextPtr.Next = head
				greaterNextPtr = head
			}
		}
		head = next
	}

	resPtr := lessPtr
	nextPtr := lessNextPtr
	if resPtr == nil {
		resPtr = greaterPtr
		nextPtr = greaterNextPtr
	} else if greaterPtr != nil {
		nextPtr.Next = greaterPtr
	}


	return resPtr
}

//跑题了，此解法是分割链表，左边小于x, 右边大于x
func partition2(head *ListNode, x int) *ListNode {
	//判断
	if head == nil {
		return nil
	}

	//寻找对应的节点
	var lessPtr *ListNode = nil
	var lessNextPtr *ListNode = nil
	var greaterPtr *ListNode = nil
	var greaterNextPtr *ListNode = nil
	var equPtr *ListNode = nil
	var equNextPtr *ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val > x {
			if greaterPtr == nil {
				greaterPtr = head
				greaterNextPtr = head
			} else {
				greaterNextPtr.Next = head
				greaterNextPtr = head
			}
		} else if head.Val == x {
			if equPtr == nil {
				equPtr = head
				equNextPtr = head
			} else {
				equNextPtr.Next = head
				equNextPtr = head
			}
		} else {
			if lessPtr == nil {
				lessPtr = head
				lessNextPtr = head
			} else {
				lessNextPtr.Next = head
				lessNextPtr = head
			}
		}
		head = next
	}

	resPtr := lessPtr
	nextPtr := lessNextPtr
	if resPtr == nil {
		resPtr = equNextPtr
		nextPtr = equNextPtr
	} else if equPtr != nil {
		nextPtr.Next = equPtr
		nextPtr = equNextPtr
	}

	if resPtr == nil {
		resPtr = greaterPtr
		nextPtr = greaterNextPtr
	} else {
		nextPtr.Next = greaterPtr
		nextPtr = greaterNextPtr
	}

	return resPtr
}

func main() {
	node := GetList([]int{1, 4, 3, 2, 5, 2})
	PrintList(node)
	PrintList(partition(node, 3))

	println("hello")
}

func partition1(head *ListNode, x int) *ListNode {
	//判断
	if head == nil {
		return nil
	}

	//寻找对应的节点
	tarptr := head
	for tarptr != nil && tarptr.Val != x {
		tarptr = tarptr.Next
	}
	if tarptr == nil {
		return nil
	}

	back := head
	front := tarptr
	for back != tarptr {
		if back.Val <= tarptr.Val {
			continue
		}
		if back == head {
			head = back.Next
		}
		front.Next = back
		front = front.Next
		back = head
	}

	return nil
}
