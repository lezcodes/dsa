# Binary Tree Breadth-First Search

## Description

Implementation of breadth-first search (BFS) algorithms for binary trees, also known as level-order traversal. This module includes various BFS-based tree algorithms that explore nodes level by level using a queue data structure.

BFS is fundamental for solving many tree problems including:

- **Level-order traversal**: Visit all nodes level by level from left to right
- **Level-by-level processing**: Group nodes by their depth in the tree
- **Tree views**: Get left-side or right-side view of the tree
- **Zigzag traversal**: Alternate direction for each level
- **Depth calculations**: Find minimum and maximum depth efficiently
- **Level-specific operations**: Calculate sums or perform operations on specific levels

## Tree Structure Example

```
      3
     / \
    9   20
   / \  / \
  1   2 15  7
```

**BFS Results:**

- **Traversal**: [3, 9, 20, 1, 2, 15, 7]
- **Levels**: [[3], [9, 20], [1, 2, 15, 7]]
- **Right View**: [3, 20, 7]
- **Left View**: [3, 9, 1]
- **Zigzag**: [[3], [20, 9], [1, 2, 15, 7]]

## Algorithms Implemented

### Core BFS Functions

- `BFS()`: Standard level-order traversal
- `BFSLevels()`: Group nodes by level
- `BFSRightSideView()`: Rightmost node at each level
- `BFSLeftSideView()`: Leftmost node at each level
- `BFSZigzag()`: Alternating left-to-right and right-to-left traversal

### Depth and Level Operations

- `MaxDepth()`: Maximum depth of the tree
- `MinDepth()`: Minimum depth to a leaf node
- `LevelSum()`: Sum of all nodes at a specific level

## Complexity

### Time Complexity

- **All BFS operations**: O(n) - visits each node exactly once
- **Space complexity varies by operation**

### Space Complexity

- **BFS traversal**: O(w) - where w is the maximum width of the tree
- **Level-based operations**: O(w) for queue + O(n) for result storage
- **Best case (balanced tree)**: O(n/2) ≈ O(n)
- **Worst case (complete tree)**: O(n)

## Implementation Details

### Queue-Based Approach

- Uses explicit queue (slice) for level-by-level processing
- FIFO (First In, First Out) ensures proper level ordering
- Level tracking enables grouped operations

### Key Techniques

- **Level size tracking**: Process one complete level at a time
- **Direction alternation**: For zigzag traversal using index manipulation
- **Early termination**: For minimum depth calculation
- **View algorithms**: Track first/last nodes per level

## Use Cases

### Level-Order Traversal

- Tree serialization and deserialization
- Printing tree structure level by level
- Building tree from level-order input

### Tree Views

- **Right Side View**: What you see from the right side
- **Left Side View**: What you see from the left side
- UI rendering and tree visualization

### Zigzag Traversal

- Spiral tree traversal
- Alternative tree printing formats
- Specific algorithmic challenges

### Depth Calculations

- **Max Depth**: Tree height, balancing checks
- **Min Depth**: Shortest path to leaf, optimization problems
- Tree analysis and validation

### Level Operations

- Level-wise processing and aggregation
- Tree statistics and analysis
- Conditional operations based on depth

## Performance Characteristics

BFS is optimal for:

- Finding shortest path to any node (unweighted)
- Level-wise processing requirements
- Tree width and depth analysis
- Problems requiring complete level exploration

BFS uses more memory than DFS but provides level-wise guarantees that DFS cannot offer.

## Usage

```bash
make run NAME=0015-binary-tree-breadth-first-search
```

## Testing

```bash
make test NAME=0015-binary-tree-breadth-first-search
```

## Benchmarking

```bash
go test -bench=. ./0015-binary-tree-breadth-first-search/
```

## Visual Representation

### BFS Traversal Order

```mermaid
graph TD
    A["1 (Level 0)"] --> B["2 (Level 1)"]
    A --> C["3 (Level 1)"]
    B --> D["4 (Level 2)"]
    B --> E["5 (Level 2)"]
    C --> F["6 (Level 2)"]
    C --> G["7 (Level 2)"]

    H["BFS Order: 1 → 2 → 3 → 4 → 5 → 6 → 7"]

    style A fill:#e1f5fe
    style B fill:#f3e5f5
    style C fill:#f3e5f5
    style D fill:#e8f5e8
    style E fill:#e8f5e8
    style F fill:#e8f5e8
    style G fill:#e8f5e8
    style H fill:#c8e6c9
```

### BFS Algorithm Flow

```mermaid
graph TD
    A[Start BFS] --> B[Create empty queue]
    B --> C[Add root to queue]
    C --> D{Queue empty?}
    D -->|Yes| E[Traversal complete]
    D -->|No| F[Dequeue front node]
    F --> G[Visit/Process node]
    G --> H{Has left child?}
    H -->|Yes| I[Enqueue left child]
    H -->|No| J{Has right child?}
    I --> J
    J -->|Yes| K[Enqueue right child]
    J -->|No| L[Continue to next iteration]
    K --> L
    L --> D

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Queue State During Traversal

```mermaid
graph LR
    subgraph "Step 1: Initialize"
        Q1["Queue: [1]"]
        V1["Visited: []"]
    end

    subgraph "Step 2: Process 1"
        Q2["Queue: [2, 3]"]
        V2["Visited: [1]"]
    end

    subgraph "Step 3: Process 2"
        Q3["Queue: [3, 4, 5]"]
        V3["Visited: [1, 2]"]
    end

    subgraph "Step 4: Process 3"
        Q4["Queue: [4, 5, 6, 7]"]
        V4["Visited: [1, 2, 3]"]
    end

    style Q1 fill:#e1f5fe
    style V4 fill:#c8e6c9
