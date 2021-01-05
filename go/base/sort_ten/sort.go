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
		for j := 0; j < llen-1-i; j++ {
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

	for i := 0; i < llen-1; i++ {
		maxIndex := i
		for j := i + 1; j < llen; j++ {
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

	for i := 0; i < llen-1; i++ {
		for j := i + 1; j > 0 && datas[j] < datas[j-1]; j-- {
			datas[j], datas[j-1] = datas[j-1], datas[j]
		}
	}
	return datas
}

// 希尔，第一个突破O(n2)的排序算法，是插入排序的改进版本，优先比较较远的元素，又称为缩小增量排序
func shellSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 {
		return datas
	}

	gap := llen / 2
	for gap > 0 {
		for l := gap; l < llen; l++ {
			for k := l; k-gap > 0 && datas[k] < datas[k-gap]; k -= gap {
				datas[k], datas[k-gap] = datas[k-gap], datas[k]
			}
		}
		// for i := 0; i < gap; i++ {
		// 	for j := i+1; j < llen; j+= gap {
		// 	}
		// }

		fmt.Println("pivot", gap)
		gap = gap / 2
	}
	// for i := 0; i < llen -1 ; i+=

	return datas
}

// 归并排序
func mergeSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 || llen == 1 {
		return datas
	}

	length := len(datas) / 2
	left := mergeSort(datas[0:length])
	right := mergeSort(datas[length:llen])
	return mergeSortSub(left, right)
}

//
func mergeSortSub(datas1 []int, datas2 []int) []int {
	res := []int{}
	i,j := 0, 0

	//先保证没问题的排列好，最后把剩下的元素加到最后的结果中
	for i < len(datas1) && j < len(datas2) {
		if datas1[i] < datas2[j] {
			res = append(res, datas1[i])
			i++
		} else {
			res = append(res, datas2[j])
			j++
		}
	}

	if i < len(datas1) {
		res = append(res, datas1[i:len(datas1)] ...)
	}

	if j < len(datas2) {
		res = append(res, datas1[j:len(datas1)] ...)
	}


	return res
}


// 可能出现的边界条件，为空，i 到达上届 这种写法容易出错，
func mergeSortSub1(datas1 []int, datas2 []int) []int {
	res := []int{}
	for i, j := 0, 0; i < len(datas1) || j < len(datas2); {
		if len(datas1) == 0 || i >= len(datas1) {
			res = append(res, datas2[j])
			j++
		} else if len(datas2) == 0 || j >= len(datas2) {
			res = append(res, datas1[i])
			i++
		} else if datas1[i] < datas2[j] {
			res = append(res, datas1[i])
			i++
		} else {
			res = append(res, datas2[j])
			j++
		}
	}
	return res
}


//快速排序 inplace
func quickSort(datas []int) []int {
	llen := len(datas)
	if llen == 0 || llen == 1 {
		return datas
	}

	quickSortSub(datas, 0, len(datas)-1)

	return datas
}

func quickSortSub(datas []int, left, right int) {
	if left >= right {
		return
	}

	tempindex := left
	i, j := left, right
	for i < j {
		for datas[j] > datas[tempindex] && i < j {
			j --
		}

		for datas[i] < datas[tempindex] && i < j {
			i ++
		}

		datas[j], datas[i] = datas[i], datas[j]
	}
	datas[tempindex], datas[i] =  datas[i], datas[tempindex]

	quickSortSub(datas, left, i)
	quickSortSub(datas, i+1, right)
}

// 堆 近似完全二叉树，同时满足子节点的减值或索引总是大于（小于）他的父节点
func heapSort(datas []int) []int {
	if len(datas) <= 1 {
		return datas
	}

	res := []int{}



	return res
}


// 基数排序: 非比较型整数排序算法


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


	//
	res = []int{1, 2, 19, 3, 6, 5, 4}
	fmt.Printf("%v \n", mergeSort(res))
	res = []int{}
	fmt.Printf("%v \n", mergeSort(res))
	res = []int{2}
	fmt.Printf("%v \n", mergeSort(res))
	res = []int{3, 2}
	fmt.Printf("%v \n", mergeSort(res))


	res = []int{1, 2, 19, 3, 6, 5, 4}
	fmt.Printf("%v \n", shellSort(res))
	// res = []int{}
	// fmt.Printf("%v \n", insertSort(res))
	// res = []int{2}
	// fmt.Printf("%v \n", insertSort(res))
	// res = []int{3, 2}
	// fmt.Printf("%v \n", insertSort(res))
	//da := []int{1, 2, 3, 4}
	//fmt.Println("da : ", da[0:1])
	//fmt.Println("da : ", da[1:2])

	fmt.Println("quickSort")
	res = []int{1, 2, 19, 3, 6, 5, 4}
	fmt.Printf("%v \n", quickSort(res))
	res = []int{}
	fmt.Printf("%v \n", quickSort(res))
	res = []int{2}
	fmt.Printf("%v \n", quickSort(res))
	res = []int{3, 2}
	fmt.Printf("%v \n", quickSort(res))

}
