package main

import "fmt"

/**
	思路是遍历，但是不要漏掉边界条件，最后一次是没有没有遍历到的aaa
 */
func largeGroupPositions(s string) [][]int {
	res := [][]int{}

	lastchar := int32(0)
	lastIndex := -1
	for index, char := range s {
		fmt.Printf("%V \n", char)
		if lastchar == 0 {
			lastchar = char
			lastIndex = index
		} else if char != lastchar {
			if index - lastIndex >= 3 {
				res = append(res, []int{lastIndex, index-1})
			}
			lastIndex = index
			lastchar = char
		}

		//边界条件
		if index == len(s)-1 && char == lastchar && index - lastIndex+1 >= 3 {
			res = append(res, []int{lastIndex, index})
		}
	}

	return res
}

func main() {
	fmt.Println("%v", largeGroupPositions("aaa"))

}