package queue

import (
	"reflect"
	"testing"
)

func TestLinkedListQueue(t *testing.T) {
	q := NewLinkedListQueue()

	if !q.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	_, err = q.Front()
	if err == nil {
		t.Error("Expected error when getting front of empty queue")
	}

	_, err = q.Rear()
	if err == nil {
		t.Error("Expected error when getting rear of empty queue")
	}

	values := []int{10, 20, 30, 40}
	for _, val := range values {
		q.Enqueue(val)
	}

	if q.Size() != 4 {
		t.Errorf("Expected size to be 4, got %d", q.Size())
	}

	if q.IsEmpty() {
		t.Error("Expected queue not to be empty")
	}

	front, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 10 {
		t.Errorf("Expected front to be 10, got %d", front)
	}

	rear, err := q.Rear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if rear != 40 {
		t.Errorf("Expected rear to be 40, got %d", rear)
	}

	expected := []int{10, 20, 30, 40}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, q.ToSlice())
	}

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if dequeued != 10 {
		t.Errorf("Expected dequeued to be 10, got %d", dequeued)
	}

	expected = []int{20, 30, 40}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v after dequeue, got %v", expected, q.ToSlice())
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after clear")
	}
}

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(3)

	if !q.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}

	if q.IsFull() {
		t.Error("Expected new queue not to be full")
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	if q.Capacity() != 3 {
		t.Errorf("Expected capacity to be 3, got %d", q.Capacity())
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	values := []int{10, 20, 30}
	for _, val := range values {
		err := q.Enqueue(val)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	if !q.IsFull() {
		t.Error("Expected queue to be full")
	}

	err = q.Enqueue(40)
	if err == nil {
		t.Error("Expected error when enqueueing to full queue")
	}

	front, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 10 {
		t.Errorf("Expected front to be 10, got %d", front)
	}

	rear, err := q.Rear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if rear != 30 {
		t.Errorf("Expected rear to be 30, got %d", rear)
	}

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, q.ToSlice())
	}

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if dequeued != 10 {
		t.Errorf("Expected dequeued to be 10, got %d", dequeued)
	}

	err = q.Enqueue(40)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected = []int{20, 30, 40}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v after circular operation, got %v", expected, q.ToSlice())
	}
}

func TestArrayQueueCircular(t *testing.T) {
	q := NewArrayQueue(3)

	for i := 1; i <= 3; i++ {
		q.Enqueue(i * 10)
	}

	for i := 0; i < 5; i++ {
		dequeued, _ := q.Dequeue()
		q.Enqueue(dequeued + 100)
	}

	expected := []int{130, 210, 220}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v after circular operations, got %v", expected, q.ToSlice())
	}
}

func TestDynamicQueue(t *testing.T) {
	q := NewDynamicQueue()

	if !q.IsEmpty() {
		t.Error("Expected new queue to be empty")
	}

	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeuing from empty queue")
	}

	values := []int{10, 20, 30, 40, 50}
	for _, val := range values {
		q.Enqueue(val)
	}

	if q.Size() != 5 {
		t.Errorf("Expected size to be 5, got %d", q.Size())
	}

	front, err := q.Front()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if front != 10 {
		t.Errorf("Expected front to be 10, got %d", front)
	}

	rear, err := q.Rear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if rear != 50 {
		t.Errorf("Expected rear to be 50, got %d", rear)
	}

	expected := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, q.ToSlice())
	}

	dequeued, err := q.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if dequeued != 10 {
		t.Errorf("Expected dequeued to be 10, got %d", dequeued)
	}

	expected = []int{20, 30, 40, 50}
	if !reflect.DeepEqual(q.ToSlice(), expected) {
		t.Errorf("Expected %v after dequeue, got %v", expected, q.ToSlice())
	}

	q.Clear()
	if !q.IsEmpty() {
		t.Error("Expected queue to be empty after clear")
	}
}

