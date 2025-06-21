# Tree Traversal

## Description

Tree traversal refers to the process of visiting each node in a tree data structure exactly once in a systematic way. This implementation covers the three main depth-first traversal methods: inorder, preorder, and postorder.

## Visual Representation

### Binary Tree Structure

```mermaid
graph TD
    A[1] --> B[2]
    A --> C[3]
    B --> D[4]
    B --> E[5]
    C --> F[6]
    C --> G[7]

    style A fill:#e1f5fe
    style B fill:#f3e5f5
    style C fill:#f3e5f5
    style D fill:#e8f5e8
    style E fill:#e8f5e8
    style F fill:#e8f5e8
    style G fill:#e8f5e8
```

### Inorder Traversal (Left → Root → Right)

```mermaid
graph TD
    A["1<br/>(4th)"] --> B["2<br/>(2nd)"]
    A --> C["3<br/>(6th)"]
    B --> D["4<br/>(1st)"]
    B --> E["5<br/>(3rd)"]
    C --> F["6<br/>(5th)"]
    C --> G["7<br/>(7th)"]

    H[Inorder: 4, 2, 5, 1, 6, 3, 7]

    style D fill:#c8e6c9
    style B fill:#e8f5e8
    style E fill:#fff3e0
    style A fill:#ffcdd2
    style F fill:#e1f5fe
    style C fill:#f3e5f5
    style G fill:#fce4ec
```

### Preorder Traversal (Root → Left → Right)

```mermaid
graph TD
    A["1<br/>(1st)"] --> B["2<br/>(2nd)"]
    A --> C["3<br/>(5th)"]
    B --> D["4<br/>(3rd)"]
    B --> E["5<br/>(4th)"]
    C --> F["6<br/>(6th)"]
    C --> G["7<br/>(7th)"]

    H[Preorder: 1, 2, 4, 5, 3, 6, 7]

    style A fill:#c8e6c9
    style B fill:#e8f5e8
    style D fill:#fff3e0
    style E fill:#ffcdd2
    style C fill:#e1f5fe
    style F fill:#f3e5f5
    style G fill:#fce4ec
```

### Postorder Traversal (Left → Right → Root)

```mermaid
graph TD
    A["1<br/>(7th)"] --> B["2<br/>(3rd)"]
    A --> C["3<br/>(6th)"]
    B --> D["4<br/>(1st)"]
    B --> E["5<br/>(2nd)"]
    C --> F["6<br/>(4th)"]
    C --> G["7<br/>(5th)"]

    H[Postorder: 4, 5, 2, 6, 7, 3, 1]

    style D fill:#c8e6c9
    style E fill:#e8f5e8
    style B fill:#fff3e0
    style F fill:#ffcdd2
    style G fill:#e1f5fe
    style C fill:#f3e5f5
    style A fill:#fce4ec
```

### Traversal Algorithm Flow

```mermaid
graph TD
    A[Start Traversal] --> B{Choose Method}

    B -->|Inorder| C[Visit Left Subtree]
    C --> D[Visit Root]
    D --> E[Visit Right Subtree]

    B -->|Preorder| F[Visit Root]
    F --> G[Visit Left Subtree]
    G --> H[Visit Right Subtree]

    B -->|Postorder| I[Visit Left Subtree]
    I --> J[Visit Right Subtree]
    J --> K[Visit Root]

    E --> L[Complete]
    H --> L
    K --> L

    style A fill:#e1f5fe
    style L fill:#c8e6c9
```

### Recursive Implementation Pattern

```mermaid
graph TD
    A["traverse(node)"] --> B{"node == null?"}
    B -->|Yes| C[Return]
    B -->|No| D[Apply traversal pattern]

    D --> E["Inorder:<br/>traverse(left)<br/>visit(node)<br/>traverse(right)"]
    D --> F["Preorder:<br/>visit(node)<br/>traverse(left)<br/>traverse(right)"]
    D --> G["Postorder:<br/>traverse(left)<br/>traverse(right)<br/>visit(node)"]

    style A fill:#e1f5fe
    style C fill:#c8e6c9
```

### Use Cases for Different Traversals

```mermaid
graph TD
    A[Tree Traversal Applications] --> B[Inorder]
    A --> C[Preorder]
    A --> D[Postorder]

    B --> B1["Binary Search Trees<br/>Get sorted sequence"]
    B --> B2["Expression Trees<br/>Infix notation"]

    C --> C1["Tree Serialization<br/>Save tree structure"]
    C --> C2["Directory Listing<br/>Folder before contents"]

    D --> D1["Tree Deletion<br/>Children before parent"]
    D --> D2["Expression Trees<br/>Postfix notation"]
    D --> D3["Calculate Directory Size<br/>Subdirs before parent"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

Tree traversal refers to the process of visiting each node in a tree data structure exactly once in a systematic way.

## Tree Structure

```
      1
     / \
    2   3
   / \ / \
  4  5 6  7
```

**Traversal Results:**

- Pre-order: [1, 2, 4, 5, 3, 6, 7]
- In-order: [4, 2, 5, 1, 6, 3, 7]
- Post-order: [4, 5, 2, 6, 7, 3, 1]

## Complexity

### Recursive Implementations

- **Time Complexity**: O(n) - visits each node exactly once
- **Space Complexity**: O(h) - where h is the height of the tree (call stack)
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)

### Iterative Implementations

- **Time Complexity**: O(n) - visits each node exactly once
- **Space Complexity**: O(h) - explicit stack storage
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)

## Implementation Details

### Recursive Approach

- Natural and intuitive implementation
- Uses system call stack
- Prone to stack overflow for very deep trees
- Generally more readable and easier to understand

### Iterative Approach

- Uses explicit stack data structure
- More memory efficient for deep trees
- Avoids stack overflow issues
- Pre-order is straightforward, post-order is most complex

## Use Cases

### Pre-order Traversal

- Tree serialization and copying
- Prefix expression evaluation
- Directory listing (depth-first)
- Tree validation

### In-order Traversal

- Binary Search Tree (BST) sorted output
- Expression tree evaluation (infix notation)
- Tree flattening to sorted array

### Post-order Traversal

- Tree deletion (children before parent)
- Directory size calculation
- Postfix expression evaluation
- Dependency resolution

## Performance Characteristics

The iterative implementations generally have similar performance to recursive ones but offer better control over memory usage. Post-order iterative traversal is the most complex due to the need to track the last visited node.

## Usage

```bash
make run NAME=0014-tree-traversal
```

## Testing

```bash
make test NAME=0014-tree-traversal
```

## Benchmarking

```bash
go test -bench=. ./0014-tree-traversal/
```
