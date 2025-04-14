package graphs

import (
	"fmt"
)

// An Adjacency List is a common way to represent a graph:
// Each node (vertex) stores a list of neighbors.
// Efficient for sparse graphs (i.e., fewer edges).

type Graph struct {
	adjacencyList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

// Add an edge (u -> v)
func (g *Graph) AddEdge(u, v int, directed bool) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	if !directed {
		g.adjacencyList[v] = append(g.adjacencyList[v], u)
	}
}

func (g *Graph) PrintGraph() {
	for node, neighbors := range g.adjacencyList {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}

func createAdjacencyListGraph() {
	graph := NewGraph()

	// Add some edges
	graph.AddEdge(0, 1, false)
	graph.AddEdge(0, 2, false)
	graph.AddEdge(1, 3, false)
	graph.AddEdge(2, 3, false)
	graph.AddEdge(3, 4, false)

	graph.PrintGraph()
}
