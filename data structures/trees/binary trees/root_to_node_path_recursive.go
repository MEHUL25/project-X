package trees

// We use recursion with backtracking:
// At each node, we append it to the path.
// If it’s not the target, we recursively call left and right child.
// If the target is not found down that path, we pop (backtrack) that node from the path.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rootToNodePath(root *TreeNode, target int) []int {
	path := []int{}
	if dfsPath(root, target, &path) {
		return path
	}
	return []int{} // target not found
}

func dfsPath(node *TreeNode, target int, path *[]int) bool {
	if node == nil {
		return false
	}

	*path = append(*path, node.Val)

	if node.Val == target {
		return true
	}

	// Check in left or right subtree
	if dfsPath(node.Left, target, path) || dfsPath(node.Right, target, path) {
		return true
	}

	// Backtrack if target not found in either subtree
	*path = (*path)[:len(*path)-1]
	return false
}
