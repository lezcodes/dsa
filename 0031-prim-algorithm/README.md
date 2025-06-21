# prim-algorithm

## Description

Prim's algorithm is a greedy algorithm that finds a minimum spanning tree (MST) for a weighted undirected graph. It builds the MST by starting from an arbitrary vertex and repeatedly adding the minimum-weight edge that connects a vertex in the MST to a vertex outside the MST.

This implementation provides two approaches:

- **Priority Queue Approach**: Uses a min-heap for efficient edge selection (O(E log V))
- **Simple Approach**: Uses linear search for minimum edge selection (O(V²))

## Key Features

- Weighted undirected graph support
- Two algorithm implementations for different performance characteristics
- Comprehensive connectivity checking
- Support for floating-point and negative weights
- Detailed MST properties and validation
- Step-by-step algorithm visualization support

## Complexity

- **Priority Queue Approach**:
  - Time Complexity: O(E log V) where E = edges, V = vertices
  - Space Complexity: O(V + E)
- **Simple Approach**:
  - Time Complexity: O(V²)
  - Space Complexity: O(V + E)

## Algorithm Steps

1. Start with an arbitrary vertex (typically vertex 0)
2. Mark the starting vertex as visited
3. Add all edges from visited vertices to the priority queue/candidate set
4. Repeat until MST is complete:
   - Select the minimum-weight edge connecting visited to unvisited vertex
   - Add the edge to MST and mark the destination vertex as visited
   - Add all edges from the newly visited vertex to candidates
5. Return the MST with total cost

## Real-World Applications

- **Network Design**: Designing minimum-cost communication networks
- **Circuit Design**: Connecting components with minimum wire length
- **Transportation**: Building road/railway networks with minimum cost
- **Clustering**: Creating hierarchical clusters in data analysis
- **Approximation Algorithms**: TSP approximation and other optimization problems

## API Reference

### Graph Operations

```go
g := NewGraph(vertices)           // Create graph with n vertices
g.AddEdge(from, to, weight)       // Add weighted edge
g.GetVertexCount()                // Get number of vertices
g.GetEdgeCount()                  // Get number of edges
g.GetNeighbors(vertex)            // Get adjacent edges
g.IsConnected()                   // Check if graph is connected
```

### MST Algorithms

```go
mst, err := g.PrimMST()           // Priority queue approach
mst, err := g.PrimMSTSimple()     // Simple O(V²) approach
```

### MST Properties

```go
mst.GetEdges()                    // Get MST edges
mst.GetTotalCost()                // Get total weight
mst.GetVertexCount()              // Get vertex count
mst.GetEdgeCount()                // Get edge count
mst.IsComplete()                  // Check if MST is complete
```

### Utility Functions

```go
g.PrintGraph()                    // Print graph structure
mst.PrintMST()                    // Print MST details
```

## Usage

```bash
make run n=0031-prim-algorithm
```

## Testing

```bash
make test n=0031-prim-algorithm
```

## Implementation Details

- Uses adjacency list representation for efficient neighbor access
- Priority queue implemented with Go's container/heap package
- Supports disconnected graph detection with appropriate error handling
- Handles edge cases: empty graphs, single vertices, negative weights
- Comprehensive test coverage including large graphs and performance benchmarks

## Visual Representation

### Sample Graph for MST

```mermaid
graph LR
    A[A] ---|2| B[B]
    A ---|3| C[C]
    A ---|1| D[D]
    B ---|4| C
    B ---|5| E[E]
    C ---|6| E
    D ---|7| E

    F["Graph with edge weights"]
    G["Goal: Find MST with minimum total weight"]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
```

### Prim's Algorithm Steps

```mermaid
graph TD
    A[Start with arbitrary vertex] --> B[Add to MST set]
    B --> C[Initialize priority queue with adjacent edges]
    C --> D{Queue empty?}
    D -->|Yes| E[MST Complete]
    D -->|No| F[Extract minimum weight edge]
    F --> G{Both vertices in MST?}
    G -->|Yes| H[Skip edge (would create cycle)]
    G -->|No| I[Add edge to MST]
    I --> J[Add new vertex to MST set]
    J --> K[Add new adjacent edges to queue]
    K --> L[Update priorities if needed]
    L --> D
    H --> D

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Step-by-Step MST Construction

```mermaid
graph LR
    subgraph "Step 1: Start with A"
        A1[A]
        E1["MST: {A}"]
        E1_W["Weight: 0"]
    end

    subgraph "Step 2: Add edge A-D (weight 1)"
        A2[A] ---|1| D2[D]
        E2["MST: {A, D}"]
        E2_W["Weight: 1"]
    end

    subgraph "Step 3: Add edge A-B (weight 2)"
        A3[A] ---|1| D3[D]
        A3 ---|2| B3[B]
        E3["MST: {A, D, B}"]
        E3_W["Weight: 3"]
    end

    subgraph "Step 4: Add edge A-C (weight 3)"
        A4[A] ---|1| D4[D]
        A4 ---|2| B4[B]
        A4 ---|3| C4[C]
        E4["MST: {A, D, B, C}"]
        E4_W["Weight: 6"]
    end

    style A1 fill:#e1f5fe
    style E4_W fill:#c8e6c9
