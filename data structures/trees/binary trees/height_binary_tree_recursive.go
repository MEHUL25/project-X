package trees

// The height of a binary tree is the number of edges on the longest path from the root to a leaf.
// Alternatively, if you're counting in terms of nodes, height is the number of nodes on the longest path minus 1.

// If the node is nil, return -1 (height of an empty tree).
// If you prefer to count in terms of nodes: return 0 here.

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(root *TreeNode) int {
	if root == nil {
		return -1 // height in terms of edges. If counting nodes, return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
