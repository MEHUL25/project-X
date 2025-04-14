package graphs

import (
	"fmt"
)

func dfsTopSort(n int, graph map[int][]int) ([]int, bool) {
	visited := make([]bool, n)
	onStack := make([]bool, n) // for cycle detection
	var result []int
	var hasCycle bool

	var dfs func(int)
	dfs = func(node int) {
		if onStack[node] {
			hasCycle = true
			return
		}
		if visited[node] || hasCycle {
			return
		}
		onStack[node] = true
		visited[node] = true
		for _, neighbor := range graph[node] {
			dfs(neighbor)
		}
		onStack[node] = false
		result = append(result, node)
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	if hasCycle {
		return nil, false
	}
	// reverse result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result, true
}

func dfsExample() {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {4},
		4: {},
	}
	order, ok := dfsTopSort(5, graph)
	if !ok {
		fmt.Println("Cycle detected, topological sort not possible.")
	} else {
		fmt.Println("Topological order (DFS):", order)
	}
}
