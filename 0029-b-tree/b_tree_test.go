package b_tree

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}
}

func TestNewBTree(t *testing.T) {
	bt := NewBTree(3)
	if bt.T != 3 {
		t.Errorf("Expected minimum degree 3, got %d", bt.T)
	}
	if bt.Root != nil {
		t.Error("Expected nil root for new tree")
	}
	if bt.Size != 0 {
		t.Errorf("Expected size 0, got %d", bt.Size)
	}
}

func TestNewBTreeMinimumDegree(t *testing.T) {
	bt := NewBTree(1)
	if bt.T != 2 {
		t.Errorf("Expected minimum degree 2 for input 1, got %d", bt.T)
	}
}

func TestInsertSingle(t *testing.T) {
	bt := NewBTree(3)
	bt.Insert(10)

	if bt.Size != 1 {
		t.Errorf("Expected size 1, got %d", bt.Size)
	}
	if !bt.Search(10) {
		t.Error("Expected to find inserted key 10")
	}
}

func TestInsertMultiple(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, value := range values {
		bt.Insert(value)
	}

	if bt.Size != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), bt.Size)
	}

	for _, value := range values {
		if !bt.Search(value) {
			t.Errorf("Expected to find inserted key %d", value)
		}
	}
}

func TestInsertDuplicates(t *testing.T) {
	bt := NewBTree(3)
	bt.Insert(10)
	bt.Insert(10)
	bt.Insert(10)

	if bt.Size != 3 {
		t.Errorf("Expected size 3 with duplicates, got %d", bt.Size)
	}
}

func TestInsertCausesSplit(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 30, 40, 50}

	for _, value := range values {
		bt.Insert(value)
	}

	if !bt.Validate() {
		t.Error("Tree should be valid after splits")
	}

	for _, value := range values {
		if !bt.Search(value) {
			t.Errorf("Expected to find key %d after splits", value)
		}
	}
}

func TestSearch(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		bt.Insert(value)
	}

	for _, value := range values {
		if !bt.Search(value) {
			t.Errorf("Expected to find key %d", value)
		}
	}

	if bt.Search(100) {
		t.Error("Should not find non-existent key 100")
	}
}

func TestSearchIterative(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		bt.Insert(value)
	}

	for _, value := range values {
		if !bt.SearchIterative(value) {
			t.Errorf("Expected to find key %d with iterative search", value)
		}
	}

	if bt.SearchIterative(100) {
		t.Error("Should not find non-existent key 100 with iterative search")
	}
}

func TestSearchEmptyTree(t *testing.T) {
	bt := NewBTree(3)

	if bt.Search(10) {
		t.Error("Should not find key in empty tree")
	}

	if bt.SearchIterative(10) {
		t.Error("Should not find key in empty tree with iterative search")
	}
}

func TestDelete(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, value := range values {
		bt.Insert(value)
	}

	initialSize := bt.Size

	if !bt.Delete(12) {
		t.Error("Expected successful deletion of existing key 12")
	}

	if bt.Size != initialSize-1 {
		t.Errorf("Expected size %d after deletion, got %d", initialSize-1, bt.Size)
	}

	if bt.Search(12) {
		t.Error("Should not find deleted key 12")
	}

	if !bt.Validate() {
		t.Error("Tree should be valid after deletion")
	}
}

func TestDeleteNonExistent(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5}

	for _, value := range values {
		bt.Insert(value)
	}

	initialSize := bt.Size

	if bt.Delete(100) {
		t.Error("Should not successfully delete non-existent key")
	}

	if bt.Size != initialSize {
		t.Errorf("Size should remain %d after failed deletion, got %d", initialSize, bt.Size)
	}
}

func TestDeleteFromEmptyTree(t *testing.T) {
	bt := NewBTree(3)

	if bt.Delete(10) {
		t.Error("Should not successfully delete from empty tree")
	}
}

func TestDeleteCausesRebalancing(t *testing.T) {
	bt := NewBTree(3)
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, value := range values {
		bt.Insert(value)
	}

	toDelete := []int{1, 3, 5, 7, 9, 11, 13, 15}
	for _, value := range toDelete {
		if !bt.Delete(value) {
			t.Errorf("Expected successful deletion of key %d", value)
		}

		if !bt.Validate() {
			t.Errorf("Tree should be valid after deleting %d", value)
		}
	}
}

