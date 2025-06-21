# Same Tree Problem

## Description

The Same Tree Problem determines if two binary trees are structurally identical and have the same values at corresponding nodes. This fundamental tree problem demonstrates recursive thinking and tree comparison algorithms.

## Visual Representation

### Identical Trees Example

```mermaid
graph TD
    subgraph "Tree A"
        A1[1] --> A2[2]
        A1 --> A3[3]
        A2 --> A4[4]
        A2 --> A5[5]
    end

    subgraph "Tree B"
        B1[1] --> B2[2]
        B1 --> B3[3]
        B2 --> B4[4]
        B2 --> B5[5]
    end

    C["Trees A and B are IDENTICAL ✓"]

    style A1 fill:#c8e6c9
    style B1 fill:#c8e6c9
    style C fill:#c8e6c9
```

### Different Trees Examples

```mermaid
graph LR
    subgraph "Case 1: Different Values"
        A1[1] --> A2[2]
        A1 --> A3[3]

        B1[1] --> B2[2]
        B1 --> B3[4]

        C1["Different values at same position ✗"]
    end

    subgraph "Case 2: Different Structure"
        D1[1] --> D2[2]
        D1 --> D3[3]
        D2 --> D4[4]

        E1[1] --> E2[2]
        E1 --> E3[3]
        E3 --> E4[4]

        C2["Different structure ✗"]
    end

    style A3 fill:#ffcdd2
    style B3 fill:#ffcdd2
    style C1 fill:#ffcdd2
    style C2 fill:#ffcdd2
```

### Recursive Algorithm Flow

```mermaid
graph TD
    A["isSameTree(tree1, tree2)"] --> B{"Both trees null?"}
    B -->|Yes| C["Return true"]
    B -->|No| D{"One tree null?"}
    D -->|Yes| E["Return false"]
    D -->|No| F{"Values equal?"}
    F -->|No| E
    F -->|Yes| G["Check left subtrees"]
    G --> H["isSameTree(tree1.left, tree2.left)"]
    H --> I{"Left subtrees same?"}
    I -->|No| E
    I -->|Yes| J["Check right subtrees"]
    J --> K["isSameTree(tree1.right, tree2.right)"]
    K --> L["Return result"]

    style A fill:#e1f5fe
    style C fill:#c8e6c9
    style L fill:#c8e6c9
    style E fill:#ffcdd2
```

### Step-by-Step Comparison

```mermaid
graph TD
    subgraph "Comparison Steps"
        A["Step 1: Compare roots (1 == 1) ✓"]
        B["Step 2: Compare left children (2 == 2) ✓"]
        C["Step 3: Compare left-left children (4 == 4) ✓"]
        D["Step 4: Compare left-right children (5 == 5) ✓"]
        E["Step 5: Compare right children (3 == 3) ✓"]
        F["Result: Trees are identical ✓"]
    end

    A --> B --> C --> D --> E --> F

    style A fill:#e1f5fe
    style F fill:#c8e6c9
```

### Base Cases Visualization

```mermaid
graph LR
    subgraph "Base Case 1: Both Null"
        A1[null]
        A2[null]
        A3["return true"]
    end

    subgraph "Base Case 2: One Null"
        B1[node]
        B2[null]
        B3["return false"]
    end

    subgraph "Base Case 3: Different Values"
        C1[5]
        C2[3]
        C3["return false"]
    end

    style A3 fill:#c8e6c9
    style B3 fill:#ffcdd2
    style C3 fill:#ffcdd2
```

### Iterative Implementation

```mermaid
graph TD
    A[Iterative Approach] --> B[Use stack/queue for traversal]
    B --> C[Push both root nodes]
    C --> D{Stack empty?}
    D -->|Yes| E[Return true]
    D -->|No| F[Pop two nodes]
    F --> G{Both null?}
    G -->|Yes| D
    G -->|No| H{One null or values different?}
    H -->|Yes| I[Return false]
    H -->|No| J[Push children pairs]
    J --> D

    style A fill:#e1f5fe
    style E fill:#c8e6c9
    style I fill:#ffcdd2
```

### Complexity Analysis

