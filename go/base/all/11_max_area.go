package all

import (
	"fmt"
)

func getArea(left, right, leftHeight, rightHeight int) int {
	height := leftHeight
	if leftHeight > rightHeight {
		height = rightHeight
	}
	return height * (right - left)
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := getArea(i, j, height[i], height[j])
			if area > max {
				max = area
			}
		}
	}
	return max
}

// @lc code=end

func main() {
	// str := "LEETCODEISHIRING"
	fmt.Println("", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