func TestInOrderTraversal(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, value := range values {
		bt.Insert(value)
	}

	result := bt.InOrderTraversal()
	expected := make([]int, len(values))
	copy(expected, values)
	sort.Ints(expected)

	if len(result) != len(expected) {
		t.Errorf("Expected %d elements in traversal, got %d", len(expected), len(result))
	}

	for i, value := range result {
		if value != expected[i] {
			t.Errorf("Expected %d at position %d, got %d", expected[i], i, value)
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5}

	for _, value := range values {
		bt.Insert(value)
	}

	result := bt.PreOrderTraversal()

	if len(result) != len(values) {
		t.Errorf("Expected %d elements in pre-order traversal, got %d", len(values), len(result))
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		bt.Insert(value)
	}

	result := bt.LevelOrderTraversal()

	if len(result) != len(values) {
		t.Errorf("Expected %d elements in level-order traversal, got %d", len(values), len(result))
	}
}

func TestEmptyTreeTraversals(t *testing.T) {
	bt := NewBTree(3)

	inOrder := bt.InOrderTraversal()
	preOrder := bt.PreOrderTraversal()
	levelOrder := bt.LevelOrderTraversal()

	if len(inOrder) != 0 {
		t.Errorf("Expected empty in-order traversal, got %d elements", len(inOrder))
	}

	if len(preOrder) != 0 {
		t.Errorf("Expected empty pre-order traversal, got %d elements", len(preOrder))
	}

	if len(levelOrder) != 0 {
		t.Errorf("Expected empty level-order traversal, got %d elements", len(levelOrder))
	}
}

func TestFindMinMax(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, value := range values {
		bt.Insert(value)
	}

	min, hasMin := bt.FindMin()
	max, hasMax := bt.FindMax()

	if !hasMin {
		t.Error("Expected to find minimum value")
	}

	if !hasMax {
		t.Error("Expected to find maximum value")
	}

	expectedMin := 5
	expectedMax := 30

	if min != expectedMin {
		t.Errorf("Expected minimum %d, got %d", expectedMin, min)
	}

	if max != expectedMax {
		t.Errorf("Expected maximum %d, got %d", expectedMax, max)
	}
}

func TestFindMinMaxEmptyTree(t *testing.T) {
	bt := NewBTree(3)

	_, hasMin := bt.FindMin()
	_, hasMax := bt.FindMax()

	if hasMin {
		t.Error("Should not find minimum in empty tree")
	}

	if hasMax {
		t.Error("Should not find maximum in empty tree")
	}
}

func TestGetHeight(t *testing.T) {
	bt := NewBTree(3)

	if bt.GetHeight() != -1 {
		t.Errorf("Expected height -1 for empty tree, got %d", bt.GetHeight())
	}

	bt.Insert(10)
	if bt.GetHeight() != 0 {
		t.Errorf("Expected height 0 for single node, got %d", bt.GetHeight())
	}

	values := []int{20, 5, 6, 12, 30, 7, 17, 25, 35, 40}
	for _, value := range values {
		bt.Insert(value)
	}

	height := bt.GetHeight()
	if height < 0 {
		t.Errorf("Expected non-negative height, got %d", height)
	}
}

func TestGetSize(t *testing.T) {
	bt := NewBTree(3)

	if bt.GetSize() != 0 {
		t.Errorf("Expected size 0 for empty tree, got %d", bt.GetSize())
	}

	values := []int{10, 20, 5, 6, 12}
	for i, value := range values {
		bt.Insert(value)
		expectedSize := i + 1
		if bt.GetSize() != expectedSize {
			t.Errorf("Expected size %d after inserting %d elements, got %d", expectedSize, i+1, bt.GetSize())
		}
	}
}

func TestIsEmpty(t *testing.T) {
	bt := NewBTree(3)

	if !bt.IsEmpty() {
		t.Error("Expected empty tree to be empty")
	}

	bt.Insert(10)

	if bt.IsEmpty() {
		t.Error("Expected non-empty tree to not be empty")
	}
}

func TestGetNodeCount(t *testing.T) {
	bt := NewBTree(3)

	if bt.GetNodeCount() != 0 {
		t.Errorf("Expected 0 nodes in empty tree, got %d", bt.GetNodeCount())
	}

	bt.Insert(10)
	if bt.GetNodeCount() != 1 {
		t.Errorf("Expected 1 node after single insert, got %d", bt.GetNodeCount())
	}

	values := []int{20, 5, 6, 12, 30, 7, 17}
	for _, value := range values {
		bt.Insert(value)
	}

	nodeCount := bt.GetNodeCount()
	if nodeCount <= 0 {
		t.Errorf("Expected positive node count, got %d", nodeCount)
	}
}

func TestGetLeafCount(t *testing.T) {
	bt := NewBTree(3)

	if bt.GetLeafCount() != 0 {
		t.Errorf("Expected 0 leaves in empty tree, got %d", bt.GetLeafCount())
	}

	bt.Insert(10)
	if bt.GetLeafCount() != 1 {
		t.Errorf("Expected 1 leaf for single node, got %d", bt.GetLeafCount())
	}

	values := []int{20, 5, 6, 12, 30, 7, 17, 25, 35}
	for _, value := range values {
		bt.Insert(value)
	}

	leafCount := bt.GetLeafCount()
	if leafCount <= 0 {
		t.Errorf("Expected positive leaf count, got %d", leafCount)
	}
}

func TestGetAllKeys(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		bt.Insert(value)
	}

	allKeys := bt.GetAllKeys()
	inOrder := bt.InOrderTraversal()

	if len(allKeys) != len(inOrder) {
		t.Errorf("GetAllKeys and InOrderTraversal should return same length")
	}

	for i, key := range allKeys {
		if key != inOrder[i] {
			t.Errorf("GetAllKeys should match InOrderTraversal at position %d", i)
		}
	}
}

