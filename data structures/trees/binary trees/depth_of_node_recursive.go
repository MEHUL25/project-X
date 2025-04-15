package trees

// Depth of a node = Number of edges from the root to that node.
// Maximum depth of the tree = Depth of the deepest node (a leaf usually).
// 📌 In most contexts, "maximum depth of the tree" = height of the tree (if root is at depth 0 and leaf is the deepest).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depthOfNode(root *TreeNode, target int) int {
	return findDepth(root, target, 0)
}

func findDepth(node *TreeNode, target, depth int) int {
	if node == nil {
		return -1 // Not found
	}

	if node.Val == target {
		return depth
	}

	// Search left subtree
	left := findDepth(node.Left, target, depth+1)
	if left != -1 {
		return left
	}

	// Search right subtree
	return findDepth(node.Right, target, depth+1)
}

// Time Complexity: O(n) → Must visit every node in the worst case.

// Space Complexity: O(h) → Due to recursion stack.
