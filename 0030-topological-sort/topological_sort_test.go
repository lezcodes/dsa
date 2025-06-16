package topological_sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}

func TestNewGraph(t *testing.T) {
	g := NewGraph(5, false)
	if g.GetVertexCount() != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.GetVertexCount())
	}
	if g.GetEdgeCount() != 0 {
		t.Errorf("Expected 0 edges, got %d", g.GetEdgeCount())
	}
}

func TestNewGraphWithMatrix(t *testing.T) {
	g := NewGraph(3, true)
	if g.GetVertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", g.GetVertexCount())
	}
	if !g.useMatrix {
		t.Error("Expected graph to use adjacency matrix")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(3, false)

	err := g.AddEdge(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to exist")
	}

	if g.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", g.GetEdgeCount())
	}
}

func TestAddEdgeMatrix(t *testing.T) {
	g := NewGraph(3, true)

	err := g.AddEdge(0, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !g.HasEdge(0, 2) {
		t.Error("Expected edge (0,2) to exist in matrix")
	}
}

func TestAddEdgeInvalidVertex(t *testing.T) {
	g := NewGraph(3, false)

	err := g.AddEdge(0, 5)
	if err == nil {
		t.Error("Expected error for invalid vertex")
	}

	err = g.AddEdge(-1, 1)
	if err == nil {
		t.Error("Expected error for negative vertex")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph(3, false)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)

	err := g.RemoveEdge(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if g.HasEdge(0, 1) {
		t.Error("Expected edge (0,1) to be removed")
	}

	if g.GetEdgeCount() != 1 {
		t.Errorf("Expected 1 edge after removal, got %d", g.GetEdgeCount())
	}
}

func TestGetNeighbors(t *testing.T) {
	g := NewGraph(4, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)

	neighbors := g.GetNeighbors(0)
	expected := []int{1, 2, 3}

	sort.Ints(neighbors)
	sort.Ints(expected)

	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Expected neighbors %v, got %v", expected, neighbors)
	}
}

func TestGetInOutDegree(t *testing.T) {
	g := NewGraph(4, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	if g.GetOutDegree(0) != 2 {
		t.Errorf("Expected out-degree 2 for vertex 0, got %d", g.GetOutDegree(0))
	}

	if g.GetInDegree(2) != 2 {
		t.Errorf("Expected in-degree 2 for vertex 2, got %d", g.GetInDegree(2))
	}

	if g.GetInDegree(0) != 0 {
		t.Errorf("Expected in-degree 0 for vertex 0, got %d", g.GetInDegree(0))
	}

	if g.GetOutDegree(3) != 0 {
		t.Errorf("Expected out-degree 0 for vertex 3, got %d", g.GetOutDegree(3))
	}
}

func TestVertexNames(t *testing.T) {
	g := NewGraph(3, false)

	err := g.SetVertexName(0, "Start")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	name := g.GetVertexName(0)
	if name != "Start" {
		t.Errorf("Expected name 'Start', got '%s'", name)
	}

	err = g.SetVertexName(5, "Invalid")
	if err == nil {
		t.Error("Expected error for invalid vertex index")
	}
}

func TestKahnSortSimple(t *testing.T) {
	g := NewGraph(4, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	sorter := NewTopologicalSorter(g)
	result, err := sorter.KahnSort()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 4 vertices in result, got %d", len(result))
	}

	if !isValidTopologicalOrder(g, result) {
		t.Errorf("Result is not a valid topological order: %v", result)
	}
}

func TestDFSSortSimple(t *testing.T) {
	g := NewGraph(4, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	sorter := NewTopologicalSorter(g)
	result, err := sorter.DFSSort()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 4 vertices in result, got %d", len(result))
	}

	if !isValidTopologicalOrder(g, result) {
		t.Errorf("Result is not a valid topological order: %v", result)
	}
}

func TestCyclicGraph(t *testing.T) {
	g := NewGraph(3, false)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)

	sorter := NewTopologicalSorter(g)

	if !sorter.HasCycle() {
		t.Error("Expected cycle to be detected")
	}

	if g.IsDAG() {
		t.Error("Expected graph to not be a DAG")
	}

	_, err := sorter.KahnSort()
	if err == nil {
		t.Error("Expected error for cyclic graph in Kahn's algorithm")
	}

	_, err = sorter.DFSSort()
	if err == nil {
		t.Error("Expected error for cyclic graph in DFS algorithm")
	}
}

func TestSelfLoop(t *testing.T) {
	g := NewGraph(2, false)
	g.AddEdge(0, 0)

	sorter := NewTopologicalSorter(g)

	if !sorter.HasCycle() {
		t.Error("Expected self-loop to be detected as cycle")
	}
}

func TestEmptyGraph(t *testing.T) {
	g := NewGraph(3, false)
	sorter := NewTopologicalSorter(g)

	result, err := sorter.KahnSort()
	if err != nil {
		t.Errorf("Expected no error for empty graph, got %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 vertices in result, got %d", len(result))
	}
}

func TestSingleVertex(t *testing.T) {
	g := NewGraph(1, false)
	sorter := NewTopologicalSorter(g)

	result, err := sorter.KahnSort()
	if err != nil {
		t.Errorf("Expected no error for single vertex, got %v", err)
	}

	if len(result) != 1 || result[0] != 0 {
		t.Errorf("Expected [0], got %v", result)
	}
}

func TestLinearGraph(t *testing.T) {
	g := NewGraph(5, false)
	for i := range 4 {
		g.AddEdge(i, i+1)
	}

	sorter := NewTopologicalSorter(g)
	result, err := sorter.KahnSort()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComplexDAG(t *testing.T) {
	g := NewGraph(6, false)
	edges := [][2]int{
		{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	sorter := NewTopologicalSorter(g)

	kahnResult, kahnErr := sorter.KahnSort()
	dfsResult, dfsErr := sorter.DFSSort()

	if kahnErr != nil {
		t.Errorf("Kahn's algorithm failed: %v", kahnErr)
	}

	if dfsErr != nil {
		t.Errorf("DFS algorithm failed: %v", dfsErr)
	}

	if !isValidTopologicalOrder(g, kahnResult) {
		t.Errorf("Kahn's result is invalid: %v", kahnResult)
	}

	if !isValidTopologicalOrder(g, dfsResult) {
		t.Errorf("DFS result is invalid: %v", dfsResult)
	}
}

func TestAllTopologicalSorts(t *testing.T) {
	g := NewGraph(3, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	sorter := NewTopologicalSorter(g)
	allSorts, err := sorter.AllTopologicalSorts()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(allSorts) != 2 {
		t.Errorf("Expected 2 possible sorts, got %d", len(allSorts))
	}

	for _, sort := range allSorts {
		if !isValidTopologicalOrder(g, sort) {
			t.Errorf("Invalid topological sort: %v", sort)
		}
	}
}

func TestFindLongestPath(t *testing.T) {
	g := NewGraph(4, false)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	sorter := NewTopologicalSorter(g)
	path, length, err := sorter.FindLongestPath()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if length != 3 {
		t.Errorf("Expected path length 3, got %d", length)
	}

	expectedPath := []int{0, 1, 2, 3}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
}

func TestMatrixVsListConsistency(t *testing.T) {
	edges := [][2]int{
		{0, 1}, {0, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {4, 5},
	}

	gList := NewGraph(6, false)
	gMatrix := NewGraph(6, true)

	for _, edge := range edges {
		gList.AddEdge(edge[0], edge[1])
		gMatrix.AddEdge(edge[0], edge[1])
	}

	sorterList := NewTopologicalSorter(gList)
	sorterMatrix := NewTopologicalSorter(gMatrix)

	listResult, listErr := sorterList.KahnSort()
	matrixResult, matrixErr := sorterMatrix.KahnSort()

	if listErr != nil || matrixErr != nil {
		t.Errorf("Errors: list=%v, matrix=%v", listErr, matrixErr)
	}

	if !isValidTopologicalOrder(gList, listResult) {
		t.Error("List-based result is invalid")
	}

	if !isValidTopologicalOrder(gMatrix, matrixResult) {
		t.Error("Matrix-based result is invalid")
	}
}

func TestLargeDAG(t *testing.T) {
	n := 100
	g := NewGraph(n, false)

	rand.Seed(time.Now().UnixNano())
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			if rand.Float32() < 0.1 {
				g.AddEdge(i, j)
			}
		}
	}

	sorter := NewTopologicalSorter(g)

	if sorter.HasCycle() {
		t.Error("Large DAG should not have cycles")
	}

	result, err := sorter.KahnSort()
	if err != nil {
		t.Errorf("Large DAG sort failed: %v", err)
	}

	if len(result) != n {
		t.Errorf("Expected %d vertices, got %d", n, len(result))
	}

	if !isValidTopologicalOrder(g, result) {
		t.Error("Large DAG result is invalid")
	}
}

func TestDisconnectedGraph(t *testing.T) {
	g := NewGraph(6, false)
	g.AddEdge(0, 1)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)

	sorter := NewTopologicalSorter(g)
	result, err := sorter.KahnSort()

	if err != nil {
		t.Errorf("Expected no error for disconnected graph, got %v", err)
	}

	if len(result) != 6 {
		t.Errorf("Expected 6 vertices, got %d", len(result))
	}

	if !isValidTopologicalOrder(g, result) {
		t.Error("Disconnected graph result is invalid")
	}
}

func TestEdgeCases(t *testing.T) {
	g := NewGraph(0, false)
	if g.GetVertexCount() != 0 {
		t.Error("Expected 0 vertices for empty graph")
	}

	g2 := NewGraph(1, false)
	sorter := NewTopologicalSorter(g2)
	result, err := sorter.KahnSort()

	if err != nil {
		t.Errorf("Single vertex should work, got error: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 vertex, got %d", len(result))
	}
}

func isValidTopologicalOrder(g *Graph, order []int) bool {
	if len(order) != g.GetVertexCount() {
		return false
	}

	position := make(map[int]int)
	for i, vertex := range order {
		position[vertex] = i
	}

	for u := range g.GetVertexCount() {
		neighbors := g.GetNeighbors(u)
		for _, v := range neighbors {
			if position[u] >= position[v] {
				return false
			}
		}
	}

	return true
}

func BenchmarkKahnSort(b *testing.B) {
	g := NewGraph(1000, false)
	for i := range 999 {
		g.AddEdge(i, i+1)
	}

	sorter := NewTopologicalSorter(g)

	for b.Loop() {
		sorter.KahnSort()
	}
}

func BenchmarkDFSSort(b *testing.B) {
	g := NewGraph(1000, false)
	for i := range 999 {
		g.AddEdge(i, i+1)
	}

	sorter := NewTopologicalSorter(g)

	for b.Loop() {
		sorter.DFSSort()
	}
}

func BenchmarkHasCycle(b *testing.B) {
	g := NewGraph(1000, false)
	rand.Seed(42)
	for i := range 999 {
		for j := i + 1; j < 1000; j++ {
			if rand.Float32() < 0.01 {
				g.AddEdge(i, j)
			}
		}
	}

	sorter := NewTopologicalSorter(g)

	for b.Loop() {
		sorter.HasCycle()
	}
}

func BenchmarkAddEdge(b *testing.B) {
	g := NewGraph(10000, false)
	b.ResetTimer()

	for i := range b.N {
		from := i % 9999
		to := (i + 1) % 10000
		if from < to {
			g.AddEdge(from, to)
		}
	}
}

func BenchmarkGetInDegree(b *testing.B) {
	g := NewGraph(1000, false)
	for i := range 999 {
		g.AddEdge(i, i+1)
	}

	b.ResetTimer()
	for i := range b.N {
		g.GetInDegree(i % 1000)
	}
}