```

### Priority Queue Operations

```mermaid
graph TD
    A[Priority Queue Management] --> B[Edge Representation]
    A --> C[Heap Operations]
    A --> D[Update Strategies]

    B --> B1["Store (weight, vertex_from, vertex_to)"]
    B --> B2["Min-heap based on weight"]
    B --> B3["Handle duplicate edges"]

    C --> C1["Insert: O(log E)"]
    C --> C2["Extract-min: O(log E)"]
    C --> C3["Decrease-key: O(log E)"]

    D --> D1["Lazy approach: Allow duplicates"]
    D --> D2["Eager approach: Update priorities"]
    D --> D3["Fibonacci heap for better performance"]

    style A fill:#e1f5fe
    style C2 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Lazy vs Eager Implementation

```mermaid
graph LR
    subgraph "Lazy Prim's"
        A["Add all edges to priority queue"]
        B["Allow duplicate edges"]
        C["Check if edge creates cycle when extracted"]
        D["Simpler implementation"]
        E["Higher space complexity"]
    end

    subgraph "Eager Prim's"
        F["Track minimum edge to each vertex"]
        G["Update priorities when better edge found"]
        H["No duplicates in queue"]
        I["More complex implementation"]
        J["Lower space complexity"]
    end

    style D fill:#c8e6c9
    style J fill:#c8e6c9
```

### MST Properties and Invariants

```mermaid
graph TD
    A[MST Properties] --> B[Cut Property]
    A --> C[Cycle Property]
    A --> D[Uniqueness]

    B --> B1["Minimum edge crossing any cut is in MST"]
    B --> B2["Justifies greedy choice"]
    B --> B3["Foundation of Prim's algorithm"]

    C --> C1["Maximum edge in any cycle is not in MST"]
    C --> C2["Helps identify unnecessary edges"]

    D --> D1["MST is unique if all edge weights distinct"]
    D --> D2["Multiple MSTs possible with equal weights"]
    D --> D3["All MSTs have same total weight"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style D3 fill:#c8e6c9
```

### Algorithm Complexity Analysis

```mermaid
graph TD
    A[Prim's Algorithm Complexity] --> B[Time Complexity]
    A --> C[Space Complexity]

    B --> B1["Binary Heap: O(E log V)"]
    B --> B2["Fibonacci Heap: O(E + V log V)"]
    B --> B3["Dense graphs: O(V²) with arrays"]
    B --> B4["Depends on priority queue implementation"]

    C --> C1["Adjacency list: O(V + E)"]
    C --> C2["Priority queue: O(E) lazy, O(V) eager"]
    C --> C3["MST storage: O(V)"]

    D[Performance Notes] --> E["Dense graphs favor array implementation"]
    D --> F["Sparse graphs favor heap implementation"]
    D --> G["Fibonacci heap best asymptotically"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style F fill:#c8e6c9
```

### Prim's vs Kruskal's Algorithm

```mermaid
graph TD
    A[MST Algorithm Comparison] --> B[Prim's Algorithm]
    A --> C[Kruskal's Algorithm]

    B --> B1["Vertex-based approach"]
    B --> B2["Grows single tree"]
    B --> B3["Better for dense graphs"]
    B --> B4["Uses priority queue"]
    B --> B5["O(E log V) with binary heap"]

    C --> C1["Edge-based approach"]
    C --> C2["Merges forest of trees"]
    C --> C3["Better for sparse graphs"]
    C --> C4["Uses Union-Find"]
    C --> C5["O(E log E) time"]

    D[Selection Criteria] --> E["Dense graphs: Use Prim's"]
    D --> F["Sparse graphs: Use Kruskal's"]
    D --> G["Both produce same MST weight"]

    style B3 fill:#c8e6c9
    style C3 fill:#c8e6c9
    style G fill:#c8e6c9
```

### Cut Property Visualization

