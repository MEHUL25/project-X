package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodePath struct {
	Node *TreeNode
	Path []int
}

func rootToNodePathIterative(root *TreeNode, target int) []int {
	if root == nil {
		return []int{}
	}

	stack := []NodePath{{Node: root, Path: []int{root.Val}}}

	for len(stack) > 0 {
		// Pop from stack
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr := top.Node
		path := top.Path

		if curr.Val == target {
			return path
		}

		// Right first (so left is processed first due to stack LIFO)
		if curr.Right != nil {
			stack = append(stack, NodePath{
				Node: curr.Right,
				Path: append([]int{}, append(path, curr.Right.Val)...),
			})
		}

		if curr.Left != nil {
			stack = append(stack, NodePath{
				Node: curr.Left,
				Path: append([]int{}, append(path, curr.Left.Val)...),
			})
		}
	}

	return []int{} // Not found
}
