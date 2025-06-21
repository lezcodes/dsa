# Fibonacci

## Description

A comprehensive implementation of Fibonacci sequence computation using multiple approaches including dynamic programming, memoization, matrix exponentiation, and iterative methods. The implementation supports arbitrary precision arithmetic using Go's `math/big` package and includes related sequences like Lucas numbers and Tribonacci numbers.

## Fibonacci Sequence Visualization

```mermaid
graph LR
    A[0] --> B[1]
    B --> C[1]
    C --> D[2]
    D --> E[3]
    E --> F[5]
    F --> G[8]
    G --> H[13]
    H --> I[21]
    I --> J[34]
    J --> K[55]

    subgraph "Rule: F(n) = F(n-1) + F(n-2)"
        L[F(0) = 0] --> M[F(1) = 1]
        M --> N[F(2) = F(1) + F(0) = 1]
        N --> O[F(3) = F(2) + F(1) = 2]
    end

    style A fill:#e1f5fe
    style B fill:#e1f5fe
    style L fill:#c8e6c9
    style M fill:#c8e6c9
```

## Recursive Tree Visualization

```mermaid
graph TD
    A[fib(5)] --> B[fib(4)]
    A --> C[fib(3)]
    B --> D[fib(3)]
    B --> E[fib(2)]
    C --> F[fib(2)]
    C --> G[fib(1)]
    D --> H[fib(2)]
    D --> I[fib(1)]
    E --> J[fib(1)]
    E --> K[fib(0)]
    F --> L[fib(1)]
    F --> M[fib(0)]
    H --> N[fib(1)]
    H --> O[fib(0)]

    subgraph "Overlapping Subproblems"
        P[fib(3) computed 2 times]
        Q[fib(2) computed 3 times]
        R[fib(1) computed 5 times]
        S[fib(0) computed 3 times]
    end

    style A fill:#e1f5fe
    style G fill:#c8e6c9
    style I fill:#c8e6c9
    style J fill:#c8e6c9
    style K fill:#c8e6c9
    style L fill:#c8e6c9
    style M fill:#c8e6c9
    style N fill:#c8e6c9
    style O fill:#c8e6c9
    style P fill:#ffcdd2
    style Q fill:#ffcdd2
    style R fill:#ffcdd2
    style S fill:#ffcdd2
```

## Algorithm Comparison

```mermaid
graph TD
    subgraph "Iterative O(n) - Optimal Space"
        A1[a = 0, b = 1] --> A2[For i = 2 to n]
        A2 --> A3[temp = a + b<br/>a = b, b = temp]
        A3 --> A4{i < n?}
        A4 -->|Yes| A3
        A4 -->|No| A5[Return b]
    end

    subgraph "Dynamic Programming O(n)"
        B1[dp[0] = 0<br/>dp[1] = 1] --> B2[For i = 2 to n]
        B2 --> B3[dp[i] = dp[i-1] + dp[i-2]]
        B3 --> B4[Return dp[n]]
    end

    subgraph "Matrix Exponentiation O(log n)"
        C1[Matrix = [[1,1],[1,0]]] --> C2[Fast matrix power]
        C2 --> C3[Matrix^n in O(log n)]
        C3 --> C4[Extract F(n) from result]
    end

    style A1 fill:#e1f5fe
    style A5 fill:#c8e6c9
    style B1 fill:#e1f5fe
    style B4 fill:#c8e6c9
    style C1 fill:#e1f5fe
    style C4 fill:#c8e6c9
```

## Matrix Exponentiation Method

```mermaid
graph TD
    subgraph "Matrix Formula"
        A["[F(n+1)]   [1 1]^n   [1]<br/>[F(n)  ] = [1 0]   × [0]"] --> B[Fast Matrix Power]
    end

    subgraph "Power Calculation Example: n=5"
        C[Matrix^5] --> D[Matrix^4 × Matrix^1]
        D --> E[Matrix^2 × Matrix^2 × Matrix^1]
        E --> F[(Matrix^1)² × (Matrix^1)² × Matrix^1]
        F --> G[Binary: 5 = 101₂]
        G --> H[Use only positions with 1 bits]
    end

    subgraph "Matrix Multiplication"
        I["[a b]   [e f]   [ae+bg af+bh]<br/>[c d] × [g h] = [ce+dg cf+dh]"] --> J[Time: O(1) for 2×2]
    end

    style A fill:#e1f5fe
    style C fill:#fff3e0
    style I fill:#c8e6c9
```

