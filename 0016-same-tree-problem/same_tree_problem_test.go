package same_tree_problem

import "testing"

func TestIsSameTree(t *testing.T) {
	tree1 := NewTreeNode(1)
	tree1.Left = NewTreeNode(2)
	tree1.Right = NewTreeNode(3)

	tree2 := NewTreeNode(1)
	tree2.Left = NewTreeNode(2)
	tree2.Right = NewTreeNode(3)

	tree3 := NewTreeNode(1)
	tree3.Left = NewTreeNode(2)
	tree3.Right = NewTreeNode(4)

	tests := []struct {
		name     string
		tree1    *TreeNode
		tree2    *TreeNode
		expected bool
	}{
		{"Identical trees", tree1, tree2, true},
		{"Different values", tree1, tree3, false},
		{"Both nil", nil, nil, true},
		{"One nil", tree1, nil, false},
		{"Single node same", NewTreeNode(42), NewTreeNode(42), true},
		{"Single node different", NewTreeNode(42), NewTreeNode(24), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameTree(tt.tree1, tt.tree2)
			if result != tt.expected {
				t.Errorf("IsSameTree() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRun(t *testing.T) {
	result := Run()
	resultMap, ok := result.(map[string]bool)
	if !ok {
		t.Fatalf("Expected map[string]bool, got %T", result)
	}

	if !resultMap["same_trees"] {
		t.Error("Expected same_trees to be true")
	}
	if resultMap["different_trees"] {
		t.Error("Expected different_trees to be false")
	}
}

func BenchmarkIsSameTree(b *testing.B) {
	tree1 := NewTreeNode(1)
	tree1.Left = NewTreeNode(2)
	tree1.Right = NewTreeNode(3)

	tree2 := NewTreeNode(1)
	tree2.Left = NewTreeNode(2)
	tree2.Right = NewTreeNode(3)

	for b.Loop() {
		IsSameTree(tree1, tree2)
	}
}
