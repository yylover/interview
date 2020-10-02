package all


/**
如果整数只有int32, 如何判断溢出，可以判断进位之前的一次数据
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

func reverse(x int) int {
	d := (x)
	if x < 0 {
		d = -(x)
	}

	nums := []int{}
	for d != 0 {
		nums = append(nums, d%10)
		d = d/10;
	}

	res := int(0)
	for i := 0; i < len(nums); i++ {
		if isFull(res, x < 0, nums[i]) {
			return 0
		}
		res = res* 10 + nums[i]
		if res < 0 {
			return 0
		}
	}

	if x < 0 {
		res = -res
	}

	return res

}