```

### Level-by-Level Processing

```mermaid
graph TD
    subgraph "Level 0"
        L0["Node: 1"]
    end

    subgraph "Level 1"
        L1A["Node: 2"]
        L1B["Node: 3"]
    end

    subgraph "Level 2"
        L2A["Node: 4"]
        L2B["Node: 5"]
        L2C["Node: 6"]
        L2D["Node: 7"]
    end

    A[Process by Level] --> B["Level 0: [1]"]
    B --> C["Level 1: [2, 3]"]
    C --> D["Level 2: [4, 5, 6, 7]"]

    style A fill:#e1f5fe
    style D fill:#c8e6c9
```

### BFS vs DFS Comparison

```mermaid
graph LR
    subgraph "BFS (Queue - FIFO)"
        A1["Visit by levels"]
        A2["Shortest path in unweighted trees"]
        A3["More memory usage"]
        A4["Iterative implementation"]
    end

    subgraph "DFS (Stack - LIFO)"
        B1["Visit depth-first"]
        B2["Natural recursion"]
        B3["Less memory usage"]
        B4["Goes deeper before wider"]
    end

    subgraph "Use Cases"
        C1["BFS: Level-order printing"]
        C2["BFS: Shortest path"]
        C3["DFS: Tree serialization"]
        C4["DFS: Expression evaluation"]
    end

    style A1 fill:#c8e6c9
    style C1 fill:#c8e6c9
```

### Memory Usage Analysis

```mermaid
graph TD
    A[BFS Memory Analysis] --> B[Queue Size]
    A --> C[Space Complexity]

    B --> B1["Best case: O(1) - skewed tree"]
    B --> B2["Worst case: O(n/2) - complete tree"]
    B --> B3["Average: O(w) where w = max width"]

    C --> C1["Space: O(w) for queue"]
    C --> C2["Time: O(n) visit each node once"]

    D[Tree Width Examples] --> E["Complete tree: width = n/2"]
    D --> F["Skewed tree: width = 1"]
    D --> G["Balanced tree: width ≈ 2^(h-1)"]

    style A fill:#e1f5fe
    style B2 fill:#ffcdd2
    style C2 fill:#c8e6c9
```

### Implementation Variations

```mermaid
graph TD
    A[BFS Implementations] --> B[Standard BFS]
    A --> C[Level-wise BFS]
    A --> D[Right-to-Left BFS]
    A --> E[Zigzag BFS]

    B --> B1["Single queue<br/>Visit all nodes<br/>Order: left to right"]

    C --> C1["Track level boundaries<br/>Group nodes by level<br/>Return 2D array"]

    D --> D1["Add right child first<br/>Then left child<br/>Reverse order"]

    E --> E1["Alternate direction<br/>Use two stacks/deque<br/>Level 0: L→R, Level 1: R→L"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
```

### Applications and Use Cases

```mermaid
graph TD
    A[BFS Applications] --> B[Tree Problems]
    A --> C[Graph Problems]
    A --> D[Practical Uses]

    B --> B1["Level-order printing<br/>Tree serialization<br/>Find tree width"]

    C --> C1["Shortest path in unweighted graphs<br/>Connected components<br/>Minimum spanning tree"]

    D --> D1["Web crawling<br/>Social network analysis<br/>GPS navigation"]

    E[Specific Tree Applications] --> F["Find leftmost node at each level"]
    E --> G["Print tree in level order"]
    E --> H["Check if tree is complete"]
    E --> I["Find maximum width of tree"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Pseudocode Visualization

```mermaid
graph TD
    A["function BFS(root)"] --> B["if root == null: return"]
    B --> C["queue = new Queue()"]
    C --> D["queue.enqueue(root)"]
    D --> E["while !queue.isEmpty()"]
    E --> F["node = queue.dequeue()"]
    F --> G["visit(node)"]
    G --> H["if node.left != null"]
    H --> I["queue.enqueue(node.left)"]
    I --> J["if node.right != null"]
    J --> K["queue.enqueue(node.right)"]
    K --> E
    E --> L["end"]

    style A fill:#e1f5fe
    style L fill:#c8e6c9
```

### Level-Order with Level Tracking

```mermaid
graph LR
    subgraph "Enhanced BFS"
        A["Add (node, level) to queue"]
        B["Track current level"]
        C["Group nodes by level"]
        D["Return list of lists"]
    end

    subgraph "Example Output"
        E["Level 0: [1]"]
        F["Level 1: [2, 3]"]
        G["Level 2: [4, 5, 6, 7]"]
        H["Result: [[1], [2,3], [4,5,6,7]]"]
    end

    style A fill:#e1f5fe
    style H fill:#c8e6c9
```

Binary Tree Breadth-First Search (BFS), also known as level-order traversal, visits nodes level by level from left to right.
