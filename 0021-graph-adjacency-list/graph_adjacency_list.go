package graph_adjacency_list

import "slices"

import "fmt"

type Graph struct {
	adjList  map[int][]int
	vertices int
	directed bool
}

func NewGraph(vertices int, directed bool) *Graph {
	return &Graph{
		adjList:  make(map[int][]int),
		vertices: vertices,
		directed: directed,
	}
}

func (g *Graph) AddEdge(from, to int) {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		g.adjList[from] = append(g.adjList[from], to)
		if !g.directed && from != to {
			g.adjList[to] = append(g.adjList[to], from)
		}
	}
}

func (g *Graph) RemoveEdge(from, to int) {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		g.adjList[from] = g.removeFromSlice(g.adjList[from], to)
		if !g.directed && from != to {
			g.adjList[to] = g.removeFromSlice(g.adjList[to], from)
		}
	}
}

func (g *Graph) removeFromSlice(slice []int, value int) []int {
	for i, v := range slice {
		if v == value {
			return slices.Delete(slice, i, i+1)
		}
	}
	return slice
}

func (g *Graph) HasEdge(from, to int) bool {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		neighbors := g.adjList[from]
		if slices.Contains(neighbors, to) {
			return true
		}
	}
	return false
}

func (g *Graph) GetNeighbors(vertex int) []int {
	if vertex >= 0 && vertex < g.vertices {
		return g.adjList[vertex]
	}
	return []int{}
}

func (g *Graph) BFS(start int) []int {
	if start < 0 || start >= g.vertices {
		return []int{}
	}

	visited := make([]bool, g.vertices)
	queue := []int{start}
	result := []int{}

	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)

		for _, neighbor := range g.GetNeighbors(vertex) {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func (g *Graph) DFS(start int) []int {
	if start < 0 || start >= g.vertices {
		return []int{}
	}

	visited := make([]bool, g.vertices)
	result := []int{}
	g.dfsRecursive(start, visited, &result)
	return result
}

func (g *Graph) dfsRecursive(vertex int, visited []bool, result *[]int) {
	visited[vertex] = true
	*result = append(*result, vertex)

	for _, neighbor := range g.GetNeighbors(vertex) {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited, result)
		}
	}
}

func (g *Graph) DFSIterative(start int) []int {
	if start < 0 || start >= g.vertices {
		return []int{}
	}

	visited := make([]bool, g.vertices)
	stack := []int{start}
	result := []int{}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)

			neighbors := g.GetNeighbors(vertex)
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}

	return result
}

func (g *Graph) IsConnected() bool {
	if g.vertices == 0 {
		return true
	}

	visited := g.BFS(0)
	return len(visited) == g.vertices
}

func (g *Graph) HasCycle() bool {
	visited := make([]bool, g.vertices)
	recStack := make([]bool, g.vertices)

	for i := range g.vertices {
		if !visited[i] {
			if g.hasCycleDFS(i, visited, recStack) {
				return true
			}
		}
	}
	return false
}

func (g *Graph) hasCycleDFS(vertex int, visited, recStack []bool) bool {
	visited[vertex] = true
	recStack[vertex] = true

	for _, neighbor := range g.GetNeighbors(vertex) {
		if !visited[neighbor] {
			if g.hasCycleDFS(neighbor, visited, recStack) {
				return true
			}
		} else if recStack[neighbor] {
			return true
		}
	}

	recStack[vertex] = false
	return false
}

func (g *Graph) GetVertexCount() int {
	return g.vertices
}

func (g *Graph) GetEdgeCount() int {
	count := 0
	for _, neighbors := range g.adjList {
		count += len(neighbors)
	}
	if !g.directed {
		count /= 2
	}
	return count
}

func (g *Graph) Print() {
	fmt.Printf("Graph (%d vertices, %s):\n", g.vertices, map[bool]string{true: "directed", false: "undirected"}[g.directed])
	for i := range g.vertices {
		fmt.Printf("%d: %v\n", i, g.adjList[i])
	}
}

func (g *Graph) TopologicalSort() []int {
	if !g.directed {
		return []int{}
	}

	visited := make([]bool, g.vertices)
	stack := []int{}

	for i := range g.vertices {
		if !visited[i] {
			g.topologicalSortDFS(i, visited, &stack)
		}
	}

	result := make([]int, len(stack))
	for i := range stack {
		result[i] = stack[len(stack)-1-i]
	}
	return result
}

func (g *Graph) topologicalSortDFS(vertex int, visited []bool, stack *[]int) {
	visited[vertex] = true

	for _, neighbor := range g.GetNeighbors(vertex) {
		if !visited[neighbor] {
			g.topologicalSortDFS(neighbor, visited, stack)
		}
	}

	*stack = append(*stack, vertex)
}

func Run() any {
	fmt.Println("=== Adjacency List Graph Implementation ===")

	fmt.Println("\n--- Undirected Graph ---")
	undirected := NewGraph(5, false)
	undirected.AddEdge(0, 1)
	undirected.AddEdge(0, 2)
	undirected.AddEdge(1, 3)
	undirected.AddEdge(2, 4)
	undirected.AddEdge(3, 4)

	undirected.Print()
	fmt.Printf("BFS from 0: %v\n", undirected.BFS(0))
	fmt.Printf("DFS from 0: %v\n", undirected.DFS(0))
	fmt.Printf("DFS Iterative from 0: %v\n", undirected.DFSIterative(0))
	fmt.Printf("Is Connected: %v\n", undirected.IsConnected())
	fmt.Printf("Edge Count: %d\n", undirected.GetEdgeCount())

	fmt.Println("\n--- Directed Graph ---")
	directed := NewGraph(4, true)
	directed.AddEdge(0, 1)
	directed.AddEdge(0, 2)
	directed.AddEdge(1, 2)
	directed.AddEdge(2, 3)

	directed.Print()
	fmt.Printf("BFS from 0: %v\n", directed.BFS(0))
	fmt.Printf("DFS from 0: %v\n", directed.DFS(0))
	fmt.Printf("DFS Iterative from 0: %v\n", directed.DFSIterative(0))
	fmt.Printf("Has Cycle: %v\n", directed.HasCycle())
	fmt.Printf("Topological Sort: %v\n", directed.TopologicalSort())

	fmt.Println("\n--- DAG Example ---")
	dag := NewGraph(6, true)
	dag.AddEdge(5, 2)
	dag.AddEdge(5, 0)
	dag.AddEdge(4, 0)
	dag.AddEdge(4, 1)
	dag.AddEdge(2, 3)
	dag.AddEdge(3, 1)

	dag.Print()
	fmt.Printf("Topological Sort: %v\n", dag.TopologicalSort())
	fmt.Printf("Has Cycle: %v\n", dag.HasCycle())

	return "Graph adjacency list implementation complete"
}
