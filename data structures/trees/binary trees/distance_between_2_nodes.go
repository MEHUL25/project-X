package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// First, find LCA (Lowest Common Ancestor), then compute distance from LCA to both nodes and add.

func findLCA(root *TreeNode, a, b int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > a && root.Val > b {
		return findLCA(root.Left, a, b)
	} else if root.Val < a && root.Val < b {
		return findLCA(root.Right, a, b)
	}
	return root
}

func findLevel(root *TreeNode, val, level int) int {
	if root == nil {
		return -1
	}
	if root.Val == val {
		return level
	}
	if val < root.Val {
		return findLevel(root.Left, val, level+1)
	}
	return findLevel(root.Right, val, level+1)
}

func findDistanceBetweenNodes(root *TreeNode, a, b int) int {
	lca := findLCA(root, a, b)
	d1 := findLevel(lca, a, 0)
	d2 := findLevel(lca, b, 0)
	return d1 + d2
}
