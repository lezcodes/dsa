package binary_search_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{Root: nil}
}

func (bst *BST) Insert(value int) {
	bst.Root = bst.insertRecursive(bst.Root, value)
}

func (bst *BST) insertRecursive(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = bst.insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertRecursive(node.Right, value)
	}

	return node
}

func (bst *BST) Search(value int) bool {
	return bst.searchRecursive(bst.Root, value)
}

func (bst *BST) searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return bst.searchRecursive(node.Left, value)
	}
	return bst.searchRecursive(node.Right, value)
}

func (bst *BST) Delete(value int) {
	bst.Root = bst.deleteRecursive(bst.Root, value)
}

func (bst *BST) deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = bst.deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.deleteRecursive(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		successor := bst.findMin(node.Right)
		node.Value = successor.Value
		node.Right = bst.deleteRecursive(node.Right, successor.Value)
	}

	return node
}

func (bst *BST) findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (bst *BST) InOrderTraversal() []int {
	var result []int
	bst.inOrderRecursive(bst.Root, &result)
	return result
}

func (bst *BST) inOrderRecursive(node *Node, result *[]int) {
	if node != nil {
		bst.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		bst.inOrderRecursive(node.Right, result)
	}
}

func Run() any {
	bst := NewBST()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		bst.Insert(v)
	}

	inOrder := bst.InOrderTraversal()
	found := bst.Search(40)
	bst.Delete(30)
	afterDelete := bst.InOrderTraversal()

	return map[string]any{
		"inserted":     values,
		"inorder":      inOrder,
		"found_40":     found,
		"after_delete": afterDelete,
	}
}
