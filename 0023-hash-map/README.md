# Hash Map

## Description

A complete hash map (hash table) implementation from scratch using separate chaining for collision resolution. This implementation provides O(1) average-case performance for insertion, lookup, and deletion operations with automatic resizing to maintain optimal load factors.

## Visual Representation

### Hash Map Structure

```mermaid
graph LR
    subgraph "Hash Map with Chaining"
        A["Key: 'apple'"] --> H1["Hash Function"] --> I1["Index: 2"]
        B["Key: 'banana'"] --> H2["Hash Function"] --> I2["Index: 5"]
        C["Key: 'cherry'"] --> H3["Hash Function"] --> I3["Index: 2"]
    end

    subgraph "Bucket Array"
        B0["0: []"]
        B1["1: []"]
        B2["2: apple→cherry"]
        B3["3: []"]
        B4["4: []"]
        B5["5: banana"]
        B6["6: []"]
        B7["7: []"]
    end

    I1 -.-> B2
    I2 -.-> B5
    I3 -.-> B2

    style A fill:#e1f5fe
    style B2 fill:#fff3e0
    style B5 fill:#c8e6c9
```

### Hash Function Process (FNV-1a)

```mermaid
graph TD
    A["Input: key string"] --> B["Initialize hash = FNV_OFFSET"]
    B --> C["For each byte in key"]
    C --> D["hash = hash XOR byte"]
    D --> E["hash = hash * FNV_PRIME"]
    E --> F{"More bytes?"}
    F -->|Yes| C
    F -->|No| G["hash = hash % capacity"]
    G --> H["Return bucket index"]

    subgraph "FNV-1a Constants"
        I["FNV_OFFSET = 2166136261"]
        J["FNV_PRIME = 16777619"]
    end

    style A fill:#e1f5fe
    style H fill:#c8e6c9
```

### Collision Resolution: Separate Chaining

```mermaid
graph LR
    subgraph "No Collision"
        A1["Key: 'apple'"] --> B1["Hash: 2"] --> C1["Bucket 2: [apple]"]
    end

    subgraph "Collision Occurs"
        A2["Key: 'cherry'"] --> B2["Hash: 2"] --> C2["Bucket 2: [apple→cherry]"]
    end

    subgraph "Chain Traversal"
        D["Search 'cherry'"] --> E["Go to bucket 2"]
        E --> F["Check 'apple' ≠ 'cherry'"]
        F --> G["Follow next pointer"]
        G --> H["Check 'cherry' = 'cherry' ✓"]
    end

    style C1 fill:#c8e6c9
    style C2 fill:#fff3e0
    style H fill:#c8e6c9
```

### Dynamic Resizing Process

```mermaid
graph TD
    A["Insert new key-value"] --> B["Calculate load factor"]
    B --> C{"Load factor > 0.75?"}
    C -->|No| D["Insert into current table"]
    C -->|Yes| E["Trigger resize"]
    E --> F["Create new table: 2x capacity"]
    F --> G["Rehash all existing entries"]
    G --> H["Insert into new table"]
    H --> I["Replace old table"]
    I --> J["Insert new key-value pair"]

    style A fill:#e1f5fe
    style D fill:#c8e6c9
    style E fill:#fff3e0
    style J fill:#c8e6c9
```

### Load Factor Visualization

```mermaid
graph LR
    subgraph "Before Resize (Load Factor = 0.8)"
        A["Capacity: 8"]
        B["Size: 6"]
        C["Load Factor: 6/8 = 0.75"]
        D["Status: Approaching limit"]
    end

    subgraph "After Resize (Load Factor = 0.4)"
        E["Capacity: 16"]
        F["Size: 6"]
        G["Load Factor: 6/16 = 0.375"]
        H["Status: Optimal range"]
    end

    subgraph "Performance Impact"
        I["High load: More collisions"]
        J["Low load: Memory waste"]
        K["Target: 0.5-0.75 range"]
    end

    style C fill:#ffcdd2
    style G fill:#c8e6c9
    style K fill:#c8e6c9
```

### Hash Map Operations Complexity

