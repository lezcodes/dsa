package prim_algorithm

import (
	"container/heap"
	"errors"
	"fmt"
	"math"
)

type Edge struct {
	From   int
	To     int
	Weight float64
}

type Graph struct {
	vertices int
	adjList  map[int][]Edge
	edges    []Edge
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Edge))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type MST struct {
	edges      []Edge
	totalCost  float64
	vertices   int
	isComplete bool
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]Edge),
		edges:    []Edge{},
	}
}

func (g *Graph) AddEdge(from, to int, weight float64) error {
	if from < 0 || from >= g.vertices || to < 0 || to >= g.vertices {
		return errors.New("vertex index out of range")
	}

	if from == to {
		return errors.New("self-loops not allowed in MST")
	}

	edge1 := Edge{From: from, To: to, Weight: weight}
	edge2 := Edge{From: to, To: from, Weight: weight}

	g.adjList[from] = append(g.adjList[from], edge1)
	g.adjList[to] = append(g.adjList[to], edge2)

	g.edges = append(g.edges, edge1)

	return nil
}

func (g *Graph) GetVertexCount() int {
	return g.vertices
}

func (g *Graph) GetEdgeCount() int {
	return len(g.edges)
}

func (g *Graph) GetNeighbors(vertex int) []Edge {
	if vertex < 0 || vertex >= g.vertices {
		return []Edge{}
	}
	return g.adjList[vertex]
}

func (g *Graph) IsConnected() bool {
	if g.vertices == 0 {
		return true
	}

	visited := make([]bool, g.vertices)
	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for _, edge := range g.adjList[v] {
			if !visited[edge.To] {
				dfs(edge.To)
			}
		}
	}

	dfs(0)

	for i := range g.vertices {
		if !visited[i] {
			return false
		}
	}
	return true
}

func (g *Graph) PrimMST() (*MST, error) {
	if g.vertices == 0 {
		return &MST{edges: []Edge{}, totalCost: 0, vertices: 0, isComplete: true}, nil
	}

	if !g.IsConnected() {
		return nil, errors.New("graph is not connected")
	}

	visited := make([]bool, g.vertices)
	mstEdges := []Edge{}
	totalCost := 0.0

	pq := &PriorityQueue{}
	heap.Init(pq)

	visited[0] = true
	for _, edge := range g.adjList[0] {
		heap.Push(pq, &edge)
	}

	for pq.Len() > 0 && len(mstEdges) < g.vertices-1 {
		edge := heap.Pop(pq).(*Edge)

		if visited[edge.To] {
			continue
		}

		visited[edge.To] = true
		mstEdges = append(mstEdges, *edge)
		totalCost += edge.Weight

		for _, nextEdge := range g.adjList[edge.To] {
			if !visited[nextEdge.To] {
				heap.Push(pq, &nextEdge)
			}
		}
	}

	isComplete := len(mstEdges) == g.vertices-1
	return &MST{
		edges:      mstEdges,
		totalCost:  totalCost,
		vertices:   g.vertices,
		isComplete: isComplete,
	}, nil
}

