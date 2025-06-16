package ford_fulkerson_algorithm

import (
	"errors"
	"fmt"
	"math"
)

type Edge struct {
	From     int
	To       int
	Capacity float64
	Flow     float64
}

type FlowNetwork struct {
	vertices int
	adjList  map[int][]int
	edges    []Edge
	edgeMap  map[string]int
}

type MaxFlowResult struct {
	MaxFlow   float64
	MinCut    []Edge
	FlowEdges []Edge
	Source    int
	Sink      int
}

type AugmentingPath struct {
	Path       []int
	Bottleneck float64
	Edges      []int
}

func NewFlowNetwork(vertices int) *FlowNetwork {
	return &FlowNetwork{
		vertices: vertices,
		adjList:  make(map[int][]int),
		edges:    []Edge{},
		edgeMap:  make(map[string]int),
	}
}

func (fn *FlowNetwork) AddEdge(from, to int, capacity float64) error {
	if from < 0 || from >= fn.vertices || to < 0 || to >= fn.vertices {
		return errors.New("vertex index out of range")
	}

	if capacity < 0 {
		return errors.New("capacity must be non-negative")
	}

	edgeKey := fmt.Sprintf("%d-%d", from, to)
	if existingIdx, exists := fn.edgeMap[edgeKey]; exists {
		fn.edges[existingIdx].Capacity += capacity
		return nil
	}

	edge := Edge{
		From:     from,
		To:       to,
		Capacity: capacity,
		Flow:     0,
	}

	edgeIdx := len(fn.edges)
	fn.edges = append(fn.edges, edge)
	fn.edgeMap[edgeKey] = edgeIdx

	fn.adjList[from] = append(fn.adjList[from], edgeIdx)

	reverseKey := fmt.Sprintf("%d-%d", to, from)
	if _, exists := fn.edgeMap[reverseKey]; !exists {
		reverseEdge := Edge{
			From:     to,
			To:       from,
			Capacity: 0,
			Flow:     0,
		}
		reverseIdx := len(fn.edges)
		fn.edges = append(fn.edges, reverseEdge)
		fn.edgeMap[reverseKey] = reverseIdx
		fn.adjList[to] = append(fn.adjList[to], reverseIdx)
	}

	return nil
}

func (fn *FlowNetwork) GetVertexCount() int {
	return fn.vertices
}

func (fn *FlowNetwork) GetEdgeCount() int {
	return len(fn.edges)
}

func (fn *FlowNetwork) GetEdges() []Edge {
	return fn.edges
}

func (fn *FlowNetwork) GetResidualCapacity(edgeIdx int) float64 {
	edge := fn.edges[edgeIdx]
	return edge.Capacity - edge.Flow
}

func (fn *FlowNetwork) FindAugmentingPathDFS(source, sink int) *AugmentingPath {
	visited := make([]bool, fn.vertices)
	path := []int{}
	edgePath := []int{}

	var dfs func(int, float64) float64
	dfs = func(u int, minCapacity float64) float64 {
		if u == sink {
			return minCapacity
		}

		visited[u] = true

		for _, edgeIdx := range fn.adjList[u] {
			edge := fn.edges[edgeIdx]
			residualCap := fn.GetResidualCapacity(edgeIdx)

			if !visited[edge.To] && residualCap > 0 {
				path = append(path, edge.To)
				edgePath = append(edgePath, edgeIdx)

				bottleneck := math.Min(minCapacity, residualCap)
				result := dfs(edge.To, bottleneck)

				if result > 0 {
					return result
				}

				path = path[:len(path)-1]
				edgePath = edgePath[:len(edgePath)-1]
			}
		}

		return 0
	}

	path = append(path, source)
	bottleneck := dfs(source, math.Inf(1))

	if bottleneck > 0 {
		return &AugmentingPath{
			Path:       path,
			Bottleneck: bottleneck,
			Edges:      edgePath,
		}
	}

	return nil
}

func (fn *FlowNetwork) FindAugmentingPathBFS(source, sink int) *AugmentingPath {
	visited := make([]bool, fn.vertices)
	parent := make([]int, fn.vertices)
	parentEdge := make([]int, fn.vertices)

	for i := range fn.vertices {
		parent[i] = -1
		parentEdge[i] = -1
	}

	queue := []int{source}
	visited[source] = true

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		if u == sink {
			break
		}

		for _, edgeIdx := range fn.adjList[u] {
			edge := fn.edges[edgeIdx]
			residualCap := fn.GetResidualCapacity(edgeIdx)

			if !visited[edge.To] && residualCap > 0 {
				visited[edge.To] = true
				parent[edge.To] = u
				parentEdge[edge.To] = edgeIdx
				queue = append(queue, edge.To)
			}
		}
	}

	if !visited[sink] {
		return nil
	}

	path := []int{}
	edgePath := []int{}
	bottleneck := math.Inf(1)

	current := sink
	for current != source {
		path = append([]int{current}, path...)
		edgeIdx := parentEdge[current]
		edgePath = append([]int{edgeIdx}, edgePath...)
		residualCap := fn.GetResidualCapacity(edgeIdx)
		bottleneck = math.Min(bottleneck, residualCap)
		current = parent[current]
	}
	path = append([]int{source}, path...)

	return &AugmentingPath{
		Path:       path,
		Bottleneck: bottleneck,
		Edges:      edgePath,
	}
}

