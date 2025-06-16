package topological_sort

import (
	"errors"
	"fmt"
	"slices"
)

type Graph struct {
	vertices    int
	adjList     map[int][]int
	adjMatrix   [][]bool
	useMatrix   bool
	vertexNames map[int]string
}

type TopologicalSorter struct {
	graph *Graph
}

func NewGraph(vertices int, useMatrix bool) *Graph {
	g := &Graph{
		vertices:    vertices,
		adjList:     make(map[int][]int),
		useMatrix:   useMatrix,
		vertexNames: make(map[int]string),
	}

	if useMatrix {
		g.adjMatrix = make([][]bool, vertices)
		for i := range vertices {
			g.adjMatrix[i] = make([]bool, vertices)
		}
	}

	for i := range vertices {
		g.adjList[i] = []int{}
		g.vertexNames[i] = fmt.Sprintf("V%d", i)
	}

	return g
}

func NewTopologicalSorter(graph *Graph) *TopologicalSorter {
	return &TopologicalSorter{graph: graph}
}

func (g *Graph) AddEdge(from, to int) error {
	if from < 0 || from >= g.vertices || to < 0 || to >= g.vertices {
		return errors.New("vertex index out of range")
	}

	if g.useMatrix {
		g.adjMatrix[from][to] = true
	} else {
		if slices.Contains(g.adjList[from], to) {
			return nil
		}
		g.adjList[from] = append(g.adjList[from], to)
	}

	return nil
}

func (g *Graph) RemoveEdge(from, to int) error {
	if from < 0 || from >= g.vertices || to < 0 || to >= g.vertices {
		return errors.New("vertex index out of range")
	}

	if g.useMatrix {
		g.adjMatrix[from][to] = false
	} else {
		neighbors := g.adjList[from]
		for i, neighbor := range neighbors {
			if neighbor == to {
				g.adjList[from] = slices.Delete(neighbors, i, i+1)
				break
			}
		}
	}

	return nil
}

func (g *Graph) HasEdge(from, to int) bool {
	if from < 0 || from >= g.vertices || to < 0 || to >= g.vertices {
		return false
	}

	if g.useMatrix {
		return g.adjMatrix[from][to]
	}

	return slices.Contains(g.adjList[from], to)
}

func (g *Graph) GetNeighbors(vertex int) []int {
	if vertex < 0 || vertex >= g.vertices {
		return []int{}
	}

	if g.useMatrix {
		neighbors := []int{}
		for i := range g.vertices {
			if g.adjMatrix[vertex][i] {
				neighbors = append(neighbors, i)
			}
		}
		return neighbors
	}

	result := make([]int, len(g.adjList[vertex]))
	copy(result, g.adjList[vertex])
	return result
}

func (g *Graph) GetInDegree(vertex int) int {
	if vertex < 0 || vertex >= g.vertices {
		return 0
	}

	inDegree := 0
	if g.useMatrix {
		for i := range g.vertices {
			if g.adjMatrix[i][vertex] {
				inDegree++
			}
		}
	} else {
		for i := range g.vertices {
			if slices.Contains(g.adjList[i], vertex) {
				inDegree++
			}
		}
	}

	return inDegree
}

func (g *Graph) GetOutDegree(vertex int) int {
	if vertex < 0 || vertex >= g.vertices {
		return 0
	}

	if g.useMatrix {
		outDegree := 0
		for i := range g.vertices {
			if g.adjMatrix[vertex][i] {
				outDegree++
			}
		}
		return outDegree
	}

	return len(g.adjList[vertex])
}

func (g *Graph) GetVertexCount() int {
	return g.vertices
}

func (g *Graph) GetEdgeCount() int {
	edgeCount := 0
	if g.useMatrix {
		for i := range g.vertices {
			for j := range g.vertices {
				if g.adjMatrix[i][j] {
					edgeCount++
				}
			}
		}
	} else {
		for i := range g.vertices {
			edgeCount += len(g.adjList[i])
		}
	}
	return edgeCount
}

func (g *Graph) SetVertexName(vertex int, name string) error {
	if vertex < 0 || vertex >= g.vertices {
		return errors.New("vertex index out of range")
	}
	g.vertexNames[vertex] = name
	return nil
}

