package traversals

// It’s like normal level order traversal (BFS), but:

// Left to right for level 0

// Right to left for level 1

// Left to right for level 2

// We’ll use a queue for BFS and a boolean flag to toggle direction.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Zigzag Level Order Traversal
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Place value in correct position based on direction
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight // Flip direction
	}

	return result
}
