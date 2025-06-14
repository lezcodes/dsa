package doubly_linked_list

import (
	"reflect"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList()

	if dll.Head != nil {
		t.Error("Expected head to be nil for new list")
	}

	if dll.Tail != nil {
		t.Error("Expected tail to be nil for new list")
	}

	if dll.Length() != 0 {
		t.Errorf("Expected length to be 0, got %d", dll.Length())
	}

	if !dll.IsEmpty() {
		t.Error("Expected new list to be empty")
	}
}

func TestPrepend(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.Prepend(10)
	if dll.Head.Data != 10 {
		t.Errorf("Expected head data to be 10, got %d", dll.Head.Data)
	}

	if dll.Tail.Data != 10 {
		t.Errorf("Expected tail data to be 10, got %d", dll.Tail.Data)
	}

	if dll.Length() != 1 {
		t.Errorf("Expected length to be 1, got %d", dll.Length())
	}

	dll.Prepend(20)
	if dll.Head.Data != 20 {
		t.Errorf("Expected head data to be 20, got %d", dll.Head.Data)
	}

	if dll.Tail.Data != 10 {
		t.Errorf("Expected tail data to be 10, got %d", dll.Tail.Data)
	}

	if dll.Head.Prev != nil {
		t.Error("Expected head prev to be nil")
	}

	if dll.Tail.Next != nil {
		t.Error("Expected tail next to be nil")
	}

	expected := []int{20, 10}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}
}

func TestAppend(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.Append(10)
	if dll.Head.Data != 10 {
		t.Errorf("Expected head data to be 10, got %d", dll.Head.Data)
	}

	if dll.Tail.Data != 10 {
		t.Errorf("Expected tail data to be 10, got %d", dll.Tail.Data)
	}

	dll.Append(20)
	dll.Append(30)

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}

	if dll.Head.Data != 10 {
		t.Errorf("Expected head data to be 10, got %d", dll.Head.Data)
	}

	if dll.Tail.Data != 30 {
		t.Errorf("Expected tail data to be 30, got %d", dll.Tail.Data)
	}

	if dll.Length() != 3 {
		t.Errorf("Expected length to be 3, got %d", dll.Length())
	}
}

