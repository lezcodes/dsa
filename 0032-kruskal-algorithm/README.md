# kruskal-algorithm

## Description

Kruskal's algorithm is a greedy algorithm that finds a minimum spanning tree (MST) for a weighted undirected graph. It works by sorting all edges by weight and then adding edges to the MST in order of increasing weight, using a Union-Find data structure to detect and avoid cycles.

This implementation features:

- **Union-Find with Path Compression**: Optimized disjoint set operations
- **Union by Rank**: Balanced tree structure for efficient operations
- **Step-by-Step Visualization**: Track algorithm decisions and cycle detection
- **Comprehensive Edge Sorting**: Handles duplicate weights correctly

## Key Features

- Weighted undirected graph support
- Efficient Union-Find data structure with path compression and union by rank
- Cycle detection during MST construction
- Algorithm step tracking for educational purposes
- Support for floating-point and negative weights
- Comprehensive connectivity analysis

## Complexity

- **Time Complexity**: O(E log E) where E = edges (dominated by edge sorting)
- **Space Complexity**: O(V + E) where V = vertices
- **Union-Find Operations**: Nearly O(1) amortized with path compression and union by rank

## Algorithm Steps

1. Sort all edges in the graph by weight in ascending order
2. Initialize Union-Find data structure with all vertices as separate components
3. For each edge in sorted order:
   - Check if adding the edge would create a cycle (using Union-Find)
   - If no cycle: add edge to MST and union the components
   - If cycle: reject the edge
4. Continue until MST has V-1 edges (for V vertices)
5. Return the MST with total cost

## Real-World Applications

- **Network Design**: Building minimum-cost communication networks
- **Circuit Design**: Connecting electronic components with minimum wire
- **Transportation**: Designing efficient road/railway networks
- **Clustering**: Creating hierarchical data clusters
- **Image Segmentation**: Computer vision and image processing
- **Social Networks**: Finding minimum connection costs between groups

## API Reference

### Graph Operations

```go
g := NewGraph(vertices)           // Create graph with n vertices
g.AddEdge(from, to, weight)       // Add weighted edge
g.GetVertexCount()                // Get number of vertices
g.GetEdgeCount()                  // Get number of edges
g.GetEdges()                      // Get all edges
g.IsConnected()                   // Check if graph is connected
```

### MST Algorithms

```go
mst, err := g.KruskalMST()                    // Standard Kruskal's algorithm
mst, steps, err := g.KruskalMSTWithSteps()    // With step-by-step tracking
```

### Union-Find Operations

```go
uf := NewUnionFind(n)             // Create Union-Find for n elements
uf.Find(x)                        // Find root with path compression
uf.Union(x, y)                    // Union two components
uf.Connected(x, y)                // Check if two elements are connected
uf.ComponentCount()               // Get number of separate components
```

### MST Properties

```go
mst.GetEdges()                    // Get MST edges
mst.GetTotalCost()                // Get total weight
mst.GetVertexCount()              // Get vertex count
mst.GetEdgeCount()                // Get edge count
mst.IsComplete()                  // Check if MST is complete
```

### Algorithm Steps

```go
type KruskalStep struct {
    Edge     Edge    // The edge being considered
    Accepted bool    // Whether edge was added to MST
    Reason   string  // Explanation of decision
}
```

### Utility Functions

```go
g.PrintGraph()                    // Print graph structure
mst.PrintMST()                    // Print MST details
PrintKruskalSteps(steps)          // Print algorithm steps
```

## Usage

```bash
make run n=0032-kruskal-algorithm
```

## Testing

```bash
make test n=0032-kruskal-algorithm
```

## Implementation Details

- **Edge Sorting**: Uses Go's sort.Slice for stable sorting of edges by weight
- **Union-Find Optimization**: Path compression flattens trees during Find operations
- **Union by Rank**: Attaches smaller trees under larger trees to maintain balance
- **Cycle Detection**: Efficiently detects cycles using Union-Find connectivity checks
- **Memory Efficient**: Stores only necessary edge information, no adjacency matrix
- **Error Handling**: Comprehensive validation for disconnected graphs and invalid inputs

## Performance Characteristics

