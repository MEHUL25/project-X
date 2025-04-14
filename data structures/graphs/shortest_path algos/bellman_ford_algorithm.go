package graphs

// Handles Negative Weights)

import (
	"fmt"
	"math"
)

type BFEdge struct {
	from, to, weight int
}

func bellmanFord(n int, edges []BFEdge, src int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if dist[e.from] != math.MaxInt32 && dist[e.from]+e.weight < dist[e.to] {
				dist[e.to] = dist[e.from] + e.weight
			}
		}
	}

	// Optional: detect negative cycles
	for _, e := range edges {
		if dist[e.from] != math.MaxInt32 && dist[e.from]+e.weight < dist[e.to] {
			fmt.Println("Negative weight cycle detected")
			break
		}
	}

	return dist
}

func bellmanFordAlgorithmExample() {
	edges := []BFEdge{
		{0, 1, 4},
		{0, 2, 5},
		{1, 2, -3},
		{2, 3, 4},
	}
	dist := bellmanFord(4, edges, 0)
	fmt.Println("Shortest distances from source:", dist)
}