func TestInsertAt(t *testing.T) {
	dll := NewDoublyLinkedList()

	err := dll.InsertAt(10, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = dll.InsertAt(30, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = dll.InsertAt(20, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}

	err = dll.InsertAt(-1, -1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	err = dll.InsertAt(5, 10)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestGet(t *testing.T) {
	dll := NewDoublyLinkedList()

	_, err := dll.Get(0)
	if err == nil {
		t.Error("Expected error when getting from empty list")
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)
	dll.Append(40)
	dll.Append(50)

	tests := []struct {
		index    int
		expected int
		hasError bool
	}{
		{0, 10, false},
		{1, 20, false},
		{2, 30, false},
		{3, 40, false},
		{4, 50, false},
		{5, 0, true},
		{-1, 0, true},
	}

	for _, test := range tests {
		value, err := dll.Get(test.index)
		if test.hasError && err == nil {
			t.Errorf("Get(%d): expected error", test.index)
		} else if !test.hasError && err != nil {
			t.Errorf("Get(%d): unexpected error: %v", test.index, err)
		} else if !test.hasError && value != test.expected {
			t.Errorf("Get(%d): expected %d, got %d", test.index, test.expected, value)
		}
	}
}

func TestRemove(t *testing.T) {
	dll := NewDoublyLinkedList()

	value, found := dll.Remove(10)
	if found || value != 0 {
		t.Error("Expected false and 0 when removing from empty list")
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	value, found = dll.Remove(20)
	if !found || value != 20 {
		t.Errorf("Expected true and 20, got %t and %d", found, value)
	}

	expected := []int{10, 30}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}

	value, found = dll.Remove(10)
	if !found || value != 10 {
		t.Errorf("Expected true and 10, got %t and %d", found, value)
	}

	expected = []int{30}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}

	value, found = dll.Remove(40)
	if found || value != 0 {
		t.Errorf("Expected false and 0, got %t and %d", found, value)
	}
}

func TestRemoveAt(t *testing.T) {
	dll := NewDoublyLinkedList()

	_, err := dll.RemoveAt(0)
	if err == nil {
		t.Error("Expected error when removing from empty list")
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)
	dll.Append(40)
	dll.Append(50)

	value, err := dll.RemoveAt(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected removed value to be 30, got %d", value)
	}

	expected := []int{10, 20, 40, 50}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}

	value, err = dll.RemoveAt(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected removed value to be 10, got %d", value)
	}

	value, err = dll.RemoveAt(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 50 {
		t.Errorf("Expected removed value to be 50, got %d", value)
	}

	_, err = dll.RemoveAt(5)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestRemoveHeadAndTail(t *testing.T) {
	dll := NewDoublyLinkedList()

	_, err := dll.RemoveHead()
	if err == nil {
		t.Error("Expected error when removing head from empty list")
	}

	_, err = dll.RemoveTail()
	if err == nil {
		t.Error("Expected error when removing tail from empty list")
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	value, err := dll.RemoveHead()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 10 {
		t.Errorf("Expected removed head to be 10, got %d", value)
	}

	value, err = dll.RemoveTail()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != 30 {
		t.Errorf("Expected removed tail to be 30, got %d", value)
	}

	expected := []int{20}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, dll.ToSlice())
	}
}

func TestSearch(t *testing.T) {
	dll := NewDoublyLinkedList()

	index := dll.Search(10)
	if index != -1 {
		t.Errorf("Expected -1 for search in empty list, got %d", index)
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	tests := []struct {
		value    int
		expected int
	}{
		{10, 0},
		{20, 1},
		{30, 2},
		{40, -1},
	}

	for _, test := range tests {
		index := dll.Search(test.value)
		if index != test.expected {
			t.Errorf("Search(%d): expected %d, got %d", test.value, test.expected, index)
		}
	}
}

func TestToSliceReverse(t *testing.T) {
	dll := NewDoublyLinkedList()

	reverse := dll.ToSliceReverse()
	if len(reverse) != 0 {
		t.Errorf("Expected empty slice for empty list, got %v", reverse)
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	forward := dll.ToSlice()
	reverse = dll.ToSliceReverse()

	expectedForward := []int{10, 20, 30}
	expectedReverse := []int{30, 20, 10}

	if !reflect.DeepEqual(forward, expectedForward) {
		t.Errorf("Expected forward %v, got %v", expectedForward, forward)
	}

	if !reflect.DeepEqual(reverse, expectedReverse) {
		t.Errorf("Expected reverse %v, got %v", expectedReverse, reverse)
	}
}

func TestReverse(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.Reverse()
	if dll.Head != nil {
		t.Error("Expected head to remain nil after reversing empty list")
	}

	dll.Append(10)
	dll.Reverse()
	expected := []int{10}
	if !reflect.DeepEqual(dll.ToSlice(), expected) {
		t.Errorf("Expected %v after reversing single element, got %v", expected, dll.ToSlice())
	}

	dll.Clear()
	dll.Append(10)
	dll.Append(20)
	dll.Append(30)
	dll.Append(40)

	originalForward := dll.ToSlice()
	originalReverse := dll.ToSliceReverse()

	dll.Reverse()

	newForward := dll.ToSlice()
	newReverse := dll.ToSliceReverse()

	if !reflect.DeepEqual(newForward, originalReverse) {
		t.Errorf("Expected forward after reverse %v, got %v", originalReverse, newForward)
	}

	if !reflect.DeepEqual(newReverse, originalForward) {
		t.Errorf("Expected reverse after reverse %v, got %v", originalForward, newReverse)
	}
}

func TestGetMiddle(t *testing.T) {
	dll := NewDoublyLinkedList()

	_, err := dll.GetMiddle()
	if err == nil {
		t.Error("Expected error when getting middle of empty list")
	}

	dll.Append(10)
	middle, err := dll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 10 {
		t.Errorf("Expected middle to be 10, got %d", middle)
	}

	dll.Append(20)
	middle, err = dll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 20 {
		t.Errorf("Expected middle to be 20, got %d", middle)
	}

	dll.Append(30)
	middle, err = dll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 20 {
		t.Errorf("Expected middle to be 20, got %d", middle)
	}

	dll.Append(40)
	dll.Append(50)
	middle, err = dll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 30 {
		t.Errorf("Expected middle to be 30, got %d", middle)
	}
}

func TestGetHeadAndTail(t *testing.T) {
	dll := NewDoublyLinkedList()

	_, err := dll.GetHead()
	if err == nil {
		t.Error("Expected error when getting head of empty list")
	}

	_, err = dll.GetTail()
	if err == nil {
		t.Error("Expected error when getting tail of empty list")
	}

	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	head, err := dll.GetHead()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if head != 10 {
		t.Errorf("Expected head to be 10, got %d", head)
	}

	tail, err := dll.GetTail()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tail != 30 {
		t.Errorf("Expected tail to be 30, got %d", tail)
	}
}

func TestDisplay(t *testing.T) {
	dll := NewDoublyLinkedList()

	display := dll.Display()
	if display != "[]" {
		t.Errorf("Expected '[]' for empty list, got '%s'", display)
	}

	dll.Append(10)
	display = dll.Display()
	if display != "[10]" {
		t.Errorf("Expected '[10]' for single element, got '%s'", display)
	}

	dll.Append(20)
	dll.Append(30)
	display = dll.Display()
	if display != "[10 <-> 20 <-> 30]" {
		t.Errorf("Expected '[10 <-> 20 <-> 30]', got '%s'", display)
	}
}

func TestClear(t *testing.T) {
	dll := NewDoublyLinkedList()
	dll.Append(10)
	dll.Append(20)
	dll.Append(30)

	dll.Clear()

	if dll.Head != nil {
		t.Error("Expected head to be nil after clear")
	}

	if dll.Tail != nil {
		t.Error("Expected tail to be nil after clear")
	}

	if dll.Length() != 0 {
		t.Errorf("Expected length to be 0 after clear, got %d", dll.Length())
	}

	if !dll.IsEmpty() {
		t.Error("Expected list to be empty after clear")
	}
}

func TestBidirectionalIntegrity(t *testing.T) {
	dll := NewDoublyLinkedList()

	for i := 1; i <= 5; i++ {
		dll.Append(i * 10)
	}

	current := dll.Head
	for current != nil {
		if current.Next != nil && current.Next.Prev != current {
			t.Error("Forward link integrity broken")
		}
		if current.Prev != nil && current.Prev.Next != current {
			t.Error("Backward link integrity broken")
		}
		current = current.Next
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

	requiredKeys := []string{"original_list", "length", "data_structure", "bidirectional"}
	for _, key := range requiredKeys {
		if _, exists := resultMap[key]; !exists {
			t.Errorf("Expected result to contain '%s' key", key)
		}
	}
}

func BenchmarkPrepend(b *testing.B) {
	dll := NewDoublyLinkedList()

	for b.Loop() {
		dll.Prepend(42)
	}
}

func BenchmarkAppend(b *testing.B) {
	dll := NewDoublyLinkedList()

	for b.Loop() {
		dll.Append(42)
	}
}

func BenchmarkGetFromStart(b *testing.B) {
	dll := NewDoublyLinkedList()
	for i := range 1000 {
		dll.Append(i)
	}

	for b.Loop() {
		dll.Get(100)
	}
}

func BenchmarkGetFromEnd(b *testing.B) {
	dll := NewDoublyLinkedList()
	for i := range 1000 {
		dll.Append(i)
	}

	for b.Loop() {
		dll.Get(900)
	}
}

func BenchmarkRemove(b *testing.B) {
	for b.Loop() {
		dll := NewDoublyLinkedList()
		for i := range 100 {
			dll.Append(i)
		}

		dll.Remove(50)
	}
}

func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		dll := NewDoublyLinkedList()
		for i := range 1000 {
			dll.Append(i)
		}

		dll.Reverse()
	}
}
