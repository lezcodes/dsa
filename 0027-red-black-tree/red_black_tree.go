package red_black_tree

import (
	"fmt"
)

type Color bool

const (
	RED   Color = false
	BLACK Color = true
)

type Node struct {
	Value  int
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

type RBTree struct {
	Root *Node
	NIL  *Node
	Size int
}

func NewRBTree() *RBTree {
	nil_node := &Node{Color: BLACK}
	return &RBTree{
		Root: nil_node,
		NIL:  nil_node,
		Size: 0,
	}
}

func (rb *RBTree) rotateLeft(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != rb.NIL {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == rb.NIL {
		rb.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

func (rb *RBTree) rotateRight(y *Node) {
	x := y.Left
	y.Left = x.Right
	if x.Right != rb.NIL {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == rb.NIL {
		rb.Root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}
	x.Right = y
	y.Parent = x
}

func (rb *RBTree) Insert(value int) {
	rb.insertRecursive(value)
}

func (rb *RBTree) insertRecursive(value int) {
	node := &Node{
		Value:  value,
		Color:  RED,
		Left:   rb.NIL,
		Right:  rb.NIL,
		Parent: rb.NIL,
	}

	y := rb.NIL
	x := rb.Root

	for x != rb.NIL {
		y = x
		if node.Value < x.Value {
			x = x.Left
		} else if node.Value > x.Value {
			x = x.Right
		} else {
			return
		}
	}

	node.Parent = y
	if y == rb.NIL {
		rb.Root = node
	} else if node.Value < y.Value {
		y.Left = node
	} else {
		y.Right = node
	}

	rb.Size++
	rb.insertFixup(node)
}

func (rb *RBTree) insertFixup(z *Node) {
	for z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					rb.rotateLeft(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rb.rotateRight(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rb.rotateRight(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				rb.rotateLeft(z.Parent.Parent)
			}
		}
	}
	rb.Root.Color = BLACK
}

func (rb *RBTree) InsertIterative(value int) {
	node := &Node{
		Value:  value,
		Color:  RED,
		Left:   rb.NIL,
		Right:  rb.NIL,
		Parent: rb.NIL,
	}

	y := rb.NIL
	x := rb.Root

	for x != rb.NIL {
		y = x
		if node.Value < x.Value {
			x = x.Left
		} else if node.Value > x.Value {
			x = x.Right
		} else {
			return
		}
	}

	node.Parent = y
	if y == rb.NIL {
		rb.Root = node
	} else if node.Value < y.Value {
		y.Left = node
	} else {
		y.Right = node
	}

	rb.Size++
	rb.insertFixup(node)
}

func (rb *RBTree) Delete(value int) bool {
	node := rb.findNode(rb.Root, value)
	if node == rb.NIL {
		return false
	}

	rb.deleteNode(node)
	rb.Size--
	return true
}

func (rb *RBTree) deleteNode(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node

	if z.Left == rb.NIL {
		x = z.Right
		rb.transplant(z, z.Right)
	} else if z.Right == rb.NIL {
		x = z.Left
		rb.transplant(z, z.Left)
	} else {
		y = rb.minimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right
		if y.Parent == z {
			x.Parent = y
		} else {
			rb.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		rb.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		rb.deleteFixup(x)
	}
}

func (rb *RBTree) transplant(u, v *Node) {
	if u.Parent == rb.NIL {
		rb.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

func (rb *RBTree) deleteFixup(x *Node) {
	for x != rb.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rb.rotateLeft(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rb.rotateRight(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rb.rotateLeft(x.Parent)
				x = rb.Root
			}
		} else {
			w := x.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rb.rotateRight(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					rb.rotateLeft(w)
					w = x.Parent.Left
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				rb.rotateRight(x.Parent)
				x = rb.Root
			}
		}
	}
	x.Color = BLACK
}

func (rb *RBTree) minimum(node *Node) *Node {
	for node.Left != rb.NIL {
		node = node.Left
	}
	return node
}

func (rb *RBTree) Search(value int) bool {
	return rb.searchRecursive(rb.Root, value)
}

func (rb *RBTree) searchRecursive(node *Node, value int) bool {
	if node == rb.NIL {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return rb.searchRecursive(node.Left, value)
	}
	return rb.searchRecursive(node.Right, value)
}

func (rb *RBTree) SearchIterative(value int) bool {
	current := rb.Root
	for current != rb.NIL {
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

func (rb *RBTree) findNode(node *Node, value int) *Node {
	if node == rb.NIL || value == node.Value {
		return node
	}

	if value < node.Value {
		return rb.findNode(node.Left, value)
	}
	return rb.findNode(node.Right, value)
}

func (rb *RBTree) GetSize() int {
	return rb.Size
}

func (rb *RBTree) IsEmpty() bool {
	return rb.Root == rb.NIL
}

func (rb *RBTree) GetHeight() int {
	return rb.getHeightRecursive(rb.Root)
}

func (rb *RBTree) getHeightRecursive(node *Node) int {
	if node == rb.NIL {
		return -1
	}

	leftHeight := rb.getHeightRecursive(node.Left)
	rightHeight := rb.getHeightRecursive(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (rb *RBTree) GetBlackHeight() int {
	return rb.getBlackHeightRecursive(rb.Root)
}

func (rb *RBTree) getBlackHeightRecursive(node *Node) int {
	if node == rb.NIL {
		return 0
	}

	leftBlackHeight := rb.getBlackHeightRecursive(node.Left)
	if node.Color == BLACK {
		return leftBlackHeight + 1
	}
	return leftBlackHeight
}

func (rb *RBTree) IsValidRBTree() bool {
	if rb.Root == rb.NIL {
		return true
	}

	if rb.Root.Color != BLACK {
		return false
	}

	_, valid := rb.validateRBProperties(rb.Root)
	return valid
}

func (rb *RBTree) validateRBProperties(node *Node) (int, bool) {
	if node == rb.NIL {
		return 0, true
	}

	if node.Color == RED {
		if (node.Left != rb.NIL && node.Left.Color == RED) ||
			(node.Right != rb.NIL && node.Right.Color == RED) {
			return 0, false
		}
	}

	leftBlackHeight, leftValid := rb.validateRBProperties(node.Left)
	if !leftValid {
		return 0, false
	}

	rightBlackHeight, rightValid := rb.validateRBProperties(node.Right)
	if !rightValid {
		return 0, false
	}

	if leftBlackHeight != rightBlackHeight {
		return 0, false
	}

	blackHeight := leftBlackHeight
	if node.Color == BLACK {
		blackHeight++
	}

	return blackHeight, true
}

func (rb *RBTree) InOrderTraversal() []int {
	var result []int
	rb.inOrderRecursive(rb.Root, &result)
	return result
}

func (rb *RBTree) inOrderRecursive(node *Node, result *[]int) {
	if node != rb.NIL {
		rb.inOrderRecursive(node.Left, result)
		*result = append(*result, node.Value)
		rb.inOrderRecursive(node.Right, result)
	}
}

func (rb *RBTree) PreOrderTraversal() []int {
	var result []int
	rb.preOrderRecursive(rb.Root, &result)
	return result
}

func (rb *RBTree) preOrderRecursive(node *Node, result *[]int) {
	if node != rb.NIL {
		*result = append(*result, node.Value)
		rb.preOrderRecursive(node.Left, result)
		rb.preOrderRecursive(node.Right, result)
	}
}

func (rb *RBTree) LevelOrderTraversal() []int {
	if rb.Root == rb.NIL {
		return []int{}
	}

	var result []int
	queue := []*Node{rb.Root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != rb.NIL {
			queue = append(queue, node.Left)
		}
		if node.Right != rb.NIL {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func (rb *RBTree) FindMin() (int, bool) {
	if rb.Root == rb.NIL {
		return 0, false
	}
	node := rb.minimum(rb.Root)
	return node.Value, true
}

func (rb *RBTree) FindMax() (int, bool) {
	if rb.Root == rb.NIL {
		return 0, false
	}
	node := rb.maximum(rb.Root)
	return node.Value, true
}

func (rb *RBTree) maximum(node *Node) *Node {
	for node.Right != rb.NIL {
		node = node.Right
	}
	return node
}

func (rb *RBTree) Clear() {
	rb.Root = rb.NIL
	rb.Size = 0
}

func (rb *RBTree) PrintTree() {
	rb.printTreeRecursive(rb.Root, "", true)
}

func (rb *RBTree) printTreeRecursive(node *Node, prefix string, isLast bool) {
	if node == rb.NIL {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	colorStr := "R"
	if node.Color == BLACK {
		colorStr = "B"
	}

	fmt.Printf("%s%s%d (%s)\n", prefix, connector, node.Value, colorStr)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	children := []*Node{}
	if node.Left != rb.NIL {
		children = append(children, node.Left)
	}
	if node.Right != rb.NIL {
		children = append(children, node.Right)
	}

	for i, child := range children {
		isLastChild := i == len(children)-1
		rb.printTreeRecursive(child, childPrefix, isLastChild)
	}
}

func Run() any {
	rb := NewRBTree()

	values := []int{10, 20, 30, 40, 50, 25, 15, 35}
	for _, value := range values {
		rb.Insert(value)
	}

	result := make(map[string]any)
	result["size"] = rb.GetSize()
	result["height"] = rb.GetHeight()
	result["blackHeight"] = rb.GetBlackHeight()
	result["isValidRBTree"] = rb.IsValidRBTree()
	result["inOrder"] = rb.InOrderTraversal()
	result["preOrder"] = rb.PreOrderTraversal()
	result["levelOrder"] = rb.LevelOrderTraversal()

	result["searchExisting"] = rb.Search(30)
	result["searchNonExisting"] = rb.Search(100)

	min, hasMin := rb.FindMin()
	max, hasMax := rb.FindMax()
	result["min"] = map[string]any{"value": min, "exists": hasMin}
	result["max"] = map[string]any{"value": max, "exists": hasMax}

	result["deleteSuccess"] = rb.Delete(20)
	result["sizeAfterDelete"] = rb.GetSize()
	result["heightAfterDelete"] = rb.GetHeight()
	result["isValidAfterDelete"] = rb.IsValidRBTree()
	result["inOrderAfterDelete"] = rb.InOrderTraversal()

	rb.InsertIterative(12)
	result["sizeAfterIterativeInsert"] = rb.GetSize()
	result["searchIterative"] = rb.SearchIterative(12)
	result["isValidAfterIterativeInsert"] = rb.IsValidRBTree()

	return result
}
