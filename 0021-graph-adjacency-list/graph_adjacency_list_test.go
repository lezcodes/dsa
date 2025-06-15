package graph_adjacency_list

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewGraph(t *testing.T) {
	tests := []struct {
		vertices int
		directed bool
	}{
		{5, true},
		{3, false},
		{0, true},
		{1, false},
	}

	for _, test := range tests {
		g := NewGraph(test.vertices, test.directed)
		if g.vertices != test.vertices {
			t.Errorf("Expected %d vertices, got %d", test.vertices, g.vertices)
		}
		if g.directed != test.directed {
			t.Errorf("Expected directed=%v, got %v", test.directed, g.directed)
		}
		if g.adjList == nil {
			t.Error("Expected non-nil adjacency list")
		}
	}
}

func TestAddEdge(t *testing.T) {
	t.Run("Directed Graph", func(t *testing.T) {
		g := NewGraph(3, true)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		if !g.HasEdge(0, 1) {
			t.Error("Expected edge 0->1")
		}
		if g.HasEdge(1, 0) {
			t.Error("Unexpected edge 1->0 in directed graph")
		}
		if !g.HasEdge(1, 2) {
			t.Error("Expected edge 1->2")
		}
	})

	t.Run("Undirected Graph", func(t *testing.T) {
		g := NewGraph(3, false)
		g.AddEdge(0, 1)

		if !g.HasEdge(0, 1) {
			t.Error("Expected edge 0->1")
		}
		if !g.HasEdge(1, 0) {
			t.Error("Expected edge 1->0 in undirected graph")
		}
	})

	t.Run("Self Loop", func(t *testing.T) {
		g := NewGraph(3, true)
		g.AddEdge(0, 0)

		if !g.HasEdge(0, 0) {
			t.Error("Expected self-loop edge 0->0")
		}
	})

	t.Run("Invalid Vertices", func(t *testing.T) {
		g := NewGraph(3, true)
		g.AddEdge(-1, 0)
		g.AddEdge(0, 3)
		g.AddEdge(3, 0)

		if g.HasEdge(-1, 0) || g.HasEdge(0, 3) || g.HasEdge(3, 0) {
			t.Error("Should not add edges with invalid vertices")
		}
	})
}

func TestRemoveEdge(t *testing.T) {
	t.Run("Directed Graph", func(t *testing.T) {
		g := NewGraph(3, true)
		g.AddEdge(0, 1)
		g.AddEdge(1, 0)
		g.RemoveEdge(0, 1)

		if g.HasEdge(0, 1) {
			t.Error("Edge 0->1 should be removed")
		}
		if !g.HasEdge(1, 0) {
			t.Error("Edge 1->0 should still exist")
		}
	})

	t.Run("Undirected Graph", func(t *testing.T) {
		g := NewGraph(3, false)
		g.AddEdge(0, 1)
		g.RemoveEdge(0, 1)

		if g.HasEdge(0, 1) || g.HasEdge(1, 0) {
			t.Error("Both directions should be removed in undirected graph")
		}
	})

	t.Run("Remove Non-existent Edge", func(t *testing.T) {
		g := NewGraph(3, true)
		g.RemoveEdge(0, 1)

		if g.HasEdge(0, 1) {
			t.Error("Should handle removal of non-existent edge gracefully")
		}
	})
}

func TestGetNeighbors(t *testing.T) {
	g := NewGraph(4, true)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)

	neighbors := g.GetNeighbors(0)
	sort.Ints(neighbors)
	expected := []int{1, 2}
	if !reflect.DeepEqual(neighbors, expected) {
		t.Errorf("Expected neighbors %v, got %v", expected, neighbors)
	}

	neighbors = g.GetNeighbors(3)
	if len(neighbors) != 0 {
		t.Errorf("Expected no neighbors for vertex 3, got %v", neighbors)
	}

	neighbors = g.GetNeighbors(-1)
	if len(neighbors) != 0 {
		t.Error("Should return empty slice for invalid vertex")
	}
}

func TestBFS(t *testing.T) {
	g := NewGraph(5, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	result := g.BFS(0)
	if len(result) != 5 {
		t.Errorf("Expected 5 vertices in BFS, got %d", len(result))
	}
	if result[0] != 0 {
		t.Errorf("Expected BFS to start with 0, got %d", result[0])
	}

	result = g.BFS(-1)
	if len(result) != 0 {
		t.Error("BFS with invalid start should return empty slice")
	}

	disconnected := NewGraph(4, false)
	disconnected.AddEdge(0, 1)
	disconnected.AddEdge(2, 3)
	result = disconnected.BFS(0)
	if len(result) != 2 {
		t.Errorf("Expected 2 vertices in disconnected BFS, got %d", len(result))
	}
}

func TestDFS(t *testing.T) {
	g := NewGraph(4, true)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)

	result := g.DFS(0)
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices in DFS, got %d", len(result))
	}
	if result[0] != 0 {
		t.Errorf("Expected DFS to start with 0, got %d", result[0])
	}

	result = g.DFS(-1)
	if len(result) != 0 {
		t.Error("DFS with invalid start should return empty slice")
	}
}

func TestDFSIterative(t *testing.T) {
	g := NewGraph(4, true)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)

	result := g.DFSIterative(0)
	if len(result) != 4 {
		t.Errorf("Expected 4 vertices in iterative DFS, got %d", len(result))
	}
	if result[0] != 0 {
		t.Errorf("Expected iterative DFS to start with 0, got %d", result[0])
	}

	result = g.DFSIterative(-1)
	if len(result) != 0 {
		t.Error("Iterative DFS with invalid start should return empty slice")
	}
}

