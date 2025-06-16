package b_tree

import (
	"fmt"
	"sort"
)

type Node struct {
	Keys     []int
	Children []*Node
	IsLeaf   bool
	Parent   *Node
}

type BTree struct {
	Root *Node
	T    int
	Size int
}

func NewBTree(t int) *BTree {
	if t < 2 {
		t = 2
	}
	return &BTree{
		Root: nil,
		T:    t,
		Size: 0,
	}
}

func (bt *BTree) newNode(isLeaf bool) *Node {
	return &Node{
		Keys:     make([]int, 0, 2*bt.T-1),
		Children: make([]*Node, 0, 2*bt.T),
		IsLeaf:   isLeaf,
		Parent:   nil,
	}
}

func (bt *BTree) Insert(key int) {
	if bt.Root == nil {
		bt.Root = bt.newNode(true)
		bt.Root.Keys = append(bt.Root.Keys, key)
		bt.Size++
		return
	}

	if len(bt.Root.Keys) == 2*bt.T-1 {
		newRoot := bt.newNode(false)
		newRoot.Children = append(newRoot.Children, bt.Root)
		bt.Root.Parent = newRoot
		bt.splitChild(newRoot, 0)
		bt.Root = newRoot
	}

	bt.insertNonFull(bt.Root, key)
}

func (bt *BTree) insertNonFull(node *Node, key int) {
	i := len(node.Keys) - 1

	if node.IsLeaf {
		node.Keys = append(node.Keys, 0)
		for i >= 0 && node.Keys[i] > key {
			node.Keys[i+1] = node.Keys[i]
			i--
		}
		node.Keys[i+1] = key
		bt.Size++
	} else {
		for i >= 0 && node.Keys[i] > key {
			i--
		}
		i++

		if len(node.Children[i].Keys) == 2*bt.T-1 {
			bt.splitChild(node, i)
			if node.Keys[i] < key {
				i++
			}
		}
		bt.insertNonFull(node.Children[i], key)
	}
}

func (bt *BTree) splitChild(parent *Node, index int) {
	fullChild := parent.Children[index]
	newChild := bt.newNode(fullChild.IsLeaf)

	medianKey := fullChild.Keys[bt.T-1]

	newChild.Keys = append(newChild.Keys, fullChild.Keys[bt.T:]...)
	fullChild.Keys = fullChild.Keys[:bt.T-1]

	if !fullChild.IsLeaf {
		newChild.Children = append(newChild.Children, fullChild.Children[bt.T:]...)
		fullChild.Children = fullChild.Children[:bt.T]

		for _, child := range newChild.Children {
			child.Parent = newChild
		}
	}

	parent.Children = append(parent.Children, nil)
	copy(parent.Children[index+2:], parent.Children[index+1:])
	parent.Children[index+1] = newChild
	newChild.Parent = parent

	parent.Keys = append(parent.Keys, 0)
	copy(parent.Keys[index+1:], parent.Keys[index:])
	parent.Keys[index] = medianKey
}

func (bt *BTree) Delete(key int) bool {
	if bt.Root == nil {
		return false
	}

	found := bt.deleteFromNode(bt.Root, key)

	if len(bt.Root.Keys) == 0 && !bt.Root.IsLeaf {
		bt.Root = bt.Root.Children[0]
		bt.Root.Parent = nil
	}

	if found {
		bt.Size--
	}

	return found
}

func (bt *BTree) deleteFromNode(node *Node, key int) bool {
	i := bt.findKeyIndex(node.Keys, key)

	if i < len(node.Keys) && node.Keys[i] == key {
		if node.IsLeaf {
			copy(node.Keys[i:], node.Keys[i+1:])
			node.Keys = node.Keys[:len(node.Keys)-1]
			return true
		} else {
			return bt.deleteFromInternalNode(node, i)
		}
	} else if node.IsLeaf {
		return false
	} else {
		shouldDeleteFromSubtree := (i == len(node.Keys))

		if len(node.Children[i].Keys) < bt.T {
			bt.fill(node, i)
		}

		if shouldDeleteFromSubtree && i > len(node.Keys) {
			return bt.deleteFromNode(node.Children[i-1], key)
		} else {
			return bt.deleteFromNode(node.Children[i], key)
		}
	}
}

func (bt *BTree) deleteFromInternalNode(node *Node, index int) bool {
	key := node.Keys[index]

	if len(node.Children[index].Keys) >= bt.T {
		pred := bt.getPredecessor(node, index)
		node.Keys[index] = pred
		return bt.deleteFromNode(node.Children[index], pred)
	} else if len(node.Children[index+1].Keys) >= bt.T {
		succ := bt.getSuccessor(node, index)
		node.Keys[index] = succ
		return bt.deleteFromNode(node.Children[index+1], succ)
	} else {
		bt.merge(node, index)
		return bt.deleteFromNode(node.Children[index], key)
	}
}

func (bt *BTree) getPredecessor(node *Node, index int) int {
	current := node.Children[index]
	for !current.IsLeaf {
		current = current.Children[len(current.Children)-1]
	}
	return current.Keys[len(current.Keys)-1]
}