func TestQueueFIFOProperty(t *testing.T) {
	queues := []struct {
		name  string
		queue interface {
			Enqueue(int) error
			Dequeue() (int, error)
			IsEmpty() bool
		}
	}{
		{"LinkedList", &LinkedListQueueAdapter{NewLinkedListQueue()}},
		{"Array", NewArrayQueue(10)},
		{"Dynamic", &DynamicQueueAdapter{NewDynamicQueue()}},
	}

	for _, test := range queues {
		t.Run(test.name, func(t *testing.T) {
			q := test.queue

			input := []int{1, 2, 3, 4, 5}
			for _, val := range input {
				q.Enqueue(val)
			}

			var output []int
			for !q.IsEmpty() {
				val, _ := q.Dequeue()
				output = append(output, val)
			}

			if !reflect.DeepEqual(input, output) {
				t.Errorf("FIFO property violated: input %v, output %v", input, output)
			}
		})
	}
}

type LinkedListQueueAdapter struct {
	*LinkedListQueue
}

func (a *LinkedListQueueAdapter) Enqueue(data int) error {
	a.LinkedListQueue.Enqueue(data)
	return nil
}

type DynamicQueueAdapter struct {
	*DynamicQueue
}

func (a *DynamicQueueAdapter) Enqueue(data int) error {
	a.DynamicQueue.Enqueue(data)
	return nil
}

func TestQueueDisplay(t *testing.T) {
	q := NewLinkedListQueue()

	display := q.Display()
	if display != "Queue: []" {
		t.Errorf("Expected 'Queue: []', got '%s'", display)
	}

	q.Enqueue(10)
	q.Enqueue(20)

	display = q.Display()
	expected := "Queue: [10 <- 20] (front <- rear)"
	if display != expected {
		t.Errorf("Expected '%s', got '%s'", expected, display)
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

	requiredKeys := []string{"linked_list_queue", "array_queue", "dynamic_queue", "data_structure"}
	for _, key := range requiredKeys {
		if _, exists := resultMap[key]; !exists {
			t.Errorf("Expected result to contain '%s' key", key)
		}
	}
}

func BenchmarkLinkedListQueueEnqueue(b *testing.B) {
	q := NewLinkedListQueue()

	for b.Loop() {
		q.Enqueue(42)
	}
}

func BenchmarkLinkedListQueueDequeue(b *testing.B) {
	q := NewLinkedListQueue()
	for i := range b.N {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for b.Loop() {
		q.Dequeue()
	}
}

func BenchmarkArrayQueueEnqueue(b *testing.B) {
	q := NewArrayQueue(b.N)

	for b.Loop() {
		q.Enqueue(42)
	}
}

func BenchmarkArrayQueueDequeue(b *testing.B) {
	q := NewArrayQueue(b.N)
	for i := range b.N {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for b.Loop() {
		q.Dequeue()
	}
}

func BenchmarkDynamicQueueEnqueue(b *testing.B) {
	q := NewDynamicQueue()

	for b.Loop() {
		q.Enqueue(42)
	}
}

func BenchmarkDynamicQueueDequeue(b *testing.B) {
	q := NewDynamicQueue()
	for i := range b.N {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for b.Loop() {
		q.Dequeue()
	}
}

func BenchmarkQueueOperations(b *testing.B) {
	benchmarks := []struct {
		name string
		fn   func(*testing.B)
	}{
		{"LinkedListMixed", func(b *testing.B) {
			q := NewLinkedListQueue()
			for b.Loop() {
				q.Enqueue(42)
				if !q.IsEmpty() {
					q.Dequeue()
				}
			}
		}},
		{"ArrayMixed", func(b *testing.B) {
			q := NewArrayQueue(1000)
			for b.Loop() {
				q.Enqueue(42)
				if !q.IsEmpty() {
					q.Dequeue()
				}
			}
		}},
		{"DynamicMixed", func(b *testing.B) {
			q := NewDynamicQueue()
			for b.Loop() {
				q.Enqueue(42)
				if !q.IsEmpty() {
					q.Dequeue()
				}
			}
		}},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, bm.fn)
	}
}
