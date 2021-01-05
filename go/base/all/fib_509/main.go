package main

import "fmt"

/**
递归和循环
 */


func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	sum := 0
	lasta := 0
	lastb := 1
	for i := 2; i <= n; i++ {
		sum = lasta+lastb
		lastb, lasta = sum, lastb
	}
	return sum
}

func main() {
	fmt.Println(fib2(4))
}