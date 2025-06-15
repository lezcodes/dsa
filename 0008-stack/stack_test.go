package stack

import (
	"reflect"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	s := NewLinkedListStack()

	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", s.Size())
	}

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}

	_, err = s.Peek()
	if err == nil {
		t.Error("Expected error when peeking empty stack")
	}

	values := []int{10, 20, 30, 40}
	for _, val := range values {
		s.Push(val)
	}

	if s.Size() != 4 {
		t.Errorf("Expected size to be 4, got %d", s.Size())
	}

	if s.IsEmpty() {
		t.Error("Expected stack not to be empty")
	}

	peek, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if peek != 40 {
		t.Errorf("Expected peek to be 40, got %d", peek)
	}

	expected := []int{40, 30, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, s.ToSlice())
	}

	popped, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if popped != 40 {
		t.Errorf("Expected popped to be 40, got %d", popped)
	}

	expected = []int{30, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v after pop, got %v", expected, s.ToSlice())
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty after clear")
	}
}

func TestArrayStack(t *testing.T) {
	s := NewArrayStack(3)

	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	if s.IsFull() {
		t.Error("Expected new stack not to be full")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", s.Size())
	}

	if s.Capacity() != 3 {
		t.Errorf("Expected capacity to be 3, got %d", s.Capacity())
	}

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}

	values := []int{10, 20, 30}
	for _, val := range values {
		err := s.Push(val)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if !s.IsFull() {
		t.Error("Expected stack to be full")
	}

	err = s.Push(40)
	if err == nil {
		t.Error("Expected error when pushing to full stack")
	}

	peek, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if peek != 30 {
		t.Errorf("Expected peek to be 30, got %d", peek)
	}

	expected := []int{30, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, s.ToSlice())
	}

	popped, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if popped != 30 {
		t.Errorf("Expected popped to be 30, got %d", popped)
	}

	err = s.Push(40)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected = []int{40, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v after push, got %v", expected, s.ToSlice())
	}
}

func TestDynamicStack(t *testing.T) {
	s := NewDynamicStack()

	if !s.IsEmpty() {
		t.Error("Expected new stack to be empty")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", s.Size())
	}

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}

	values := []int{10, 20, 30, 40, 50}
	for _, val := range values {
		s.Push(val)
	}

	if s.Size() != 5 {
		t.Errorf("Expected size to be 5, got %d", s.Size())
	}

	peek, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if peek != 50 {
		t.Errorf("Expected peek to be 50, got %d", peek)
	}

	expected := []int{50, 40, 30, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, s.ToSlice())
	}

	popped, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if popped != 50 {
		t.Errorf("Expected popped to be 50, got %d", popped)
	}

	expected = []int{40, 30, 20, 10}
	if !reflect.DeepEqual(s.ToSlice(), expected) {
		t.Errorf("Expected %v after pop, got %v", expected, s.ToSlice())
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty after clear")
	}
}

func TestStackLIFOProperty(t *testing.T) {
	stacks := []struct {
		name  string
		stack interface {
			Push(int) error
			Pop() (int, error)
			IsEmpty() bool
		}
	}{
		{"LinkedList", &LinkedListStackAdapter{NewLinkedListStack()}},
		{"Array", NewArrayStack(10)},
		{"Dynamic", &DynamicStackAdapter{NewDynamicStack()}},
	}

	for _, test := range stacks {
		t.Run(test.name, func(t *testing.T) {
			s := test.stack

			input := []int{1, 2, 3, 4, 5}
			for _, val := range input {
				s.Push(val)
			}

			var output []int
			for !s.IsEmpty() {
				val, _ := s.Pop()
				output = append(output, val)
			}

			expected := []int{5, 4, 3, 2, 1}
			if !reflect.DeepEqual(output, expected) {
				t.Errorf("LIFO property violated: expected %v, got %v", expected, output)
			}
		})
	}
}

type LinkedListStackAdapter struct {
	*LinkedListStack
}

func (a *LinkedListStackAdapter) Push(data int) error {
	a.LinkedListStack.Push(data)
	return nil
}

type DynamicStackAdapter struct {
	*DynamicStack
}

func (a *DynamicStackAdapter) Push(data int) error {
	a.DynamicStack.Push(data)
	return nil
}