func (fn *FlowNetwork) AugmentFlow(augPath *AugmentingPath) {
	for _, edgeIdx := range augPath.Edges {
		fn.edges[edgeIdx].Flow += augPath.Bottleneck

		edge := fn.edges[edgeIdx]
		reverseKey := fmt.Sprintf("%d-%d", edge.To, edge.From)
		if reverseIdx, exists := fn.edgeMap[reverseKey]; exists {
			fn.edges[reverseIdx].Flow -= augPath.Bottleneck
		}
	}
}

func (fn *FlowNetwork) FordFulkersonDFS(source, sink int) (*MaxFlowResult, error) {
	if source < 0 || source >= fn.vertices || sink < 0 || sink >= fn.vertices {
		return nil, errors.New("invalid source or sink vertex")
	}

	if source == sink {
		return nil, errors.New("source and sink cannot be the same")
	}

	fn.resetFlow()
	maxFlow := 0.0

	for {
		augPath := fn.FindAugmentingPathDFS(source, sink)
		if augPath == nil {
			break
		}

		fn.AugmentFlow(augPath)
		maxFlow += augPath.Bottleneck
	}

	minCut := fn.FindMinCut(source)
	flowEdges := fn.GetFlowEdges()

	return &MaxFlowResult{
		MaxFlow:   maxFlow,
		MinCut:    minCut,
		FlowEdges: flowEdges,
		Source:    source,
		Sink:      sink,
	}, nil
}

func (fn *FlowNetwork) FordFulkersonBFS(source, sink int) (*MaxFlowResult, error) {
	if source < 0 || source >= fn.vertices || sink < 0 || sink >= fn.vertices {
		return nil, errors.New("invalid source or sink vertex")
	}

	if source == sink {
		return nil, errors.New("source and sink cannot be the same")
	}

	fn.resetFlow()
	maxFlow := 0.0

	for {
		augPath := fn.FindAugmentingPathBFS(source, sink)
		if augPath == nil {
			break
		}

		fn.AugmentFlow(augPath)
		maxFlow += augPath.Bottleneck
	}

	minCut := fn.FindMinCut(source)
	flowEdges := fn.GetFlowEdges()

	return &MaxFlowResult{
		MaxFlow:   maxFlow,
		MinCut:    minCut,
		FlowEdges: flowEdges,
		Source:    source,
		Sink:      sink,
	}, nil
}

func (fn *FlowNetwork) FindMinCut(source int) []Edge {
	visited := make([]bool, fn.vertices)

	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		for _, edgeIdx := range fn.adjList[u] {
			edge := fn.edges[edgeIdx]
			if !visited[edge.To] && fn.GetResidualCapacity(edgeIdx) > 0 {
				dfs(edge.To)
			}
		}
	}

	dfs(source)

	minCut := []Edge{}
	for i, edge := range fn.edges {
		if visited[edge.From] && !visited[edge.To] && edge.Capacity > 0 {
			minCut = append(minCut, fn.edges[i])
		}
	}

	return minCut
}

func (fn *FlowNetwork) GetFlowEdges() []Edge {
	flowEdges := []Edge{}
	for _, edge := range fn.edges {
		if edge.Flow > 0 && edge.Capacity > 0 {
			flowEdges = append(flowEdges, edge)
		}
	}
	return flowEdges
}

func (fn *FlowNetwork) resetFlow() {
	for i := range fn.edges {
		fn.edges[i].Flow = 0
	}
}

func (fn *FlowNetwork) GetTotalCapacity() float64 {
	total := 0.0
	for _, edge := range fn.edges {
		if edge.Capacity > 0 {
			total += edge.Capacity
		}
	}
	return total
}

func (fn *FlowNetwork) IsValidFlow() bool {
	for i := range fn.vertices {
		inFlow := 0.0
		outFlow := 0.0

		for _, edge := range fn.edges {
			if edge.To == i && edge.Capacity > 0 {
				inFlow += edge.Flow
			}
			if edge.From == i && edge.Capacity > 0 {
				outFlow += edge.Flow
			}
		}

		if math.Abs(inFlow-outFlow) > 1e-9 {
			return false
		}
	}
	return true
}

func (result *MaxFlowResult) GetMaxFlow() float64 {
	return result.MaxFlow
}

func (result *MaxFlowResult) GetMinCut() []Edge {
	return result.MinCut
}

func (result *MaxFlowResult) GetMinCutCapacity() float64 {
	capacity := 0.0
	for _, edge := range result.MinCut {
		capacity += edge.Capacity
	}
	return capacity
}

func (result *MaxFlowResult) GetFlowEdges() []Edge {
	return result.FlowEdges
}

