package main


func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructTree(nums, 0, len(nums))
}

func constructTree(nums []int , left int, right int) *TreeNode {
	if left >= right {
		return nil
	}

	index := findBiggestIndex(nums, left, right)
	node := &TreeNode{Val:nums[index], Left:nil, Right:nil}
	node.Left = constructTree(nums, left, index)
	node.Right = constructTree(nums, index+1, right)
	return node;
}

func findBiggestIndex(nums []int, left, right int) int {
	index := left
	for i := index + 1; i < right; i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}

func main() {
	nums := []int{3,2,1,6,0,5}
	node := constructMaximumBinaryTree(nums)
	PrePrintTree(node)
}