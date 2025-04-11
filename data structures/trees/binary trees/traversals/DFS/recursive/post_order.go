package trees

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// PostOrder Traversal (Left, Right, Root)
func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Print(root.data, " ")
}
