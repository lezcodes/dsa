# Fibonacci

## Description

A comprehensive implementation of Fibonacci sequence computation using multiple approaches including dynamic programming, memoization, matrix exponentiation, and iterative methods. The implementation supports arbitrary precision arithmetic using Go's `math/big` package and includes related sequences like Lucas numbers and Tribonacci numbers.

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
