package path

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths int

// Note: 从根节点开始，查找匹配的路径
func pathSum(root *TreeNode, sum int) int {
	paths = 0

	if root == nil {
		return paths
	}
	dfs(root, sum)

	return paths
}

func dfs(node *TreeNode, target int) {
	if node == nil {
		return
	}

	test(node, target)
	dfs(node.Left, target)
	dfs(node.Right, target)
}

func test(node *TreeNode, target int) {
	if node == nil {
		return
	}
	if node.Val == target {
		paths++
	}
	test(node.Left, target-node.Val)
	test(node.Right, target-node.Val)
}
