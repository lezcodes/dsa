# Stack

## Description

A Stack is a linear data structure that follows the Last In, First Out (LIFO) principle. Elements are added (pushed) and removed (popped) from the same end, called the top of the stack.

## Visual Representation

### Stack Structure

```mermaid
graph TD
    Top["Top (Push/Pop)"] --> A[40]
    A --> B[30]
    B --> C[20]
    C --> D[10]
    D --> Bottom[Bottom]

    style Top fill:#e1f5fe
    style Bottom fill:#f5f5f5
    style A fill:#c8e6c9
    style B fill:#e8f5e8
    style C fill:#f1f8e9
    style D fill:#f9f9f9
```

### Push Operation

```mermaid
graph TD
    A[New Element: 50] --> B[Check if stack is full]
    B --> C{Is Full?}
    C -->|Yes| D[Return Stack Overflow Error]
    C -->|No| E[Increment top pointer]
    E --> F[Add element at top]
    F --> G[Increment size]
    G --> H[Operation Complete]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style D fill:#ffcdd2
```

### Pop Operation

```mermaid
graph TD
    A[Pop Request] --> B[Check if stack is empty]
    B --> C{Is Empty?}
    C -->|Yes| D[Return Stack Underflow Error]
    C -->|No| E[Get element from top]
    E --> F[Decrement top pointer]
    F --> G[Decrement size]
    G --> H[Return element]

    style A fill:#e1f5fe
    style H fill:#c8e6c9
    style D fill:#ffcdd2
```

### Stack Operations Flow

```mermaid
graph LR
    subgraph "Initial State"
        S1[Empty Stack]
    end

    subgraph "After Push(10, 20, 30)"
        S2["30 ← Top<br/>20<br/>10"]
    end

    subgraph "After Pop()"
        S3["20 ← Top<br/>10"]
    end

    subgraph "After Push(40)"
        S4["40 ← Top<br/>20<br/>10"]
    end

    S1 --> S2
    S2 --> S3
    S3 --> S4

    style S1 fill:#f5f5f5
    style S4 fill:#e8f5e8
```

### LIFO Principle Visualization

```mermaid
graph LR
    subgraph "Stack of Plates"
        direction TB
        P1[Plate 3 - Last In]
        P2[Plate 2]
        P3[Plate 1 - First In]
    end

    subgraph "Removal Order"
        direction TB
        R1[First Out ← Plate 3]
        R2[Second Out ← Plate 2]
        R3[Last Out ← Plate 1]
    end

    P1 -.->|Pop| R1
    P2 -.->|Pop| R2
    P3 -.->|Pop| R3

    style P1 fill:#ffcdd2
    style R1 fill:#ffcdd2
```

### Applications Visualization

```mermaid
graph TD
    A[Stack Applications] --> B[Function Call Stack]
    A --> C[Expression Evaluation]
    A --> D[Undo Operations]
    A --> E[Browser History]
    A --> F[Balanced Parentheses]

    B --> B1["main() calls func1()<br/>func1() calls func2()<br/>Return in reverse order"]
    C --> C1["Convert infix to postfix<br/>Evaluate postfix expressions"]
    D --> D1["Text editor undo<br/>Game state restoration"]
    E --> E1["Back button navigation<br/>Page history tracking"]
    F --> F1["Check matching brackets<br/>Validate syntax"]

    style A fill:#e1f5fe
```

A Stack is a linear data structure that follows the Last In, First Out (LIFO) principle.

This implementation provides three different stack variants:

1. **LinkedListStack**: Uses a linked list with a top pointer
2. **ArrayStack**: Uses a fixed-size array with capacity management
3. **DynamicStack**: Uses Go's slice with automatic resizing

## Key Operations

- **Push**: Add element to the top of the stack
- **Pop**: Remove and return element from the top of the stack
- **Peek**: Get the top element without removing it
- **Size**: Get the number of elements in the stack
- **IsEmpty**: Check if the stack is empty
- **Clear**: Remove all elements from the stack

## Complexity

### LinkedListStack

- **Push**: O(1) - Constant time insertion at top
- **Pop**: O(1) - Constant time removal from top
- **Peek**: O(1) - Direct access to top pointer
- **Space**: O(n) - Linear space for n elements

### ArrayStack (Fixed Size)

- **Push**: O(1) - Constant time insertion (when not full)
- **Pop**: O(1) - Constant time removal
- **Peek**: O(1) - Direct array access
- **Space**: O(capacity) - Fixed space based on capacity

