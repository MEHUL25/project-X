package graphs

import (
	"fmt"
)

func countComponents(n int, edges [][]int) int {
	// Step 1: Build adjacency list
	graph := make(map[int][]int)
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Step 2: Track visited nodes
	visited := make([]bool, n)

	// Step 3: DFS function
	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Step 4: Count connected components
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}

func dfsExample() {
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}

	fmt.Println("Number of connected components:", countComponents(n, edges))
}
