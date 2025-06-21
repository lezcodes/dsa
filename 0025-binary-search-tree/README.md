# Binary Search Tree

## Description

A complete binary search tree (BST) implementation from scratch with recursive insertion and deletion operations. The BST maintains the fundamental property that for any node, all values in the left subtree are smaller and all values in the right subtree are larger. This implementation includes all three deletion cases and comprehensive tree operations.

## Key Features

- **Recursive Operations**: All core operations implemented using recursion
- **Complete Deletion**: Handles all three deletion cases properly
- **Tree Traversals**: In-order, pre-order, post-order, and level-order traversals
- **Tree Validation**: Ensures BST property is maintained
- **Successor/Predecessor**: Find next/previous elements in sorted order
- **Comprehensive API**: Standard BST operations plus utility methods

## Implementation Details

- **Node Structure**: Each node contains value, left child, and right child pointers
- **Recursive Design**: Insert, delete, search, and traversal operations use recursion
- **Size Tracking**: Maintains count of nodes for O(1) size queries
- **BST Property**: Left subtree < node < right subtree maintained at all times

## Complexity

- **Time Complexity**:
  - Average case: O(log n) for Insert, Delete, Search
  - Worst case: O(n) for unbalanced tree (degenerates to linked list)
  - Traversals: O(n) to visit all nodes
- **Space Complexity**: O(n) for storage, O(h) for recursion stack where h is height

## Deletion Cases

### Case 1: No Children (Leaf Node)

