package binary_tree_depth_first_search

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func NewBST() *BST {
	return &BST{}
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func (bst *BST) Insert(value int) {
	bst.Root = bst.insertNode(bst.Root, value)
}

func (bst *BST) insertNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return NewTreeNode(value)
	}

	if value < node.Value {
		node.Left = bst.insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertNode(node.Right, value)
	}

	return node
}

func (bst *BST) Delete(value int) bool {
	var deleted bool
	bst.Root, deleted = bst.deleteNode(bst.Root, value)
	return deleted
}

func (bst *BST) deleteNode(node *TreeNode, value int) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}

	var deleted bool

	if value < node.Value {
		node.Left, deleted = bst.deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right, deleted = bst.deleteNode(node.Right, value)
	} else {
		deleted = true

		if node.Left == nil {
			return node.Right, deleted
		}
		if node.Right == nil {
			return node.Left, deleted
		}

		minNode := bst.findMin(node.Right)
		node.Value = minNode.Value
		node.Right, _ = bst.deleteNode(node.Right, minNode.Value)
	}

	return node, deleted
}

func (bst *BST) findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *BST) Find(value int) *TreeNode {
	return bst.findNode(bst.Root, value)
}

func (bst *BST) findNode(node *TreeNode, value int) *TreeNode {
	if node == nil || node.Value == value {
		return node
	}

	if value < node.Value {
		return bst.findNode(node.Left, value)
	}
	return bst.findNode(node.Right, value)
}

func (bst *BST) InOrderTraversal() []int {
	var result []int
	bst.inOrder(bst.Root, &result)
	return result
}

func (bst *BST) inOrder(node *TreeNode, result *[]int) {
	if node != nil {
		bst.inOrder(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrder(node.Right, result)
	}
}

func (bst *BST) PreOrderTraversal() []int {
	var result []int
	bst.preOrder(bst.Root, &result)
	return result
}

func (bst *BST) preOrder(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		bst.preOrder(node.Left, result)
		bst.preOrder(node.Right, result)
	}
}

func (bst *BST) PostOrderTraversal() []int {
	var result []int
	bst.postOrder(bst.Root, &result)
	return result
}

func (bst *BST) postOrder(node *TreeNode, result *[]int) {
	if node != nil {
		bst.postOrder(node.Left, result)
		bst.postOrder(node.Right, result)
		*result = append(*result, node.Value)
	}
}

func (bst *BST) Height() int {
	return bst.height(bst.Root)
}

func (bst *BST) height(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bst.height(node.Left)
	rightHeight := bst.height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bst *BST) Size() int {
	return bst.size(bst.Root)
}

func (bst *BST) size(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + bst.size(node.Left) + bst.size(node.Right)
}

func Run() any {
	bst := NewBST()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		bst.Insert(v)
	}

	return map[string]any{
		"inserted_values": values,
		"in_order":        bst.InOrderTraversal(),
		"pre_order":       bst.PreOrderTraversal(),
		"post_order":      bst.PostOrderTraversal(),
		"find_30":         bst.Find(30) != nil,
		"find_90":         bst.Find(90) != nil,
		"height":          bst.Height(),
		"size":            bst.Size(),
		"deleted_40":      bst.Delete(40),
		"after_delete":    bst.InOrderTraversal(),
	}
}
