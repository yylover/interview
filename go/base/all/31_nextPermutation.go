package all

import "fmt"

func reverse(nums[]int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int) {
	found := false
	for j := len(nums) - 1; j > 0 ; j-- {
		if nums[j-1] >= nums[j] {
			continue
		}
		fmt.Println(j)
		//从右侧找到第一个大于当前值的数，进行交换
		k := len(nums)-1
		for ;nums[j-1] >= nums[k]&& k>0 ; k-- {
		}
		fmt.Println(k)
		nums[k], nums[j-1] = nums[j-1], nums[k]
		fmt.Println(nums)

		reverse(nums, j, len(nums)-1)
		found = true
		break
	}

	if !found {
		reverse(nums, 0, len(nums)-1)
	}
}