# LRU Cache

## Description

A complete LRU (Least Recently Used) cache implementation from scratch using a combination of a doubly-linked list and hash map. This provides O(1) time complexity for both get and put operations while maintaining the LRU eviction policy. When the cache reaches its capacity, the least recently used item is automatically evicted to make room for new entries.

## Visual Representation

### LRU Cache Architecture

```mermaid
graph TD
    subgraph "LRU Cache Structure"
        A[Hash Map] --> B[Key → Node Mapping]
        C[Doubly Linked List] --> D[Access Order Tracking]
    end

    subgraph "Memory Layout"
        E["head ↔ Node1 ↔ Node2 ↔ Node3 ↔ tail"]
        F["MRU (Most Recent) ← → LRU (Least Recent)"]
    end

    subgraph "Node Structure"
        G[key: string]
        H[value: interface{}]
        I[prev: *Node]
        J[next: *Node]
    end

    style A fill:#e1f5fe
    style E fill:#c8e6c9
    style G fill:#fff3e0
```

### Get Operation Flow

```mermaid
graph TD
    A[Get(key)] --> B{Key exists in hash map?}
    B -->|No| C[Return nil, false]
    B -->|Yes| D[Get node from hash map]
    D --> E[Remove node from current position]
    E --> F[Move node to head (most recent)]
    F --> G[Return node.value, true]

    subgraph "List Update"
        H[node.prev.next = node.next]
        I[node.next.prev = node.prev]
        J[Insert at head]
    end

    E --> H

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style C fill:#ffcdd2
```

### Put Operation Flow

```mermaid
graph TD
    A[Put(key, value)] --> B{Key exists?}
    B -->|Yes| C[Update existing node]
    B -->|No| D{Cache at capacity?}

    C --> E[Update value]
    E --> F[Move to head]

    D -->|No| G[Create new node]
    D -->|Yes| H[Remove LRU node]
    H --> I[Remove from hash map]
    I --> G

    G --> J[Add to hash map]
    J --> K[Insert at head]
    K --> L[Update size]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style K fill:#c8e6c9
    style H fill:#fff3e0
```

### LRU Eviction Process

```mermaid
graph LR
    subgraph "Before Eviction (Capacity: 3)"
        A1["head ↔ C ↔ B ↔ A ↔ tail"]
        A2["MRU: C, LRU: A"]
    end

    subgraph "Add new item D (triggers eviction)"
        B1["Remove A (LRU)"]
        B2["Add D at head"]
    end

    subgraph "After Eviction"
        C1["head ↔ D ↔ C ↔ B ↔ tail"]
        C2["MRU: D, LRU: B"]
    end

    A1 --> B1 --> C1

    style A2 fill:#ffcdd2
    style C2 fill:#c8e6c9
    style B1 fill:#fff3e0
```

### Access Pattern Visualization

```mermaid
graph TD
    A[Access Pattern Example] --> B[Initial: empty cache]
    B --> C["Put(A, 1): [A]"]
    C --> D["Put(B, 2): [B, A]"]
    D --> E["Put(C, 3): [C, B, A]"]
    E --> F["Get(A): [A, C, B] - A moves to front"]
    F --> G["Put(D, 4): [D, A, C] - B evicted"]
    G --> H["Get(C): [C, D, A] - C moves to front"]

    subgraph "Cache States"
        I["MRU → LRU order"]
        J["Capacity: 3"]
        K["Eviction: Remove tail"]
    end

    style B fill:#e1f5fe
    style H fill:#c8e6c9
    style G fill:#fff3e0
```

### Hash Map + Doubly Linked List Synergy

```mermaid
graph LR
    subgraph "Hash Map (O(1) lookup)"
        A["key1 → Node1"]
        B["key2 → Node2"]
        C["key3 → Node3"]
    end

    subgraph "Doubly Linked List (O(1) reorder)"
        D[head] --> E[Node3]
        E --> F[Node1]
        F --> G[Node2]
        G --> H[tail]
        E -.->|prev| D
        F -.->|prev| E
        G -.->|prev| F
        H -.->|prev| G
    end

    A -.-> F
    B -.-> G
    C -.-> E

    style A fill:#e1f5fe
    style E fill:#c8e6c9
```

