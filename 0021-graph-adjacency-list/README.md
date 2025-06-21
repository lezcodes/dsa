# Graph Adjacency List

## Description

A Graph Adjacency List represents a graph as an array of linked lists or dynamic arrays. Each vertex has a list of its adjacent vertices, making it memory-efficient for sparse graphs and providing fast neighbor iteration.

## Visual Representation

### Basic Graph and List Representation

```mermaid
graph LR
    subgraph "Sample Graph"
        A[0] --> B[1]
        A --> C[2]
        B --> C
        B --> D[3]
        C --> D
    end

    subgraph "Adjacency List"
        L0["0: [1, 2]"]
        L1["1: [2, 3]"]
        L2["2: [3]"]
        L3["3: []"]
    end

    style A fill:#e1f5fe
    style L0 fill:#c8e6c9
    style L1 fill:#c8e6c9
    style L2 fill:#c8e6c9
    style L3 fill:#c8e6c9
```

### Directed vs Undirected Representation

```mermaid
graph LR
    subgraph "Undirected Graph"
        A1[A] --- B1[B]
        A1 --- C1[C]
        B1 --- C1

        U_List["A: [B, C]<br/>B: [A, C]<br/>C: [A, B]"]
        U_Note["Each edge appears twice"]
    end

    subgraph "Directed Graph"
        A2[A] --> B2[B]
        A2 --> C2[C]
        B2 --> C2

        D_List["A: [B, C]<br/>B: [C]<br/>C: []"]
        D_Note["Each edge appears once"]
    end

    style U_List fill:#c8e6c9
    style D_List fill:#fff3e0
```

### Weighted Graph Representation

```mermaid
graph LR
    subgraph "Weighted Graph"
        A[A] -->|5| B[B]
        A -->|3| C[C]
        B -->|2| C
        B -->|7| D[D]
        C -->|1| D
    end

    subgraph "Weighted Adjacency List"
        W_List["A: [(B,5), (C,3)]<br/>B: [(C,2), (D,7)]<br/>C: [(D,1)]<br/>D: []"]
        W_Note["Store (neighbor, weight) pairs"]
    end

    style A fill:#e1f5fe
    style W_List fill:#c8e6c9
```

### List Operations

```mermaid
graph TD
    A[List Operations] --> B[Add Edge]
    A --> C[Remove Edge]
    A --> D[Check Edge]
    A --> E[Get Neighbors]

    B --> B1["Append to adjacency list"]
    B --> B2["For undirected: add both directions"]
    B --> B3["Time: O(1) for dynamic arrays"]

    C --> C1["Find and remove from list"]
    C --> C2["Time: O(degree) linear search"]
    C --> C3["Optimize with hash sets"]

    D --> D1["Search in adjacency list"]
    D --> D2["Time: O(degree) worst case"]
    D --> D3["O(1) with hash sets"]

    E --> E1["Return entire adjacency list"]
    E --> E2["Time: O(degree)"]

    style A fill:#e1f5fe
    style B3 fill:#c8e6c9
    style E2 fill:#c8e6c9
```

### Memory Layout and Structure

