package ring_buffers

import "errors"

type RingBuffer[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
	cap    int
}

func NewRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity <= 0 {
		capacity = 1
	}
	return &RingBuffer[T]{
		buffer: make([]T, capacity),
		head:   0,
		tail:   0,
		size:   0,
		cap:    capacity,
	}
}

func (rb *RingBuffer[T]) Enqueue(item T) error {
	if rb.IsFull() {
		return errors.New("ring buffer is full")
	}

	rb.buffer[rb.tail] = item
	rb.tail = (rb.tail + 1) % rb.cap
	rb.size++
	return nil
}

func (rb *RingBuffer[T]) Dequeue() (T, error) {
	var zero T
	if rb.IsEmpty() {
		return zero, errors.New("ring buffer is empty")
	}

	item := rb.buffer[rb.head]
	rb.buffer[rb.head] = zero
	rb.head = (rb.head + 1) % rb.cap
	rb.size--
	return item, nil
}

func (rb *RingBuffer[T]) Peek() (T, error) {
	var zero T
	if rb.IsEmpty() {
		return zero, errors.New("ring buffer is empty")
	}
	return rb.buffer[rb.head], nil
}

func (rb *RingBuffer[T]) IsEmpty() bool {
	return rb.size == 0
}

func (rb *RingBuffer[T]) IsFull() bool {
	return rb.size == rb.cap
}

func (rb *RingBuffer[T]) Size() int {
	return rb.size
}

func (rb *RingBuffer[T]) Capacity() int {
	return rb.cap
}

func (rb *RingBuffer[T]) Clear() {
	var zero T
	for i := range rb.cap {
		rb.buffer[i] = zero
	}
	rb.head = 0
	rb.tail = 0
	rb.size = 0
}

func (rb *RingBuffer[T]) ToSlice() []T {
	if rb.IsEmpty() {
		return []T{}
	}

	result := make([]T, rb.size)
	for i := range rb.size {
		result[i] = rb.buffer[(rb.head+i)%rb.cap]
	}
	return result
}

func Run() any {
	rb := NewRingBuffer[int](3)

	rb.Enqueue(1)
	rb.Enqueue(2)
	rb.Enqueue(3)

	first, _ := rb.Dequeue()
	rb.Enqueue(4)

	return map[string]any{
		"first_dequeued": first,
		"remaining":      rb.ToSlice(),
		"is_full":        rb.IsFull(),
		"size":           rb.Size(),
	}
}
