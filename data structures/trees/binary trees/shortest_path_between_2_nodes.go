package trees

// Find the Lowest Common Ancestor (LCA) of val1 and val2.
// Trace path from LCA to val1 and from LCA to val2.
// Combine paths to get the full shortest path.

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func shortestPath(root *TreeNode, val1, val2 int) []int {
	lca := findLCA(root, val1, val2)
	if lca == nil {
		return []int{}
	}

	pathToVal1 := []int{}
	pathToVal2 := []int{}

	findPath(lca, val1, &pathToVal1)
	findPath(lca, val2, &pathToVal2)

	// Reverse pathToVal1 (except the LCA to avoid duplication)
	reverse(pathToVal1)
	pathToVal2 = pathToVal2[1:] // skip LCA duplicate
	return append(pathToVal1, pathToVal2...)
}

// Standard LCA function for binary tree
func findLCA(root *TreeNode, val1, val2 int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val1 || root.Val == val2 {
		return root
	}

	left := findLCA(root.Left, val1, val2)
	right := findLCA(root.Right, val1, val2)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// Helper to get path from root to target value
func findPath(root *TreeNode, target int, path *[]int) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root.Val)

	if root.Val == target {
		return true
	}

	if findPath(root.Left, target, path) || findPath(root.Right, target, path) {
		return true
	}

	// Backtrack
	*path = (*path)[:len(*path)-1]
	return false
}

// Reverses a slice in place
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