- **Best Case**: O(E log E) for any connected graph
- **Average Case**: O(E log E) with nearly O(1) Union-Find operations
- **Worst Case**: O(E log E) - Union-Find operations remain efficient
- **Space Usage**: Linear in vertices and edges, no additional data structures needed

## Comparison with Prim's Algorithm

- **Kruskal's**: Better for sparse graphs, global edge perspective
- **Prim's**: Better for dense graphs, local vertex perspective
- **Both**: Produce optimal MST with same total weight
- **Edge Selection**: Kruskal sorts all edges first, Prim grows incrementally

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

    F["Edges sorted by weight:"]
    G["(A,D):1, (A,B):2, (A,C):3, (B,C):4, (B,E):5, (C,E):6, (D,E):7"]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
```

### Kruskal's Algorithm Steps

```mermaid
graph TD
    A[Sort all edges by weight] --> B[Initialize Union-Find structure]
    B --> C[For each edge in sorted order]
    C --> D{Do endpoints belong to same component?}
    D -->|Yes| E["Skip edge (would create cycle)"]
    D -->|No| F[Add edge to MST]
    F --> G[Union the components]
    G --> H{MST has V-1 edges?}
    H -->|No| I[Continue to next edge]
    H -->|Yes| J[MST Complete]
    E --> I
    I --> C

    style A fill:#e1f5fe
    style J fill:#c8e6c9
```

### Step-by-Step MST Construction

```mermaid
graph LR
    subgraph "Step 1: Process edge (A,D) weight 1"
        A1[A] ---|1| D1[D]
        S1["Components: {A,D}, {B}, {C}, {E}"]
        S1_Add["✓ Add to MST"]
    end

    subgraph "Step 2: Process edge (A,B) weight 2"
        A2[A] ---|1| D2[D]
        A2 ---|2| B2[B]
        S2["Components: {A,D,B}, {C}, {E}"]
        S2_Add["✓ Add to MST"]
    end

    subgraph "Step 3: Process edge (A,C) weight 3"
        A3[A] ---|1| D3[D]
        A3 ---|2| B3[B]
        A3 ---|3| C3[C]
        S3["Components: {A,D,B,C}, {E}"]
        S3_Add["✓ Add to MST"]
    end

    subgraph "Step 4: Process edge (B,C) weight 4"
        S4["B and C already connected"]
        S4_Skip["✗ Skip (would create cycle)"]
    end

    style S1_Add fill:#c8e6c9
    style S4_Skip fill:#ffcdd2
```

### Union-Find Data Structure

```mermaid
graph TD
    A[Union-Find Operations] --> B[Find Operation]
    A --> C[Union Operation]
    A --> D[Optimizations]

    B --> B1["Find root of component"]
    B --> B2["Path compression"]
    B --> B3["O(α(n)) amortized"]

    C --> C1["Merge two components"]
    C --> C2["Union by rank/size"]
    C --> C3["Keep trees balanced"]

    D --> D1["Path compression"]
    D --> D2["Union by rank"]
    D --> D3["Nearly constant time"]

    E[Cycle Detection] --> F["If Find(u) == Find(v)"]
    E --> G["Then u and v in same component"]
    E --> H["Adding edge creates cycle"]

    style A fill:#e1f5fe
    style B3 fill:#c8e6c9
    style H fill:#ffcdd2
```

### Union-Find Tree Evolution

```mermaid
graph LR
    subgraph "Initial State"
        A1[A]
        B1[B]
        C1[C]
        D1[D]
        E1[E]
        I1["Each vertex is its own component"]
    end

    subgraph "After Union(A,D)"
        A2[A] --> D2[D]
        B2[B]
        C2[C]
        E2[E]
        I2["Components: {A,D}, {B}, {C}, {E}"]
    end

    subgraph "After Union(A,B)"
        A3[A] --> D3[D]
        A3 --> B3[B]
        C3[C]
        E3[E]
        I3["Components: {A,D,B}, {C}, {E}"]
    end

    style I1 fill:#e1f5fe
    style I3 fill:#c8e6c9
