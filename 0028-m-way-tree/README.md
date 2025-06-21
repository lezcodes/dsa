# m-way-tree

## Description

M-way Tree (also known as Multiway Tree or General Tree) is a tree data structure where each node can have up to M children. This implementation provides a configurable branching factor and supports automatic node splitting when nodes exceed capacity, similar to B-tree behavior.

### Key Features

- **Configurable branching factor**: Set maximum number of children per node (M)
- **Automatic node splitting**: Nodes split when they exceed M-1 keys
- **Sorted keys**: Keys within each node are maintained in sorted order
- **Dual search methods**: Both recursive and iterative search implementations
- **Complete operations**: Insert, delete, search, traversals, and validation
- **Tree validation**: Ensures structural integrity and ordering properties

### M-way Tree Properties

1. **Branching Factor**: Each node can have at most M children
2. **Key Capacity**: Each node can store at most M-1 keys
3. **Sorted Order**: Keys within each node are sorted in ascending order
4. **Leaf Identification**: Leaf nodes have no children
5. **Parent-Child Relationships**: Maintained for efficient operations

### Node Structure

- **Keys**: Array of sorted integer values (max M-1)
- **Children**: Array of child node pointers (max M)
- **IsLeaf**: Boolean flag indicating if node is a leaf
- **Parent**: Pointer to parent node for upward traversal

## Complexity

- **Time Complexity**:

  - Insert: O(log_M n) - logarithmic with base M
  - Delete: O(log_M n) - includes predecessor finding
  - Search: O(log_M n) - both recursive and iterative
  - Traversals: O(n) - visit each key once
  - Height: O(log_M n) - computed recursively
  - Validation: O(n) - checks all nodes and relationships

- **Space Complexity**:
  - Storage: O(n) - one key per element plus node overhead
  - Recursive operations: O(log_M n) - call stack depth
  - Node capacity: O(M) - keys and children arrays per node

## Implementation Details

### Core Methods

- `NewMWayTree(m)` - Creates tree with branching factor M (minimum 3)
- `Insert(key)` - Inserts key with automatic node splitting
- `Delete(key)` - Removes key with predecessor replacement for internal nodes
- `Search(key)` - Recursive search through tree
- `SearchIterative(key)` - Iterative search alternative

### Node Management

- `splitNode(node)` - Splits overfull nodes, promotes middle key
- `findKeyPosition(keys, key)` - Binary search within node keys
- `findPredecessor(node, pos)` - Finds in-order predecessor for deletion

### Traversal Methods

- `InOrderTraversal()` - Returns keys in sorted order
- `PreOrderTraversal()` - Node-first traversal
- `PostOrderTraversal()` - Children-first traversal
- `LevelOrderTraversal()` - Breadth-first traversal

### Utility Methods

- `GetHeight()` - Returns tree height
- `GetBranchingFactor()` - Returns M value
- `GetSize()` - Returns total number of keys
- `GetNodeCount()` - Returns total number of nodes
- `GetLeafCount()` - Returns number of leaf nodes
- `FindMin()/FindMax()` - Find minimum/maximum values
- `GetAllKeys()` - Returns all keys in sorted order
- `Validate()` - Validates tree structure and properties
- `Clear()` - Removes all nodes
- `PrintTree()` - Visual tree representation

### Splitting Algorithm

When a node reaches M keys:

1. Find middle key (index M/2)
2. Create new right node with keys after middle
3. Keep left keys in original node
4. Promote middle key to parent
5. If no parent exists, create new root

## Usage

```bash
make run n=0028-m-way-tree
```

### Example Operations

```go
// Create M-way tree with branching factor 4
tree := NewMWayTree(4)

// Insert values (triggers automatic splitting)
values := []int{10, 20, 5, 6, 12, 30, 7, 17}
for _, value := range values {
    tree.Insert(value)
}

// Search operations
found := tree.Search(20)              // true
foundIter := tree.SearchIterative(20) // true

// Tree properties
size := tree.GetSize()               // 8
height := tree.GetHeight()           // Depends on splits
branchingFactor := tree.GetBranchingFactor() // 4
nodeCount := tree.GetNodeCount()     // Number of internal nodes
leafCount := tree.GetLeafCount()     // Number of leaf nodes
valid := tree.Validate()             // true

// Traversals
inOrder := tree.InOrderTraversal()    // [5, 6, 7, 10, 12, 17, 20, 30]
preOrder := tree.PreOrderTraversal()  // Node-first order
postOrder := tree.PostOrderTraversal() // Children-first order
levelOrder := tree.LevelOrderTraversal() // Level by level
allKeys := tree.GetAllKeys()          // Sorted array of all keys

// Min/Max operations
min, hasMin := tree.FindMin() // 5, true
max, hasMax := tree.FindMax() // 30, true

// Deletion
deleted := tree.Delete(12) // true
newSize := tree.GetSize()  // 7
stillValid := tree.Validate() // true

// Different branching factors
tree3 := NewMWayTree(3)  // Ternary tree
tree10 := NewMWayTree(10) // Wider tree
```

