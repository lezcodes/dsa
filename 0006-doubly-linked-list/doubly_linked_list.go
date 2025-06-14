package doubly_linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func (dll *DoublyLinkedList) Length() int {
	return dll.size
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList) Prepend(data int) {
	newNode := &Node{Data: data, Next: nil, Prev: nil}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}

	dll.size++
}

func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil, Prev: nil}

	if dll.Tail == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}

	dll.size++
}

func (dll *DoublyLinkedList) InsertAt(data int, index int) error {
	if index < 0 || index > dll.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	if index == 0 {
		dll.Prepend(data)
		return nil
	}

	if index == dll.size {
		dll.Append(data)
		return nil
	}

	newNode := &Node{Data: data, Next: nil, Prev: nil}

	if index <= dll.size/2 {
		current := dll.Head
		for range index {
			current = current.Next
		}

		newNode.Next = current
		newNode.Prev = current.Prev
		current.Prev.Next = newNode
		current.Prev = newNode
	} else {
		current := dll.Tail
		for range dll.size - index - 1 {
			current = current.Prev
		}

		newNode.Next = current
		newNode.Prev = current.Prev
		current.Prev.Next = newNode
		current.Prev = newNode
	}

	dll.size++
	return nil
}

func (dll *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index >= dll.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}

	var current *Node

	if index <= dll.size/2 {
		current = dll.Head
		for range index {
			current = current.Next
		}
	} else {
		current = dll.Tail
		for range dll.size - index - 1 {
			current = current.Prev
		}
	}

	return current.Data, nil
}

func (dll *DoublyLinkedList) Remove(data int) (int, bool) {
	if dll.Head == nil {
		return 0, false
	}

	current := dll.Head
	for current != nil {
		if current.Data == data {
			dll.removeNode(current)
			return data, true
		}
		current = current.Next
	}

	return 0, false
}

func (dll *DoublyLinkedList) RemoveAt(index int) (int, error) {
	if index < 0 || index >= dll.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}

	var current *Node

	if index <= dll.size/2 {
		current = dll.Head
		for range index {
			current = current.Next
		}
	} else {
		current = dll.Tail
		for range dll.size - index - 1 {
			current = current.Prev
		}
	}

	data := current.Data
	dll.removeNode(current)
	return data, nil
}

func (dll *DoublyLinkedList) removeNode(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.Tail = node.Prev
	}

	dll.size--
}

func (dll *DoublyLinkedList) RemoveHead() (int, error) {
	if dll.Head == nil {
		return 0, fmt.Errorf("list is empty")
	}

	data := dll.Head.Data
	dll.removeNode(dll.Head)
	return data, nil
}

func (dll *DoublyLinkedList) RemoveTail() (int, error) {
	if dll.Tail == nil {
		return 0, fmt.Errorf("list is empty")
	}

	data := dll.Tail.Data
	dll.removeNode(dll.Tail)
	return data, nil
}

func (dll *DoublyLinkedList) Search(data int) int {
	current := dll.Head
	index := 0

	for current != nil {
		if current.Data == data {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

func (dll *DoublyLinkedList) Contains(data int) bool {
	return dll.Search(data) != -1
}

func (dll *DoublyLinkedList) Clear() {
	dll.Head = nil
	dll.Tail = nil
	dll.size = 0
}

func (dll *DoublyLinkedList) ToSlice() []int {
	result := make([]int, 0, dll.size)
	current := dll.Head

	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}

	return result
}

func (dll *DoublyLinkedList) ToSliceReverse() []int {
	result := make([]int, 0, dll.size)
	current := dll.Tail

	for current != nil {
		result = append(result, current.Data)
		current = current.Prev
	}

	return result
}

func (dll *DoublyLinkedList) Display() string {
	if dll.Head == nil {
		return "[]"
	}

	result := "["
	current := dll.Head

	for current != nil {
		result += fmt.Sprintf("%d", current.Data)
		if current.Next != nil {
			result += " <-> "
		}
		current = current.Next
	}

	result += "]"
	return result
}

func (dll *DoublyLinkedList) Reverse() {
	if dll.Head == nil || dll.Head.Next == nil {
		return
	}

	current := dll.Head

	for current != nil {
		current.Next, current.Prev = current.Prev, current.Next
		current = current.Prev
	}

	dll.Head, dll.Tail = dll.Tail, dll.Head
}

func (dll *DoublyLinkedList) GetMiddle() (int, error) {
	if dll.Head == nil {
		return 0, fmt.Errorf("list is empty")
	}

	slow := dll.Head
	fast := dll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Data, nil
}

func (dll *DoublyLinkedList) GetHead() (int, error) {
	if dll.Head == nil {
		return 0, fmt.Errorf("list is empty")
	}
	return dll.Head.Data, nil
}

func (dll *DoublyLinkedList) GetTail() (int, error) {
	if dll.Tail == nil {
		return 0, fmt.Errorf("list is empty")
	}
	return dll.Tail.Data, nil
}

func Run() any {
	dll := NewDoublyLinkedList()

	dll.Append(10)
	dll.Append(20)
	dll.Prepend(5)
	dll.InsertAt(15, 2)
	dll.Append(30)

	originalList := dll.ToSlice()

	searchResult := dll.Search(15)
	middle, _ := dll.GetMiddle()

	removedValue, found := dll.Remove(20)
	afterRemove := dll.ToSlice()

	dll.Reverse()
	reversedList := dll.ToSlice()

	reverseTraversal := dll.ToSliceReverse()

	return map[string]any{
		"original_list":     originalList,
		"search_15_index":   searchResult,
		"middle_element":    middle,
		"removed_value":     removedValue,
		"removed_found":     found,
		"after_remove":      afterRemove,
		"reversed_list":     reversedList,
		"reverse_traversal": reverseTraversal,
		"length":            dll.Length(),
		"contains_30":       dll.Contains(30),
		"display":           dll.Display(),
		"data_structure":    "Doubly Linked List",
		"bidirectional":     true,
	}
}
