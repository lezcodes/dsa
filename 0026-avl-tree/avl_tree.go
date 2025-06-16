package avl_tree

import (
	"fmt"
	"math"
)

type Node struct {
	Value  int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
	Size int
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		Root: nil,
		Size: 0,
	}
}

func (avl *AVLTree) getHeight(node *Node) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func (avl *AVLTree) updateHeight(node *Node) {
	if node != nil {
		leftHeight := avl.getHeight(node.Left)
		rightHeight := avl.getHeight(node.Right)
		if leftHeight > rightHeight {
			node.Height = leftHeight + 1
		} else {
			node.Height = rightHeight + 1
		}
	}
}

func (avl *AVLTree) getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return avl.getHeight(node.Left) - avl.getHeight(node.Right)
}

func (avl *AVLTree) rotateRight(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	avl.updateHeight(y)
	avl.updateHeight(x)

	return x
}

func (avl *AVLTree) rotateLeft(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	avl.updateHeight(x)
	avl.updateHeight(y)

	return y
}

func (avl *AVLTree) Insert(value int) {
	avl.Root = avl.insertRecursive(avl.Root, value)
}

func (avl *AVLTree) insertRecursive(node *Node, value int) *Node {
	if node == nil {
		avl.Size++
		return &Node{Value: value, Height: 0}
	}

	if value < node.Value {
		node.Left = avl.insertRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = avl.insertRecursive(node.Right, value)
	} else {
		return node
	}

	avl.updateHeight(node)

	balance := avl.getBalance(node)

	if balance > 1 && value < node.Left.Value {
		return avl.rotateRight(node)
	}

	if balance < -1 && value > node.Right.Value {
		return avl.rotateLeft(node)
	}

	if balance > 1 && value > node.Left.Value {
		node.Left = avl.rotateLeft(node.Left)
		return avl.rotateRight(node)
	}

	if balance < -1 && value < node.Right.Value {
		node.Right = avl.rotateRight(node.Right)
		return avl.rotateLeft(node)
	}

	return node
}

func (avl *AVLTree) InsertIterative(value int) {
	if avl.Root == nil {
		avl.Root = &Node{Value: value, Height: 0}
		avl.Size++
		return
	}

	stack := []*Node{}
	current := avl.Root

	for current != nil {
		stack = append(stack, current)
		if value < current.Value {
			if current.Left == nil {
				current.Left = &Node{Value: value, Height: 0}
				avl.Size++
				break
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &Node{Value: value, Height: 0}
				avl.Size++
				break
			}
			current = current.Right
		} else {
			return
		}
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		avl.updateHeight(node)
		balance := avl.getBalance(node)

		if balance > 1 {
			if avl.getBalance(node.Left) >= 0 {
				node = avl.rotateRight(node)
			} else {
				node.Left = avl.rotateLeft(node.Left)
				node = avl.rotateRight(node)
			}
		} else if balance < -1 {
			if avl.getBalance(node.Right) <= 0 {
				node = avl.rotateLeft(node)
			} else {
				node.Right = avl.rotateRight(node.Right)
				node = avl.rotateLeft(node)
			}
		}

		if len(stack) == 0 {
			avl.Root = node
		} else {
			parent := stack[len(stack)-1]
			if node.Value < parent.Value {
				parent.Left = node
			} else {
				parent.Right = node
			}
		}
	}
}

func (avl *AVLTree) Delete(value int) bool {
	initialSize := avl.Size
	avl.Root = avl.deleteRecursive(avl.Root, value)
	return avl.Size < initialSize
}

func (avl *AVLTree) deleteRecursive(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = avl.deleteRecursive(node.Left, value)
	} else if value > node.Value {
		node.Right = avl.deleteRecursive(node.Right, value)
	} else {
		avl.Size--

		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		successor := avl.findMin(node.Right)
		node.Value = successor.Value
		node.Right = avl.deleteRecursive(node.Right, successor.Value)
		avl.Size++
	}

	avl.updateHeight(node)

	balance := avl.getBalance(node)

	if balance > 1 && avl.getBalance(node.Left) >= 0 {
		return avl.rotateRight(node)
	}

	if balance > 1 && avl.getBalance(node.Left) < 0 {
		node.Left = avl.rotateLeft(node.Left)
		return avl.rotateRight(node)
	}

	if balance < -1 && avl.getBalance(node.Right) <= 0 {
		return avl.rotateLeft(node)
	}

	if balance < -1 && avl.getBalance(node.Right) > 0 {
		node.Right = avl.rotateRight(node.Right)
		return avl.rotateLeft(node)
	}

	return node
}

