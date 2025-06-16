package kruskal_algorithm

import (
	"errors"
	"fmt"
	"sort"
)

type Edge struct {
	From   int
	To     int
	Weight float64
}

type Graph struct {
	vertices int
	edges    []Edge
}

type UnionFind struct {
	parent []int
	rank   []int
	size   int
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
		edges:    []Edge{},
	}
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		size:   n,
	}

	for i := range n {
		uf.parent[i] = i
		uf.rank[i] = 0
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	uf.size--
	return true
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) ComponentCount() int {
	return uf.size
}

func (g *Graph) AddEdge(from, to int, weight float64) error {
	if from < 0 || from >= g.vertices || to < 0 || to >= g.vertices {
		return errors.New("vertex index out of range")
	}

	if from == to {
		return errors.New("self-loops not allowed in MST")
	}

	edge := Edge{From: from, To: to, Weight: weight}
	g.edges = append(g.edges, edge)

	return nil
}

func (g *Graph) GetVertexCount() int {
	return g.vertices
}

func (g *Graph) GetEdgeCount() int {
	return len(g.edges)
}

func (g *Graph) GetEdges() []Edge {
	return g.edges
}

func (g *Graph) IsConnected() bool {
	if g.vertices == 0 {
		return true
	}

	uf := NewUnionFind(g.vertices)

	for _, edge := range g.edges {
		uf.Union(edge.From, edge.To)
	}

	return uf.ComponentCount() == 1
}

func (g *Graph) KruskalMST() (*MST, error) {
	if g.vertices == 0 {
		return &MST{edges: []Edge{}, totalCost: 0, vertices: 0, isComplete: true}, nil
	}

	if !g.IsConnected() {
		return nil, errors.New("graph is not connected")
	}

	sortedEdges := make([]Edge, len(g.edges))
	copy(sortedEdges, g.edges)

	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i].Weight < sortedEdges[j].Weight
	})

	uf := NewUnionFind(g.vertices)
	mstEdges := []Edge{}
	totalCost := 0.0

	for _, edge := range sortedEdges {
		if uf.Union(edge.From, edge.To) {
			mstEdges = append(mstEdges, edge)
			totalCost += edge.Weight

			if len(mstEdges) == g.vertices-1 {
				break
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

func (g *Graph) KruskalMSTWithSteps() (*MST, []KruskalStep, error) {
	if g.vertices == 0 {
		return &MST{edges: []Edge{}, totalCost: 0, vertices: 0, isComplete: true}, []KruskalStep{}, nil
	}

	if !g.IsConnected() {
		return nil, nil, errors.New("graph is not connected")
	}

	sortedEdges := make([]Edge, len(g.edges))
	copy(sortedEdges, g.edges)

	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i].Weight < sortedEdges[j].Weight
	})

	uf := NewUnionFind(g.vertices)
	mstEdges := []Edge{}
	totalCost := 0.0
	steps := []KruskalStep{}

	for _, edge := range sortedEdges {
		step := KruskalStep{
			Edge:     edge,
			Accepted: false,
			Reason:   "",
		}

		if uf.Connected(edge.From, edge.To) {
			step.Reason = "creates cycle"
		} else {
			uf.Union(edge.From, edge.To)
			mstEdges = append(mstEdges, edge)
			totalCost += edge.Weight
			step.Accepted = true
			step.Reason = "added to MST"

			if len(mstEdges) == g.vertices-1 {
				step.Reason = "added to MST (complete)"
			}
		}

		steps = append(steps, step)

		if len(mstEdges) == g.vertices-1 {
			for i := len(steps); i < len(sortedEdges); i++ {
				remainingEdge := sortedEdges[i]
				remainingStep := KruskalStep{
					Edge:     remainingEdge,
					Accepted: false,
					Reason:   "creates cycle",
				}
				steps = append(steps, remainingStep)
			}
			break
		}
	}

	isComplete := len(mstEdges) == g.vertices-1
	mst := &MST{
		edges:      mstEdges,
		totalCost:  totalCost,
		vertices:   g.vertices,
		isComplete: isComplete,
	}

	return mst, steps, nil
}

type KruskalStep struct {
	Edge     Edge
	Accepted bool
	Reason   string
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
	fmt.Printf("Graph with %d vertices and %d edges:\n", g.vertices, len(g.edges))
	for i, edge := range g.edges {
		fmt.Printf("  %d: %d -- %d (%.2f)\n", i+1, edge.From, edge.To, edge.Weight)
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

func PrintKruskalSteps(steps []KruskalStep) {
	fmt.Println("Kruskal's Algorithm Steps:")
	for i, step := range steps {
		status := "REJECTED"
		if step.Accepted {
			status = "ACCEPTED"
		}
		fmt.Printf("  %d: Edge (%d--%d, %.2f) %s - %s\n",
			i+1, step.Edge.From, step.Edge.To, step.Edge.Weight, status, step.Reason)
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

	kruskalMST, kruskalErr := g.KruskalMST()
	if kruskalErr == nil {
		result["kruskalMSTCost"] = kruskalMST.GetTotalCost()
		result["kruskalMSTEdgeCount"] = kruskalMST.GetEdgeCount()
		result["kruskalMSTComplete"] = kruskalMST.IsComplete()

		kruskalEdges := make([]map[string]any, len(kruskalMST.GetEdges()))
		for i, edge := range kruskalMST.GetEdges() {
			kruskalEdges[i] = map[string]any{
				"from":   edge.From,
				"to":     edge.To,
				"weight": edge.Weight,
			}
		}
		result["kruskalMSTEdges"] = kruskalEdges
	} else {
		result["kruskalMSTError"] = kruskalErr.Error()
	}

	stepsMST, steps, stepsErr := g.KruskalMSTWithSteps()
	if stepsErr == nil {
		result["stepsMSTCost"] = stepsMST.GetTotalCost()
		result["stepsCount"] = len(steps)

		stepsData := make([]map[string]any, len(steps))
		for i, step := range steps {
			stepsData[i] = map[string]any{
				"edge": map[string]any{
					"from":   step.Edge.From,
					"to":     step.Edge.To,
					"weight": step.Edge.Weight,
				},
				"accepted": step.Accepted,
				"reason":   step.Reason,
			}
		}
		result["algorithmSteps"] = stepsData
	}

	disconnectedGraph := NewGraph(4)
	disconnectedGraph.AddEdge(0, 1, 1)
	disconnectedGraph.AddEdge(2, 3, 2)

	result["disconnectedGraphConnected"] = disconnectedGraph.IsConnected()
	_, disconnectedErr := disconnectedGraph.KruskalMST()
	result["disconnectedMSTError"] = disconnectedErr != nil
	if disconnectedErr != nil {
		result["disconnectedErrorMessage"] = disconnectedErr.Error()
	}

	singleVertexGraph := NewGraph(1)
	singleMST, _ := singleVertexGraph.KruskalMST()
	result["singleVertexMSTCost"] = singleMST.GetTotalCost()
	result["singleVertexMSTComplete"] = singleMST.IsComplete()

	uf := NewUnionFind(5)
	result["unionFindInitialComponents"] = uf.ComponentCount()
	uf.Union(0, 1)
	uf.Union(2, 3)
	result["unionFindAfterUnions"] = uf.ComponentCount()
	result["unionFind0And1Connected"] = uf.Connected(0, 1)
	result["unionFind0And2Connected"] = uf.Connected(0, 2)

	return result
}