```mermaid
graph TD
    A["Hash Map Operations"] --> B["Average Case O(1)"]
    A --> C["Worst Case O(n)"]

    B --> B1["Get: Direct bucket access"]
    B --> B2["Set: Hash + bucket insert"]
    B --> B3["Delete: Hash + list removal"]
    B --> B4["Good hash distribution"]

    C --> C1["All keys hash to same bucket"]
    C --> C2["Linear search through chain"]
    C --> C3["Poor hash function"]
    C --> C4["Adversarial input"]

    D[Resize Operation] --> E["Time: O(n)"]
    D --> F["Frequency: Rare"]
    D --> G["Amortized: Still O(1)"]

    style B fill:#c8e6c9
    style C fill:#ffcdd2
    style G fill:#c8e6c9
```

### Memory Layout

```mermaid
graph TD
    subgraph "Hash Map Memory Structure"
        A["HashMap struct"] --> B["buckets: []*Node"]
        A --> C["size: int"]
        A --> D["capacity: int"]
    end

    subgraph "Bucket Array"
        E["buckets[0] → Node → Node → nil"]
        F["buckets[1] → nil"]
        G["buckets[2] → Node → nil"]
        H["buckets[3] → Node → Node → Node → nil"]
    end

    subgraph "Node Structure"
        I["key: string"]
        J["value: interface{}"]
        K["next: *Node"]
    end

    B --> E

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Comparison with Other Data Structures

```mermaid
graph TD
    A["Key-Value Storage Comparison"] --> B[Hash Map]
    A --> C[Binary Search Tree]
    A --> D["Array/Slice"]

    B --> B1["Average: O(1) operations"]
    B --> B2["Memory: Sparse arrays"]
    B --> B3["Ordering: No guaranteed order"]
    B --> B4["Use case: Fast lookups"]

    C --> C1["Consistent: O(log n) operations"]
    C --> C2["Memory: Compact nodes"]
    C --> C3["Ordering: Sorted keys"]
    C --> C4["Use case: Ordered iteration"]

    D --> D1["Index access: O(1)"]
    D --> D2["Search: O(n)"]
    D --> D3["Memory: Contiguous"]
    D --> D4["Use case: Sequential data"]

    style B1 fill:#c8e6c9
    style C1 fill:#fff3e0
    style D1 fill:#c8e6c9
```

## Key Features

- **Separate Chaining**: Uses linked lists to handle hash collisions
- **Dynamic Resizing**: Automatically doubles capacity when load factor exceeds 0.75
- **Generic Values**: Supports any value type with string keys
- **FNV Hash Function**: Uses FNV-1a hash algorithm for good distribution
- **Complete API**: Standard hash map operations plus utility methods

## Implementation Details

- **Hash Function**: FNV-1a (Fowler-Noll-Vo) hash algorithm
- **Collision Resolution**: Separate chaining with linked lists
- **Load Factor**: Maintains load factor below 0.75 via automatic resizing
- **Resize Strategy**: Doubles capacity and rehashes all elements
- **Default Capacity**: 16 buckets initially

## Complexity

- **Time Complexity**:
  - Average case: O(1) for Get, Set, Delete, Has
  - Worst case: O(n) when all keys hash to the same bucket
  - Resize operation: O(n) to rehash all elements
- **Space Complexity**: O(n) where n is the number of key-value pairs

## Core Operations

### Basic Operations

- `Set(key, value)` - Insert or update a key-value pair
- `Get(key)` - Retrieve value by key, returns (value, exists)
- `Delete(key)` - Remove key-value pair, returns success boolean
- `Has(key)` - Check if key exists
- `Size()` - Get number of stored pairs
- `IsEmpty()` - Check if hash map is empty
- `Clear()` - Remove all key-value pairs

### Utility Operations

- `Keys()` - Get slice of all keys
- `Values()` - Get slice of all values
- `Entries()` - Get slice of all key-value pairs
- `ForEach(func)` - Iterate over all pairs with callback
- `LoadFactor()` - Get current load factor
- `Capacity()` - Get current bucket capacity

### Advanced Operations

- `GetBucketDistribution()` - Get distribution of items across buckets (for analysis)

## Usage

```bash
make run n=hash-map
```

## Testing

```bash
make test n=hash-map
```

## Benchmarking

```bash
make bench n=hash-map
```

## Example Usage in Go

```go
// Create a new hash map
hm := NewHashMap()