### Cache Hit vs Miss Scenarios

```mermaid
graph TD
    A[Cache Access] --> B{Data in cache?}

    B -->|Yes - Cache Hit| C[Fast retrieval]
    C --> D[Move to MRU position]
    C --> E[Return cached data]

    B -->|No - Cache Miss| F[Expensive operation]
    F --> G[Fetch from source]
    G --> H[Store in cache]
    H --> I{Cache full?}
    I -->|Yes| J[Evict LRU item]
    I -->|No| K[Add to cache]
    J --> K
    K --> L[Return data]

    style C fill:#c8e6c9
    style F fill:#ffcdd2
    style J fill:#fff3e0
```

### LRU vs Other Eviction Policies

```mermaid
graph TD
    A[Cache Eviction Policies] --> B[LRU]
    A --> C[FIFO]
    A --> D[LFU]
    A --> E[Random]

    B --> B1["Evict: Least recently used"]
    B --> B2["Good: Temporal locality"]
    B --> B3["Complexity: O(1) with good design"]

    C --> C1["Evict: First inserted"]
    C --> C2["Simple: Queue-based"]
    C --> C3["Poor: Ignores access patterns"]

    D --> D1["Evict: Least frequently used"]
    D --> D2["Good: Frequency matters"]
    D --> D3["Complex: Counter maintenance"]

    E --> E1["Evict: Random selection"]
    E --> E2["Simple: No bookkeeping"]
    E --> E3["Poor: No pattern consideration"]

    style B2 fill:#c8e6c9
    style B3 fill:#c8e6c9
    style C3 fill:#ffcdd2
    style E3 fill:#ffcdd2
```

### Performance Analysis

```mermaid
graph LR
    subgraph "Time Complexity"
        A["Get: O(1)"]
        B["Put: O(1)"]
        C["Delete: O(1)"]
        D["All operations constant time"]
    end

    subgraph "Space Complexity"
        E["Storage: O(capacity)"]
        F["Hash map: O(n)"]
        G["Linked list: O(n)"]
        H["Total: O(capacity)"]
    end

    subgraph "Real-world Benefits"
        I["Predictable memory usage"]
        J["Fast access times"]
        K["Automatic cleanup"]
        L["Temporal locality exploitation"]
    end

    style A fill:#c8e6c9
    style B fill:#c8e6c9
    style C fill:#c8e6c9
    style I fill:#c8e6c9
```

## Key Features

- **O(1) Operations**: Both get and put operations run in constant time
- **Doubly-Linked List**: Maintains access order efficiently
- **Hash Map**: Provides fast key-to-node lookups
- **Automatic Eviction**: Removes least recently used items when capacity is exceeded
- **Generic Values**: Supports any value type with string keys
- **Complete API**: Standard cache operations plus utility methods

## Implementation Details

- **Data Structure**: Combination of hash map and doubly-linked list
- **Eviction Policy**: Least Recently Used (LRU)
- **Capacity Management**: Fixed capacity with automatic eviction
- **Access Order**: Most recent items at head, least recent at tail
- **Sentinel Nodes**: Dummy head and tail nodes simplify list operations

## Complexity

- **Time Complexity**:
  - Get: O(1) - Hash map lookup + list node movement
  - Put: O(1) - Hash map insertion + list operations
  - Delete: O(1) - Hash map removal + list node removal
  - All other operations: O(1) or O(n) for iteration-based methods
- **Space Complexity**: O(capacity) for the cache storage

## Core Operations

### Basic Operations

