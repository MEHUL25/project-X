package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h) where h is the height of the tree (O(log n) for balanced, O(n) for skewed).
// Space: Recursive: O(h) stack space

func searchBST(root *TreeNode, target int) *TreeNode {
	if root == nil || root.Val == target {
		return root
	}

	if target < root.Val {
		return searchBST(root.Left, target)
	}
	return searchBST(root.Right, target)
}
