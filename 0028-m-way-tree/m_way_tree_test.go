package m_way_tree

import (
	"reflect"
	"testing"
)

func TestNewMWayTree(t *testing.T) {
	tree := NewMWayTree(4)
	if tree.M != 4 {
		t.Errorf("Expected branching factor 4, got %d", tree.M)
	}
	if tree.Size != 0 {
		t.Error("Expected size to be 0")
	}
	if !tree.IsEmpty() {
		t.Error("Expected tree to be empty")
	}
}

func TestNewMWayTreeMinimumBranching(t *testing.T) {
	tree := NewMWayTree(1)
	if tree.M != 3 {
		t.Errorf("Expected minimum branching factor 3, got %d", tree.M)
	}
}

func TestInsertSingle(t *testing.T) {
	tree := NewMWayTree(3)
	tree.Insert(10)

	if tree.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", tree.GetSize())
	}
	if !tree.Search(10) {
		t.Error("Expected to find value 10")
	}
	if tree.GetHeight() != 0 {
		t.Errorf("Expected height 0, got %d", tree.GetHeight())
	}
}

func TestInsertMultiple(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	if tree.GetSize() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), tree.GetSize())
	}

	for _, value := range values {
		if !tree.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := NewMWayTree(3)
	tree.Insert(10)
	tree.Insert(10)

	if tree.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", tree.GetSize())
	}
}

func TestNodeSplitting(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		tree.Insert(value)
		if !tree.Validate() {
			t.Errorf("Tree invalid after inserting %d", value)
		}
	}

	if tree.GetHeight() < 1 {
		t.Error("Expected tree to have grown in height due to splitting")
	}
}

func TestDelete(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	if !tree.Delete(12) {
		t.Error("Expected successful deletion")
	}

	if tree.Search(12) {
		t.Error("Expected value 12 to be deleted")
	}

	if tree.GetSize() != len(values)-1 {
		t.Errorf("Expected size %d, got %d", len(values)-1, tree.GetSize())
	}

	if !tree.Validate() {
		t.Error("Expected tree to remain valid after deletion")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	tree := NewMWayTree(3)
	tree.Insert(10)

	if tree.Delete(20) {
		t.Error("Expected deletion to fail for non-existent value")
	}

	if tree.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", tree.GetSize())
	}
}

func TestDeleteFromLeaf(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		tree.Insert(value)
	}

	if !tree.Delete(50) {
		t.Error("Expected successful deletion from leaf")
	}

	if !tree.Validate() {
		t.Error("Expected tree to remain valid after leaf deletion")
	}
}

func TestDeleteFromInternalNode(t *testing.T) {
	tree := NewMWayTree(3)
	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}

	initialSize := tree.GetSize()
	if !tree.Delete(5) {
		t.Error("Expected successful deletion from internal node")
	}

	if tree.GetSize() != initialSize-1 {
		t.Error("Expected size to decrease by 1")
	}

	if !tree.Validate() {
		t.Error("Expected tree to remain valid after internal node deletion")
	}
}

