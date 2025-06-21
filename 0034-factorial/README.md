# Factorial

## Description

A comprehensive implementation of factorial computation using multiple approaches including dynamic programming, memoization, iterative, and recursive methods. The implementation supports arbitrary precision arithmetic using Go's `math/big` package and includes additional factorial variants like double factorial and subfactorial.

## Factorial Concept

```mermaid
graph TD
    A[n! = n × (n-1) × (n-2) × ... × 2 × 1] --> B[Base Cases]
    B --> C[0! = 1]
    B --> D[1! = 1]
    A --> E[Examples]
    E --> F[5! = 5 × 4 × 3 × 2 × 1 = 120]
    E --> G[3! = 3 × 2 × 1 = 6]

    style A fill:#e1f5fe
    style C fill:#c8e6c9
    style D fill:#c8e6c9
    style F fill:#fff3e0
    style G fill:#fff3e0
```

## Recursive Breakdown Visualization

```mermaid
graph TD
    A[factorial(5)] --> B[5 × factorial(4)]
    B --> C[5 × 4 × factorial(3)]
    C --> D[5 × 4 × 3 × factorial(2)]
    D --> E[5 × 4 × 3 × 2 × factorial(1)]
    E --> F[5 × 4 × 3 × 2 × 1]
    F --> G[120]

    H[Call Stack] --> I[factorial(5) waits for factorial(4)]
    I --> J[factorial(4) waits for factorial(3)]
    J --> K[factorial(3) waits for factorial(2)]
    K --> L[factorial(2) waits for factorial(1)]
    L --> M[factorial(1) returns 1]
    M --> N[Stack unwinds with multiplication]

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style M fill:#c8e6c9
    style N fill:#c8e6c9
```

## Algorithm Comparison

```mermaid
graph LR
    subgraph "Iterative Approach"
        A1[Start: result = 1<br/>i = 1] --> A2[result *= i<br/>i++]
        A2 --> A3{i <= n?}
        A3 -->|Yes| A2
        A3 -->|No| A4[Return result]
    end

    subgraph "Dynamic Programming"
        B1[dp[0] = 1<br/>dp[1] = 1] --> B2[For i = 2 to n]
        B2 --> B3[dp[i] = i × dp[i-1]]
        B3 --> B4[Return dp[n]]
    end

    subgraph "Memoization"
        C1[Check memo[n]] --> C2{Exists?}
        C2 -->|Yes| C3[Return memo[n]]
        C2 -->|No| C4[Compute n × factorial(n-1)]
        C4 --> C5[Store in memo[n]]
        C5 --> C6[Return result]
    end

    style A1 fill:#e1f5fe
    style A4 fill:#c8e6c9
    style B1 fill:#e1f5fe
    style B4 fill:#c8e6c9
    style C1 fill:#e1f5fe
    style C3 fill:#c8e6c9
    style C6 fill:#c8e6c9
```

## Performance Analysis

```mermaid
graph TD
    subgraph "Time Complexity Comparison"
        A[Iterative: O(n)] --> A1[Single loop execution]
        B[Recursive: O(n)] --> B1[n function calls]
        C[DP: O(n)] --> C1[Fill table once]
        D[Memoized: O(1)] --> D1[After first computation]
    end

    subgraph "Space Complexity Comparison"
        E[Iterative: O(1)] --> E1[Constant space]
        F[Recursive: O(n)] --> F1[Call stack depth]
        G[DP: O(n)] --> G1[Array storage]
        H[Memoized: O(n)] --> H1[Cache storage]
    end

    style A fill:#c8e6c9
    style E fill:#c8e6c9
    style D fill:#c8e6c9
```

## Factorial Variants

```mermaid
graph TD
    subgraph "Double Factorial (n!!)"
        A[n!! = n × (n-2) × (n-4) × ...] --> B[Even: 8!! = 8×6×4×2 = 384]
        A --> C[Odd: 7!! = 7×5×3×1 = 105]
    end

    subgraph "Subfactorial (!n)"
        D[Derangements formula] --> E[!n = (n-1) × (!⁡(n-1) + !(n-2))]
        E --> F[!0 = 1, !1 = 0, !2 = 1]
        F --> G[!3 = 2, !4 = 9, !5 = 44]
    end

    subgraph "Falling Factorial"
        H[n^(k) = n × (n-1) × ... × (n-k+1)] --> I[5^(3) = 5×4×3 = 60]
    end

    style A fill:#e1f5fe
    style D fill:#e1f5fe
    style H fill:#e1f5fe
    style B fill:#fff3e0
    style C fill:#fff3e0
    style G fill:#fff3e0
    style I fill:#fff3e0
```