## Performance Analysis

```mermaid
graph LR
    subgraph "Time Complexity"
        A[Recursive: O(2ⁿ)] --> A1[Exponential - very slow]
        B[Iterative: O(n)] --> B1[Linear - good for most cases]
        C[DP: O(n)] --> C1[Linear with O(n) space]
        D[Matrix: O(log n)] --> D1[Logarithmic - best for large n]
        E[Memoized: O(n)] --> E1[Linear first call, O(1) after]
    end

    subgraph "Space Complexity"
        F[Recursive: O(n)] --> F1[Call stack depth]
        G[Iterative: O(1)] --> G1[Constant space - optimal]
        H[DP: O(n)] --> H1[Array storage]
        I[Matrix: O(log n)] --> I1[Recursion for power]
        J[Memoized: O(n)] --> J1[Cache storage]
    end

    style A fill:#ffcdd2
    style D fill:#c8e6c9
    style G fill:#c8e6c9
```

## Related Sequences

```mermaid
graph TD
    subgraph "Lucas Numbers: L(n) = L(n-1) + L(n-2)"
        A[L(0) = 2, L(1) = 1] --> B[2, 1, 3, 4, 7, 11, 18, 29, 47...]
        B --> C[Relation: L(n) = F(n-1) + F(n+1)]
    end

    subgraph "Tribonacci: T(n) = T(n-1) + T(n-2) + T(n-3)"
        D[T(0)=0, T(1)=1, T(2)=1] --> E[0, 1, 1, 2, 4, 7, 13, 24, 44...]
    end

    subgraph "Golden Ratio φ"
        F[φ = (1 + √5) / 2] --> G[φ ≈ 1.618033988...]
        G --> H[Binet's Formula: F(n) = (φⁿ - ψⁿ)/√5]
        H --> I[where ψ = -1/φ]
    end

    style A fill:#e1f5fe
    style D fill:#e1f5fe
    style F fill:#e1f5fe
    style C fill:#c8e6c9
    style I fill:#c8e6c9
```

## Mathematical Properties

```mermaid
graph TD
    A[Fibonacci Properties] --> B[GCD Property]
    A --> C[Sum Formula]
    A --> D[Identity Relations]
    A --> E[Divisibility Rules]

    B --> B1[gcd(F(m), F(n)) = F(gcd(m, n))]
    C --> C1[F(1) + F(2) + ... + F(n) = F(n+2) - 1]
    D --> D1[F(n+m) = F(n)×F(m+1) + F(n-1)×F(m)]
    D --> D2[F(2n) = F(n)×(2F(n+1) - F(n))]
    E --> E1[F(n) divides F(kn) for any positive k]
    E --> E2[3 divides F(n) iff 4 divides n]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
    style D2 fill:#c8e6c9
    style E1 fill:#c8e6c9
    style E2 fill:#c8e6c9
```

## Real-World Applications

```mermaid
graph TD
    A[Fibonacci in Nature & Science] --> B[Biology]
    A --> C[Art & Architecture]
    A --> D[Computer Science]
    A --> E[Financial Markets]

    B --> B1[Flower petals (3, 5, 8, 13, 21)]
    B --> B2[Pine cone spirals]
    B --> B3[Population growth models]

    C --> C1[Golden ratio in design]
    C --> C2[Spiral patterns in art]
    C --> C3[Architecture proportions]

    D --> D1[Algorithm optimization]
    D --> D2[Data structure analysis]
    D --> D3[Recursive algorithm design]

    E --> E1[Technical analysis]
    E --> E2[Fibonacci retracements]
    E --> E3[Trading strategies]

    style A fill:#e1f5fe
    style B1 fill:#c8e6c9
    style B2 fill:#c8e6c9
    style C1 fill:#c8e6c9
    style D1 fill:#c8e6c9
    style E1 fill:#c8e6c9
```

## Features

