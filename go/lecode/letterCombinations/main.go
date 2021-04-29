package main

import (
	"fmt"
	"strings"
)

/**
map 声明
string index to string
递归是否可以改成循环
 */

var nummap  = map[string][]string  {
	"2" : {"a", "b", "c"},
	"3" : {"d", "e", "f"},
	"4" : {"g", "h", "i"},
	"5" : {"j", "k", "l"},
	"6" : {"m", "n", "o"},
	"7" : {"p", "q", "r", "s"},
	"8" : {"t", "u", "v"},
	"9" : {"w", "x", "y", "z"},
}

func loopFunc(index int, digits string, str []string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, strings.Join(str, ""))
		return
	}

	nums, ok :=  nummap[string(digits[index])]
	if !ok {
		panic(fmt.Sprintf("invalid input : %v", string(digits[index])))
	}

	for _, c := range nums {
		str[index] = c
		loopFunc(index+1, digits, str, res)
	}

}

func letterCombinations(digits string) []string {
	res := []string{}
	str := make([]string, len(digits))
	if len(digits) == 0 {
		return res
	}
	loopFunc(0, digits, str, &res)
	return res
}

func main() {
	nums := letterCombinations("")
	fmt.Println("%v", nums)
}
