package dijkstra_algorithm

import (
	"math"
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	if g.Vertices != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.Vertices)
	}
	if len(g.AdjList) != 5 {
		t.Errorf("Expected adjacency list of length 5, got %d", len(g.AdjList))
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 3)

	if len(g.AdjList[0]) != 1 {
		t.Errorf("Expected 1 edge from vertex 0, got %d", len(g.AdjList[0]))
	}

	edge := g.AdjList[0][0]
	if edge.To != 1 || edge.Weight != 5 {
		t.Errorf("Expected edge to vertex 1 with weight 5, got to %d with weight %d", edge.To, edge.Weight)
	}
}

func TestAddBidirectionalEdge(t *testing.T) {
	g := NewGraph(2)
	g.AddBidirectionalEdge(0, 1, 10)

	if len(g.AdjList[0]) != 1 || len(g.AdjList[1]) != 1 {
		t.Error("Bidirectional edge not added correctly")
	}

	if g.AdjList[0][0].To != 1 || g.AdjList[0][0].Weight != 10 {
		t.Error("Forward edge incorrect")
	}

	if g.AdjList[1][0].To != 0 || g.AdjList[1][0].Weight != 10 {
		t.Error("Backward edge incorrect")
	}
}

func TestDijkstraSimple(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 1)

	result := g.Dijkstra(0)

	expectedDistances := []int{0, 1, 3, 4}
	if !reflect.DeepEqual(result.Distances, expectedDistances) {
		t.Errorf("Expected distances %v, got %v", expectedDistances, result.Distances)
	}
}

func TestDijkstraDisconnectedGraph(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(2, 3, 1)

	result := g.Dijkstra(0)

	if result.Distances[0] != 0 {
		t.Errorf("Expected distance to source to be 0, got %d", result.Distances[0])
	}
	if result.Distances[1] != 1 {
		t.Errorf("Expected distance to vertex 1 to be 1, got %d", result.Distances[1])
	}
	if result.Distances[2] != math.MaxInt32 {
		t.Errorf("Expected distance to unreachable vertex 2 to be MaxInt32, got %d", result.Distances[2])
	}
	if result.Distances[3] != math.MaxInt32 {
		t.Errorf("Expected distance to unreachable vertex 3 to be MaxInt32, got %d", result.Distances[3])
	}
}

func TestDijkstraSingleVertex(t *testing.T) {
	g := NewGraph(1)
	result := g.Dijkstra(0)

	if result.Distances[0] != 0 {
		t.Errorf("Expected distance to source to be 0, got %d", result.Distances[0])
	}
}

func TestDijkstraInvalidSource(t *testing.T) {
	g := NewGraph(3)

	result := g.Dijkstra(-1)
	if result != nil {
		t.Error("Expected nil result for invalid source")
	}

	result = g.Dijkstra(3)
	if result != nil {
		t.Error("Expected nil result for out of bounds source")
	}
}

func TestGetPath(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 4, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 2)
	g.AddEdge(2, 3, 4)
	g.AddEdge(3, 2, 6)
	g.AddEdge(3, 0, 7)
	g.AddEdge(4, 1, 3)
	g.AddEdge(4, 2, 9)
	g.AddEdge(4, 3, 2)

	result := g.Dijkstra(0)

	pathTo3 := result.GetPath(3)
	expectedPath := []int{0, 4, 3}
	if !reflect.DeepEqual(pathTo3, expectedPath) {
		t.Errorf("Expected path to vertex 3: %v, got %v", expectedPath, pathTo3)
	}

	pathTo2 := result.GetPath(2)
	expectedPath2 := []int{0, 4, 1, 2}
	if !reflect.DeepEqual(pathTo2, expectedPath2) {
		t.Errorf("Expected path to vertex 2: %v, got %v", expectedPath2, pathTo2)
	}
}

func TestGetPathUnreachable(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1)

	result := g.Dijkstra(0)
	path := result.GetPath(2)

	if path != nil {
		t.Errorf("Expected nil path to unreachable vertex, got %v", path)
	}
}

func TestGetDistance(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 3)

	result := g.Dijkstra(0)

	if result.GetDistance(0) != 0 {
		t.Errorf("Expected distance to source to be 0, got %d", result.GetDistance(0))
	}
	if result.GetDistance(1) != 5 {
		t.Errorf("Expected distance to vertex 1 to be 5, got %d", result.GetDistance(1))
	}
	if result.GetDistance(2) != 8 {
		t.Errorf("Expected distance to vertex 2 to be 8, got %d", result.GetDistance(2))
	}
}

