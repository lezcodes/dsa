package dijkstra_algorithm

import (
	"container/heap"
	"math"
)

type Edge struct {
	To     int
	Weight int
}

type Graph struct {
	Vertices int
	AdjList  [][]Edge
}

type Item struct {
	Vertex   int
	Distance int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

type DijkstraResult struct {
	Distances []int
	Previous  []int
	Source    int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make([][]Edge, vertices),
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	if from >= 0 && from < g.Vertices && to >= 0 && to < g.Vertices {
		g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
	}
}

func (g *Graph) AddBidirectionalEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

func (g *Graph) Dijkstra(source int) *DijkstraResult {
	if source < 0 || source >= g.Vertices {
		return nil
	}

	distances := make([]int, g.Vertices)
	previous := make([]int, g.Vertices)
	visited := make([]bool, g.Vertices)

	for i := range g.Vertices {
		distances[i] = math.MaxInt32
		previous[i] = -1
	}

	distances[source] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Vertex: source, Distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)

		if visited[current.Vertex] {
			continue
		}

		visited[current.Vertex] = true

		for _, edge := range g.AdjList[current.Vertex] {
			if visited[edge.To] {
				continue
			}

			newDistance := distances[current.Vertex] + edge.Weight

			if newDistance < distances[edge.To] {
				distances[edge.To] = newDistance
				previous[edge.To] = current.Vertex
				heap.Push(&pq, &Item{Vertex: edge.To, Distance: newDistance})
			}
		}
	}

	return &DijkstraResult{
		Distances: distances,
		Previous:  previous,
		Source:    source,
	}
}

func (r *DijkstraResult) GetPath(target int) []int {
	if target < 0 || target >= len(r.Distances) || r.Distances[target] == math.MaxInt32 {
		return nil
	}

	path := []int{}
	current := target

	for current != -1 {
		path = append([]int{current}, path...)
		current = r.Previous[current]
	}

	if len(path) > 0 && path[0] == r.Source {
		return path
	}

	return nil
}

func (r *DijkstraResult) GetDistance(target int) int {
	if target < 0 || target >= len(r.Distances) {
		return math.MaxInt32
	}
	return r.Distances[target]
}

func (r *DijkstraResult) HasPath(target int) bool {
	return target >= 0 && target < len(r.Distances) && r.Distances[target] != math.MaxInt32
}

func Run() any {
	graph := NewGraph(6)

	graph.AddBidirectionalEdge(0, 1, 4)
	graph.AddBidirectionalEdge(0, 2, 3)
	graph.AddBidirectionalEdge(1, 2, 1)
	graph.AddBidirectionalEdge(1, 3, 2)
	graph.AddBidirectionalEdge(2, 3, 4)
	graph.AddBidirectionalEdge(3, 4, 2)
	graph.AddBidirectionalEdge(4, 5, 6)

	result := graph.Dijkstra(0)

	shortestPaths := make(map[string]any)
	shortestPaths["source"] = 0
	shortestPaths["distances"] = result.Distances

	paths := make(map[int][]int)
	for i := 1; i < 6; i++ {
		if result.HasPath(i) {
			paths[i] = result.GetPath(i)
		}
	}
	shortestPaths["paths"] = paths

	return shortestPaths
}
