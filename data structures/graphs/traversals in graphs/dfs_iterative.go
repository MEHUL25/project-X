package traversals

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func (g *Graph) DFSIterative(start int) {
	visited := make(map[int]bool)
	stack := []int{start}

	fmt.Print("DFS (Iterative) Traversal: ")
	for len(stack) > 0 {
		// Pop from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			fmt.Printf("%d ", node)
			visited[node] = true

			// Push unvisited neighbors (in reverse for same order as recursion)
			for i := len(g.adjList[node]) - 1; i >= 0; i-- {
				neighbor := g.adjList[node][i]
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}
	fmt.Println()
}
