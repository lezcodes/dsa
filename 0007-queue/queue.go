package queue

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

func (q *Queue) Front() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func Run() any {
	queue := NewQueue()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	front, _ := queue.Front()
	dequeue, _ := queue.Dequeue()

	return map[string]any{
		"size":    queue.Size(),
		"front":   front,
		"dequeue": dequeue,
		"empty":   queue.IsEmpty(),
	}
}
