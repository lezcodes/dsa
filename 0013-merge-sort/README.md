# Merge Sort

## Description

Merge Sort is a stable, divide-and-conquer sorting algorithm that works by recursively dividing the array into halves, sorting each half, and then merging the sorted halves back together. It guarantees O(n log n) time complexity in all cases.

This implementation includes multiple variants:

- **Top-Down Merge Sort**: Classic recursive implementation
- **Bottom-Up Merge Sort**: Iterative implementation that avoids recursion
- **Stable Merge Sort**: Explicitly maintains stability for equal elements
- **Optimized Merge Sort**: Uses insertion sort for small subarrays and skips unnecessary merges
- **In-Place Merge Sort**: Sorts the array without creating a new array

## Visual Representation

### Divide and Conquer Process

```mermaid
graph TD
    A["[38, 27, 43, 3, 9, 82, 10]"] --> B["[38, 27, 43, 3]"]
    A --> C["[9, 82, 10]"]

    B --> D["[38, 27]"]
    B --> E["[43, 3]"]

    C --> F["[9, 82]"]
    C --> G["[10]"]

    D --> H["[38]"]
    D --> I["[27]"]

    E --> J["[43]"]
    E --> K["[3]"]

    F --> L["[9]"]
    F --> M["[82]"]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style I fill:#c8e6c9
    style J fill:#c8e6c9
    style K fill:#c8e6c9
    style L fill:#c8e6c9
    style M fill:#c8e6c9
    style G fill:#c8e6c9
```

### Merge Process (Conquer Phase)

```mermaid
graph TD
    A["[27]"] --> C["[27, 38]"]
    B["[38]"] --> C

    D["[3]"] --> F["[3, 43]"]
    E["[43]"] --> F

    G["[9]"] --> I["[9, 82]"]
    H["[82]"] --> I

    C --> J["[3, 27, 38, 43]"]
    F --> J

    I --> K["[9, 10, 82]"]
    L["[10]"] --> K

    J --> M["[3, 9, 10, 27, 38, 43, 82]"]
    K --> M

    style A fill:#c8e6c9
    style B fill:#c8e6c9
    style M fill:#4caf50
```

### Merge Algorithm Step-by-Step

```mermaid
graph TD
    A[Two Sorted Arrays] --> B[Initialize pointers: i=0, j=0]
    B --> C["Compare arr1[i] vs arr2[j]"]
    C --> D{"arr1[i] <= arr2[j]?"}
    D -->|Yes| E["Add arr1[i] to result, i++"]
    D -->|No| F["Add arr2[j] to result, j++"]
    E --> G{More elements?}
    F --> G
    G -->|Yes| C
    G -->|No| H[Copy remaining elements]
    H --> I[Merge Complete]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
```

### Merge Sort Algorithm Flow

```mermaid
graph TD
    A["mergeSort(array)"] --> B{"array.length <= 1?"}
    B -->|Yes| C[Return array]
    B -->|No| D[Find middle point]
    D --> E[Split into left and right halves]
    E --> F["mergeSort(left)"]
    E --> G["mergeSort(right)"]
    F --> H["merge(left, right)"]
    G --> H
    H --> I[Return merged array]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
    style C fill:#c8e6c9
```

### Time Complexity Analysis

```mermaid
graph TD
    A[Merge Sort Complexity] --> B[Divide Phase]
    A --> C[Conquer Phase]

    B --> B1["log n levels<br/>(halving each time)"]
    C --> C1["O(n) merge at each level<br/>(compare and copy)"]

    B1 --> D["Total: O(n) × log n = O(n log n)"]
    C1 --> D

    E[Space Complexity] --> F["O(n) - requires extra arrays"]

    style A fill:#e1f5fe
    style D fill:#c8e6c9
    style F fill:#fff3e0
```

### Comparison with Other Sorts

```mermaid
graph LR
    subgraph "Merge Sort Advantages"
        A1[Stable sorting]
        A2["Guaranteed O(n log n)"]
        A3[Predictable performance]
        A4[Good for large datasets]
    end

    subgraph "Merge Sort Disadvantages"
        B1["O(n) extra space"]
        B2[Not in-place]
        B3[Overhead for small arrays]
    end

    subgraph "Best Use Cases"
        C1[Large datasets]
        C2[Stability required]
        C3[Linked lists]
        C4[External sorting]
    end

    style A1 fill:#c8e6c9
    style A2 fill:#c8e6c9
    style B1 fill:#ffcdd2
    style B2 fill:#ffcdd2
```

### External Merge Sort

```mermaid
graph LR
    subgraph "Phase 1: Create Sorted Runs"
        A[Large File] --> B[Read chunks into memory]
        B --> C[Sort each chunk]
        C --> D[Write sorted runs to disk]
    end

    subgraph "Phase 2: Merge Runs"
        E[Read runs from disk] --> F[K-way merge]
        F --> G[Write final sorted file]
    end

    D --> E

    style A fill:#e1f5fe
    style G fill:#c8e6c9
```

## Algorithm Steps

1. **Divide**: Split the array into two halves at the middle point
2. **Conquer**: Recursively sort both halves
3. **Combine**: Merge the two sorted halves back together in sorted order

## Complexity

- **Time Complexity**:
  - Best Case: O(n log n)
  - Average Case: O(n log n)
  - Worst Case: O(n log n) - guaranteed!
- **Space Complexity**:
  - Standard: O(n) - for temporary arrays during merging
  - Optimized: O(n) - reuses temporary array

## Key Features

- **Stable**: Equal elements maintain their relative order
- **Predictable**: Always O(n log n) time complexity
- **Parallelizable**: Can be easily parallelized
- **External Sorting**: Suitable for sorting large datasets that don't fit in memory
- **Not In-Place**: Requires additional memory for merging

## Variants Implemented

### 1. Top-Down (Recursive)

- Classic divide-and-conquer approach
- Uses recursion to split and merge
- Most intuitive implementation

### 2. Bottom-Up (Iterative)

- Avoids recursion overhead
- Starts with small subarrays and builds up
- Better for systems with limited stack space

### 3. Stable Merge Sort

- Explicitly preserves stability
- Important when sorting objects with multiple keys
- Same performance as standard merge sort

### 4. Optimized Merge Sort

- Uses insertion sort for small subarrays (≤10 elements)
- Skips merge step if array is already sorted
- Reuses temporary array to reduce allocations
- Better practical performance

### 5. In-Place Merge Sort

- Sorts without creating new arrays
- Uses the same merge logic but modifies original array
- Space-efficient variant

## Usage

```bash
make run NAME=0013-merge-sort
```

## Testing

```bash
make test NAME=0013-merge-sort
```

## Benchmarking

```

```
