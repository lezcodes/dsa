package ford_fulkerson_algorithm

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

func TestNewFlowNetwork(t *testing.T) {
	fn := NewFlowNetwork(5)
	if fn.GetVertexCount() != 5 {
		t.Errorf("Expected 5 vertices, got %d", fn.GetVertexCount())
	}
	if fn.GetEdgeCount() != 0 {
		t.Errorf("Expected 0 edges, got %d", fn.GetEdgeCount())
	}
}

func TestAddEdge(t *testing.T) {
	fn := NewFlowNetwork(3)

	err := fn.AddEdge(0, 1, 10.5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if fn.GetEdgeCount() != 2 {
		t.Errorf("Expected 2 edges (forward and reverse), got %d", fn.GetEdgeCount())
	}

	edges := fn.GetEdges()
	if len(edges) < 1 || edges[0].From != 0 || edges[0].To != 1 || edges[0].Capacity != 10.5 {
		t.Errorf("Expected edge (0->1, 10.5), got %v", edges[0])
	}
}

func TestAddEdgeInvalidVertex(t *testing.T) {
	fn := NewFlowNetwork(3)

	err := fn.AddEdge(0, 5, 10)
	if err == nil {
		t.Error("Expected error for invalid vertex")
	}

	err = fn.AddEdge(-1, 1, 10)
	if err == nil {
		t.Error("Expected error for negative vertex")
	}
}

func TestAddEdgeNegativeCapacity(t *testing.T) {
	fn := NewFlowNetwork(3)

	err := fn.AddEdge(0, 1, -5)
	if err == nil {
		t.Error("Expected error for negative capacity")
	}
}

func TestAddDuplicateEdge(t *testing.T) {
	fn := NewFlowNetwork(3)

	fn.AddEdge(0, 1, 10)
	initialCount := fn.GetEdgeCount()

	fn.AddEdge(0, 1, 5)

	if fn.GetEdgeCount() != initialCount {
		t.Error("Expected same edge count after adding duplicate")
	}

	edges := fn.GetEdges()
	found := false
	for _, edge := range edges {
		if edge.From == 0 && edge.To == 1 && edge.Capacity == 15 {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected capacities to be summed for duplicate edges")
	}
}

func TestSimpleMaxFlow(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 20)
	fn.AddEdge(0, 2, 10)
	fn.AddEdge(1, 2, 30)
	fn.AddEdge(1, 3, 10)
	fn.AddEdge(2, 3, 20)

	result, err := fn.FordFulkersonDFS(0, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedFlow := 30.0
	if result.GetMaxFlow() != expectedFlow {
		t.Errorf("Expected max flow %.1f, got %.1f", expectedFlow, result.GetMaxFlow())
	}

	if result.GetMaxFlow() != result.GetMinCutCapacity() {
		t.Errorf("Max flow (%.1f) should equal min cut capacity (%.1f)",
			result.GetMaxFlow(), result.GetMinCutCapacity())
	}
}

func TestBothAlgorithmsConsistency(t *testing.T) {
	fn := NewFlowNetwork(6)
	edges := []struct {
		from, to int
		capacity float64
	}{
		{0, 1, 16}, {0, 2, 13}, {1, 2, 10}, {1, 3, 12},
		{2, 1, 4}, {2, 4, 14}, {3, 2, 9}, {3, 5, 20},
		{4, 3, 7}, {4, 5, 4},
	}

	for _, e := range edges {
		fn.AddEdge(e.from, e.to, e.capacity)
	}

	dfsResult, dfsErr := fn.FordFulkersonDFS(0, 5)
	bfsResult, bfsErr := fn.FordFulkersonBFS(0, 5)

	if dfsErr != nil || bfsErr != nil {
		t.Errorf("Errors: dfs=%v, bfs=%v", dfsErr, bfsErr)
	}

	if dfsResult.GetMaxFlow() != bfsResult.GetMaxFlow() {
		t.Errorf("Different max flows: dfs=%.1f, bfs=%.1f",
			dfsResult.GetMaxFlow(), bfsResult.GetMaxFlow())
	}

	if dfsResult.GetMinCutCapacity() != bfsResult.GetMinCutCapacity() {
		t.Errorf("Different min cut capacities: dfs=%.1f, bfs=%.1f",
			dfsResult.GetMinCutCapacity(), bfsResult.GetMinCutCapacity())
	}
}

func TestInvalidSourceSink(t *testing.T) {
	fn := NewFlowNetwork(3)
	fn.AddEdge(0, 1, 10)

	_, err := fn.FordFulkersonDFS(0, 0)
	if err == nil {
		t.Error("Expected error for same source and sink")
	}

	_, err = fn.FordFulkersonDFS(-1, 2)
	if err == nil {
		t.Error("Expected error for invalid source")
	}

	_, err = fn.FordFulkersonDFS(0, 5)
	if err == nil {
		t.Error("Expected error for invalid sink")
	}
}

func TestDisconnectedGraph(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(2, 3, 5)

	result, err := fn.FordFulkersonDFS(0, 3)
	if err != nil {
		t.Errorf("Expected no error for disconnected graph, got %v", err)
	}

	if result.GetMaxFlow() != 0 {
		t.Errorf("Expected max flow 0 for disconnected graph, got %.1f", result.GetMaxFlow())
	}
}

func TestSingleEdge(t *testing.T) {
	fn := NewFlowNetwork(2)
	fn.AddEdge(0, 1, 15)

	result, err := fn.FordFulkersonDFS(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.GetMaxFlow() != 15 {
		t.Errorf("Expected max flow 15, got %.1f", result.GetMaxFlow())
	}
}

func TestLinearPath(t *testing.T) {
	fn := NewFlowNetwork(5)
	capacities := []float64{10, 5, 15, 8}

	for i := range 4 {
		fn.AddEdge(i, i+1, capacities[i])
	}

	result, err := fn.FordFulkersonDFS(0, 4)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedFlow := 5.0
	if result.GetMaxFlow() != expectedFlow {
		t.Errorf("Expected max flow %.1f, got %.1f", expectedFlow, result.GetMaxFlow())
	}
}

func TestParallelEdges(t *testing.T) {
	fn := NewFlowNetwork(2)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(0, 1, 5)

	result, err := fn.FordFulkersonDFS(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.GetMaxFlow() != 15 {
		t.Errorf("Expected max flow 15, got %.1f", result.GetMaxFlow())
	}
}

func TestComplexNetwork(t *testing.T) {
	fn := NewFlowNetwork(6)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(0, 2, 10)
	fn.AddEdge(1, 3, 25)
	fn.AddEdge(2, 3, 6)
	fn.AddEdge(2, 4, 10)
	fn.AddEdge(3, 5, 10)
	fn.AddEdge(4, 5, 10)

	result, err := fn.FordFulkersonDFS(0, 5)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedFlow := 20.0
	if result.GetMaxFlow() != expectedFlow {
		t.Errorf("Expected max flow %.1f, got %.1f", expectedFlow, result.GetMaxFlow())
	}

	if len(result.GetMinCut()) == 0 {
		t.Error("Expected non-empty min cut")
	}
}

func TestFloatingPointCapacities(t *testing.T) {
	fn := NewFlowNetwork(3)
	fn.AddEdge(0, 1, 10.5)
	fn.AddEdge(1, 2, 7.3)
	fn.AddEdge(0, 2, 5.8)

	result, err := fn.FordFulkersonDFS(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedFlow := 13.1
	if math.Abs(result.GetMaxFlow()-expectedFlow) > 0.001 {
		t.Errorf("Expected max flow %.1f, got %.1f", expectedFlow, result.GetMaxFlow())
	}
}

func TestZeroCapacityEdges(t *testing.T) {
	fn := NewFlowNetwork(3)
	fn.AddEdge(0, 1, 0)
	fn.AddEdge(1, 2, 10)
	fn.AddEdge(0, 2, 5)

	result, err := fn.FordFulkersonDFS(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.GetMaxFlow() != 5 {
		t.Errorf("Expected max flow 5, got %.1f", result.GetMaxFlow())
	}
}

func TestResidualCapacity(t *testing.T) {
	fn := NewFlowNetwork(2)
	fn.AddEdge(0, 1, 10)

	edges := fn.GetEdges()
	if len(edges) < 1 {
		t.Fatal("Expected at least one edge")
	}

	residualCap := fn.GetResidualCapacity(0)
	if residualCap != 10 {
		t.Errorf("Expected residual capacity 10, got %.1f", residualCap)
	}

	fn.edges[0].Flow = 3
	residualCap = fn.GetResidualCapacity(0)
	if residualCap != 7 {
		t.Errorf("Expected residual capacity 7, got %.1f", residualCap)
	}
}

func TestAugmentingPathDFS(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(1, 2, 5)
	fn.AddEdge(2, 3, 15)

	path := fn.FindAugmentingPathDFS(0, 3)
	if path == nil {
		t.Error("Expected to find augmenting path")
	}

	if path.Bottleneck != 5 {
		t.Errorf("Expected bottleneck 5, got %.1f", path.Bottleneck)
	}

	if len(path.Path) != 4 {
		t.Errorf("Expected path length 4, got %d", len(path.Path))
	}
}

func TestAugmentingPathBFS(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(1, 2, 5)
	fn.AddEdge(2, 3, 15)

	path := fn.FindAugmentingPathBFS(0, 3)
	if path == nil {
		t.Error("Expected to find augmenting path")
	}

	if path.Bottleneck != 5 {
		t.Errorf("Expected bottleneck 5, got %.1f", path.Bottleneck)
	}

	if len(path.Path) != 4 {
		t.Errorf("Expected path length 4, got %d", len(path.Path))
	}
}

func TestNoAugmentingPath(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(2, 3, 5)

	path := fn.FindAugmentingPathDFS(0, 3)
	if path != nil {
		t.Error("Expected no augmenting path for disconnected graph")
	}

	path = fn.FindAugmentingPathBFS(0, 3)
	if path != nil {
		t.Error("Expected no augmenting path for disconnected graph")
	}
}

func TestFlowConservation(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(0, 2, 10)
	fn.AddEdge(1, 3, 10)
	fn.AddEdge(2, 3, 10)

	result, err := fn.FordFulkersonDFS(0, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	flowEdges := result.GetFlowEdges()
	sourceOutFlow := 0.0
	sinkInFlow := 0.0

	for _, edge := range flowEdges {
		if edge.From == 0 {
			sourceOutFlow += edge.Flow
		}
		if edge.To == 3 {
			sinkInFlow += edge.Flow
		}
	}

	if sourceOutFlow != sinkInFlow {
		t.Errorf("Flow conservation violated: source out=%.1f, sink in=%.1f",
			sourceOutFlow, sinkInFlow)
	}

	if sourceOutFlow != result.GetMaxFlow() {
		t.Errorf("Source outflow (%.1f) should equal max flow (%.1f)",
			sourceOutFlow, result.GetMaxFlow())
	}
}

func TestMinCutProperties(t *testing.T) {
	fn := NewFlowNetwork(4)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(0, 2, 10)
	fn.AddEdge(1, 3, 10)
	fn.AddEdge(2, 3, 10)

	result, err := fn.FordFulkersonDFS(0, 3)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	minCut := result.GetMinCut()
	if len(minCut) == 0 {
		t.Error("Expected non-empty min cut")
	}

	minCutCapacity := result.GetMinCutCapacity()
	if minCutCapacity != result.GetMaxFlow() {
		t.Errorf("Min cut capacity (%.1f) should equal max flow (%.1f)",
			minCutCapacity, result.GetMaxFlow())
	}
}

func TestLargeNetwork(t *testing.T) {
	n := 50
	fn := NewFlowNetwork(n)

	for i := range n - 1 {
		fn.AddEdge(i, i+1, float64(i+1))
	}

	for i := range n / 2 {
		fn.AddEdge(i, i+n/2, float64(n-i))
	}

	result, err := fn.FordFulkersonDFS(0, n-1)
	if err != nil {
		t.Errorf("Expected no error for large network, got %v", err)
	}

	if result.GetMaxFlow() <= 0 {
		t.Error("Expected positive max flow for large network")
	}

	if result.GetMaxFlow() != result.GetMinCutCapacity() {
		t.Error("Max flow should equal min cut capacity")
	}
}

func TestGetTotalCapacity(t *testing.T) {
	fn := NewFlowNetwork(3)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(1, 2, 5)
	fn.AddEdge(0, 2, 8)

	totalCapacity := fn.GetTotalCapacity()
	expectedTotal := 23.0
	if totalCapacity != expectedTotal {
		t.Errorf("Expected total capacity %.1f, got %.1f", expectedTotal, totalCapacity)
	}
}

func TestMaxFlowResultMethods(t *testing.T) {
	fn := NewFlowNetwork(3)
	fn.AddEdge(0, 1, 10)
	fn.AddEdge(1, 2, 5)

	result, err := fn.FordFulkersonDFS(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.GetMaxFlow() != 5 {
		t.Errorf("Expected max flow 5, got %.1f", result.GetMaxFlow())
	}

	if len(result.GetFlowEdges()) == 0 {
		t.Error("Expected non-empty flow edges")
	}

	if len(result.GetMinCut()) == 0 {
		t.Error("Expected non-empty min cut")
	}

	if result.GetMinCutCapacity() != result.GetMaxFlow() {
		t.Error("Min cut capacity should equal max flow")
	}
}

func BenchmarkFordFulkersonDFS(b *testing.B) {
	fn := NewFlowNetwork(100)
	for i := range 99 {
		fn.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		fn.FordFulkersonDFS(0, 99)
	}
}

func BenchmarkFordFulkersonBFS(b *testing.B) {
	fn := NewFlowNetwork(100)
	for i := range 99 {
		fn.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		fn.FordFulkersonBFS(0, 99)
	}
}

func BenchmarkAddEdge(b *testing.B) {
	fn := NewFlowNetwork(10000)

	for i := range b.N {
		from := i % 9999
		to := (i + 1) % 10000
		if from != to {
			fn.AddEdge(from, to, float64(i))
		}
	}
}

func BenchmarkFindAugmentingPathDFS(b *testing.B) {
	fn := NewFlowNetwork(50)
	for i := range 49 {
		fn.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		fn.FindAugmentingPathDFS(0, 49)
	}
}

func BenchmarkFindAugmentingPathBFS(b *testing.B) {
	fn := NewFlowNetwork(50)
	for i := range 49 {
		fn.AddEdge(i, i+1, float64(i+1))
	}

	for b.Loop() {
		fn.FindAugmentingPathBFS(0, 49)
	}
}