## Testing

```bash
make test n=0028-m-way-tree
```

### Test Coverage

- **Basic Operations**: Insert, delete, search functionality
- **Node Splitting**: Automatic splitting when nodes exceed capacity
- **Different Branching Factors**: Testing with M = 3, 4, 5, 10
- **Edge Cases**: Empty tree, single node, duplicates
- **Traversals**: All traversal methods validated
- **Large Datasets**: 1000+ key stress testing
- **Validation**: Tree structure and property checking
- **Performance**: Benchmarks for all major operations

## Real-World Applications

- **Database Indexing**: Foundation for B-tree and B+ tree implementations
- **File Systems**: Directory structures with configurable fanout
- **Decision Trees**: Multi-way decision nodes in machine learning
- **Parsing**: Abstract syntax trees with variable children
- **Network Routing**: Routing tables with multiple next-hop options
- **Memory Management**: Allocation trees with configurable branching

## Advantages

- **Configurable Performance**: Adjust branching factor for specific use cases
- **Reduced Height**: Higher branching factor means shorter trees
- **Cache Efficiency**: Fewer levels mean fewer memory accesses
- **Flexible Structure**: Adapts to different data distributions
- **Foundation for B-trees**: Core concepts apply to database indexing

## Comparison with Other Trees

- **vs Binary Trees**: Reduced height, more complex node management
- **vs B-Trees**: Simpler implementation, no minimum fill requirements
- **vs Trie**: More general purpose, not limited to string prefixes
- **vs Heap**: Maintains sorted order, supports range queries

## Performance Characteristics

- **Height**: O(log_M n) - decreases as M increases
- **Node Utilization**: Variable, depends on insertion order
- **Split Frequency**: Decreases with higher M values
- **Search Efficiency**: Fewer comparisons per level, more comparisons per node
- **Memory Overhead**: Increases with M due to larger node arrays

## Branching Factor Selection

- **M = 3**: Similar to binary tree, frequent splits
- **M = 4-8**: Good balance for most applications
- **M = 16-64**: Suitable for disk-based storage (B-tree territory)
- **M > 64**: Diminishing returns, increased node search time

## Tree Validation Rules

1. **Key Count**: Each node has at most M-1 keys
2. **Child Count**: Internal nodes have exactly (keys + 1) children
3. **Key Ordering**: Keys within nodes are sorted
4. **Leaf Property**: Leaf nodes have no children
5. **Parent Consistency**: All parent-child relationships are bidirectional

## Visual Representation

### Basic M-Way Tree Structure (M=4)

```mermaid
graph TD
    A["[10, 20, 30]"] --> B["[5, 7]"]
    A --> C["[15, 17]"]
    A --> D["[25, 27]"]
    A --> E["[35, 37, 40]"]

    B --> F["[1, 2]"]
    B --> G["[6]"]
    B --> H["[8, 9]"]

    C --> I["[12, 13]"]
    C --> J["[16]"]
    C --> K["[18, 19]"]

    L["M = 4 (4-way tree)"]
    M["Each node has at most 3 keys"]
    N["Each node has at most 4 children"]

    style A fill:#e1f5fe
    style L fill:#c8e6c9
    style M fill:#c8e6c9
```

### Node Structure and Properties

