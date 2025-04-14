package graphs

// Shortest Path from Source in Weighted Graph

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node, dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func dijkstra(n int, graph map[int][]Edge, src int) []int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[src] = 0

	pq := &PriorityQueue{{src, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(Item)
		if curr.dist > dist[curr.node] {
			continue
		}
		for _, edge := range graph[curr.node] {
			if dist[curr.node]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[curr.node] + edge.weight
				heap.Push(pq, Item{edge.to, dist[edge.to]})
			}
		}
	}
	return dist
}

func dijstraExample() {
	graph := map[int][]Edge{
		0: {{1, 2}, {2, 4}},
		1: {{2, 1}, {3, 7}},
		2: {{4, 3}},
		3: {{5, 1}},
		4: {{3, 2}, {5, 5}},
		5: {},
	}
	distances := dijkstra(6, graph, 0)
	fmt.Println("Shortest distances from source:", distances)
}