func (bt *BTree) getSuccessor(node *Node, index int) int {
	current := node.Children[index+1]
	for !current.IsLeaf {
		current = current.Children[0]
	}
	return current.Keys[0]
}

func (bt *BTree) fill(node *Node, index int) {
	if index != 0 && len(node.Children[index-1].Keys) >= bt.T {
		bt.borrowFromPrev(node, index)
	} else if index != len(node.Children)-1 && len(node.Children[index+1].Keys) >= bt.T {
		bt.borrowFromNext(node, index)
	} else {
		if index != len(node.Children)-1 {
			bt.merge(node, index)
		} else {
			bt.merge(node, index-1)
		}
	}
}

func (bt *BTree) borrowFromPrev(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index-1]

	child.Keys = append([]int{node.Keys[index-1]}, child.Keys...)
	node.Keys[index-1] = sibling.Keys[len(sibling.Keys)-1]
	sibling.Keys = sibling.Keys[:len(sibling.Keys)-1]

	if !child.IsLeaf {
		child.Children = append([]*Node{sibling.Children[len(sibling.Children)-1]}, child.Children...)
		child.Children[0].Parent = child
		sibling.Children = sibling.Children[:len(sibling.Children)-1]
	}
}

func (bt *BTree) borrowFromNext(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index+1]

	child.Keys = append(child.Keys, node.Keys[index])
	node.Keys[index] = sibling.Keys[0]
	sibling.Keys = sibling.Keys[1:]

	if !child.IsLeaf {
		child.Children = append(child.Children, sibling.Children[0])
		child.Children[len(child.Children)-1].Parent = child
		sibling.Children = sibling.Children[1:]
	}
}

func (bt *BTree) merge(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index+1]

	child.Keys = append(child.Keys, node.Keys[index])
	child.Keys = append(child.Keys, sibling.Keys...)

	if !child.IsLeaf {
		child.Children = append(child.Children, sibling.Children...)
		for _, grandchild := range sibling.Children {
			grandchild.Parent = child
		}
	}

	copy(node.Keys[index:], node.Keys[index+1:])
	node.Keys = node.Keys[:len(node.Keys)-1]

	copy(node.Children[index+1:], node.Children[index+2:])
	node.Children = node.Children[:len(node.Children)-1]
}

func (bt *BTree) Search(key int) bool {
	return bt.searchInNode(bt.Root, key)
}

func (bt *BTree) searchInNode(node *Node, key int) bool {
	if node == nil {
		return false
	}

	i := 0
	for i < len(node.Keys) && key > node.Keys[i] {
		i++
	}

	if i < len(node.Keys) && key == node.Keys[i] {
		return true
	}

	if node.IsLeaf {
		return false
	}

	return bt.searchInNode(node.Children[i], key)
}

func (bt *BTree) SearchIterative(key int) bool {
	current := bt.Root

	for current != nil {
		i := 0
		for i < len(current.Keys) && key > current.Keys[i] {
			i++
		}

		if i < len(current.Keys) && key == current.Keys[i] {
			return true
		}

		if current.IsLeaf {
			return false
		}

		current = current.Children[i]
	}

	return false
}

func (bt *BTree) findKeyIndex(keys []int, key int) int {
	return sort.SearchInts(keys, key)
}

func (bt *BTree) GetSize() int {
	return bt.Size
}

func (bt *BTree) IsEmpty() bool {
	return bt.Root == nil
}

func (bt *BTree) GetHeight() int {
	return bt.getHeightRecursive(bt.Root)
}

func (bt *BTree) getHeightRecursive(node *Node) int {
	if node == nil {
		return -1
	}

	if node.IsLeaf {
		return 0
	}

	return bt.getHeightRecursive(node.Children[0]) + 1
}

func (bt *BTree) GetMinimumDegree() int {
	return bt.T
}

func (bt *BTree) InOrderTraversal() []int {
	var result []int
	bt.inOrderRecursive(bt.Root, &result)
	return result
}

func (bt *BTree) inOrderRecursive(node *Node, result *[]int) {
	if node == nil {
		return
	}

	if node.IsLeaf {
		*result = append(*result, node.Keys...)
		return
	}

	for i := range node.Keys {
		bt.inOrderRecursive(node.Children[i], result)
		*result = append(*result, node.Keys[i])
	}

	if len(node.Children) > len(node.Keys) {
		bt.inOrderRecursive(node.Children[len(node.Keys)], result)
	}
}

func (bt *BTree) PreOrderTraversal() []int {
	var result []int
	bt.preOrderRecursive(bt.Root, &result)
	return result
}

func (bt *BTree) preOrderRecursive(node *Node, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Keys...)

	for _, child := range node.Children {
		bt.preOrderRecursive(child, result)
	}
}

func (bt *BTree) LevelOrderTraversal() []int {
	if bt.Root == nil {
		return []int{}
	}

	var result []int
	queue := []*Node{bt.Root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Keys...)

		for _, child := range node.Children {
			if child != nil {
				queue = append(queue, child)
			}
		}
	}

	return result
}

