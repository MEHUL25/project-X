package graphs

// For early termination and memory efficiency, DFS is usually faster in practice for this specific use-case (especially for sparse graphs).

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

// ✅ Detect if a path exists from src to dst
func (g *Graph) PathExists(src, dst int) bool {
	if src == dst {
		return true
	}

	visited := make(map[int]bool)
	return g.dfsPath(src, dst, visited)
}

func (g *Graph) dfsPath(current, target int, visited map[int]bool) bool {
	if current == target {
		return true
	}
	visited[current] = true

	for _, neighbor := range g.adjList[current] {
		if !visited[neighbor] {
			if g.dfsPath(neighbor, target, visited) {
				return true
			}
		}
	}

	return false
}

func pathExistsExample() {
	g := NewGraph()
	g.AddEdge(0, 1, false)
	g.AddEdge(0, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(3, 4, false)

	fmt.Println("Path from 0 to 4:", g.PathExists(0, 4)) // true
	fmt.Println("Path from 2 to 4:", g.PathExists(2, 4)) // true
	fmt.Println("Path from 4 to 0:", g.PathExists(4, 0)) // true (because undirected)
	fmt.Println("Path from 2 to 5:", g.PathExists(2, 5)) // false
}

// Time: O(V + E), where V = vertices and E = edges

// Space: O(V) for visited map and recursion stack
