# B-Tree

## Description

A B-Tree is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. It's optimized for systems that read and write large blocks of data, making it ideal for databases and file systems.

## Visual Representation

### B-Tree Structure (Order 3)

```mermaid
graph TD
    A["[10, 20]"] --> B["[3, 7]"]
    A --> C["[13, 17]"]
    A --> D["[22, 25, 30]"]

    B --> E["[1, 2]"]
    B --> F["[5, 6]"]
    B --> G["[8, 9]"]

    C --> H["[11, 12]"]
    C --> I["[15, 16]"]
    C --> J["[18, 19]"]

    D --> K["[21]"]
    D --> L["[23, 24]"]
    D --> M["[27, 28]"]
    D --> N["[32, 35]"]

    O["Order = 3 (minimum degree)"]
    P["Each node: 2-5 keys (except root)"]
    Q["All leaves at same level"]

    style A fill:#e1f5fe
    style O fill:#c8e6c9
    style Q fill:#c8e6c9
```

### B-Tree Properties

```mermaid
graph TD
    A[B-Tree Properties] --> B[Node Constraints]
    A --> C[Balance Property]
    A --> D[Ordering Property]

    B --> B1["Min degree t ≥ 2"]
    B --> B2["Root: 1 to 2t-1 keys"]
    B --> B3["Internal nodes: t-1 to 2t-1 keys"]
    B --> B4["Leaf nodes: t-1 to 2t-1 keys"]
    B --> B5["Internal nodes: t to 2t children"]

    C --> C1["All leaves at same level"]
    C --> C2["Height: O(log_t n)"]
    C --> C3["Perfectly balanced"]

    D --> D1["Keys sorted within nodes"]
    D --> D2["Subtree ordering preserved"]
    D --> D3["Enables efficient search"]

    style A fill:#e1f5fe
    style C1 fill:#c8e6c9
    style C2 fill:#c8e6c9
```

### Minimum Degree and Node Capacity

```mermaid
graph LR
    subgraph "Minimum Degree t = 3"
        A["Min keys per node: t-1 = 2"]
        B["Max keys per node: 2t-1 = 5"]
        C["Min children: t = 3"]
        D["Max children: 2t = 6"]
    end

    subgraph "Node Examples"
        E["Valid: [10, 20] (2 keys)"]
        F["Valid: [10, 20, 30, 40, 50] (5 keys)"]
        G["Invalid: [10] (too few, except root)"]
        H["Invalid: [10, 20, 30, 40, 50, 60] (too many)"]
    end

    style A fill:#c8e6c9
    style B fill:#c8e6c9
    style G fill:#ffcdd2
    style H fill:#ffcdd2
```

### B-Tree Search Algorithm

```mermaid
graph TD
    A[Search(key, node)] --> B{node is null?}
    B -->|Yes| C[Key not found]
    B -->|No| D[Binary search in node]
    D --> E{Key found?}
    E -->|Yes| F[Return success]
    E -->|No| G{Is leaf node?}
    G -->|Yes| C
    G -->|No| H[Find child to search]
    H --> I[Recursively search child]
    I --> J[Return result]

    K[Binary Search in Node] --> L["left = 0, right = keyCount-1"]
    L --> M["while left <= right"]
    M --> N["  mid = (left + right) / 2"]
    N --> O["  if key == keys[mid]: return mid"]
    O --> P["  else if key < keys[mid]: right = mid-1"]
    P --> Q["  else: left = mid+1"]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style C fill:#ffcdd2
```

### B-Tree Insertion Process

