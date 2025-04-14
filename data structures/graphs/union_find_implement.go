package graphs

// Absolutely! The Union-Find (or Disjoint Set Union – DSU) is a powerful data structure used to:

// Detect cycles in undirected graphs

// Manage disjoint sets

// Solve connectivity problems
import (
	"fmt"
)

// UnionFind struct
type UnionFind struct {
	parent []int
	rank   []int
}

// Initialize UnionFind with n elements (0 to n-1)
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{parent, rank}
}

// Find with path compression
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Union by rank
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in the same set
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	return true
}

func unionFindExample() {
	n := 5 // Number of nodes
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 4}, // This will introduce a cycle
	}

	uf := NewUnionFind(n)
	hasCycle := false

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !uf.Union(u, v) {
			fmt.Printf("Cycle detected while processing edge (%d, %d)\n", u, v)
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		fmt.Println("No cycles found in the graph.")
	}
}
