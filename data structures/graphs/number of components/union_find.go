package graphs

import (
	"fmt"
)

func countComponents(n int, edges [][]int) int {
	// Step 1: Initialize parent array
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i // each node is its own parent initially
	}

	// Step 2: Find function with path compression
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // path compression
		}
		return parent[x]
	}

	// Step 3: Union function
	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX != rootY {
			parent[rootX] = rootY // merge sets
		}
	}

	// Step 4: Perform unions for all edges
	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	// Step 5: Count unique roots
	uniqueRoots := make(map[int]bool)
	for i := 0; i < n; i++ {
		root := find(i)
		uniqueRoots[root] = true
	}

	return len(uniqueRoots)
}

func unionFindExample() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}

	fmt.Println("Number of connected components:", countComponents(n, edges))
}
