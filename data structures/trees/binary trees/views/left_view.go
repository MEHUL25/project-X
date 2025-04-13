package views

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leftViewIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

func leftViewRecursive(root *TreeNode) []int {
	result := []int{}
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(result) {
			result = append(result, node.Val)
		}
		helper(node.Left, level+1)
		helper(node.Right, level+1)
	}
	helper(root, 0)
	return result
}
