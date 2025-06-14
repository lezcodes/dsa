package linked_list

import (
	"reflect"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()

	if ll.Head != nil {
		t.Error("Expected head to be nil for new list")
	}

	if ll.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", ll.Size())
	}

	if !ll.IsEmpty() {
		t.Error("Expected new list to be empty")
	}
}

func TestInsertAtHead(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtHead(10)
	if ll.Head.Data != 10 {
		t.Errorf("Expected head data to be 10, got %d", ll.Head.Data)
	}

	if ll.Size() != 1 {
		t.Errorf("Expected size to be 1, got %d", ll.Size())
	}

	ll.InsertAtHead(20)
	if ll.Head.Data != 20 {
		t.Errorf("Expected head data to be 20, got %d", ll.Head.Data)
	}

	if ll.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", ll.Size())
	}

	expected := []int{20, 10}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}
}

func TestInsertAtTail(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAtTail(10)
	if ll.Head.Data != 10 {
		t.Errorf("Expected head data to be 10, got %d", ll.Head.Data)
	}

	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	if ll.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", ll.Size())
	}
}

func TestInsertAtIndex(t *testing.T) {
	ll := NewLinkedList()

	err := ll.InsertAtIndex(0, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = ll.InsertAtIndex(1, 30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = ll.InsertAtIndex(1, 20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{10, 20, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	err = ll.InsertAtIndex(-1, 5)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	err = ll.InsertAtIndex(10, 5)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestDeleteByValue(t *testing.T) {
	ll := NewLinkedList()

	found := ll.DeleteByValue(10)
	if found {
		t.Error("Expected false when deleting from empty list")
	}

	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	found = ll.DeleteByValue(20)
	if !found {
		t.Error("Expected true when deleting existing element")
	}

	expected := []int{10, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	found = ll.DeleteByValue(10)
	if !found {
		t.Error("Expected true when deleting head element")
	}

	expected = []int{30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	found = ll.DeleteByValue(40)
	if found {
		t.Error("Expected false when deleting non-existent element")
	}
}

func TestDeleteAtIndex(t *testing.T) {
	ll := NewLinkedList()

	err := ll.DeleteAtIndex(0)
	if err == nil {
		t.Error("Expected error when deleting from empty list")
	}

	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	err = ll.DeleteAtIndex(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []int{10, 30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	err = ll.DeleteAtIndex(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected = []int{30}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v, got %v", expected, ll.ToSlice())
	}

	err = ll.DeleteAtIndex(5)
	if err == nil {
		t.Error("Expected error for index out of bounds")
	}
}

func TestSearch(t *testing.T) {
	ll := NewLinkedList()

	index := ll.Search(10)
	if index != -1 {
		t.Errorf("Expected -1 for search in empty list, got %d", index)
	}

	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

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
		index := ll.Search(test.value)
		if index != test.expected {
			t.Errorf("Search(%d): expected %d, got %d", test.value, test.expected, index)
		}
	}
}

func TestGetAt(t *testing.T) {
	ll := NewLinkedList()

	_, err := ll.GetAt(0)
	if err == nil {
		t.Error("Expected error when getting from empty list")
	}

	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	tests := []struct {
		index    int
		expected int
		hasError bool
	}{
		{0, 10, false},
		{1, 20, false},
		{2, 30, false},
		{3, 0, true},
		{-1, 0, true},
	}

	for _, test := range tests {
		value, err := ll.GetAt(test.index)
		if test.hasError && err == nil {
			t.Errorf("GetAt(%d): expected error", test.index)
		} else if !test.hasError && err != nil {
			t.Errorf("GetAt(%d): unexpected error: %v", test.index, err)
		} else if !test.hasError && value != test.expected {
			t.Errorf("GetAt(%d): expected %d, got %d", test.index, test.expected, value)
		}
	}
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList()

	ll.Reverse()
	if ll.Head != nil {
		t.Error("Expected head to remain nil after reversing empty list")
	}

	ll.InsertAtTail(10)
	ll.Reverse()
	expected := []int{10}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v after reversing single element, got %v", expected, ll.ToSlice())
	}

	ll.Clear()
	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)
	ll.InsertAtTail(40)

	ll.Reverse()
	expected = []int{40, 30, 20, 10}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("Expected %v after reversing, got %v", expected, ll.ToSlice())
	}
}

func TestGetMiddle(t *testing.T) {
	ll := NewLinkedList()

	_, err := ll.GetMiddle()
	if err == nil {
		t.Error("Expected error when getting middle of empty list")
	}

	ll.InsertAtTail(10)
	middle, err := ll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 10 {
		t.Errorf("Expected middle to be 10, got %d", middle)
	}

	ll.InsertAtTail(20)
	middle, err = ll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 20 {
		t.Errorf("Expected middle to be 20, got %d", middle)
	}

	ll.InsertAtTail(30)
	middle, err = ll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 20 {
		t.Errorf("Expected middle to be 20, got %d", middle)
	}

	ll.InsertAtTail(40)
	ll.InsertAtTail(50)
	middle, err = ll.GetMiddle()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if middle != 30 {
		t.Errorf("Expected middle to be 30, got %d", middle)
	}
}

func TestContains(t *testing.T) {
	ll := NewLinkedList()

	if ll.Contains(10) {
		t.Error("Expected false for empty list")
	}

	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	if !ll.Contains(20) {
		t.Error("Expected true for existing element")
	}

	if ll.Contains(40) {
		t.Error("Expected false for non-existing element")
	}
}

func TestDisplay(t *testing.T) {
	ll := NewLinkedList()

	display := ll.Display()
	if display != "[]" {
		t.Errorf("Expected '[]' for empty list, got '%s'", display)
	}

	ll.InsertAtTail(10)
	display = ll.Display()
	if display != "[10]" {
		t.Errorf("Expected '[10]' for single element, got '%s'", display)
	}

	ll.InsertAtTail(20)
	ll.InsertAtTail(30)
	display = ll.Display()
	if display != "[10 -> 20 -> 30]" {
		t.Errorf("Expected '[10 -> 20 -> 30]', got '%s'", display)
	}
}

func TestClear(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertAtTail(10)
	ll.InsertAtTail(20)
	ll.InsertAtTail(30)

	ll.Clear()

	if ll.Head != nil {
		t.Error("Expected head to be nil after clear")
	}

	if ll.Size() != 0 {
		t.Errorf("Expected size to be 0 after clear, got %d", ll.Size())
	}

	if !ll.IsEmpty() {
		t.Error("Expected list to be empty after clear")
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

	requiredKeys := []string{"original_list", "size", "data_structure"}
	for _, key := range requiredKeys {
		if _, exists := resultMap[key]; !exists {
			t.Errorf("Expected result to contain '%s' key", key)
		}
	}
}

func BenchmarkInsertAtHead(b *testing.B) {
	ll := NewLinkedList()

	for b.Loop() {
		ll.InsertAtHead(42)
	}
}

func BenchmarkInsertAtTail(b *testing.B) {
	ll := NewLinkedList()

	for b.Loop() {
		ll.InsertAtTail(42)
	}
}

func BenchmarkSearch(b *testing.B) {
	ll := NewLinkedList()
	for i := range 1000 {
		ll.InsertAtTail(i)
	}

	for b.Loop() {
		ll.Search(500)
	}
}

func BenchmarkDeleteByValue(b *testing.B) {
	for b.Loop() {
		ll := NewLinkedList()
		for i := range 100 {
			ll.InsertAtTail(i)
		}

		ll.DeleteByValue(50)
	}
}

func BenchmarkReverse(b *testing.B) {
	for b.Loop() {
		ll := NewLinkedList()
		for i := range 1000 {
			ll.InsertAtTail(i)
		}

		ll.Reverse()
	}
}
