package all

import "sort"

type NumList []int

func (l NumList) Len() int {
	return len(l)
}
func (l NumList) Less(i, j int) bool {
	return l[i] < l[j]
}
func (l NumList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Sort(NumList(nums))
	// fmt.Printf("%v", nums)

	if len(nums) < 3 {
		return 0
	}

	min := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k && j < len(nums) {
			tmp := nums[i] + nums[j] + nums[k] - target
			if AbsInt(tmp) < AbsInt(min-target) {
				min = nums[i] + nums[j] + nums[k]
			}
			if tmp == 0 {
				return nums[i] + nums[j] + nums[k]
			} else if tmp < 0 {
				j++
			} else {
				k--
			}
		}
	}

	// fmt.Println(data)
	return min

}