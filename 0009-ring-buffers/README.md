# Ring Buffers (Circular Buffers)

## Description

A ring buffer, also known as a circular buffer, is a fixed-size data structure that efficiently implements a FIFO (First-In-First-Out) queue. The key characteristic is its circular nature - when the buffer reaches its capacity, new elements wrap around to the beginning, overwriting the oldest data if the buffer is full.

This implementation provides:

- **Generic support** for any data type using Go generics
- **Fixed capacity** with automatic wrap-around behavior
- **O(1) operations** for all basic operations
- **Memory efficient** by reusing the same buffer space
- **Thread-unsafe** design for maximum performance (external synchronization required if used concurrently)

## Key Features

- `Enqueue(item)` - Add item to the tail (returns error if full)
- `Dequeue()` - Remove and return item from head (returns error if empty)
- `Peek()` - View head item without removing it
- `IsEmpty()` / `IsFull()` - Check buffer state
- `Size()` / `Capacity()` - Get current size and maximum capacity
- `Clear()` - Reset buffer to empty state
- `ToSlice()` - Convert buffer contents to a slice in FIFO order

## Complexity

- **Time Complexity**: O(1) for all operations (enqueue, dequeue, peek, isEmpty, isFull)
- **Space Complexity**: O(n) where n is the fixed capacity
- **Access Pattern**: Sequential access with circular wrapping

## Algorithm Details

The ring buffer uses two pointers:

- **Head**: Points to the next item to dequeue
- **Tail**: Points to the position for the next enqueue
- **Size counter**: Tracks current number of elements

Circular behavior is achieved using modulo arithmetic: `(index + 1) % capacity`

## Usage

```bash
make run NAME=0009-ring-buffers
make test NAME=0009-ring-buffers
```

### Code Example

```go
// Create ring buffer with capacity 3
rb := NewRingBuffer[int](3)

// Fill the buffer
rb.Enqueue(1) // [1, _, _] head=0, tail=1
rb.Enqueue(2) // [1, 2, _] head=0, tail=2
rb.Enqueue(3) // [1, 2, 3] head=0, tail=0 (wrapped)

// Buffer is now full
rb.IsFull() // true

// Dequeue and add new item (demonstrates circular wrapping)
first, _ := rb.Dequeue() // first=1, [_, 2, 3] head=1, tail=0
rb.Enqueue(4)            // [4, 2, 3] head=1, tail=1

// Current contents in FIFO order: [2, 3, 4]
contents := rb.ToSlice() // [2, 3, 4]
```

## Applications

### 1. **Producer-Consumer Scenarios**

```go
rb := NewRingBuffer[Task](1000)
// Producer adds tasks, consumer processes them
```

### 2. **Streaming Data Buffers**

```go
rb := NewRingBuffer[AudioSample](4096)
// Audio processing with fixed-size circular buffer
```

### 3. **Network Packet Buffers**

```go
rb := NewRingBuffer[Packet](512)
// Network driver buffering incoming packets
```

### 4. **Real-time Systems**

```go
rb := NewRingBuffer[SensorReading](100)
// Keep latest 100 sensor readings, automatically drop oldest
```

### 5. **Rate Limiting / Sliding Window**

```go
rb := NewRingBuffer[time.Time](60)
// Track timestamps for rate limiting (60 requests per minute)
```

## Advantages

- **Constant time operations** - All operations are O(1)
- **Memory efficient** - Fixed memory allocation, no dynamic resizing
- **Cache friendly** - Sequential access pattern with good locality
- **Predictable performance** - No garbage collection pressure from frequent allocations
- **Simple implementation** - Straightforward logic with minimal overhead

## Trade-offs

- **Fixed capacity** - Cannot grow beyond initial size
- **Error handling required** - Operations can fail when full/empty
- **Thread safety** - Requires external synchronization for concurrent access
- **Memory usage** - Always allocates full capacity regardless of current size

## Comparison with Other Structures

| Operation | Ring Buffer | Slice Queue | Linked Queue |
| --------- | ----------- | ----------- | ------------ |
| Enqueue   | O(1)        | O(1)\*      | O(1)         |
| Dequeue   | O(1)        | O(n)        | O(1)         |
| Memory    | Fixed       | Growing     | Dynamic      |
| Cache     | Excellent   | Good        | Poor         |

\*Slice queue enqueue is O(n) when resizing occurs

## Testing

```bash
make test NAME=0009-ring-buffers
```

The test suite covers:

- Basic enqueue/dequeue operations
- Circular wrapping behavior
- Overflow/underflow error conditions
- Edge cases (empty, full, single element)
- Generic type support
- Performance benchmarks

## Visual Representation

### Ring Buffer Structure

```mermaid
graph LR
    subgraph "Ring Buffer (Capacity: 6)"
        A[0] --> B[1]
        B --> C[2]
        C --> D[3]
        D --> E[4]
        E --> F[5]
        F --> A
    end

    subgraph "Pointers"
        Head[Head: 1]
        Tail[Tail: 4]
        Size[Size: 3]
    end

    style A fill:#f5f5f5
    style B fill:#c8e6c9
    style C fill:#c8e6c9
    style D fill:#c8e6c9
    style E fill:#f5f5f5
    style F fill:#f5f5f5
```

### Ring Buffer States

