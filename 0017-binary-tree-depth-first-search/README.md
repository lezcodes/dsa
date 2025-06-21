# Binary Tree Depth-First Search

## Description

Binary Tree Depth-First Search (DFS) is a tree traversal algorithm that explores as far down each branch as possible before backtracking. It includes three main traversal orders: preorder, inorder, and postorder, each serving different purposes in tree processing.

## Visual Representation

### DFS Traversal Orders

```mermaid
graph TD
    A[1] --> B[2]
    A --> C[3]
    B --> D[4]
    B --> E[5]
    C --> F[6]
    C --> G[7]

    H["Preorder (NLR): 1, 2, 4, 5, 3, 6, 7"]
    I["Inorder (LNR): 4, 2, 5, 1, 6, 3, 7"]
    J["Postorder (LRN): 4, 5, 2, 6, 7, 3, 1"]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style I fill:#fff3e0
    style J fill:#f3e5f5
```

### DFS Algorithm Flow (Recursive)

```mermaid
graph TD
    A["dfs(node)"] --> B{"node == null?"}
    B -->|Yes| C["Return"]
    B -->|No| D{"Traversal Type?"}

    D -->|Preorder| E["Visit node"]
    E --> F["dfs(left)"]
    F --> G["dfs(right)"]

    D -->|Inorder| H["dfs(left)"]
    H --> I["Visit node"]
    I --> J["dfs(right)"]

    D -->|Postorder| K["dfs(left)"]
    K --> L["dfs(right)"]
    L --> M["Visit node"]

    style A fill:#e1f5fe
    style C fill:#c8e6c9
```

### Recursive Call Stack Visualization

```mermaid
graph LR
    subgraph "Call Stack for Preorder DFS"
        A["dfs(1)"] --> B["dfs(2)"]
        B --> C["dfs(4)"]
        C --> D["dfs(null)"]
    end

    subgraph "Stack State"
        E["Bottom: dfs(1)"]
        F["Middle: dfs(2)"]
        G["Top: dfs(4)"]
    end

    subgraph "Execution Order"
        H["1. Visit 1"]
        I["2. Visit 2"]
        J["3. Visit 4"]
        K["4. Backtrack to 2"]
        L["5. Visit 5"]
    end

    style A fill:#e1f5fe
    style D fill:#ffcdd2
    style L fill:#c8e6c9
```

### Iterative DFS Implementation

```mermaid
graph TD
    A[Iterative DFS] --> B[Create stack]
    B --> C[Push root to stack]
    C --> D{Stack empty?}
    D -->|Yes| E[Traversal complete]
    D -->|No| F[Pop node from stack]
    F --> G[Visit/Process node]
    G --> H[Push right child]
    H --> I[Push left child]
    I --> D

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Stack State During Iterative DFS

```mermaid
graph LR
    subgraph "Step 1: Initialize"
        S1["Stack: [1]"]
        V1["Visited: []"]
    end

    subgraph "Step 2: Process 1"
        S2["Stack: [3, 2]"]
        V2["Visited: [1]"]
    end

    subgraph "Step 3: Process 2"
        S3["Stack: [3, 5, 4]"]
        V3["Visited: [1, 2]"]
    end

    subgraph "Step 4: Process 4"
        S4["Stack: [3, 5]"]
        V4["Visited: [1, 2, 4]"]
    end

    style S1 fill:#e1f5fe
    style V4 fill:#c8e6c9
```

### DFS vs BFS Comparison

```mermaid
graph TD
    A[Tree Traversal Comparison] --> B[DFS]
    A --> C[BFS]

    B --> B1["Uses Stack (LIFO)"]
    B --> B2["Goes deep first"]
    B --> B3["Less memory usage"]
    B --> B4["Natural recursion"]
    B --> B5["Good for tree structure"]

    C --> C1["Uses Queue (FIFO)"]
    C --> C2["Goes wide first"]
    C --> C3["More memory usage"]
    C --> C4["Iterative approach"]
    C --> C5["Good for level-wise processing"]

    style B1 fill:#c8e6c9
    style C1 fill:#fff3e0
