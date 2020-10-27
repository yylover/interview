package main

import (
	"fmt"
)

func loopFunc(targetNum int) {
	stack := []string{}
	for i := 0; i < targetNum; i++ {
		tmp := []string{"A", "B"}
		for j := 0; j < len(tmp); j++ {
			tmp
		}
	}
}

func innerFunc(str string, targetNum int, res *[]string) {
	if len(str) == targetNum {
		*res = append(*res, str)
		return
	}

	tmp := []string{"A", "B"}
	for j := 0; j < len(tmp); j++ {
		innerFunc(str+tmp[j], targetNum, res)
	}
}

func loopNums() []string {
	res := []string{}
	innerFunc("", 2, &res)
	return res
}

// @lc code=end

func main() {
	// str := "LEETCODEISHIRING"
	fmt.Println("", loopNums())
}
