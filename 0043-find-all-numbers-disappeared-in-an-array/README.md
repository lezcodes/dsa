# [Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)

## Description

Implementation of the "Find All Numbers Disappeared in an Array" problem (LeetCode 448) using two approaches. Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

## Visual Representation

```mermaid
graph TD
    A[Input: Array with duplicates and missing numbers] --> B{Choose Algorithm}
    B -->|Extra Space| C[Hash Set Approach]
    B -->|In-Place| D[Negative Marking Approach]

    C --> C1[Create set of seen numbers]
    C1 --> C2[Iterate through range 1 to n]
    C2 --> C3{Number in set?}
    C3 -->|No| C4[Add to result]
    C3 -->|Yes| C5[Continue]
    C4 --> C6{More numbers?}
    C5 --> C6
    C6 -->|Yes| C2
    C6 -->|No| C7[Return missing numbers]

    D --> D1[For each number in array]
    D1 --> D2["Use abs(number) as index"]
    D2 --> D3["Mark nums[index-1] as negative"]
    D3 --> D4{More numbers?}
    D4 -->|Yes| D1
    D4 -->|No| D5[Scan array for positive values]
    D5 --> D6["Positive at index i means i+1 is missing"]

    style C7 fill:#c8e6c9
    style D6 fill:#c8e6c9
```

```mermaid
graph LR
    subgraph "Example: [4,3,2,7,8,2,3,1] → Missing: [5,6]"
        A1[4] --> A2[3] --> A3[2] --> A4[7] --> A5[8] --> A6[2] --> A7[3] --> A8[1]
    end

    subgraph "In-Place Negative Marking Process"
        B1["Step 1: 4 → mark index 3 negative"]
        B2["Step 2: 3 → mark index 2 negative"]
        B3["Step 3: 2 → mark index 1 negative"]
        B4["Step 4: 7 → mark index 6 negative"]
        B5["Step 5: 8 → mark index 7 negative"]
        B6["...continue marking..."]
        B7["Result: indices 4,5 remain positive → missing 5,6"]
    end

    style B7 fill:#c8e6c9
```

## In-Place Algorithm Step-by-Step

```mermaid
graph TD
    A["Array: [4,3,2,7,8,2,3,1]"] --> B["Process 4: mark index 3"]
    B --> C["Array: [4,3,2,-7,8,2,3,1]"]
    C --> D["Process 3: mark index 2"]
    D --> E["Array: [4,3,-2,-7,8,2,3,1]"]
    E --> F["Process 2: mark index 1"]
    F --> G["Array: [4,-3,-2,-7,8,2,3,1]"]
    G --> H["Continue for all elements..."]
    H --> I["Final: [-4,-3,-2,-7,8,2,-3,-1]"]
    I --> J["Scan: indices 4,5 are positive"]
    J --> K["Result: [5,6] are missing"]

    style I fill:#fff3e0
    style K fill:#c8e6c9
```

## Algorithms

### 1. Hash Set Approach (Clear and Simple)

```go
func FindDisappearedNumbersHashSet(nums []int) []int {
    seen := make(map[int]bool)
    for _, num := range nums {
        seen[num] = true
    }

    result := []int{}
    for i := 1; i <= len(nums); i++ {
        if !seen[i] {
            result = append(result, i)
        }
    }

    return result
}
```

### 2. In-Place Negative Marking (Space Optimized)

```go
func FindDisappearedNumbersInPlace(nums []int) []int {
    // Mark seen numbers by negating values at corresponding indices
    for _, num := range nums {
        index := abs(num) - 1  // Convert to 0-based index
        if index >= 0 && index < len(nums) && nums[index] > 0 {
            nums[index] = -nums[index]  // Mark as seen
        }
    }

    // Collect indices with positive values (missing numbers)
    result := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            result = append(result, i+1)  // Convert back to 1-based
        }
    }

    return result
}
```

## How In-Place Algorithm Works

1. **Marking Phase**: For each number `num` in the array:

   - Use `abs(num) - 1` as an index (convert 1-based to 0-based)
   - Mark `nums[index]` as negative to indicate the number is present

2. **Collection Phase**: Scan the array:
   - If `nums[i]` is positive, then `i+1` is missing from the original array
   - Add `i+1` to the result

**Key Insight**: We use the array itself as a hash table where the sign indicates presence/absence.

## Complexity

### Hash Set Approach

- **Time Complexity**: O(n) - two passes through the data
- **Space Complexity**: O(n) - hash set storage

### In-Place Negative Marking

- **Time Complexity**: O(n) - two passes through array
- **Space Complexity**: O(1) - only modifies input array (result doesn't count)
- **Constraint**: Input array is modified during processing

## Usage

```bash
make run n=0043-find-all-numbers-disappeared-in-an-array
```

```bash
make check n=0043-find-all-numbers-disappeared-in-an-array
```

## Testing

```bash
make test n=0043-find-all-numbers-disappeared-in-an-array
```

The test suite includes:

- Problem examples: [4,3,2,7,8,2,3,1]→[5,6], [1,1]→[2]
- Edge cases: no missing numbers, all missing except one, single element
- Both algorithm implementations tested with identical test cases
- Performance benchmarks comparing both approaches

## When to Use

**Hash Set Approach:**

- ✅ Input array remains unchanged
- ✅ Clear and intuitive logic
- ✅ Easy to understand and maintain
- ❌ Requires O(n) extra space

**In-Place Negative Marking:**

- ✅ Meets O(1) space requirement (follow-up)
- ✅ Optimal space complexity
- ✅ Clever use of array as hash table
- ❌ Modifies input array
- ❌ More complex logic

**Choose Hash Set when:** You need to preserve the input array or want simpler code
**Choose In-Place when:** Space is critical and you can modify the input (follow-up requirement)
