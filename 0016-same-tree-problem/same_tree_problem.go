package same_tree_problem

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func IsSameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return IsSameTree(a.Left, b.Left) && IsSameTree(a.Right, b.Right)
}

func Run() any {
	tree1 := NewTreeNode(1)
	tree1.Left = NewTreeNode(2)
	tree1.Right = NewTreeNode(3)

	tree2 := NewTreeNode(1)
	tree2.Left = NewTreeNode(2)
	tree2.Right = NewTreeNode(3)

	tree3 := NewTreeNode(1)
	tree3.Left = NewTreeNode(2)
	tree3.Right = NewTreeNode(4)

	return map[string]bool{
		"same_trees":      IsSameTree(tree1, tree2),
		"different_trees": IsSameTree(tree1, tree3),
	}
}
