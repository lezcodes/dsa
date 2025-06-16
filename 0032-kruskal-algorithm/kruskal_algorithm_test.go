package kruskal_algorithm

import (
	"math"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	if g.GetVertexCount() != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.GetVertexCount())
	}
	if g.GetEdgeCount() != 0 {
		t.Errorf("Expected 0 edges, got %d", g.GetEdgeCount())
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(3)

	err := g.AddEdge(0, 1, 2.5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if g.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", g.GetEdgeCount())
	}

	edges := g.GetEdges()
	if len(edges) != 1 || edges[0].From != 0 || edges[0].To != 1 || edges[0].Weight != 2.5 {
		t.Errorf("Expected edge (0->1, 2.5), got %v", edges[0])
	}
}

func TestAddEdgeInvalidVertex(t *testing.T) {
	g := NewGraph(3)

	err := g.AddEdge(0, 5, 1.0)
	if err == nil {
		t.Error("Expected error for invalid vertex")
	}

	err = g.AddEdge(-1, 1, 1.0)
	if err == nil {
		t.Error("Expected error for negative vertex")
	}
}

func TestAddSelfLoop(t *testing.T) {
	g := NewGraph(3)

	err := g.AddEdge(0, 0, 1.0)
	if err == nil {
		t.Error("Expected error for self-loop")
	}
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(5)

	if uf.ComponentCount() != 5 {
		t.Errorf("Expected 5 components initially, got %d", uf.ComponentCount())
	}

	if uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to be disconnected initially")
	}

	success := uf.Union(0, 1)
	if !success {
		t.Error("Expected successful union of 0 and 1")
	}

	if !uf.Connected(0, 1) {
		t.Error("Expected 0 and 1 to be connected after union")
	}

	if uf.ComponentCount() != 4 {
		t.Errorf("Expected 4 components after union, got %d", uf.ComponentCount())
	}

	success = uf.Union(0, 1)
	if success {
		t.Error("Expected unsuccessful union of already connected components")
	}
}

func TestUnionFindPathCompression(t *testing.T) {
	uf := NewUnionFind(4)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	root0 := uf.Find(0)
	root3 := uf.Find(3)

	if root0 != root3 {
		t.Error("Expected all elements to have same root after chaining")
	}

	if uf.ComponentCount() != 1 {
		t.Errorf("Expected 1 component, got %d", uf.ComponentCount())
	}
}

func TestIsConnected(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 3)

	if !g.IsConnected() {
		t.Error("Expected connected graph")
	}

	disconnected := NewGraph(4)
	disconnected.AddEdge(0, 1, 1)
	disconnected.AddEdge(2, 3, 2)

	if disconnected.IsConnected() {
		t.Error("Expected disconnected graph")
	}
}

func TestKruskalMST(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 3)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !mst.IsComplete() {
		t.Error("Expected complete MST")
	}

	if mst.GetEdgeCount() != 3 {
		t.Errorf("Expected 3 edges in MST, got %d", mst.GetEdgeCount())
	}

	expectedCost := 6.0
	if mst.GetTotalCost() != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.GetTotalCost())
	}
}

func TestKruskalMSTWithSteps(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 3)

	mst, steps, err := g.KruskalMSTWithSteps()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !mst.IsComplete() {
		t.Error("Expected complete MST")
	}

	if len(steps) == 0 {
		t.Error("Expected algorithm steps")
	}

	acceptedCount := 0
	for _, step := range steps {
		if step.Accepted {
			acceptedCount++
		}
	}

	if acceptedCount != 3 {
		t.Errorf("Expected 3 accepted edges, got %d", acceptedCount)
	}
}

func TestDisconnectedGraph(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 3, 2)

	_, err := g.KruskalMST()
	if err == nil {
		t.Error("Expected error for disconnected graph")
	}

	_, _, err = g.KruskalMSTWithSteps()
	if err == nil {
		t.Error("Expected error for disconnected graph with steps")
	}
}

func TestSingleVertex(t *testing.T) {
	g := NewGraph(1)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error for single vertex, got %v", err)
	}

	if mst.GetTotalCost() != 0 {
		t.Errorf("Expected cost 0 for single vertex, got %.1f", mst.GetTotalCost())
	}

	if !mst.IsComplete() {
		t.Error("Single vertex MST should be complete")
	}
}

func TestEmptyGraph(t *testing.T) {
	g := NewGraph(0)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error for empty graph, got %v", err)
	}

	if mst.GetTotalCost() != 0 {
		t.Errorf("Expected cost 0 for empty graph, got %.1f", mst.GetTotalCost())
	}

	if !mst.IsComplete() {
		t.Error("Empty graph MST should be complete")
	}
}

func TestLinearGraph(t *testing.T) {
	g := NewGraph(5)
	for i := range 4 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedCost := 10.0
	if mst.GetTotalCost() != expectedCost {
		t.Errorf("Expected cost %.1f, got %.1f", expectedCost, mst.GetTotalCost())
	}

	if mst.GetEdgeCount() != 4 {
		t.Errorf("Expected 4 edges, got %d", mst.GetEdgeCount())
	}
}

func TestCompleteGraph(t *testing.T) {
	n := 4
	g := NewGraph(n)

	for i := range n {
		for j := i + 1; j < n; j++ {
			weight := float64(i*n + j)
			g.AddEdge(i, j, weight)
		}
	}

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if mst.GetEdgeCount() != n-1 {
		t.Errorf("Expected %d edges, got %d", n-1, mst.GetEdgeCount())
	}

	if !mst.IsComplete() {
		t.Error("Complete graph MST should be complete")
	}
}