```mermaid
graph TD
    A[Adjacency List Structure] --> B[Array of Lists]
    A --> C[Array of Sets]
    A --> D[Hash Map Implementation]

    B --> B1["Array indices = vertices"]
    B --> B2["Each element = linked list"]
    B --> B3["Memory: O(V + E)"]

    C --> C1["Fast edge checking"]
    C --> C2["No duplicate edges"]
    C --> C3["Higher memory overhead"]

    D --> D1["Dynamic vertex addition"]
    D --> D2["String vertex names"]
    D --> D3["Flexible implementation"]

    style A fill:#e1f5fe
    style B3 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Space and Time Complexity

```mermaid
graph TD
    A[Adjacency List Complexity] --> B[Space Complexity]
    A --> C[Time Complexity]

    B --> B1["O(V + E) total space"]
    B --> B2["Efficient for sparse graphs"]
    B --> B3["Each edge stored once (directed)"]

    C --> C1["Add Edge: O(1) append"]
    C --> C2["Remove Edge: O(degree)"]
    C --> C3["Check Edge: O(degree)"]
    C --> C4["Get Neighbors: O(degree)"]
    C --> C5["DFS/BFS: O(V + E)"]

    D[Optimization] --> E["Use hash sets for O(1) operations"]
    D --> F["Maintain sorted lists for binary search"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style C5 fill:#c8e6c9
```

### Graph Traversal with Adjacency Lists

```mermaid
graph LR
    subgraph "DFS Traversal"
        A["For each neighbor in adj[v]:"]
        B["  if not visited[neighbor]:"]
        C["    dfs(neighbor)"]
    end

    subgraph "BFS Traversal"
        D["queue.add(start)"]
        E["while queue not empty:"]
        F["  v = queue.pop()"]
        G["  for neighbor in adj[v]:"]
        H["    if not visited: queue.add(neighbor)"]
    end

    subgraph "Efficiency"
        I["Natural iteration over neighbors"]
        J["No need to check all vertices"]
        K["Optimal O(V + E) complexity"]
    end

    style A fill:#e1f5fe
    style I fill:#c8e6c9
    style K fill:#c8e6c9
```

### List vs Matrix Comparison

```mermaid
graph TD
    A[Representation Comparison] --> B[Adjacency List]
    A --> C[Adjacency Matrix]

    B --> B1["Space: O(V + E)"]
    B --> B2["Sparse graph friendly"]
    B --> B3["Edge check: O(degree)"]
    B --> B4["Neighbor iteration: O(degree)"]
    B --> B5["Dynamic graph support"]

    C --> C1["Space: O(V²)"]
    C --> C2["Dense graph suitable"]
    C --> C3["Edge check: O(1)"]
    C --> C4["Neighbor iteration: O(V)"]
    C --> C5["Fixed vertex count"]

    D[Best Use Cases] --> E["Lists: Sparse graphs, traversals"]
    D --> F["Matrix: Dense graphs, edge queries"]

    style B1 fill:#c8e6c9
    style B4 fill:#c8e6c9
    style C3 fill:#fff3e0
```

### Implementation Variants

```mermaid
graph TD
    A[Implementation Options] --> B[Data Structure Choice]
    A --> C[Storage Strategy]
    A --> D[Access Patterns]

    B --> B1["Dynamic Arrays (vectors)"]
    B --> B2["Linked Lists"]
    B --> B3["Hash Sets"]
    B --> B4["Sorted Arrays"]

    C --> C1["Edge objects with weights"]
    C --> C2["Simple neighbor indices"]
    C --> C3["Tuple/pair representations"]

    D --> D1["Sequential access: arrays"]
    D --> D2["Frequent lookups: sets"]
    D --> D3["Range queries: sorted structures"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style B3 fill:#c8e6c9
    style D2 fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[Adjacency List Applications] --> B[Social Networks]
    A --> C[Web Graphs]
    A --> D[Biological Networks]
    A --> E[Transportation]

    B --> B1["Friend connections"]
    B --> B2["Follow relationships"]
    B --> B3["Recommendation systems"]

    C --> C1["Page links"]
    C --> C2["Web crawling"]
    C --> C3["PageRank algorithm"]

    D --> D1["Protein interactions"]
    D --> D2["Gene networks"]
    D --> D3["Neural connections"]

    E --> E1["Route planning"]
    E --> E2["Public transit"]
    E --> E3["Flight networks"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Memory Optimization Techniques

```mermaid
graph LR
    subgraph "Memory Optimizations"
        A["Compressed representations"]
        B["Bit-packed edge lists"]
        C["Offset arrays for CSR format"]
    end

    subgraph "Access Optimizations"
        D["Cache-friendly layouts"]
        E["Locality-preserving ordering"]
        F["Block-based storage"]
    end

    subgraph "Dynamic Optimizations"
        G["Incremental updates"]
        H["Batch modifications"]
        I["Lazy deletions"]
    end

    style A fill:#e1f5fe
    style D fill:#c8e6c9
    style G fill:#c8e6c9
```

### Common Algorithms Implementation

```mermaid
graph TD
    A[Graph Algorithms] --> B[Shortest Path]
    A --> C[Connectivity]
    A --> D[Topological Sort]

    B --> B1["Dijkstra with priority queue"]
    B --> B2["Bellman-Ford with edge relaxation"]
    B --> B3["BFS for unweighted graphs"]

    C --> C1["DFS for connected components"]
    C --> C2["Union-Find with adjacency lists"]
    C --> C3["Strongly connected components"]

    D --> D1["Kahn's algorithm with in-degrees"]
    D --> D2["DFS-based ordering"]

    E[Benefits] --> F["Natural edge iteration"]
    E --> G["Optimal space usage"]
    E --> H["Dynamic graph support"]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style G fill:#c8e6c9
    style H fill:#c8e6c9
```

A Graph Adjacency List represents a graph as an array of linked lists or dynamic arrays.

### Features

- **Graph Types**: Directed and undirected graphs
- **Core Operations**: Add/remove edges, check edge existence, get neighbors
- **Traversal Algorithms**:
  - Breadth-First Search (BFS) - iterative implementation
  - Depth-First Search (DFS) - recursive implementation
  - DFS Iterative - stack-based implementation
- **Graph Analysis**: Connectivity checking, cycle detection, topological sorting
- **Utility Methods**: Print graph structure, vertex/edge counting, neighbor enumeration
- **Advanced Features**: Topological sort for DAGs (Directed Acyclic Graphs)

### Data Structure

- **Adjacency List**: Map where each vertex maps to a slice of its neighbors
- **Space Efficient**: For sparse graphs (few edges relative to vertices)
- **Dynamic Storage**: Only stores actual edges, no wasted space for non-existent edges

## Complexity

### Time Complexity

- **Add Edge**: O(1) average case
- **Remove Edge**: O(degree) where degree is number of neighbors
- **Has Edge**: O(degree) where degree is number of neighbors
- **Get Neighbors**: O(1) - direct map lookup
- **BFS Traversal**: O(V + E) where V is vertices, E is edges
- **DFS Traversal**: O(V + E) where V is vertices, E is edges
- **Is Connected**: O(V + E) - performs BFS traversal
- **Has Cycle**: O(V + E) - performs DFS with recursion stack tracking
- **Topological Sort**: O(V + E) - DFS-based algorithm

### Space Complexity

- **Graph Storage**: O(V + E) - vertices plus edges
- **BFS**: O(V) - queue and visited array
- **DFS**: O(V) - recursion stack and visited array
- **Topological Sort**: O(V) - visited array and result stack
- **Overall**: O(V + E) - optimal for sparse graphs

## Usage

```bash
make run NAME=0021-graph-adjacency-list
```

## Testing

```bash
make test NAME=0021-graph-adjacency-list
```
