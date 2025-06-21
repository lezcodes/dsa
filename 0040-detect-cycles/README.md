# detect-cycles

## Description

Implementation of cycle detection in directed graphs using Depth-First Search (DFS) with three-color marking.

This solves the issue tracking circular dependency problem where:

- Each issue is a node in a directed graph
- If issue i is blocked by issue j, there's a directed edge from j to i
- We need to find all circular dependency loops

## Graph Representation

```mermaid
graph TD
    subgraph "Issue Dependency Example"
        A[Issue 0] --> B[Issue 1]
        B --> C[Issue 2]
        C --> A

        D[Issue 3] --> E[Issue 4]
        E --> D

        F[Issue 5] --> G[Issue 6]
    end

    subgraph "Cycles Found"
        H["Cycle 1: 0 → 1 → 2 → 0"]
        I["Cycle 2: 3 → 4 → 3"]
        J["No cycle: 5 → 6"]
    end

    style A fill:#ffcdd2
    style B fill:#ffcdd2
    style C fill:#ffcdd2
    style D fill:#fff3e0
    style E fill:#fff3e0
    style F fill:#e1f5fe
    style G fill:#c8e6c9
    style H fill:#ffcdd2
    style I fill:#fff3e0
```

## Three-Color DFS Algorithm

```mermaid
graph TD
    subgraph "Color States"
        W["WHITE: Unvisited"] --> G["GRAY: In current path"]
        G --> B["BLACK: Completely processed"]

        W2[Color = 0] --> G2[Color = 1]
        G2 --> B2[Color = 2]
    end

    subgraph "Cycle Detection Logic"
        A["Visit node"] --> B["Mark as GRAY"]
        B --> C["Explore neighbors"]
        C --> D{"Neighbor color?"}
        D -->|WHITE| E["Recursive DFS"]
        D -->|GRAY| F["Back edge found - CYCLE!"]
        D -->|BLACK| G["Cross edge - continue"]
        E --> H["Mark as BLACK"]
        F --> I["Extract cycle path"]
        G --> J["Check next neighbor"]
        H --> J
    end

    style W fill:#ffffff
    style G fill:#e0e0e0
    style B fill:#000000
    style B2 fill:#000000
    style F fill:#ffcdd2
    style I fill:#ffcdd2
```

## Algorithm Flow

```mermaid
flowchart TD
    A["Initialize: All nodes WHITE"] --> B["Start DFS from unvisited node"]
    B --> C["Mark current node GRAY"]
    C --> D["For each neighbor"]

    D --> E{Neighbor color?}
    E -->|WHITE| F["Recursive DFS on neighbor"]
    E -->|GRAY| G["Back edge - Cycle detected!"]
    E -->|BLACK| H["Cross edge - Continue"]

    F --> I{Cycle found in recursion?}
    I -->|Yes| J["Propagate cycle info"]
    I -->|No| K["Continue with next neighbor"]

    G --> L[Extract cycle from current path]
    H --> M[Check next neighbor]
    K --> M
    J --> M
    L --> N["Add cycle to results"]

    M --> O{More neighbors?}
    O -->|Yes| D
    O -->|No| P["Mark current node BLACK"]

    P --> Q{More unvisited nodes?}
    Q -->|Yes| B
    Q -->|No| R["Return all cycles"]

    style A fill:#e1f5fe
    style G fill:#ffcdd2
    style L fill:#ffcdd2
    style R fill:#c8e6c9
```

## Cycle Extraction Process

```mermaid
graph TD
    subgraph "DFS Stack When Cycle Found"
        S1[Node 0: GRAY] --> S2[Node 1: GRAY]
        S2 --> S3[Node 2: GRAY]
        S3 --> S4[Back edge to Node 0]
    end

    subgraph "Cycle Extraction"
        A["Back edge: 2 → 0"] --> B["Find 0 in current path"]
        B --> C["Extract path from 0 to 2"]
        C --> D["Cycle: 0 → 1 → 2 → 0"]
    end

    subgraph "Parent Tracking"
        E["parent[1] = 0"] --> F["parent[2] = 1"]
        F --> G["Use parent array to rebuild path"]
    end

    style S4 fill:#ffcdd2
    style A fill:#ffcdd2
    style D fill:#c8e6c9
```

## Step-by-Step Execution Example

```mermaid
graph LR
    subgraph "Step 1: Start DFS from node 0"
        A1[0: WHITE → GRAY] --> B1[Visit neighbor 1]
    end

    subgraph "Step 2: Process node 1"
        A2[1: WHITE → GRAY] --> B2[Visit neighbor 2]
    end

    subgraph "Step 3: Process node 2"
        A3[2: WHITE → GRAY] --> B3[Visit neighbor 0]
    end

    subgraph "Step 4: Back edge detected"
        A4["0 is GRAY - Back edge!"] --> B4["Cycle: 0→1→2→0"]
    end

    style A1 fill:#e1f5fe
    style A2 fill:#fff3e0
    style A3 fill:#fff3e0
    style A4 fill:#ffcdd2
    style B4 fill:#c8e6c9
```

## Cycle Normalization

```mermaid
graph TD
    subgraph "Raw Cycles Found"
        R1["Cycle: 0→1→2→0"]
        R2["Cycle: 1→2→0→1"]
        R3["Cycle: 2→0→1→2"]
    end

    subgraph "Normalization Process"
        N1["Find minimum node in cycle"] --> N2["Rotate cycle to start with min"]
        N2 --> N3["Canonical form: 0→1→2→0"]
    end

    subgraph "Duplicate Removal"
        D1["Use Set/Map for uniqueness"] --> D2["Only keep canonical forms"]
        D2 --> D3["Result: [0→1→2→0]"]
    end

    style R1 fill:#fff3e0
    style R2 fill:#fff3e0
    style R3 fill:#fff3e0
    style N3 fill:#c8e6c9
    style D3 fill:#c8e6c9
```

## Complexity Analysis

```mermaid
graph TD
    subgraph "Time Complexity"
        T1["DFS: O(V + E)"] --> T2["Each node visited once"]
        T2 --> T3["Each edge explored once"]
        T3 --> T4["Cycle extraction: O(V) per cycle"]
    end

    subgraph "Space Complexity"
        S1["Color array: O(V)"] --> S2["Parent array: O(V)"]
        S2 --> S3["Recursion stack: O(V)"]
        S3 --> S4["Total: O(V)"]
    end

    subgraph "Practical Performance"
        P1["Best case: O(V + E) - No cycles"]
        P2["Worst case: O(V + E) - Many cycles"]
        P3["Space efficient compared to other methods"]
    end

    style T1 fill:#e1f5fe
    style S1 fill:#e1f5fe
    style P1 fill:#c8e6c9
    style P2 fill:#c8e6c9
    style P3 fill:#c8e6c9
```

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
