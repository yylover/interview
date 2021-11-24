package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
)




// 二分查找
func binarySearch(nums []int, target, i, j int) int {
	for i <= j  {
		index := (i+j)/2
		if nums[index] == target {
			return index
		} else if nums[index] < target {
			i = index+1
		} else {
			j = index-1
		}
	}
	return -1
}

func search(nums []int, target int) int {

	i, j := 0, len(nums)-1
	for i <= j {
		index := (i+j)/2
		fmt.Println("index : ", index, i, j)
		if nums[i] <= nums[index] {
			res := binarySearch(nums, target, i, index)
			if (res >= 0) {
				return res
			}
			i = index+1
		} else {
			res := binarySearch(nums, target, index+1, j)
			if (res >= 0) {
				return res
			}
			j = index
		}
		fmt.Println(index)
	}
	return -1;
}

func incrP(p *int)int{
	*p ++
	return *p
}
type Slice []int
func NewSlice() Slice {
	return make(Slice, 0)
}

func (s*Slice)Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

type RandomizedCollection struct {
	datas []int
	index map[int][]int
}

func NewRandomizedCollection() *RandomizedCollection{
	rc := &RandomizedCollection{
		datas: make([]int, 0),
		index: make(map[int][]int, 0),
	}
	return rc
}

func (rc *RandomizedCollection) insert(val int) bool {
	rc.datas = append(rc.datas, val)

	//索引更新
	retVal := false
	indexs, ok := rc.index[val]
	if !ok {
		indexs = make([]int, 0)
		rc.index[val] = indexs
		retVal = true
	}
	indexs = append(indexs, len(rc.datas)-1)
	rc.index[val] = indexs
	return retVal
}

func (rc *RandomizedCollection) remove(val int) bool {
	indexs, ok := rc.index[val]
	if !ok {
		return false
	}
	if len(indexs) <=0 {
		return false
	}

	loc := indexs[len(indexs)-1]
	indexs = indexs[0:len(indexs)-1]
	rc.index[val] = indexs

	//删除指定位置元素
	rc.datas = append(rc.datas[:loc], rc.datas[loc+1:]...)
	return true
}

func (rc *RandomizedCollection) getRandom() int {
	return rc.datas[rand.Intn(len(rc.datas))]
}



func main() {
	rc := NewRandomizedCollection()
	rc.insert(1)
	fmt.Println(rc.datas)
	rc.insert(2)
	fmt.Println(rc.datas)
	rc.insert(3)
	fmt.Println(rc.datas)
	rc.insert(1)
	fmt.Println(rc.datas)
	rc.remove(3)
	fmt.Println(rc.datas)
	rc.remove(2)
	fmt.Println(rc.datas)
	fmt.Println(rc.getRandom())



	return ;
	br := bufio.NewReader(strings.NewReader("1+2"))
	w, err := br.ReadSlice('+')
	fmt.Println("%q ", string(w), err)
	w, err = br.ReadSlice('+')
	fmt.Println("%q ", string(w), err)

}
