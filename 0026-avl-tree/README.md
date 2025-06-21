# avl-tree

## Description

AVL Tree (Adelson-Velsky and Landis Tree) is a self-balancing binary search tree where the heights of the two child subtrees of any node differ by at most one. This implementation provides both recursive and iterative approaches for key operations.

### Key Features

- **Self-balancing**: Automatically maintains balance through rotations
- **Height-balanced**: Height difference between left and right subtrees ≤ 1
- **Dual implementation**: Both recursive (default) and iterative methods
- **Complete operations**: Insert, delete, search, traversals, and utility functions
- **Rotation support**: Left, right, left-right, and right-left rotations

### AVL Tree Properties

1. **Binary Search Tree Property**: Left child < parent < right child
2. **Balance Factor**: For any node, |height(left) - height(right)| ≤ 1
3. **Height Tracking**: Each node stores its height for efficient balance calculations
4. **Automatic Rebalancing**: Rotations performed during insertion/deletion to maintain balance

### Rotation Types

- **Left Rotation (LL)**: When right subtree is heavier
- **Right Rotation (RR)**: When left subtree is heavier
- **Left-Right Rotation (LR)**: Left child's right subtree is heavier
- **Right-Left Rotation (RL)**: Right child's left subtree is heavier

## Complexity

- **Time Complexity**:

  - Insert: O(log n) - guaranteed due to balanced height
  - Delete: O(log n) - includes rebalancing operations
  - Search: O(log n) - both recursive and iterative
  - Traversals: O(n) - visit each node once
  - Height: O(1) - stored in each node
  - Balance Check: O(n) - validates entire tree

- **Space Complexity**:
  - Storage: O(n) - one node per element
  - Recursive operations: O(log n) - call stack depth
  - Iterative operations: O(log n) - explicit stack for rebalancing

## Implementation Details

### Core Methods

- `Insert(value)` - Recursive insertion with automatic rebalancing
- `InsertIterative(value)` - Iterative insertion alternative
- `Delete(value)` - Recursive deletion with rebalancing
- `Search(value)` - Recursive search
- `SearchIterative(value)` - Iterative search alternative

### Traversal Methods

- `InOrderTraversal()` - Returns sorted array of values
- `PreOrderTraversal()` - Root-first traversal
- `LevelOrderTraversal()` - Breadth-first traversal

### Utility Methods

- `GetHeight()` - Returns tree height
- `IsBalanced()` - Validates AVL properties
- `GetSize()` - Returns number of nodes
- `FindMin()/FindMax()` - Find minimum/maximum values
- `Clear()` - Removes all nodes
- `PrintTree()` - Visual tree representation with heights and balance factors

### Balance Factor Calculation

```
Balance Factor = Height(Left Subtree) - Height(Right Subtree)
```

- Balance Factor > 1: Left-heavy (needs right rotation)
- Balance Factor < -1: Right-heavy (needs left rotation)
- Balance Factor ∈ [-1, 0, 1]: Balanced

## Usage

```bash
make run n=0026-avl-tree
```

### Example Operations

```go
avl := NewAVLTree()

// Insert values (triggers automatic rebalancing)
values := []int{10, 20, 30, 40, 50, 25}
for _, value := range values {
    avl.Insert(value)
}

// Search operations
found := avl.Search(30)              // true
foundIter := avl.SearchIterative(30) // true

// Tree properties
size := avl.GetSize()        // 6
height := avl.GetHeight()    // Balanced height
balanced := avl.IsBalanced() // true

// Traversals
inOrder := avl.InOrderTraversal()    // [10, 20, 25, 30, 40, 50]
preOrder := avl.PreOrderTraversal()  // Root-first order
levelOrder := avl.LevelOrderTraversal() // Level by level

// Min/Max operations
min, hasMin := avl.FindMin() // 10, true
max, hasMax := avl.FindMax() // 50, true

// Deletion (maintains balance)
deleted := avl.Delete(20) // true
newSize := avl.GetSize()  // 5

// Iterative insertion
avl.InsertIterative(15)
```

