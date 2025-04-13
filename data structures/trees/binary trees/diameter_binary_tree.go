package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// Helper function to calculate height and update maxDiameter
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)

		// Update max diameter at this node
		if leftHeight+rightHeight > maxDiameter {
			maxDiameter = leftHeight + rightHeight
		}

		// Return height of current subtree
		return 1 + max(leftHeight, rightHeight)
	}

	height(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
