# Queue

## Description

A Queue is a linear data structure that follows the First In, First Out (FIFO) principle. Elements are added (enqueued) at the rear and removed (dequeued) from the front.

## Visual Representation

### Queue Structure

```mermaid
graph LR
    subgraph "Queue (FIFO)"
        Front["Front<br/>Dequeue"] --> A[10]
        A --> B[20]
        B --> C[30]
        C --> D[40]
        D --> Rear["Rear<br/>Enqueue"]
    end

    style Front fill:#ffcdd2
    style Rear fill:#c8e6c9
    style A fill:#f3e5f5
    style B fill:#f3e5f5
    style C fill:#f3e5f5
    style D fill:#f3e5f5
```

### Enqueue Operation

```mermaid
graph TD
    A[New Element: 50] --> B[Check if queue is full]
    B --> C{Is Full?}
    C -->|Yes| D[Return Error]
    C -->|No| E[Add element at rear]
    E --> F[Increment rear pointer]
    F --> G[Increment size]
    G --> H[Operation Complete]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style D fill:#ffcdd2
```

### Dequeue Operation

```mermaid
graph TD
    A[Dequeue Request] --> B[Check if queue is empty]
    B --> C{Is Empty?}
    C -->|Yes| D[Return Error]
    C -->|No| E[Get element from front]
    E --> F[Increment front pointer]
    F --> G[Decrement size]
    G --> H[Return element]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style D fill:#ffcdd2
```

### Queue Operations Flow

```mermaid
graph LR
    subgraph "Initial State"
        Q1[Empty Queue]
    end

    subgraph "After Enqueue(10, 20, 30)"
        Q2["[10, 20, 30]<br/>Front → 10, Rear → 30"]
    end

    subgraph "After Dequeue()"
        Q3["[20, 30]<br/>Front → 20, Rear → 30"]
    end

    subgraph "After Enqueue(40)"
        Q4["[20, 30, 40]<br/>Front → 20, Rear → 40"]
    end

    Q1 --> Q2
    Q2 --> Q3
    Q3 --> Q4

    style Q1 fill:#f5f5f5
    style Q4 fill:#e8f5e8
```

### Implementation Types

```mermaid
graph TD
    A[Queue Implementations] --> B[Array-based]
    A --> C[Linked List-based]
    A --> D[Circular Buffer]

    B --> B1["Fixed size<br/>Fast access<br/>Memory waste"]
    C --> C1["Dynamic size<br/>Extra memory overhead<br/>Pointer management"]
    D --> D1["Fixed size<br/>Memory efficient<br/>Circular indexing"]

    style A fill:#e1f5fe
    style B1 fill:#fff3e0
    style C1 fill:#fff3e0
    style D1 fill:#fff3e0
```

This implementation provides three different queue variants:

1. **LinkedListQueue**: Uses a linked list with separate front and rear pointers
2. **ArrayQueue**: Uses a circular buffer with fixed capacity
3. **DynamicQueue**: Uses Go's slice with dynamic resizing

## Key Operations

- **Enqueue**: Add element to the rear of the queue
- **Dequeue**: Remove and return element from the front of the queue
- **Front**: Get the front element without removing it
- **Rear**: Get the rear element without removing it
- **Size**: Get the number of elements in the queue
- **IsEmpty**: Check if the queue is empty
- **Clear**: Remove all elements from the queue

## Complexity

### LinkedListQueue

- **Enqueue**: O(1) - Constant time insertion at rear
- **Dequeue**: O(1) - Constant time removal from front
- **Front/Rear**: O(1) - Direct access to front/rear pointers
- **Space**: O(n) - Linear space for n elements

### ArrayQueue (Circular Buffer)

- **Enqueue**: O(1) - Constant time insertion (when not full)
- **Dequeue**: O(1) - Constant time removal
- **Front/Rear**: O(1) - Direct array access
- **Space**: O(capacity) - Fixed space based on capacity

### DynamicQueue

- **Enqueue**: O(1) amortized - May require slice expansion
- **Dequeue**: O(n) - Requires shifting all elements
- **Front/Rear**: O(1) - Direct slice access
- **Space**: O(n) - Dynamic space allocation

## Implementation Details

### LinkedListQueue

Uses a doubly-pointed linked list where:

- `front` points to the first node (for dequeue operations)
- `rear` points to the last node (for enqueue operations)
- Maintains a `size` counter for O(1) size queries

### ArrayQueue

Uses a circular buffer approach where:

- `front` and `rear` indices wrap around using modulo arithmetic
- Avoids shifting elements by reusing array positions
- Has a fixed capacity to prevent unbounded growth

### DynamicQueue

Uses Go's built-in slice operations:

- Appends to the end for enqueue (O(1) amortized)
- Uses slice re-slicing for dequeue (O(n) due to shifting)
- Automatically handles memory allocation

## Performance Comparison

| Operation | LinkedList | Array (Circular) | Dynamic  |
| --------- | ---------- | ---------------- | -------- |
| Enqueue   | O(1)       | O(1)             | O(1)\*   |
| Dequeue   | O(1)       | O(1)             | O(n)     |
| Front     | O(1)       | O(1)             | O(1)     |
| Rear      | O(1)       | O(1)             | O(1)     |
| Memory    | Variable   | Fixed            | Variable |

\*Amortized complexity

## Real-World Applications

### Task Scheduling

```go
taskQueue := NewLinkedListQueue()
taskQueue.Enqueue(task1)
taskQueue.Enqueue(task2)
nextTask, _ := taskQueue.Dequeue()
```

### Breadth-First Search (BFS)

```go
bfsQueue := NewLinkedListQueue()
bfsQueue.Enqueue(startNode)
for !bfsQueue.IsEmpty() {
    node, _ := bfsQueue.Dequeue()
    // Process node and add neighbors
}
```

### Print Queue Management

```go
printQueue := NewArrayQueue(100)
printQueue.Enqueue(document1)
printQueue.Enqueue(document2)
nextDoc, _ := printQueue.Dequeue()
```

### Buffer Management

```go
buffer := NewDynamicQueue()
for data := range inputStream {
    buffer.Enqueue(data)
    if buffer.Size() > threshold {
        processedData, _ := buffer.Dequeue()
        // Process data
    }
}
```

## When to Use Each Implementation

### LinkedListQueue

- **Best for**: General-purpose queuing, unlimited size requirements
- **Pros**: No capacity limit, O(1) all operations, memory efficient
- **Cons**: Extra memory overhead for pointers

### ArrayQueue

- **Best for**: High-performance scenarios with known capacity limits
- **Pros**: Cache-friendly, minimal memory overhead, fastest operations
- **Cons**: Fixed capacity, potential for queue full errors

### DynamicQueue

- **Best for**: Simple implementations, variable workloads
- **Pros**: Automatic resizing, simple to understand
- **Cons**: O(n) dequeue operation, potential memory overhead

## Usage

```bash
make run NAME=0007-queue
```

## Testing

```bash
make test NAME=0007-queue
```

## Advanced Features

- **Display**: Visual representation of queue contents
- **ToSlice**: Convert queue to slice for inspection
- **Circular Buffer**: Efficient array utilization in ArrayQueue
- **Error Handling**: Proper error reporting for edge cases
- **Benchmarking**: Performance comparison between implementations