## Testing

```bash
make test n=0026-avl-tree
```

### Test Coverage

- **Basic Operations**: Insert, delete, search functionality
- **Balance Validation**: All four rotation types tested
- **Edge Cases**: Empty tree, single node, duplicates
- **Traversals**: All traversal methods validated
- **Large Datasets**: 1000+ node stress testing
- **Performance**: Benchmarks for all major operations
- **Iterative Methods**: Alternative implementations tested

## Real-World Applications

- **Database Indexing**: Maintaining sorted indexes with frequent updates
- **Memory Management**: Balanced allocation trees
- **Compiler Design**: Symbol tables requiring fast lookups
- **Graphics**: Spatial partitioning with balanced access
- **File Systems**: Directory structures with balanced access times

## Advantages over Regular BST

- **Guaranteed Performance**: O(log n) operations even with sorted input
- **Predictable Height**: Maximum height is 1.44 \* log₂(n)
- **No Degeneration**: Cannot become a linear chain like unbalanced BST
- **Consistent Performance**: Uniform operation times regardless of insertion order

## Comparison with Other Trees

- **vs Red-Black Tree**: Stricter balancing, faster lookups, slower insertions/deletions
- **vs B-Tree**: Better for in-memory operations, simpler implementation
- **vs Splay Tree**: More predictable performance, no amortized analysis needed

## Visual Representation

### AVL Tree Balance Property

```mermaid
graph TD
    A[10] --> B[5]
    A --> C[15]
    B --> D[2]
    B --> E[7]
    C --> F[12]
    C --> G[20]

    H["Balance Factor = height(left) - height(right)"]
    I["Valid AVL: BF ∈ {-1, 0, 1} for all nodes"]

    style A fill:#c8e6c9
    style H fill:#e1f5fe
    style I fill:#c8e6c9
```

### Balance Factors and Heights

```mermaid
graph TD
    A["10 (BF: 0, h: 2)"] --> B["5 (BF: 0, h: 1)"]
    A --> C["15 (BF: 0, h: 1)"]
    B --> D["2 (BF: 0, h: 0)"]
    B --> E["7 (BF: 0, h: 0)"]
    C --> F["12 (BF: 0, h: 0)"]
    C --> G["20 (BF: 0, h: 0)"]

    H["BF = Balance Factor"]
    I["h = Height"]

    style A fill:#c8e6c9
    style H fill:#e1f5fe
```

### Unbalanced Tree Example

```mermaid
graph TD
    A["10 (BF: -2)"] --> B["5 (BF: 0)"]
    A --> C["15 (BF: -1)"]
    C --> D["12 (BF: 0)"]
    C --> E["20 (BF: -1)"]
    E --> F["25 (BF: 0)"]

    G["Node 10 has BF = -2, violates AVL property"]
    H["Needs right rotation to rebalance"]

    style A fill:#ffcdd2
    style G fill:#ffcdd2
    style H fill:#fff3e0
```

### Four Types of Rotations

```mermaid
graph LR
    subgraph "Left-Left (LL) Case"
        A1["z"] --> B1["y"]
        A1 --> C1["T4"]
        B1 --> D1["x"]
        B1 --> E1["T3"]
        D1 --> F1["T1"]
        D1 --> G1["T2"]

        H1["Right Rotation needed"]
    end

    subgraph "Right-Right (RR) Case"
        A2["z"] --> B2["T1"]
        A2 --> C2["y"]
        C2 --> D2["T2"]
        C2 --> E2["x"]
        E2 --> F2["T3"]
        E2 --> G2["T4"]

        H2["Left Rotation needed"]
    end

    style H1 fill:#fff3e0
    style H2 fill:#fff3e0
```

### Left Rotation (RR Case)

