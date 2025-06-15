package array_list

import (
	"errors"
	"fmt"
)

const (
	defaultCapacity = 10
	growthFactor    = 2
	shrinkThreshold = 0.25
)

type ArrayList[T comparable] struct {
	data     []T
	size     int
	capacity int
}

func NewArrayList[T comparable]() *ArrayList[T] {
	return &ArrayList[T]{
		data:     make([]T, defaultCapacity),
		size:     0,
		capacity: defaultCapacity,
	}
}

func NewArrayListWithCapacity[T comparable](capacity int) *ArrayList[T] {
	if capacity < 1 {
		capacity = defaultCapacity
	}
	return &ArrayList[T]{
		data:     make([]T, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (al *ArrayList[T]) Add(item T) {
	al.ensureCapacity()
	al.data[al.size] = item
	al.size++
}

func (al *ArrayList[T]) Insert(index int, item T) error {
	if index < 0 || index > al.size {
		return errors.New("index out of bounds")
	}

	al.ensureCapacity()

	for i := al.size; i > index; i-- {
		al.data[i] = al.data[i-1]
	}

	al.data[index] = item
	al.size++
	return nil
}

func (al *ArrayList[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index >= al.size {
		return zero, errors.New("index out of bounds")
	}

	item := al.data[index]

	for i := index; i < al.size-1; i++ {
		al.data[i] = al.data[i+1]
	}

	al.size--
	al.data[al.size] = zero
	al.shrinkIfNeeded()

	return item, nil
}

func (al *ArrayList[T]) RemoveItem(item T) bool {
	index := al.IndexOf(item)
	if index == -1 {
		return false
	}
	al.Remove(index)
	return true
}

func (al *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= al.size {
		return zero, errors.New("index out of bounds")
	}
	return al.data[index], nil
}

func (al *ArrayList[T]) Set(index int, item T) error {
	if index < 0 || index >= al.size {
		return errors.New("index out of bounds")
	}
	al.data[index] = item
	return nil
}

func (al *ArrayList[T]) IndexOf(item T) int {
	for i := range al.size {
		if al.data[i] == item {
			return i
		}
	}
	return -1
}

func (al *ArrayList[T]) Contains(item T) bool {
	return al.IndexOf(item) != -1
}

func (al *ArrayList[T]) Size() int {
	return al.size
}

func (al *ArrayList[T]) Capacity() int {
	return al.capacity
}

func (al *ArrayList[T]) IsEmpty() bool {
	return al.size == 0
}

func (al *ArrayList[T]) Clear() {
	var zero T
	for i := range al.size {
		al.data[i] = zero
	}
	al.size = 0
}

func (al *ArrayList[T]) ToSlice() []T {
	result := make([]T, al.size)
	copy(result, al.data[:al.size])
	return result
}

func (al *ArrayList[T]) String() string {
	return fmt.Sprintf("ArrayList{size: %d, capacity: %d, data: %v}",
		al.size, al.capacity, al.data[:al.size])
}

func (al *ArrayList[T]) ensureCapacity() {
	if al.size >= al.capacity {
		al.resize(al.capacity * growthFactor)
	}
}

func (al *ArrayList[T]) shrinkIfNeeded() {
	if al.capacity > defaultCapacity && float64(al.size) <= float64(al.capacity)*shrinkThreshold {
		newCapacity := al.capacity / growthFactor
		if newCapacity < defaultCapacity {
			newCapacity = defaultCapacity
		}
		al.resize(newCapacity)
	}
}

func (al *ArrayList[T]) resize(newCapacity int) {
	newData := make([]T, newCapacity)
	copy(newData, al.data[:al.size])
	al.data = newData
	al.capacity = newCapacity
}

func (al *ArrayList[T]) First() (T, error) {
	var zero T
	if al.IsEmpty() {
		return zero, errors.New("list is empty")
	}
	return al.data[0], nil
}

func (al *ArrayList[T]) Last() (T, error) {
	var zero T
	if al.IsEmpty() {
		return zero, errors.New("list is empty")
	}
	return al.data[al.size-1], nil
}

func (al *ArrayList[T]) Prepend(item T) {
	al.Insert(0, item)
}

func (al *ArrayList[T]) Pop() (T, error) {
	var zero T
	if al.IsEmpty() {
		return zero, errors.New("list is empty")
	}
	return al.Remove(al.size - 1)
}

func Run() any {
	al := NewArrayList[string]()

	al.Add("apple")
	al.Add("banana")
	al.Add("cherry")

	al.Insert(1, "blueberry")

	removed, _ := al.Remove(2)

	al.Set(0, "apricot")

	return map[string]any{
		"final_list":         al.ToSlice(),
		"size":               al.Size(),
		"capacity":           al.Capacity(),
		"removed":            removed,
		"contains_cherry":    al.Contains("cherry"),
		"index_of_blueberry": al.IndexOf("blueberry"),
	}
}
