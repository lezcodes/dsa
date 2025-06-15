package queue

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *LinkedListQueue) Enqueue(data int) {
	newNode := &Node{Data: data, Next: nil}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}

	q.size++
}

func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}

	data := q.front.Data
	q.front = q.front.Next

	if q.front == nil {
		q.rear = nil
	}

	q.size--
	return data, nil
}

func (q *LinkedListQueue) Front() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.front.Data, nil
}

func (q *LinkedListQueue) Rear() (int, error) {
	if q.rear == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.rear.Data, nil
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Clear() {
	q.front = nil
	q.rear = nil
	q.size = 0
}

func (q *LinkedListQueue) ToSlice() []int {
	result := make([]int, 0, q.size)
	current := q.front

	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}

	return result
}

func (q *LinkedListQueue) Display() string {
	if q.front == nil {
		return "Queue: []"
	}

	result := "Queue: ["
	current := q.front

	for current != nil {
		result += fmt.Sprintf("%d", current.Data)
		if current.Next != nil {
			result += " <- "
		}
		current = current.Next
	}

	result += "] (front <- rear)"
	return result
}

type ArrayQueue struct {
	data     []int
	front    int
	rear     int
	size     int
	capacity int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	if capacity <= 0 {
		capacity = 10
	}

	return &ArrayQueue{
		data:     make([]int, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

func (q *ArrayQueue) Enqueue(data int) error {
	if q.size >= q.capacity {
		return fmt.Errorf("queue is full")
	}

	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = data
	q.size++

	return nil
}

func (q *ArrayQueue) Dequeue() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	data := q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--

	return data, nil
}

func (q *ArrayQueue) Front() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[q.front], nil
}

func (q *ArrayQueue) Rear() (int, error) {
	if q.size == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[q.rear], nil
}

func (q *ArrayQueue) Size() int {
	return q.size
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue) IsFull() bool {
	return q.size == q.capacity
}

func (q *ArrayQueue) Clear() {
	q.front = 0
	q.rear = -1
	q.size = 0
}

func (q *ArrayQueue) ToSlice() []int {
	result := make([]int, 0, q.size)

	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result = append(result, q.data[index])
	}

	return result
}

func (q *ArrayQueue) Display() string {
	if q.size == 0 {
		return "Queue: []"
	}

	result := "Queue: ["

	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.capacity
		result += fmt.Sprintf("%d", q.data[index])
		if i < q.size-1 {
			result += " <- "
		}
	}

	result += "] (front <- rear)"
	return result
}

func (q *ArrayQueue) Capacity() int {
	return q.capacity
}

type DynamicQueue struct {
	data []int
}

func NewDynamicQueue() *DynamicQueue {
	return &DynamicQueue{
		data: make([]int, 0),
	}
}

func (q *DynamicQueue) Enqueue(data int) {
	q.data = append(q.data, data)
}

func (q *DynamicQueue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	data := q.data[0]
	q.data = q.data[1:]

	return data, nil
}

func (q *DynamicQueue) Front() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[0], nil
}

func (q *DynamicQueue) Rear() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[len(q.data)-1], nil
}

func (q *DynamicQueue) Size() int {
	return len(q.data)
}

func (q *DynamicQueue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *DynamicQueue) Clear() {
	q.data = q.data[:0]
}

func (q *DynamicQueue) ToSlice() []int {
	result := make([]int, len(q.data))
	copy(result, q.data)
	return result
}

func (q *DynamicQueue) Display() string {
	if len(q.data) == 0 {
		return "Queue: []"
	}

	result := "Queue: ["

	for i, val := range q.data {
		result += fmt.Sprintf("%d", val)
		if i < len(q.data)-1 {
			result += " <- "
		}
	}

	result += "] (front <- rear)"
	return result
}

func Run() any {
	llQueue := NewLinkedListQueue()
	arrayQueue := NewArrayQueue(5)
	dynamicQueue := NewDynamicQueue()

	operations := []int{10, 20, 30, 40}

	for _, val := range operations {
		llQueue.Enqueue(val)
		arrayQueue.Enqueue(val)
		dynamicQueue.Enqueue(val)
	}

	llFront, _ := llQueue.Front()
	arrayFront, _ := arrayQueue.Front()
	dynamicFront, _ := dynamicQueue.Front()

	llDequeued, _ := llQueue.Dequeue()
	arrayDequeued, _ := arrayQueue.Dequeue()
	dynamicDequeued, _ := dynamicQueue.Dequeue()

	return map[string]any{
		"linked_list_queue": map[string]any{
			"original":      []int{10, 20, 30, 40},
			"front_element": llFront,
			"dequeued":      llDequeued,
			"after_dequeue": llQueue.ToSlice(),
			"size":          llQueue.Size(),
			"display":       llQueue.Display(),
		},
		"array_queue": map[string]any{
			"original":      []int{10, 20, 30, 40},
			"front_element": arrayFront,
			"dequeued":      arrayDequeued,
			"after_dequeue": arrayQueue.ToSlice(),
			"size":          arrayQueue.Size(),
			"capacity":      arrayQueue.Capacity(),
			"display":       arrayQueue.Display(),
		},
		"dynamic_queue": map[string]any{
			"original":      []int{10, 20, 30, 40},
			"front_element": dynamicFront,
			"dequeued":      dynamicDequeued,
			"after_dequeue": dynamicQueue.ToSlice(),
			"size":          dynamicQueue.Size(),
			"display":       dynamicQueue.Display(),
		},
		"data_structure":  "Queue (FIFO)",
		"implementations": 3,
		"principle":       "First In, First Out",
	}
}
