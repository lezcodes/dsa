# ford-fulkerson-algorithm

## Description

The Ford-Fulkerson algorithm is a method for computing the maximum flow in a flow network. It uses the concept of augmenting paths to iteratively increase the flow until no more augmenting paths can be found. The algorithm also finds the minimum cut of the network, which by the Max-Flow Min-Cut theorem has the same value as the maximum flow.

## Flow Network Visualization

```mermaid
graph LR
    subgraph "Flow Network Example"
        S((Source<br/>S)) --> A((A))
        S --> B((B))
        A --> C((C))
        B --> C
        A --> T((Sink<br/>T))
        C --> T
    end

    subgraph "Edge Labels: capacity/flow"
        S2((S)) --"16/0"--> A2((A))
        S2 --"13/0"--> B2((B))
        A2 --"10/0"--> C2((C))
        B2 --"4/0"--> C2
        A2 --"12/0"--> T2((T))
        C2 --"14/0"--> T2
    end

    style S fill:#e1f5fe
    style T fill:#c8e6c9
    style S2 fill:#e1f5fe
    style T2 fill:#c8e6c9
```

## Algorithm Flow

```mermaid
flowchart TD
    A["Initialize Flow Network<br/>All flows = 0"] --> B{Find Augmenting Path<br/>using DFS or BFS}
    B -->|Path Found| C["Find Bottleneck<br/>min capacity along path"]
    C --> D["Augment Flow<br/>along the path"]
    D --> E["Update Residual Graph<br/>forward & backward edges"]
    E --> B
    B -->|No Path Found| F["Maximum Flow Found<br/>Compute Min-Cut"]
    F --> G["Return Result<br/>Max Flow = Min Cut"]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style B fill:#fff3e0
    style F fill:#c8e6c9
```

## Step-by-Step Execution

```mermaid
graph TD
    subgraph "Step 1: Find First Augmenting Path"
        S1((S)) --"16"--> A1((A))
        S1 --"13"--> B1((B))
        A1 --"12"--> T1((T))
        A1 --"10"--> C1((C))
        B1 --"4"--> C1
        C1 --"14"--> T1
    end

    subgraph "Step 2: Augment Flow (Bottleneck = 12)"
        S2((S)) --"16/12"--> A2((A))
        S2 --"13/0"--> B2((B))
        A2 --"12/12"--> T2((T))
        A2 --"10/0"--> C2((C))
        B2 --"4/0"--> C2
        C2 --"14/0"--> T2
    end

    subgraph "Step 3: Find Second Path & Final Flow"
        S3((S)) --"16/16"--> A3((A))
        S3 --"13/13"--> B3((B))
        A3 --"12/12"--> T3((T))
        A3 --"10/4"--> C3((C))
        B3 --"4/4"--> C3
        C3 --"14/4"--> T3
    end

    style S1 fill:#e1f5fe
    style S2 fill:#fff3e0
    style S3 fill:#c8e6c9
    style T1 fill:#c8e6c9
    style T2 fill:#c8e6c9
    style T3 fill:#c8e6c9
```

## Residual Graph Concept

```mermaid
graph LR
    subgraph "Original Graph"
        A1((A)) --"10/6"--> B1((B))
    end

    subgraph "Residual Graph"
        A2((A)) --"4<br/>(10-6)"--> B2((B))
        B2 --"6<br/>(flow)"--> A2
    end

    subgraph "Legend"
        L1[Forward Edge: remaining capacity] --- L2[Backward Edge: current flow]
    end

    style A1 fill:#e1f5fe
    style B1 fill:#c8e6c9
    style A2 fill:#e1f5fe
    style B2 fill:#c8e6c9
```

## Min-Cut Visualization

```mermaid
graph LR
    subgraph "Minimum Cut Partition"
        subgraph "Source Side (S-reachable)"
            S((S))
            A((A))
        end

        subgraph "Sink Side (T-reachable)"
            B((B))
            C((C))
            T((T))
        end

        S -.->|Cut Edge<br/>16/16| A
        A -.->|Cut Edge<br/>10/4| C
        S --> B
        B --> C
        C --> T
    end

    style S fill:#e1f5fe
    style A fill:#e1f5fe
    style B fill:#ffcdd2
    style C fill:#ffcdd2
    style T fill:#ffcdd2
```

## Algorithm Comparison

```mermaid
graph TB
    subgraph "Ford-Fulkerson Variants"
        A["Ford-Fulkerson<br/>with DFS"] --> A1["Time: O(E × f)<br/>Can be exponential"]
        B["Edmonds-Karp<br/>with BFS"] --> B1["Time: O(V × E²)<br/>Polynomial guarantee"]
        C["Dinic's Algorithm"] --> C1["Time: O(V² × E)<br/>Best for dense graphs"]
        D["Push-Relabel"] --> D1["Time: O(V³)<br/>Good for sparse graphs"]
    end

    style A fill:#fff3e0
    style B fill:#c8e6c9
    style C fill:#e1f5fe
    style D fill:#e1f5fe
```

This implementation provides:

- **DFS-based Path Finding**: Uses depth-first search to find augmenting paths
- **BFS-based Path Finding**: Uses breadth-first search (Edmonds-Karp) for better performance
- **Residual Graph Management**: Automatically handles forward and backward edges
- **Min-Cut Computation**: Finds the minimum cut that separates source from sink

