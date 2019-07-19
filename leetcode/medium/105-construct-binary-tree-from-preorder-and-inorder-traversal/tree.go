package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0, 0, len(inorder)-1, preorder, inorder)
}

func helper(preStart, inStart, inEnd int, preorder, inorder []int) *TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}
	root := &TreeNode{preorder[preStart], nil, nil}
	inIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			inIndex = i
			break
		}
	}
	root.Left = helper(preStart+1, inStart, inIndex-1, preorder, inorder)
	root.Right = helper(preStart+inIndex-inStart+1, inIndex+1, inEnd, preorder, inorder) // (inIndex - inStart) is the size of root's left subtree
	return root
}
