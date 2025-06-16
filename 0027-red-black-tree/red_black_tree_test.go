package red_black_tree

import (
	"reflect"
	"testing"
)

func TestNewRBTree(t *testing.T) {
	rb := NewRBTree()
	if rb.Root != rb.NIL {
		t.Error("Expected root to be NIL")
	}
	if rb.Size != 0 {
		t.Error("Expected size to be 0")
	}
	if !rb.IsEmpty() {
		t.Error("Expected tree to be empty")
	}
}

func TestInsertSingle(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(10)

	if rb.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", rb.GetSize())
	}
	if !rb.Search(10) {
		t.Error("Expected to find value 10")
	}
	if rb.Root.Color != BLACK {
		t.Error("Expected root to be BLACK")
	}
}

func TestInsertMultiple(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		rb.Insert(value)
	}

	if rb.GetSize() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), rb.GetSize())
	}

	for _, value := range values {
		if !rb.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties")
	}
}

func TestInsertDuplicate(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(10)
	rb.Insert(10)

	if rb.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", rb.GetSize())
	}
}

func TestInsertIterative(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		rb.InsertIterative(value)
	}

	if rb.GetSize() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), rb.GetSize())
	}

	for _, value := range values {
		if !rb.SearchIterative(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after iterative insertions")
	}
}

func TestRootAlwaysBlack(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		rb.Insert(value)
		if rb.Root.Color != BLACK {
			t.Errorf("Expected root to be BLACK after inserting %d", value)
		}
	}
}

func TestRedNodeHasBlackChildren(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50, 25, 15, 35}

	for _, value := range values {
		rb.Insert(value)
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties")
	}
}

func TestDelete(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50, 25}

	for _, value := range values {
		rb.Insert(value)
	}

	if !rb.Delete(20) {
		t.Error("Expected successful deletion")
	}

	if rb.Search(20) {
		t.Error("Expected value 20 to be deleted")
	}

	if rb.GetSize() != len(values)-1 {
		t.Errorf("Expected size %d, got %d", len(values)-1, rb.GetSize())
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after deletion")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(10)

	if rb.Delete(20) {
		t.Error("Expected deletion to fail for non-existent value")
	}

	if rb.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", rb.GetSize())
	}
}

func TestDeleteLeaf(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(20)
	rb.Insert(10)
	rb.Insert(30)

	if !rb.Delete(10) {
		t.Error("Expected successful deletion of leaf")
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after leaf deletion")
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(20)
	rb.Insert(10)
	rb.Insert(30)
	rb.Insert(25)

	if !rb.Delete(30) {
		t.Error("Expected successful deletion of node with one child")
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after deletion")
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	rb := NewRBTree()
	values := []int{20, 10, 30, 5, 15, 25, 35}

	for _, value := range values {
		rb.Insert(value)
	}

	if !rb.Delete(20) {
		t.Error("Expected successful deletion of node with two children")
	}

	if rb.Search(20) {
		t.Error("Expected value 20 to be deleted")
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after deletion")
	}
}

func TestSearchRecursive(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		rb.Insert(value)
	}

	for _, value := range values {
		if !rb.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if rb.Search(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestSearchIterative(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		rb.Insert(value)
	}

	for _, value := range values {
		if !rb.SearchIterative(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if rb.SearchIterative(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestInOrderTraversal(t *testing.T) {
	rb := NewRBTree()
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		rb.Insert(value)
	}

	expected := []int{5, 10, 20, 30, 35, 40, 50}
	actual := rb.InOrderTraversal()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	rb := NewRBTree()
	values := []int{20, 10, 30}

	for _, value := range values {
		rb.Insert(value)
	}

	actual := rb.PreOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}

	if actual[0] != 20 {
		t.Errorf("Expected root to be 20, got %d", actual[0])
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	rb := NewRBTree()
	values := []int{20, 10, 30}

	for _, value := range values {
		rb.Insert(value)
	}

	actual := rb.LevelOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}

	if actual[0] != 20 {
		t.Errorf("Expected first element to be 20, got %d", actual[0])
	}
}

func TestFindMinMax(t *testing.T) {
	rb := NewRBTree()
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		rb.Insert(value)
	}

	min, hasMin := rb.FindMin()
	if !hasMin || min != 5 {
		t.Errorf("Expected min 5, got %d (exists: %v)", min, hasMin)
	}

	max, hasMax := rb.FindMax()
	if !hasMax || max != 50 {
		t.Errorf("Expected max 50, got %d (exists: %v)", max, hasMax)
	}
}

func TestFindMinMaxEmpty(t *testing.T) {
	rb := NewRBTree()

	_, hasMin := rb.FindMin()
	if hasMin {
		t.Error("Expected no min in empty tree")
	}

	_, hasMax := rb.FindMax()
	if hasMax {
		t.Error("Expected no max in empty tree")
	}
}

func TestGetBlackHeight(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		rb.Insert(value)
	}

	blackHeight := rb.GetBlackHeight()
	if blackHeight < 1 {
		t.Errorf("Expected black height >= 1, got %d", blackHeight)
	}
}

func TestIsValidRBTree(t *testing.T) {
	rb := NewRBTree()

	if !rb.IsValidRBTree() {
		t.Error("Expected empty tree to be valid")
	}

	values := []int{10, 20, 30, 40, 50, 25, 15, 35}
	for _, value := range values {
		rb.Insert(value)
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to be valid after inserting %d", value)
		}
	}

	for i := range len(values) / 2 {
		rb.Delete(values[i])
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to be valid after deleting %d", values[i])
		}
	}
}

func TestClear(t *testing.T) {
	rb := NewRBTree()
	values := []int{10, 20, 30}

	for _, value := range values {
		rb.Insert(value)
	}

	rb.Clear()

	if !rb.IsEmpty() {
		t.Error("Expected tree to be empty after clear")
	}

	if rb.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", rb.GetSize())
	}

	if rb.GetHeight() != -1 {
		t.Errorf("Expected height -1, got %d", rb.GetHeight())
	}
}