```

### Memory Usage Analysis

```mermaid
graph TD
    A[DFS Memory Analysis] --> B[Recursive Implementation]
    A --> C[Iterative Implementation]

    B --> B1["Stack space: O(h)"]
    B --> B2["Best case: O(log n) balanced tree"]
    B --> B3["Worst case: O(n) skewed tree"]

    C --> C1["Explicit stack: O(h)"]
    C --> C2["Better control over stack"]
    C --> C3["Avoid stack overflow"]

    D[Space Complexity Summary] --> E["Balanced tree: O(log n)"]
    D --> F["Skewed tree: O(n)"]
    D --> G["Complete tree: O(log n)"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style E fill:#c8e6c9
```

### DFS Applications

```mermaid
graph TD
    A[DFS Applications] --> B[Tree Operations]
    A --> C[Graph Problems]
    A --> D[Parsing & Evaluation]

    B --> B1["Tree serialization"]
    B --> B2["Directory traversal"]
    B --> B3["Expression trees"]
    B --> B4["Find tree properties"]

    C --> C1["Topological sorting"]
    C --> C2["Cycle detection"]
    C --> C3["Connected components"]
    C --> C4["Path finding"]

    D --> D1["Mathematical expressions"]
    D --> D2["Compiler design"]
    D --> D3["Syntax analysis"]
    D --> D4["Decision trees"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Search vs Traversal

```mermaid
graph LR
    subgraph "DFS Traversal"
        A["Visit all nodes"]
        B["Process entire tree"]
        C["Different visit orders"]
        D["Return path/values"]
    end

    subgraph "DFS Search"
        E["Find specific value"]
        F["Stop when found"]
        G["Return true/false"]
        H["Early termination"]
    end

    subgraph "Applications"
        I["Traversal: Tree printing"]
        J["Traversal: Serialization"]
        K["Search: Find node"]
        L["Search: Path exists"]
    end

    style A fill:#c8e6c9
    style E fill:#fff3e0
```

### DFS Path Tracking

```mermaid
graph TD
    A[Path Tracking in DFS] --> B[Current Path]
    A --> C[All Paths]
    A --> D[Path to Target]

    B --> B1["Maintain current path"]
    B --> B2["Add node on visit"]
    B --> B3["Remove on backtrack"]

    C --> C1["Store all root-to-leaf paths"]
    C --> C2["Collect complete paths"]
    C --> C3["Return list of paths"]

    D --> D1["Find path to specific node"]
    D --> D2["Return when target found"]
    D --> D3["Include parent information"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Implementation Variations

```mermaid
graph TD
    A[DFS Implementations] --> B[Recursive]
    A --> C[Iterative with Stack]
    A --> D[Morris Traversal]

    B --> B1["Simple and clean"]
    B --> B2["Natural for trees"]
    B --> B3["Stack overflow risk"]

    C --> C1["Explicit stack control"]
    C --> C2["No recursion limit"]
    C --> C3["More complex code"]

    D --> D1["O(1) space complexity"]
    D --> D2["Modifies tree temporarily"]
    D --> D3["Advanced technique"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#fff3e0
    style D1 fill:#f3e5f5
```

### Time Complexity Analysis

```mermaid
graph LR
    subgraph "DFS Time Complexity"
        A["Visit each node once: O(n)"]
        B["Process each edge once: O(n-1)"]
        C["Total time: O(n)"]
    end

    subgraph "Space Complexity"
        D["Recursive stack: O(h)"]
        E["Best case: O(log n)"]
        F["Worst case: O(n)"]
    end

    subgraph "Practical Considerations"
        G["Stack overflow in deep trees"]
        H["Memory-efficient traversal"]
        I["Early termination benefits"]
    end

    style C fill:#c8e6c9
    style F fill:#ffcdd2
    style I fill:#c8e6c9
```

Binary Tree Depth-First Search (DFS) is a tree traversal algorithm that explores as far down each branch as possible before backtracking.

## Features

- **Insert**: Add new values while maintaining BST property
- **Delete**: Remove values with proper tree restructuring
- **Search**: Boolean check for value existence
- **Find**: Return node reference for a given value
- **Traversals**: In-order, pre-order, and post-order depth-first traversals
- **Utility Methods**: Height calculation and size counting

## BST Properties

- Left subtree contains only values less than the parent node
- Right subtree contains only values greater than the parent node
- No duplicate values allowed
- In-order traversal produces sorted sequence

## Tree Structure Example

```
       50
      /  \
    30    70
   / \   / \
  20 40 60 80
```

## Traversal Methods

### In-Order (Left → Root → Right)

Produces sorted sequence: `[20, 30, 40, 50, 60, 70, 80]`

### Pre-Order (Root → Left → Right)

Useful for tree copying: `[50, 30, 20, 40, 70, 60, 80]`

### Post-Order (Left → Right → Root)

Useful for tree deletion: `[20, 40, 30, 60, 80, 70, 50]`

## Operations

### Insert

- **Time**: O(log n) average, O(n) worst case
- **Space**: O(log n) recursion stack
- Maintains BST property by comparing values

### Delete

- **Time**: O(log n) average, O(n) worst case
- **Space**: O(log n) recursion stack
- Handles three cases:
  - Leaf node: Simple removal
  - One child: Replace with child
  - Two children: Replace with in-order successor

### Search/Find

- **Time**: O(log n) average, O(n) worst case
- **Space**: O(log n) recursion stack
- Efficient binary search through tree structure

## Complexity

### Time Complexity

- **Insert**: O(log n) average, O(n) worst case
- **Delete**: O(log n) average, O(n) worst case
- **Search/Find**: O(log n) average, O(n) worst case
- **Traversals**: O(n) - visit each node once
- **Height**: O(n) - may visit all nodes
- **Size**: O(n) - visit each node once

### Space Complexity

- **Storage**: O(n) for n nodes
- **Operations**: O(log n) average recursion depth, O(n) worst case
- **Traversals**: O(n) for result storage + O(log n) recursion

## Use Cases

- **Sorted Data**: Maintaining sorted collections with dynamic updates
- **Range Queries**: Finding values within specific ranges
- **Database Indexing**: B-tree variants for database indexes
- **Expression Trees**: Parsing and evaluating mathematical expressions
- **File Systems**: Directory structure representation

## Usage

```bash
make run NAME=0017-binary-tree-depth-first
```

## Testing

```bash
make test NAME=0017-binary-tree-depth-first
```