```mermaid
graph LR
    subgraph "Internal Node Structure"
        A["key₁ | key₂ | ... | keyₖ"]
        B["ptr₀ | ptr₁ | ... | ptrₖ"]
        C["k ≤ M-1 keys"]
        D["k+1 ≤ M pointers"]
    end

    subgraph "Key Ordering"
        E["key₁ < key₂ < ... < keyₖ"]
        F["ptr₀: children < key₁"]
        G["ptrᵢ: keyᵢ < children < keyᵢ₊₁"]
        H["ptrₖ: children > keyₖ"]
    end

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### M-Way Tree Properties

```mermaid
graph TD
    A[M-Way Tree Properties] --> B[Node Constraints]
    A --> C[Key Ordering]
    A --> D[Search Properties]

    B --> B1["Max M children per node"]
    B --> B2["Max M-1 keys per node"]
    B --> B3["Min 1 key per node (except root)"]
    B --> B4["Leaf nodes have no children"]

    C --> C1["Keys sorted within each node"]
    C --> C2["All keys in subtree follow BST property"]
    C --> C3["Left subtree < key < right subtree"]

    D --> D1["Search time: O(log_M n)"]
    D --> D2["Height: O(log_M n)"]
    D --> D3["Disk I/O efficient"]

    style A fill:#e1f5fe
    style D1 fill:#c8e6c9
    style D2 fill:#c8e6c9
```

### Search Algorithm

```mermaid
graph TD
    A[Search(key, node)] --> B{node is null?}
    B -->|Yes| C[Key not found]
    B -->|No| D[Search keys in node]
    D --> E{Key found?}
    E -->|Yes| F[Return success]
    E -->|No| G[Find appropriate child pointer]
    G --> H[Recursively search child]
    H --> I[Return result]

    J[Linear Search in Node] --> K["for i = 0 to node.keyCount-1"]
    K --> L["  if key == node.keys[i]: return found"]
    L --> M["  if key < node.keys[i]: child = node.children[i]"]
    M --> N["child = node.children[node.keyCount] // rightmost"]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style C fill:#ffcdd2
```

### Insertion Process

```mermaid
graph TD
    A[Insert(key)] --> B[Find leaf node position]
    B --> C[Insert key in sorted order]
    C --> D{Node overflow? (keys > M-1)}
    D -->|No| E[Insertion complete]
    D -->|Yes| F[Split node]
    F --> G[Promote middle key to parent]
    G --> H{Parent overflow?}
    H -->|No| I[Insertion complete]
    H -->|Yes| J[Split parent recursively]
    J --> K{Reached root?}
    K -->|No| H
    K -->|Yes| L[Create new root]
    L --> I

    style A fill:#e1f5fe
    style E fill:#c8e6c9
    style I fill:#c8e6c9
```

### Node Splitting Example (M=4)

```mermaid
graph LR
    subgraph "Before Split"
        A["[10, 20, 30, 40]"]
        B["Overflow! 4 keys in M=4 tree"]
    end

    subgraph "After Split"
        C["[20]"] --> D["[10]"]
        C --> E["[30, 40]"]
        F["Promote middle key (20) to parent"]
    end

    subgraph "Split Algorithm"
        G["1. Find middle key"]
        H["2. Create two new nodes"]
        I["3. Distribute keys and children"]
        J["4. Promote middle key"]
    end

    style A fill:#ffcdd2
    style B fill:#ffcdd2
    style C fill:#c8e6c9
    style F fill:#c8e6c9
```

### Deletion Algorithm

```mermaid
graph TD
    A[Delete(key)] --> B[Find key location]
    B --> C{Key in leaf node?}
    C -->|Yes| D[Remove key directly]
    C -->|No| E[Find inorder successor/predecessor]
    E --> F[Replace key with successor]
    F --> G[Delete successor from leaf]

    D --> H{Node underflow?}
    G --> H
    H -->|No| I[Deletion complete]
    H -->|Yes| J[Try borrowing from sibling]
    J --> K{Sibling has extra keys?}
    K -->|Yes| L[Borrow key through parent]
    K -->|No| M[Merge with sibling]
    L --> I
    M --> N{Parent underflow?}
    N -->|No| I
    N -->|Yes| O[Fix parent recursively]
    O --> N

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Borrowing and Merging

```mermaid
graph LR
    subgraph "Borrowing Example"
        A["Parent: [20]"]
        A --> B["Left: [10]"]
        A --> C["Right: [30, 40, 50]"]

        D["After Borrowing"]
        E["Parent: [30]"]
        E --> F["Left: [10, 20]"]
        E --> G["Right: [40, 50]"]
    end

    subgraph "Merging Example"
        H["Parent: [20, 40]"]
        H --> I["Left: [10]"]
        H --> J["Middle: [30]"]
        H --> K["Right: [50, 60]"]

        L["After Merging"]
        M["Parent: [40]"]
        M --> N["Merged: [10, 20, 30]"]
        M --> O["Right: [50, 60]"]
    end

    style E fill:#c8e6c9
    style M fill:#c8e6c9
```

