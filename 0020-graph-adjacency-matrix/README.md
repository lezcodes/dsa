# Graph Adjacency Matrix

## Description

A Graph Adjacency Matrix is a 2D array representation of a graph where matrix[i][j] indicates whether there is an edge between vertex i and vertex j. This representation is efficient for dense graphs and provides O(1) edge lookup time.

## Visual Representation

### Basic Graph and Matrix Representation

```mermaid
graph LR
    subgraph "Sample Graph"
        A[0] --> B[1]
        A --> C[2]
        B --> C
        B --> D[3]
        C --> D
    end

    subgraph "Adjacency Matrix"
        M["  0 1 2 3<br/>0 [0 1 1 0]<br/>1 [0 0 1 1]<br/>2 [0 0 0 1]<br/>3 [0 0 0 0]"]
    end

    style A fill:#e1f5fe
    style M fill:#c8e6c9
```

### Directed vs Undirected Graphs

```mermaid
graph LR
    subgraph "Undirected Graph"
        A1[A] --- B1[B]
        A1 --- C1[C]
        B1 --- C1

        M1["  A B C<br/>A [0 1 1]<br/>B [1 0 1]<br/>C [1 1 0]"]
        M1_Note["Symmetric Matrix"]
    end

    subgraph "Directed Graph"
        A2[A] --> B2[B]
        A2 --> C2[C]
        B2 --> C2

        M2["  A B C<br/>A [0 1 1]<br/>B [0 0 1]<br/>C [0 0 0]"]
        M2_Note["Asymmetric Matrix"]
    end

    style M1 fill:#c8e6c9
    style M2 fill:#fff3e0
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

    subgraph "Weighted Adjacency Matrix"
        W["  A B C D<br/>A [0 5 3 ∞]<br/>B [∞ 0 2 7]<br/>C [∞ ∞ 0 1]<br/>D [∞ ∞ ∞ 0]"]
        WNote["∞ = no edge, numbers = weights"]
    end

    style A fill:#e1f5fe
    style W fill:#c8e6c9
```

### Matrix Operations

```mermaid
graph TD
    A[Matrix Operations] --> B[Add Edge]
    A --> C[Remove Edge]
    A --> D[Check Edge]
    A --> E[Get Neighbors]

    B --> B1["matrix[i][j] = 1 (or weight)"]
    B --> B2["For undirected: matrix[j][i] = 1"]

    C --> C1["matrix[i][j] = 0"]
    C --> C2["For undirected: matrix[j][i] = 0"]

    D --> D1["return matrix[i][j] != 0"]

    E --> E1["scan row i for non-zero values"]
    E --> E2["return list of connected vertices"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Space and Time Complexity

```mermaid
graph TD
    A[Adjacency Matrix Complexity] --> B[Space Complexity]
    A --> C[Time Complexity]

    B --> B1["Always O(V²)"]
    B --> B2["Independent of edge count"]
    B --> B3["Memory intensive for sparse graphs"]

    C --> C1["Add/Remove Edge: O(1)"]
    C --> C2["Check Edge: O(1)"]
    C --> C3["Get All Neighbors: O(V)"]
    C --> C4["Get All Edges: O(V²)"]

    D[Efficiency Analysis] --> E["Dense graphs: Efficient"]
    D --> F["Sparse graphs: Wasteful"]
    D --> G["Fast edge queries"]

    style A fill:#e1f5fe
    style C1 fill:#c8e6c9
    style C2 fill:#c8e6c9
    style B3 fill:#ffcdd2
```

### Matrix Properties

```mermaid
graph LR
    subgraph "Matrix Properties"
        A["Diagonal Elements"]
        B["Self-loops: matrix[i][i] = 1"]
        C["No self-loops: matrix[i][i] = 0"]
    end

    subgraph "Symmetry"
        D["Undirected: matrix[i][j] = matrix[j][i]"]
        E["Directed: matrix[i][j] ≠ matrix[j][i]"]
    end

    subgraph "Sparsity"
        F["Dense: Many 1s"]
        G["Sparse: Many 0s"]
        H["Density = |E| / |V|²"]
    end

    style B fill:#c8e6c9
    style D fill:#c8e6c9
    style H fill:#fff3e0
```

### Common Graph Algorithms with Matrix

```mermaid
graph TD
    A[Graph Algorithms] --> B[Traversal]
    A --> C[Shortest Path]
    A --> D[Connectivity]

    B --> B1["DFS: O(V²) to check all edges"]
    B --> B2["BFS: O(V²) scanning matrix"]

    C --> C1["Floyd-Warshall: O(V³)"]
    C --> C2["Matrix multiplication approach"]

    D --> D1["Check connected components"]
    D --> D2["Transitive closure"]

    E[Matrix Powers] --> F["A² gives 2-hop paths"]
    E --> G["A^k gives k-hop paths"]

    style A fill:#e1f5fe
    style C1 fill:#c8e6c9
    style F fill:#c8e6c9
