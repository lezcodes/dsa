package binary_tree_breadth_first_search

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func BFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}

func BFSLevels(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for range levelSize {
			current := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, current.Value)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := range levelSize {
			current := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, current.Value)
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := range levelSize {
			current := queue[0]
			queue = queue[1:]

			if i == 0 {
				result = append(result, current.Value)
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, levelSize)

		for i := range levelSize {
			current := queue[0]
			queue = queue[1:]

			index := i
			if !leftToRight {
				index = levelSize - 1 - i
			}
			currentLevel[index] = current.Value

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		for range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return depth
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current.Left == nil && current.Right == nil {
				return depth
			}

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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

	queue := []*TreeNode{root}
	currentLevel := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		if currentLevel == level {
			sum := 0
			for i := range levelSize {
				current := queue[i]
				sum += current.Value
			}
			return sum
		}

		for range levelSize {
			current := queue[0]
			queue = queue[1:]

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
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
