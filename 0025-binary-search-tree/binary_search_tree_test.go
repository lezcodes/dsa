package binary_search_tree

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
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
}

func TestInsert(t *testing.T) {
	bst := NewBST()

	bst.Insert(50)
	if bst.Root == nil || bst.Root.Value != 50 {
		t.Error("Expected root to be 50")
	}

	bst.Insert(30)
	if bst.Root.Left == nil || bst.Root.Left.Value != 30 {
		t.Error("Expected left child to be 30")
	}

	bst.Insert(70)
	if bst.Root.Right == nil || bst.Root.Right.Value != 70 {
		t.Error("Expected right child to be 70")
	}

	bst.Insert(50)
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

	bst.Delete(20)

	if bst.Search(20) {
		t.Error("Expected node 20 to be deleted")
	}

	if bst.Root.Left.Left != nil {
		t.Error("Expected left child of 30 to be nil after deleting 20")
	}
}

func TestDeleteCase2OneChild(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 60, 80, 10}
	for _, value := range values {
		bst.Insert(value)
	}

	bst.Delete(20)

	if bst.Search(20) {
		t.Error("Expected node 20 to be deleted")
	}

	if bst.Root.Left.Left == nil || bst.Root.Left.Left.Value != 10 {
		t.Error("Expected left child of 30 to be 10 after deleting 20")
	}
}

func TestDeleteCase3TwoChildren(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80, 35, 45}
	for _, value := range values {
		bst.Insert(value)
	}

	bst.Delete(30)

	if bst.Search(30) {
		t.Error("Expected node 30 to be deleted")
	}

	if bst.Root.Left == nil || bst.Root.Left.Value != 35 {
		t.Error("Expected left child of root to be 35 (successor of 30)")
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

	bst.Delete(100)
}

func TestDeleteRoot(t *testing.T) {
	bst := NewBST()
	bst.Insert(50)

	bst.Delete(50)

	if bst.Root != nil {
		t.Error("Expected root to be nil after deletion")
	}
}

func TestInOrderTraversal(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, value := range values {
		bst.Insert(value)
	}

	result := bst.InOrderTraversal()
	expected := []int{20, 30, 40, 50, 60, 70, 80}

	if len(result) != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), len(result))
	}

	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Expected element %d to be %d, got %d", i, expected[i], val)
		}
	}
}
