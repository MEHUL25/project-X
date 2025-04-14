package graphs

// To find the largest connected component in an undirected graph, you can use DFS or BFS to explore each component and keep track of the size of the largest one you encounter.

import (
	"fmt"
)

// Graph using adjacency list
type Graph struct {
	adj map[int][]int
}

// NewGraph initializes a new empty graph
func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

// AddEdge adds an undirected edge
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// DFS returns the size of a connected component starting from 'v'
func (g *Graph) dfs(v int, visited map[int]bool) int {
	stack := []int{v}
	size := 0

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[curr] {
			continue
		}
		visited[curr] = true
		size++

		for _, neighbor := range g.adj[curr] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
	return size
}

// LargestComponent returns the size of the largest component
func (g *Graph) LargestComponent() int {
	visited := make(map[int]bool)
	largest := 0

	for node := range g.adj {
		if !visited[node] {
			size := g.dfs(node, visited)
			if size > largest {
				largest = size
			}
		}
	}
	return largest
}

func lagestComponentExample() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	// Node 5 is isolated

	fmt.Printf("Size of largest component: %d\n", g.LargestComponent())
}
