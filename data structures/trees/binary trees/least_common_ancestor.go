package trees

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Standard LCA function for binary tree
func findLCA(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val1 || root.Val == val2 {
		return root
	}

	left := findLCA(root.Left, val1, val2)
	right := findLCA(root.Right, val1, val2)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
