package detect_cycles

import (
	"reflect"
	"sort"
	"testing"
)

func sortCycles(cycles [][]int) {
	for i := range cycles {
		sort.Ints(cycles[i])
	}
	sort.Slice(cycles, func(i, j int) bool {
		if len(cycles[i]) != len(cycles[j]) {
			return len(cycles[i]) < len(cycles[j])
		}
		for k := range cycles[i] {
			if cycles[i][k] != cycles[j][k] {
				return cycles[i][k] < cycles[j][k]
			}
		}
		return false
	})
}

func TestDetectCycles(t *testing.T) {
	testCases := []struct {
		name     string
		blockers [][]bool
		expected [][]int
	}{
		{
			name: "Simple 3-node cycle",
			blockers: [][]bool{
				{false, false, true},
				{true, false, false},
				{false, true, false},
			},
			expected: [][]int{{0, 1, 2}},
		},
		{
			name: "Two separate cycles",
			blockers: [][]bool{
				{false, false, true, false, false},
				{true, false, false, false, false},
				{false, true, false, false, false},
				{false, false, false, false, true},
				{false, false, false, true, false},
			},
			expected: [][]int{{0, 1, 2}, {3, 4}},
		},
		{
			name: "Self-loop",
			blockers: [][]bool{
				{true, false},
				{false, false},
			},
			expected: [][]int{{0}},
		},
		{
			name: "No cycles",
			blockers: [][]bool{
				{false, true, false},
				{false, false, true},
				{false, false, false},
			},
			expected: [][]int{},
		},
		{
			name: "Complex graph with one cycle",
			blockers: [][]bool{
				{false, true, false, false},
				{false, false, true, false},
				{false, false, false, true},
				{true, false, false, false},
			},
			expected: [][]int{{0, 1, 2, 3}},
		},
		{
			name:     "Empty graph",
			blockers: [][]bool{},
			expected: [][]int{},
		},
		{
			name: "Single node no cycle",
			blockers: [][]bool{
				{false},
			},
			expected: [][]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := detectCycles(tc.blockers)

			sortCycles(result)
			sortCycles(tc.expected)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestBuildGraphFromBlockers(t *testing.T) {
	blockers := [][]bool{
		{false, false, true},
		{true, false, false},
		{false, true, false},
	}

	graph := buildGraphFromBlockers(blockers)

	if graph.numNodes != 3 {
		t.Errorf("Expected 3 nodes, got %d", graph.numNodes)
	}

	expectedEdges := [][]int{
		{1},
		{2},
		{0},
	}

	for i, expected := range expectedEdges {
		if !reflect.DeepEqual(graph.adjacencyList[i], expected) {
			t.Errorf("Node %d: expected %v, got %v", i, expected, graph.adjacencyList[i])
		}
	}
}

func TestCycleNormalization(t *testing.T) {
	detector := &CycleDetector{}

	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{1, 2, 0}, []int{0, 1, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{5}, []int{5}},
	}

	for _, tc := range testCases {
		result := detector.normalizeCycle(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Input %v: expected %v, got %v", tc.input, tc.expected, result)
		}
	}
}

func TestExtractCycle(t *testing.T) {
	detector := &CycleDetector{}

	path := []int{0, 1, 2, 3}
	backEdgeTarget := 1

	result := detector.extractCycle(path, backEdgeTarget)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	results, ok := result.([][][]int)
	if !ok {
		t.Error("Expected result to be [][][]int")
		return
	}

	if len(results) == 0 {
		t.Error("Expected at least one result")
	}

	for i, cycles := range results {
		for j, cycle := range cycles {
			if len(cycle) == 0 {
				t.Errorf("Result %d, cycle %d: empty cycle", i, j)
			}
		}
	}
}

func TestLargeGraph(t *testing.T) {
	size := 100
	blockers := make([][]bool, size)
	for i := range size {
		blockers[i] = make([]bool, size)
	}

	for i := range size - 1 {
		blockers[i+1][i] = true
	}
	blockers[0][size-1] = true

	result := detectCycles(blockers)

	if len(result) != 1 {
		t.Errorf("Expected 1 cycle, got %d", len(result))
	}

	if len(result) > 0 && len(result[0]) != size {
		t.Errorf("Expected cycle of length %d, got %d", size, len(result[0]))
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}

func BenchmarkDetectCycles(b *testing.B) {
	blockers := [][]bool{
		{false, false, true, false, false},
		{true, false, false, false, false},
		{false, true, false, false, false},
		{false, false, false, false, true},
		{false, false, false, true, false},
	}

	b.ResetTimer()
	for b.Loop() {
		detectCycles(blockers)
	}
}

func BenchmarkLargeGraph(b *testing.B) {
	size := 1000
	blockers := make([][]bool, size)
	for i := range size {
		blockers[i] = make([]bool, size)
	}

	for i := range size - 1 {
		blockers[i+1][i] = true
	}
	blockers[0][size-1] = true

	b.ResetTimer()
	for b.Loop() {
		detectCycles(blockers)
	}
}
