package all



import "fmt"

/**
1. 终止条件
2. 测试
""， a, aa, aabb, abbc, abbbc,
 */

func isPalindrom(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left ++
		right --
	}
	return true
}

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxLen := len(s)
	for maxLen > 1 {
		// TODO 这里的终止条件要注意
		for i := 0; i + maxLen -1 < len(s); i++ {
			if isPalindrom(s[i: i+maxLen]) {
				return s[i: i+maxLen]
			}
		}
		maxLen --
	}

	return s[0:1]
}
// @lc code=end

func main() {
	str := "bb"
	fmt.Println("longtest : ", longestPalindrome(str))
}
