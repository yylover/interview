package all


/**
	1. 目前使用递归的方式来实验的，怎么用循环的方式来实现。
	2.
 */

func innerFunc2(str string, sum int, leftNum int, targetNum int, res *[]string) {
	if sum < 0 {
		return
	}

	if len(str) == targetNum {
		*res = append(*res, str)
		return
	}

	tmp := []string{"(", ")"}
	if leftNum == targetNum/2 {
		tmp = []string{")"}
	} else if sum == 0 {
		tmp = []string{"("}
	}
	for j := 0; j < len(tmp); j++ {
		if tmp[j] == "(" {
			innerFunc2(str+tmp[j], sum+1, leftNum+1, targetNum, res)
		} else {
			innerFunc2(str+tmp[j], sum-1, leftNum, targetNum, res)
		}
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	innerFunc2("", 0, 0, 2*n, &res)
	return res
}