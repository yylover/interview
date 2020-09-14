package main

import "fmt"

func coutingSort(arr []int, maxValue int)[]int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}

	res := []int{}
	for i := 0; i < bucketLen; i++ {
		for j := 0; j < bucket[i]; j++ {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	arr := []int{1,2, 3, 4, 2, 3, 4, 2, 1, 3}
	arr = coutingSort(arr, 5)
	fmt.Println("coutingSort : ", arr)
}