### DynamicStack

- **Push**: O(1) amortized - May require slice expansion
- **Pop**: O(1) - Direct slice modification
- **Peek**: O(1) - Direct slice access
- **Space**: O(n) - Dynamic space allocation

## Implementation Details

### LinkedListStack

Uses a singly-linked list where:

- `top` points to the most recently added node
- New nodes are inserted at the beginning (top)
- Maintains a `size` counter for O(1) size queries

### ArrayStack

Uses a fixed-size array where:

- `top` index tracks the position of the top element (-1 when empty)
- Has a fixed capacity to prevent unbounded growth
- Provides `IsFull()` method to check capacity limits

### DynamicStack

Uses Go's built-in slice operations:

- Appends to the end for push operations
- Uses slice re-slicing for pop operations
- Automatically handles memory allocation and expansion

## Performance Comparison

| Operation | LinkedList | Array (Fixed) | Dynamic  |
| --------- | ---------- | ------------- | -------- |
| Push      | O(1)       | O(1)          | O(1)\*   |
| Pop       | O(1)       | O(1)          | O(1)     |
| Peek      | O(1)       | O(1)          | O(1)     |
| Memory    | Variable   | Fixed         | Variable |

\*Amortized complexity

## Real-World Applications

### Function Call Stack

```go
callStack := NewLinkedListStack()
callStack.Push(functionA)
callStack.Push(functionB)
returnTo, _ := callStack.Pop()
```

### Expression Evaluation

```go
result, err := EvaluatePostfix([]string{"3", "4", "+", "2", "*"})
// result = 14, evaluates (3 + 4) * 2
```

### Balanced Parentheses Checking

```go
isBalanced := IsBalancedParentheses("({[]})")
// isBalanced = true
```

### Undo/Redo Operations

```go
undoStack := NewDynamicStack()
undoStack.Push(lastAction)
if !undoStack.IsEmpty() {
    action, _ := undoStack.Pop()
    // Undo the action
}
```

### Browser History

```go
history := NewArrayStack(50)
history.Push(currentPage)
previousPage, _ := history.Pop()
```

## Advanced Features

### Postfix Expression Evaluation

The implementation includes a complete postfix (Reverse Polish Notation) evaluator:

- Supports +, -, \*, / operations
- Handles division by zero errors
- Validates expression correctness

### Balanced Parentheses Checker

Validates if parentheses, brackets, and braces are properly balanced:

- Supports (), [], {} bracket types
- Handles nested structures
- Ignores non-bracket characters

## Usage Examples

### Basic Stack Operations

```go
stack := NewLinkedListStack()
stack.Push(10)
stack.Push(20)
stack.Push(30)

top, _ := stack.Peek()    // top = 30
popped, _ := stack.Pop()  // popped = 30
fmt.Println(stack.Display()) // Stack: [20 | 10] (top | bottom)
```

### Postfix Evaluation

```go
expression := []string{"15", "7", "1", "1", "+", "-", "/", "3", "*"}
result, err := EvaluatePostfix(expression)
// Evaluates: 15 / (7 - (1 + 1)) * 3 = 9
```

### Parentheses Validation

```go
expressions := []string{
    "({[]})",    // true - balanced
    "({[})",     // false - wrong order
    "(((",       // false - unbalanced
}

for _, expr := range expressions {
    balanced := IsBalancedParentheses(expr)
    fmt.Printf("%s: %v\n", expr, balanced)
}
```

## When to Use Each Implementation

### LinkedListStack

- **Best for**: General-purpose stacking, unlimited size requirements
- **Pros**: No capacity limit, O(1) all operations, memory efficient
- **Cons**: Extra memory overhead for pointers

### ArrayStack

- **Best for**: High-performance scenarios with known capacity limits
- **Pros**: Cache-friendly, minimal memory overhead, fastest operations
- **Cons**: Fixed capacity, potential for stack overflow errors

### DynamicStack

- **Best for**: Variable workloads, simple implementations
- **Pros**: Automatic resizing, simple to understand
- **Cons**: Potential memory overhead from slice expansion

## Usage

```bash
make run NAME=0008-stack
```

## Testing

```bash
make test NAME=0008-stack
```

## Practical Applications

- **Compiler Design**: Function call management, operator precedence
- **Algorithm Implementation**: Depth-First Search, backtracking
- **System Programming**: Memory management, interrupt handling
- **Text Processing**: Parsing, syntax checking, markup validation
- **Mathematical Computing**: Expression evaluation, calculator implementations