- `Put(key, value)` - Insert or update a key-value pair (moves to front)
- `Get(key)` - Retrieve value by key and mark as recently used
- `Delete(key)` - Remove key-value pair from cache
- `Has(key)` - Check if key exists (without affecting access order)
- `Size()` - Get number of stored pairs
- `Capacity()` - Get maximum capacity
- `IsEmpty()` - Check if cache is empty
- `IsFull()` - Check if cache is at capacity
- `Clear()` - Remove all key-value pairs

### Utility Operations

- `Keys()` - Get slice of all keys in MRU order
- `Values()` - Get slice of all values in MRU order
- `Entries()` - Get slice of all key-value pairs in MRU order
- `ForEach(func)` - Iterate over all pairs in MRU order
- `Peek(key)` - Get value without affecting access order
- `GetMostRecentKey()` - Get the most recently used key
- `GetLeastRecentKey()` - Get the least recently used key

### Advanced Operations

- `SetCapacity(newCapacity)` - Change cache capacity (evicts items if needed)

## Usage

```bash
make run n=lru-cache
```

## Testing

```bash
make test n=lru-cache
```

## Benchmarking

```bash
make bench n=lru-cache
```

## Example Usage in Go

```go
// Create a new LRU cache with capacity 3
cache := NewLRUCache(3)

// Insert key-value pairs
cache.Put("user1", "Alice")
cache.Put("user2", "Bob")
cache.Put("user3", "Charlie")

// Retrieve values (marks as recently used)
user, exists := cache.Get("user1")
if exists {
    fmt.Printf("User: %s\n", user)
}

// Check if cache is full
if cache.IsFull() {
    fmt.Println("Cache is at capacity")
}

// Add another item (will evict least recently used)
cache.Put("user4", "David")

// user2 will be evicted since user1 was accessed recently
_, exists = cache.Get("user2")
fmt.Printf("User2 exists: %t\n", exists) // false

// Peek at value without affecting order
value, exists := cache.Peek("user3")
fmt.Printf("User3 (peek): %s\n", value)

// Get most and least recently used keys
mostRecent, _ := cache.GetMostRecentKey()
leastRecent, _ := cache.GetLeastRecentKey()
fmt.Printf("Most recent: %s, Least recent: %s\n", mostRecent, leastRecent)

// Iterate over all entries (in MRU order)
cache.ForEach(func(key string, value any) {
    fmt.Printf("%s: %v\n", key, value)
})

// Get cache statistics
fmt.Printf("Size: %d, Capacity: %d\n", cache.Size(), cache.Capacity())
```

## LRU Eviction Policy

The LRU (Least Recently Used) policy works as follows:

1. **Access Order**: Items are ordered by recency of access
2. **Most Recent**: Newly added or accessed items move to the front
3. **Least Recent**: Items not accessed for the longest time move to the back
4. **Eviction**: When capacity is exceeded, the item at the back is removed
5. **Update**: Updating an existing key moves it to the front

## Performance Characteristics

### Time Complexity

- **Get Operation**: O(1) - Hash lookup + list node movement
- **Put Operation**: O(1) - Hash insertion + list operations
- **Delete Operation**: O(1) - Hash removal + list node removal
- **Eviction**: O(1) - Remove tail node and hash entry
- **Iteration**: O(n) - Must visit all stored elements

### Space Complexity

- **Storage**: O(capacity) - Fixed maximum memory usage
- **Overhead**: O(1) per item - Hash entry + list node pointers
- **Total**: O(capacity) - Predictable memory footprint

## Implementation Strategy

### Data Structures

1. **Hash Map**: Maps keys to doubly-linked list nodes for O(1) lookup
2. **Doubly-Linked List**: Maintains access order with O(1) insertion/removal
3. **Sentinel Nodes**: Dummy head and tail simplify edge cases

### Key Operations

1. **Get**: Hash lookup → Move node to head → Return value
2. **Put**: Check if exists → Update or create → Move to head → Evict if needed
3. **Eviction**: Remove tail node → Delete from hash map → Decrement size

## Use Cases

