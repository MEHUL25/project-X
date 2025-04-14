package graphs

// Use a queue to traverse the graph level-by-level.

// Keep a parent map to track the parent of each node.

// If we visit a visited neighbor that is not the parent, a cycle exists.

import (
	"container/list"
	"fmt"
)

// Graph structure
type Graph struct {
	adjacency map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacency: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjacency[u] = append(g.adjacency[u], v)
	g.adjacency[v] = append(g.adjacency[v], u) // Undirected
}

// BFS function to detect cycle from a start node
func (g *Graph) hasCycleBFS(start int, visited map[int]bool) bool {
	queue := list.New()
	parent := make(map[int]int)

	// Start BFS from 'start'
	visited[start] = true
	queue.PushBack(start)
	parent[start] = -1

	for queue.Len() > 0 {
		element := queue.Front()
		current := element.Value.(int)
		queue.Remove(element)

		for _, neighbor := range g.adjacency[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
				parent[neighbor] = current
			} else if parent[current] != neighbor {
				// If visited neighbor is not the parent -> cycle
				return true
			}
		}
	}

	return false
}

// Check all components
func (g *Graph) HasCycleBFS() bool {
	visited := make(map[int]bool)
	for node := range g.adjacency {
		if !visited[node] {
			if g.hasCycleBFS(node, visited) {
				return true
			}
		}
	}
	return false
}

// Example usage
func undirectedBFS() {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1) // Creates a cycle

	if graph.HasCycleBFS() {
		fmt.Println("Cycle detected in the graph (BFS).")
	} else {
		fmt.Println("No cycle in the graph (BFS).")
	}
}
