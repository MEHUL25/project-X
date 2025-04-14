package graphs

// We use DFS with a recursion stack:

// A visited map to mark nodes we've visited.

// A recStack (recursion stack) to track the current DFS path.

// If a node is visited and also in the recursion stack, we found a cycle.

import (
	"fmt"
)

// Directed Graph structure
type Graph struct {
	adjacency map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjacency: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	// Directed edge: u → v
	g.adjacency[u] = append(g.adjacency[u], v)
}

// DFS function to detect cycle
func (g *Graph) hasCycleDFS(node int, visited, recStack map[int]bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, neighbor := range g.adjacency[node] {
		if !visited[neighbor] {
			if g.hasCycleDFS(neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			// Found a back edge
			return true
		}
	}

	recStack[node] = false // backtrack
	return false
}

// Main function to check all components
func (g *Graph) HasCycle() bool {
	visited := make(map[int]bool)
	recStack := make(map[int]bool)

	for node := range g.adjacency {
		if !visited[node] {
			if g.hasCycleDFS(node, visited, recStack) {
				return true
			}
		}
	}
	return false
}

// Example usage
func directedDFS() {
	graph := NewGraph()
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 2) // Back edge creating a cycle

	if graph.HasCycle() {
		fmt.Println("Cycle detected in the directed graph.")
	} else {
		fmt.Println("No cycle in the directed graph.")
	}
}
