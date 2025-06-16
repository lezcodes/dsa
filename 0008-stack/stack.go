package stack

import "fmt"

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func Run() any {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	peek, _ := stack.Peek()
	pop, _ := stack.Pop()

	return map[string]any{
		"size":  stack.Size(),
		"peek":  peek,
		"pop":   pop,
		"empty": stack.IsEmpty(),
	}
}
