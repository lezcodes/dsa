package tree_traversal

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.Value)
	result = append(result, PreOrderTraversal(root.Left)...)
	result = append(result, PreOrderTraversal(root.Right)...)

	return result
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, InOrderTraversal(root.Left)...)
	result = append(result, root.Value)
	result = append(result, InOrderTraversal(root.Right)...)

	return result
}

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, PostOrderTraversal(root.Left)...)
	result = append(result, PostOrderTraversal(root.Right)...)
	result = append(result, root.Value)

	return result
}

func PreOrderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, current.Value)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}

func InOrderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Value)
		current = current.Right
	}

	return result
}

func PostOrderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{}
	lastVisited := (*TreeNode)(nil)
	current := root

	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			peekNode := stack[len(stack)-1]
			if peekNode.Right != nil && lastVisited != peekNode.Right {
				current = peekNode.Right
			} else {
				result = append(result, peekNode.Value)
				lastVisited = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
	}

	return result
}

type TraversalResults struct {
	PreOrder  []int
	InOrder   []int
	PostOrder []int
}

func Run() any {
	root := NewTreeNode(7)
	root.Left = NewTreeNode(23)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(5)
	root.Left.Right = NewTreeNode(4)
	root.Right.Left = NewTreeNode(18)
	root.Right.Right = NewTreeNode(21)

	return TraversalResults{
		PreOrder:  PreOrderTraversal(root),
		InOrder:   InOrderTraversal(root),
		PostOrder: PostOrderTraversal(root),
	}
}
