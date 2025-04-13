package bst

// Given a sorted array in ascending order, convert it to a height-balanced Binary Search Tree (BST).

// A height-balanced BST means the depth of the two subtrees of every node never differs by more than 1.

// TreeNode structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Convert sorted array to balanced BST
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

// Recursive helper function
func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}
