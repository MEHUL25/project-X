package traversals

// Start from a source vertex.
// Push it into a queue and mark as visited.
// While the queue is not empty:
// Pop the front node.
// Visit all its unvisited neighbors, enqueue them, and mark as visited.

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int, directed bool) {
	g.adjList[u] = append(g.adjList[u], v)
	if !directed {
		g.adjList[v] = append(g.adjList[v], u)
	}
}

// 🔁 Breadth-First Search
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	fmt.Print("BFS Traversal: ")

	for len(queue) > 0 {
		// Dequeue
		node := queue[0]
		queue = queue[1:]

		// If not visited
		if !visited[node] {
			fmt.Printf("%d ", node)
			visited[node] = true

			// Enqueue neighbors
			for _, neighbor := range g.adjList[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	fmt.Println()
}

func BFSTraversal() {
	g := NewGraph()
	g.AddEdge(0, 1, false)
	g.AddEdge(0, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(2, 3, false)
	g.AddEdge(3, 4, false)

	g.BFS(0)
}
