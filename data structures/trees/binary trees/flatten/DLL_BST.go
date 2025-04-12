package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Use the inorder strategy.

// Recursively build left subtree → assign root → build right subtree.

var dllHead *TreeNode

func countDLLNodes(head *TreeNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Right
	}
	return count
}

func sortedDLLToBST(n int) *TreeNode {
	if n <= 0 {
		return nil
	}

	// Recursively build left subtree
	left := sortedDLLToBST(n / 2)

	// Root node
	root := dllHead
	root.Left = left

	// Move dllHead forward
	dllHead = dllHead.Right

	// Recursively build right subtree
	root.Right = sortedDLLToBST(n - n/2 - 1)
	return root
}
