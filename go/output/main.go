package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)


func getTodayUnix() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findSmallList(lists *[]*ListNode) *ListNode {
	var ptr **ListNode = nil
	for _, node := range *lists {
		if node != nil {
			if node == nil || node.Val < *ptr.Val {
				*ptr = node
			}
		}
	}
	*ptr = *ptr.Next
	return ptr
}
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	nextPtr := head
	node := findSmallList(lists)
	for node != nil {
		nextPtr.Next = node
		nextPtr = nextPtr.Next
		node = findSmallList(lists)
	}
	return head.Next
}

// @lc code=end




func main() {

	fmt.Println(getTodayUnix())
	return ;

	file, err := os.Open("/Users/yangqiang/Downloads/testfile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		datas := strings.Split(scanner.Text(), "|")
		if len(datas) != 40 {
			fmt.Println(datas)
		}
		fmt.Println(len(datas))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}