func TestHasPath(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)

	result := g.Dijkstra(0)

	if !result.HasPath(0) {
		t.Error("Expected path to source vertex")
	}
	if !result.HasPath(1) {
		t.Error("Expected path to connected vertex 1")
	}
	if !result.HasPath(2) {
		t.Error("Expected path to connected vertex 2")
	}
	if result.HasPath(3) {
		t.Error("Expected no path to disconnected vertex 3")
	}
}

func TestDijkstraComplexGraph(t *testing.T) {
	g := NewGraph(6)
	g.AddBidirectionalEdge(0, 1, 4)
	g.AddBidirectionalEdge(0, 2, 3)
	g.AddBidirectionalEdge(1, 2, 1)
	g.AddBidirectionalEdge(1, 3, 2)
	g.AddBidirectionalEdge(2, 3, 4)
	g.AddBidirectionalEdge(3, 4, 2)
	g.AddBidirectionalEdge(4, 5, 6)

	result := g.Dijkstra(0)

	expectedDistances := []int{0, 4, 3, 6, 8, 14}
	if !reflect.DeepEqual(result.Distances, expectedDistances) {
		t.Errorf("Expected distances %v, got %v", expectedDistances, result.Distances)
	}

	pathTo5 := result.GetPath(5)
	expectedPath := []int{0, 1, 3, 4, 5}
	if !reflect.DeepEqual(pathTo5, expectedPath) {
		t.Errorf("Expected path to vertex 5: %v, got %v", expectedPath, pathTo5)
	}
}

func TestEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		edges    [][]int
		source   int
		target   int
		expected int
	}{
		{"Self loop", 1, [][]int{}, 0, 0, 0},
		{"Two vertices connected", 2, [][]int{{0, 1, 5}}, 0, 1, 5},
		{"Two vertices disconnected", 2, [][]int{}, 0, 1, math.MaxInt32},
		{"Linear chain", 4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := NewGraph(test.vertices)
			for _, edge := range test.edges {
				g.AddEdge(edge[0], edge[1], edge[2])
			}

			result := g.Dijkstra(test.source)
			distance := result.GetDistance(test.target)

			if distance != test.expected {
				t.Errorf("Expected distance %d, got %d", test.expected, distance)
			}
		})
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
	}

	if resultMap["source"] != 0 {
		t.Error("Expected source to be 0")
	}

	distances, ok := resultMap["distances"].([]int)
	if !ok || len(distances) != 6 {
		t.Error("Expected distances array of length 6")
	}

	if distances[0] != 0 {
		t.Error("Expected distance to source to be 0")
	}
}

func BenchmarkDijkstraSmallGraph(b *testing.B) {
	g := NewGraph(10)
	for i := range 9 {
		g.AddBidirectionalEdge(i, i+1, i+1)
	}

	for b.Loop() {
		g.Dijkstra(0)
	}
}

func BenchmarkDijkstraMediumGraph(b *testing.B) {
	g := NewGraph(100)
	for i := range 99 {
		g.AddBidirectionalEdge(i, i+1, 1)
		if i < 98 {
			g.AddBidirectionalEdge(i, i+2, 3)
		}
	}

	for b.Loop() {
		g.Dijkstra(0)
	}
}

func BenchmarkDijkstraLargeGraph(b *testing.B) {
	g := NewGraph(1000)
	for i := range 999 {
		g.AddBidirectionalEdge(i, i+1, 1)
		if i%10 == 0 && i < 990 {
			g.AddBidirectionalEdge(i, i+10, 2)
		}
	}

	for b.Loop() {
		g.Dijkstra(0)
	}
}

func BenchmarkDijkstraDenseGraph(b *testing.B) {
	vertices := 50
	g := NewGraph(vertices)

	for i := range vertices {
		for j := i + 1; j < vertices; j++ {
			g.AddBidirectionalEdge(i, j, i+j+1)
		}
	}

	for b.Loop() {
		g.Dijkstra(0)
	}
}

func BenchmarkGetPath(b *testing.B) {
	g := NewGraph(100)
	for i := range 99 {
		g.AddBidirectionalEdge(i, i+1, 1)
	}

	result := g.Dijkstra(0)

	for b.Loop() {
		result.GetPath(99)
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
