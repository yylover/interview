package main

import (
	"fmt"
)

// bubbleSort
func bubbleSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 {
		return datas
	}

	// for j := llen - 1; j > 0; j-- {
	// 	for i := 0; i < j; i++ {
	// 		if datas[i] > datas[j] {
	// 			datas[i], datas[j] = datas[j], datas[i]
	// 		}
	// 	}
	// }

	for i := 0; i < llen-1; i++ {
		for j := 0; j < llen - 1 -i; j++ {
			if datas[j] > datas[j+1] {
				datas[j], datas[j+1] = datas[j+1], datas[j]
			}
		}
	}
	return datas
}

// 选择， 表现稳定，无论数据如何都是O(n2)复杂度
func selectionSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 {
		return datas
	}

	for i := 0; i < llen-1; i ++ {
		maxIndex := i
		for j := i+1; j < llen; j++ {
			if datas[j] < datas[maxIndex] {
				maxIndex = j
			}
		}
		datas[i], datas[maxIndex] = datas[maxIndex], datas[i]
	}

	return datas
}

// 插入排序
func insertSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 {
		return datas
	}

	for i := 0; i < llen-1; i ++ {
		for j := i+1; j > 0 && datas[j] < datas[j-1]; j-- {
			datas[j],datas[j-1] = datas[j-1], datas[j]
		}
	}
	return datas
}

// 希尔，第一个突破O(n2)的排序算法，是插入排序的改进版本，优先比较较远的元素，又称为缩小增量排序
func shellSort(datas []int) []int{
	llen := len(datas)
	if llen == 0 {
		return datas
	}


	gap := llen /  2
	for gap > 0 {
		for l := gap; l < llen; l++ {
			for k := l; k -gap > 0 && datas[k] < datas[k-gap]; k-=gap {
				datas[k], datas[k-gap] = datas[k-gap], datas[k]
			}
		}
		// for i := 0; i < gap; i++ {
		// 	for j := i+1; j < llen; j+= gap {
		// 	}
		// }

		fmt.Println("pivot", gap)
		gap = gap /2
	}
	// for i := 0; i < llen -1 ; i+=

	return datas
}

//

func main() {
	res := []int{1, 2, 19, 3, 4, 5, 6}
	fmt.Printf("%v \n", bubbleSort(res))
	res = []int{}
	fmt.Printf("%v \n", bubbleSort(res))
	res = []int{2}
	fmt.Printf("%v \n", bubbleSort(res))
	res = []int{3, 2}
	fmt.Printf("%v \n", bubbleSort(res))

	res = []int{1, 2, 19, 3, 4, 5, 6}
	fmt.Printf("%v \n", selectionSort(res))
	res = []int{}
	fmt.Printf("%v \n", selectionSort(res))
	res = []int{2}
	fmt.Printf("%v \n", selectionSort(res))
	res = []int{3, 2}
	fmt.Printf("%v \n", selectionSort(res))

	res = []int{1, 2, 19, 3, 6, 5, 4}
	fmt.Printf("%v \n", insertSort(res))
	res = []int{}
	fmt.Printf("%v \n", insertSort(res))
	res = []int{2}
	fmt.Printf("%v \n", insertSort(res))
	res = []int{3, 2}
	fmt.Printf("%v \n", insertSort(res))


	res = []int{1, 2, 19, 3, 6, 5, 4}
	fmt.Printf("%v \n", shellSort(res))
	// res = []int{}
	// fmt.Printf("%v \n", insertSort(res))
	// res = []int{2}
	// fmt.Printf("%v \n", insertSort(res))
	// res = []int{3, 2}
	// fmt.Printf("%v \n", insertSort(res))

	v := 42
	fmt.Println(int(float32(v)*1.2))
}
