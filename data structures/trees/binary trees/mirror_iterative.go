package trees

// TreeNode defines a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// mirrorTreeIterative mirrors the tree using a queue
func mirrorTreeIterative(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Swap left and right
		current.Left, current.Right = current.Right, current.Left

		// Add children to the queue
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return root
}
