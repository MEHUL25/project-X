package graphs

// An Adjacency Matrix is a 2D array where:
// matrix[i][j] == 1 means there is an edge from vertex i to vertex j.
// It's space-heavy: takes O(V^2) space.
// Great for dense graphs.

import "fmt"

type Graph struct {
	vertices int
	matrix   [][]int
}

// Constructor for graph
func NewGraph(v int) *Graph {
	matrix := make([][]int, v)
	for i := range matrix {
		matrix[i] = make([]int, v)
	}
	return &Graph{
		vertices: v,
		matrix:   matrix,
	}
}

// AddEdge adds an edge from u to v
func (g *Graph) AddEdge(u, v int, directed bool) {
	g.matrix[u][v] = 1
	if !directed {
		g.matrix[v][u] = 1
	}
}

// PrintGraph prints the adjacency matrix
func (g *Graph) PrintGraph() {
	fmt.Println("Adjacency Matrix:")
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

func createAdjacencyMatrixGraph() {
	g := NewGraph(5) // 5 vertices: 0 to 4

	g.AddEdge(0, 1, false)
	g.AddEdge(0, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(2, 3, false)
	g.AddEdge(3, 4, false)

	g.PrintGraph()
}

// Output

// Adjacency Matrix:
// 0 1 1 0 0
// 1 0 0 1 0
// 1 0 0 1 0
// 0 1 1 0 1
// 0 0 0 1 0
