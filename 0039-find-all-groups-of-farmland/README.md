# find-all-groups-of-farmland

## Description

Implementation of the "Find All Groups of Farmland" problem using a greedy rectangular expansion algorithm.

Given a binary matrix where 1 represents farmland and 0 represents forest, finds all rectangular groups of connected farmland. Each group is guaranteed to be rectangular and non-adjacent to other groups.

**Algorithm approach**:

1. **Scan the matrix** from top-left to bottom-right
2. **When farmland found** (value 1), start a new rectangular group
3. **Expand the rectangle** by moving right and down to find the complete boundaries
4. **Mark cells as visited** by setting them to 0 during expansion
5. **Record coordinates** as [top-left-row, top-left-col, bottom-right-row, bottom-right-col]

**Key insight**: Since groups are guaranteed to be rectangular and non-adjacent, we can use a greedy approach to expand each rectangle to its maximum size without worrying about complex connectivity patterns.

## Complexity

- Time Complexity: O(m × n) where m and n are the matrix dimensions - each cell is visited exactly once
- Space Complexity: O(m × n) for the matrix copy, plus O(k) for storing k groups in the result

## Usage

```bash
make run n=0039-find-all-groups-of-farmland
```

## Testing

```bash
make test n=0039-find-all-groups-of-farmland
```
