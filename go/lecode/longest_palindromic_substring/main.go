package main

import "fmt"

/**
1. 特殊case ""  "a" "aa"
2. 有没有更好的方法
3. golang 单步调试
4.
 */

func isPalindrom(s string) bool {
	if len(s) <= 1{
		return true
	}

	for i, j := 0, len(s)-1; i < j;  i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}


func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	res := s[0:1]
	for maxLen := 1; maxLen < len(s) ; maxLen++ {
		//fmt.Println(maxLen)
		for i := 0; i < len(s)-maxLen  ; i++ {
			if isPalindrom(s[i: i+maxLen+1]) {
				res = s[i:i+maxLen+1]
				break
			}
		}

	}


	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}