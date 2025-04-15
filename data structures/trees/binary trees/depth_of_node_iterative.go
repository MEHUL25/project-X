package trees

// Depth of a node = Number of edges from the root to that node.
// Maximum depth of the tree = Depth of the deepest node (a leaf usually).
// 📌 In most contexts, "maximum depth of the tree" = height of the tree (if root is at depth 0 and leaf is the deepest).

// finding the depth of a specific node in a binary tree using BFS (level-order traversal).

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func depthOfNodeIterative(root *TreeNode, target int) int {
	if root == nil {
		return -1
	}

	type nodeDepth struct {
		node  *TreeNode
		depth int
	}

	queue := []nodeDepth{{root, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.node.Val == target {
			return current.depth
		}

		if current.node.Left != nil {
			queue = append(queue, nodeDepth{current.node.Left, current.depth + 1})
		}

		if current.node.Right != nil {
			queue = append(queue, nodeDepth{current.node.Right, current.depth + 1})
		}
	}

	return -1 // Target not found
}

// Time Complexity: O(n) — each node visited once.

// Space Complexity: O(w) — where w is the maximum width of the tree (queue size).
