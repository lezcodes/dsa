# bloom-filter

## Description

Implementation of a Bloom Filter - a space-efficient probabilistic data structure used to test whether an element is a member of a set.

Key characteristics:

- **False positives possible**: May return "possibly in set" for items not actually added
- **No false negatives**: If it says "not in set", the item is definitely not in the set
- **Space efficient**: Uses a bit array much smaller than storing all elements
- **Fast operations**: O(k) time for both insertion and lookup, where k is the number of hash functions

## Bloom Filter Structure

```mermaid
graph TD
    subgraph "Bloom Filter Components"
        A["Bit Array of size m"] --> B["Initially all bits = 0"]
        C["k Hash Functions"] --> D["h₁, h₂, h₃, ..., hₖ"]
        E["Elements to Add"] --> F["Apply all k hash functions"]
        F --> G["Set corresponding bits to 1"]
    end

    subgraph "Bit Array Example (m=16)"
        H["Index: 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15"]
        I["Bits:  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0  0"]
    end

    style A fill:#e1f5fe
    style C fill:#e1f5fe
    style H fill:#fff3e0
    style I fill:#c8e6c9
```

## Adding Elements Process

```mermaid
graph TD
    A["Add Element 'apple'"] --> B["Apply Hash Functions"]
    B --> C["h₁('apple') = 3"]
    B --> D["h₂('apple') = 7"]
    B --> E["h₃('apple') = 12"]

    C --> F["Set bit[3] = 1"]
    D --> G["Set bit[7] = 1"]
    E --> H["Set bit[12] = 1"]

    I["Bit Array After Adding 'apple'"] --> J["0 0 0 1 0 0 0 1 0 0 0 0 1 0 0 0"]

    K["Add Element 'banana'"] --> L["h₁('banana') = 1"]
    K --> M["h₂('banana') = 7"]
    K --> N["h₃('banana') = 14"]

    O["Final Bit Array"] --> P["0 1 0 1 0 0 0 1 0 0 0 0 1 0 1 0"]

    style A fill:#e1f5fe
    style K fill:#e1f5fe
    style F fill:#c8e6c9
    style G fill:#c8e6c9
    style H fill:#c8e6c9
    style J fill:#fff3e0
    style P fill:#c8e6c9
```

## Lookup Process

```mermaid
graph TD
    A["Query Element 'apple'"] --> B["Apply Same Hash Functions"]
    B --> C["h₁('apple') = 3"]
    B --> D["h₂('apple') = 7"]
    B --> E["h₃('apple') = 12"]

    C --> F["Check bit[3]"]
    D --> G["Check bit[7]"]
    E --> H["Check bit[12]"]

    F -->|1| I["✓ Bit is set"]
    G -->|1| J["✓ Bit is set"]
    H -->|1| K["✓ Bit is set"]

    I --> L["All bits set?"]
    J --> L
    K --> L
    L -->|Yes| M["Possibly in set"]

    N["Query Element 'cherry'"] --> O["h₁('cherry') = 2"]
    N --> P["h₂('cherry') = 5"]
    N --> Q["h₃('cherry') = 9"]

    O --> R["Check bit[2]"]
    R -->|0| S["Definitely NOT in set"]

    style A fill:#e1f5fe
    style N fill:#e1f5fe
    style M fill:#fff3e0
    style S fill:#c8e6c9
```

## False Positive Scenario

```mermaid
graph TD
    subgraph "Elements Added"
        A[Added: 'apple', 'banana']
        B[Bit positions set: 1,3,7,12,14]
    end

    subgraph "False Positive Example"
        C["Query: 'orange'"] --> D["h₁('orange') = 1"]
        C --> E["h₂('orange') = 12"]
        C --> F["h₃('orange') = 14"]

        D --> G["bit[1] = 1?"]
        E --> H["bit[12] = 1?"]
        F --> I["bit[14] = 1?"]

        G -->|Yes| J["✓ Set by 'banana'"]
        H -->|Yes| K["✓ Set by 'apple'"]
        I -->|Yes| L["✓ Set by 'banana'"]
    end

    J --> M["All bits are 1"]
    K --> M
    L --> M
    M --> N["FALSE POSITIVE<br/>'orange' not actually added"]

    style A fill:#e1f5fe
    style C fill:#fff3e0
    style N fill:#ffcdd2
```

