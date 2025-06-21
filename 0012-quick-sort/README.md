# Quick Sort

## Description

Quick Sort is a highly efficient divide-and-conquer sorting algorithm that works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

## Visual Representation

### Algorithm Flow

```mermaid
graph TD
    A[Choose Pivot] --> B[Partition Array]
    B --> C[Elements < Pivot]
    B --> D[Pivot]
    B --> E[Elements > Pivot]
    C --> F[Recursively Sort Left]
    E --> G[Recursively Sort Right]
    F --> H[Combine Results]
    G --> H
    D --> H
    H --> I[Sorted Array]

    style A fill:#e1f5fe
    style D fill:#fff3e0
    style I fill:#c8e6c9
```

### Partitioning Process

```mermaid
graph LR
    subgraph "Initial Array"
        A1[3] --- A2[6] --- A3[8] --- A4[10] --- A5[1] --- A6[2] --- A7[1]
    end

    subgraph "Choose Pivot (last element = 1)"
        B1[3] --- B2[6] --- B3[8] --- B4[10] --- B5[1] --- B6[2] --- B7["1 (pivot)"]
    end

    subgraph "After Partitioning"
        C1["1 (pivot)"] --- C2[3] --- C3[6] --- C4[8] --- C5[10] --- C6[2]
        C7["< pivot"]
        C8["> pivot"]
    end

    style B7 fill:#fff3e0
    style C1 fill:#fff3e0
    style C7 fill:#e8f5e8
    style C8 fill:#ffebee
```

### Divide and Conquer Visualization

```mermaid
graph TD
    A["[3,6,8,10,1,2,1]"] --> B["Pivot: 1"]
    B --> C["Left: []"]
    B --> D["Right: [3,6,8,10,2,1]"]

    D --> E["Pivot: 1"]
    E --> F["Left: []"]
    E --> G["Right: [3,6,8,10,2]"]

    G --> H["Pivot: 2"]
    H --> I["Left: []"]
    H --> J["Right: [3,6,8,10]"]

    J --> K["Pivot: 10"]
    K --> L["Left: [3,6,8]"]
    K --> M["Right: []"]

    L --> N["Continue recursion..."]

    style A fill:#e1f5fe
    style B fill:#fff3e0
    style E fill:#fff3e0
    style H fill:#fff3e0
    style K fill:#fff3e0
```

### Time Complexity Scenarios

```mermaid
graph LR
    subgraph "Best Case: O(n log n)"
        A1[Balanced partitions] --> A2[log n levels] --> A3[n work per level]
    end

    subgraph "Worst Case: O(n²)"
        B1[Unbalanced partitions] --> B2[n levels] --> B3[n work per level]
    end

    style A1 fill:#c8e6c9
    style B1 fill:#ffcdd2
```

This implementation includes multiple variants:

- **Basic Quick Sort**: Uses the last element as pivot
- **Random Pivot**: Randomly selects a pivot to avoid worst-case performance
- **Median-of-Three**: Uses the median of first, middle, and last elements as pivot
- **In-Place Sorting**: Sorts the array without creating a new array

## Algorithm Steps

1. **Choose a pivot** element from the array
2. **Partition** the array so that:
   - Elements smaller than the pivot come before it
   - Elements greater than the pivot come after it
3. **Recursively apply** the same process to the sub-arrays

## Complexity

- **Time Complexity**:
  - Best/Average Case: O(n log n)
  - Worst Case: O(n²) - occurs when the pivot is always the smallest or largest element
- **Space Complexity**:
  - Average Case: O(log n) - due to recursion stack
  - Worst Case: O(n) - in case of unbalanced partitions

## Key Features

- **Not Stable**: Equal elements may not maintain their relative order
- **In-Place**: Can sort with O(1) extra space (excluding recursion stack)
- **Adaptive**: Performance can be improved with better pivot selection strategies
- **Cache-Efficient**: Good locality of reference

## Pivot Selection Strategies

### 1. Last Element (Basic)

- Simple implementation
- Worst case: O(n²) for already sorted arrays

### 2. Random Pivot

- Reduces probability of worst-case performance
- Expected time complexity: O(n log n)

### 3. Median-of-Three

- Takes median of first, middle, and last elements
- Better performance on partially sorted arrays
- Reduces worst-case scenarios

## Usage

```bash
make run NAME=0012-quick-sort
```

## Testing

```bash
make test NAME=0012-quick-sort
```

## Benchmarking

```bash
go test -bench=. -benchmem
```

## Implementation Details

### Functions Available:

- `QuickSort(arr []int) []int` - Basic quick sort returning new array
- `QuickSortRandomPivot(arr []int) []int` - Quick sort with random pivot
- `QuickSortMedianOfThree(arr []int) []int` - Quick sort with median-of-three pivot
- `QuickSortInPlace(arr []int)` - In-place quick sort
- `IsSorted(arr []int) bool` - Utility to check if array is sorted

### When to Use Quick Sort:

- **Good for**: Large datasets, when average-case performance is important
- **Avoid when**: Stability is required, or when worst-case performance must be guaranteed

### Comparison with Other Sorting Algorithms:

- **vs Merge Sort**: Faster in practice, but not stable and has worse worst-case
- **vs Heap Sort**: Better cache performance, but not stable
- **vs Insertion Sort**: Much faster for large arrays, but slower for very small arrays
