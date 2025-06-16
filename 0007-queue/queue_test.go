package queue

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("Expected non-nil Queue")
	}
	if q.Size() != 0 {
		t.Error("Expected empty queue size to be 0")
	}
	if !q.IsEmpty() {
		t.Error("Expected empty queue to be empty")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	val, err := q.Dequeue()
	if err != nil || val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, err = q.Front()
	if err != nil || val != 20 {
		t.Errorf("Expected front to be 20, got %d", val)
	}
}

func TestEmptyQueue(t *testing.T) {
	q := NewQueue()

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	_, err = q.Front()
	if err == nil {
		t.Error("Expected error when getting front of empty queue")
	}
}
