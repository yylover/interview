package main

import "fmt"

// buildMaxHeap 建堆
func buildMaxHeap(arr []int, arrLen int) {
	// for i := 0; i <
	for i := arrLen/2; i >= 0; i -- {
		heapify(arr, i, arrLen)
	}

}

// heapify堆调整
func heapify(arr []int, left int, right int) {
	if left >= right {
		return
	}

	index := left
	for index < right {
		leftIndex :=  index * 2 + 1
		rightIndex := index *2+ 2
		maxIndex := index
		if leftIndex < right && arr[leftIndex] > arr[maxIndex] {
			maxIndex = leftIndex
		}

		if rightIndex < right && arr[rightIndex] > arr[maxIndex] {
			maxIndex = rightIndex
		}
		if maxIndex == index {
			break
		} else {
			arr[maxIndex], arr[index] =  arr[index], arr[maxIndex]
			index = maxIndex
		}
	}
}



func heapSort(arr []int) []int {
	arrLen := len(arr)
	// 1. 建堆
	buildMaxHeap(arr, arrLen)
	fmt.Println("arr : ", arr)
	for i := arrLen - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		//依次输出，调整
		heapify(arr, 0, i)
	}
	return arr
}


func main() {
	arr := []int{3,2,1,4,8,7,5,6}
	heapSort(arr)
	fmt.Println("sort : ", arr)
}