```

### Matrix vs List Comparison

```mermaid
graph TD
    A[Representation Comparison] --> B[Adjacency Matrix]
    A --> C[Adjacency List]

    B --> B1["Space: O(V²)"]
    B --> B2["Edge check: O(1)"]
    B --> B3["Add/Remove: O(1)"]
    B --> B4["Memory: High for sparse"]

    C --> C1["Space: O(V + E)"]
    C --> C2["Edge check: O(degree)"]
    C --> C3["Add/Remove: O(1)"]
    C --> C4["Memory: Efficient"]

    D[When to Use Matrix] --> E["Dense graphs"]
    D --> F["Frequent edge queries"]
    D --> G["Matrix operations needed"]

    style B2 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style E fill:#c8e6c9
```

### Implementation Details

```mermaid
graph TD
    A[Implementation Considerations] --> B[Data Types]
    A --> C[Memory Layout]
    A --> D[Optimization]

    B --> B1["Boolean: 0/1 for unweighted"]
    B --> B2["Integer: weights for weighted"]
    B --> B3["Float: fractional weights"]

    C --> C1["Row-major order"]
    C --> C2["Cache-friendly access"]
    C --> C3["Contiguous memory"]

    D --> D1["Bit vectors for boolean"]
    D --> D2["Sparse matrix formats"]
    D --> D3["Symmetric storage"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C2 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[Adjacency Matrix Applications] --> B[Social Networks]
    A --> C[Computer Networks]
    A --> D[Transportation]
    A --> E[Game Development]

    B --> B1["Friendship matrices"]
    B --> B2["Influence networks"]
    B --> B3["Community detection"]

    C --> C1["Network topology"]
    C --> C2["Routing tables"]
    C --> C3["Connectivity matrices"]

    D --> D1["Road networks"]
    D --> D2["Flight connections"]
    D --> D3["Distance matrices"]

    E --> E1["Game state spaces"]
    E --> E2["AI pathfinding"]
    E --> E3["Strategy games"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Matrix Initialization and Updates

```mermaid
graph LR
    subgraph "Initialization"
        A["Create V×V matrix"]
        B["Fill with zeros"]
        C["Set diagonal if needed"]
    end

    subgraph "Adding Edges"
        D["matrix[from][to] = weight"]
        E["If undirected: matrix[to][from] = weight"]
    end

    subgraph "Removing Edges"
        F["matrix[from][to] = 0"]
        G["If undirected: matrix[to][from] = 0"]
    end

    style A fill:#e1f5fe
    style D fill:#c8e6c9
    style F fill:#fff3e0
```

A Graph Adjacency Matrix is a 2D array representation of a graph where matrix[i][j] indicates whether there is an edge between vertex i and vertex j.

### Features

- **Graph Types**: Directed and undirected graphs
- **Core Operations**: Add/remove edges, check edge existence, get neighbors
- **Traversal Algorithms**:
  - Breadth-First Search (BFS) - iterative implementation
  - Depth-First Search (DFS) - recursive implementation
  - DFS Iterative - stack-based implementation
- **Graph Analysis**: Connectivity checking, cycle detection
- **Utility Methods**: Print graph structure, neighbor enumeration

### Data Structure

- **Adjacency Matrix**: 2D array where `matrix[i][j] = 1` indicates an edge from vertex i to vertex j
- **Space Efficient**: For dense graphs (many edges relative to vertices)
- **Fast Edge Lookup**: O(1) time complexity for checking if edge exists

## Complexity

### Time Complexity

- **Add Edge**: O(1)
- **Remove Edge**: O(1)
- **Has Edge**: O(1)
- **Get Neighbors**: O(V) where V is number of vertices
- **BFS Traversal**: O(V²) - visits each vertex once, checks all V neighbors for each
- **DFS Traversal**: O(V²) - visits each vertex once, checks all V neighbors for each
- **Is Connected**: O(V²) - performs BFS traversal
- **Has Cycle**: O(V²) - performs DFS with recursion stack tracking

### Space Complexity

- **Graph Storage**: O(V²) - adjacency matrix requires V×V space
- **BFS**: O(V) - queue and visited array
- **DFS**: O(V) - recursion stack and visited array
- **Overall**: O(V²) - dominated by matrix storage

## Usage

```bash
make run NAME=0020-graph-adjacency-matrix
```

## Testing

```bash
make test NAME=0020-graph-adjacency-matrix
```