func TestClear(t *testing.T) {
	bt := NewBTree(3)
	values := []int{10, 20, 5, 6, 12, 30}

	for _, value := range values {
		bt.Insert(value)
	}

	bt.Clear()

	if bt.Root != nil {
		t.Error("Expected nil root after clear")
	}

	if bt.Size != 0 {
		t.Errorf("Expected size 0 after clear, got %d", bt.Size)
	}

	if !bt.IsEmpty() {
		t.Error("Expected tree to be empty after clear")
	}
}

func TestValidate(t *testing.T) {
	bt := NewBTree(3)

	if !bt.Validate() {
		t.Error("Empty tree should be valid")
	}

	values := []int{10, 20, 5, 6, 12, 30, 7, 17, 25, 35, 40, 50}
	for _, value := range values {
		bt.Insert(value)
		if !bt.Validate() {
			t.Errorf("Tree should be valid after inserting %d", value)
		}
	}

	toDelete := []int{12, 25, 40}
	for _, value := range toDelete {
		bt.Delete(value)
		if !bt.Validate() {
			t.Errorf("Tree should be valid after deleting %d", value)
		}
	}
}

func TestDifferentMinimumDegrees(t *testing.T) {
	degrees := []int{2, 3, 4, 5, 10}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, degree := range degrees {
		bt := NewBTree(degree)

		for _, value := range values {
			bt.Insert(value)
		}

		if bt.GetMinimumDegree() != degree {
			t.Errorf("Expected minimum degree %d, got %d", degree, bt.GetMinimumDegree())
		}

		if !bt.Validate() {
			t.Errorf("Tree with degree %d should be valid", degree)
		}

		if bt.Size != len(values) {
			t.Errorf("Expected size %d for degree %d tree, got %d", len(values), degree, bt.Size)
		}

		for _, value := range values {
			if !bt.Search(value) {
				t.Errorf("Should find value %d in degree %d tree", value, degree)
			}
		}
	}
}

func TestLargeDataset(t *testing.T) {
	bt := NewBTree(5)
	n := 1000

	values := make([]int, n)
	for i := range n {
		values[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	for _, value := range values {
		bt.Insert(value)
	}

	if bt.Size != n {
		t.Errorf("Expected size %d, got %d", n, bt.Size)
	}

	if !bt.Validate() {
		t.Error("Large tree should be valid")
	}

	for i := 1; i <= n; i++ {
		if !bt.Search(i) {
			t.Errorf("Should find value %d in large tree", i)
		}
	}

	inOrder := bt.InOrderTraversal()
	if len(inOrder) != n {
		t.Errorf("Expected %d elements in traversal, got %d", n, len(inOrder))
	}

	for i := range n {
		if inOrder[i] != i+1 {
			t.Errorf("Expected %d at position %d in sorted order, got %d", i+1, i, inOrder[i])
		}
	}
}

func TestComplexDeletions(t *testing.T) {
	bt := NewBTree(3)
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, value := range values {
		bt.Insert(value)
	}

	toDelete := []int{10, 5, 15, 3, 7, 12, 18, 1, 20, 8}

	for _, value := range toDelete {
		initialSize := bt.Size
		if !bt.Delete(value) {
			t.Errorf("Expected successful deletion of %d", value)
		}

		if bt.Size != initialSize-1 {
			t.Errorf("Expected size to decrease by 1 after deleting %d", value)
		}

		if bt.Search(value) {
			t.Errorf("Should not find deleted value %d", value)
		}

		if !bt.Validate() {
			t.Errorf("Tree should be valid after deleting %d", value)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	bt := NewBTree(100)
	b.ResetTimer()

	for i := range b.N {
		bt.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bt := NewBTree(100)
	for i := range 10000 {
		bt.Insert(i)
	}

	b.ResetTimer()
	for i := range b.N {
		bt.Search(i % 10000)
	}
}

func BenchmarkSearchIterative(b *testing.B) {
	bt := NewBTree(100)
	for i := range 10000 {
		bt.Insert(i)
	}

	b.ResetTimer()
	for i := range b.N {
		bt.SearchIterative(i % 10000)
	}
}

func BenchmarkDelete(b *testing.B) {
	values := make([]int, b.N)
	for i := range b.N {
		values[i] = i
	}

	b.ResetTimer()
	for i := range b.N {
		bt := NewBTree(100)
		for j := range 1000 {
			bt.Insert(j)
		}
		bt.Delete(values[i] % 1000)
	}
}

func BenchmarkInOrderTraversal(b *testing.B) {
	bt := NewBTree(100)
	for i := range 10000 {
		bt.Insert(i)
	}

	for b.Loop() {
		bt.InOrderTraversal()
	}
}
