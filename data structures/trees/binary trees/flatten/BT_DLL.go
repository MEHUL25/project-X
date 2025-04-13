package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Convert a binary tree to a sorted doubly linked list in-place. The left and right pointers of the nodes should act as prev and next pointers.

// For a BST, inorder traversal gives nodes in sorted order. So we can connect nodes while doing inorder traversal.

type DLLNode = TreeNode // Reuse TreeNode

func bTreeToDLL(root *TreeNode) *TreeNode {
	var head, prev *TreeNode

	var convert func(node *TreeNode)
	convert = func(node *TreeNode) {
		if node == nil {
			return
		}

		// In-order traversal
		convert(node.Left)

		if prev == nil {
			head = node // first node becomes head
		} else {
			prev.Right = node
			node.Left = prev
		}
		prev = node

		convert(node.Right)
	}

	convert(root)
	return head
}
