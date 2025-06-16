package prim_algorithm

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

	neighbors := g.GetNeighbors(0)
	if len(neighbors) != 1 || neighbors[0].To != 1 || neighbors[0].Weight != 2.5 {
		t.Errorf("Expected neighbor (1, 2.5), got %v", neighbors)
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

func TestPrimMSTSimple(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 3)

	mst, err := g.PrimMST()
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

func TestPrimMSTSimpleAlgorithm(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 3)

	mst, err := g.PrimMSTSimple()
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

func TestBothAlgorithmsConsistency(t *testing.T) {
	g := NewGraph(5)
	edges := []struct {
		from, to int
		weight   float64
	}{
		{0, 1, 2}, {0, 3, 6}, {1, 2, 3}, {1, 3, 8}, {1, 4, 5}, {2, 4, 7}, {3, 4, 9},
	}

	for _, e := range edges {
		g.AddEdge(e.from, e.to, e.weight)
	}

	primMST, primErr := g.PrimMST()
	simpleMST, simpleErr := g.PrimMSTSimple()

	if primErr != nil || simpleErr != nil {
		t.Errorf("Errors: prim=%v, simple=%v", primErr, simpleErr)
	}

	if primMST.GetTotalCost() != simpleMST.GetTotalCost() {
		t.Errorf("Different costs: prim=%.1f, simple=%.1f",
			primMST.GetTotalCost(), simpleMST.GetTotalCost())
	}

	if primMST.GetEdgeCount() != simpleMST.GetEdgeCount() {
		t.Errorf("Different edge counts: prim=%d, simple=%d",
			primMST.GetEdgeCount(), simpleMST.GetEdgeCount())
	}
}

func TestDisconnectedGraph(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 3, 2)

	_, err := g.PrimMST()
	if err == nil {
		t.Error("Expected error for disconnected graph")
	}

	_, err = g.PrimMSTSimple()
	if err == nil {
		t.Error("Expected error for disconnected graph")
	}
}

func TestSingleVertex(t *testing.T) {
	g := NewGraph(1)

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

	mst, err := g.PrimMST()
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

func BenchmarkPrimMST(b *testing.B) {
	g := NewGraph(1000)
	for i := range 999 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		g.PrimMST()
	}
}

func BenchmarkPrimMSTSimple(b *testing.B) {
	g := NewGraph(1000)
	for i := range 999 {
		g.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		g.PrimMSTSimple()
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
