package all

import "fmt"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 * 1. 1单独处理
 * 2. 遇到边界的处理流程
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	lenStr := len(s)
	fmt.Println(lenStr)
	arr := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]byte, lenStr)
	}

	forwardType := 1
	for k, i, j := 0, 0, 0; k < lenStr; k++ {
		arr[i][j] = s[k]
		if (numRows) == 1 {
			j++
		} else if forwardType == 1 {
			if i >= numRows-1 {
				forwardType = 2
				i--
				j++
			} else {
				i++
			}
		} else {
			if i <= 0 {
				forwardType = 1
				i++
			} else {
				i--
				j++
			}
		}
	}

	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < lenStr; j++ {
			fmt.Print(" ", arr[i][j])
			if arr[i][j] != 0 {
				res += string(arr[i][j])
			}
		}
		fmt.Println()
	}

	return res
}

// @lc code=end