func TestIsConnected(t *testing.T) {
	connected := NewGraph(3, false)
	connected.AddEdge(0, 1)
	connected.AddEdge(1, 2)

	if !connected.IsConnected() {
		t.Error("Expected connected graph to return true")
	}

	disconnected := NewGraph(4, false)
	disconnected.AddEdge(0, 1)
	disconnected.AddEdge(2, 3)

	if disconnected.IsConnected() {
		t.Error("Expected disconnected graph to return false")
	}

	empty := NewGraph(0, false)
	if !empty.IsConnected() {
		t.Error("Empty graph should be considered connected")
	}

	single := NewGraph(1, false)
	if !single.IsConnected() {
		t.Error("Single vertex graph should be connected")
	}
}

func TestHasCycle(t *testing.T) {
	noCycle := NewGraph(3, true)
	noCycle.AddEdge(0, 1)
	noCycle.AddEdge(1, 2)

	if noCycle.HasCycle() {
		t.Error("Expected no cycle in linear graph")
	}

	withCycle := NewGraph(3, true)
	withCycle.AddEdge(0, 1)
	withCycle.AddEdge(1, 2)
	withCycle.AddEdge(2, 0)

	if !withCycle.HasCycle() {
		t.Error("Expected cycle in circular graph")
	}

	selfLoop := NewGraph(2, true)
	selfLoop.AddEdge(0, 0)

	if !selfLoop.HasCycle() {
		t.Error("Expected cycle with self-loop")
	}
}

func TestGetVertexCount(t *testing.T) {
	g := NewGraph(5, true)
	if g.GetVertexCount() != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.GetVertexCount())
	}
}

func TestGetEdgeCount(t *testing.T) {
	t.Run("Directed Graph", func(t *testing.T) {
		g := NewGraph(3, true)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(0, 2)

		if g.GetEdgeCount() != 3 {
			t.Errorf("Expected 3 edges in directed graph, got %d", g.GetEdgeCount())
		}
	})

	t.Run("Undirected Graph", func(t *testing.T) {
		g := NewGraph(3, false)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)

		if g.GetEdgeCount() != 2 {
			t.Errorf("Expected 2 edges in undirected graph, got %d", g.GetEdgeCount())
		}
	})
}

func TestTopologicalSort(t *testing.T) {
	t.Run("DAG", func(t *testing.T) {
		g := NewGraph(6, true)
		g.AddEdge(5, 2)
		g.AddEdge(5, 0)
		g.AddEdge(4, 0)
		g.AddEdge(4, 1)
		g.AddEdge(2, 3)
		g.AddEdge(3, 1)

		result := g.TopologicalSort()
		if len(result) != 6 {
			t.Errorf("Expected 6 vertices in topological sort, got %d", len(result))
		}

		positions := make(map[int]int)
		for i, v := range result {
			positions[v] = i
		}

		if positions[5] >= positions[2] || positions[5] >= positions[0] {
			t.Error("Vertex 5 should come before 2 and 0")
		}
		if positions[4] >= positions[0] || positions[4] >= positions[1] {
			t.Error("Vertex 4 should come before 0 and 1")
		}
		if positions[2] >= positions[3] {
			t.Error("Vertex 2 should come before 3")
		}
		if positions[3] >= positions[1] {
			t.Error("Vertex 3 should come before 1")
		}
	})

	t.Run("Undirected Graph", func(t *testing.T) {
		g := NewGraph(3, false)
		g.AddEdge(0, 1)
		result := g.TopologicalSort()
		if len(result) != 0 {
			t.Error("Topological sort should return empty for undirected graph")
		}
	})
}

func TestComplexGraphOperations(t *testing.T) {
	g := NewGraph(6, false)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	bfs := g.BFS(0)
	dfs := g.DFS(0)
	dfsIter := g.DFSIterative(0)

	if len(bfs) != 6 || len(dfs) != 6 || len(dfsIter) != 6 {
		t.Error("All traversals should visit all 6 vertices")
	}

	if !g.IsConnected() {
		t.Error("Graph should be connected")
	}

	g.RemoveEdge(3, 5)
	g.RemoveEdge(4, 5)

	bfsAfterRemoval := g.BFS(0)
	if len(bfsAfterRemoval) != 5 {
		t.Errorf("Expected 5 vertices after edge removal, got %d", len(bfsAfterRemoval))
	}
}

func BenchmarkBFS(b *testing.B) {
	g := NewGraph(100, false)
	for i := range 99 {
		g.AddEdge(i, i+1)
	}

	for b.Loop() {
		g.BFS(0)
	}
}

func BenchmarkDFS(b *testing.B) {
	g := NewGraph(100, false)
	for i := range 99 {
		g.AddEdge(i, i+1)
	}

	for b.Loop() {
		g.DFS(0)
	}
}

func BenchmarkDFSIterative(b *testing.B) {
	g := NewGraph(100, false)
	for i := range 99 {
		g.AddEdge(i, i+1)
	}

	for b.Loop() {
		g.DFSIterative(0)
	}
}

func BenchmarkAddEdge(b *testing.B) {
	g := NewGraph(1000, false)

	for i := 0; b.Loop(); i++ {
		g.AddEdge(i%1000, (i+1)%1000)
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	g := NewGraph(100, true)
	for i := range 99 {
		g.AddEdge(i, i+1)
	}

	for b.Loop() {
		g.TopologicalSort()
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}
