package trees

// TreeNode defines a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mirrorTree recursively mirrors the binary tree
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
