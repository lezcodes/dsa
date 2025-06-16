package binary_search_tree

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
		return
	}

	if initialSize, exists := resultMap["initialSize"]; !exists || initialSize != 11 {
		t.Errorf("Expected initial size to be 11, got %v", initialSize)
	}

	if isValid, exists := resultMap["isValidBST"]; !exists || isValid != true {
		t.Errorf("Expected BST to be valid after operations, got %v", isValid)
	}
}

func TestNewBST(t *testing.T) {
	bst := NewBST()
	if bst == nil {
		t.Error("Expected non-nil BST")
	}
	if bst.Root != nil {
		t.Error("Expected empty BST root to be nil")
	}
	if bst.GetSize() != 0 {
		t.Errorf("Expected size 0, got %d", bst.GetSize())
	}
	if !bst.IsEmpty() {
		t.Error("Expected empty BST")
	}
}

func TestInsert(t *testing.T) {
	bst := NewBST()

	bst.Insert(50)
	if bst.Root == nil || bst.Root.Value != 50 {
		t.Error("Expected root to be 50")
	}
	if bst.GetSize() != 1 {
		t.Errorf("Expected size 1, got %d", bst.GetSize())
	}

	bst.Insert(30)
	if bst.Root.Left == nil || bst.Root.Left.Value != 30 {
		t.Error("Expected left child to be 30")
	}

	bst.Insert(70)
	if bst.Root.Right == nil || bst.Root.Right.Value != 70 {
		t.Error("Expected right child to be 70")
	}

	if bst.GetSize() != 3 {
		t.Errorf("Expected size 3, got %d", bst.GetSize())
	}

	bst.Insert(50)
	if bst.GetSize() != 3 {
		t.Errorf("Expected size to remain 3 after duplicate insert, got %d", bst.GetSize())
	}
}

func TestSearch(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	for _, value := range values {
		if !bst.Search(value) {
			t.Errorf("Expected to find value %d", value)
		}
	}

	nonExistentValues := []int{10, 25, 35, 45, 55, 65, 75, 90}
	for _, value := range nonExistentValues {
		if bst.Search(value) {
			t.Errorf("Expected not to find value %d", value)
		}
	}

	if bst.Search(100) {
		t.Error("Expected not to find value 100 in empty subtree")
	}
}

func TestDeleteCase1NoChild(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	initialSize := bst.GetSize()

	deleted := bst.Delete(20)
	if !deleted {
		t.Error("Expected deletion of leaf node 20 to succeed")
	}

	if bst.GetSize() != initialSize-1 {
		t.Errorf("Expected size to decrease by 1, got %d", bst.GetSize())
	}

	if bst.Search(20) {
		t.Error("Expected node 20 to be deleted")
	}

	if bst.Root.Left.Left != nil {
		t.Error("Expected left child of 30 to be nil after deleting 20")
	}

	if !bst.IsValidBST() {
		t.Error("Expected BST to remain valid after leaf deletion")
	}
}

func TestDeleteCase2OneChild(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 60, 80, 10}
	for _, value := range values {
		bst.Insert(value)
	}

	initialSize := bst.GetSize()

	deleted := bst.Delete(20)
	if !deleted {
		t.Error("Expected deletion of node 20 with one child to succeed")
	}

	if bst.GetSize() != initialSize-1 {
		t.Errorf("Expected size to decrease by 1, got %d", bst.GetSize())
	}

	if bst.Search(20) {
		t.Error("Expected node 20 to be deleted")
	}

	if bst.Root.Left.Left == nil || bst.Root.Left.Left.Value != 10 {
		t.Error("Expected left child of 30 to be 10 after deleting 20")
	}

	if !bst.IsValidBST() {
		t.Error("Expected BST to remain valid after one-child deletion")
	}
}

func TestDeleteCase3TwoChildren(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80, 35, 45}
	for _, value := range values {
		bst.Insert(value)
	}

	initialSize := bst.GetSize()

	deleted := bst.Delete(30)
	if !deleted {
		t.Error("Expected deletion of node 30 with two children to succeed")
	}

	if bst.GetSize() != initialSize-1 {
		t.Errorf("Expected size to decrease by 1, got %d", bst.GetSize())
	}

	if bst.Search(30) {
		t.Error("Expected node 30 to be deleted")
	}

	if bst.Root.Left == nil || bst.Root.Left.Value != 35 {
		t.Error("Expected left child of root to be 35 (successor of 30)")
	}

	if !bst.IsValidBST() {
		t.Error("Expected BST to remain valid after two-children deletion")
	}

	inOrder := bst.InOrderTraversal()
	for i := 1; i < len(inOrder); i++ {
		if inOrder[i-1] >= inOrder[i] {
			t.Error("Expected in-order traversal to be sorted after deletion")
		}
	}
}

