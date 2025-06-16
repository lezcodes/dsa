package stack

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("Expected non-nil Stack")
	}
	if s.Size() != 0 {
		t.Error("Expected empty stack size to be 0")
	}
	if !s.IsEmpty() {
		t.Error("Expected empty stack to be empty")
	}
}

func TestPushPop(t *testing.T) {
	s := NewStack()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	val, err := s.Pop()
	if err != nil || val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	val, err = s.Peek()
	if err != nil || val != 20 {
		t.Errorf("Expected peek to be 20, got %d", val)
	}
}

func TestEmptyStack(t *testing.T) {
	s := NewStack()

	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}

	_, err = s.Peek()
	if err == nil {
		t.Error("Expected error when peeking empty stack")
	}
}
