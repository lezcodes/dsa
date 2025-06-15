package ring_buffers

import (
	"reflect"
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](5)
	if rb.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %d", rb.Capacity())
	}
	if rb.Size() != 0 {
		t.Errorf("Expected size 0, got %d", rb.Size())
	}
	if !rb.IsEmpty() {
		t.Error("Expected empty buffer")
	}
	if rb.IsFull() {
		t.Error("Expected non-full buffer")
	}
}

func TestNewRingBufferInvalidCapacity(t *testing.T) {
	rb := NewRingBuffer[int](0)
	if rb.Capacity() != 1 {
		t.Errorf("Expected capacity 1 for invalid input, got %d", rb.Capacity())
	}

	rb2 := NewRingBuffer[int](-5)
	if rb2.Capacity() != 1 {
		t.Errorf("Expected capacity 1 for negative input, got %d", rb2.Capacity())
	}
}

func TestEnqueueDequeue(t *testing.T) {
	rb := NewRingBuffer[int](3)

	err := rb.Enqueue(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if rb.Size() != 1 {
		t.Errorf("Expected size 1, got %d", rb.Size())
	}

	err = rb.Enqueue(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = rb.Enqueue(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !rb.IsFull() {
		t.Error("Expected full buffer")
	}

	item, err := rb.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if item != 1 {
		t.Errorf("Expected 1, got %d", item)
	}
	if rb.Size() != 2 {
		t.Errorf("Expected size 2, got %d", rb.Size())
	}
}

func TestCircularWrapping(t *testing.T) {
	rb := NewRingBuffer[int](3)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)

	rb.Dequeue()
	rb.Dequeue()

	rb.Enqueue(4)
	rb.Enqueue(5)

	expected := []int{3, 4, 5}
	actual := rb.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestOverflowError(t *testing.T) {
	rb := NewRingBuffer[int](2)

	rb.Enqueue(1)
	rb.Enqueue(2)

	err := rb.Enqueue(3)
	if err == nil {
		t.Error("Expected overflow error")
	}
}

func TestUnderflowError(t *testing.T) {
	rb := NewRingBuffer[int](3)

	_, err := rb.Dequeue()
	if err == nil {
		t.Error("Expected underflow error")
	}

	_, err = rb.Peek()
	if err == nil {
		t.Error("Expected peek error on empty buffer")
	}
}

func TestPeek(t *testing.T) {
	rb := NewRingBuffer[int](3)

	rb.Enqueue(42)
	rb.Enqueue(43)

	item, err := rb.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if item != 42 {
		t.Errorf("Expected 42, got %d", item)
	}

	if rb.Size() != 2 {
		t.Errorf("Expected size unchanged after peek, got %d", rb.Size())
	}
}

func TestClear(t *testing.T) {
	rb := NewRingBuffer[int](3)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)

	rb.Clear()

	if !rb.IsEmpty() {
		t.Error("Expected empty buffer after clear")
	}
	if rb.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", rb.Size())
	}
	if rb.Capacity() != 3 {
		t.Errorf("Expected capacity unchanged after clear, got %d", rb.Capacity())
	}
}

func TestToSliceEmpty(t *testing.T) {
	rb := NewRingBuffer[int](3)

	slice := rb.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %v", slice)
	}
}

func TestToSliceWrapped(t *testing.T) {
	rb := NewRingBuffer[int](4)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)
	rb.Enqueue(4)

	rb.Dequeue()
	rb.Dequeue()

	rb.Enqueue(5)
	rb.Enqueue(6)

	expected := []int{3, 4, 5, 6}
	actual := rb.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestStringType(t *testing.T) {
	rb := NewRingBuffer[string](2)

	rb.Enqueue("hello")
	rb.Enqueue("world")

	item, _ := rb.Dequeue()
	if item != "hello" {
		t.Errorf("Expected 'hello', got '%s'", item)
	}

	rb.Enqueue("!")

	expected := []string{"world", "!"}
	actual := rb.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	data, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected map result")
	}

	if data["first_dequeued"] != 1 {
		t.Errorf("Expected first_dequeued to be 1, got %v", data["first_dequeued"])
	}
}

func BenchmarkEnqueue(b *testing.B) {
	rb := NewRingBuffer[int](1000)

	for i := 0; b.Loop(); i++ {
		if rb.IsFull() {
			rb.Dequeue()
		}
		rb.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	rb := NewRingBuffer[int](1000)
	for i := range 1000 {
		rb.Enqueue(i)
	}

	for i := 0; b.Loop(); i++ {
		if rb.IsEmpty() {
			rb.Enqueue(i)
		}
		rb.Dequeue()
	}
}

func BenchmarkMixedOperations(b *testing.B) {
	rb := NewRingBuffer[int](100)

	for i := 0; b.Loop(); i++ {
		if i%2 == 0 {
			if !rb.IsFull() {
				rb.Enqueue(i)
			}
		} else {
			if !rb.IsEmpty() {
				rb.Dequeue()
			}
		}
	}
}
