package binary_tree_breadth_first_search

import (
	"reflect"
	"testing"
)

func createTestTree() *TreeNode {
	root := NewTreeNode(3)
	root.Left = NewTreeNode(9)
	root.Right = NewTreeNode(20)
	root.Right.Left = NewTreeNode(15)
	root.Right.Right = NewTreeNode(7)
	return root
}

func createCompleteTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)
	root.Right.Left = NewTreeNode(6)
	root.Right.Right = NewTreeNode(7)
	return root
}

func createUnbalancedTree() *TreeNode {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Left = NewTreeNode(3)
	root.Left.Left.Left = NewTreeNode(4)
	return root
}

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: []int{3, 9, 20, 15, 7},
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
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
			result := BFS(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFS() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBFSLevels(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
			expected: [][]int{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BFSLevels(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFSLevels() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBFSRightSideView(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: []int{3, 20, 7},
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: []int{1, 3, 7},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
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
			result := BFSRightSideView(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFSRightSideView() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBFSLeftSideView(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: []int{3, 9, 15},
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: []int{1, 2, 4},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
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
			result := BFSLeftSideView(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFSLeftSideView() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBFSZigzag(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: [][]int{{1}, {3, 2}, {4, 5, 6, 7}},
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: [][]int{},
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
			expected: [][]int{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BFSZigzag(tt.root)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFSZigzag() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: 3,
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: 3,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
			expected: 1,
		},
		{
			name:     "Unbalanced tree",
			root:     createUnbalancedTree(),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxDepth(tt.root)
			if result != tt.expected {
				t.Errorf("MaxDepth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Test tree",
			root:     createTestTree(),
			expected: 2,
		},
		{
			name:     "Complete tree",
			root:     createCompleteTree(),
			expected: 3,
		},
		{
			name:     "Empty tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Single node",
			root:     NewTreeNode(42),
			expected: 1,
		},
		{
			name:     "Unbalanced tree",
			root:     createUnbalancedTree(),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinDepth(tt.root)
			if result != tt.expected {
				t.Errorf("MinDepth() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLevelSum(t *testing.T) {
	tree := createTestTree()

	tests := []struct {
		name     string
		level    int
		expected int
	}{
		{
			name:     "Level 0 (root)",
			level:    0,
			expected: 3,
		},
		{
			name:     "Level 1",
			level:    1,
			expected: 29,
		},
		{
			name:     "Level 2",
			level:    2,
			expected: 22,
		},
		{
			name:     "Level 3 (non-existent)",
			level:    3,
			expected: 0,
		},
		{
			name:     "Negative level",
			level:    -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LevelSum(tree, tt.level)
			if result != tt.expected {
				t.Errorf("LevelSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLevelSumEmptyTree(t *testing.T) {
	result := LevelSum(nil, 0)
	if result != 0 {
		t.Errorf("LevelSum(nil, 0) = %v, want 0", result)
	}
}

func TestRun(t *testing.T) {
	result := Run()
	bfsResults, ok := result.(BFSResults)
	if !ok {
		t.Fatalf("Expected BFSResults, got %T", result)
	}

	if len(bfsResults.Traversal) == 0 {
		t.Error("Expected non-empty traversal")
	}
	if len(bfsResults.Levels) == 0 {
		t.Error("Expected non-empty levels")
	}
	if len(bfsResults.RightView) == 0 {
		t.Error("Expected non-empty right view")
	}
	if len(bfsResults.LeftView) == 0 {
		t.Error("Expected non-empty left view")
	}
	if len(bfsResults.ZigzagLevels) == 0 {
		t.Error("Expected non-empty zigzag levels")
	}
	if bfsResults.MaxDepth <= 0 {
		t.Error("Expected positive max depth")
	}
	if bfsResults.MinDepth <= 0 {
		t.Error("Expected positive min depth")
	}

	expectedTraversal := []int{3, 9, 20, 1, 2, 15, 7}
	if !reflect.DeepEqual(bfsResults.Traversal, expectedTraversal) {
		t.Errorf("Traversal = %v, want %v", bfsResults.Traversal, expectedTraversal)
	}

	expectedLevels := [][]int{{3}, {9, 20}, {1, 2, 15, 7}}
	if !reflect.DeepEqual(bfsResults.Levels, expectedLevels) {
		t.Errorf("Levels = %v, want %v", bfsResults.Levels, expectedLevels)
	}
}

func BenchmarkBFS(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		BFS(tree)
	}
}

func BenchmarkBFSLevels(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		BFSLevels(tree)
	}
}

func BenchmarkBFSRightSideView(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		BFSRightSideView(tree)
	}
}

func BenchmarkBFSZigzag(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		BFSZigzag(tree)
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		MaxDepth(tree)
	}
}

func BenchmarkMinDepth(b *testing.B) {
	tree := createCompleteTree()
	for b.Loop() {
		MinDepth(tree)
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

func BenchmarkLargeTreeBFS(b *testing.B) {
	tree := createLargeTree(10)
	for b.Loop() {
		BFS(tree)
	}
}

func BenchmarkLargeTreeBFSLevels(b *testing.B) {
	tree := createLargeTree(10)
	for b.Loop() {
		BFSLevels(tree)
	}
}
