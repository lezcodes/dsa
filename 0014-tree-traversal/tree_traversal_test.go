package tree_traversal

import (
	"reflect"
	"testing"
)

func createTestTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)
	return root
}

func createSingleNodeTree() *TreeNode {
	return NewTreeNode(42)
}

func createUnbalancedTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(3)
	root.Left.Left.Left = NewTreeNode(4)
	return root
}

func TestPreOrderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Balanced tree",
			root:     createTestTree(),
			expected: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     createSingleNodeTree(),
			expected: []int{42},
		},
		{
			name:     "Unbalanced tree",
			root:     createUnbalancedTree(),
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PreOrderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PreOrderTraversal() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInOrderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Balanced tree",
			root:     createTestTree(),
			expected: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     createSingleNodeTree(),
			expected: []int{42},
		},
		{
			name:     "Unbalanced tree",
			root:     createUnbalancedTree(),
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InOrderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InOrderTraversal() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPostOrderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Balanced tree",
			root:     createTestTree(),
			expected: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     createSingleNodeTree(),
			expected: []int{42},
		},
		{
			name:     "Unbalanced tree",
			root:     createUnbalancedTree(),
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PostOrderTraversal(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PostOrderTraversal() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIterativeVsRecursiveTraversals(t *testing.T) {
	tree := createTestTree()

	preOrderRecursive := PreOrderTraversal(tree)
	preOrderIterative := PreOrderIterative(tree)
	if !reflect.DeepEqual(preOrderRecursive, preOrderIterative) {
		t.Errorf("Pre-order mismatch: recursive %v, iterative %v", preOrderRecursive, preOrderIterative)
	}

	inOrderRecursive := InOrderTraversal(tree)
	inOrderIterative := InOrderIterative(tree)
	if !reflect.DeepEqual(inOrderRecursive, inOrderIterative) {
		t.Errorf("In-order mismatch: recursive %v, iterative %v", inOrderRecursive, inOrderIterative)
	}

	postOrderRecursive := PostOrderTraversal(tree)
	postOrderIterative := PostOrderIterative(tree)
	if !reflect.DeepEqual(postOrderRecursive, postOrderIterative) {
		t.Errorf("Post-order mismatch: recursive %v, iterative %v", postOrderRecursive, postOrderIterative)
	}
}

func TestIterativeTraversalsEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
	}{
		{"Empty tree", nil},
		{"Single node", createSingleNodeTree()},
		{"Unbalanced tree", createUnbalancedTree()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preOrderRecursive := PreOrderTraversal(tt.root)
			preOrderIterative := PreOrderIterative(tt.root)
			if !reflect.DeepEqual(preOrderRecursive, preOrderIterative) {
				t.Errorf("Pre-order mismatch for %s: recursive %v, iterative %v",
					tt.name, preOrderRecursive, preOrderIterative)
			}

			inOrderRecursive := InOrderTraversal(tt.root)
			inOrderIterative := InOrderIterative(tt.root)
			if !reflect.DeepEqual(inOrderRecursive, inOrderIterative) {
				t.Errorf("In-order mismatch for %s: recursive %v, iterative %v",
					tt.name, inOrderRecursive, inOrderIterative)
			}

			postOrderRecursive := PostOrderTraversal(tt.root)
			postOrderIterative := PostOrderIterative(tt.root)
			if !reflect.DeepEqual(postOrderRecursive, postOrderIterative) {
				t.Errorf("Post-order mismatch for %s: recursive %v, iterative %v",
					tt.name, postOrderRecursive, postOrderIterative)
			}
		})
	}
}

func TestRun(t *testing.T) {
	result := Run()
	traversalResults, ok := result.(TraversalResults)
	if !ok {
		t.Fatalf("Expected TraversalResults, got %T", result)
	}

	expectedPreOrder := []int{1, 2, 4, 5, 3, 6, 7}
	expectedInOrder := []int{4, 2, 5, 1, 6, 3, 7}
	expectedPostOrder := []int{4, 5, 2, 6, 7, 3, 1}

	if !reflect.DeepEqual(traversalResults.PreOrder, expectedPreOrder) {
		t.Errorf("PreOrder = %v, want %v", traversalResults.PreOrder, expectedPreOrder)
	}
	if !reflect.DeepEqual(traversalResults.InOrder, expectedInOrder) {
		t.Errorf("InOrder = %v, want %v", traversalResults.InOrder, expectedInOrder)
	}
	if !reflect.DeepEqual(traversalResults.PostOrder, expectedPostOrder) {
		t.Errorf("PostOrder = %v, want %v", traversalResults.PostOrder, expectedPostOrder)
	}
}

func BenchmarkPreOrderRecursive(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		PreOrderTraversal(tree)
	}
}

func BenchmarkPreOrderIterative(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		PreOrderIterative(tree)
	}
}

func BenchmarkInOrderRecursive(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		InOrderTraversal(tree)
	}
}

func BenchmarkInOrderIterative(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		InOrderIterative(tree)
	}
}

func BenchmarkPostOrderRecursive(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		PostOrderTraversal(tree)
	}
}

func BenchmarkPostOrderIterative(b *testing.B) {
	tree := createTestTree()
	for b.Loop() {
		PostOrderIterative(tree)
	}
}

func createLargeTree(depth int) *TreeNode {
	if depth <= 0 {
		return nil
	}

	root := NewTreeNode(depth)
	root.Left = createLargeTree(depth - 1)
	root.Right = createLargeTree(depth - 1)
	return root
}

func BenchmarkLargeTreePreOrder(b *testing.B) {
	tree := createLargeTree(10)
	for b.Loop() {
		PreOrderTraversal(tree)
	}
}

func BenchmarkLargeTreeInOrder(b *testing.B) {
	tree := createLargeTree(10)
	for b.Loop() {
		InOrderTraversal(tree)
	}
}

func BenchmarkLargeTreePostOrder(b *testing.B) {
	tree := createLargeTree(10)
	for b.Loop() {
		PostOrderTraversal(tree)
	}
}