```mermaid
graph LR
    subgraph "Before Left Rotation"
        A["x"] --> B["T1"]
        A --> C["y"]
        C --> D["T2"]
        C --> E["T3"]
    end

    subgraph "After Left Rotation"
        F["y"] --> G["x"]
        F --> H["T3"]
        G --> I["T1"]
        G --> J["T2"]
    end

    K["y becomes new root"]
    L["x becomes left child of y"]

    style A fill:#ffcdd2
    style F fill:#c8e6c9
    style K fill:#c8e6c9
```

### Right Rotation (LL Case)

```mermaid
graph LR
    subgraph "Before Right Rotation"
        A["z"] --> B["y"]
        A --> C["T3"]
        B --> D["T1"]
        B --> E["T2"]
    end

    subgraph "After Right Rotation"
        F["y"] --> G["T1"]
        F --> H["z"]
        H --> I["T2"]
        H --> J["T3"]
    end

    K["y becomes new root"]
    L["z becomes right child of y"]

    style A fill:#ffcdd2
    style F fill:#c8e6c9
    style K fill:#c8e6c9
```

### Left-Right Rotation (LR Case)

```mermaid
graph LR
    subgraph "Step 1: Left Rotation on y"
        A1["z"] --> B1["y"]
        A1 --> C1["T4"]
        B1 --> D1["T1"]
        B1 --> E1["x"]
        E1 --> F1["T2"]
        E1 --> G1["T3"]
    end

    subgraph "Step 2: Right Rotation on z"
        A2["z"] --> B2["x"]
        A2 --> C2["T4"]
        B2 --> D2["y"]
        B2 --> E2["T3"]
        D2 --> F2["T1"]
        D2 --> G2["T2"]
    end

    subgraph "Final Result"
        A3["x"] --> B3["y"]
        A3 --> C3["z"]
        B3 --> D3["T1"]
        B3 --> E3["T2"]
        C3 --> F3["T3"]
        C3 --> G3["T4"]
    end

    style A3 fill:#c8e6c9
```

### Right-Left Rotation (RL Case)

```mermaid
graph LR
    subgraph "Step 1: Right Rotation on y"
        A1["z"] --> B1["T1"]
        A1 --> C1["y"]
        C1 --> D1["x"]
        C1 --> E1["T4"]
        D1 --> F1["T2"]
        D1 --> G1["T3"]
    end

    subgraph "Step 2: Left Rotation on z"
        A2["z"] --> B2["T1"]
        A2 --> C2["x"]
        C2 --> D2["T2"]
        C2 --> E2["y"]
        E2 --> F2["T3"]
        E2 --> G2["T4"]
    end

    subgraph "Final Result"
        A3["x"] --> B3["z"]
        A3 --> C3["y"]
        B3 --> D3["T1"]
        B3 --> E3["T2"]
        C3 --> F3["T3"]
        C3 --> G3["T4"]
    end

    style A3 fill:#c8e6c9
```

### AVL Insertion Algorithm

```mermaid
graph TD
    A[Insert node like BST] --> B[Update heights bottom-up]
    B --> C[Calculate balance factors]
    C --> D{BF > 1 or BF < -1?}
    D -->|No| E[Tree remains balanced]
    D -->|Yes| F{Which case?}

    F -->|LL| G[Right rotation]
    F -->|RR| H[Left rotation]
    F -->|LR| I[Left-Right rotation]
    F -->|RL| J[Right-Left rotation]

    G --> K[Update heights]
    H --> K
    I --> K
    J --> K
    K --> L[Tree balanced]

    style A fill:#e1f5fe
    style L fill:#c8e6c9
    style E fill:#c8e6c9
```

### AVL Deletion Algorithm