```mermaid
graph LR
    subgraph "Graph with Cut"
        A[A] ---|2| B[B]
        A ---|1| C[C]
        B ---|3| D[D]
        C ---|4| D

        Cut["Cut separates {A,C} from {B,D}"]
        MinEdge["Minimum crossing edge: A-B (weight 2)"]
    end

    subgraph "Cut Property"
        Prop["The minimum weight edge crossing any cut"]
        Prop2["is guaranteed to be in some MST"]
        Proof["This justifies Prim's greedy choice"]
    end

    style MinEdge fill:#c8e6c9
    style Proof fill:#c8e6c9
```

### Implementation Variants

```mermaid
graph TD
    A[Prim's Implementation Variants] --> B[Data Structure Choice]
    A --> C[Graph Representation]
    A --> D[Optimization Level]

    B --> B1["Binary heap (simple)"]
    B --> B2["Fibonacci heap (optimal)"]
    B --> B3["Array-based (dense graphs)"]

    C --> C1["Adjacency matrix"]
    C --> C2["Adjacency list"]
    C --> C3["Edge list"]

    D --> D1["Basic implementation"]
    D --> D2["Lazy deletion optimization"]
    D --> D3["Path compression techniques"]

    E[Practical Considerations] --> F["Use binary heap for most cases"]
    E --> G["Array-based for very dense graphs"]
    E --> H["Consider memory vs time trade-offs"]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style H fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[Prim's Algorithm Applications] --> B[Network Design]
    A --> C[Infrastructure Planning]
    A --> D[Circuit Design]
    A --> E[Clustering]

    B --> B1["Computer networks"]
    B --> B2["Telecommunication networks"]
    B --> B3["Internet backbone design"]

    C --> C1["Road networks"]
    C --> C2["Power grid design"]
    C --> C3["Water distribution systems"]

    D --> D1["VLSI circuit layout"]
    D --> D2["PCB routing"]
    D --> D3["Electronic component placement"]

    E --> E1["Hierarchical clustering"]
    E --> E2["Data mining"]
    E --> E3["Image segmentation"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Detailed Algorithm Trace

```mermaid
graph LR
    subgraph "Initial Graph"
        A[A] ---|2| B[B]
        A ---|3| C[C]
        A ---|1| D[D]
        B ---|4| C
        B ---|5| E[E]
        C ---|6| E
        D ---|7| E
    end

    subgraph "Priority Queue States"
        Q1["Initial: []"]
        Q2["After A: [(1,A,D), (2,A,B), (3,A,C)]"]
        Q3["After D: [(2,A,B), (3,A,C), (7,D,E)]"]
        Q4["After B: [(3,A,C), (4,B,C), (5,B,E), (7,D,E)]"]
        Q5["After C: [(5,B,E), (6,C,E), (7,D,E)]"]
        Q6["After E: []"]
    end

    style A fill:#e1f5fe
    style Q6 fill:#c8e6c9
```

### Error Handling and Edge Cases

```mermaid
graph TD
    A[Edge Cases and Errors] --> B[Disconnected Graph]
    A --> C[Single Vertex]
    A --> D[Negative Weights]
    A --> E[Duplicate Edges]

    B --> B1["No spanning tree exists"]
    B --> B2["Return forest of MSTs"]
    B --> B3["Detect using vertex count"]

    C --> C1["MST is the single vertex"]
    C --> C2["No edges needed"]

    D --> D1["Algorithm still works"]
    D --> D2["MST weight can be negative"]

    E --> E1["Keep minimum weight edge"]
    E --> E2["Handle in preprocessing"]

    F[Implementation Checks] --> G["Verify graph connectivity"]
    F --> H["Handle empty graphs"]
    F --> I["Validate edge weights"]

    style A fill:#e1f5fe
    style D1 fill:#c8e6c9
    style G fill:#c8e6c9
```

### Optimization Techniques

```mermaid
graph TD
    A[Prim's Optimizations] --> B[Data Structure]
    A --> C[Algorithm Variants]
    A --> D[Implementation]

    B --> B1["Fibonacci heap for dense graphs"]
    B --> B2["Binary heap for sparse graphs"]
    B --> B3["Simple array for very dense graphs"]

    C --> C1["Lazy approach (simpler)"]
    C --> C2["Eager approach (efficient)"]
    C --> C3["Hybrid approaches"]

    D --> D1["Precompute adjacency lists"]
    D --> D2["Use bit manipulation"]
    D --> D3["Memory pool allocation"]

    E[Performance Tips] --> F["Profile with actual data"]
    E --> G["Consider graph characteristics"]
    E --> H["Balance complexity vs performance"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style F fill:#c8e6c9
```

Prim's Algorithm is a greedy algorithm that finds a minimum spanning tree (MST) for a weighted undirected graph.
