# Dijkstra's Algorithm

## Description

Dijkstra's Algorithm is a graph search algorithm that finds the shortest path between nodes in a weighted graph with non-negative edge weights. It's widely used in network routing protocols, GPS navigation systems, and social networking analysis.

## Visual Representation

### Graph Example

```mermaid
graph LR
    A[A] -->|4| B[B]
    A -->|2| C[C]
    B -->|1| C[C]
    B -->|5| D[D]
    C -->|8| D[D]
    C -->|10| E[E]
    D -->|2| E[E]
    D -->|6| F[F]
    E -->|3| F[F]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
```

### Algorithm Step-by-Step Execution

```mermaid
graph TD
    A[Initialize] --> B[Set source distance to 0]
    B --> C[Add all vertices to priority queue]
    C --> D[Extract minimum distance vertex]
    D --> E[Update distances to neighbors]
    E --> F{Queue empty?}
    F -->|No| D
    F -->|Yes| G[Algorithm Complete]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
```

### Dijkstra's Execution Example

```mermaid
graph LR
    subgraph "Step 1: Initialize (Source: A)"
        A1[A:0] --> B1[B:∞]
        A1 --> C1[C:∞]
        B1 --> D1[D:∞]
        C1 --> D1
        D1 --> E1[E:∞]
        D1 --> F1[F:∞]
        E1 --> F1
    end

    subgraph "Step 2: Process A"
        A2[A:0✓] -->|4| B2[B:4]
        A2 -->|2| C2[C:2]
        B2 --> D2[D:∞]
        C2 --> D2
        D2 --> E2[E:∞]
        D2 --> F2[F:∞]
        E2 --> F2
    end

    subgraph "Step 3: Process C (min distance)"
        A3[A:0✓] --> B3[B:3]
        A3 --> C3[C:2✓]
        B3 --> D3[D:10]
        C3 --> D3
        D3 --> E3[E:12]
        D3 --> F3[F:∞]
        E3 --> F3
    end

    style A1 fill:#e1f5fe
    style A2 fill:#c8e6c9
    style C3 fill:#c8e6c9
```

### Priority Queue Operations

```mermaid
graph TD
    A[Priority Queue in Dijkstra] --> B[Min-Heap Implementation]
    B --> C[Extract-Min: O(log V)]
    B --> D[Decrease-Key: O(log V)]

    E[Alternative: Simple Array] --> F[Extract-Min: O(V)]
    E --> G[Decrease-Key: O(1)]

    H[Total Complexity] --> I["Heap: O((V + E) log V)"]
    H --> J["Array: O(V²)"]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
    style J fill:#fff3e0
```

### Distance Relaxation Process

```mermaid
graph LR
    subgraph "Current State"
        A[Source] -->|5| B[Current: dist=8]
        A -->|3| C[Via: dist=3]
        C -->|2| B
    end

    subgraph "Relaxation Check"
        D["dist[B] = 8"]
        E["dist[C] + weight(C,B) = 3 + 2 = 5"]
        F["5 < 8? YES → Update dist[B] = 5"]
    end

    subgraph "After Relaxation"
        G[Source] -->|3| H[Via: dist=3]
        H -->|2| I[Updated: dist=5]
    end

    style F fill:#c8e6c9
    style I fill:#c8e6c9
```

### Path Reconstruction

```mermaid
graph TD
    A[Dijkstra Complete] --> B[Backtrack from destination]
    B --> C[Follow parent pointers]
    C --> D[Build path in reverse]
    D --> E[Reverse to get final path]

    F[Example: A to F] --> G["F ← E ← D ← B ← A"]
    G --> H["Path: A → B → D → E → F"]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
```

### Dijkstra vs Other Algorithms

```mermaid
graph TD
    A[Shortest Path Algorithms] --> B[Dijkstra]
    A --> C[Bellman-Ford]
    A --> D[Floyd-Warshall]
    A --> E[A*]

    B --> B1["Single-source<br/>Non-negative weights<br/>O((V+E)log V)"]
    C --> C1["Single-source<br/>Negative weights OK<br/>O(VE)"]
    D --> D1["All-pairs<br/>Any weights<br/>O(V³)"]
    E --> E1["Single-pair<br/>Heuristic-guided<br/>Often faster in practice"]

    style B1 fill:#c8e6c9
    style A fill:#e1f5fe
```

### Algorithm Limitations