```mermaid
graph TD
    A[Delete node like BST] --> B[Update heights from deleted position]
    B --> C[Check balance at each ancestor]
    C --> D{Unbalanced node found?}
    D -->|No| E[Deletion complete]
    D -->|Yes| F[Determine rotation case]
    F --> G[Perform rotation]
    G --> H[Update heights]
    H --> I[Continue checking ancestors]
    I --> D

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Complexity Analysis

```mermaid
graph TD
    A[AVL Tree Complexity] --> B[Time Complexity]
    A --> C[Space Complexity]

    B --> B1["Search: O(log n)"]
    B --> B2["Insert: O(log n)"]
    B --> B3["Delete: O(log n)"]
    B --> B4["All operations guaranteed logarithmic"]

    C --> C1["Space: O(n) for storing nodes"]
    C --> C2["Height: O(log n) guaranteed"]
    C --> C3["Recursion depth: O(log n)"]

    D[Balance Guarantee] --> E["Max height difference: 1"]
    D --> F["Height ≤ 1.44 * log₂(n)"]
    D --> G["Better balance than Red-Black trees"]

    style A fill:#e1f5fe
    style B4 fill:#c8e6c9
    style F fill:#c8e6c9
```

### AVL vs Other Trees

```mermaid
graph TD
    A[Tree Comparison] --> B[AVL Tree]
    A --> C[Red-Black Tree]
    A --> D[Regular BST]

    B --> B1["Strictly balanced"]
    B --> B2["More rotations on insert/delete"]
    B --> B3["Faster lookups"]
    B --> B4["Height: ~1.44 log n"]

    C --> C1["Approximately balanced"]
    C --> C2["Fewer rotations"]
    C --> C3["Good for frequent modifications"]
    C --> C4["Height: ~2 log n"]

    D --> D1["No balance guarantee"]
    D --> D2["Can degenerate to O(n)"]
    D --> D3["Fastest insertions"]
    D --> D4["Height: O(n) worst case"]

    style B1 fill:#c8e6c9
    style C2 fill:#fff3e0
    style D2 fill:#ffcdd2
```

### Height Calculation and Update

```mermaid
graph LR
    subgraph "Height Calculation"
        A["height(node) = 1 + max(height(left), height(right))"]
        B["height(null) = -1"]
        C["leaf node height = 0"]
    end

    subgraph "Balance Factor"
        D["BF(node) = height(left) - height(right)"]
        E["BF = -1: right-heavy"]
        F["BF = 0: balanced"]
        G["BF = 1: left-heavy"]
    end

    subgraph "Update Strategy"
        H["Update heights bottom-up"]
        I["Check balance at each level"]
        J["Rotate if |BF| > 1"]
    end

    style A fill:#e1f5fe
    style D fill:#c8e6c9
    style H fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[AVL Tree Applications] --> B[Database Indexing]
    A --> C[Memory Management]
    A --> D[Language Implementations]
    A --> E[Graphics & Games]

    B --> B1["Database B+ trees"]
    B --> B2["Search indexing"]
    B --> B3["Range queries"]

    C --> C1["Memory allocators"]
    C --> C2["Virtual memory management"]
    C --> C3["Free block tracking"]

    D --> D1["Symbol tables"]
    D --> D2["Compiler implementations"]
    D --> D3["Standard libraries"]

    E --> E1["3D graphics engines"]
    E --> E2["Collision detection"]
    E --> E3["Spatial partitioning"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Implementation Considerations

```mermaid
graph TD
    A[Implementation Details] --> B[Node Structure]
    A --> C[Rotation Efficiency]
    A --> D[Height Storage]

    B --> B1["Store balance factor vs height"]
    B --> B2["Parent pointers optional"]
    B --> B3["Key-value pairs"]

    C --> C1["Single vs double rotations"]
    C --> C2["Iterative vs recursive"]
    C --> C3["Minimize pointer updates"]

    D --> D1["Cache heights in nodes"]
    D --> D2["Recalculate on demand"]
    D --> D3["Balance factor only"]

    E[Optimization Tips] --> F["Avoid unnecessary rotations"]
    E --> G["Batch operations when possible"]
    E --> H["Use iterative traversals"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

An AVL Tree is a self-balancing binary search tree where the heights of the two child subtrees of any node differ by at most one.
