package avl_tree

import (
	"reflect"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	avl := NewAVLTree()
	if avl.Root != nil {
		t.Error("Expected root to be nil")
	}
	if avl.Size != 0 {
		t.Error("Expected size to be 0")
	}
	if !avl.IsEmpty() {
		t.Error("Expected tree to be empty")
	}
}

func TestInsertSingle(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(10)

	if avl.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", avl.GetSize())
	}
	if !avl.Search(10) {
		t.Error("Expected to find value 10")
	}
	if avl.GetHeight() != 0 {
		t.Errorf("Expected height 0, got %d", avl.GetHeight())
	}
}

func TestInsertMultiple(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		avl.Insert(value)
	}

	if avl.GetSize() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), avl.GetSize())
	}

	for _, value := range values {
		if !avl.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}
}

func TestInsertDuplicate(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(10)
	avl.Insert(10)

	if avl.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", avl.GetSize())
	}
}

func TestInsertIterative(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		avl.InsertIterative(value)
	}

	if avl.GetSize() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), avl.GetSize())
	}

	for _, value := range values {
		if !avl.SearchIterative(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced after iterative insertions")
	}
}

func TestRightRotation(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(30)
	avl.Insert(20)
	avl.Insert(10)

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced after right rotation")
	}

	expected := []int{10, 20, 30}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLeftRotation(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(10)
	avl.Insert(20)
	avl.Insert(30)

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced after left rotation")
	}

	expected := []int{10, 20, 30}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestLeftRightRotation(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(30)
	avl.Insert(10)
	avl.Insert(20)

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced after left-right rotation")
	}

	expected := []int{10, 20, 30}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestRightLeftRotation(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(20)

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced after right-left rotation")
	}

	expected := []int{10, 20, 30}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDelete(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		avl.Insert(value)
	}

	if !avl.Delete(20) {
		t.Error("Expected successful deletion")
	}

	if avl.Search(20) {
		t.Error("Expected value 20 to be deleted")
	}

	if avl.GetSize() != len(values)-1 {
		t.Errorf("Expected size %d, got %d", len(values)-1, avl.GetSize())
	}

	if !avl.IsBalanced() {
		t.Error("Expected tree to remain balanced after deletion")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(10)

	if avl.Delete(20) {
		t.Error("Expected deletion to fail for non-existent value")
	}

	if avl.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", avl.GetSize())
	}
}

func TestDeleteLeaf(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)

	if !avl.Delete(10) {
		t.Error("Expected successful deletion of leaf")
	}

	expected := []int{20, 30}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	avl := NewAVLTree()
	avl.Insert(20)
	avl.Insert(10)
	avl.Insert(30)
	avl.Insert(25)

	if !avl.Delete(30) {
		t.Error("Expected successful deletion of node with one child")
	}

	expected := []int{10, 20, 25}
	actual := avl.InOrderTraversal()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	avl := NewAVLTree()
	values := []int{20, 10, 30, 5, 15, 25, 35}

	for _, value := range values {
		avl.Insert(value)
	}

	if !avl.Delete(20) {
		t.Error("Expected successful deletion of node with two children")
	}

	if avl.Search(20) {
		t.Error("Expected value 20 to be deleted")
	}

	if !avl.IsBalanced() {
		t.Error("Expected tree to remain balanced after deletion")
	}
}

