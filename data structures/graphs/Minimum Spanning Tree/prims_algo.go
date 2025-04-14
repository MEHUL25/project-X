package graph

import (
	"container/heap"
	"fmt"
)

type PrimEdge struct {
	to     int
	weight int
}

type PrimItem struct {
	node, cost int
}
type PrimPQ []PrimItem

func (pq PrimPQ) Len() int            { return len(pq) }
func (pq PrimPQ) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq PrimPQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PrimPQ) Push(x interface{}) { *pq = append(*pq, x.(PrimItem)) }
func (pq *PrimPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func prim(n int, graph map[int][]PrimEdge) int {
	visited := make([]bool, n)
	pq := &PrimPQ{{0, 0}}
	heap.Init(pq)
	total := 0

	for pq.Len() > 0 {
		item := heap.Pop(pq).(PrimItem)
		if visited[item.node] {
			continue
		}
		visited[item.node] = true
		total += item.cost

		for _, edge := range graph[item.node] {
			if !visited[edge.to] {
				heap.Push(pq, PrimItem{edge.to, edge.weight})
			}
		}
	}

	return total
}

func primAlgoExample() {
	graph := map[int][]PrimEdge{
		0: {{1, 1}, {2, 2}},
		1: {{0, 1}, {2, 4}, {3, 6}},
		2: {{0, 2}, {1, 4}, {3, 3}},
		3: {{1, 6}, {2, 3}},
	}
	fmt.Println("Minimum cost of MST (Prim):", prim(4, graph))
}
