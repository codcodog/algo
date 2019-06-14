package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	convert(root)
	return root
}

func convert(cur *TreeNode) *TreeNode {
	if cur == nil {
		return nil
	}
	convert(cur.Right)
	cur.Val += sum
	sum = cur.Val
	convert(cur.Left)
	return cur
}