func TestDeleteNonExistent(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70}
	for _, value := range values {
		bst.Insert(value)
	}

	initialSize := bst.GetSize()

	deleted := bst.Delete(100)
	if deleted {
		t.Error("Expected deletion of non-existent node to fail")
	}

	if bst.GetSize() != initialSize {
		t.Errorf("Expected size to remain %d, got %d", initialSize, bst.GetSize())
	}
}

func TestDeleteRoot(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)

	deleted := bst.Delete(50)
	if !deleted {
		t.Error("Expected deletion of root to succeed")
	}

	if bst.Root != nil {
		t.Error("Expected root to be nil after deletion")
	}

	if bst.GetSize() != 0 {
		t.Errorf("Expected size 0 after root deletion, got %d", bst.GetSize())
	}

	if !bst.IsEmpty() {
		t.Error("Expected BST to be empty after root deletion")
	}
}

func TestFindMinMax(t *testing.T) {
	bst := NewBST()

	_, hasMin := bst.FindMin()
	_, hasMax := bst.FindMax()
	if hasMin || hasMax {
		t.Error("Expected no min/max in empty BST")
	}

	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 90}
	for _, value := range values {
		bst.Insert(value)
	}

	min, hasMin := bst.FindMin()
	if !hasMin || min != 10 {
		t.Errorf("Expected min to be 10, got %d", min)
	}

	max, hasMax := bst.FindMax()
	if !hasMax || max != 90 {
		t.Errorf("Expected max to be 90, got %d", max)
	}
}

func TestHeight(t *testing.T) {
	bst := NewBST()

	if bst.Height() != -1 {
		t.Errorf("Expected height -1 for empty BST, got %d", bst.Height())
	}

	bst.Insert(50)
	if bst.Height() != 0 {
		t.Errorf("Expected height 0 for single node, got %d", bst.Height())
	}

	bst.Insert(30)
	bst.Insert(70)
	if bst.Height() != 1 {
		t.Errorf("Expected height 1, got %d", bst.Height())
	}

	bst.Insert(20)
	bst.Insert(40)
	if bst.Height() != 2 {
		t.Errorf("Expected height 2, got %d", bst.Height())
	}
}

func TestInOrderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	inOrder := bst.InOrderTraversal()
	expected := []int{20, 30, 40, 50, 60, 70, 80}

	if !reflect.DeepEqual(inOrder, expected) {
		t.Errorf("Expected in-order %v, got %v", expected, inOrder)
	}

	for i := 1; i < len(inOrder); i++ {
		if inOrder[i-1] >= inOrder[i] {
			t.Error("Expected in-order traversal to be sorted")
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40}
	for _, value := range values {
		bst.Insert(value)
	}

	preOrder := bst.PreOrderTraversal()
	expected := []int{50, 30, 20, 40, 70}

	if !reflect.DeepEqual(preOrder, expected) {
		t.Errorf("Expected pre-order %v, got %v", expected, preOrder)
	}
}

func TestPostOrderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40}
	for _, value := range values {
		bst.Insert(value)
	}

	postOrder := bst.PostOrderTraversal()
	expected := []int{20, 40, 30, 70, 50}

	if !reflect.DeepEqual(postOrder, expected) {
		t.Errorf("Expected post-order %v, got %v", expected, postOrder)
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	levelOrder := bst.LevelOrderTraversal()
	expected := []int{50, 30, 70, 20, 40, 60, 80}

	if !reflect.DeepEqual(levelOrder, expected) {
		t.Errorf("Expected level-order %v, got %v", expected, levelOrder)
	}
}

func TestIsValidBST(t *testing.T) {
	bst := NewBST()
	if !bst.IsValidBST() {
		t.Error("Expected empty BST to be valid")
	}

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	if !bst.IsValidBST() {
		t.Error("Expected properly constructed BST to be valid")
	}

	bst.Root.Left.Value = 60
	if bst.IsValidBST() {
		t.Error("Expected BST with invalid left child to be invalid")
	}
}

func TestCountNodes(t *testing.T) {
	bst := NewBST()
	if bst.CountNodes() != 0 {
		t.Errorf("Expected 0 nodes in empty BST, got %d", bst.CountNodes())
	}

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	if bst.CountNodes() != len(values) {
		t.Errorf("Expected %d nodes, got %d", len(values), bst.CountNodes())
	}

	if bst.CountNodes() != bst.GetSize() {
		t.Error("Expected CountNodes to match GetSize")
	}
}