```mermaid
graph TD
    A[Insert(key)] --> B{Root is full?}
    B -->|Yes| C[Split root]
    B -->|No| D[Insert into non-full root]
    C --> E[Create new root]
    E --> D
    D --> F[Find appropriate leaf]
    F --> G[Insert key in leaf]
    G --> H{Leaf is full?}
    H -->|No| I[Insertion complete]
    H -->|Yes| J[Split leaf]
    J --> K[Propagate split upward]
    K --> L{Parent is full?}
    L -->|No| I
    L -->|Yes| M[Split parent]
    M --> K

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Node Splitting Example

```mermaid
graph LR
    subgraph "Before Split (t=3, max=5 keys)"
        A["[10, 20, 30, 40, 50]"]
        B["Node is full, need to split"]
    end

    subgraph "After Split"
        C["Parent gets: [30]"]
        C --> D["Left: [10, 20]"]
        C --> E["Right: [40, 50]"]
    end

    subgraph "Split Algorithm"
        F["1. Find median key"]
        G["2. Create new node"]
        H["3. Move upper half to new node"]
        I["4. Promote median to parent"]
        J["5. Update parent pointers"]
    end

    style A fill:#ffcdd2
    style B fill:#ffcdd2
    style C fill:#c8e6c9
    style I fill:#c8e6c9
```

### B-Tree Deletion Cases

```mermaid
graph TD
    A[Delete(key)] --> B[Find key location]
    B --> C{Key in leaf?}
    C -->|Yes| D[Case 1: Delete from leaf]
    C -->|No| E[Case 2: Delete from internal]

    D --> F{Leaf has >= t keys?}
    F -->|Yes| G[Simply remove key]
    F -->|No| H[Case 3: Borrow or merge]

    E --> I[Replace with predecessor/successor]
    I --> J[Delete predecessor/successor]
    J --> H

    H --> K{Sibling has >= t keys?}
    K -->|Yes| L[Borrow from sibling]
    K -->|No| M[Merge with sibling]

    L --> N[Deletion complete]
    M --> O{Parent underflows?}
    O -->|No| N
    O -->|Yes| P[Recursively fix parent]
    P --> O

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style N fill:#c8e6c9
```

### Borrowing from Sibling

```mermaid
graph LR
    subgraph "Before Borrowing"
        A["Parent: [20]"]
        A --> B["Left: [10] (underflow)"]
        A --> C["Right: [30, 40, 50]"]
    end

    subgraph "After Borrowing"
        D["Parent: [30]"]
        D --> E["Left: [10, 20]"]
        D --> F["Right: [40, 50]"]
    end

    subgraph "Borrowing Process"
        G["1. Move parent key to underflow node"]
        H["2. Move sibling's key to parent"]
        I["3. Move sibling's child if internal"]
    end

    style B fill:#ffcdd2
    style E fill:#c8e6c9
    style H fill:#c8e6c9
```

### Merging Nodes

```mermaid
graph LR
    subgraph "Before Merge"
        A["Parent: [20, 40]"]
        A --> B["Left: [10] (underflow)"]
        A --> C["Middle: [30]"]
        A --> D["Right: [50, 60]"]
    end

    subgraph "After Merge"
        E["Parent: [40]"]
        E --> F["Merged: [10, 20, 30]"]
        E --> G["Right: [50, 60]"]
    end

    subgraph "Merge Process"
        H["1. Pull down parent key"]
        I["2. Combine with sibling"]
        J["3. Remove parent key"]
        K["4. Update parent pointers"]
    end

    style B fill:#ffcdd2
    style F fill:#c8e6c9
    style I fill:#c8e6c9
```

### B-Tree vs B+ Tree

```mermaid
graph TD
    A[Comparison] --> B[B-Tree]
    A --> C[B+ Tree]

    B --> B1["Data in all nodes"]
    B --> B2["Smaller tree height"]
    B --> B3["Direct key access"]
    B --> B4["Complex deletion"]

    C --> C1["Data only in leaves"]
    C --> C2["Larger tree height"]
    C --> C3["Sequential leaf scanning"]
    C --> C4["Simpler operations"]
    C --> C5["Better for range queries"]

    D[Use Cases] --> E["B-Tree: In-memory databases"]
    D --> F["B+ Tree: File systems, disk-based DBs"]

    style B2 fill:#c8e6c9
    style C5 fill:#c8e6c9
    style F fill:#c8e6c9