## Growth Rate Visualization

```mermaid
graph LR
    subgraph "Factorial Growth Pattern"
        A[1! = 1] --> B[2! = 2]
        B --> C[3! = 6]
        C --> D[4! = 24]
        D --> E[5! = 120]
        E --> F[10! = 3,628,800]
        F --> G[20! = 2.4 × 10¹⁸]
        G --> H[100! = 9.3 × 10¹⁵⁷]
    end

    subgraph "Growth Rate"
        I[Factorial grows faster than<br/>exponential functions] --> J[n! > 2ⁿ for n ≥ 4]
        J --> K[Stirling's Approximation:<br/>n! ≈ √(2πn) × (n/e)ⁿ]
    end

    style A fill:#e1f5fe
    style H fill:#ffcdd2
    style I fill:#fff3e0
    style K fill:#c8e6c9
```

## Real-World Applications

```mermaid
graph TD
    A[Factorial Applications] --> B[Combinatorics]
    A --> C[Probability]
    A --> D[Series Expansions]
    A --> E[Algorithm Analysis]

    B --> B1[Permutations: P(n,r) = n!/(n-r)!]
    B --> B2[Combinations: C(n,r) = n!/(r!(n-r)!)]

    C --> C1[Birthday Problem]
    C --> C2[Poisson Distribution]

    D --> D1[Taylor Series: eˣ = Σ(xⁿ/n!)]
    D --> D2[Sine/Cosine Expansions]

    E --> E1[Time Complexity of Algorithms]
    E --> E2[Recursive Function Analysis]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style B2 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style C2 fill:#c8e6c9
```

## Features

- **Multiple Algorithms**: Iterative, recursive, dynamic programming, and memoized approaches
- **Arbitrary Precision**: Uses `math/big` for handling very large factorials
- **Optimized DP**: Space-optimized dynamic programming with O(1) space complexity
- **Memoization**: Efficient caching with `FactorialCalculator` struct
- **Factorial Variants**: Double factorial (n!!), subfactorial (!n), and factorial range computation
- **Utility Functions**: Factorial detection, range computation, and error handling
- **Comprehensive Testing**: 25+ test functions with edge cases and performance benchmarks

## Algorithms Implemented

### 1. Dynamic Programming (FactorialDP)

- **Time Complexity**: O(n)
- **Space Complexity**: O(n)
- Builds factorial table from bottom-up

### 2. Optimized Dynamic Programming (FactorialDPOptimized)

- **Time Complexity**: O(n)
- **Space Complexity**: O(1)
- Uses only two variables instead of full array

### 3. Memoization (FactorialMemoized)

- **Time Complexity**: O(n) first call, O(1) subsequent calls
- **Space Complexity**: O(n) for memo table
- Top-down approach with caching

### 4. Iterative (FactorialIterative)

- **Time Complexity**: O(n)
- **Space Complexity**: O(1)
- Simple loop-based computation

### 5. Recursive (FactorialRecursive)

- **Time Complexity**: O(n)
- **Space Complexity**: O(n) call stack
- Classic recursive implementation

## Additional Features

### Double Factorial (n!!)

Computes n × (n-2) × (n-4) × ... × 2 or 1

### Subfactorial (!n)

Number of derangements of n objects using DP: !n = (n-1) × (!(n-1) + !(n-2))

### Factorial Detection

Determines if a given number is a factorial and returns the corresponding n

### Range Computation

Efficiently computes factorials for a range of values

## Usage

```bash
make run n=0034-factorial
```

## Testing

```bash
make test n=0034-factorial
```

## Performance Benchmarks

- **FactorialIterative**: ~500ns for n=20
- **FactorialDP**: ~800ns for n=20
- **FactorialMemoized**: ~50ns for n=20 (cached)
- **FactorialLarge**: ~2μs for n=100

## Real-World Applications

- **Combinatorics**: Permutations and combinations
- **Probability**: Statistical calculations
- **Mathematics**: Series expansions and mathematical proofs
- **Computer Science**: Algorithm analysis and complexity theory
- **Cryptography**: Key generation and mathematical foundations

## Error Handling

All functions return appropriate errors for:

- Negative inputs
- Invalid ranges
- Computation failures

## Example Output

```
0! = 1
1! = 1
5! = 120
10! = 3628800
100! = 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000 (158 digits)
```
