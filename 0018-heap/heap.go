package heap

import (
	"container/heap"
	"fmt"
)

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]int, 0)}
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := h.parent(index)
		if h.data[index] >= h.data[parentIndex] {
			break
		}
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *MinHeap) ExtractMin() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	min := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	if len(h.data) > 0 {
		h.heapifyDown(0)
	}

	return min, true
}

func (h *MinHeap) heapifyDown(index int) {
	for {
		smallest := index
		left := h.leftChild(index)
		right := h.rightChild(index)

		if left < len(h.data) && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < len(h.data) && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.swap(index, smallest)
		index = smallest
	}
}

func (h *MinHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *MinHeap) Size() int {
	return len(h.data)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeap) ToSlice() []int {
	result := make([]int, len(h.data))
	copy(result, h.data)
	return result
}

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: make([]int, 0)}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Insert(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := h.parent(index)
		if h.data[index] <= h.data[parentIndex] {
			break
		}
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *MaxHeap) ExtractMax() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}

	max := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	if len(h.data) > 0 {
		h.heapifyDown(0)
	}

	return max, true
}

func (h *MaxHeap) heapifyDown(index int) {
	for {
		largest := index
		left := h.leftChild(index)
		right := h.rightChild(index)

		if left < len(h.data) && h.data[left] > h.data[largest] {
			largest = left
		}

		if right < len(h.data) && h.data[right] > h.data[largest] {
			largest = right
		}

		if largest == index {
			break
		}

		h.swap(index, largest)
		index = largest
	}
}

func (h *MaxHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

func (h *MaxHeap) Size() int {
	return len(h.data)
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MaxHeap) ToSlice() []int {
	result := make([]int, len(h.data))
	copy(result, h.data)
	return result
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSort(arr []int) []int {
	maxHeap := NewMaxHeap()
	for _, v := range arr {
		maxHeap.Insert(v)
	}

	result := make([]int, 0, len(arr))
	for !maxHeap.IsEmpty() {
		val, _ := maxHeap.ExtractMax()
		result = append(result, val)
	}

	return result
}

func Run() any {
	minHeap := NewMinHeap()
	maxHeap := NewMaxHeap()

	values := []int{15, 10, 20, 8, 25, 5, 12}

	for _, v := range values {
		minHeap.Insert(v)
		maxHeap.Insert(v)
	}

	goHeap := &IntHeap{}
	heap.Init(goHeap)
	for _, v := range values {
		heap.Push(goHeap, v)
	}

	minExtracted := []int{}
	for !minHeap.IsEmpty() {
		val, _ := minHeap.ExtractMin()
		minExtracted = append(minExtracted, val)
	}

	maxExtracted := []int{}
	for !maxHeap.IsEmpty() {
		val, _ := maxHeap.ExtractMax()
		maxExtracted = append(maxExtracted, val)
	}

	goExtracted := []int{}
	for goHeap.Len() > 0 {
		val := heap.Pop(goHeap).(int)
		goExtracted = append(goExtracted, val)
	}

	sorted := HeapSort(values)

	return map[string]any{
		"original_values":  values,
		"min_heap_extract": minExtracted,
		"max_heap_extract": maxExtracted,
		"go_heap_extract":  goExtracted,
		"heap_sort_desc":   sorted,
		"comparison": fmt.Sprintf("Raw min-heap vs Go heap: %v vs %v",
			minExtracted, goExtracted),
	}
}
