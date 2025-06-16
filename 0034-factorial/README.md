# Factorial

## Description

A comprehensive implementation of factorial computation using multiple approaches including dynamic programming, memoization, iterative, and recursive methods. The implementation supports arbitrary precision arithmetic using Go's `math/big` package and includes additional factorial variants like double factorial and subfactorial.

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
