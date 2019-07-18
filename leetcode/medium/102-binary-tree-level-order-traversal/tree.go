package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}

	var helper func(node *TreeNode, level int)
	helper = func(node *TreeNode, level int) {
		if len(levels) == level {
			levels = append(levels, make([]int, 0))
		}
		levels[level] = append(levels[level], node.Val)

		if node.Left != nil {
			helper(node.Left, level+1)
		}
		if node.Right != nil {
			helper(node.Right, level+1)
		}
	}
	helper(root, 0)
	return levels
}