```mermaid
graph TD
    subgraph "Empty Buffer"
        A1[Head = Tail = 0, Size = 0]
        B1["[_, _, _, _, _, _]"]
    end

    subgraph "Partially Filled"
        A2[Head = 1, Tail = 4, Size = 3]
        B2["[_, X, X, X, _, _]"]
    end

    subgraph "Full Buffer"
        A3[Head = 0, Tail = 0, Size = 6]
        B3["[X, X, X, X, X, X]"]
    end

    subgraph "Wrapped Around"
        A4[Head = 4, Tail = 2, Size = 4]
        B4["[X, X, _, _, X, X]"]
    end

    style A1 fill:#f5f5f5
    style A2 fill:#c8e6c9
    style A3 fill:#fff3e0
    style A4 fill:#e1f5fe
```

### Enqueue Operation

```mermaid
graph TD
    A[Enqueue Element] --> B{Is buffer full?}
    B -->|Yes| C[Return error/overwrite]
    B -->|No| D[Insert at tail position]
    D --> E[Increment tail pointer]
    E --> F[tail = (tail + 1) % capacity]
    F --> G[Increment size]
    G --> H[Operation complete]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style C fill:#ffcdd2
```

### Dequeue Operation

```mermaid
graph TD
    A[Dequeue Element] --> B{Is buffer empty?}
    B -->|Yes| C[Return error]
    B -->|No| D[Get element at head]
    D --> E[Increment head pointer]
    E --> F[head = (head + 1) % capacity]
    F --> G[Decrement size]
    G --> H[Return element]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style C fill:#ffcdd2
```

### Circular Indexing Visualization

```mermaid
graph LR
    subgraph "Linear Array View"
        L0[0] --- L1[1] --- L2[2] --- L3[3] --- L4[4] --- L5[5]
    end

    subgraph "Circular View"
        C0[0] --> C1[1]
        C1 --> C2[2]
        C2 --> C3[3]
        C3 --> C4[4]
        C4 --> C5[5]
        C5 --> C0
    end

    subgraph "Modulo Operation"
        M1["Index = (head + offset) % capacity"]
        M2["Next = (current + 1) % capacity"]
    end

    style C0 fill:#e1f5fe
    style M1 fill:#c8e6c9
```

### Buffer Operations Example

```mermaid
graph TD
    A[Initial: Empty Buffer] --> B[Enqueue: A, B, C]
    B --> C["Buffer: [A, B, C, _, _, _]<br/>Head: 0, Tail: 3"]
    C --> D[Dequeue: 2 elements]
    D --> E["Buffer: [_, _, C, _, _, _]<br/>Head: 2, Tail: 3"]
    E --> F[Enqueue: D, E, F, G, H]
    F --> G["Buffer: [G, H, C, D, E, F]<br/>Head: 2, Tail: 2 (wrapped)"]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
```

### Full vs Empty Distinction

```mermaid
graph LR
    subgraph "Problem: Same Pointers"
        A1["Empty: head = tail = 0"]
        A2["Full: head = tail = 0"]
    end

    subgraph "Solution 1: Size Counter"
        B1["Track size separately"]
        B2["Empty: size = 0"]
        B3["Full: size = capacity"]
    end

    subgraph "Solution 2: Waste One Slot"
        C1["Full: (tail + 1) % cap = head"]
        C2["Empty: tail = head"]
        C3["Capacity reduced by 1"]
    end

    style A1 fill:#ffcdd2
    style B1 fill:#c8e6c9
    style C1 fill:#fff3e0
```

### Performance Comparison

```mermaid
graph TD
    A[Ring Buffer vs Alternatives] --> B[Ring Buffer]
    A --> C[Dynamic Queue]
    A --> D[Fixed Array Queue]

    B --> B1["Enqueue: O(1)<br/>Dequeue: O(1)<br/>Space: O(n) fixed<br/>Memory: Excellent locality"]

    C --> C1["Enqueue: O(1) amortized<br/>Dequeue: O(1)<br/>Space: O(n) variable<br/>Memory: Allocation overhead"]

    D --> D1["Enqueue: O(n) shift<br/>Dequeue: O(n) shift<br/>Space: O(n) fixed<br/>Memory: Good locality"]

    style B1 fill:#c8e6c9
    style A fill:#e1f5fe
```

### Real-World Applications

```mermaid
graph TD
    A[Ring Buffer Applications] --> B[Audio/Video Streaming]
    A --> C[Hardware Interfaces]
    A --> D[Producer-Consumer]
    A --> E[Network Buffers]
    A --> F[Logging Systems]

    B --> B1["Continuous data flow<br/>Real-time processing<br/>Buffer underrun/overrun"]
    C --> C1["UART buffers<br/>Keyboard input<br/>Interrupt handlers"]
    D --> D1["Threading systems<br/>Pipeline processing<br/>Rate limiting"]
    E --> E1["TCP/UDP buffers<br/>Packet queuing<br/>Flow control"]
    F --> F1["Circular log files<br/>Recent event tracking<br/>Memory-bounded logging"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```

### Thread Safety Considerations

```mermaid
graph TD
    A[Ring Buffer Concurrency] --> B[Single Producer/Consumer]
    A --> C[Multiple Producers/Consumers]

    B --> B1["Lock-free possible<br/>Atomic head/tail updates<br/>Memory barriers needed"]

    C --> C1["Locks required<br/>Mutex protection<br/>Condition variables"]

    D[Common Issues] --> E[ABA Problem]
    D --> F[Memory Ordering]
    D --> G[Cache Line Contention]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#fff3e0
```
