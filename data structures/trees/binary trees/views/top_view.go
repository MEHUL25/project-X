package view

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iterative
// Perform level-order traversal.
// Keep track of horizontal distance (hd) from the root:
// Left child → hd - 1
// Right child → hd + 1
// For top view, store the first node encountered at each horizontal distance.

type Pair struct {
	Node *TreeNode
	HD   int
}

func topViewIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	hdMap := map[int]int{}
	queue := list.New()
	queue.PushBack(Pair{root, 0})

	minHD, maxHD := 0, 0

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(Pair)
		node, hd := curr.Node, curr.HD

		if _, found := hdMap[hd]; !found {
			hdMap[hd] = node.Val
		}
		if node.Left != nil {
			queue.PushBack(Pair{node.Left, hd - 1})
		}
		if node.Right != nil {
			queue.PushBack(Pair{node.Right, hd + 1})
		}
		if hd < minHD {
			minHD = hd
		}
		if hd > maxHD {
			maxHD = hd
		}
	}

	result := []int{}
	for i := minHD; i <= maxHD; i++ {
		result = append(result, hdMap[i])
	}
	return result
}

// Top View – Recursive is not straightforward and typically avoided because:
// Level-order traversal (iterative) is the most efficient.

// You’d need complex tracking of levels + HD with recursive stack, which can be messy and inefficient.
