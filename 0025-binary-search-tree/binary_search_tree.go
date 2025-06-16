package binary_search_tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
	Size int
}

func NewBST() *BST {
	return &BST{
		Root: nil,
		Size: 0,
	}
}

func (bst *BST) Insert(value int) {
	bst.Root = bst.insertRecursive(bst.Root, value)
}

func (bst *BST) insertRecursive(node *Node, value int) *Node {
	if node == nil {
		bst.Size++
		return &Node{Value: value}
	}

	if value < node.Value {
		node.Left = bst.insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.insertRecursive(node.Right, value)
	}

	return node
}

func (bst *BST) Delete(value int) bool {
	initialSize := bst.Size
	bst.Root = bst.deleteRecursive(bst.Root, value)
	return bst.Size < initialSize
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
		bst.Size--

		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		successor := bst.findMin(node.Right)
		node.Value = successor.Value
		node.Right = bst.deleteRecursive(node.Right, successor.Value)
		bst.Size++
	}

	return node
}

func (bst *BST) findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
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

func (bst *BST) FindMin() (int, bool) {
	if bst.Root == nil {
		return 0, false
	}
	node := bst.findMin(bst.Root)
	return node.Value, true
}

func (bst *BST) FindMax() (int, bool) {
	if bst.Root == nil {
		return 0, false
	}
	node := bst.findMax(bst.Root)
	return node.Value, true
}

func (bst *BST) findMax(node *Node) *Node {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (bst *BST) Height() int {
	return bst.heightRecursive(bst.Root)
}

func (bst *BST) heightRecursive(node *Node) int {
	if node == nil {
		return -1
	}

	leftHeight := bst.heightRecursive(node.Left)
	rightHeight := bst.heightRecursive(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (bst *BST) IsEmpty() bool {
	return bst.Root == nil
}

func (bst *BST) GetSize() int {
	return bst.Size
}

func (bst *BST) Clear() {
	bst.Root = nil
	bst.Size = 0
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

func (bst *BST) PreOrderTraversal() []int {
	var result []int
	bst.preOrderRecursive(bst.Root, &result)
	return result
}

func (bst *BST) preOrderRecursive(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		bst.preOrderRecursive(node.Left, result)
		bst.preOrderRecursive(node.Right, result)
	}
}

func (bst *BST) PostOrderTraversal() []int {
	var result []int
	bst.postOrderRecursive(bst.Root, &result)
	return result
}

func (bst *BST) postOrderRecursive(node *Node, result *[]int) {
	if node != nil {
		bst.postOrderRecursive(node.Left, result)
		bst.postOrderRecursive(node.Right, result)
		*result = append(*result, node.Value)
	}
}

func (bst *BST) LevelOrderTraversal() []int {
	if bst.Root == nil {
		return []int{}
	}

	var result []int
	queue := []*Node{bst.Root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func (bst *BST) IsValidBST() bool {
	return bst.isValidBSTRecursive(bst.Root, nil, nil)
}

func (bst *BST) isValidBSTRecursive(node *Node, min, max *int) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
		return false
	}

	return bst.isValidBSTRecursive(node.Left, min, &node.Value) &&
		bst.isValidBSTRecursive(node.Right, &node.Value, max)
}

func (bst *BST) CountNodes() int {
	return bst.countNodesRecursive(bst.Root)
}

func (bst *BST) countNodesRecursive(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + bst.countNodesRecursive(node.Left) + bst.countNodesRecursive(node.Right)
}

func (bst *BST) CountLeaves() int {
	return bst.countLeavesRecursive(bst.Root)
}

func (bst *BST) countLeavesRecursive(node *Node) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}
	return bst.countLeavesRecursive(node.Left) + bst.countLeavesRecursive(node.Right)
}

func (bst *BST) GetSuccessor(value int) (int, bool) {
	node := bst.findNode(bst.Root, value)
	if node == nil {
		return 0, false
	}

	if node.Right != nil {
		successor := bst.findMin(node.Right)
		return successor.Value, true
	}

	var successor *Node
	current := bst.Root
	for current != nil {
		if value < current.Value {
			successor = current
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			break
		}
	}

	if successor != nil {
		return successor.Value, true
	}
	return 0, false
}

func (bst *BST) GetPredecessor(value int) (int, bool) {
	node := bst.findNode(bst.Root, value)
	if node == nil {
		return 0, false
	}

	if node.Left != nil {
		predecessor := bst.findMax(node.Left)
		return predecessor.Value, true
	}

	var predecessor *Node
	current := bst.Root
	for current != nil {
		if value > current.Value {
			predecessor = current
			current = current.Right
		} else if value < current.Value {
			current = current.Left
		} else {
			break
		}
	}

	if predecessor != nil {
		return predecessor.Value, true
	}
	return 0, false
}

func (bst *BST) findNode(node *Node, value int) *Node {
	if node == nil || node.Value == value {
		return node
	}

	if value < node.Value {
		return bst.findNode(node.Left, value)
	}
	return bst.findNode(node.Right, value)
}

func (bst *BST) PrintTree() {
	bst.printTreeRecursive(bst.Root, "", true)
}

func (bst *BST) printTreeRecursive(node *Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	fmt.Printf("%s%s%d\n", prefix, connector, node.Value)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	children := []*Node{}
	if node.Left != nil {
		children = append(children, node.Left)
	}
	if node.Right != nil {
		children = append(children, node.Right)
	}

	for i, child := range children {
		isLastChild := i == len(children)-1
		bst.printTreeRecursive(child, childPrefix, isLastChild)
	}
}

func Run() any {
	bst := NewBST()

	values := []int{50, 30, 70, 20, 40, 60, 80, 10, 25, 35, 45}
	for _, value := range values {
		bst.Insert(value)
	}

	result := make(map[string]any)
	result["initialSize"] = bst.GetSize()
	result["height"] = bst.Height()
	result["inOrder"] = bst.InOrderTraversal()

	result["searchExisting"] = bst.Search(40)
	result["searchNonExisting"] = bst.Search(100)

	min, hasMin := bst.FindMin()
	max, hasMax := bst.FindMax()
	result["min"] = map[string]any{"value": min, "exists": hasMin}
	result["max"] = map[string]any{"value": max, "exists": hasMax}

	result["deleteLeaf"] = bst.Delete(10)
	result["sizeAfterDeleteLeaf"] = bst.GetSize()

	result["deleteOneChild"] = bst.Delete(25)
	result["sizeAfterDeleteOneChild"] = bst.GetSize()

	result["deleteTwoChildren"] = bst.Delete(30)
	result["sizeAfterDeleteTwoChildren"] = bst.GetSize()

	result["inOrderAfterDeletions"] = bst.InOrderTraversal()
	result["isValidBST"] = bst.IsValidBST()

	successor, hasSuccessor := bst.GetSuccessor(50)
	result["successor50"] = map[string]any{"value": successor, "exists": hasSuccessor}

	predecessor, hasPredecessor := bst.GetPredecessor(50)
	result["predecessor50"] = map[string]any{"value": predecessor, "exists": hasPredecessor}

	result["countNodes"] = bst.CountNodes()
	result["countLeaves"] = bst.CountLeaves()
	result["levelOrder"] = bst.LevelOrderTraversal()

	return result
}