func TestSearchRecursive(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		avl.Insert(value)
	}

	for _, value := range values {
		if !avl.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if avl.Search(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestSearchIterative(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		avl.Insert(value)
	}

	for _, value := range values {
		if !avl.SearchIterative(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if avl.SearchIterative(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestInOrderTraversal(t *testing.T) {
	avl := NewAVLTree()
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		avl.Insert(value)
	}

	expected := []int{5, 10, 20, 30, 35, 40, 50}
	actual := avl.InOrderTraversal()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	avl := NewAVLTree()
	values := []int{20, 10, 30}

	for _, value := range values {
		avl.Insert(value)
	}

	actual := avl.PreOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}

	if actual[0] != 20 {
		t.Errorf("Expected root to be 20, got %d", actual[0])
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	avl := NewAVLTree()
	values := []int{20, 10, 30}

	for _, value := range values {
		avl.Insert(value)
	}

	actual := avl.LevelOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}

	if actual[0] != 20 {
		t.Errorf("Expected first element to be 20, got %d", actual[0])
	}
}

func TestFindMinMax(t *testing.T) {
	avl := NewAVLTree()
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		avl.Insert(value)
	}

	min, hasMin := avl.FindMin()
	if !hasMin || min != 5 {
		t.Errorf("Expected min 5, got %d (exists: %v)", min, hasMin)
	}

	max, hasMax := avl.FindMax()
	if !hasMax || max != 50 {
		t.Errorf("Expected max 50, got %d (exists: %v)", max, hasMax)
	}
}

func TestFindMinMaxEmpty(t *testing.T) {
	avl := NewAVLTree()

	_, hasMin := avl.FindMin()
	if hasMin {
		t.Error("Expected no min in empty tree")
	}

	_, hasMax := avl.FindMax()
	if hasMax {
		t.Error("Expected no max in empty tree")
	}
}

func TestIsBalanced(t *testing.T) {
	avl := NewAVLTree()

	if !avl.IsBalanced() {
		t.Error("Expected empty tree to be balanced")
	}

	values := []int{10, 20, 30, 40, 50, 25, 15, 35}
	for _, value := range values {
		avl.Insert(value)
		if !avl.IsBalanced() {
			t.Errorf("Expected tree to be balanced after inserting %d", value)
		}
	}
}

func TestClear(t *testing.T) {
	avl := NewAVLTree()
	values := []int{10, 20, 30}

	for _, value := range values {
		avl.Insert(value)
	}

	avl.Clear()

	if !avl.IsEmpty() {
		t.Error("Expected tree to be empty after clear")
	}

	if avl.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", avl.GetSize())
	}

	if avl.GetHeight() != -1 {
		t.Errorf("Expected height -1, got %d", avl.GetHeight())
	}
}

func TestLargeDataset(t *testing.T) {
	avl := NewAVLTree()

	for i := 1; i <= 1000; i++ {
		avl.Insert(i)
	}

	if avl.GetSize() != 1000 {
		t.Errorf("Expected size 1000, got %d", avl.GetSize())
	}

	if !avl.IsBalanced() {
		t.Error("Expected tree to be balanced with large dataset")
	}

	for i := 1; i <= 1000; i++ {
		if !avl.Search(i) {
			t.Errorf("Expected to find value %d", i)
		}
	}

	for i := 1; i <= 500; i++ {
		if !avl.Delete(i) {
			t.Errorf("Expected to delete value %d", i)
		}
	}

	if avl.GetSize() != 500 {
		t.Errorf("Expected size 500, got %d", avl.GetSize())
	}

	if !avl.IsBalanced() {
		t.Error("Expected tree to remain balanced after deletions")
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

	if resultMap["size"].(int) != 6 {
		t.Errorf("Expected size 6, got %v", resultMap["size"])
	}

	if !resultMap["isBalanced"].(bool) {
		t.Error("Expected tree to be balanced")
	}

	if !resultMap["searchExisting"].(bool) {
		t.Error("Expected to find existing value")
	}

	if resultMap["searchNonExisting"].(bool) {
		t.Error("Expected not to find non-existing value")
	}
}

func BenchmarkInsert(b *testing.B) {
	avl := NewAVLTree()
	for b.Loop() {
		avl.Insert(b.N)
	}
}

func BenchmarkInsertIterative(b *testing.B) {
	avl := NewAVLTree()
	for b.Loop() {
		avl.InsertIterative(b.N)
	}
}

func BenchmarkSearch(b *testing.B) {
	avl := NewAVLTree()
	for i := range 1000 {
		avl.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		avl.Search(b.N % 1000)
	}
}

func BenchmarkSearchIterative(b *testing.B) {
	avl := NewAVLTree()
	for i := range 1000 {
		avl.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		avl.SearchIterative(b.N % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	for b.Loop() {
		avl := NewAVLTree()
		for i := range 100 {
			avl.Insert(i)
		}
		avl.Delete(50)
	}
}
