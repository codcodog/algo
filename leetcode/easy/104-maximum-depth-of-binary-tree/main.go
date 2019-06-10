package depthtree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	var depth int
	if right > left {
		depth = right
	} else {
		depth = left
	}

	return 1 + depth
}
