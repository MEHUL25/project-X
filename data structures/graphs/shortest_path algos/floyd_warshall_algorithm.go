package graphs

// All-Pairs Shortest Path

import (
	"fmt"
	"math"
)

func floydWarshall(n int, graph [][]int) [][]int {
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else if graph[i][j] != 0 {
				dist[i][j] = graph[i][j]
			} else {
				dist[i][j] = math.MaxInt32 / 2 // to prevent overflow
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func floydWarshallAlgorithmExample() {
	graph := [][]int{
		{0, 3, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 7},
		{2, 0, 0, 0},
	}
	result := floydWarshall(4, graph)
	fmt.Println("All-pairs shortest distances:")
	for _, row := range result {
		fmt.Println(row)
	}
}
