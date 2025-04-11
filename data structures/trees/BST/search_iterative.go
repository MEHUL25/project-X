package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h) where h is the height of the tree (O(log n) for balanced, O(n) for skewed).
// Space: Iterative: O(1)

func searchBSTIterative(root *TreeNode, target int) *TreeNode {
	current := root

	for current != nil {
		if current.Val == target {
			return current
		} else if target < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil // Not found
}
