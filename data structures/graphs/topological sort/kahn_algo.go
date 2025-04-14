package graphs

import (
	"fmt"
)

func topologicalSortKahn(n int, graph map[int][]int) ([]int, bool) {
	indegree := make([]int, n)
	for _, neighbors := range graph {
		for _, v := range neighbors {
			indegree[v]++
		}
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var result []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != n {
		return nil, false // Cycle detected
	}
	return result, true
}

func kahnExample() {
	graph := map[int][]int{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {4},
		4: {},
	}
	order, ok := topologicalSortKahn(5, graph)
	if !ok {
		fmt.Println("Cycle detected, topological sort not possible.")
	} else {
		fmt.Println("Topological order (Kahn):", order)
	}
}
