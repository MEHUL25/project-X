package traversals

// Note: DFS traversal order may vary depending on the order of neighbors in the adjacency list.

import "fmt"

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

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Print("DFS Traversal: ")
	g.dfsHelper(start, visited)
	fmt.Println()
}

func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	fmt.Printf("%d ", node)
	visited[node] = true
	for _, neighbor := range g.adjList[node] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited)
		}
	}
}

func DFSTraversal() {
	g := NewGraph()
	g.AddEdge(0, 1, false)
	g.AddEdge(0, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(2, 3, false)
	g.AddEdge(3, 4, false)

	g.DFS(0)
}
