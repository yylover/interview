package all


/**
移动索引的时候判断终止条件
"" "+" "-" " --234" " -+234" " -+91283472332"
*/


const INT_MAX = 1<<31 - 1
const INT_MIN = - 1 << 31

func isFull(base int , isNegative bool,  mod int) bool {
	if !isNegative {
		if base > INT_MAX / 10 || (base == INT_MAX / 10 && mod > INT_MAX % 10) {
			return true
		} else {
			return false
		}
	} else {
		if base > (INT_MAX+1) / 10 || (base == (INT_MAX+1) / 10 && mod > (INT_MAX+1) % 10) {
			return true
		} else {
			return false
		}
	}
}


func isValidNum(charnum byte)bool {
	if charnum >= '0' && charnum <= '9' {
		return true
	}
	return false
}

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	isNegative := false

	i := 0
	if str[i] == ' ' {
		for i < len(str) && str[i] == ' ' {
			i++
		}
	}

	if i >= len(str) {
		return 0
	}

	if  str[i] == '-' {
		isNegative = true
		i++
	}

	if i >= len(str) {
		return 0
	}
	if str[i] == '+' {
		if isNegative { //已经处理过一个符号
			return 0
		}
		i++
	}

	res := 0
	for ; i < len(str); i++ {
		if (isValidNum(str[i])) {
			if isFull(res, isNegative, int(str[i]-'0')) {
				if isNegative {
					return INT_MIN
				} else {
					return INT_MAX
				}
			}

			res = res * 10 + int(str[i]-'0')
		} else {
			break
		}
	}

	if isNegative {
		res = -res
	}

	return res
}
