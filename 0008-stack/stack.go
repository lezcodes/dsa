package stack

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedListStack struct {
	top  *Node
	size int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		top:  nil,
		size: 0,
	}
}

func (s *LinkedListStack) Push(data int) {
	newNode := &Node{Data: data, Next: s.top}
	s.top = newNode
	s.size++
}

func (s *LinkedListStack) Pop() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("stack is empty")
	}

	data := s.top.Data
	s.top = s.top.Next
	s.size--

	return data, nil
}

func (s *LinkedListStack) Peek() (int, error) {
	if s.top == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.top.Data, nil
}

func (s *LinkedListStack) Size() int {
	return s.size
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedListStack) Clear() {
	s.top = nil
	s.size = 0
}

func (s *LinkedListStack) ToSlice() []int {
	result := make([]int, 0, s.size)
	current := s.top

	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}

	return result
}

func (s *LinkedListStack) Display() string {
	if s.top == nil {
		return "Stack: []"
	}

	result := "Stack: ["
	current := s.top

	for current != nil {
		result += fmt.Sprintf("%d", current.Data)
		if current.Next != nil {
			result += " | "
		}
		current = current.Next
	}

	result += "] (top | bottom)"
	return result
}

type ArrayStack struct {
	data     []int
	top      int
	capacity int
}

func NewArrayStack(capacity int) *ArrayStack {
	if capacity <= 0 {
		capacity = 10
	}

	return &ArrayStack{
		data:     make([]int, capacity),
		top:      -1,
		capacity: capacity,
	}
}

func (s *ArrayStack) Push(data int) error {
	if s.top >= s.capacity-1 {
		return fmt.Errorf("stack is full")
	}

	s.top++
	s.data[s.top] = data

	return nil
}

func (s *ArrayStack) Pop() (int, error) {
	if s.top < 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	data := s.data[s.top]
	s.top--

	return data, nil
}

func (s *ArrayStack) Peek() (int, error) {
	if s.top < 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[s.top], nil
}

func (s *ArrayStack) Size() int {
	return s.top + 1
}

func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

func (s *ArrayStack) IsFull() bool {
	return s.top >= s.capacity-1
}

func (s *ArrayStack) Clear() {
	s.top = -1
}

func (s *ArrayStack) ToSlice() []int {
	result := make([]int, 0, s.Size())

	for i := s.top; i >= 0; i-- {
		result = append(result, s.data[i])
	}

	return result
}

func (s *ArrayStack) Display() string {
	if s.top < 0 {
		return "Stack: []"
	}

	result := "Stack: ["

	for i := s.top; i >= 0; i-- {
		result += fmt.Sprintf("%d", s.data[i])
		if i > 0 {
			result += " | "
		}
	}

	result += "] (top | bottom)"
	return result
}

func (s *ArrayStack) Capacity() int {
	return s.capacity
}

type DynamicStack struct {
	data []int
}

func NewDynamicStack() *DynamicStack {
	return &DynamicStack{
		data: make([]int, 0),
	}
}

func (s *DynamicStack) Push(data int) {
	s.data = append(s.data, data)
}

func (s *DynamicStack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return data, nil
}

func (s *DynamicStack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *DynamicStack) Size() int {
	return len(s.data)
}

func (s *DynamicStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *DynamicStack) Clear() {
	s.data = s.data[:0]
}

func (s *DynamicStack) ToSlice() []int {
	result := make([]int, len(s.data))

	for i := len(s.data) - 1; i >= 0; i-- {
		result[len(s.data)-1-i] = s.data[i]
	}

	return result
}

func (s *DynamicStack) Display() string {
	if len(s.data) == 0 {
		return "Stack: []"
	}

	result := "Stack: ["

	for i := len(s.data) - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d", s.data[i])
		if i > 0 {
			result += " | "
		}
	}

	result += "] (top | bottom)"
	return result
}

func EvaluatePostfix(expression []string) (int, error) {
	stack := NewDynamicStack()

	for _, token := range expression {
		switch token {
		case "+":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a + b)
		case "-":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a - b)
		case "*":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			stack.Push(a * b)
		case "/":
			if stack.Size() < 2 {
				return 0, fmt.Errorf("invalid postfix expression")
			}
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			if b == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			stack.Push(a / b)
		default:
			var num int
			_, err := fmt.Sscanf(token, "%d", &num)
			if err != nil {
				return 0, fmt.Errorf("invalid token: %s", token)
			}
			stack.Push(num)
		}
	}

	if stack.Size() != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack.Pop()
}

func IsBalancedParentheses(expression string) bool {
	stack := NewDynamicStack()
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range expression {
		switch char {
		case '(', '{', '[':
			stack.Push(int(char))
		case ')', '}', ']':
			if stack.IsEmpty() {
				return false
			}

			top, _ := stack.Pop()
			if rune(top) != pairs[char] {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

func Run() any {
	llStack := NewLinkedListStack()
	arrayStack := NewArrayStack(5)
	dynamicStack := NewDynamicStack()

	operations := []int{10, 20, 30, 40}

	for _, val := range operations {
		llStack.Push(val)
		arrayStack.Push(val)
		dynamicStack.Push(val)
	}

	llPeek, _ := llStack.Peek()
	arrayPeek, _ := arrayStack.Peek()
	dynamicPeek, _ := dynamicStack.Peek()

	llPopped, _ := llStack.Pop()
	arrayPopped, _ := arrayStack.Pop()
	dynamicPopped, _ := dynamicStack.Pop()

	postfixResult, _ := EvaluatePostfix([]string{"3", "4", "+", "2", "*", "7", "-"})

	balancedTest1 := IsBalancedParentheses("({[]})")
	balancedTest2 := IsBalancedParentheses("({[})")

	return map[string]any{
		"linked_list_stack": map[string]any{
			"original":     []int{10, 20, 30, 40},
			"peek_element": llPeek,
			"popped":       llPopped,
			"after_pop":    llStack.ToSlice(),
			"size":         llStack.Size(),
			"display":      llStack.Display(),
		},
		"array_stack": map[string]any{
			"original":     []int{10, 20, 30, 40},
			"peek_element": arrayPeek,
			"popped":       arrayPopped,
			"after_pop":    arrayStack.ToSlice(),
			"size":         arrayStack.Size(),
			"capacity":     arrayStack.Capacity(),
			"display":      arrayStack.Display(),
		},
		"dynamic_stack": map[string]any{
			"original":     []int{10, 20, 30, 40},
			"peek_element": dynamicPeek,
			"popped":       dynamicPopped,
			"after_pop":    dynamicStack.ToSlice(),
			"size":         dynamicStack.Size(),
			"display":      dynamicStack.Display(),
		},
		"postfix_evaluation": map[string]any{
			"expression": "3 4 + 2 * 7 -",
			"result":     postfixResult,
		},
		"balanced_parentheses": map[string]any{
			"test1": map[string]any{
				"expression": "({[]})",
				"balanced":   balancedTest1,
			},
			"test2": map[string]any{
				"expression": "({[})",
				"balanced":   balancedTest2,
			},
		},
		"data_structure":  "Stack (LIFO)",
		"implementations": 3,
		"principle":       "Last In, First Out",
	}
}
