package graph

import (
	"fmt"
	"sort"
)

type KruskalEdge struct {
	u, v, weight int
}

func find(parent []int, x int) int {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent, rank []int, x, y int) bool {
	rootX := find(parent, x)
	rootY := find(parent, y)
	if rootX == rootY {
		return false
	}
	if rank[rootX] < rank[rootY] {
		parent[rootX] = rootY
	} else if rank[rootX] > rank[rootY] {
		parent[rootY] = rootX
	} else {
		parent[rootY] = rootX
		rank[rootX]++
	}
	return true
}

func kruskal(n int, edges []KruskalEdge) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	totalWeight := 0
	for _, edge := range edges {
		if union(parent, rank, edge.u, edge.v) {
			totalWeight += edge.weight
		}
	}

	return totalWeight
}

func krushkalAlgorithmExample() {
	edges := []KruskalEdge{
		{0, 1, 1}, {0, 2, 2}, {1, 2, 4}, {1, 3, 6}, {2, 3, 3},
	}
	fmt.Println("Minimum cost of MST (Kruskal):", kruskal(4, edges))
}
