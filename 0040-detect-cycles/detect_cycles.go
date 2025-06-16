package detect_cycles

const (
	WHITE = iota
	GRAY
	BLACK
)

type Graph struct {
	adjacencyList [][]int
	numNodes      int
}

type CycleDetector struct {
	graph  *Graph
	color  []int
	parent []int
	cycles [][]int
}

func NewGraph(numNodes int) *Graph {
	return &Graph{
		adjacencyList: make([][]int, numNodes),
		numNodes:      numNodes,
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
}

func buildGraphFromBlockers(blockers [][]bool) *Graph {
	n := len(blockers)
	if n == 0 {
		return NewGraph(0)
	}

	graph := NewGraph(n)

	for i := range n {
		for j := range len(blockers[i]) {
			if blockers[i][j] {
				graph.AddEdge(j, i)
			}
		}
	}

	return graph
}

func NewCycleDetector(graph *Graph) *CycleDetector {
	return &CycleDetector{
		graph:  graph,
		color:  make([]int, graph.numNodes),
		parent: make([]int, graph.numNodes),
		cycles: [][]int{},
	}
}

func (cd *CycleDetector) dfs(node int, path []int) {
	cd.color[node] = GRAY
	path = append(path, node)

	for _, neighbor := range cd.graph.adjacencyList[node] {
		if cd.color[neighbor] == GRAY {
			cycle := cd.extractCycle(path, neighbor)
			if len(cycle) > 0 {
				cd.cycles = append(cd.cycles, cycle)
			}
		} else if cd.color[neighbor] == WHITE {
			cd.parent[neighbor] = node
			cd.dfs(neighbor, path)
		}
	}

	cd.color[node] = BLACK
}

func (cd *CycleDetector) extractCycle(path []int, backEdgeTarget int) []int {
	cycle := []int{}
	found := false

	for _, node := range path {
		if node == backEdgeTarget {
			found = true
		}
		if found {
			cycle = append(cycle, node)
		}
	}

	return cycle
}

func (cd *CycleDetector) findAllCycles() [][]int {
	for i := range cd.graph.numNodes {
		cd.color[i] = WHITE
		cd.parent[i] = -1
	}

	for i := range cd.graph.numNodes {
		if cd.color[i] == WHITE {
			cd.dfs(i, []int{})
		}
	}

	return cd.removeDuplicateCycles(cd.cycles)
}

func (cd *CycleDetector) removeDuplicateCycles(cycles [][]int) [][]int {
	seen := make(map[string]bool)
	uniqueCycles := [][]int{}

	for _, cycle := range cycles {
		if len(cycle) == 0 {
			continue
		}

		normalized := cd.normalizeCycle(cycle)
		key := cd.cycleToString(normalized)

		if !seen[key] {
			seen[key] = true
			uniqueCycles = append(uniqueCycles, cycle)
		}
	}

	return uniqueCycles
}

func (cd *CycleDetector) normalizeCycle(cycle []int) []int {
	if len(cycle) == 0 {
		return cycle
	}

	minIdx := 0
	for i := 1; i < len(cycle); i++ {
		if cycle[i] < cycle[minIdx] {
			minIdx = i
		}
	}

	normalized := make([]int, len(cycle))
	for i := range len(cycle) {
		normalized[i] = cycle[(minIdx+i)%len(cycle)]
	}

	return normalized
}

func (cd *CycleDetector) cycleToString(cycle []int) string {
	result := ""
	for i, node := range cycle {
		if i > 0 {
			result += ","
		}
		result += string(rune('0' + node))
	}
	return result
}

func detectCycles(blockers [][]bool) [][]int {
	if len(blockers) == 0 {
		return [][]int{}
	}

	graph := buildGraphFromBlockers(blockers)
	detector := NewCycleDetector(graph)
	return detector.findAllCycles()
}

type TestCase struct {
	Name     string
	Blockers [][]bool
	Expected [][]int
}

func Run() any {
	testCases := []TestCase{
		{
			Name: "Simple 3-node cycle",
			Blockers: [][]bool{
				{false, false, true},
				{true, false, false},
				{false, true, false},
			},
			Expected: [][]int{{0, 2, 1}},
		},
		{
			Name: "Two separate cycles",
			Blockers: [][]bool{
				{false, false, true, false, false},
				{true, false, false, false, false},
				{false, true, false, false, false},
				{false, false, false, false, true},
				{false, false, false, true, false},
			},
			Expected: [][]int{{0, 2, 1}, {3, 4}},
		},
		{
			Name: "Self-loop",
			Blockers: [][]bool{
				{true, false},
				{false, false},
			},
			Expected: [][]int{{0}},
		},
		{
			Name: "No cycles",
			Blockers: [][]bool{
				{false, true, false},
				{false, false, true},
				{false, false, false},
			},
			Expected: [][]int{},
		},
		{
			Name: "Complex graph with one cycle",
			Blockers: [][]bool{
				{false, true, false, false},
				{false, false, true, false},
				{false, false, false, true},
				{true, false, false, false},
			},
			Expected: [][]int{{0, 1, 2, 3}},
		},
	}

	results := make([][][]int, len(testCases))
	for i, tc := range testCases {
		results[i] = detectCycles(tc.Blockers)
	}

	return results
}
