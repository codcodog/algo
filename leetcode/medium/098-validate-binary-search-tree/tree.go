package tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}

	val := node.Val
	if val <= lower {
		return false
	}
	if val >= upper {
		return false
	}
	return helper(node.Left, lower, val) && helper(node.Right, val, upper)
}
