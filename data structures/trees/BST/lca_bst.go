package trees

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLCAinBST(root *TreeNode, val1, val2 int) *TreeNode {
	for root != nil {
		if val1 < root.Val && val2 < root.Val {
			root = root.Left
		} else if val1 > root.Val && val2 > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