## Hash Functions Implementation

```mermaid
graph TD
    subgraph "Multiple Hash Functions Strategy"
        A["Primary Hash Functions"] --> B["FNV-1a Hash"]
        A --> C["MD5-based Hash"]
        A --> D["Polynomial Rolling Hash"]

        E["Additional Hash Functions"] --> F["Linear Combination"]
        F --> G["hᵢ(x) = (hash1(x) + i × hash2(x)) mod m"]

        B --> H["Fast, good distribution"]
        C --> I["Cryptographic quality"]
        D --> J["Good for strings"]
        G --> K["Generates k functions from 2 base functions"]
    end

    style A fill:#e1f5fe
    style E fill:#e1f5fe
    style H fill:#c8e6c9
    style I fill:#c8e6c9
    style J fill:#c8e6c9
    style K fill:#c8e6c9
```

## Algorithm Flow

```mermaid
flowchart TD
    A["Create Bloom Filter<br/>size m, k hash functions"] --> B["Initialize bit array<br/>all bits = 0"]

    B --> C["Operation?"]
    C -->|Add| D["Hash element with all k functions"]
    C -->|Contains| H["Hash element with all k functions"]

    D --> E["Set bits at hash positions to 1"]
    E --> F["Element added"]
    F --> C

    H --> I["All bits at hash positions = 1?"]
    I -->|Yes| J["Possibly in set<br/>might be false positive"]
    I -->|No| K["Definitely NOT in set<br/>guaranteed accurate"]

    J --> C
    K --> C

    style A fill:#e1f5fe
    style F fill:#c8e6c9
    style J fill:#fff3e0
    style K fill:#c8e6c9
```

## Optimal Parameters

```mermaid
graph TD
    subgraph "Parameter Selection"
        A["Given: n elements, desired false positive rate p"] --> B["Optimal bit array size"]
        B --> C["m = -(n × ln(p)) / (ln(2))²"]

        A --> D["Optimal number of hash functions"]
        D --> E["k = (m/n) × ln(2)"]

        C --> F["Trade-off: Larger m = less false positives"]
        E --> G["Trade-off: More k = better accuracy but slower"]
    end

    subgraph "Example Calculation"
        H["n = 1000 elements"] --> I["p = 0.01 (1% false positive)"]
        I --> J["m ≈ 9585 bits ≈ 1.2 KB"]
        J --> K["k ≈ 7 hash functions"]
    end

    style A fill:#e1f5fe
    style H fill:#fff3e0
    style J fill:#c8e6c9
    style K fill:#c8e6c9
```

## Use Cases and Applications

```mermaid
graph TD
    A["Bloom Filter Applications"] --> B["Database Systems"]
    A --> C["Web Caching"]
    A --> D["Network Systems"]
    A --> E["Distributed Systems"]

    B --> B1["Avoid expensive disk lookups"]
    B --> B2["LSM-tree optimizations"]

    C --> C1["CDN cache filtering"]
    C --> C2["Web crawler URL deduplication"]

    D --> D1["Network packet filtering"]
    D --> D2["Malicious URL detection"]

    E --> E1["Distributed cache coordination"]
    E --> E2["Set membership in P2P networks"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
    style E1 fill:#c8e6c9
```

The implementation uses multiple hash functions (FNV-1a, MD5-based, and polynomial rolling hash) to distribute elements across the bit array. Additional hash functions are generated using linear combinations of the base hashes.

## How it works

1. **Initialization**: Create a bit array of size m and choose k hash functions
2. **Adding elements**: Hash the element with all k functions, set corresponding bits to 1
3. **Testing membership**: Hash the element with all k functions, check if all bits are 1
4. **Result interpretation**:
   - All bits set → "possibly in set" (could be false positive)
   - Any bit unset → "definitely not in set" (guaranteed accurate)

## Complexity

- Time Complexity: O(k) for both Add and Contains operations, where k is the number of hash functions
- Space Complexity: O(m) where m is the size of the bit array

## Usage

```bash
make run n=0038-bloom-filter
```

## Testing

```bash
make test n=0038-bloom-filter
```