func TestCountLeaves(t *testing.T) {
	bst := NewBST()
	if bst.CountLeaves() != 0 {
		t.Errorf("Expected 0 leaves in empty BST, got %d", bst.CountLeaves())
	}

	bst.Insert(50)
	if bst.CountLeaves() != 1 {
		t.Errorf("Expected 1 leaf for single node, got %d", bst.CountLeaves())
	}

	values := []int{30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	if bst.CountLeaves() != 4 {
		t.Errorf("Expected 4 leaves, got %d", bst.CountLeaves())
	}
}

func TestGetSuccessor(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80, 35, 45}
	for _, value := range values {
		bst.Insert(value)
	}

	successor, exists := bst.GetSuccessor(30)
	if !exists || successor != 35 {
		t.Errorf("Expected successor of 30 to be 35, got %d", successor)
	}

	successor, exists = bst.GetSuccessor(40)
	if !exists || successor != 45 {
		t.Errorf("Expected successor of 40 to be 45, got %d", successor)
	}

	successor, exists = bst.GetSuccessor(50)
	if !exists || successor != 60 {
		t.Errorf("Expected successor of 50 to be 60, got %d", successor)
	}

	_, exists = bst.GetSuccessor(80)
	if exists {
		t.Error("Expected no successor for maximum value 80")
	}

	_, exists = bst.GetSuccessor(100)
	if exists {
		t.Error("Expected no successor for non-existent value 100")
	}
}

func TestGetPredecessor(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80, 35, 45}
	for _, value := range values {
		bst.Insert(value)
	}

	predecessor, exists := bst.GetPredecessor(50)
	if !exists || predecessor != 45 {
		t.Errorf("Expected predecessor of 50 to be 45, got %d", predecessor)
	}

	predecessor, exists = bst.GetPredecessor(40)
	if !exists || predecessor != 35 {
		t.Errorf("Expected predecessor of 40 to be 35, got %d", predecessor)
	}

	predecessor, exists = bst.GetPredecessor(30)
	if !exists || predecessor != 20 {
		t.Errorf("Expected predecessor of 30 to be 20, got %d", predecessor)
	}

	_, exists = bst.GetPredecessor(20)
	if exists {
		t.Error("Expected no predecessor for minimum value 20")
	}

	_, exists = bst.GetPredecessor(100)
	if exists {
		t.Error("Expected no predecessor for non-existent value 100")
	}
}

func TestClear(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40}
	for _, value := range values {
		bst.Insert(value)
	}

	if bst.IsEmpty() {
		t.Error("Expected BST not to be empty before clear")
	}

	bst.Clear()

	if !bst.IsEmpty() {
		t.Error("Expected BST to be empty after clear")
	}

	if bst.Root != nil {
		t.Error("Expected root to be nil after clear")
	}

	if bst.GetSize() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", bst.GetSize())
	}
}

func TestComplexDeletionScenarios(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45, 55, 65, 75, 85}
	for _, value := range values {
		bst.Insert(value)
	}

	initialInOrder := bst.InOrderTraversal()

	bst.Delete(10)
	bst.Delete(25)
	bst.Delete(35)

	if !bst.IsValidBST() {
		t.Error("Expected BST to remain valid after multiple deletions")
	}

	finalInOrder := bst.InOrderTraversal()
	for i := 1; i < len(finalInOrder); i++ {
		if finalInOrder[i-1] >= finalInOrder[i] {
			t.Error("Expected in-order traversal to remain sorted after deletions")
		}
	}

	if len(finalInOrder) != len(initialInOrder)-3 {
		t.Errorf("Expected %d nodes after deletions, got %d", len(initialInOrder)-3, len(finalInOrder))
	}
}

func BenchmarkInsert(b *testing.B) {
	bst := NewBST()
	for b.Loop() {
		bst.Insert(b.N % 10000)
	}
}

func BenchmarkSearch(b *testing.B) {
	bst := NewBST()
	for i := range 1000 {
		bst.Insert(i)
	}

	for b.Loop() {
		bst.Search(b.N % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	values := make([]int, b.N)
	for i := range b.N {
		values[i] = i
	}

	b.ResetTimer()
	for b.Loop() {
		bst := NewBST()
		for _, value := range values {
			bst.Insert(value)
		}
		bst.Delete(b.N / 2)
	}
}

func BenchmarkInOrderTraversal(b *testing.B) {
	bst := NewBST()
	for i := range 1000 {
		bst.Insert(i)
	}

	for b.Loop() {
		bst.InOrderTraversal()
	}
}

func BenchmarkHeight(b *testing.B) {
	bst := NewBST()
	for i := range 1000 {
		bst.Insert(i)
	}

	for b.Loop() {
		bst.Height()
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
