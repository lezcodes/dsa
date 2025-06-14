package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		size: 0,
	}
}

func (ll *LinkedList) InsertAtHead(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.size++
}

func (ll *LinkedList) InsertAtTail(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.size++
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.size++
}

func (ll *LinkedList) InsertAtIndex(index int, data int) error {
	if index < 0 || index > ll.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	if index == 0 {
		ll.InsertAtHead(data)
		return nil
	}

	newNode := &Node{Data: data, Next: nil}
	current := ll.Head

	for range index - 1 {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	ll.size++
	return nil
}

func (ll *LinkedList) DeleteByValue(data int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.size--
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			ll.size--
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList) DeleteAtIndex(index int) error {
	if index < 0 || index >= ll.size {
		return fmt.Errorf("index out of bounds: %d", index)
	}

	if index == 0 {
		ll.Head = ll.Head.Next
		ll.size--
		return nil
	}

	current := ll.Head
	for range index - 1 {
		current = current.Next
	}

	current.Next = current.Next.Next
	ll.size--
	return nil
}

func (ll *LinkedList) Search(data int) int {
	current := ll.Head
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

func (ll *LinkedList) GetAt(index int) (int, error) {
	if index < 0 || index >= ll.size {
		return 0, fmt.Errorf("index out of bounds: %d", index)
	}

	current := ll.Head
	for range index {
		current = current.Next
	}

	return current.Data, nil
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList) Clear() {
	ll.Head = nil
	ll.size = 0
}

func (ll *LinkedList) ToSlice() []int {
	result := make([]int, 0, ll.size)
	current := ll.Head

	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}

	return result
}

func (ll *LinkedList) Display() string {
	if ll.Head == nil {
		return "[]"
	}

	result := "["
	current := ll.Head

	for current != nil {
		result += fmt.Sprintf("%d", current.Data)
		if current.Next != nil {
			result += " -> "
		}
		current = current.Next
	}

	result += "]"
	return result
}

func (ll *LinkedList) Reverse() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	var prev *Node
	current := ll.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	ll.Head = prev
}

func (ll *LinkedList) Contains(data int) bool {
	return ll.Search(data) != -1
}

func (ll *LinkedList) GetMiddle() (int, error) {
	if ll.Head == nil {
		return 0, fmt.Errorf("list is empty")
	}

	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Data, nil
}

func Run() any {
	ll := NewLinkedList()

	ll.InsertAtHead(10)
	ll.InsertAtHead(20)
	ll.InsertAtTail(30)
	ll.InsertAtTail(40)
	ll.InsertAtIndex(2, 25)

	originalList := ll.ToSlice()

	searchResult := ll.Search(25)
	middle, _ := ll.GetMiddle()

	ll.DeleteByValue(20)
	afterDelete := ll.ToSlice()

	ll.Reverse()
	reversed := ll.ToSlice()

	return map[string]any{
		"original_list":   originalList,
		"search_25_index": searchResult,
		"middle_element":  middle,
		"after_delete_20": afterDelete,
		"reversed_list":   reversed,
		"size":            ll.Size(),
		"contains_30":     ll.Contains(30),
		"display":         ll.Display(),
		"data_structure":  "Single Linked List",
		"operations":      "Insert, Delete, Search, Reverse",
	}
}
