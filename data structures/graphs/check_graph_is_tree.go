package graphs

// To check if an undirected graph is a tree in Go, you need to validate two conditions:

// The graph is connected – every node can be reached from any other node.

// The graph has no cycles – it's acyclic.

import (
	"fmt"
)

// Graph is represented as an adjacency list
type Graph struct {
	adj map[int][]int
}

// NewGraph initializes a new graph
func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// isCyclicDFS checks for cycles using DFS
func (g *Graph) isCyclicDFS(v int, visited map[int]bool, parent int) bool {
	visited[v] = true

	for _, neighbor := range g.adj[v] {
		if !visited[neighbor] {
			if g.isCyclicDFS(neighbor, visited, v) {
				return true
			}
		} else if neighbor != parent {
			// A visited neighbor that is not the parent indicates a cycle
			return true
		}
	}
	return false
}

// IsTree checks if the graph is a tree
func (g *Graph) IsTree() bool {
	visited := make(map[int]bool)

	// Pick any starting node (first key in adj map)
	var startNode int
	for k := range g.adj {
		startNode = k
		break
	}

	// Check for cycle
	if g.isCyclicDFS(startNode, visited, -1) {
		return false
	}

	// Check if the graph is connected
	if len(visited) != len(g.adj) {
		return false
	}

	return true
}

func example() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)

	if g.IsTree() {
		fmt.Println("The graph is a tree")
	} else {
		fmt.Println("The graph is NOT a tree")
	}
}
