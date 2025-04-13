package traversals

// DFS with Level Tracking:
// The idea is to traverse the tree recursively.

// While traversing, we keep track of the level we're currently on.

// We toggle the direction (left-to-right or right-to-left) based on the level's parity (even or odd).

// Use a recursive approach where we visit the left child first, then the right child, to mimic DFS traversal.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS Helper Function
func dfs(root *TreeNode, level int, result *[][]int, leftToRight bool) {
	if root == nil {
		return
	}

	// Ensure the result has enough levels
	if len(*result) <= level {
		*result = append(*result, []int{})
	}

	// Append the current node's value based on the current direction
	if leftToRight {
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		// For right-to-left, insert at the beginning (reverse order)
		(*result)[level] = append([]int{root.Val}, (*result)[level]...)
	}

	// Recursively visit left and right children
	dfs(root.Left, level+1, result, !leftToRight)
	dfs(root.Right, level+1, result, !leftToRight)
}
