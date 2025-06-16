package m_way_tree

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

type MWayTree struct {
	Root *Node
	M    int
	Size int
}

func NewMWayTree(m int) *MWayTree {
	if m < 2 {
		m = 3
	}
	return &MWayTree{
		Root: nil,
		M:    m,
		Size: 0,
	}
}

func (tree *MWayTree) newNode(isLeaf bool) *Node {
	return &Node{
		Keys:     make([]int, 0, tree.M-1),
		Children: make([]*Node, 0, tree.M),
		IsLeaf:   isLeaf,
		Parent:   nil,
	}
}

func (tree *MWayTree) Insert(key int) {
	if tree.Root == nil {
		tree.Root = tree.newNode(true)
		tree.Root.Keys = append(tree.Root.Keys, key)
		tree.Size++
		return
	}

	if tree.insertRecursive(tree.Root, key) {
		tree.Size++
	}
}

func (tree *MWayTree) insertRecursive(node *Node, key int) bool {
	pos := tree.findKeyPosition(node.Keys, key)

	if pos < len(node.Keys) && node.Keys[pos] == key {
		return false
	}

	if node.IsLeaf {
		node.Keys = append(node.Keys, 0)
		copy(node.Keys[pos+1:], node.Keys[pos:])
		node.Keys[pos] = key

		if len(node.Keys) >= tree.M {
			tree.splitNode(node)
		}
		return true
	}

	if tree.insertRecursive(node.Children[pos], key) {
		if len(node.Keys) >= tree.M {
			tree.splitNode(node)
		}
		return true
	}
	return false
}

func (tree *MWayTree) splitNode(node *Node) {
	if len(node.Keys) < tree.M {
		return
	}

	mid := len(node.Keys) / 2
	midKey := node.Keys[mid]

	rightNode := tree.newNode(node.IsLeaf)
	rightNode.Keys = append(rightNode.Keys, node.Keys[mid+1:]...)
	node.Keys = node.Keys[:mid]

	if !node.IsLeaf {
		rightNode.Children = append(rightNode.Children, node.Children[mid+1:]...)
		node.Children = node.Children[:mid+1]

		for _, child := range rightNode.Children {
			child.Parent = rightNode
		}
	}

	if node.Parent == nil {
		newRoot := tree.newNode(false)
		newRoot.Keys = append(newRoot.Keys, midKey)
		newRoot.Children = append(newRoot.Children, node, rightNode)
		node.Parent = newRoot
		rightNode.Parent = newRoot
		tree.Root = newRoot
	} else {
		parent := node.Parent
		pos := tree.findKeyPosition(parent.Keys, midKey)

		parent.Keys = append(parent.Keys, 0)
		copy(parent.Keys[pos+1:], parent.Keys[pos:])
		parent.Keys[pos] = midKey

		parent.Children = append(parent.Children, nil)
		copy(parent.Children[pos+2:], parent.Children[pos+1:])
		parent.Children[pos+1] = rightNode
		rightNode.Parent = parent
	}
}

func (tree *MWayTree) Delete(key int) bool {
	if tree.Root == nil {
		return false
	}

	if tree.deleteRecursive(tree.Root, key) {
		tree.Size--
		if tree.Root != nil && len(tree.Root.Keys) == 0 && !tree.Root.IsLeaf {
			tree.Root = tree.Root.Children[0]
			tree.Root.Parent = nil
		}
		return true
	}
	return false
}

func (tree *MWayTree) deleteRecursive(node *Node, key int) bool {
	pos := tree.findKeyPosition(node.Keys, key)

	if pos < len(node.Keys) && node.Keys[pos] == key {
		if node.IsLeaf {
			copy(node.Keys[pos:], node.Keys[pos+1:])
			node.Keys = node.Keys[:len(node.Keys)-1]
			return true
		} else {
			return false
		}
	}

	if node.IsLeaf {
		return false
	}

	if pos < len(node.Children) {
		return tree.deleteRecursive(node.Children[pos], key)
	}

	return false
}

