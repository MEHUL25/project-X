package bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBSTHelper checks if the tree is a valid BST within min and max range
func isValidBSTHelper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// Node must be within the range (exclusive)
	if node.Val <= min || node.Val >= max {
		return false
	}

	// Check left and right subtrees recursively with updated ranges
	return isValidBSTHelper(node.Left, min, node.Val) &&
		isValidBSTHelper(node.Right, node.Val, max)
}

// isValidBST initiates the check with the full integer range
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}
