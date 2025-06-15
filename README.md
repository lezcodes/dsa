# Data Structures and Algorithms (DSA) Repository

A Go-based repository for implementing and testing various data structures and algorithms with an organized, serialized structure.

## Repository Structure

Each algorithm or data structure is contained in its own directory as a package within the main `github.com/celj/dsa` module:

```
XXXX-algorithm-name/
├── algorithm_name.go          # Implementation
├── algorithm_name_test.go     # Tests and benchmarks
└── README.md                  # Algorithm-specific documentation
```

**Example:**

```
0001-linear-search/
├── linear_search.go           # Linear search implementation
├── linear_search_test.go      # Comprehensive tests and benchmarks
└── README.md                  # Algorithm documentation
```

## Quick Start

### Creating a New Algorithm/Data Structure

```bash
make new n=binary-search
```

This will create a new directory `0002-binary-search` with all the necessary files. The directory number is automatically incremented based on existing directories.

### Testing

Test all algorithms:

```bash
make test
```

Test specific algorithms by name:

```bash
make test n=linear-search,binary-search
```

Test specific algorithms by index:

```bash
make test n=1,2,15,20
```

Mix names and indices:

```bash
make test n=linear-search,2,bubble-sort,12
```

### Running Algorithms

Run all algorithms:

```bash
make run
```

Run specific algorithms by name:

```bash
make run n=linear-search,binary-search
```

Run specific algorithms by index:

```bash
make run n=1,2,15,20
```

### Benchmarking

Run benchmarks for all algorithms:

```bash
make bench
```

Run benchmarks for specific algorithms by name:

```bash
make bench n=linear-search,binary-search
```

Run benchmarks for specific algorithms by index:

```bash
make bench n=1,2,15,20
```

### Other Commands

View available commands:

```bash
make help
```

## Selection Formats

The `n` parameter supports multiple selection formats:

- **By algorithm names**: `n=linear-search,binary-search`
- **By directory indices**: `n=1,2,15,20` (where 1 = 0001-_, 2 = 0002-_, etc.)
- **Mixed format**: `n=linear-search,2,bubble-sort,15`
- **Single selection**: `n=linear-search` or `n=1`
- **Full directory name**: `n=0001-linear-search`

## Features

- **Automatic Serialization**: Directories are automatically numbered (0001, 0002, etc.)
- **Complete Package Structure**: Each algorithm gets its own Go package within the main module
- **Flexible Selection**: Use algorithm names, indices, or mix both for multi-selection
- **Comma-separated Multi-selection**: Run multiple algorithms in a single command
- **Comprehensive Testing**: Includes unit tests, benchmarks, go vet, and formatting checks
- **Template Generation**: Automatically generates boilerplate code with TODO markers
- **Colorized Output**: Test runner provides clear, colorized feedback
- **Dynamic Execution**: Run command creates temporary main.go files with proper imports

## Example Usage

1. Create a new algorithm:

   ```bash
   make new n=binary-search
   ```

2. Navigate to the directory and implement:

   ```bash
   cd 0002-binary-search
   # Edit binary_search.go to implement your algorithm
   # Update tests in binary_search_test.go
   ```

3. Test your implementation:

   ```bash
   make test n=binary-search
   ```

4. Run your algorithm:

   ```bash
   make run n=binary-search
   ```

5. Run multiple algorithms:
   ```bash
   make run n=1,2,5
   make test n=linear-search,binary-search,bubble-sort
   make bench n=1,3,quick-sort,15
   ```

## Implemented Algorithms

### 0001-linear-search

A complete implementation of linear search with:

- Multiple data type support (int, string, float64, generic)
- Comprehensive test suite with edge cases
- Performance benchmarks for different array sizes
- Detailed complexity analysis and documentation

**Usage:**

```bash
make run n=linear-search
# or
make run n=1
# Output: map[array:[64 34 25 12 22 11 90] found:true index:4 target:22]

make test n=linear-search
# or
make test n=1
# Runs comprehensive tests and benchmarks
```

## Directory Contents

When you create a new algorithm, you'll get:

- **algorithm_name.go**: Your implementation goes here
- **algorithm_name_test.go**: Unit tests and benchmarks
- **README.md**: Documentation template

## Test Suite

The test runner performs:

- Unit tests (`go test -v`)
- Benchmarks (`go test -bench=. -benchmem`)
- Code analysis (`go vet`)
- Format checking (`gofmt`)

## How It Works

- Each directory is a Go package within the `github.com/celj/dsa` module
- The `make run` command creates a temporary main.go file with proper import aliasing
- Import paths use the directory name: `github.com/celj/dsa/XXXX-algorithm-name`
- Package names use underscores: `algorithm_name`
- Selection supports both full directory names, algorithm names, and numeric indices
- Temporary files are automatically cleaned up after execution

## Repository Scripts

The repository uses three main scripts:

- `scripts/new.sh` - Creates new algorithm directories
- `scripts/test.sh` - Runs tests and benchmarks with multi-selection support
- `scripts/run.sh` - Executes algorithms with multi-selection support
- `scripts/bench.sh` - Runs benchmarks with multi-selection support

## Contributing

1. Create a new algorithm with `make new n=your-algorithm`
2. Implement the algorithm in the generated files
3. Add comprehensive tests and benchmarks
4. Update the algorithm's README.md with complexity analysis
5. Run `make test` to ensure everything passes

## Best Practices

Based on the linear-search implementation:

**Code Structure:**

- Provide multiple functions for different data types
- Include a generic implementation when possible
- Use clear, descriptive function names
- Return consistent result formats

**Testing:**

- Test all edge cases (empty arrays, single elements, not found)
- Include benchmarks for different input sizes
- Use table-driven tests for comprehensive coverage
- Test both success and failure scenarios

**Documentation:**

- Include algorithm description and complexity analysis
- Provide usage examples with code snippets
- Document when to use vs. alternatives
- Include performance characteristics
