package traversals

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Boundary = left boundary (excluding leaves) + all leaves + right boundary (excluding leaves, reversed)

func printBoundary(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")

	// Print left boundary (excluding leaves)
	printLeftBoundary(root.Left)

	// Print all leaves
	printLeaves(root.Left)
	printLeaves(root.Right)

	// Print right boundary (excluding leaves, in reverse)
	printRightBoundary(root.Right)
}

func printLeftBoundary(node *TreeNode) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	fmt.Print(node.Val, " ")
	if node.Left != nil {
		printLeftBoundary(node.Left)
	} else {
		printLeftBoundary(node.Right)
	}
}

func printRightBoundary(node *TreeNode) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	if node.Right != nil {
		printRightBoundary(node.Right)
	} else {
		printRightBoundary(node.Left)
	}
	fmt.Print(node.Val, " ") // Print after recursion (bottom-up)
}

func printLeaves(node *TreeNode) {
	if node == nil {
		return
	}
	printLeaves(node.Left)
	if node.Left == nil && node.Right == nil {
		fmt.Print(node.Val, " ")
	}
	printLeaves(node.Right)
}
