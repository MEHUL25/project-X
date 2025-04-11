package trees

// A binary tree is height-balanced if for every node, the height difference between the left and right subtrees is no more than 1.

// Recursively compute the height of left and right subtrees.
// If any subtree is unbalanced, propagate the failure up early (using a sentinel like -1).
// This avoids recomputing heights (unlike a naive top-down approach).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return checkHeight(root) != -1
}

func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