## Key Features

- Directed weighted flow networks with capacity constraints
- Automatic residual graph construction with reverse edges
- Both DFS and BFS implementations for path finding
- Maximum flow and minimum cut computation
- Flow conservation validation
- Support for floating-point capacities
- Comprehensive edge case handling

## Complexity

- **Ford-Fulkerson (DFS)**: O(E × f) where E = edges, f = maximum flow value
- **Edmonds-Karp (BFS)**: O(V × E²) where V = vertices, E = edges
- **Space Complexity**: O(V + E) for graph representation and auxiliary structures

## Algorithm Steps

1. **Initialize**: Set all edge flows to 0, create residual graph with reverse edges
2. **Find Augmenting Path**: Use DFS or BFS to find path from source to sink with positive residual capacity
3. **Compute Bottleneck**: Find minimum residual capacity along the augmenting path
4. **Augment Flow**: Increase flow along forward edges, decrease along backward edges
5. **Repeat**: Continue until no augmenting path exists
6. **Compute Min-Cut**: Find all vertices reachable from source in residual graph

## Real-World Applications

- **Network Routing**: Internet traffic optimization and bandwidth allocation
- **Transportation**: Railway/airline capacity planning and scheduling
- **Supply Chain**: Distribution network optimization and bottleneck analysis
- **Bipartite Matching**: Maximum matching in bipartite graphs
- **Image Segmentation**: Computer vision and medical imaging
- **Project Scheduling**: Resource allocation and critical path analysis

## API Reference

### Flow Network Operations

```go
fn := NewFlowNetwork(vertices)     // Create flow network with n vertices
fn.AddEdge(from, to, capacity)     // Add directed edge with capacity
fn.GetVertexCount()                // Get number of vertices
fn.GetEdgeCount()                  // Get number of edges (including reverse)
fn.GetEdges()                      // Get all edges in network
fn.GetTotalCapacity()              // Get sum of all edge capacities
```

### Maximum Flow Algorithms

```go
result, err := fn.FordFulkersonDFS(source, sink)    // DFS-based Ford-Fulkerson
result, err := fn.FordFulkersonBFS(source, sink)    // BFS-based Edmonds-Karp
```

### Augmenting Path Finding

```go
path := fn.FindAugmentingPathDFS(source, sink)      // DFS path finding
path := fn.FindAugmentingPathBFS(source, sink)      // BFS path finding
```

### Result Analysis

```go
result.GetMaxFlow()                // Get maximum flow value
result.GetMinCut()                 // Get minimum cut edges
result.GetMinCutCapacity()         // Get minimum cut capacity
result.GetFlowEdges()              // Get edges with positive flow
```

### Flow Network Properties

```go
fn.GetResidualCapacity(edgeIdx)    // Get residual capacity of edge
fn.IsValidFlow()                   // Validate flow conservation
fn.FindMinCut(source)              // Find minimum cut from source
```

### Data Structures

```go
type Edge struct {
    From     int     // Source vertex
    To       int     // Destination vertex
    Capacity float64 // Maximum capacity
    Flow     float64 // Current flow
}

type AugmentingPath struct {
    Path       []int     // Vertex sequence
    Bottleneck float64   // Minimum capacity along path
    Edges      []int     // Edge indices in path
}

type MaxFlowResult struct {
    MaxFlow   float64   // Maximum flow value
    MinCut    []Edge    // Minimum cut edges
    FlowEdges []Edge    // Edges with positive flow
    Source    int       // Source vertex
    Sink      int       // Sink vertex
}
```

### Utility Functions

```go
fn.PrintNetwork()                  // Print network structure
result.PrintResult()               // Print flow and cut results
```

## Usage

```bash
make run n=0033-ford-fulkerson-algorithm
```

## Testing

```bash
make test n=0033-ford-fulkerson-algorithm
```

## Implementation Details

- **Residual Graph**: Automatically creates reverse edges with 0 capacity for flow cancellation
- **Edge Management**: Uses edge mapping to handle parallel edges and capacity aggregation
- **Flow Augmentation**: Updates both forward and backward edge flows simultaneously
- **Min-Cut Finding**: Uses DFS on residual graph to identify reachable vertices from source
- **Error Handling**: Comprehensive validation for invalid vertices, negative capacities, and edge cases

## Performance Characteristics

- **DFS Approach**: Faster per iteration but may find longer paths
- **BFS Approach**: Finds shortest augmenting paths, guaranteeing polynomial time
- **Memory Usage**: Linear in vertices and edges with efficient adjacency list representation
- **Numerical Stability**: Handles floating-point capacities with appropriate precision

## Max-Flow Min-Cut Theorem

The algorithm demonstrates the fundamental theorem that in any flow network:

- **Maximum Flow Value** = **Minimum Cut Capacity**
- The minimum cut represents the bottleneck that limits the maximum flow
- Removing min-cut edges disconnects source from sink with minimum capacity loss

## Comparison with Other Algorithms

- **vs. Dinic's Algorithm**: Simpler implementation but potentially slower on dense graphs
- **vs. Push-Relabel**: More intuitive flow augmentation approach
- **vs. Capacity Scaling**: Better for networks with large capacity ranges
- **DFS vs. BFS**: BFS (Edmonds-Karp) guarantees O(VE²) while DFS can be exponential in worst case