func TestStackDisplay(t *testing.T) {
	s := NewLinkedListStack()

	display := s.Display()
	if display != "Stack: []" {
		t.Errorf("Expected 'Stack: []', got '%s'", display)
	}

	s.Push(10)
	s.Push(20)

	display = s.Display()
	expected := "Stack: [20 | 10] (top | bottom)"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
	}
}

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		name       string
		expression []string
		expected   int
		hasError   bool
	}{
		{
			name:       "Simple addition",
			expression: []string{"3", "4", "+"},
			expected:   7,
			hasError:   false,
		},
		{
			name:       "Complex expression",
			expression: []string{"3", "4", "+", "2", "*", "7", "-"},
			expected:   7,
			hasError:   false,
		},
		{
			name:       "Division",
			expression: []string{"15", "3", "/"},
			expected:   5,
			hasError:   false,
		},
		{
			name:       "Multiple operations",
			expression: []string{"2", "3", "+", "4", "5", "+", "*"},
			expected:   45,
			hasError:   false,
		},
		{
			name:       "Division by zero",
			expression: []string{"5", "0", "/"},
			expected:   0,
			hasError:   true,
		},
		{
			name:       "Invalid expression - insufficient operands",
			expression: []string{"3", "+"},
			expected:   0,
			hasError:   true,
		},
		{
			name:       "Invalid expression - too many operands",
			expression: []string{"3", "4", "5", "+"},
			expected:   0,
			hasError:   true,
		},
		{
			name:       "Invalid token",
			expression: []string{"3", "a", "+"},
			expected:   0,
			hasError:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := EvaluatePostfix(test.expression)

			if test.hasError {
				if err == nil {
					t.Errorf("Expected error for expression %v", test.expression)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("Expected %d, got %d for expression %v", test.expected, result, test.expression)
				}
			}
		})
	}
}

func TestIsBalancedParentheses(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   bool
	}{
		{
			name:       "Simple balanced",
			expression: "()",
			expected:   true,
		},
		{
			name:       "Multiple types balanced",
			expression: "({[]})",
			expected:   true,
		},
		{
			name:       "Complex balanced",
			expression: "{[()]}",
			expected:   true,
		},
		{
			name:       "Empty string",
			expression: "",
			expected:   true,
		},
		{
			name:       "Simple unbalanced",
			expression: "(",
			expected:   false,
		},
		{
			name:       "Wrong order",
			expression: "({[})",
			expected:   false,
		},
		{
			name:       "Extra closing",
			expression: "())",
			expected:   false,
		},
		{
			name:       "Mixed with other characters",
			expression: "a(b{c[d]e}f)g",
			expected:   true,
		},
		{
			name:       "Nested unbalanced",
			expression: "(()",
			expected:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsBalancedParentheses(test.expression)
			if result != test.expected {
				t.Errorf("Expected %v for expression '%s', got %v", test.expected, test.expression, result)
			}
		})
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
	}

	requiredKeys := []string{"linked_list_stack", "array_stack", "dynamic_stack", "postfix_evaluation", "balanced_parentheses", "data_structure"}
	for _, key := range requiredKeys {
		if _, exists := resultMap[key]; !exists {
			t.Errorf("Expected result to contain '%s' key", key)
		}
	}
}

func BenchmarkLinkedListStackPush(b *testing.B) {
	s := NewLinkedListStack()

	for b.Loop() {
		s.Push(42)
	}
}

func BenchmarkLinkedListStackPop(b *testing.B) {
	s := NewLinkedListStack()
	for i := range b.N {
		s.Push(i)
	}

	b.ResetTimer()
	for b.Loop() {
		s.Pop()
	}
}

func BenchmarkArrayStackPush(b *testing.B) {
	s := NewArrayStack(b.N)

	for b.Loop() {
		s.Push(42)
	}
}

func BenchmarkArrayStackPop(b *testing.B) {
	s := NewArrayStack(b.N)
	for i := range b.N {
		s.Push(i)
	}

	b.ResetTimer()
	for b.Loop() {
		s.Pop()
	}
}

func BenchmarkDynamicStackPush(b *testing.B) {
	s := NewDynamicStack()

	for b.Loop() {
		s.Push(42)
	}
}

func BenchmarkDynamicStackPop(b *testing.B) {
	s := NewDynamicStack()
	for i := range b.N {
		s.Push(i)
	}

	b.ResetTimer()
	for b.Loop() {
		s.Pop()
	}
}

func BenchmarkStackOperations(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(*testing.B)
	}{
		{"LinkedListMixed", func(b *testing.B) {
			s := NewLinkedListStack()
			for b.Loop() {
				s.Push(42)
				if !s.IsEmpty() {
					s.Pop()
				}
			}
		}},
		{"ArrayMixed", func(b *testing.B) {
			s := NewArrayStack(1000)
			for b.Loop() {
				s.Push(42)
				if !s.IsEmpty() {
					s.Pop()
				}
			}
		}},
		{"DynamicMixed", func(b *testing.B) {
			s := NewDynamicStack()
			for b.Loop() {
				s.Push(42)
				if !s.IsEmpty() {
					s.Pop()
				}
			}
		}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, bm.fn)
	}
}

func BenchmarkPostfixEvaluation(b *testing.B) {
	expression := []string{"3", "4", "+", "2", "*", "7", "-"}

	for b.Loop() {
		EvaluatePostfix(expression)
	}
}

func BenchmarkBalancedParentheses(b *testing.B) {
	expression := "({[()()]})"

	for b.Loop() {
		IsBalancedParentheses(expression)
	}
}
