package linked_list

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) Delete(data int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		return true
	}

	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}

	return false
}

func (ll *LinkedList) Search(data int) bool {
	current := ll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList) ToSlice() []int {
	var result []int
	current := ll.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}

func Run() any {
	ll := NewLinkedList()

	ll.Insert(30)
	ll.Insert(20)
	ll.Insert(10)

	list := ll.ToSlice()
	found := ll.Search(20)
	ll.Delete(20)
	afterDelete := ll.ToSlice()

	return map[string]any{
		"original":     list,
		"found_20":     found,
		"after_delete": afterDelete,
	}
}