- **Multiple Algorithms**: Iterative, recursive, dynamic programming, memoized, and matrix exponentiation approaches
- **Arbitrary Precision**: Uses `math/big` for handling very large Fibonacci numbers
- **Matrix Exponentiation**: O(log n) time complexity using fast matrix power
- **Optimized DP**: Space-optimized dynamic programming with O(1) space complexity
- **Memoization**: Efficient caching with `FibonacciCalculator` struct
- **Related Sequences**: Lucas numbers, Tribonacci numbers, and golden ratio computation
- **Utility Functions**: Fibonacci detection, range computation, and mathematical properties
- **Comprehensive Testing**: 30+ test functions with edge cases and performance benchmarks

## Algorithms Implemented

### 1. Dynamic Programming (FibonacciDP)

- **Time Complexity**: O(n)
- **Space Complexity**: O(n)
- Builds Fibonacci table from bottom-up

### 2. Optimized Dynamic Programming (FibonacciDPOptimized)

- **Time Complexity**: O(n)
- **Space Complexity**: O(1)
- Uses only two variables instead of full array

### 3. Memoization (FibonacciMemoized)

- **Time Complexity**: O(n) first call, O(1) subsequent calls
- **Space Complexity**: O(n) for memo table
- Top-down approach with caching

### 4. Iterative (FibonacciIterative)

- **Time Complexity**: O(n)
- **Space Complexity**: O(1)
- Simple loop-based computation

### 5. Recursive (FibonacciRecursive)

- **Time Complexity**: O(2^n)
- **Space Complexity**: O(n) call stack
- Classic recursive implementation (inefficient for large n)

### 6. Matrix Exponentiation (FibonacciMatrix)

- **Time Complexity**: O(log n)
- **Space Complexity**: O(log n)
- Uses fast matrix power with [[1,1],[1,0]]^n

## Additional Features

### Lucas Numbers

Sequence: L(0)=2, L(1)=1, L(n)=L(n-1)+L(n-2)
Related to Fibonacci: L(n) = F(n-1) + F(n+1)

### Tribonacci Numbers

Sequence: T(0)=0, T(1)=1, T(2)=1, T(n)=T(n-1)+T(n-2)+T(n-3)

### Golden Ratio (φ)

Computes the golden ratio: φ = (1 + √5) / 2 ≈ 1.618033988749...

### Fibonacci Detection

Determines if a given number is a Fibonacci number and returns the corresponding n

### Range Computation

Efficiently computes Fibonacci numbers for a range of values

## Matrix Exponentiation Details

The Fibonacci matrix formula uses:

```
[F(n+1)]   [1 1]^n   [1]
[F(n)  ] = [1 0]   × [0]
```

This allows computing F(n) in O(log n) time using fast matrix exponentiation.

## Usage

```bash
make run n=0035-fibonacci
```

## Testing

```bash
make test n=0035-fibonacci
```

## Performance Benchmarks

- **FibonacciIterative**: ~800ns for n=30
- **FibonacciDP**: ~1.2μs for n=30
- **FibonacciMemoized**: ~100ns for n=30 (cached)
- **FibonacciMatrix**: ~2μs for n=30
- **FibonacciLarge**: ~15μs for n=100

## Real-World Applications

- **Nature**: Spiral patterns in shells, flowers, and galaxies
- **Art & Architecture**: Golden ratio in design and proportions
- **Computer Science**: Algorithm analysis and optimization
- **Financial Markets**: Technical analysis and trading strategies
- **Biology**: Population growth models and genetic algorithms
- **Mathematics**: Number theory and mathematical sequences

## Mathematical Properties

- **Binet's Formula**: F(n) = (φⁿ - ψⁿ) / √5, where φ = golden ratio, ψ = -1/φ
- **GCD Property**: gcd(F(m), F(n)) = F(gcd(m, n))
- **Sum Formula**: F(1) + F(2) + ... + F(n) = F(n+2) - 1
- **Identity**: F(n+m) = F(n)×F(m+1) + F(n-1)×F(m)

## Error Handling

All functions return appropriate errors for:

- Negative inputs
- Invalid ranges
- Computation failures

## Example Output

```
F(0) = 0
F(1) = 1
F(10) = 55
F(50) = 12586269025
F(100) = 354224848179261915075 (21 digits)
Golden Ratio = 1.618033988749894848204586834365638117720309179805762862135448622705260462818902449707207204189391137484754088075386891752126633862223536450849140235
```