func (g *Graph) PrimMSTSimple() (*MST, error) {
	if g.vertices == 0 {
		return &MST{edges: []Edge{}, totalCost: 0, vertices: 0, isComplete: true}, nil
	}

	if !g.IsConnected() {
		return nil, errors.New("graph is not connected")
	}

	visited := make([]bool, g.vertices)
	key := make([]float64, g.vertices)
	parent := make([]int, g.vertices)

	for i := range g.vertices {
		key[i] = math.Inf(1)
		parent[i] = -1
	}

	key[0] = 0

	for range g.vertices {
		u := -1
		minKey := math.Inf(1)

		for v := range g.vertices {
			if !visited[v] && key[v] < minKey {
				minKey = key[v]
				u = v
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true

		for _, edge := range g.adjList[u] {
			v := edge.To
			if !visited[v] && edge.Weight < key[v] {
				key[v] = edge.Weight
				parent[v] = u
			}
		}
	}

	mstEdges := []Edge{}
	totalCost := 0.0

	for i := 1; i < g.vertices; i++ {
		if parent[i] != -1 {
			weight := key[i]
			edge := Edge{From: parent[i], To: i, Weight: weight}
			mstEdges = append(mstEdges, edge)
			totalCost += weight
		}
	}

	isComplete := len(mstEdges) == g.vertices-1
	return &MST{
		edges:      mstEdges,
		totalCost:  totalCost,
		vertices:   g.vertices,
		isComplete: isComplete,
	}, nil
}

func (mst *MST) GetEdges() []Edge {
	return mst.edges
}

func (mst *MST) GetTotalCost() float64 {
	return mst.totalCost
}

func (mst *MST) GetVertexCount() int {
	return mst.vertices
}

func (mst *MST) IsComplete() bool {
	return mst.isComplete
}

func (mst *MST) GetEdgeCount() int {
	return len(mst.edges)
}

func (g *Graph) PrintGraph() {
	fmt.Printf("Graph with %d vertices:\n", g.vertices)
	for i := range g.vertices {
		fmt.Printf("Vertex %d: ", i)
		neighbors := g.GetNeighbors(i)
		if len(neighbors) == 0 {
			fmt.Println("[]")
		} else {
			fmt.Print("[")
			for j, edge := range neighbors {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("(%d, %.1f)", edge.To, edge.Weight)
			}
			fmt.Println("]")
		}
	}
}

func (mst *MST) PrintMST() {
	fmt.Printf("Minimum Spanning Tree:\n")
	fmt.Printf("Total Cost: %.2f\n", mst.totalCost)
	fmt.Printf("Edges (%d):\n", len(mst.edges))
	for i, edge := range mst.edges {
		fmt.Printf("  %d: %d -- %d (%.2f)\n", i+1, edge.From, edge.To, edge.Weight)
	}
}

func Run() any {
	g := NewGraph(6)

	edges := []struct {
		from, to int
		weight   float64
	}{
		{0, 1, 4}, {0, 2, 2}, {1, 2, 1}, {1, 3, 5},
		{2, 3, 8}, {2, 4, 10}, {3, 4, 2}, {3, 5, 6}, {4, 5, 3},
	}

	for _, e := range edges {
		g.AddEdge(e.from, e.to, e.weight)
	}

	result := make(map[string]any)
	result["vertexCount"] = g.GetVertexCount()
	result["edgeCount"] = g.GetEdgeCount()
	result["isConnected"] = g.IsConnected()

	primMST, primErr := g.PrimMST()
	if primErr == nil {
		result["primMSTCost"] = primMST.GetTotalCost()
		result["primMSTEdgeCount"] = primMST.GetEdgeCount()
		result["primMSTComplete"] = primMST.IsComplete()

		primEdges := make([]map[string]any, len(primMST.GetEdges()))
		for i, edge := range primMST.GetEdges() {
			primEdges[i] = map[string]any{
				"from":   edge.From,
				"to":     edge.To,
				"weight": edge.Weight,
			}
		}
		result["primMSTEdges"] = primEdges
	} else {
		result["primMSTError"] = primErr.Error()
	}

	simpleMST, simpleErr := g.PrimMSTSimple()
	if simpleErr == nil {
		result["simpleMSTCost"] = simpleMST.GetTotalCost()
		result["simpleMSTEdgeCount"] = simpleMST.GetEdgeCount()
		result["simpleMSTComplete"] = simpleMST.IsComplete()

		simpleEdges := make([]map[string]any, len(simpleMST.GetEdges()))
		for i, edge := range simpleMST.GetEdges() {
			simpleEdges[i] = map[string]any{
				"from":   edge.From,
				"to":     edge.To,
				"weight": edge.Weight,
			}
		}
		result["simpleMSTEdges"] = simpleEdges
	} else {
		result["simpleMSTError"] = simpleErr.Error()
	}

	disconnectedGraph := NewGraph(4)
	disconnectedGraph.AddEdge(0, 1, 1)
	disconnectedGraph.AddEdge(2, 3, 2)

	result["disconnectedGraphConnected"] = disconnectedGraph.IsConnected()
	_, disconnectedErr := disconnectedGraph.PrimMST()
	result["disconnectedMSTError"] = disconnectedErr != nil
	if disconnectedErr != nil {
		result["disconnectedErrorMessage"] = disconnectedErr.Error()
	}

	singleVertexGraph := NewGraph(1)
	singleMST, _ := singleVertexGraph.PrimMST()
	result["singleVertexMSTCost"] = singleMST.GetTotalCost()
	result["singleVertexMSTComplete"] = singleMST.IsComplete()

	return result
}