func (bt *BTree) FindMin() (int, bool) {
	if bt.Root == nil {
		return 0, false
	}

	current := bt.Root
	for !current.IsLeaf {
		current = current.Children[0]
	}

	if len(current.Keys) > 0 {
		return current.Keys[0], true
	}
	return 0, false
}

func (bt *BTree) FindMax() (int, bool) {
	if bt.Root == nil {
		return 0, false
	}

	current := bt.Root
	for !current.IsLeaf {
		current = current.Children[len(current.Children)-1]
	}

	if len(current.Keys) > 0 {
		return current.Keys[len(current.Keys)-1], true
	}
	return 0, false
}

func (bt *BTree) GetAllKeys() []int {
	return bt.InOrderTraversal()
}

func (bt *BTree) GetNodeCount() int {
	return bt.getNodeCountRecursive(bt.Root)
}

func (bt *BTree) getNodeCountRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	count := 1
	for _, child := range node.Children {
		count += bt.getNodeCountRecursive(child)
	}
	return count
}

func (bt *BTree) GetLeafCount() int {
	return bt.getLeafCountRecursive(bt.Root)
}

func (bt *BTree) getLeafCountRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	if node.IsLeaf {
		return 1
	}

	count := 0
	for _, child := range node.Children {
		count += bt.getLeafCountRecursive(child)
	}
	return count
}

func (bt *BTree) Clear() {
	bt.Root = nil
	bt.Size = 0
}

func (bt *BTree) PrintTree() {
	bt.printTreeRecursive(bt.Root, "", true, 0)
}

func (bt *BTree) printTreeRecursive(node *Node, prefix string, isLast bool, level int) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	leafIndicator := ""
	if node.IsLeaf {
		leafIndicator = " (leaf)"
	}

	fmt.Printf("%s%s%v%s\n", prefix, connector, node.Keys, leafIndicator)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		bt.printTreeRecursive(child, childPrefix, isLastChild, level+1)
	}
}

func (bt *BTree) Validate() bool {
	if bt.Root == nil {
		return true
	}
	return bt.validateRecursive(bt.Root, nil, nil)
}

func (bt *BTree) validateRecursive(node *Node, minVal, maxVal *int) bool {
	if len(node.Keys) > 2*bt.T-1 {
		return false
	}

	if node != bt.Root && len(node.Keys) < bt.T-1 {
		return false
	}

	for i := 1; i < len(node.Keys); i++ {
		if node.Keys[i-1] >= node.Keys[i] {
			return false
		}
	}

	if minVal != nil && len(node.Keys) > 0 && node.Keys[0] <= *minVal {
		return false
	}

	if maxVal != nil && len(node.Keys) > 0 && node.Keys[len(node.Keys)-1] >= *maxVal {
		return false
	}

	if node.IsLeaf {
		return len(node.Children) == 0
	}

	if len(node.Children) != len(node.Keys)+1 {
		return false
	}

	for i, child := range node.Children {
		if child.Parent != node {
			return false
		}

		var childMin, childMax *int
		if i > 0 {
			childMin = &node.Keys[i-1]
		} else {
			childMin = minVal
		}
		if i < len(node.Keys) {
			childMax = &node.Keys[i]
		} else {
			childMax = maxVal
		}

		if !bt.validateRecursive(child, childMin, childMax) {
			return false
		}
	}

	return true
}

func Run() any {
	bt := NewBTree(3)

	values := []int{10, 20, 5, 6, 12, 30, 7, 17, 25, 35, 40, 50, 15, 18, 22, 27}
	for _, value := range values {
		bt.Insert(value)
	}

	result := make(map[string]any)
	result["minimumDegree"] = bt.GetMinimumDegree()
	result["size"] = bt.GetSize()
	result["height"] = bt.GetHeight()
	result["nodeCount"] = bt.GetNodeCount()
	result["leafCount"] = bt.GetLeafCount()
	result["isValid"] = bt.Validate()

	result["inOrder"] = bt.InOrderTraversal()
	result["preOrder"] = bt.PreOrderTraversal()
	result["levelOrder"] = bt.LevelOrderTraversal()
	result["allKeys"] = bt.GetAllKeys()

	result["searchExisting"] = bt.Search(20)
	result["searchNonExisting"] = bt.Search(100)
	result["searchIterative"] = bt.SearchIterative(25)

	min, hasMin := bt.FindMin()
	max, hasMax := bt.FindMax()
	result["min"] = map[string]any{"value": min, "exists": hasMin}
	result["max"] = map[string]any{"value": max, "exists": hasMax}

	result["deleteSuccess"] = bt.Delete(12)
	result["sizeAfterDelete"] = bt.GetSize()
	result["isValidAfterDelete"] = bt.Validate()
	result["allKeysAfterDelete"] = bt.GetAllKeys()

	bt2 := NewBTree(2)
	for _, value := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		bt2.Insert(value)
	}
	result["btree2Height"] = bt2.GetHeight()
	result["btree2Size"] = bt2.GetSize()
	result["btree2Valid"] = bt2.Validate()

	return result
}
