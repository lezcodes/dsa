package linked_list

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Error("Expected non-nil LinkedList")
	}
	if ll.Head != nil {
		t.Error("Expected empty LinkedList head to be nil")
	}
}

func TestInsert(t *testing.T) {
	ll := NewLinkedList()

	ll.Insert(1)
	if ll.Head == nil || ll.Head.Data != 1 {
		t.Error("Expected head to be 1")
	}

	ll.Insert(2)
	if ll.Head.Data != 2 {
		t.Error("Expected new head to be 2")
	}
	if ll.Head.Next == nil || ll.Head.Next.Data != 1 {
		t.Error("Expected second node to be 1")
	}
}

func TestSearch(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	if !ll.Search(2) {
		t.Error("Expected to find 2")
	}
	if ll.Search(4) {
		t.Error("Expected not to find 4")
	}
}

func TestDelete(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	if !ll.Delete(2) {
		t.Error("Expected to delete 2")
	}
	if ll.Search(2) {
		t.Error("Expected 2 to be deleted")
	}

	if ll.Delete(4) {
		t.Error("Expected deletion of non-existent element to return false")
	}
}
