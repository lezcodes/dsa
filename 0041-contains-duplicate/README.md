# [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

## Description

Implementation of the "Contains Duplicate" problem (LeetCode 217) using two different approaches. Given an integer array, determine if any value appears at least twice in the array, returning true if duplicates exist and false if all elements are distinct.

## Visual Representation

```mermaid
graph TD
    A[Input: Array of integers] --> B{Choose Algorithm}
    B -->|Hash Set| C[Hash Set Approach]
    B -->|In-Place| D[In-Place Negative Marking]

    C --> C1[Create empty set]
    C1 --> C2[For each element]
    C2 --> C3{Element in set?}
    C3 -->|Yes| C4[Return true - Duplicate found]
    C3 -->|No| C5[Add element to set]
    C5 --> C6{More elements?}
    C6 -->|Yes| C2
    C6 -->|No| C7[Return false - No duplicates]

    D --> D1[For each element]
    D1 --> D2["Use abs(element) as index"]
    D2 --> D3{Element at index negative?}
    D3 -->|Yes| D4[Return true - Duplicate found]
    D3 -->|No| D5[Mark element at index negative]
    D5 --> D6{More elements?}
    D6 -->|Yes| D1
    D6 -->|No| D7[Return false - No duplicates]

    style C4 fill:#ffcdd2
    style C7 fill:#c8e6c9
    style D4 fill:#ffcdd2
    style D7 fill:#c8e6c9
```

```mermaid
graph LR
    subgraph "Example: [1, 2, 3, 1]"
        A1[1] --> A2[2] --> A3[3] --> A4[1]
    end

    subgraph "Hash Set Approach"
        B1["Set: {1}"] --> B2["Set: {1,2}"] --> B3["Set: {1,2,3}"] --> B4[1 already exists ✗]
    end

    subgraph "Result"
        C1[true - Duplicate found]
    end

    A1 -.->|Add| B1
    A2 -.->|Add| B2
    A3 -.->|Add| B3
    A4 -.->|Duplicate!| B4
    B4 --> C1

    style A4 fill:#ffcdd2
    style B4 fill:#ffcdd2
    style C1 fill:#ffcdd2
```

## Algorithms

### 1. Hash Set Approach (Set Length Comparison)

```go
func ContainsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        seen[num] = true
    }
    return len(seen) != len(nums)
}
```

### 2. In-Place Negative Marking (Space Optimized)

```go
func ContainsDuplicateInPlace(nums []int) bool {
    for _, num := range nums {
        index := abs(num)
        if index > 0 && index <= len(nums) {
            if nums[index-1] < 0 {
                return true  // Already marked negative
            }
            nums[index-1] = -nums[index-1]  // Mark as seen
        }
    }
    return false
}
```

## Complexity

### Hash Set Approach

- **Time Complexity**: O(n) - single pass through array
- **Space Complexity**: O(n) - hash set storage

### In-Place Negative Marking

- **Time Complexity**: O(n) - single pass through array
- **Space Complexity**: O(1) - constant extra space
- **Constraint**: Only works when array elements are positive integers ≤ array length

## Usage

```bash
make run n=0041-contains-duplicate
```

```bash
make check n=0041-contains-duplicate
```

## Testing

```bash
make test n=0041-contains-duplicate
```

The test suite includes:

- Problem examples: [1,2,3,1], [1,2,3,4], [1,1,1,3,3,4,3,2,4,2]
- Edge cases: empty arrays, single elements, immediate duplicates
- Both algorithm implementations tested with identical test cases
- Performance benchmarks for both approaches

## When to Use

**Hash Set Approach:**

- ✅ Works with any integer values (negative, zero, large numbers)
- ✅ Clear and intuitive logic
- ✅ No constraints on input values
- ❌ Requires O(n) extra space

**In-Place Negative Marking:**

- ✅ O(1) space complexity
- ✅ No additional data structures needed
- ❌ Limited to positive integers ≤ array length
- ❌ Mutates the input array

**Choose Hash Set when:** Input has negative numbers, large numbers, or space isn't critical
**Choose In-Place when:** Memory is constrained and input meets the positive integer constraint
