package all

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	nums := []uint8{}
	for x > 0 {
		nums = append(nums, uint8(x%10))
		x = x/10
	}

	for i, j := 0, len(nums)-1; i < j;  {
		if (nums[i] != nums[j]) {
			return false
		}
		i ++
		j--
	}

	return true

}