func (g *Graph) GetVertexName(vertex int) string {
	if vertex < 0 || vertex >= g.vertices {
		return ""
	}
	return g.vertexNames[vertex]
}

func (g *Graph) IsDAG() bool {
	sorter := NewTopologicalSorter(g)
	_, err := sorter.KahnSort()
	return err == nil
}

func (ts *TopologicalSorter) KahnSort() ([]int, error) {
	graph := ts.graph
	inDegree := make([]int, graph.vertices)

	for i := range graph.vertices {
		inDegree[i] = graph.GetInDegree(i)
	}

	queue := []int{}
	for i := range graph.vertices {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := []int{}
	processedCount := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)
		processedCount++

		neighbors := graph.GetNeighbors(current)
		for _, neighbor := range neighbors {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if processedCount != graph.vertices {
		return nil, errors.New("graph contains a cycle")
	}

	return result, nil
}

func (ts *TopologicalSorter) DFSSort() ([]int, error) {
	graph := ts.graph
	visited := make([]bool, graph.vertices)
	recStack := make([]bool, graph.vertices)
	result := []int{}

	var dfs func(int) error
	dfs = func(vertex int) error {
		visited[vertex] = true
		recStack[vertex] = true

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if recStack[neighbor] {
				return errors.New("graph contains a cycle")
			}
			if !visited[neighbor] {
				if err := dfs(neighbor); err != nil {
					return err
				}
			}
		}

		recStack[vertex] = false
		result = append([]int{vertex}, result...)
		return nil
	}

	for i := range graph.vertices {
		if !visited[i] {
			if err := dfs(i); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func (ts *TopologicalSorter) AllTopologicalSorts() ([][]int, error) {
	graph := ts.graph
	inDegree := make([]int, graph.vertices)

	for i := range graph.vertices {
		inDegree[i] = graph.GetInDegree(i)
	}

	var allSorts [][]int
	var currentSort []int
	visited := make([]bool, graph.vertices)

	var backtrack func()
	backtrack = func() {
		available := []int{}
		for i := range graph.vertices {
			if !visited[i] && inDegree[i] == 0 {
				available = append(available, i)
			}
		}

		if len(available) == 0 {
			if len(currentSort) == graph.vertices {
				sortCopy := make([]int, len(currentSort))
				copy(sortCopy, currentSort)
				allSorts = append(allSorts, sortCopy)
			}
			return
		}

		for _, vertex := range available {
			visited[vertex] = true
			currentSort = append(currentSort, vertex)

			neighbors := graph.GetNeighbors(vertex)
			for _, neighbor := range neighbors {
				inDegree[neighbor]--
			}

			backtrack()

			visited[vertex] = false
			currentSort = currentSort[:len(currentSort)-1]
			for _, neighbor := range neighbors {
				inDegree[neighbor]++
			}
		}
	}

	backtrack()

	if len(allSorts) == 0 {
		return nil, errors.New("graph contains a cycle")
	}

	return allSorts, nil
}

func (ts *TopologicalSorter) HasCycle() bool {
	graph := ts.graph
	visited := make([]bool, graph.vertices)
	recStack := make([]bool, graph.vertices)

	var dfs func(int) bool
	dfs = func(vertex int) bool {
		visited[vertex] = true
		recStack[vertex] = true

		neighbors := graph.GetNeighbors(vertex)
		for _, neighbor := range neighbors {
			if recStack[neighbor] {
				return true
			}
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}

		recStack[vertex] = false
		return false
	}

	for i := range graph.vertices {
		if !visited[i] && dfs(i) {
			return true
		}
	}

	return false
}

func (ts *TopologicalSorter) FindLongestPath() ([]int, int, error) {
	topOrder, err := ts.KahnSort()
	if err != nil {
		return nil, 0, err
	}

	graph := ts.graph
	dist := make([]int, graph.vertices)
	parent := make([]int, graph.vertices)

	for i := range graph.vertices {
		dist[i] = -1
		parent[i] = -1
	}

	for _, u := range topOrder {
		if dist[u] == -1 {
			dist[u] = 0
		}

		neighbors := graph.GetNeighbors(u)
		for _, v := range neighbors {
			if dist[v] < dist[u]+1 {
				dist[v] = dist[u] + 1
				parent[v] = u
			}
		}
	}

	maxDist := 0
	endVertex := 0
	for i := range graph.vertices {
		if dist[i] > maxDist {
			maxDist = dist[i]
			endVertex = i
		}
	}

	path := []int{}
	current := endVertex
	for current != -1 {
		path = append([]int{current}, path...)
		current = parent[current]
	}

	return path, maxDist, nil
}

func (g *Graph) PrintGraph() {
	fmt.Printf("Graph with %d vertices:\n", g.vertices)
	for i := range g.vertices {
		neighbors := g.GetNeighbors(i)
		fmt.Printf("%s (%d) -> ", g.GetVertexName(i), i)
		if len(neighbors) == 0 {
			fmt.Println("[]")
		} else {
			fmt.Print("[")
			for j, neighbor := range neighbors {
				if j > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%s (%d)", g.GetVertexName(neighbor), neighbor)
			}
			fmt.Println("]")
		}
	}
}

func Run() any {
	g := NewGraph(6, false)

	g.SetVertexName(0, "A")
	g.SetVertexName(1, "B")
	g.SetVertexName(2, "C")
	g.SetVertexName(3, "D")
	g.SetVertexName(4, "E")
	g.SetVertexName(5, "F")

	edges := [][2]int{
		{0, 1}, {0, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {4, 5},
	}

	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	sorter := NewTopologicalSorter(g)

	result := make(map[string]any)
	result["vertexCount"] = g.GetVertexCount()
	result["edgeCount"] = g.GetEdgeCount()
	result["isDAG"] = g.IsDAG()
	result["hasCycle"] = sorter.HasCycle()

	kahnSort, kahnErr := sorter.KahnSort()
	if kahnErr == nil {
		kahnNames := make([]string, len(kahnSort))
		for i, v := range kahnSort {
			kahnNames[i] = g.GetVertexName(v)
		}
		result["kahnSort"] = kahnNames
		result["kahnSortIndices"] = kahnSort
	} else {
		result["kahnSortError"] = kahnErr.Error()
	}

	dfsSort, dfsErr := sorter.DFSSort()
	if dfsErr == nil {
		dfsNames := make([]string, len(dfsSort))
		for i, v := range dfsSort {
			dfsNames[i] = g.GetVertexName(v)
		}
		result["dfsSort"] = dfsNames
		result["dfsSortIndices"] = dfsSort
	} else {
		result["dfsSortError"] = dfsErr.Error()
	}

	longestPath, pathLength, pathErr := sorter.FindLongestPath()
	if pathErr == nil {
		pathNames := make([]string, len(longestPath))
		for i, v := range longestPath {
			pathNames[i] = g.GetVertexName(v)
		}
		result["longestPath"] = pathNames
		result["longestPathIndices"] = longestPath
		result["longestPathLength"] = pathLength
	}

	inDegrees := make(map[string]int)
	outDegrees := make(map[string]int)
	for i := range g.GetVertexCount() {
		name := g.GetVertexName(i)
		inDegrees[name] = g.GetInDegree(i)
		outDegrees[name] = g.GetOutDegree(i)
	}
	result["inDegrees"] = inDegrees
	result["outDegrees"] = outDegrees

	g2 := NewGraph(4, true)
	g2.SetVertexName(0, "X")
	g2.SetVertexName(1, "Y")
	g2.SetVertexName(2, "Z")
	g2.SetVertexName(3, "W")

	cyclicEdges := [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}}
	for _, edge := range cyclicEdges {
		g2.AddEdge(edge[0], edge[1])
	}

	sorter2 := NewTopologicalSorter(g2)
	result["cyclicGraphHasCycle"] = sorter2.HasCycle()
	result["cyclicGraphIsDAG"] = g2.IsDAG()

	_, cyclicErr := sorter2.KahnSort()
	result["cyclicSortError"] = cyclicErr != nil
	if cyclicErr != nil {
		result["cyclicErrorMessage"] = cyclicErr.Error()
	}

	return result
}
