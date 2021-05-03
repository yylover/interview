package main

/**
	递归解法，有没有迭代的解决方式？
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	leftSame := isSameTree(p.Left, q.Left)
	rightSame := isSameTree(p.Right, q.Right)
	return leftSame && rightSame
}
