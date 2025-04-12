package flatten

// Given a binary tree, flatten it into a "linked list" in-place, where:

// The "linked list" should use the right pointers.

// It should follow the preorder traversal: Root → Left → Right.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// Save original left and right
	left := root.Left
	right := root.Right

	// Flatten left subtree becomes the right
	root.Left = nil
	root.Right = left

	// Attach the original right subtree at the end of the new right chain
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = right
}