- **Description**: Node to delete has no left or right child
- **Solution**: Simply remove the node (set parent's pointer to nil)
- **Example**: Deleting leaf nodes like 10, 25, 35, 45 in a tree

### Case 2: One Child

- **Description**: Node to delete has exactly one child (left or right)
- **Solution**: Replace the node with its single child
- **Example**: If deleting node 20 with only left child 10, replace 20 with 10

### Case 3: Two Children

- **Description**: Node to delete has both left and right children
- **Solution**: Replace with in-order successor (smallest node in right subtree)
- **Process**:
  1. Find the in-order successor (leftmost node in right subtree)
  2. Replace the node's value with successor's value
  3. Recursively delete the successor (which has at most one child)

## Core Operations

### Basic Operations

- `Insert(value)` - Add a new value to the BST
- `Delete(value)` - Remove a value from the BST (handles all 3 cases)
- `Search(value)` - Check if a value exists in the BST
- `FindMin()` - Get the minimum value (leftmost node)
- `FindMax()` - Get the maximum value (rightmost node)
- `GetSize()` - Get the number of nodes
- `IsEmpty()` - Check if the tree is empty
- `Clear()` - Remove all nodes

### Traversal Operations

- `InOrderTraversal()` - Returns values in sorted order (left, root, right)
- `PreOrderTraversal()` - Returns values in pre-order (root, left, right)
- `PostOrderTraversal()` - Returns values in post-order (left, right, root)
- `LevelOrderTraversal()` - Returns values level by level (breadth-first)

### Utility Operations

- `Height()` - Get the height of the tree
- `IsValidBST()` - Verify the BST property is maintained
- `CountNodes()` - Count total nodes (alternative to GetSize)
- `CountLeaves()` - Count leaf nodes
- `GetSuccessor(value)` - Find the next larger value
- `GetPredecessor(value)` - Find the next smaller value
- `PrintTree()` - Visual representation of the tree structure

## Usage

```bash
make run n=binary-search-tree
```

## Testing

```bash
make test n=binary-search-tree
```

## Benchmarking

```bash
make bench n=binary-search-tree
```

## Example Usage in Go

```go
// Create a new BST
bst := NewBST()

// Insert values
values := []int{50, 30, 70, 20, 40, 60, 80}
for _, value := range values {
    bst.Insert(value)
}

// Search for values
exists := bst.Search(40)  // true
exists = bst.Search(100) // false

// Find min and max
min, hasMin := bst.FindMin() // 20, true
max, hasMax := bst.FindMax() // 80, true

// Get tree information
size := bst.GetSize()    // 7
height := bst.Height()   // 2
isEmpty := bst.IsEmpty() // false

// Traversals
inOrder := bst.InOrderTraversal()     // [20, 30, 40, 50, 60, 70, 80]
preOrder := bst.PreOrderTraversal()   // [50, 30, 20, 40, 70, 60, 80]
postOrder := bst.PostOrderTraversal() // [20, 40, 30, 60, 80, 70, 50]
levelOrder := bst.LevelOrderTraversal() // [50, 30, 70, 20, 40, 60, 80]

// Deletion examples for all three cases
bst.Delete(20)  // Case 1: No children (leaf)
bst.Delete(30)  // Case 2: One child (if 30 has only one child)
bst.Delete(50)  // Case 3: Two children (root with both children)

// Successor and predecessor
successor, exists := bst.GetSuccessor(40)   // Next larger value
predecessor, exists := bst.GetPredecessor(40) // Next smaller value

// Validation
isValid := bst.IsValidBST() // Verify BST property

// Tree statistics
nodeCount := bst.CountNodes()
leafCount := bst.CountLeaves()

// Clear the tree
bst.Clear()
```

## Deletion Algorithm Details

### Recursive Deletion Process

```
deleteRecursive(node, value):
    if node is null:
        return null

    if value < node.value:
        node.left = deleteRecursive(node.left, value)
    else if value > node.value:
        node.right = deleteRecursive(node.right, value)
    else:  // Found the node to delete
        // Case 1: No children
        if node.left is null and node.right is null:
            return null

        // Case 2: One child
        if node.left is null:
            return node.right
        if node.right is null:
            return node.left

        // Case 3: Two children
        successor = findMin(node.right)
        node.value = successor.value
        node.right = deleteRecursive(node.right, successor.value)

    return node
```

## Performance Characteristics

### Average Case (Balanced Tree)

- **Insert**: O(log n) - Navigate down one path to leaf
- **Delete**: O(log n) - Find node + potential successor lookup
- **Search**: O(log n) - Navigate down one path
- **Traversals**: O(n) - Visit every node once

### Worst Case (Unbalanced Tree)

- **All Operations**: O(n) - Tree degenerates to linked list
- **Space**: O(n) - Recursion stack depth equals tree height

### Best Case

- **Search/Insert/Delete**: O(1) - When target is root (for search/delete)

## Tree Traversal Explanations

### In-Order Traversal (Left, Root, Right)

- **Purpose**: Returns values in sorted order
- **Use Case**: Getting all values in ascending order
- **Property**: For BST, always produces sorted sequence

### Pre-Order Traversal (Root, Left, Right)

- **Purpose**: Visits root before children
- **Use Case**: Tree serialization, creating copy of tree
- **Property**: First element is always the root

### Post-Order Traversal (Left, Right, Root)

- **Purpose**: Visits children before root
- **Use Case**: Tree deletion, calculating tree properties
- **Property**: Last element is always the root

### Level-Order Traversal (Breadth-First)

- **Purpose**: Visits nodes level by level
- **Use Case**: Tree printing, finding nodes at specific depth
- **Implementation**: Uses queue for breadth-first search

## Use Cases

- **Database Indexing**: B-trees and B+ trees are extensions of BSTs
- **Expression Parsing**: Parse and evaluate mathematical expressions
- **File Systems**: Directory structures and file organization
- **Priority Queues**: When balanced (though heaps are more common)
- **Symbol Tables**: Compiler symbol table implementation
- **Range Queries**: Find all values within a range
- **Sorted Data Maintenance**: Keep data sorted with dynamic insertions/deletions

## Advantages

- **Sorted Order**: In-order traversal gives sorted sequence
- **Efficient Operations**: O(log n) average case for basic operations
- **Dynamic**: Supports insertions and deletions efficiently
- **Range Queries**: Easy to find values within a range
- **Memory Efficient**: Only stores data and two pointers per node

## Limitations

- **Unbalanced Performance**: Can degrade to O(n) in worst case
- **No Self-Balancing**: Doesn't automatically maintain balance
- **Memory Overhead**: Requires extra memory for pointers
- **Cache Performance**: Poor cache locality compared to arrays
- **Recursion Depth**: Deep trees can cause stack overflow

## Comparison with Alternatives

- **vs. Array**: BST allows efficient insertion/deletion, array has better cache locality
- **vs. Hash Table**: BST maintains order, hash table has O(1) average operations
- **vs. AVL/Red-Black Trees**: Those are self-balancing, guarantee O(log n) operations
- **vs. Heap**: BST supports range queries, heap is better for priority queue operations
- **vs. Linked List**: BST has much faster search (O(log n) vs O(n))

## Real-World Applications

- **Database Systems**: B-tree indexes in databases like MySQL, PostgreSQL
- **File Systems**: Directory structures in operating systems
- **Compilers**: Symbol tables and syntax tree representation
- **Game Development**: Spatial partitioning and collision detection
- **Network Routing**: Routing table implementations
- **Memory Management**: Free block management in allocators

## Visual Representation

### BST Structure and Property

```mermaid
graph TD
    A[50] --> B[30]
    A --> C[70]
    B --> D[20]
    B --> E[40]
    C --> F[60]
    C --> G[80]
    D --> H[10]
    D --> I[25]
    E --> J[35]
    E --> K[45]

    style A fill:#e1f5fe
    style B fill:#f3e5f5
    style C fill:#f3e5f5
    style D fill:#e8f5e8
    style E fill:#e8f5e8
    style F fill:#e8f5e8
    style G fill:#e8f5e8
```

### BST Property Visualization

```mermaid
graph LR
    subgraph "BST Property"
        A["For every node X:"]
        B["Left subtree values < X"]
        C["Right subtree values > X"]
    end

    subgraph "Example: Node 50"
        D["Left: 30, 20, 40, 10, 25, 35, 45"]
        E["Right: 70, 60, 80"]
        F["All left < 50 < All right"]
    end

    A --> B
    A --> C

    style A fill:#e1f5fe
    style F fill:#c8e6c9
```

### Insert Operation

```mermaid
graph TD
    A[Insert value: 55] --> B[Start at root]
    B --> C{value < root?}
    C -->|Yes| D[Go to left child]
    C -->|No| E[Go to right child]
    D --> F{left child exists?}
    E --> G{right child exists?}
    F -->|Yes| H[Repeat comparison]
    F -->|No| I[Insert as left child]
    G -->|Yes| J[Repeat comparison]
    G -->|No| K[Insert as right child]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
    style K fill:#c8e6c9
```

### Search Operation

```mermaid
graph TD
    A[Search for value: 35] --> B[Start at root: 50]
    B --> C{35 < 50?}
    C -->|Yes| D[Go left to: 30]
    D --> E{35 < 30?}
    E -->|No| F[Go right to: 40]
    F --> G{35 < 40?}
    G -->|Yes| H[Go left to: 35]
    H --> I[Found! Return true]

    C -->|No| J[Go right]
    E -->|Yes| K[Go left]
    G -->|No| L[Go right]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Delete Operation Cases

```mermaid
graph TD
    A[Delete Node] --> B{Node has children?}

    B -->|No children| C[Case 1: Leaf Node]
    B -->|One child| D[Case 2: One Child]
    B -->|Two children| E[Case 3: Two Children]

    C --> C1[Simply remove node]
    D --> D1[Replace with child]
    E --> E1[Find inorder successor]
    E1 --> E2[Replace value]
    E2 --> E3[Delete successor]

    style A fill:#e1f5fe
    style C1 fill:#c8e6c9
    style D1 fill:#fff3e0
    style E3 fill:#ffcdd2
```

### Inorder Traversal (Sorted Output)

```mermaid
graph TD
    A["BST Inorder Traversal"] --> B["Visit left subtree"]
    B --> C["Visit root"]
    C --> D["Visit right subtree"]

    E["Result: 10, 20, 25, 30, 35, 40, 45, 50, 60, 70, 80"]

    D --> E

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### BST vs Array vs Linked List

```mermaid
graph TD
    A[Data Structure Comparison] --> B[BST]
    A --> C[Sorted Array]
    A --> D[Linked List]

    B --> B1["Search: O(log n) avg, O(n) worst<br/>Insert: O(log n) avg, O(n) worst<br/>Delete: O(log n) avg, O(n) worst"]

    C --> C1["Search: O(log n)<br/>Insert: O(n)<br/>Delete: O(n)"]

    D --> D1["Search: O(n)<br/>Insert: O(1)<br/>Delete: O(n)"]

    style B1 fill:#c8e6c9
    style A fill:#e1f5fe
```

### Balanced vs Unbalanced BST

```mermaid
graph LR
    subgraph "Balanced BST"
        A1[50] --> B1[30]
        A1 --> C1[70]
        B1 --> D1[20]
        B1 --> E1[40]
        C1 --> F1[60]
        C1 --> G1[80]
    end

    subgraph "Unbalanced BST (Skewed)"
        A2[10] --> B2[20]
        B2 --> C2[30]
        C2 --> D2[40]
        D2 --> E2[50]
    end

    subgraph "Performance"
        P1["Balanced: O(log n)"]
        P2["Skewed: O(n)"]
    end

    style A1 fill:#c8e6c9
    style A2 fill:#ffcdd2
    style P1 fill:#c8e6c9
    style P2 fill:#ffcdd2
```

### Common BST Applications

```mermaid
graph TD
    A[BST Applications] --> B[Database Indexing]
    A --> C[File Systems]
    A --> D[Expression Parsing]
    A --> E[Priority Queues]
    A --> F[Set Operations]

    B --> B1["B-trees and B+ trees<br/>Fast database queries"]
    C --> C1["Directory structures<br/>File organization"]
    D --> D1["Operator precedence<br/>Parse trees"]
    E --> E1["Heap alternatives<br/>Ordered processing"]
    F --> F1["Union, intersection<br/>Sorted collections"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```
