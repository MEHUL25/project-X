package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertNodeIterative(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}

	current := root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = newNode
				break
			}
			current = current.Left
		} else if val > current.Val {
			if current.Right == nil {
				current.Right = newNode
				break
			}
			current = current.Right
		} else {
			// Duplicate value, ignore or handle as needed
			break
		}
	}

	return root
}
func deleteNodeIterative(root *TreeNode, key int) *TreeNode {
	parent := (*TreeNode)(nil)
	current := root

	// Step 1: Find the node to delete
	for current != nil && current.Val != key {
		parent = current
		if key < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	if current == nil {
		return root // key not found
	}

	// Step 2: Case with 2 children
	if current.Left != nil && current.Right != nil {
		// Find inorder successor (min node in right subtree)
		successorParent := current
		successor := current.Right
		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}
		current.Val = successor.Val // Copy successor value to current node
		// Now delete successor
		current = successor
		parent = successorParent
		key = successor.Val
	}

	// Step 3: Node has at most 1 child
	var child *TreeNode
	if current.Left != nil {
		child = current.Left
	} else {
		child = current.Right
	}

	if parent == nil {
		return child // root node deleted
	}

	if parent.Left == current {
		parent.Left = child
	} else {
		parent.Right = child
	}

	return root
}
