package heap

import (
	"container/heap"
	"reflect"
	"sort"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()

	if !h.IsEmpty() {
		t.Error("New heap should be empty")
	}

	if h.Size() != 0 {
		t.Error("New heap size should be 0")
	}

	values := []int{15, 10, 20, 8, 25, 5, 12}
	for _, v := range values {
		h.Insert(v)
	}

	if h.Size() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), h.Size())
	}

	min, ok := h.Peek()
	if !ok || min != 5 {
		t.Errorf("Expected min 5, got %d", min)
	}

	extracted := []int{}
	for !h.IsEmpty() {
		val, ok := h.ExtractMin()
		if !ok {
			t.Error("ExtractMin should return true for non-empty heap")
		}
		extracted = append(extracted, val)
	}

	expected := []int{5, 8, 10, 12, 15, 20, 25}
	if !reflect.DeepEqual(extracted, expected) {
		t.Errorf("Expected %v, got %v", expected, extracted)
	}
}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	values := []int{15, 10, 20, 8, 25, 5, 12}
	for _, v := range values {
		h.Insert(v)
	}

	max, ok := h.Peek()
	if !ok || max != 25 {
		t.Errorf("Expected max 25, got %d", max)
	}

	extracted := []int{}
	for !h.IsEmpty() {
		val, ok := h.ExtractMax()
		if !ok {
			t.Error("ExtractMax should return true for non-empty heap")
		}
		extracted = append(extracted, val)
	}

	expected := []int{25, 20, 15, 12, 10, 8, 5}
	if !reflect.DeepEqual(extracted, expected) {
		t.Errorf("Expected %v, got %v", expected, extracted)
	}
}

func TestMinHeapEmptyOperations(t *testing.T) {
	h := NewMinHeap()

	_, ok := h.Peek()
	if ok {
		t.Error("Peek on empty heap should return false")
	}

	_, ok = h.ExtractMin()
	if ok {
		t.Error("ExtractMin on empty heap should return false")
	}
}

func TestMaxHeapEmptyOperations(t *testing.T) {
	h := NewMaxHeap()

	_, ok := h.Peek()
	if ok {
		t.Error("Peek on empty heap should return false")
	}

	_, ok = h.ExtractMax()
	if ok {
		t.Error("ExtractMax on empty heap should return false")
	}
}

func TestHeapWithDuplicates(t *testing.T) {
	minHeap := NewMinHeap()
	values := []int{5, 3, 5, 1, 3, 1}

	for _, v := range values {
		minHeap.Insert(v)
	}

	extracted := []int{}
	for !minHeap.IsEmpty() {
		val, _ := minHeap.ExtractMin()
		extracted = append(extracted, val)
	}

	expected := []int{1, 1, 3, 3, 5, 5}
	if !reflect.DeepEqual(extracted, expected) {
		t.Errorf("Expected %v, got %v", expected, extracted)
	}
}

func TestHeapSort(t *testing.T) {
	values := []int{15, 10, 20, 8, 25, 5, 12}
	sorted := HeapSort(values)
	expected := []int{25, 20, 15, 12, 10, 8, 5}

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

func TestGoHeapComparison(t *testing.T) {
	values := []int{15, 10, 20, 8, 25, 5, 12}

	minHeap := NewMinHeap()
	for _, v := range values {
		minHeap.Insert(v)
	}

	goHeap := &IntHeap{}
	heap.Init(goHeap)
	for _, v := range values {
		heap.Push(goHeap, v)
	}

	rawExtracted := []int{}
	for !minHeap.IsEmpty() {
		val, _ := minHeap.ExtractMin()
		rawExtracted = append(rawExtracted, val)
	}

	goExtracted := []int{}
	for goHeap.Len() > 0 {
		val := heap.Pop(goHeap).(int)
		goExtracted = append(goExtracted, val)
	}

	if !reflect.DeepEqual(rawExtracted, goExtracted) {
		t.Errorf("Raw heap %v != Go heap %v", rawExtracted, goExtracted)
	}
}

func TestHeapProperty(t *testing.T) {
	minHeap := NewMinHeap()
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}

	for _, v := range values {
		minHeap.Insert(v)
	}

	data := minHeap.ToSlice()

	for i := range data {
		left := 2*i + 1
		right := 2*i + 2

		if left < len(data) && data[i] > data[left] {
			t.Errorf("Min heap property violated: parent %d > left child %d", data[i], data[left])
		}

		if right < len(data) && data[i] > data[right] {
			t.Errorf("Min heap property violated: parent %d > right child %d", data[i], data[right])
		}
	}
}

func TestSingleElement(t *testing.T) {
	minHeap := NewMinHeap()
	minHeap.Insert(42)

	if minHeap.Size() != 1 {
		t.Error("Size should be 1")
	}

	val, ok := minHeap.Peek()
	if !ok || val != 42 {
		t.Error("Should peek 42")
	}

	val, ok = minHeap.ExtractMin()
	if !ok || val != 42 {
		t.Error("Should extract 42")
	}

	if !minHeap.IsEmpty() {
		t.Error("Should be empty after extraction")
	}
}

func TestRun(t *testing.T) {
	result := Run()
	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", result)
	}

	minExtracted, ok := resultMap["min_heap_extract"].([]int)
	if !ok {
		t.Fatal("min_heap_extract should be []int")
	}

	expected := []int{5, 8, 10, 12, 15, 20, 25}
	if !reflect.DeepEqual(minExtracted, expected) {
		t.Errorf("Expected min heap extract %v, got %v", expected, minExtracted)
	}

	maxExtracted, ok := resultMap["max_heap_extract"].([]int)
	if !ok {
		t.Fatal("max_heap_extract should be []int")
	}

	expectedMax := []int{25, 20, 15, 12, 10, 8, 5}
	if !reflect.DeepEqual(maxExtracted, expectedMax) {
		t.Errorf("Expected max heap extract %v, got %v", expectedMax, maxExtracted)
	}

	goExtracted, ok := resultMap["go_heap_extract"].([]int)
	if !ok {
		t.Fatal("go_heap_extract should be []int")
	}

	if !reflect.DeepEqual(minExtracted, goExtracted) {
		t.Errorf("Raw min heap %v != Go heap %v", minExtracted, goExtracted)
	}
}

func BenchmarkMinHeapInsert(b *testing.B) {
	h := NewMinHeap()
	for i := 0; b.Loop(); i++ {
		h.Insert(i)
	}
}

func BenchmarkMinHeapExtract(b *testing.B) {
	h := NewMinHeap()
	for i := 0; b.Loop(); i++ {
		h.Insert(i)
	}

	for b.Loop() {
		h.ExtractMin()
	}
}

func BenchmarkGoHeapPush(b *testing.B) {
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; b.Loop(); i++ {
		heap.Push(h, i)
	}
}

func BenchmarkGoHeapPop(b *testing.B) {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; b.Loop(); i++ {
		heap.Push(h, i)
	}

	for b.Loop() {
		heap.Pop(h)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	values := []int{15, 10, 20, 8, 25, 5, 12, 30, 18, 22}

	for b.Loop() {
		HeapSort(values)
	}
}

func BenchmarkStandardSort(b *testing.B) {
	values := []int{15, 10, 20, 8, 25, 5, 12, 30, 18, 22}

	for b.Loop() {
		sorted := make([]int, len(values))
		copy(sorted, values)
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	}
}