```mermaid
graph TD
    A[Same Tree Complexity] --> B[Time Complexity]
    A --> C[Space Complexity]

    B --> B1["Best case: O(1) - different roots"]
    B --> B2["Worst case: O(min(n, m))"]
    B --> B3["Average case: O(min(n, m))"]

    C --> C1["Recursive: O(min(h1, h2)) stack space"]
    C --> C2["Iterative: O(min(h1, h2)) stack/queue"]
    C --> C3["where h = height of tree"]

    D[Notes] --> E["n, m = number of nodes in trees"]
    D --> F["Algorithm stops at first difference"]
    D --> G["Must visit all nodes if trees identical"]

    style A fill:#e1f5fe
    style B2 fill:#c8e6c9
    style C1 fill:#c8e6c9
```

### Variations of the Problem

```mermaid
graph TD
    A[Tree Comparison Variations] --> B[Same Tree]
    A --> C[Symmetric Tree]
    A --> D[Subtree Check]
    A --> E[Tree Isomorphism]

    B --> B1["Exact structural match<br/>Same values at same positions"]

    C --> C1["Tree is mirror of itself<br/>Left subtree mirrors right"]

    D --> D1["Check if one tree is subtree<br/>Can appear anywhere in larger tree"]

    E --> E1["Trees have same structure<br/>Values may be different"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
```

### Edge Cases

```mermaid
graph LR
    subgraph "Edge Cases to Consider"
        A["Empty trees (both null)"]
        B["Single node trees"]
        C["One empty, one non-empty"]
        D["Same structure, different values"]
        E["Different structure, same values"]
        F["Deep vs shallow trees"]
    end

    subgraph "Expected Results"
        G["true"]
        H["depends on values"]
        I["false"]
        J["false"]
        K["false"]
        L["false"]
    end

    A -.-> G
    B -.-> H
    C -.-> I
    D -.-> J
    E -.-> K
    F -.-> L

    style G fill:#c8e6c9
    style H fill:#fff3e0
    style I fill:#ffcdd2
    style J fill:#ffcdd2
    style K fill:#ffcdd2
    style L fill:#ffcdd2
```

### Application Examples

```mermaid
graph TD
    A[Same Tree Applications] --> B[Version Control]
    A --> C[Database Indexing]
    A --> D[Compiler Design]
    A --> E[File System]

    B --> B1["Compare code trees<br/>Detect changes<br/>Merge conflicts"]

    C --> C1["Index comparison<br/>B-tree equality<br/>Optimization checks"]

    D --> D1["AST comparison<br/>Code optimization<br/>Semantic analysis"]

    E --> E1["Directory comparison<br/>Backup verification<br/>Sync operations"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Optimizations

```mermaid
graph TD
    A[Optimization Strategies] --> B[Early Termination]
    A --> C[Size Comparison]
    A --> D[Hash Comparison]

    B --> B1["Stop at first difference<br/>Return false immediately<br/>Don't check remaining nodes"]

    C --> C1["Compare tree sizes first<br/>Different sizes = not same<br/>Avoid unnecessary traversal"]

    D --> D1["Compute tree hashes<br/>Different hashes = not same<br/>Same hash = need deep check"]

    E[Advanced Techniques] --> F["Canonical form comparison"]
    E --> G["Serialization comparison"]
    E --> H["Level-order comparison"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

The Same Tree Problem determines if two binary trees are structurally identical and have the same values at corresponding nodes.

## Problem Statement

Given two binary trees, determine if they are the same tree. Two binary trees are considered the same if:

1. They have the same structure
2. Nodes at corresponding positions have the same values

## Tree Examples

### Same Trees

```
Tree 1:    Tree 2:
   1          1
  / \        / \
 2   3      2   3
```

### Different Trees

```
Tree 1:    Tree 2:
   1          1
  / \        / \
 2   3      2   4
```

## Algorithm

The solution uses a simple recursive approach:

1. If both trees are null, they are the same
2. If one tree is null and the other isn't, they are different
3. If the values at current nodes differ, they are different
4. Recursively check left and right subtrees

## Complexity

- **Time Complexity**: O(min(m, n)) where m and n are the number of nodes
- **Space Complexity**: O(h) where h is the height of the tree (recursion stack)

## Usage

```bash
make run NAME=0016-two-binary-trees-comparison
```

## Testing

```bash
make test NAME=0016-two-binary-trees-comparison
```
