package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var depth int

func diameterOfBinaryTree(root *TreeNode) int {
	depth = 0
	maxDepth(root)
	return depth
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	depth = max(depth, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