- **Web Caches**: Store frequently accessed web pages or API responses
- **Database Query Caches**: Cache expensive database query results
- **Image/Asset Caches**: Store processed images or static assets
- **Session Storage**: Maintain user session data with automatic cleanup
- **Computation Caches**: Store results of expensive calculations
- **Memory Management**: Implement page replacement in operating systems
- **CDN Systems**: Content delivery network edge caching
- **Application Caches**: General-purpose application-level caching

## Advantages

- **Optimal Performance**: O(1) for all basic operations
- **Memory Efficient**: Fixed memory usage, automatic cleanup
- **Temporal Locality**: Exploits the principle that recently used items are likely to be used again
- **Simple Interface**: Easy to use standard cache API
- **Predictable Behavior**: Clear eviction policy and capacity management
- **Thread-Safe Design**: Can be easily extended with synchronization

## Limitations

- **String Keys Only**: Current implementation limited to string keys
- **Fixed Capacity**: Cannot grow beyond initial capacity (though capacity can be changed)
- **No Persistence**: Data is lost when cache is cleared or program exits
- **Memory Overhead**: Requires additional memory for hash map and list pointers
- **No TTL**: Items don't expire based on time, only on access patterns

## Comparison with Alternatives

### vs. Other Eviction Policies

- **vs. FIFO**: LRU considers access patterns, FIFO only insertion order
- **vs. LFU**: LRU focuses on recency, LFU on frequency of access
- **vs. Random**: LRU is deterministic and exploits temporal locality

### vs. Other Implementations

- **vs. Array-based**: Doubly-linked list provides O(1) reordering vs O(n) for arrays
- **vs. Single Hash Map**: Adding list maintains access order efficiently
- **vs. Built-in Maps**: LRU provides automatic eviction and capacity management

## Real-World Applications

- **Redis**: Uses LRU as one of its eviction policies
- **CPU Caches**: Hardware caches often use LRU or LRU approximations
- **Operating Systems**: Page replacement algorithms (though often use LRU approximations)
- **Web Browsers**: Cache management for web pages and resources
- **Database Systems**: Buffer pool management and query result caching

## Visual Representation

### LRU Cache Architecture

```mermaid
graph TD
    subgraph "LRU Cache Structure"
        A[Hash Map] --> B[Key → Node Reference]
        C[Doubly Linked List] --> D[MRU ← → ← → LRU]

        B --> E["Key1 → Node1"]
        B --> F["Key2 → Node2"]
        B --> G["Key3 → Node3"]

        D --> H["Head ← → Node1 ← → Node2 ← → Node3 ← → Tail"]
    end

    style A fill:#e1f5fe
    style C fill:#f3e5f5
    style H fill:#c8e6c9
```

### Detailed Structure Visualization

```mermaid
graph LR
    subgraph "Hash Map"
        H1["'A' → Node_A"]
        H2["'B' → Node_B"]
        H3["'C' → Node_C"]
    end

    subgraph "Doubly Linked List (MRU → LRU)"
        Head[Head] --> A["Node_C<br/>Key:'C', Val:3"]
        A --> B["Node_A<br/>Key:'A', Val:1"]
        B --> C["Node_B<br/>Key:'B', Val:2"]
        C --> Tail[Tail]

        B -.->|prev| A
        C -.->|prev| B
        A -.->|prev| Head
    end

    H1 -.-> A
    H2 -.-> C
    H3 -.-> A

    style Head fill:#e1f5fe
    style Tail fill:#e1f5fe
    style A fill:#c8e6c9
```

### Get Operation Flow

```mermaid
graph TD
    A[Get(key)] --> B{Key exists in HashMap?}
    B -->|No| C[Return -1]
    B -->|Yes| D[Get node reference]
    D --> E[Remove node from current position]
    E --> F[Move node to head (MRU)]
    F --> G[Return node value]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style C fill:#ffcdd2
```

### Put Operation Flow

