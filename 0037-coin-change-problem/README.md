# Coin Change Problem

## Description

Implementation of the Coin Change Problem using Dynamic Programming.

Given an integer array of coin denominations and a target amount, finds the minimum number of coins needed to make up that amount. Returns -1 if the amount cannot be made with the given coins.

## Visual Representation

### Dynamic Programming Algorithm Flow

```mermaid
graph TD
    A["Start: coins array, amount"] --> B["Initialize dp array size amount+1"]
    B --> C["Set dp[0] = 0"]
    C --> D["Set dp[i] = infinity for i > 0"]
    D --> E["For each amount i from 1 to target"]
    E --> F["For each coin in coins"]
    F --> G{"coin <= i?"}
    G -->|No| H["Skip this coin"]
    G -->|Yes| I["Calculate dp[i-coin] + 1"]
    I --> J{"New value < dp[i]?"}
    J -->|Yes| K["Update dp[i]"]
    J -->|No| L["Keep current dp[i]"]
    K --> M["Store coin used"]
    H --> N{"More coins?"}
    L --> N
    M --> N
    N -->|Yes| F
    N -->|No| O{"More amounts?"}
    O -->|Yes| E
    O -->|No| P{"dp[amount] == infinity?"}
    P -->|Yes| Q["Return -1"]
    P -->|No| R["Return dp[amount]"]

    style A fill:#e1f5fe
    style R fill:#c8e6c9
    style Q fill:#ffcdd2
```

### DP Table Construction Example

```mermaid
graph LR
    subgraph "Coins: [1, 3, 4], Amount: 6"
        A["dp[0] = 0"]
        B["dp[1] = 1 (coin: 1)"]
        C["dp[2] = 2 (coins: 1,1)"]
        D["dp[3] = 1 (coin: 3)"]
        E["dp[4] = 1 (coin: 4)"]
        F["dp[5] = 2 (coins: 4,1)"]
        G["dp[6] = 2 (coins: 3,3)"]
    end

    subgraph "Optimal Solution"
        H["Amount 6 = coin 3 + coin 3"] --> H1["coin 3"]
        I["Minimum coins: 2"]
    end

    A --> B --> C --> D --> E --> F --> G
    G -.-> H

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style I fill:#c8e6c9
```

### Decision Making Process

```mermaid
graph TD
    A["For amount i, coin c"] --> B{"c <= i?"}
    B -->|No| C["Cannot use this coin"]
    B -->|Yes| D["Check dp[i-c]"]
    D --> E{"dp[i-c] is valid?"}
    E -->|No| F["This coin path invalid"]
    E -->|Yes| G["Calculate: dp[i-c] + 1"]
    G --> H{"New value < current dp[i]?"}
    H -->|Yes| I["Update dp[i] = dp[i-c] + 1"]
    H -->|No| J["Keep current dp[i]"]

    style A fill:#e1f5fe
    style I fill:#c8e6c9
    style C fill:#ffcdd2
    style F fill:#ffcdd2
```

### Step-by-Step Example

```mermaid
graph TD
    subgraph "Example: coins=[1,3,4], amount=6"
        A1["dp[0] = 0 (base case)"]
        A2["dp[1]: min(∞, dp[0]+1) = 1"]
        A3["dp[2]: min(∞, dp[1]+1) = 2"]
        A4["dp[3]: min(dp[2]+1, dp[0]+1) = min(3,1) = 1"]
        A5["dp[4]: min(dp[3]+1, dp[1]+1, dp[0]+1) = min(2,2,1) = 1"]
        A6["dp[5]: min(dp[4]+1, dp[2]+1, dp[1]+1) = min(2,3,2) = 2"]
        A7["dp[6]: min(dp[5]+1, dp[3]+1, dp[2]+1) = min(3,2,3) = 2"]
    end

    A1 --> A2 --> A3 --> A4 --> A5 --> A6 --> A7

    style A1 fill:#e1f5fe
    style A7 fill:#c8e6c9
```

### Solution Reconstruction

```mermaid
graph TD
    A["Backtrack from dp[amount]"] --> B["Find which coin was used"]
    B --> C["Subtract coin value from amount"]
    C --> D["Repeat until amount = 0"]
    D --> E["Collect all coins used"]

    subgraph "Example: amount=6, coins=[1,3,4]"
        F["dp[6] = 2 (used coin 3)"] --> F1["coin 3"]
        G["amount = 6-3 = 3"] --> G1["amount = 3"]
        H["dp[3] = 1 (used coin 3)"] --> H1["coin 3"]
        I["amount = 3-3 = 0"]
        J["Solution: [3, 3]"] --> J1["[3, 3]"]
    end

    A --> F

    style A fill:#e1f5fe
    style J fill:#c8e6c9
```

### Time and Space Complexity

```mermaid
graph LR
    subgraph "Complexity Analysis"
        A["Time: O(amount × coins)"]
        B["Space: O(amount)"]
    end

    subgraph "DP Table Dimensions"
        C["Rows: amount + 1"]
        D["Operations per cell: |coins|"]
        E["Total operations: amount × |coins|"]
    end

    subgraph "Space Optimization Possible"
        F["Only need previous values"]
        G["Can optimize to O(min(amount, max_coin))"]
    end

    style A fill:#c8e6c9
    style B fill:#c8e6c9
    style G fill:#fff3e0
```

### Comparison with Other Approaches

```mermaid
graph TD
    A["Coin Change Approaches"] --> B["Dynamic Programming"]
    A --> C["Greedy Algorithm"]
    A --> D["Recursive (Brute Force)"]

    B --> B1["Time: O(amount × coins)"]
    B --> B2["Space: O(amount)"]
    B --> B3["Guarantees optimal solution"]
    B --> B4["Works for any coin system"]

    C --> C1["Time: O(coins log coins)"]
    C --> C2["Space: O(1)"]
    C --> C3["May not work for all coin systems"]
    C --> C4["Fast but not always correct"]

    D --> D1["Time: O(coins^amount)"]
    D --> D2["Space: O(amount)"]
    D --> D3["Exponential time complexity"]
    D --> D4["Simple but impractical"]

    style B3 fill:#c8e6c9
    style B4 fill:#c8e6c9
    style C3 fill:#ffcdd2
    style D3 fill:#ffcdd2
```

### Real-World Applications

```mermaid
graph TD
    A["Coin Change Applications"] --> B["Currency Systems"]
    A --> C["Resource Allocation"]
    A --> D["Network Routing"]
    A --> E["Manufacturing"]

    B --> B1["Making change with fewest coins<br/>Currency conversion<br/>ATM cash dispensing"]

    C --> C1["Minimum resources to meet quota<br/>Budget optimization<br/>Inventory management"]

    D --> D1["Minimum hops in networks<br/>Bandwidth allocation<br/>Packet routing"]

    E --> E1["Production planning<br/>Material optimization<br/>Cost minimization"]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
    style E1 fill:#c8e6c9
```

The algorithm uses bottom-up dynamic programming:

- `dp[i]` represents the minimum number of coins needed to make amount `i`
- For each amount from 1 to target, tries all coin denominations
- Takes the minimum of all possible combinations
- Also tracks the actual coins used in the optimal solution

Key insight: To make amount `i`, we can use any coin `c` (where `c <= i`) and add it to the optimal solution for amount `i - c`.

## Complexity

- Time Complexity: O(amount × number of coins) - for each amount, we check all coin denominations
- Space Complexity: O(amount) for the DP array, plus O(amount) for tracking the solution path

## Usage

```bash
make run n=0037-coin-change-problem
```

## Testing

```bash
make test n=0037-coin-change-problem
```
