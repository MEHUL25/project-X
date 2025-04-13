package trees

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// PreOrder Traversal (Root, Left, Right)
func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	PreOrder(root.left)
	PreOrder(root.right)
}
