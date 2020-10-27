package main

import (
	"fmt"
	"math"
)

func isDiff(a int, b int) bool {
	if (a < 0 && b > 0) || (a > 0 && b < 0) {
		return true
	}
	return false
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 1 {
		return dividend
	}

	isNegative := isDiff(dividend, divisor)

	tmp := dividend
	res := 0
	for true {
		minus := divisor
		if isNegative {
			minus = -minus
		}
		tmp = tmp - minus

		if isDiff(tmp, dividend) {
			break
		}
		res ++
	}
	if isNegative {
		res = -res
	}
	return res
}

func main() {
	// l1 := &ListNode{Val:1, Next:nil}
	// l2 := &ListNode{Val:2, Next: nil}
	// l3 := &ListNode{Val:3, Next: nil}
	// l4 := &ListNode{Val:4, Next: nil}
	// l1.Next = l2
	// l2.Next = l3
	// l3.Next = l4
	fmt.Println(divide(-2147483648,-1))

}