func (avl *AVLTree) findMin(node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (avl *AVLTree) Search(value int) bool {
	return avl.searchRecursive(avl.Root, value)
}

func (avl *AVLTree) searchRecursive(node *Node, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return avl.searchRecursive(node.Left, value)
	}
	return avl.searchRecursive(node.Right, value)
}

func (avl *AVLTree) SearchIterative(value int) bool {
	current := avl.Root
	for current != nil {
		if value == current.Value {
			return true
		}
		if value < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

func (avl *AVLTree) GetSize() int {
	return avl.Size
}

func (avl *AVLTree) IsEmpty() bool {
	return avl.Root == nil
}

func (avl *AVLTree) GetHeight() int {
	return avl.getHeight(avl.Root)
}

func (avl *AVLTree) IsBalanced() bool {
	return avl.isBalancedRecursive(avl.Root)
}

func (avl *AVLTree) isBalancedRecursive(node *Node) bool {
	if node == nil {
		return true
	}

	balance := avl.getBalance(node)
	if int(math.Abs(float64(balance))) > 1 {
		return false
	}

	return avl.isBalancedRecursive(node.Left) && avl.isBalancedRecursive(node.Right)
}

func (avl *AVLTree) InOrderTraversal() []int {
	var result []int
	avl.inOrderRecursive(avl.Root, &result)
	return result
}

func (avl *AVLTree) inOrderRecursive(node *Node, result *[]int) {
	if node != nil {
		avl.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		avl.inOrderRecursive(node.Right, result)
	}
}

func (avl *AVLTree) PreOrderTraversal() []int {
	var result []int
	avl.preOrderRecursive(avl.Root, &result)
	return result
}

func (avl *AVLTree) preOrderRecursive(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		avl.preOrderRecursive(node.Left, result)
		avl.preOrderRecursive(node.Right, result)
	}
}

func (avl *AVLTree) LevelOrderTraversal() []int {
	if avl.Root == nil {
		return []int{}
	}

	var result []int
	queue := []*Node{avl.Root}

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

func (avl *AVLTree) FindMin() (int, bool) {
	if avl.Root == nil {
		return 0, false
	}
	node := avl.findMin(avl.Root)
	return node.Value, true
}

func (avl *AVLTree) FindMax() (int, bool) {
	if avl.Root == nil {
		return 0, false
	}
	node := avl.findMax(avl.Root)
	return node.Value, true
}

func (avl *AVLTree) findMax(node *Node) *Node {
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (avl *AVLTree) Clear() {
	avl.Root = nil
	avl.Size = 0
}

func (avl *AVLTree) PrintTree() {
	avl.printTreeRecursive(avl.Root, "", true)
}

func (avl *AVLTree) printTreeRecursive(node *Node, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	fmt.Printf("%s%s%d (h:%d, b:%d)\n", prefix, connector, node.Value, node.Height, avl.getBalance(node))

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
		avl.printTreeRecursive(child, childPrefix, isLastChild)
	}
}

func Run() any {
	avl := NewAVLTree()

	values := []int{10, 20, 30, 40, 50, 25}
	for _, value := range values {
		avl.Insert(value)
	}

	result := make(map[string]any)
	result["size"] = avl.GetSize()
	result["height"] = avl.GetHeight()
	result["isBalanced"] = avl.IsBalanced()
	result["inOrder"] = avl.InOrderTraversal()
	result["preOrder"] = avl.PreOrderTraversal()
	result["levelOrder"] = avl.LevelOrderTraversal()

	result["searchExisting"] = avl.Search(30)
	result["searchNonExisting"] = avl.Search(100)

	min, hasMin := avl.FindMin()
	max, hasMax := avl.FindMax()
	result["min"] = map[string]any{"value": min, "exists": hasMin}
	result["max"] = map[string]any{"value": max, "exists": hasMax}

	result["deleteSuccess"] = avl.Delete(20)
	result["sizeAfterDelete"] = avl.GetSize()
	result["heightAfterDelete"] = avl.GetHeight()
	result["isBalancedAfterDelete"] = avl.IsBalanced()
	result["inOrderAfterDelete"] = avl.InOrderTraversal()

	avl.InsertIterative(15)
	result["sizeAfterIterativeInsert"] = avl.GetSize()
	result["searchIterative"] = avl.SearchIterative(15)

	return result
}