func TestRotations(t *testing.T) {
	rb := NewRBTree()

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range values {
		rb.Insert(value)
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to remain valid after inserting %d", value)
		}
	}

	height := rb.GetHeight()
	if height > 6 {
		t.Errorf("Expected height <= 6 for balanced tree, got %d", height)
	}
}

func TestLargeDataset(t *testing.T) {
	rb := NewRBTree()

	for i := 1; i <= 1000; i++ {
		rb.Insert(i)
	}

	if rb.GetSize() != 1000 {
		t.Errorf("Expected size 1000, got %d", rb.GetSize())
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties with large dataset")
	}

	for i := 1; i <= 1000; i++ {
		if !rb.Search(i) {
			t.Errorf("Expected to find value %d", i)
		}
	}

	for i := 1; i <= 500; i++ {
		if !rb.Delete(i) {
			t.Errorf("Expected to delete value %d", i)
		}
	}

	if rb.GetSize() != 500 {
		t.Errorf("Expected size 500, got %d", rb.GetSize())
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to maintain Red-Black properties after deletions")
	}
}

func TestRandomInsertDelete(t *testing.T) {
	rb := NewRBTree()
	values := []int{50, 25, 75, 10, 30, 60, 80, 5, 15, 27, 35}

	for _, value := range values {
		rb.Insert(value)
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to be valid after inserting %d", value)
		}
	}

	deleteValues := []int{25, 60, 10, 80}
	for _, value := range deleteValues {
		if !rb.Delete(value) {
			t.Errorf("Expected to delete value %d", value)
		}
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to be valid after deleting %d", value)
		}
	}
}

func TestComplexDeletions(t *testing.T) {
	rb := NewRBTree()
	values := []int{20, 10, 30, 5, 15, 25, 35, 1, 7, 12, 18, 22, 27, 32, 40}

	for _, value := range values {
		rb.Insert(value)
	}

	if !rb.IsValidRBTree() {
		t.Error("Expected tree to be valid after insertions")
	}

	deleteOrder := []int{1, 7, 12, 18, 22, 27, 32, 40, 5, 15, 25, 35, 10, 30, 20}
	for _, value := range deleteOrder {
		if !rb.Delete(value) {
			t.Errorf("Expected to delete value %d", value)
		}
		if !rb.IsValidRBTree() {
			t.Errorf("Expected tree to be valid after deleting %d", value)
		}
	}

	if !rb.IsEmpty() {
		t.Error("Expected tree to be empty after deleting all values")
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

	if resultMap["size"].(int) != 8 {
		t.Errorf("Expected size 8, got %v", resultMap["size"])
	}

	if !resultMap["isValidRBTree"].(bool) {
		t.Error("Expected tree to be a valid Red-Black tree")
	}

	if !resultMap["searchExisting"].(bool) {
		t.Error("Expected to find existing value")
	}

	if resultMap["searchNonExisting"].(bool) {
		t.Error("Expected not to find non-existing value")
	}

	if !resultMap["isValidAfterDelete"].(bool) {
		t.Error("Expected tree to remain valid after deletion")
	}

	if !resultMap["isValidAfterIterativeInsert"].(bool) {
		t.Error("Expected tree to remain valid after iterative insertion")
	}
}

func BenchmarkInsert(b *testing.B) {
	rb := NewRBTree()
	for b.Loop() {
		rb.Insert(b.N)
	}
}

func BenchmarkInsertIterative(b *testing.B) {
	rb := NewRBTree()
	for b.Loop() {
		rb.InsertIterative(b.N)
	}
}

func BenchmarkSearch(b *testing.B) {
	rb := NewRBTree()
	for i := range 1000 {
		rb.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		rb.Search(b.N % 1000)
	}
}

func BenchmarkSearchIterative(b *testing.B) {
	rb := NewRBTree()
	for i := range 1000 {
		rb.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		rb.SearchIterative(b.N % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	for b.Loop() {
		rb := NewRBTree()
		for i := range 100 {
			rb.Insert(i)
		}
		rb.Delete(50)
	}
}
