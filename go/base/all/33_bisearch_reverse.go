package all

import "fmt"

func binarySearch(nums []int, target, i, j int) int {
	for i <= j {
		index := (i + j) / 2
		if nums[index] == target {
			return index
		} else if nums[index] < target {
			i = index + 1
		} else {
			j = index - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		index := (i + j) / 2
		fmt.Println("index : ", index, i, j)
		if nums[i] <= nums[index] {
			res := binarySearch(nums, target, i, index)
			if res >= 0 {
				return res
			}
			i = index + 1
		} else {
			res := binarySearch(nums, target, index+1, j)
			if res >= 0 {
				return res
			}
			j = index
		}
		fmt.Println(index)
	}
	return -1
}