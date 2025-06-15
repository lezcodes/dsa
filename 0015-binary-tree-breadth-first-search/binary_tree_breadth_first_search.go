package binary_tree_breadth_first_search

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.items) == 0 {
		var zero T
		return zero
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Peek(index int) T {
	if index < 0 || index >= len(q.items) {
		var zero T
		return zero
	}
	return q.items[index]
}

func BFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		result = append(result, current.Value)

		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}

	return result
}

func BFSLevels(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		levelSize := queue.Size()
		currentLevel := []int{}

		for range levelSize {
			current := queue.Dequeue()
			currentLevel = append(currentLevel, current.Value)

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

func BFSRightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		levelSize := queue.Size()

		for i := range levelSize {
			current := queue.Dequeue()

			if i == levelSize-1 {
				result = append(result, current.Value)
			}

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}
	}

	return result
}

func BFSLeftSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for !queue.IsEmpty() {
		levelSize := queue.Size()

		for i := range levelSize {
			current := queue.Dequeue()

			if i == 0 {
				result = append(result, current.Value)
			}

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}
	}

	return result
}

func BFSZigzag(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)
	leftToRight := true

	for !queue.IsEmpty() {
		levelSize := queue.Size()
		currentLevel := make([]int, levelSize)

		for i := range levelSize {
			current := queue.Dequeue()

			index := i
			if !leftToRight {
				index = levelSize - 1 - i
			}
			currentLevel[index] = current.Value

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}

		result = append(result, currentLevel)
		leftToRight = !leftToRight
	}

	return result
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)
	depth := 0

	for !queue.IsEmpty() {
		levelSize := queue.Size()
		depth++

		for range levelSize {
			current := queue.Dequeue()

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}
	}

	return depth
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)
	depth := 1

	for !queue.IsEmpty() {
		levelSize := queue.Size()

		for range levelSize {
			current := queue.Dequeue()

			if current.Left == nil && current.Right == nil {
				return depth
			}

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}

		depth++
	}

	return depth
}

func LevelSum(root *TreeNode, level int) int {
	if root == nil || level < 0 {
		return 0
	}

	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)
	currentLevel := 0

	for !queue.IsEmpty() {
		levelSize := queue.Size()

		if currentLevel == level {
			sum := 0
			for i := range levelSize {
				current := queue.Peek(i)
				sum += current.Value
			}
			return sum
		}

		for range levelSize {
			current := queue.Dequeue()

			if current.Left != nil {
				queue.Enqueue(current.Left)
			}
			if current.Right != nil {
				queue.Enqueue(current.Right)
			}
		}

		currentLevel++
	}

	return 0
}

type BFSResults struct {
	Traversal    []int
	Levels       [][]int
	RightView    []int
	LeftView     []int
	ZigzagLevels [][]int
	MaxDepth     int
	MinDepth     int
	Level1Sum    int
	Level2Sum    int
}

func Run() any {
	root := NewTreeNode(3)
	root.Left = NewTreeNode(9)
	root.Right = NewTreeNode(20)
	root.Right.Left = NewTreeNode(15)
	root.Right.Right = NewTreeNode(7)
	root.Left.Left = NewTreeNode(1)
	root.Left.Right = NewTreeNode(2)

	return BFSResults{
		Traversal:    BFS(root),
		Levels:       BFSLevels(root),
		RightView:    BFSRightSideView(root),
		LeftView:     BFSLeftSideView(root),
		ZigzagLevels: BFSZigzag(root),
		MaxDepth:     MaxDepth(root),
		MinDepth:     MinDepth(root),
		Level1Sum:    LevelSum(root, 1),
		Level2Sum:    LevelSum(root, 2),
	}
}
