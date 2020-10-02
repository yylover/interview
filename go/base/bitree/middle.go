package main

import "fmt"

func findMedian(nums[] int, start int, end int) (int, int, float64) {
	arrLen := end - start
	isOdd := arrLen %2 == 1
	if isOdd {
		return  arrLen/2+start -1, arrLen/2+start + 1, float64(nums[arrLen/2+start])
	} else {
		return  arrLen/2+start-1, arrLen/2+start, float64(nums[arrLen/2+start-1] + nums[arrLen/2+start]) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	leftLen := len(nums1)
	rightLen := len(nums2)
	isOdd := (leftLen+rightLen)%2 == 1
	leftStart, leftEnd := 0, leftLen
	rightStart, rightEnd := 0, rightLen
	if isOdd { // 寻找中间的
		for leftEnd - leftStart + rightEnd - rightStart != 1 {
			leftStartIndex, leftEndIndex, leftMiddle := findMedian(nums1, leftStart, leftEnd)
			rightStartIndex, rightEndIndex, rightMiddle := findMedian(nums2, rightStart, rightEnd)
			if leftMiddle < rightMiddle {
				leftStart = leftStartIndex
				rightEnd = rightEndIndex
			} else {
				rightStart = rightStartIndex
				leftEnd = leftEndIndex
			}
		}


	} else { // 寻找中间的两个数
		for leftEnd - leftStart + rightEnd - rightStart != 1 {
			leftStartIndex, leftEndIndex, leftMiddle := findMedian(nums1, leftStart, leftEnd)
			rightStartIndex, rightEndIndex, rightMiddle := findMedian(nums2, rightStart, rightEnd)
			if leftMiddle < rightMiddle {
				leftStart = leftStartIndex
				rightEnd = rightEndIndex
			} else {
				rightStart = rightStartIndex
				leftEnd = leftEndIndex
			}
		}
	}

	return
}

func main() {
	nums1 := []int{1, 2, 4, 5}
	num2 := []int{3, 7, 8}
	middle := findMedianSortedArrays(nums1, num2)
	fmt.Println("midlle : ", middle)

}