```mermaid
graph LR
    subgraph "Dijkstra Limitations"
        A[Negative Weights] --> A1["Cannot handle<br/>negative edge weights"]
        B[Single Source] --> B1["Only one source<br/>at a time"]
        C[Memory Usage] --> C1["O(V) space for<br/>distances and queue"]
    end

    subgraph "When to Use Alternatives"
        D[Negative weights] --> E[Use Bellman-Ford]
        F[All pairs] --> G[Use Floyd-Warshall]
        H[Known target] --> I[Use A* with heuristic]
    end

    style A1 fill:#ffcdd2
    style E fill:#c8e6c9
    style G fill:#c8e6c9
    style I fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[Dijkstra Applications] --> B[GPS Navigation]
    A --> C[Network Routing]
    A --> D[Social Networks]
    A --> E[Game Development]
    A --> F[Transportation]

    B --> B1["Find shortest route<br/>Traffic-aware navigation"]
    C --> C1["OSPF protocol<br/>Internet packet routing"]
    D --> D1["Friend suggestions<br/>Shortest connection path"]
    E --> E1["Pathfinding in games<br/>AI movement"]
    F --> F1["Flight connections<br/>Public transit routing"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Pseudocode Visualization

```mermaid
graph TD
    A["function dijkstra(graph, source)"] --> B["dist[source] = 0"]
    B --> C["for each vertex v ≠ source"]
    C --> D["dist[v] = ∞"]
    D --> E["add all vertices to Q"]
    E --> F["while Q is not empty"]
    F --> G["u = extract_min(Q)"]
    G --> H["for each neighbor v of u"]
    H --> I["alt = dist[u] + weight(u,v)"]
    I --> J{"alt < dist[v]?"}
    J -->|Yes| K["dist[v] = alt"]
    J -->|No| L["continue"]
    K --> L
    L --> M["next neighbor"]
    M --> H
    H --> N["next vertex"]
    N --> F
    F --> O["return dist"]

    style A fill:#e1f5fe
    style O fill:#c8e6c9
```

Dijkstra's Algorithm is a graph search algorithm that finds the shortest path between nodes in a weighted graph with non-negative edge weights.

## Implementation Details

### Data Structures

- **Graph**: Adjacency list representation with weighted edges
- **Priority Queue**: Min-heap for efficient minimum distance vertex selection
- **Result**: Contains distances array, predecessor array, and source vertex

### Core Functions

- `NewGraph(vertices)`: Creates a new graph with specified number of vertices
- `AddEdge(from, to, weight)`: Adds a directed weighted edge
- `AddBidirectionalEdge(u, v, weight)`: Adds edges in both directions
- `Dijkstra(source)`: Executes the algorithm from given source
- `GetPath(target)`: Reconstructs shortest path to target vertex
- `GetDistance(target)`: Returns shortest distance to target vertex
- `HasPath(target)`: Checks if target is reachable from source

## Complexity

- **Time Complexity**: O((V + E) log V) where V is vertices and E is edges
  - Each vertex is extracted from priority queue once: O(V log V)
  - Each edge is relaxed at most once: O(E log V)
- **Space Complexity**: O(V + E)
  - Adjacency list: O(V + E)
  - Distance and predecessor arrays: O(V)
  - Priority queue: O(V)

## Algorithm Steps

1. Initialize distances to all vertices as infinite, except source (distance 0)
2. Add source vertex to priority queue with distance 0
3. While priority queue is not empty:
   - Extract vertex with minimum distance
   - Mark as visited
   - For each unvisited neighbor:
     - Calculate new distance through current vertex
     - If new distance is shorter, update distance and predecessor
     - Add neighbor to priority queue with new distance
4. Return distances and predecessor arrays

## Usage

```bash
make run n=dijkstra-algorithm
```

### Example Output

```
Result: map[
  distances:[0 4 3 6 8 14]
  paths:map[
    1:[0 2 1]
    2:[0 2]
    3:[0 2 1 3]
    4:[0 2 1 3 4]
    5:[0 2 1 3 4 5]
  ]
  source:0
]
```

## Testing

```bash
make test n=dijkstra-algorithm
```

### Test Coverage

- Graph construction and edge addition
- Basic shortest path computation
- Disconnected graph handling
- Invalid source vertex handling
- Path reconstruction for reachable and unreachable vertices
- Distance queries and path existence checks
- Complex graph scenarios
- Edge cases (single vertex, linear chains, etc.)

## Benchmarking

```bash
make bench n=dijkstra-algorithm
```

### Benchmark Scenarios

- Small graphs (10 vertices)
- Medium graphs (100 vertices)
- Large graphs (1000 vertices)
- Dense graphs (complete graphs)
- Path reconstruction operations

## Applications

- **GPS Navigation**: Finding shortest routes between locations
- **Network Routing**: Optimal packet routing in computer networks
- **Social Networks**: Finding shortest connections between people
- **Game Development**: Pathfinding for NPCs and game mechanics
- **Airline Routes**: Finding cheapest flight connections
- **Supply Chain**: Optimizing delivery routes and costs

## Limitations

- **Non-negative weights only**: Cannot handle negative edge weights
- **Single source**: Must run multiple times for all-pairs shortest paths
- **Memory intensive**: Requires O(V²) space for dense graphs
- **Not suitable for dynamic graphs**: Recalculation needed when graph changes

## Variations

- **Bidirectional Dijkstra**: Search from both source and target
- **A\* Algorithm**: Uses heuristics for faster pathfinding
- **Johnson's Algorithm**: Handles negative weights by preprocessing
- **Dial's Algorithm**: Optimization for small integer weights
