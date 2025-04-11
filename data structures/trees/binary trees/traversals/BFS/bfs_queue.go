package traversals

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

// BFS Traversal (Level Order Traversal)
func BFS(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.data, " ")
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
