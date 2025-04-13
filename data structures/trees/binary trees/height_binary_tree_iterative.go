package trees

// The height of a binary tree is the number of edges on the longest path from the root to a leaf.
// Alternatively, if you're counting in terms of nodes, height is the number of nodes on the longest path minus 1.

// Iterative Approach (BFS Level Order)
// This method uses queue and traverses level-by-level:

// When Tree is deep and skewed	Iterative preferred

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func heightIterative(root *TreeNode) int {
	if root == nil {
		return -1
	}

	queue := []*TreeNode{root}
	height := -1

	for len(queue) > 0 {
		levelSize := len(queue)
		height++

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return height
}
