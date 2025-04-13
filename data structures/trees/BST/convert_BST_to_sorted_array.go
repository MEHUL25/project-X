package bst

// Use in-order traversal:

// In a BST, in-order traversal (Left → Node → Right) visits nodes in ascending order.

// So we can recursively traverse the tree and append values to a slice.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Convert BST to Sorted Array using in-order traversal
func bstToSortedArray(root *TreeNode) []int {
	var result []int
	inOrder(root, &result)
	return result
}

func inOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, result)
	*result = append(*result, node.Val)
	inOrder(node.Right, result)
}