```

### Path Compression Visualization

```mermaid
graph LR
    subgraph "Before Path Compression"
        A1[A] --> B1[B]
        B1 --> C1[C]
        C1 --> D1[D]
        D1 --> E1[E]
        E1 --> F1[F]

        H1["Find(A) traverses: A→B→C→D→E→F"]
    end

    subgraph "After Path Compression"
        A2[A] --> F2[F]
        B2[B] --> F2
        C2[C] --> F2
        D2[D] --> F2
        E2[E] --> F2

        H2["Find(A) now direct: A→F"]
    end

    style H1 fill:#ffcdd2
    style H2 fill:#c8e6c9
```

### Complexity Analysis

```mermaid
graph TD
    A[Kruskal's Algorithm Complexity] --> B[Time Complexity]
    A --> C[Space Complexity]

    B --> B1["Sorting edges: O(E log E)"]
    B --> B2["Union-Find operations: O(E α(V))"]
    B --> B3["Total: O(E log E)"]
    B --> B4["Dominated by sorting step"]

    C --> C1["Edge storage: O(E)"]
    C --> C2["Union-Find: O(V)"]
    C --> C3["MST storage: O(V)"]
    C --> C4["Total: O(E + V)"]

    D[Practical Notes] --> E["α(V) is inverse Ackermann"]
    D --> F["α(V) < 5 for practical inputs"]
    D --> G["Sorting is the bottleneck"]

    style A fill:#e1f5fe
    style B3 fill:#c8e6c9
    style F fill:#c8e6c9
```

### Kruskal's vs Prim's Algorithm

```mermaid
graph TD
    A[MST Algorithm Comparison] --> B[Kruskal's Algorithm]
    A --> C[Prim's Algorithm]

    B --> B1["Edge-based approach"]
    B --> B2["Global view of edges"]
    B --> B3["Better for sparse graphs"]
    B --> B4["Uses Union-Find"]
    B --> B5["O(E log E) time"]

    C --> C1["Vertex-based approach"]
    C --> C2["Local view from current tree"]
    C --> C3["Better for dense graphs"]
    C --> C4["Uses priority queue"]
    C --> C5["O(E log V) time"]

    D[Selection Guide] --> E["Sparse graphs: Kruskal's"]
    D --> F["Dense graphs: Prim's"]
    D --> G["Both find optimal MST"]

    style B3 fill:#c8e6c9
    style C3 fill:#c8e6c9
    style G fill:#c8e6c9
```

### Edge Sorting and Processing

```mermaid
graph LR
    subgraph "Edge List Representation"
        E1["(A,D,1)"]
        E2["(A,B,2)"]
        E3["(A,C,3)"]
        E4["(B,C,4)"]
        E5["(B,E,5)"]
        E6["(C,E,6)"]
        E7["(D,E,7)"]
    end

    subgraph "Processing Order"
        P1["1. Sort by weight"]
        P2["2. Process smallest first"]
        P3["3. Use Union-Find for cycle check"]
        P4["4. Add to MST if no cycle"]
    end

    subgraph "Result"
        R1["MST edges: (A,D), (A,B), (A,C), (B,E)"]
        R2["Total weight: 1+2+3+5 = 11"]
    end

    style P1 fill:#e1f5fe
    style R2 fill:#c8e6c9
```

### Cycle Detection Example

```mermaid
graph LR
    subgraph "Current MST State"
        A[A] ---|1| D[D]
        A ---|2| B[B]
        A ---|3| C[C]

        S["Components: {A,B,C,D}, {E}"]
    end

    subgraph "Consider Edge (B,C) weight 4"
        Q1["Find(B) = A (root of component)"]
        Q2["Find(C) = A (root of component)"]
        Q3["Since Find(B) == Find(C)"]
        Q4["Adding (B,C) would create cycle"]
        Q5["Skip this edge"]
    end

    style S fill:#e1f5fe
    style Q5 fill:#ffcdd2
```

### Union by Rank Optimization

```mermaid
graph TD
    A[Union by Rank] --> B[Concept]
    A --> C[Implementation]
    A --> D[Benefits]

    B --> B1["Attach smaller tree to larger tree"]
    B --> B2["Rank = height of tree"]
    B --> B3["Keeps trees balanced"]

    C --> C1["Track rank for each root"]
    C --> C2["Union smaller rank to larger"]
    C --> C3["Increment rank only when equal"]

    D --> D1["Limits tree height"]
    D --> D2["Improves Find performance"]
    D --> D3["O(log n) worst case height"]

    E[Alternative] --> F["Union by size"]
    E --> G["Track subtree size instead"]
    E --> H["Similar performance benefits"]

    style A fill:#e1f5fe
    style D2 fill:#c8e6c9
    style H fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[Kruskal's Algorithm Applications] --> B[Network Design]
    A --> C[Image Processing]
    A --> D[Cluster Analysis]
    A --> E[Approximation Algorithms]

    B --> B1["Computer network topology"]
    B --> B2["Telecommunication networks"]
    B --> B3["Transportation networks"]

    C --> C1["Image segmentation"]
    C --> C2["Feature extraction"]
    C --> C3["Computer vision"]

    D --> D1["Hierarchical clustering"]
    D --> D2["Data mining"]
    D --> D3["Social network analysis"]

    E --> E1["Traveling salesman problem"]
    E --> E2["Steiner tree approximation"]
    E --> E3["Network reliability"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Algorithm Variants and Optimizations

```mermaid
graph TD
    A[Kruskal's Variants] --> B[Standard Implementation]
    A --> C[Optimized Versions]
    A --> D[Parallel Versions]

    B --> B1["Sort all edges first"]
    B --> B2["Union-Find with optimizations"]
    B --> B3["Simple and reliable"]

    C --> C1["Filter obvious non-MST edges"]
    C --> C2["Borůvka-Kruskal hybrid"]
    C --> C3["Early termination"]

    D --> D1["Parallel sorting"]
    D --> D2["Concurrent Union-Find"]
    D --> D3["Lock-free implementations"]

    E[Performance Tips] --> F["Use efficient sorting"]
    E --> G["Implement path compression"]
    E --> H["Consider graph density"]

    style A fill:#e1f5fe
    style B3 fill:#c8e6c9
    style G fill:#c8e6c9
```

### Error Handling and Edge Cases

```mermaid
graph TD
    A[Edge Cases] --> B[Disconnected Graph]
    A --> C[Single Vertex]
    A --> D[No Edges]
    A --> E[Duplicate Edges]

    B --> B1["Returns forest of MSTs"]
    B --> B2["Each component gets MST"]

    C --> C1["MST is single vertex"]
    C --> C2["No edges in result"]

    D --> D1["Empty MST"]
    D --> D2["Each vertex separate component"]

    E --> E1["Keep minimum weight duplicate"]
    E --> E2["Or handle during preprocessing"]

    F[Implementation Checks] --> G["Validate input graph"]
    F --> H["Handle negative weights"]
    F --> I["Memory allocation errors"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style H fill:#c8e6c9
```

### Comparison with Other MST Algorithms

```mermaid
graph LR
    subgraph "Algorithm Characteristics"
        K["Kruskal's: O(E log E)"]
        P["Prim's: O(E log V)"]
        B["Borůvka's: O(E log V)"]
    end

    subgraph "Best Use Cases"
        K_Use["Kruskal's: Sparse graphs, parallel processing"]
        P_Use["Prim's: Dense graphs, streaming"]
        B_Use["Borůvka's: Parallel algorithms"]
    end

    subgraph "Data Structures"
        K_DS["Kruskal's: Union-Find"]
        P_DS["Prim's: Priority Queue"]
        B_DS["Borůvka's: Multiple techniques"]
    end

    style K_Use fill:#c8e6c9
    style K_DS fill:#c8e6c9
```

### Detailed Implementation Example

```mermaid
graph LR
    subgraph "Pseudocode Structure"
        A["function kruskal(graph):"]
        B["  edges = getAllEdges(graph)"]
        C["  sort(edges, by=weight)"]
        D["  uf = UnionFind(vertices)"]
        E["  mst = []"]
        F["  for edge in edges:"]
        G["    if not uf.connected(edge.u, edge.v):"]
        H["      mst.add(edge)"]
        I["      uf.union(edge.u, edge.v)"]
        J["  return mst"]
    end

    subgraph "Key Operations"
        Sort["O(E log E) sorting"]
        Find["O(α(V)) find operations"]
        Union["O(α(V)) union operations"]
        Total["Total: O(E log E)"]
    end

    style A fill:#e1f5fe
    style Total fill:#c8e6c9
```

Kruskal's Algorithm is a greedy algorithm that finds a minimum spanning tree (MST) for a weighted undirected graph.
