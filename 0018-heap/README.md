# Heap Data Structure

## Description

Implementation of both Min-Heap and Max-Heap data structures with comparison to Go's built-in `container/heap` package. A heap is a complete binary tree that satisfies the heap property: in a min-heap, each parent node is smaller than or equal to its children; in a max-heap, each parent node is greater than or equal to its children.

## Visual Representation

### Min-Heap Structure

```mermaid
graph TD
    A[1] --> B[3]
    A --> C[2]
    B --> D[7]
    B --> E[5]
    C --> F[4]
    C --> G[6]
    D --> H[10]
    D --> I[8]
    E --> J[9]

    style A fill:#c8e6c9
    style B fill:#e8f5e8
    style C fill:#e8f5e8
    style D fill:#f1f8e9
    style E fill:#f1f8e9
    style F fill:#f1f8e9
    style G fill:#f1f8e9
```

### Array Representation

```mermaid
graph LR
    subgraph "Index"
        I0[0] --- I1[1] --- I2[2] --- I3[3] --- I4[4] --- I5[5] --- I6[6] --- I7[7] --- I8[8] --- I9[9]
    end

    subgraph "Value"
        V0[1] --- V1[3] --- V2[2] --- V3[7] --- V4[5] --- V5[4] --- V6[6] --- V7[10] --- V8[8] --- V9[9]
    end

    subgraph "Relationships"
        R1["Parent(i) = (i-1)/2"]
        R2["Left(i) = 2*i + 1"]
        R3["Right(i) = 2*i + 2"]
    end

    style V0 fill:#c8e6c9
```

### Insert Operation (Heapify Up)

```mermaid
graph TD
    A[Insert element at end] --> B[Compare with parent]
    B --> C{Parent > child?}
    C -->|Yes| D[Swap with parent]
    C -->|No| E[Done]
    D --> F[Move up to parent position]
    F --> B

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Extract Min Operation (Heapify Down)

```mermaid
graph TD
    A[Save root value] --> B[Move last element to root]
    B --> C[Remove last element]
    C --> D[Compare with children]
    D --> E{Has smaller child?}
    E -->|Yes| F[Swap with smallest child]
    E -->|No| G[Done]
    F --> H[Move down to child position]
    H --> D
    G --> I[Return saved value]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Heap Sort Visualization

```mermaid
graph LR
    subgraph "Step 1: Build Max Heap"
        A1[Unsorted Array] --> A2[Max Heap]
    end

    subgraph "Step 2: Extract Max Repeatedly"
        B1[Extract Max] --> B2[Move to end] --> B3[Heapify remaining]
    end

    subgraph "Result"
        C1[Sorted Array]
    end

    A2 --> B1
    B3 --> B1
    B3 --> C1

    style A1 fill:#ffcdd2
    style C1 fill:#c8e6c9
```

### Complexity Comparison

```mermaid
graph TD
    subgraph "Operations"
        A["Insert: O(log n)"]
        B["Extract: O(log n)"]
        C["Peek: O(1)"]
        D["Build: O(n)"]
    end

    subgraph "Tree Height"
        E["Height = log n"]
        F[Complete binary tree]
    end

    A -.-> E
    B -.-> E

    style C fill:#c8e6c9
    style D fill:#c8e6c9
```

## Features

- **MinHeap**: Extract minimum element efficiently
- **MaxHeap**: Extract maximum element efficiently
- **Priority Queue Operations**: Insert, extract, peek
- **Heap Sort**: Sorting algorithm using heap structure
- **Go Heap Comparison**: Side-by-side comparison with standard library
- **Complete Implementation**: All fundamental heap operations

## Heap Properties

- **Complete Binary Tree**: All levels filled except possibly the last
- **Heap Property**: Parent-child relationship maintains order
- **Array Representation**: Efficient storage using slice/array
- **Index Relationships**: Parent at `(i-1)/2`, children at `2*i+1` and `2*i+2`

## Operations

### Insert (Heapify Up)

- Add element at end of array
- Bubble up to maintain heap property
- **Time**: O(log n), **Space**: O(1)

### Extract Min/Max (Heapify Down)

- Remove root element
- Replace with last element
- Bubble down to maintain heap property
- **Time**: O(log n), **Space**: O(1)

### Peek

- Return root element without removing
- **Time**: O(1), **Space**: O(1)

### Build Heap

- Convert array to heap structure
- **Time**: O(n), **Space**: O(1)

## Complexity

### Time Complexity

- **Insert**: O(log n)
- **Extract Min/Max**: O(log n)
- **Peek**: O(1)
- **Build Heap**: O(n)
- **Heap Sort**: O(n log n)

### Space Complexity

- **Storage**: O(n) for n elements
- **Operations**: O(1) auxiliary space
- **Recursive**: O(log n) if using recursive heapify

## Comparison: Raw vs Go's Heap

### Raw Implementation

```go
minHeap := NewMinHeap()
minHeap.Insert(10)
min, _ := minHeap.ExtractMin()
```

### Go's container/heap

```go
h := &IntHeap{}
heap.Init(h)
heap.Push(h, 10)
min := heap.Pop(h).(int)
```

### Key Differences

- **Raw**: Direct method calls, type-safe
- **Go Heap**: Interface-based, requires type assertions
- **Performance**: Similar O(log n) complexity
- **Flexibility**: Go heap works with any type implementing heap.Interface

## Use Cases

### Priority Queues

- Task scheduling with priorities
- Dijkstra's shortest path algorithm
- A\* pathfinding algorithm

### Sorting

- Heap sort algorithm
- Finding k largest/smallest elements
- Median maintenance

### Real-Time Systems

- Event scheduling
- Resource allocation
- Load balancing

### Graph Algorithms

- Minimum spanning tree (Prim's algorithm)
- Shortest path algorithms
- Network routing protocols

## Heap Sort Algorithm

Uses max-heap to sort in ascending order:

1. Build max-heap from input array
2. Repeatedly extract maximum
3. Result is sorted in descending order

## Usage

```bash
make run NAME=0018-heap
```

## Testing

```bash
make test NAME=0018-heap
```
