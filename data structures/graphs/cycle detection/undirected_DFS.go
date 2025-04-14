package graphs

import (
	"fmt"
)

// Graph structure using adjacency list
type Graph struct {
	adjacency map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacency: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	// Since it's undirected, add both edges
	g.adjacency[u] = append(g.adjacency[u], v)
	g.adjacency[v] = append(g.adjacency[v], u)
}

// Helper DFS function to detect cycle
func (g *Graph) hasCycleDFS(node int, visited map[int]bool, parent int) bool {
	visited[node] = true

	for _, neighbor := range g.adjacency[node] {
		if !visited[neighbor] {
			if g.hasCycleDFS(neighbor, visited, node) {
				return true
			}
		} else if neighbor != parent {
			// A back edge found (visited neighbor not parent)
			return true
		}
	}
	return false
}

// Main function to check cycle in any component
func (g *Graph) HasCycle() bool {
	visited := make(map[int]bool)
	for node := range g.adjacency {
		if !visited[node] {
			if g.hasCycleDFS(node, visited, -1) {
				return true
			}
		}
	}
	return false
}

// Example usage
func undirectedDFS() {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1) // This edge creates a cycle

	if graph.HasCycle() {
		fmt.Println("Cycle detected in the graph.")
	} else {
		fmt.Println("No cycle in the graph.")
	}
}