func TestSearchRecursive(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		tree.Insert(value)
	}

	for _, value := range values {
		if !tree.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if tree.Search(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestSearchIterative(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		tree.Insert(value)
	}

	for _, value := range values {
		if !tree.SearchIterative(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	if tree.SearchIterative(100) {
		t.Error("Expected not to find value 100")
	}
}

func TestInOrderTraversal(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		tree.Insert(value)
	}

	expected := []int{5, 10, 20, 30, 35, 40, 50}
	actual := tree.InOrderTraversal()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{20, 10, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	actual := tree.PreOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}

	allKeys := tree.GetAllKeys()
	if !reflect.DeepEqual([]int{10, 20, 30}, allKeys) {
		t.Errorf("Expected [10, 20, 30], got %v", allKeys)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{20, 10, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	actual := tree.PostOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{20, 10, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	actual := tree.LevelOrderTraversal()
	if len(actual) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(actual))
	}
}

func TestFindMinMax(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{30, 10, 40, 5, 20, 35, 50}

	for _, value := range values {
		tree.Insert(value)
	}

	min, hasMin := tree.FindMin()
	if !hasMin || min != 5 {
		t.Errorf("Expected min 5, got %d (exists: %v)", min, hasMin)
	}

	max, hasMax := tree.FindMax()
	if !hasMax || max != 50 {
		t.Errorf("Expected max 50, got %d (exists: %v)", max, hasMax)
	}
}

func TestFindMinMaxEmpty(t *testing.T) {
	tree := NewMWayTree(3)

	_, hasMin := tree.FindMin()
	if hasMin {
		t.Error("Expected no min in empty tree")
	}

	_, hasMax := tree.FindMax()
	if hasMax {
		t.Error("Expected no max in empty tree")
	}
}

func TestGetAllKeys(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{30, 10, 40, 5, 20}

	for _, value := range values {
		tree.Insert(value)
	}

	expected := []int{5, 10, 20, 30, 40}
	actual := tree.GetAllKeys()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestGetNodeCount(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		tree.Insert(value)
	}

	nodeCount := tree.GetNodeCount()
	if nodeCount < 1 {
		t.Errorf("Expected at least 1 node, got %d", nodeCount)
	}
}

func TestGetLeafCount(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{1, 2, 3, 4, 5}

	for _, value := range values {
		tree.Insert(value)
	}

	leafCount := tree.GetLeafCount()
	if leafCount < 1 {
		t.Errorf("Expected at least 1 leaf, got %d", leafCount)
	}
}

func TestValidate(t *testing.T) {
	tree := NewMWayTree(4)

	if !tree.Validate() {
		t.Error("Expected empty tree to be valid")
	}

	values := []int{10, 20, 30, 40, 50, 25, 15, 35}
	for _, value := range values {
		tree.Insert(value)
		if !tree.Validate() {
			t.Errorf("Expected tree to be valid after inserting %d", value)
		}
	}
}

func TestClear(t *testing.T) {
	tree := NewMWayTree(3)
	values := []int{10, 20, 30}

	for _, value := range values {
		tree.Insert(value)
	}

	tree.Clear()

	if !tree.IsEmpty() {
		t.Error("Expected tree to be empty after clear")
	}

	if tree.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", tree.GetSize())
	}

	if tree.GetHeight() != -1 {
		t.Errorf("Expected height -1, got %d", tree.GetHeight())
	}
}

func TestDifferentBranchingFactors(t *testing.T) {
	branchingFactors := []int{3, 4, 5, 10}

	for _, m := range branchingFactors {
		tree := NewMWayTree(m)

		for i := 1; i <= 20; i++ {
			tree.Insert(i)
		}

		if tree.GetBranchingFactor() != m {
			t.Errorf("Expected branching factor %d, got %d", m, tree.GetBranchingFactor())
		}

		if tree.GetSize() != 20 {
			t.Errorf("Expected size 20, got %d", tree.GetSize())
		}

		if !tree.Validate() {
			t.Errorf("Expected tree with branching factor %d to be valid", m)
		}

		for i := 1; i <= 20; i++ {
			if !tree.Search(i) {
				t.Errorf("Expected to find value %d in tree with branching factor %d", i, m)
			}
		}
	}
}

func TestLargeDataset(t *testing.T) {
	tree := NewMWayTree(5)

	for i := 1; i <= 100; i++ {
		tree.Insert(i)
	}

	if tree.GetSize() != 100 {
		t.Errorf("Expected size 100, got %d", tree.GetSize())
	}

	if !tree.Validate() {
		t.Error("Expected tree to be valid with large dataset")
	}

	for i := 1; i <= 100; i++ {
		if !tree.Search(i) {
			t.Errorf("Expected to find value %d", i)
		}
	}

	deletedCount := 0
	for i := 1; i <= 50; i++ {
		if tree.Delete(i) {
			deletedCount++
		}
	}

	if deletedCount == 0 {
		t.Error("Expected to delete at least some values")
	}

	if !tree.Validate() {
		t.Error("Expected tree to remain valid after deletions")
	}
}

func TestRandomInsertDelete(t *testing.T) {
	tree := NewMWayTree(4)
	values := []int{50, 25, 75, 10, 30, 60, 80, 5, 15, 27, 35}

	for _, value := range values {
		tree.Insert(value)
		if !tree.Validate() {
			t.Errorf("Expected tree to be valid after inserting %d", value)
		}
	}

	initialSize := tree.GetSize()
	deletedCount := 0
	deleteValues := []int{25, 60, 10, 80}
	for _, value := range deleteValues {
		if tree.Delete(value) {
			deletedCount++
		}
		if !tree.Validate() {
			t.Errorf("Expected tree to be valid after attempting to delete %d", value)
		}
	}

	if deletedCount == 0 {
		t.Error("Expected to delete at least some values")
	}

	if tree.GetSize() > initialSize {
		t.Error("Expected size to not increase after deletions")
	}
}

func TestHeightGrowth(t *testing.T) {
	tree := NewMWayTree(3)

	initialHeight := tree.GetHeight()

	for i := 1; i <= 10; i++ {
		tree.Insert(i)
	}

	finalHeight := tree.GetHeight()
	if finalHeight <= initialHeight {
		t.Error("Expected tree height to grow with more insertions")
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

	if resultMap["branchingFactor"].(int) != 4 {
		t.Errorf("Expected branching factor 4, got %v", resultMap["branchingFactor"])
	}

	if resultMap["size"].(int) != 12 {
		t.Errorf("Expected size 12, got %v", resultMap["size"])
	}

	if !resultMap["isValid"].(bool) {
		t.Error("Expected tree to be valid")
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
}

func BenchmarkInsert(b *testing.B) {
	tree := NewMWayTree(4)
	for b.Loop() {
		tree.Insert(b.N)
	}
}

func BenchmarkSearch(b *testing.B) {
	tree := NewMWayTree(4)
	for i := range 1000 {
		tree.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		tree.Search(b.N % 1000)
	}
}

func BenchmarkSearchIterative(b *testing.B) {
	tree := NewMWayTree(4)
	for i := range 1000 {
		tree.Insert(i)
	}

	b.ResetTimer()
	for b.Loop() {
		tree.SearchIterative(b.N % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	for b.Loop() {
		tree := NewMWayTree(4)
		for i := range 100 {
			tree.Insert(i)
		}
		tree.Delete(50)
	}
}
