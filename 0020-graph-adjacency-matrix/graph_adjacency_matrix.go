package graph_adjacency_matrix

import "fmt"

type Graph struct {
	matrix   [][]int
	vertices int
	directed bool
}

func NewGraph(vertices int, directed bool) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &Graph{
		matrix:   matrix,
		vertices: vertices,
		directed: directed,
	}
}

func (g *Graph) AddEdge(from, to int) {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		g.matrix[from][to] = 1
		if !g.directed {
			g.matrix[to][from] = 1
		}
	}
}

func (g *Graph) RemoveEdge(from, to int) {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		g.matrix[from][to] = 0
		if !g.directed {
			g.matrix[to][from] = 0
		}
	}
}

func (g *Graph) HasEdge(from, to int) bool {
	if from >= 0 && from < g.vertices && to >= 0 && to < g.vertices {
		return g.matrix[from][to] == 1
	}
	return false
}

func (g *Graph) GetNeighbors(vertex int) []int {
	neighbors := []int{}
	if vertex >= 0 && vertex < g.vertices {
		for i := range g.vertices {
			if g.matrix[vertex][i] == 1 {
				neighbors = append(neighbors, i)
			}
		}
	}
	return neighbors
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

func (g *Graph) Print() {
	fmt.Printf("Graph (%d vertices, %s):\n", g.vertices, map[bool]string{true: "directed", false: "undirected"}[g.directed])
	for i := range g.vertices {
		fmt.Printf("%d: ", i)
		for j := range g.vertices {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

func Run() any {
	fmt.Println("=== Adjacency Matrix Graph Implementation ===")

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

	return "Graph adjacency matrix implementation complete"
}