```mermaid
graph TD
    A[Put(key, value)] --> B{Key exists?}
    B -->|Yes| C[Update value]
    B -->|No| D{Cache at capacity?}

    C --> E[Move to head]
    C --> F[Operation complete]

    D -->|No| G[Create new node]
    D -->|Yes| H[Remove LRU node]

    H --> I[Remove from HashMap]
    I --> G
    G --> J[Add to head]
    J --> K[Add to HashMap]
    K --> L[Operation complete]

    E --> F

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style L fill:#c8e6c9
```

### LRU Eviction Example

```mermaid
graph LR
    subgraph "Before: Cache Full (capacity=3)"
        A1[Head] --> B1["C:3 (MRU)"]
        B1 --> C1["A:1"]
        C1 --> D1["B:2 (LRU)"]
        D1 --> E1[Tail]
    end

    subgraph "After: Put(D, 4)"
        A2[Head] --> B2["D:4 (MRU)"]
        B2 --> C2["C:3"]
        C2 --> D2["A:1"]
        D2 --> E2[Tail]
        F2["B:2 EVICTED"]
    end

    style B1 fill:#c8e6c9
    style D1 fill:#ffcdd2
    style B2 fill:#c8e6c9
    style F2 fill:#ffcdd2
```

### Access Pattern Example

```mermaid
graph TD
    A[Initial: Put A,B,C] --> B[Cache: C→A→B]
    B --> C[Get(A)] --> D[Cache: A→C→B]
    D --> E[Put(D)] --> F[Cache: D→A→C (B evicted)]
    F --> G[Get(C)] --> H[Cache: C→D→A]
    H --> I[Put(E)] --> J[Cache: E→C→D (A evicted)]

    style A fill:#e1f5fe
    style J fill:#c8e6c9
```

### Node Structure Detail

```mermaid
graph TD
    A[LRU Node] --> B[Key]
    A --> C[Value]
    A --> D[Prev Pointer]
    A --> E[Next Pointer]

    F[Example Node] --> G["Key: 'user123'"]
    F --> H["Value: UserData"]
    F --> I["Prev: Previous Node"]
    F --> J["Next: Next Node"]

    style A fill:#e1f5fe
    style F fill:#c8e6c9
```

### Time Complexity Analysis

```mermaid
graph TD
    A[LRU Cache Operations] --> B[Get Operation]
    A --> C[Put Operation]

    B --> B1["HashMap lookup: O(1)"]
    B --> B2["List update: O(1)"]
    B --> B3["Total: O(1)"]

    C --> C1["HashMap operations: O(1)"]
    C --> C2["List operations: O(1)"]
    C --> C3["Total: O(1)"]

    D[Space Complexity] --> E["O(capacity)"]

    style B3 fill:#c8e6c9
    style C3 fill:#c8e6c9
    style E fill:#c8e6c9
```

### LRU vs Other Eviction Policies

```mermaid
graph TD
    A[Cache Eviction Policies] --> B[LRU]
    A --> C[FIFO]
    A --> D[LFU]
    A --> E[Random]

    B --> B1["Evict least recently used<br/>Good temporal locality<br/>O(1) operations"]
    C --> C1["Evict first inserted<br/>Simple implementation<br/>Poor for access patterns"]
    D --> D1["Evict least frequently used<br/>Good for frequency patterns<br/>More complex tracking"]
    E --> E1["Evict random item<br/>Simple implementation<br/>Unpredictable performance"]

    style B1 fill:#c8e6c9
    style A fill:#e1f5fe
```

### Real-World Applications

```mermaid
graph TD
    A[LRU Cache Applications] --> B[CPU Caches]
    A --> C[Operating Systems]
    A --> D[Web Browsers]
    A --> E[Database Management]
    A --> F[CDN Systems]

    B --> B1["Page replacement<br/>Memory hierarchy"]
    C --> C1["Virtual memory<br/>Buffer management"]
    D --> D1["Browser cache<br/>Recently visited pages"]
    E --> E1["Buffer pools<br/>Query result caching"]
    F --> F1["Content distribution<br/>Edge caching"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
```
