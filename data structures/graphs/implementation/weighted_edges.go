package graphs

import (
	"fmt"
)

// Edge represents a weighted edge to a neighbor
type Edge struct {
	to     int
	weight int
}

// Graph represents a weighted undirected graph using an adjacency list
type Graph struct {
	adj map[int][]Edge
}

// NewGraph initializes a new weighted graph
func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]Edge),
	}
}

// AddEdge adds a weighted undirected edge to the graph
func (g *Graph) AddEdge(u, v, weight int) {
	g.adj[u] = append(g.adj[u], Edge{to: v, weight: weight})
	g.adj[v] = append(g.adj[v], Edge{to: u, weight: weight}) // Remove this line for a directed graph
}

// Print displays the graph's adjacency list
func (g *Graph) Print() {
	for u, edges := range g.adj {
		fmt.Printf("%d: ", u)
		for _, edge := range edges {
			fmt.Printf("-> %d (w=%d) ", edge.to, edge.weight)
		}
		fmt.Println()
	}
}

func example() {
	g := NewGraph()
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 8)

	fmt.Println("Weighted Graph:")
	g.Print()
}