func TestFloatingPointWeights(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.5)
	g.AddEdge(1, 2, 2.7)
	g.AddEdge(0, 2, 3.2)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedCost := 4.2
	if math.Abs(mst.GetTotalCost()-expectedCost) > 0.001 {
		t.Errorf("Expected cost %.1f, got %.1f", expectedCost, mst.GetTotalCost())
	}
}

func TestNegativeWeights(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, -1.0)
	g.AddEdge(1, 2, 2.0)
	g.AddEdge(0, 2, 3.0)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error with negative weights, got %v", err)
	}

	expectedCost := 1.0
	if mst.GetTotalCost() != expectedCost {
		t.Errorf("Expected cost %.1f, got %.1f", expectedCost, mst.GetTotalCost())
	}
}

func TestLargeGraph(t *testing.T) {
	n := 100
	g := NewGraph(n)

	for i := range n - 1 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for i := range n / 2 {
		g.AddEdge(i, i+n/2, float64(n+i))
	}

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error for large graph, got %v", err)
	}

	if mst.GetEdgeCount() != n-1 {
		t.Errorf("Expected %d edges, got %d", n-1, mst.GetEdgeCount())
	}

	if !mst.IsComplete() {
		t.Error("Large graph MST should be complete")
	}
}

func TestMSTProperties(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 3)
	g.AddEdge(0, 3, 10)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if mst.GetVertexCount() != 4 {
		t.Errorf("Expected 4 vertices, got %d", mst.GetVertexCount())
	}

	edges := mst.GetEdges()
	if len(edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(edges))
	}

	totalWeight := 0.0
	for _, edge := range edges {
		totalWeight += edge.Weight
		if edge.From < 0 || edge.From >= 4 || edge.To < 0 || edge.To >= 4 {
			t.Errorf("Invalid edge vertices: %d -> %d", edge.From, edge.To)
		}
	}

	if totalWeight != mst.GetTotalCost() {
		t.Errorf("Inconsistent total cost: edges=%.1f, mst=%.1f",
			totalWeight, mst.GetTotalCost())
	}
}

func TestEdgeCases(t *testing.T) {
	g := NewGraph(2)
	g.AddEdge(0, 1, 5.0)

	mst, err := g.KruskalMST()
	if err != nil {
		t.Errorf("Expected no error for two vertices, got %v", err)
	}

	if mst.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", mst.GetEdgeCount())
	}

	if mst.GetTotalCost() != 5.0 {
		t.Errorf("Expected cost 5.0, got %.1f", mst.GetTotalCost())
	}
}

func TestEdgeSorting(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 3)
	g.AddEdge(0, 3, 2)

	mst, steps, err := g.KruskalMSTWithSteps()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(steps) == 0 {
		t.Error("Expected algorithm steps")
	}

	for i := 1; i < len(steps); i++ {
		if steps[i-1].Edge.Weight > steps[i].Edge.Weight {
			t.Error("Expected edges to be processed in sorted order")
		}
	}

	expectedCost := 6.0
	if mst.GetTotalCost() != expectedCost {
		t.Errorf("Expected cost %.1f, got %.1f", expectedCost, mst.GetTotalCost())
	}
}

func TestCycleDetection(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(0, 2, 3)

	_, steps, err := g.KruskalMSTWithSteps()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	rejectedFound := false
	for _, step := range steps {
		if !step.Accepted && step.Reason == "creates cycle" {
			rejectedFound = true
			break
		}
	}

	if !rejectedFound {
		t.Error("Expected at least one edge to be rejected due to cycle")
	}
}

func TestUnionByRank(t *testing.T) {
	uf := NewUnionFind(8)

	uf.Union(0, 1)
	uf.Union(2, 3)
	uf.Union(4, 5)
	uf.Union(6, 7)

	uf.Union(0, 2)
	uf.Union(4, 6)

	uf.Union(0, 4)

	if uf.ComponentCount() != 1 {
		t.Errorf("Expected 1 component, got %d", uf.ComponentCount())
	}

	for i := range 8 {
		for j := i + 1; j < 8; j++ {
			if !uf.Connected(i, j) {
				t.Errorf("Expected %d and %d to be connected", i, j)
			}
		}
	}
}

func BenchmarkKruskalMST(b *testing.B) {
	g := NewGraph(1000)
	for i := range 999 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		g.KruskalMST()
	}
}

func BenchmarkKruskalMSTWithSteps(b *testing.B) {
	g := NewGraph(500)
	for i := range 499 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		g.KruskalMSTWithSteps()
	}
}

func BenchmarkUnionFind(b *testing.B) {
	uf := NewUnionFind(10000)

	for i := range b.N {
		x := i % 9999
		y := (i + 1) % 10000
		uf.Union(x, y)
	}
}

func BenchmarkAddEdge(b *testing.B) {
	g := NewGraph(10000)

	for i := range b.N {
		from := i % 9999
		to := (i + 1) % 10000
		if from != to {
			g.AddEdge(from, to, float64(i))
		}
	}
}

func BenchmarkIsConnected(b *testing.B) {
	g := NewGraph(1000)
	for i := range 999 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		g.IsConnected()
	}
}
