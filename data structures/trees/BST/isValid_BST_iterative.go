package bst

import "math"

// In a valid BST, the in-order traversal yields values in strictly increasing order.

// So, we can:

// Traverse the tree iteratively (using a stack).

// Keep track of the previous visited value.

// If current value ≤ previous value, it's not a BST.

// TreeNode defines a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBSTIterative checks if a binary tree is a valid BST using iterative in-order traversal
func isValidBSTIterative(root *TreeNode) bool {
	stack := []*TreeNode{}
	curr := root
	prevVal := math.MinInt64

	for curr != nil || len(stack) > 0 {
		// Go left as far as possible
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// Process node
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Val <= prevVal {
			return false
		}
		prevVal = curr.Val

		// Move to right subtree
		curr = curr.Right
	}

	return true
}
