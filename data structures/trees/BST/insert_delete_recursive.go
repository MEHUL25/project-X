package trees

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else if val > root.Val {
		root.Right = insertNode(root.Right, val)
	}

	return root
}

// all 3 cases:

// Node is a leaf
// Node has one child
// Node has two children – replace with inorder successor
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// Found the node to delete

		// Case 1: No child
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// Case 2: One child
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// Case 3: Two children - find inorder successor
		successor := findMin(root.Right)
		root.Val = successor.Val
		root.Right = deleteNode(root.Right, successor.Val)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func bstExample() {
	var root *TreeNode

	// Insert values into the tree
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, val := range values {
		root = insertNode(root, val)
	}

}