### M-Way Tree vs Binary Tree

```mermaid
graph TD
    A[Comparison] --> B[M-Way Tree]
    A --> C[Binary Tree]

    B --> B1["Height: O(log_M n)"]
    B --> B2["More keys per node"]
    B --> B3["Fewer disk I/O operations"]
    B --> B4["Complex node operations"]
    B --> B5["Better for external storage"]

    C --> C1["Height: O(log₂ n)"]
    C --> C2["One key per node"]
    C --> C3["More node accesses"]
    C --> C4["Simple operations"]
    C --> C5["Better for main memory"]

    D[Trade-offs] --> E["Fewer levels vs more complex nodes"]
    D --> F["I/O efficiency vs CPU complexity"]

    style B3 fill:#c8e6c9
    style C4 fill:#c8e6c9
    style E fill:#fff3e0
```

### Applications and Use Cases

```mermaid
graph TD
    A[M-Way Tree Applications] --> B[Database Systems]
    A --> C[File Systems]
    A --> D[B-Tree Foundation]
    A --> E[External Memory]

    B --> B1["Index structures"]
    B --> B2["Range queries"]
    B --> B3["Sorted data access"]

    C --> C1["Directory structures"]
    C --> C2["File allocation tables"]
    C --> C3["Metadata indexing"]

    D --> D1["B-trees for databases"]
    D --> D2["B+ trees for file systems"]
    D --> D3["Balanced variants"]

    E --> E1["Large dataset indexing"]
    E --> E2["Disk-based storage"]
    E --> E3["Minimize I/O operations"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style D1 fill:#c8e6c9
    style E3 fill:#c8e6c9
```

### Performance Analysis

```mermaid
graph LR
    subgraph "Time Complexity"
        A["Search: O(M * log_M n)"]
        B["Insert: O(M * log_M n)"]
        C["Delete: O(M * log_M n)"]
        D["M factor from linear search in node"]
    end

    subgraph "Space Complexity"
        E["Storage: O(n)"]
        F["Height: O(log_M n)"]
        G["Node size: O(M)"]
    end

    subgraph "I/O Complexity"
        H["Disk reads: O(log_M n)"]
        I["Better than binary tree: O(log₂ n)"]
        J["Optimal for external memory"]
    end

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style J fill:#c8e6c9
```

### Implementation Considerations

```mermaid
graph TD
    A[Implementation Details] --> B[Node Design]
    A --> C[Memory Management]
    A --> D[Optimization]

    B --> B1["Fixed vs variable size arrays"]
    B --> B2["Key and pointer arrays"]
    B --> B3["Key count tracking"]

    C --> C1["Node allocation strategies"]
    C --> C2["Memory pools"]
    C --> C3["Garbage collection"]

    D --> D1["Binary search in nodes"]
    D --> D2["Bulk loading techniques"]
    D --> D3["Cache-friendly layouts"]

    E[Design Choices] --> F["Order M selection"]
    E --> G["Split/merge thresholds"]
    E --> H["Balancing strategies"]

    style A fill:#e1f5fe
    style D1 fill:#c8e6c9
    style F fill:#c8e6c9
```

### Variants and Extensions

```mermaid
graph TD
    A[M-Way Tree Variants] --> B[B-Trees]
    A --> C[B+ Trees]
    A --> D[B* Trees]
    A --> E[Specialized Trees]

    B --> B1["Minimum degree constraints"]
    B --> B2["Balanced height guarantee"]
    B --> B3["Database indexing"]

    C --> C1["Data only in leaves"]
    C --> C2["Linked leaf nodes"]
    C --> C3["Sequential access optimization"]

    D --> D1["Higher space utilization"]
    D --> D2["Delayed splitting"]
    D --> D3["Better performance"]

    E --> E1["R-trees for spatial data"]
    E --> E2["LSM-trees for writes"]
    E --> E3["Fractal trees"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style C2 fill:#c8e6c9
    style D3 fill:#c8e6c9
```

An M-Way Tree is a tree data structure where each internal node can have at most M children and M-1 keys.