```

### Performance Analysis

```mermaid
graph TD
    A[B-Tree Performance] --> B[Time Complexity]
    A --> C[Space Complexity]
    A --> D[I/O Complexity]

    B --> B1["Search: O(log_t n)"]
    B --> B2["Insert: O(log_t n)"]
    B --> B3["Delete: O(log_t n)"]
    B --> B4["Range query: O(log_t n + k)"]

    C --> C1["Storage: O(n)"]
    C --> C2["Height: O(log_t n)"]
    C --> C3["Node size: O(t)"]

    D --> D1["Disk reads: O(log_t n)"]
    D --> D2["Optimal for block storage"]
    D --> D3["Minimizes disk seeks"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style D2 fill:#c8e6c9
    style D3 fill:#c8e6c9
```

### Real-World Applications

```mermaid
graph TD
    A[B-Tree Applications] --> B[Database Systems]
    A --> C[File Systems]
    A --> D[Operating Systems]
    A --> E[Search Engines]

    B --> B1["MySQL InnoDB indexes"]
    B --> B2["PostgreSQL indexes"]
    B --> B3["SQLite indexes"]
    B --> B4["MongoDB indexes"]

    C --> C1["NTFS file system"]
    C --> C2["HFS+ file system"]
    C --> C3["ext4 directory indexes"]

    D --> D1["Process scheduling"]
    D --> D2["Memory management"]
    D --> D3["Virtual memory"]

    E --> E1["Inverted indexes"]
    E --> E2["Full-text search"]
    E --> E3["Document retrieval"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style E1 fill:#c8e6c9
```

### Implementation Considerations

```mermaid
graph TD
    A[Implementation Details] --> B[Node Structure]
    A --> C[Memory Management]
    A --> D[Disk Optimization]

    B --> B1["Key and child arrays"]
    B --> B2["Key count tracking"]
    B --> B3["Leaf flag indication"]
    B --> B4["Parent pointers optional"]

    C --> C1["Pool allocation"]
    C --> C2["Cache management"]
    C --> C3["Reference counting"]

    D --> D1["Block size alignment"]
    D --> D2["Buffering strategies"]
    D --> D3["Write-ahead logging"]
    D --> D4["Compression techniques"]

    E[Optimization Tips] --> F["Choose appropriate t"]
    E --> G["Batch operations"]
    E --> H["Lazy deletion"]

    style A fill:#e1f5fe
    style D1 fill:#c8e6c9
    style F fill:#c8e6c9
    style G fill:#c8e6c9
```

### Minimum Degree Selection

```mermaid
graph LR
    subgraph "Small t (t=2)"
        A["More levels"]
        B["Less memory per node"]
        C["More disk I/O"]
        D["Simpler node operations"]
    end

    subgraph "Large t (t=100)"
        E["Fewer levels"]
        F["More memory per node"]
        G["Less disk I/O"]
        H["Complex node operations"]
    end

    subgraph "Optimal Choice"
        I["Match disk block size"]
        J["Balance CPU vs I/O"]
        K["Consider cache effects"]
        L["Typical: t = 50-200"]
    end

    style C fill:#ffcdd2
    style G fill:#c8e6c9
    style I fill:#c8e6c9
    style L fill:#c8e6c9
```

### B-Tree Variants

```mermaid
graph TD
    A[B-Tree Family] --> B[Standard B-Tree]
    A --> C[B+ Tree]
    A --> D[B* Tree]
    A --> E[Counted B-Tree]

    B --> B1["Data in all nodes"]
    B --> B2["Direct access to keys"]
    B --> B3["Compact structure"]

    C --> C1["Data only in leaves"]
    C --> C2["Linked leaf nodes"]
    C --> C3["Sequential access"]

    D --> D1["2/3 full guarantee"]
    D --> D2["Better space utilization"]
    D --> D3["Delayed splitting"]

    E --> E1["Node size tracking"]
    E --> E2["Order statistics"]
    E --> E3["Rank queries"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style C3 fill:#c8e6c9
    style D2 fill:#c8e6c9
```

A B-Tree is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time.

### Key Properties

- All leaves are at the same level
- A B-tree is defined by the term minimum degree `t` (≥ 2)
- Every node except root must contain at least `t-1` keys
- Every node may contain at most `2t-1` keys
- Number of children of a node is equal to the number of keys in it plus 1
- All keys of a node are sorted in increasing order
- The child between two keys k1 and k2 contains all keys in the range from k1 and k2
- B-tree grows and shrinks from the root which is unlike Binary Search Tree

### Features

- **Configurable Minimum Degree**: Support for different minimum degrees (t ≥ 2)
- **Automatic Node Splitting**: Nodes automatically split when they exceed capacity
- **Self-Balancing**: Maintains balanced height through splits and merges
- **Efficient Operations**: Logarithmic time complexity for all major operations
- **Both Recursive and Iterative**: Default recursive with iterative alternatives
- **Comprehensive Validation**: Built-in tree structure validation
- **Multiple Traversals**: In-order, pre-order, and level-order traversals

## Complexity

- **Time Complexity**:
  - Search: O(log n)
  - Insert: O(log n)
  - Delete: O(log n)
  - Traversal: O(n)
- **Space Complexity**: O(n)
- **Height**: O(log_t n) where t is the minimum degree

## Real-World Applications

- **Database Systems**: B-trees are widely used in database indexing (MySQL, PostgreSQL)
- **File Systems**: Many file systems use B-trees for directory structures (NTFS, HFS+)
- **Operating Systems**: Used in memory management and virtual memory systems
- **Search Engines**: Efficient indexing of large datasets
- **Distributed Systems**: B-trees work well with disk-based storage due to high branching factor

## API Reference

### Core Operations

```go
bt := NewBTree(3)                    // Create B-tree with minimum degree 3
bt.Insert(key)                       // Insert a key
found := bt.Search(key)              // Search for a key (recursive)
found := bt.SearchIterative(key)     // Search for a key (iterative)
success := bt.Delete(key)            // Delete a key
```

### Tree Information

```go
size := bt.GetSize()                 // Get number of keys
height := bt.GetHeight()             // Get tree height
degree := bt.GetMinimumDegree()      // Get minimum degree
isEmpty := bt.IsEmpty()              // Check if tree is empty
nodeCount := bt.GetNodeCount()       // Get total number of nodes
leafCount := bt.GetLeafCount()       // Get number of leaf nodes
```

### Traversals

```go
inOrder := bt.InOrderTraversal()     // Get keys in sorted order
preOrder := bt.PreOrderTraversal()   // Get keys in pre-order
levelOrder := bt.LevelOrderTraversal() // Get keys in level-order
allKeys := bt.GetAllKeys()           // Alias for InOrderTraversal
```

### Utility Operations

```go
min, hasMin := bt.FindMin()          // Find minimum key
max, hasMax := bt.FindMax()          // Find maximum key
isValid := bt.Validate()             // Validate B-tree properties
bt.PrintTree()                       // Print tree structure
bt.Clear()                           // Clear all nodes
```

## Usage

```bash
make run n=0029-b-tree
```

## Testing

```bash
make test n=0029-b-tree
```

## Implementation Details

### Node Structure

- **Keys**: Sorted array of keys (max 2t-1)
- **Children**: Array of child pointers (max 2t)
- **IsLeaf**: Boolean flag indicating leaf status
- **Parent**: Pointer to parent node for efficient operations

### Insertion Algorithm

1. If tree is empty, create root with the key
2. If root is full, split it and create new root
3. Insert into appropriate non-full node
4. Split nodes as needed during insertion

### Deletion Algorithm

1. **Case 1**: Key in leaf node - simply remove
2. **Case 2**: Key in internal node - replace with predecessor/successor
3. **Case 3**: Key not found - recursively delete from appropriate child
4. Handle underflow by borrowing from siblings or merging nodes

### Splitting Process

When a node becomes full (2t-1 keys):

1. Create new node
2. Move upper half of keys to new node
3. Move median key up to parent
4. Update child pointers appropriately

### Performance Characteristics

- **Optimal Degree**: For disk-based systems, degree is typically chosen based on page size
- **Memory Efficiency**: High branching factor reduces tree height
- **Cache Performance**: Good locality of reference within nodes
- **Disk I/O**: Minimizes disk accesses due to high branching factor

## Benchmarks

The implementation includes comprehensive benchmarks for:

- Insert operations
- Search operations (both recursive and iterative)
- Delete operations
- Traversal operations

Typical performance on modern hardware:

- Insert: ~10-50 ns/operation
- Search: ~10-30 ns/operation
- Delete: ~50-200 ns/operation