func (tree *MWayTree) Search(key int) bool {
	return tree.searchRecursive(tree.Root, key)
}

func (tree *MWayTree) searchRecursive(node *Node, key int) bool {
	if node == nil {
		return false
	}

	pos := tree.findKeyPosition(node.Keys, key)

	if pos < len(node.Keys) && node.Keys[pos] == key {
		return true
	}

	if node.IsLeaf {
		return false
	}

	return tree.searchRecursive(node.Children[pos], key)
}

func (tree *MWayTree) SearchIterative(key int) bool {
	current := tree.Root

	for current != nil {
		pos := tree.findKeyPosition(current.Keys, key)

		if pos < len(current.Keys) && current.Keys[pos] == key {
			return true
		}

		if current.IsLeaf {
			return false
		}

		current = current.Children[pos]
	}

	return false
}

func (tree *MWayTree) findKeyPosition(keys []int, key int) int {
	return sort.SearchInts(keys, key)
}

func (tree *MWayTree) GetSize() int {
	return tree.Size
}

func (tree *MWayTree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *MWayTree) GetHeight() int {
	return tree.getHeightRecursive(tree.Root)
}

func (tree *MWayTree) getHeightRecursive(node *Node) int {
	if node == nil {
		return -1
	}

	if node.IsLeaf {
		return 0
	}

	maxHeight := -1
	for _, child := range node.Children {
		height := tree.getHeightRecursive(child)
		if height > maxHeight {
			maxHeight = height
		}
	}

	return maxHeight + 1
}

func (tree *MWayTree) GetBranchingFactor() int {
	return tree.M
}

func (tree *MWayTree) InOrderTraversal() []int {
	var result []int
	tree.inOrderRecursive(tree.Root, &result)
	return result
}

func (tree *MWayTree) inOrderRecursive(node *Node, result *[]int) {
	if node == nil {
		return
	}

	if node.IsLeaf {
		*result = append(*result, node.Keys...)
		return
	}

	for i := range node.Keys {
		tree.inOrderRecursive(node.Children[i], result)
		*result = append(*result, node.Keys[i])
	}

	if len(node.Children) > len(node.Keys) {
		tree.inOrderRecursive(node.Children[len(node.Keys)], result)
	}
}

func (tree *MWayTree) PreOrderTraversal() []int {
	var result []int
	tree.preOrderRecursive(tree.Root, &result)
	return result
}

func (tree *MWayTree) preOrderRecursive(node *Node, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Keys...)

	for _, child := range node.Children {
		tree.preOrderRecursive(child, result)
	}
}

func (tree *MWayTree) PostOrderTraversal() []int {
	var result []int
	tree.postOrderRecursive(tree.Root, &result)
	return result
}

func (tree *MWayTree) postOrderRecursive(node *Node, result *[]int) {
	if node == nil {
		return
	}

	for _, child := range node.Children {
		tree.postOrderRecursive(child, result)
	}

	*result = append(*result, node.Keys...)
}

func (tree *MWayTree) LevelOrderTraversal() []int {
	if tree.Root == nil {
		return []int{}
	}

	var result []int
	queue := []*Node{tree.Root}

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

func (tree *MWayTree) FindMin() (int, bool) {
	if tree.Root == nil {
		return 0, false
	}

	current := tree.Root
	for !current.IsLeaf {
		current = current.Children[0]
	}

	if len(current.Keys) > 0 {
		return current.Keys[0], true
	}
	return 0, false
}

func (tree *MWayTree) FindMax() (int, bool) {
	if tree.Root == nil {
		return 0, false
	}

	current := tree.Root
	for !current.IsLeaf {
		current = current.Children[len(current.Children)-1]
	}

	if len(current.Keys) > 0 {
		return current.Keys[len(current.Keys)-1], true
	}
	return 0, false
}

func (tree *MWayTree) GetAllKeys() []int {
	var keys []int
	tree.collectKeysRecursive(tree.Root, &keys)
	sort.Ints(keys)
	return keys
}

func (tree *MWayTree) collectKeysRecursive(node *Node, keys *[]int) {
	if node == nil {
		return
	}

	*keys = append(*keys, node.Keys...)

	for _, child := range node.Children {
		tree.collectKeysRecursive(child, keys)
	}
}

func (tree *MWayTree) GetNodeCount() int {
	return tree.getNodeCountRecursive(tree.Root)
}

func (tree *MWayTree) getNodeCountRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	count := 1
	for _, child := range node.Children {
		count += tree.getNodeCountRecursive(child)
	}
	return count
}

