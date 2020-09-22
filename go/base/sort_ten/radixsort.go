package main

import (
	"fmt"
	"math"
)


/**
 最大基数
 */
func radixSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}

	bitNum := 0
	for max > 0 {
		max /= 10
		bitNum++
	}

	baseSort := [10][]int{}
	for i := 0; i < 0; i++ {
		baseSort[i] = []int{}
	}

	for i := 0; i < bitNum; i++ {
		for j := 0; j < len(arr); j++ {
			basenum := (arr[j] / int(math.Pow10(i))) % 10
			// basenum := arr[j] / int(math.Pow10(i))
			baseSort[basenum] = append(baseSort[basenum], arr[j])
			// fmt.Println(basenum, bitNum, arr[j])
		}

		arr = []int{}
		for j := 0; j < 10; j++ {
			for k := 0; k < len(baseSort[j]); k++ {
				arr = append(arr, baseSort[j][k])
			}
			baseSort[j] = []int{}
		}
	}

	// fmt.Println("bitNum, ", arr)

	return arr
}

func main() {
	arr := []int{112, 23, 44, 132, 113}
	radixSort(arr)
	fmt.Println(arr)

	fmt.Println(233 >> 2)
	fmt.Println(int(math.Pow10(2)))
	fmt.Println(int(math.Pow10(0)))

	fmt.Println((233 >> 2)%10)
}