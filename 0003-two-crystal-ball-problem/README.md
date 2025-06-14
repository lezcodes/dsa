# Two Crystal Ball Problem

## Description

The Two Crystal Ball Problem is a classic algorithmic puzzle that asks: Given two crystal balls that will break if dropped from a high enough floor, determine the exact floor where they will break using the minimum number of drops in the worst case.

**Problem Setup:**

- You have a building with `n` floors
- You have exactly 2 crystal balls
- There exists a critical floor `k` where:
  - Dropping from floor `k` or higher will break the ball
  - Dropping from floor `k-1` or lower will not break the ball
- Goal: Find floor `k` with minimum worst-case number of drops

**Optimal Strategy:**
The naive approach of trying each floor sequentially would take O(n) drops in the worst case. The optimal strategy uses a square root approach:

1. **First Ball**: Drop at intervals of √n (floors √n, 2√n, 3√n, ...)
2. **Second Ball**: When the first ball breaks, linearly search from the last safe floor

This gives a worst-case complexity of O(√n) drops.

**Example with 100 floors:**

- Drop first ball at floors: 10, 20, 30, 40, 50, 60, 70, 80, 90, 100
- If it breaks at floor 50, use second ball to test floors: 41, 42, 43, 44, 45, 46, 47, 48, 49
- Maximum drops: 10 (first ball) + 9 (second ball) = 19 drops

## Algorithm Implementation

The algorithm takes a boolean array where `true` indicates the ball will break at that floor:

```go
func TwoCrystalBalls(breaks []bool) int {
    jumpAmount := √n

    // Phase 1: Find the interval where breaking occurs
    for i := jumpAmount; i < n; i += jumpAmount {
        if breaks[i] {
            break
        }
    }

    // Phase 2: Linear search within the interval
    for j := i - jumpAmount; j <= i; j++ {
        if breaks[j] {
            return j
        }
    }

    return -1 // No breaking floor found
}
```

## Complexity

- **Time Complexity**: O(√n)
  - First phase: O(√n) drops to find the interval
  - Second phase: O(√n) drops to find exact floor within interval
  - Total: O(√n) + O(√n) = O(√n)
- **Space Complexity**: O(1) - only uses a constant amount of extra space

## Comparison with Other Approaches

| Approach          | Worst Case Drops | Best Case Drops  | Average Case     |
| ----------------- | ---------------- | ---------------- | ---------------- |
| Linear Search     | O(n)             | O(1)             | O(n/2)           |
| Binary Search     | Not applicable\* | Not applicable\* | Not applicable\* |
| Two Crystal Balls | O(√n)            | O(√n)            | O(√n)            |

\*Binary search cannot be used because once a ball breaks, you cannot continue using it.

## Usage

```bash
make run NAME=two-crystal-ball-problem
```

**Example Output:**

```
Running: 0003-two-crystal-ball-problem
----------------------------------------
Running 0003-two-crystal-ball-problem...
Result: map[breaking_floor:7 breaks_array:[false false false false false false false true true true] description:First floor where crystal ball breaks]
```

## Testing

```bash
make test NAME=two-crystal-ball-problem
```

**Test Coverage:**

- Edge cases: empty arrays, single elements, no breaks
- Boundary conditions: breaks at first/last floor
- Various array sizes: 10, 100, 1000, 10000 floors
- Performance benchmarks for different input sizes

## Real-World Applications

This problem models scenarios where:

- **Testing has a cost**: Each test/experiment consumes resources
- **Limited resources**: You have a fixed number of attempts
- **Monotonic property**: Once the threshold is reached, all higher values also trigger the condition

Examples:

- **Load testing**: Finding the breaking point of a system
- **Drug dosage**: Finding the minimum effective dose
- **Quality control**: Testing product failure thresholds
- **Network testing**: Finding bandwidth limits
