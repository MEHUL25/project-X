package graphs

import (
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	vertices map[string][]string
	directed bool
}

// NewGraph creates a new graph instance
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertices: make(map[string][]string),
		directed: directed,
	}
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = []string{}
	}
}

// AddEdge adds an edge from vertex1 to vertex2
func (g *Graph) AddEdge(v1, v2 string) {
	g.AddVertex(v1)
	g.AddVertex(v2)

	g.vertices[v1] = append(g.vertices[v1], v2)

	if !g.directed {
		g.vertices[v2] = append(g.vertices[v2], v1)
	}
}

// PrintGraph prints all vertices and their edges
func (g *Graph) PrintGraph() {
	for vertex, edges := range g.vertices {
		fmt.Printf("%s -> %v\n", vertex, edges)
	}
}

func createGraph() {
	graph := NewGraph(false) // set true for directed graph

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")
	graph.AddEdge("E", "F")

	graph.PrintGraph()
}