func (tree *MWayTree) GetLeafCount() int {
	return tree.getLeafCountRecursive(tree.Root)
}

func (tree *MWayTree) getLeafCountRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	if node.IsLeaf {
		return 1
	}

	count := 0
	for _, child := range node.Children {
		count += tree.getLeafCountRecursive(child)
	}
	return count
}

func (tree *MWayTree) Clear() {
	tree.Root = nil
	tree.Size = 0
}

func (tree *MWayTree) PrintTree() {
	tree.printTreeRecursive(tree.Root, "", true, 0)
}

func (tree *MWayTree) printTreeRecursive(node *Node, prefix string, isLast bool, level int) {
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
		tree.printTreeRecursive(child, childPrefix, isLastChild, level+1)
	}
}

func (tree *MWayTree) Validate() bool {
	if tree.Root == nil {
		return true
	}
	return tree.validateRecursive(tree.Root)
}

func (tree *MWayTree) validateRecursive(node *Node) bool {
	if len(node.Keys) >= tree.M {
		return false
	}

	for i := 1; i < len(node.Keys); i++ {
		if node.Keys[i-1] >= node.Keys[i] {
			return false
		}
	}

	if node.IsLeaf {
		return len(node.Children) == 0
	}

	if len(node.Children) != len(node.Keys)+1 {
		return false
	}

	for _, child := range node.Children {
		if child.Parent != node {
			return false
		}
		if !tree.validateRecursive(child) {
			return false
		}
	}

	return true
}

func Run() any {
	tree := NewMWayTree(4)

	values := []int{10, 20, 5, 6, 12, 30, 7, 17, 25, 35, 40, 50}
	for _, value := range values {
		tree.Insert(value)
	}

	result := make(map[string]any)
	result["branchingFactor"] = tree.GetBranchingFactor()
	result["size"] = tree.GetSize()
	result["height"] = tree.GetHeight()
	result["nodeCount"] = tree.GetNodeCount()
	result["leafCount"] = tree.GetLeafCount()
	result["isValid"] = tree.Validate()

	result["inOrder"] = tree.InOrderTraversal()
	result["preOrder"] = tree.PreOrderTraversal()
	result["postOrder"] = tree.PostOrderTraversal()
	result["levelOrder"] = tree.LevelOrderTraversal()
	result["allKeys"] = tree.GetAllKeys()

	result["searchExisting"] = tree.Search(20)
	result["searchNonExisting"] = tree.Search(100)
	result["searchIterative"] = tree.SearchIterative(25)

	min, hasMin := tree.FindMin()
	max, hasMax := tree.FindMax()
	result["min"] = map[string]any{"value": min, "exists": hasMin}
	result["max"] = map[string]any{"value": max, "exists": hasMax}

	result["deleteSuccess"] = tree.Delete(12)
	result["sizeAfterDelete"] = tree.GetSize()
	result["isValidAfterDelete"] = tree.Validate()
	result["allKeysAfterDelete"] = tree.GetAllKeys()

	tree3Way := NewMWayTree(3)
	for _, value := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		tree3Way.Insert(value)
	}
	result["tree3WayHeight"] = tree3Way.GetHeight()
	result["tree3WaySize"] = tree3Way.GetSize()
	result["tree3WayValid"] = tree3Way.Validate()

	return result
}
