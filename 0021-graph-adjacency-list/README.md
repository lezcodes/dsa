# Graph Adjacency List

## Description

A graph data structure implementation using an adjacency list representation. Supports both directed and undirected graphs with comprehensive traversal algorithms (BFS and DFS) and advanced graph operations.

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
