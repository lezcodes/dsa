# Merge Sort

## Description

Merge Sort is a stable, divide-and-conquer sorting algorithm that divides the input array into two halves, recursively sorts both halves, and then merges the sorted halves back together. It guarantees O(n log n) time complexity in all cases, making it one of the most reliable sorting algorithms.

This implementation includes multiple variants:

- **Top-Down Merge Sort**: Classic recursive implementation
- **Bottom-Up Merge Sort**: Iterative implementation that avoids recursion
- **Stable Merge Sort**: Explicitly maintains stability for equal elements
- **Optimized Merge Sort**: Uses insertion sort for small subarrays and skips unnecessary merges
- **In-Place Merge Sort**: Sorts the array without creating a new array

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

```bash
go test -bench=. -benchmem
```

## Implementation Details

### Functions Available:

- `MergeSort(arr []int) []int` - Standard top-down merge sort
- `MergeSortBottomUp(arr []int) []int` - Iterative bottom-up approach
- `MergeSortStable(arr []int) []int` - Explicitly stable merge sort
- `MergeSortOptimized(arr []int) []int` - Optimized with insertion sort
- `MergeSortInPlace(arr []int)` - In-place sorting
- `IsSorted(arr []int) bool` - Utility to check if array is sorted

### When to Use Merge Sort:

- **Good for**: Large datasets, when stability is required, guaranteed O(n log n) performance
- **Excellent for**: External sorting, parallel processing, linked lists
- **Consider alternatives when**: Memory is very limited, working with small arrays

### Comparison with Other Sorting Algorithms:

- **vs Quick Sort**: More predictable performance, stable, but uses more memory
- **vs Heap Sort**: Stable and better cache performance, but uses more memory
- **vs Insertion Sort**: Much faster for large arrays, but slower for very small arrays
- **vs Radix Sort**: More general (works with any comparable type), but slower for integers

## Advantages

- **Guaranteed O(n log n)**: No worst-case degradation
- **Stable**: Preserves relative order of equal elements
- **Predictable**: Performance doesn't depend on input distribution
- **Parallelizable**: Easy to implement parallel versions
- **External**: Can sort data larger than available memory

## Disadvantages

- **Space Complexity**: Requires O(n) additional memory
- **Not In-Place**: Cannot sort with O(1) extra space
- **Overhead**: Slower than quicksort for small arrays due to merge overhead
