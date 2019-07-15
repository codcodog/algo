package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	helper(root)
	return result
}

func helper(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			helper(root.Left)
		}
		result = append(result, root.Val)
		if root.Right != nil {
			helper(root.Right)
		}
	}
}
