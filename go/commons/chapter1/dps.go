package main

import (
	"fmt"
)

/**
动态规划问题的三个性质: 1. 重叠子问题 2. 状态转义方程，解决dp问题的核心 3. 最优子结构，子问题必须相互独立

*/

// 递归树可以看出，算法复杂度为O(2^n)
func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 带备忘录版本的, 算法复杂度 O(n)，算法效率跟迭代的动态规划差不多，一个自顶向下，一个自底向上
func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	// var memo = [5]int{0}
	memo := make([]int, n+1)
	return memoHelper(memo, n)
}
func memoHelper(memo []int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = memoHelper(memo, n-1) + memoHelper(memo, n-2)
	return memo[n]
}

// dp数组的迭代解法
func fib3(n int) int {
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}


//
func coinChange2(coins []int, amount int) int {
	memo := map[int]int{}
	return coinChangeDp(memo, coins, amount)
}

func min(a int, b int) int {
	if a < b  {
		return a
	} else {
		return b
	}
}

func coinChangeDp(memo map[int]int, coins []int, amount int) int {
	if num, exists := memo[amount]; exists {
		return  num
	}

	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	} else {
		res := -1
		for _, coin := range coins {
			subpro := coinChangeDp(memo, coins, amount - coin)
			if subpro < 0 {
				continue
			}

			if res == -1 {
				res = 1 + coinChangeDp(memo, coins, amount - coin)
			} else {
				res = min(res, 1 + coinChangeDp(memo, coins, amount - coin))
			}
		}
		memo[amount] = res
		return memo[amount]
	}
}

func coinChange(coins []int , amount int) int {
	memo := map[int]int{}
	for i := 0; i <= amount; i++ {
		memo[i] = -1
	}
	for i := 1 ;i <= amount; i++ {
		for _, coin := range coins {
			if i - coin < 0 {
				continue
			}

			if memo[i] == -1 {
				memo[i] = 1+memo[i-coin]
			} else {
				memo[i] = min(memo[i], 1+ memo[i-coin])
			}
		}
	}

	return memo[amount]
}

func main() {
	fmt.Println("fib1 ", fib1(20))
	fmt.Println("fib2", fib2(20))
	fmt.Println("fib3", fib3(20))
	fmt.Println("coinChange2", coinChange2([]int{1, 2, 5}, 11))
	fmt.Println("coinChangedp : ", coinChange([]int{1, 2, 3}, 11))
	fmt.Println("hello world")
}
