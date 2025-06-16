# detect-cycles

## Description

Implementation of cycle detection in directed graphs using Depth-First Search (DFS) with three-color marking.

This solves the issue tracking circular dependency problem where:

- Each issue is a node in a directed graph
- If issue i is blocked by issue j, there's a directed edge from j to i
- We need to find all circular dependency loops

**Algorithm approach**:

1. **Build graph** from blockers matrix - if blockers[i][j] is true, add edge j→i
2. **Three-color DFS** - WHITE (unvisited), GRAY (in current path), BLACK (finished)
3. **Cycle detection** - when we find a back edge to a GRAY node, extract the cycle
4. **Duplicate removal** - normalize cycles and remove duplicates

**Key insight**: Uses the classical DFS cycle detection algorithm where a back edge (edge to an ancestor in the DFS tree) indicates a cycle. The three-color approach ensures we can distinguish between cross edges and back edges.

## Complexity

- Time Complexity: O(V + E) where V is the number of nodes and E is the number of edges
- Space Complexity: O(V) for the color array, parent array, and recursion stack

## Usage

```bash
make run n=0040-detect-cycles
```

## Testing

```bash
make test n=0040-detect-cycles
```