func (fn *FlowNetwork) PrintNetwork() {
	fmt.Printf("Flow Network with %d vertices:\n", fn.vertices)
	for i, edge := range fn.edges {
		if edge.Capacity > 0 {
			fmt.Printf("  %d: %d -> %d (capacity: %.1f, flow: %.1f)\n",
				i, edge.From, edge.To, edge.Capacity, edge.Flow)
		}
	}
}

func (result *MaxFlowResult) PrintResult() {
	fmt.Printf("Maximum Flow Result:\n")
	fmt.Printf("Source: %d, Sink: %d\n", result.Source, result.Sink)
	fmt.Printf("Maximum Flow: %.2f\n", result.MaxFlow)
	fmt.Printf("Minimum Cut Capacity: %.2f\n", result.GetMinCutCapacity())

	fmt.Printf("Flow Edges:\n")
	for _, edge := range result.FlowEdges {
		fmt.Printf("  %d -> %d: %.2f/%.2f\n", edge.From, edge.To, edge.Flow, edge.Capacity)
	}

	fmt.Printf("Minimum Cut Edges:\n")
	for _, edge := range result.MinCut {
		fmt.Printf("  %d -> %d: %.2f\n", edge.From, edge.To, edge.Capacity)
	}
}

func Run() any {
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

	result := make(map[string]any)
	result["vertexCount"] = fn.GetVertexCount()
	result["edgeCount"] = fn.GetEdgeCount()
	result["totalCapacity"] = fn.GetTotalCapacity()

	dfsResult, dfsErr := fn.FordFulkersonDFS(0, 5)
	if dfsErr == nil {
		result["dfsMaxFlow"] = dfsResult.GetMaxFlow()
		result["dfsMinCutCapacity"] = dfsResult.GetMinCutCapacity()
		result["dfsFlowEdgeCount"] = len(dfsResult.GetFlowEdges())
		result["dfsMinCutEdgeCount"] = len(dfsResult.GetMinCut())

		dfsFlowEdges := make([]map[string]any, len(dfsResult.GetFlowEdges()))
		for i, edge := range dfsResult.GetFlowEdges() {
			dfsFlowEdges[i] = map[string]any{
				"from":     edge.From,
				"to":       edge.To,
				"flow":     edge.Flow,
				"capacity": edge.Capacity,
			}
		}
		result["dfsFlowEdges"] = dfsFlowEdges

		dfsMinCutEdges := make([]map[string]any, len(dfsResult.GetMinCut()))
		for i, edge := range dfsResult.GetMinCut() {
			dfsMinCutEdges[i] = map[string]any{
				"from":     edge.From,
				"to":       edge.To,
				"capacity": edge.Capacity,
			}
		}
		result["dfsMinCutEdges"] = dfsMinCutEdges
	} else {
		result["dfsError"] = dfsErr.Error()
	}

	bfsResult, bfsErr := fn.FordFulkersonBFS(0, 5)
	if bfsErr == nil {
		result["bfsMaxFlow"] = bfsResult.GetMaxFlow()
		result["bfsMinCutCapacity"] = bfsResult.GetMinCutCapacity()
		result["bfsFlowEdgeCount"] = len(bfsResult.GetFlowEdges())
		result["bfsMinCutEdgeCount"] = len(bfsResult.GetMinCut())

		bfsFlowEdges := make([]map[string]any, len(bfsResult.GetFlowEdges()))
		for i, edge := range bfsResult.GetFlowEdges() {
			bfsFlowEdges[i] = map[string]any{
				"from":     edge.From,
				"to":       edge.To,
				"flow":     edge.Flow,
				"capacity": edge.Capacity,
			}
		}
		result["bfsFlowEdges"] = bfsFlowEdges

		bfsMinCutEdges := make([]map[string]any, len(bfsResult.GetMinCut()))
		for i, edge := range bfsResult.GetMinCut() {
			bfsMinCutEdges[i] = map[string]any{
				"from":     edge.From,
				"to":       edge.To,
				"capacity": edge.Capacity,
			}
		}
		result["bfsMinCutEdges"] = bfsMinCutEdges
	} else {
		result["bfsError"] = bfsErr.Error()
	}

	invalidSourceSink := NewFlowNetwork(3)
	invalidSourceSink.AddEdge(0, 1, 10)
	_, invalidErr := invalidSourceSink.FordFulkersonDFS(0, 0)
	result["invalidSourceSinkError"] = invalidErr != nil
	if invalidErr != nil {
		result["invalidErrorMessage"] = invalidErr.Error()
	}

	singleVertex := NewFlowNetwork(1)
	_, singleErr := singleVertex.FordFulkersonDFS(0, 0)
	result["singleVertexError"] = singleErr != nil

	disconnected := NewFlowNetwork(4)
	disconnected.AddEdge(0, 1, 10)
	disconnected.AddEdge(2, 3, 5)
	disconnectedResult, _ := disconnected.FordFulkersonDFS(0, 3)
	result["disconnectedMaxFlow"] = disconnectedResult.GetMaxFlow()

	return result
}