// Insert key-value pairs
hm.Set("name", "Alice")
hm.Set("age", 30)
hm.Set("city", "New York")

// Retrieve values
name, exists := hm.Get("name")
if exists {
    fmt.Printf("Name: %s\n", name)
}

// Check if key exists
if hm.Has("age") {
    fmt.Println("Age is stored")
}

// Delete a key
deleted := hm.Delete("city")
fmt.Printf("City deleted: %t\n", deleted)

// Get all keys
keys := hm.Keys()
fmt.Printf("Keys: %v\n", keys)

// Iterate over all pairs
hm.ForEach(func(key string, value any) {
    fmt.Printf("%s: %v\n", key, value)
})

// Check statistics
fmt.Printf("Size: %d\n", hm.Size())
fmt.Printf("Load Factor: %.2f\n", hm.LoadFactor())
fmt.Printf("Capacity: %d\n", hm.Capacity())
```

## Performance Characteristics

### Average Case Performance

- **Insertion**: O(1) - Direct bucket access via hash
- **Lookup**: O(1) - Direct bucket access via hash
- **Deletion**: O(1) - Direct bucket access via hash
- **Iteration**: O(n) - Must visit all stored elements

### Worst Case Performance

- **All Operations**: O(n) - When all keys hash to same bucket
- **Resize**: O(n) - Must rehash all existing elements

### Memory Usage

- **Base Memory**: O(capacity) for bucket array
- **Element Memory**: O(n) for stored key-value pairs
- **Overhead**: Minimal - only bucket pointers and size tracking

## Load Factor and Resizing

- **Target Load Factor**: 0.75 (75% of buckets occupied on average)
- **Resize Trigger**: When load factor exceeds 0.75
- **Resize Strategy**: Double the capacity and rehash all elements
- **Benefits**: Maintains good performance by reducing collision chains

## Hash Function

Uses FNV-1a (Fowler-Noll-Vo variant 1a) hash algorithm:

- **Fast**: Simple operations (XOR and multiply)
- **Good Distribution**: Minimizes clustering and collisions
- **Deterministic**: Same key always produces same hash
- **Avalanche Effect**: Small input changes cause large hash changes

## Collision Resolution

Uses separate chaining with linked lists:

- **Pros**: Simple implementation, handles any number of collisions
- **Cons**: Extra memory overhead for pointers, potential cache misses
- **Performance**: Good average case, degrades gracefully under load

## Use Cases

- **Caches**: Fast key-value storage for computed results
- **Indexing**: Map unique identifiers to data records
- **Counting**: Count occurrences of items (frequency tables)
- **Sets**: Implement set operations using keys only
- **Lookup Tables**: Fast translation between different representations
- **Memoization**: Store function results for dynamic programming

## Advantages

- **Fast Operations**: O(1) average case for basic operations
- **Dynamic Sizing**: Grows automatically as needed
- **Simple Interface**: Easy to use standard hash map API
- **Collision Handling**: Robust separate chaining approach
- **Memory Efficient**: Only stores what's needed plus small overhead

## Limitations

- **String Keys Only**: Current implementation limited to string keys
- **Worst Case**: Can degrade to O(n) with poor hash distribution
- **Memory Overhead**: Requires extra space for buckets and pointers
- **No Ordering**: Keys are not stored in any particular order
- **Resize Cost**: Occasional O(n) resize operations

## Comparison with Alternatives

- **vs. Built-in Go map**: Similar performance, educational implementation
- **vs. Open Addressing**: Separate chaining uses more memory but handles collisions better
- **vs. Binary Search Tree**: Hash map is faster for basic operations (O(1) vs O(log n))
- **vs. Linear Search**: Hash map is much faster for large datasets (O(1) vs O(n))
