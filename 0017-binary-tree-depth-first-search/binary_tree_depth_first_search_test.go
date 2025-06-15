package binary_tree_depth_first_search

import (
	"reflect"
	"testing"
)

func TestBSTInsert(t *testing.T) {
	bst := NewBST()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		bst.Insert(v)
	}

	inOrder := bst.InOrderTraversal()
	expected := []int{20, 30, 40, 50, 60, 70, 80}

	if !reflect.DeepEqual(inOrder, expected) {
		t.Errorf("Expected %v, got %v", expected, inOrder)
	}
}

func TestBSTFind(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40}

	for _, v := range values {
		bst.Insert(v)
	}

	tests := []struct {
		value    int
		expected bool
	}{
		{50, true},
		{30, true},
		{70, true},
		{20, true},
		{40, true},
		{10, false},
		{90, false},
		{35, false},
	}

	for _, tt := range tests {
		result := bst.Find(tt.value) != nil
		if result != tt.expected {
			t.Errorf("Find(%d) != nil = %v, want %v", tt.value, result, tt.expected)
		}
	}

	node := bst.Find(30)
	if node == nil || node.Value != 30 {
		t.Errorf("Find(30) should return node with value 30")
	}

	node = bst.Find(100)
	if node != nil {
		t.Errorf("Find(100) should return nil")
	}
}

func TestBSTDelete(t *testing.T) {
	tests := []struct {
		name           string
		initialValues  []int
		deleteValue    int
		expectedResult bool
		expectedOrder  []int
	}{
		{
			name:           "Delete leaf node",
			initialValues:  []int{50, 30, 70, 20, 40},
			deleteValue:    20,
			expectedResult: true,
			expectedOrder:  []int{30, 40, 50, 70},
		},
		{
			name:           "Delete node with one child",
			initialValues:  []int{50, 30, 70, 20},
			deleteValue:    30,
			expectedResult: true,
			expectedOrder:  []int{20, 50, 70},
		},
		{
			name:           "Delete node with two children",
			initialValues:  []int{50, 30, 70, 20, 40, 60, 80},
			deleteValue:    50,
			expectedResult: true,
			expectedOrder:  []int{20, 30, 40, 60, 70, 80},
		},
		{
			name:           "Delete non-existent node",
			initialValues:  []int{50, 30, 70},
			deleteValue:    100,
			expectedResult: false,
			expectedOrder:  []int{30, 50, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.initialValues {
				bst.Insert(v)
			}

			result := bst.Delete(tt.deleteValue)
			if result != tt.expectedResult {
				t.Errorf("Delete(%d) = %v, want %v", tt.deleteValue, result, tt.expectedResult)
			}

			inOrder := bst.InOrderTraversal()
			if !reflect.DeepEqual(inOrder, tt.expectedOrder) {
				t.Errorf("After delete, expected %v, got %v", tt.expectedOrder, inOrder)
			}
		})
	}
}

func TestBSTTraversals(t *testing.T) {
	bst := NewBST()
	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range values {
		bst.Insert(v)
	}

	inOrder := bst.InOrderTraversal()
	expectedInOrder := []int{20, 30, 40, 50, 60, 70, 80}
	if !reflect.DeepEqual(inOrder, expectedInOrder) {
		t.Errorf("InOrder: expected %v, got %v", expectedInOrder, inOrder)
	}

	preOrder := bst.PreOrderTraversal()
	expectedPreOrder := []int{50, 30, 20, 40, 70, 60, 80}
	if !reflect.DeepEqual(preOrder, expectedPreOrder) {
		t.Errorf("PreOrder: expected %v, got %v", expectedPreOrder, preOrder)
	}

	postOrder := bst.PostOrderTraversal()
	expectedPostOrder := []int{20, 40, 30, 60, 80, 70, 50}
	if !reflect.DeepEqual(postOrder, expectedPostOrder) {
		t.Errorf("PostOrder: expected %v, got %v", expectedPostOrder, postOrder)
	}
}

func TestBSTHeight(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected int
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: -1,
		},
		{
			name:     "Single node",
			values:   []int{50},
			expected: 0,
		},
		{
			name:     "Balanced tree",
			values:   []int{50, 30, 70, 20, 40, 60, 80},
			expected: 2,
		},
		{
			name:     "Skewed tree",
			values:   []int{10, 20, 30, 40, 50},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := NewBST()
			for _, v := range tt.values {
				bst.Insert(v)
			}

			height := bst.Height()
			if height != tt.expected {
				t.Errorf("Height() = %d, want %d", height, tt.expected)
			}
		})
	}
}

func TestBSTSize(t *testing.T) {
	bst := NewBST()

	if bst.Size() != 0 {
		t.Errorf("Empty tree size should be 0, got %d", bst.Size())
	}

	values := []int{50, 30, 70, 20, 40}
	for i, v := range values {
		bst.Insert(v)
		expectedSize := i + 1
		if bst.Size() != expectedSize {
			t.Errorf("After inserting %d values, size should be %d, got %d", expectedSize, expectedSize, bst.Size())
		}
	}
}

func TestEmptyBST(t *testing.T) {
	bst := NewBST()

	if bst.Find(10) != nil {
		t.Error("Empty tree should not contain any values")
	}

	if bst.Find(10) != nil {
		t.Error("Find on empty tree should return nil")
	}

	if bst.Delete(10) {
		t.Error("Delete on empty tree should return false")
	}

	if len(bst.InOrderTraversal()) != 0 {
		t.Error("Empty tree traversal should return empty slice")
	}
}

func TestRun(t *testing.T) {
	result := Run()
	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("Expected map[string]interface{}, got %T", result)
	}

	if resultMap["find_30"] != true {
		t.Error("Expected find_30 to be true")
	}

	if resultMap["find_90"] != false {
		t.Error("Expected find_90 to be false")
	}

	if resultMap["deleted_40"] != true {
		t.Error("Expected deleted_40 to be true")
	}
}

func BenchmarkBSTInsert(b *testing.B) {
	bst := NewBST()
	for i := 0; b.Loop(); i++ {
		bst.Insert(i)
	}
}

func BenchmarkBSTFind(b *testing.B) {
	bst := NewBST()
	for i := range 1000 {
		bst.Insert(i)
	}

	for i := 0; b.Loop(); i++ {
		bst.Find(i % 1000)
	}
}

func BenchmarkBSTDelete(b *testing.B) {
	for b.Loop() {
		b.StopTimer()
		bst := NewBST()
		for j := range 1000 {
			bst.Insert(j)
		}
		b.StartTimer()

		bst.Delete(500)
	}
}
