package graphs

import (
	"fmt"
)

func countComponents(n int, edges [][]int) int {
	// Step 1: Build adjacency list
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Step 2: Visited tracker
	visited := make([]bool, n)

	// Step 3: BFS function
	bfs := func(start int) {
		queue := []int{start}
		visited[start] = true

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
	}

	// Step 4: Count connected components
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(i)
			count++
		}
	}

	return count
}

func main() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}

	fmt.Println("Number of connected components:", countComponents(n, edges))
}
