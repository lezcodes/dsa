package array_list

import (
	"reflect"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	al := NewArrayList[int]()
	if al.Size() != 0 {
		t.Errorf("Expected size 0, got %d", al.Size())
	}
	if al.Capacity() != 10 {
		t.Errorf("Expected capacity 10, got %d", al.Capacity())
	}
	if !al.IsEmpty() {
		t.Error("Expected empty list")
	}
}

func TestNewArrayListWithCapacity(t *testing.T) {
	al := NewArrayListWithCapacity[int](5)
	if al.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %d", al.Capacity())
	}

	al2 := NewArrayListWithCapacity[int](0)
	if al2.Capacity() != 10 {
		t.Errorf("Expected default capacity 10 for invalid input, got %d", al2.Capacity())
	}
}

func TestAddAndGet(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Add(3)

	if al.Size() != 3 {
		t.Errorf("Expected size 3, got %d", al.Size())
	}

	val, err := al.Get(0)
	if err != nil || val != 1 {
		t.Errorf("Expected 1 at index 0, got %d", val)
	}

	val, err = al.Get(2)
	if err != nil || val != 3 {
		t.Errorf("Expected 3 at index 2, got %d", val)
	}
}

func TestInsert(t *testing.T) {
	al := NewArrayList[string]()

	al.Add("first")
	al.Add("third")

	err := al.Insert(1, "second")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := []string{"first", "second", "third"}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAtBeginning(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(2)
	al.Add(3)
	al.Insert(0, 1)

	expected := []int{1, 2, 3}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertAtEnd(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Insert(2, 3)

	expected := []int{1, 2, 3}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestInsertOutOfBounds(t *testing.T) {
	al := NewArrayList[int]()

	err := al.Insert(-1, 1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	err = al.Insert(1, 1)
	if err == nil {
		t.Error("Expected error for index beyond size")
	}
}

func TestRemove(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Add(3)

	removed, err := al.Remove(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if removed != 2 {
		t.Errorf("Expected removed value 2, got %d", removed)
	}

	expected := []int{1, 3}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestRemoveOutOfBounds(t *testing.T) {
	al := NewArrayList[int]()

	_, err := al.Remove(0)
	if err == nil {
		t.Error("Expected error for empty list")
	}

	al.Add(1)

	_, err = al.Remove(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}

	_, err = al.Remove(1)
	if err == nil {
		t.Error("Expected error for index beyond size")
	}
}

func TestRemoveItem(t *testing.T) {
	al := NewArrayList[string]()

	al.Add("apple")
	al.Add("banana")
	al.Add("cherry")

	removed := al.RemoveItem("banana")
	if !removed {
		t.Error("Expected successful removal")
	}

	expected := []string{"apple", "cherry"}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

	removed = al.RemoveItem("grape")
	if removed {
		t.Error("Expected unsuccessful removal for non-existent item")
	}
}

func TestSet(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Add(3)

	err := al.Set(1, 42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	val, _ := al.Get(1)
	if val != 42 {
		t.Errorf("Expected 42 at index 1, got %d", val)
	}
}

func TestIndexOfAndContains(t *testing.T) {
	al := NewArrayList[string]()

	al.Add("apple")
	al.Add("banana")
	al.Add("cherry")

	index := al.IndexOf("banana")
	if index != 1 {
		t.Errorf("Expected index 1 for banana, got %d", index)
	}

	index = al.IndexOf("grape")
	if index != -1 {
		t.Errorf("Expected index -1 for non-existent item, got %d", index)
	}

	if !al.Contains("cherry") {
		t.Error("Expected list to contain cherry")
	}

	if al.Contains("grape") {
		t.Error("Expected list not to contain grape")
	}
}

func TestClear(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(1)
	al.Add(2)
	al.Add(3)

	al.Clear()

	if !al.IsEmpty() {
		t.Error("Expected empty list after clear")
	}
	if al.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", al.Size())
	}
}

func TestFirstAndLast(t *testing.T) {
	al := NewArrayList[int]()

	_, err := al.First()
	if err == nil {
		t.Error("Expected error for empty list")
	}

	_, err = al.Last()
	if err == nil {
		t.Error("Expected error for empty list")
	}

	al.Add(1)
	al.Add(2)
	al.Add(3)

	first, err := al.First()
	if err != nil || first != 1 {
		t.Errorf("Expected first element 1, got %d", first)
	}

	last, err := al.Last()
	if err != nil || last != 3 {
		t.Errorf("Expected last element 3, got %d", last)
	}
}

func TestPrependAndPop(t *testing.T) {
	al := NewArrayList[int]()

	al.Add(2)
	al.Add(3)
	al.Prepend(1)

	expected := []int{1, 2, 3}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v after prepend, got %v", expected, actual)
	}

	popped, err := al.Pop()
	if err != nil || popped != 3 {
		t.Errorf("Expected popped value 3, got %d", popped)
	}

	if al.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", al.Size())
	}
}

func TestDynamicGrowth(t *testing.T) {
	al := NewArrayListWithCapacity[int](2)

	al.Add(1)
	al.Add(2)

	if al.Capacity() != 2 {
		t.Errorf("Expected capacity 2, got %d", al.Capacity())
	}

	al.Add(3)

	if al.Capacity() != 4 {
		t.Errorf("Expected capacity 4 after growth, got %d", al.Capacity())
	}

	expected := []int{1, 2, 3}
	actual := al.ToSlice()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDynamicShrinking(t *testing.T) {
	al := NewArrayListWithCapacity[int](40)

	for i := 1; i <= 40; i++ {
		al.Add(i)
	}

	for range 38 {
		al.Pop()
	}

	if al.Capacity() != 10 {
		t.Errorf("Expected capacity 10 after shrinking (capped at minimum), got %d", al.Capacity())
	}

	if al.Size() != 2 {
		t.Errorf("Expected size 2 after removals, got %d", al.Size())
	}
}

func TestToSliceEmpty(t *testing.T) {
	al := NewArrayList[int]()

	slice := al.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got %v", slice)
	}
}

func TestString(t *testing.T) {
	al := NewArrayList[int]()
	al.Add(1)
	al.Add(2)

	str := al.String()
	if str == "" {
		t.Error("Expected non-empty string representation")
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

	finalList, ok := data["final_list"].([]string)
	if !ok {
		t.Error("Expected string slice for final_list")
	}

	expected := []string{"apricot", "blueberry", "cherry"}
	if !reflect.DeepEqual(expected, finalList) {
		t.Errorf("Expected %v, got %v", expected, finalList)
	}
}

func BenchmarkAdd(b *testing.B) {
	al := NewArrayList[int]()

	for i := 0; b.Loop(); i++ {
		al.Add(i)
	}
}

func BenchmarkGet(b *testing.B) {
	al := NewArrayList[int]()
	for i := range 1000 {
		al.Add(i)
	}

	for i := 0; b.Loop(); i++ {
		al.Get(i % 1000)
	}
}

func BenchmarkInsertMiddle(b *testing.B) {
	al := NewArrayList[int]()
	for i := range 1000 {
		al.Add(i)
	}

	for i := 0; b.Loop(); i++ {
		al.Insert(500, i)
		al.Remove(500)
	}
}

func BenchmarkRemoveMiddle(b *testing.B) {
	al := NewArrayList[int]()
	for i := range 2000 {
		al.Add(i)
	}

	for i := 0; b.Loop(); i++ {
		if al.Size() > 500 {
			al.Remove(al.Size() / 2)
		} else {
			al.Add(i)
		}
	}
}
