package trees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printAllPaths(root *TreeNode) {
	var path []int
	printPathsHelper(root, path)
}

func printPathsHelper(node *TreeNode, path []int) {
	if node == nil {
		return
	}

	path = append(path, node.Val)

	if node.Left == nil && node.Right == nil {
		fmt.Println(path)
		return
	}

	printPathsHelper(node.Left, path)
	printPathsHelper(node.Right, path)
}
