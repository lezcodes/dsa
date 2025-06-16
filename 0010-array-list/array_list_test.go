package array_list

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestNewArrayList(t *testing.T) {
	al := NewArrayList()
	if al == nil {
		t.Error("Expected non-nil ArrayList")
	}
}

func TestAddGet(t *testing.T) {
	al := NewArrayList()

	al.Add(10)
	al.Add(20)
	al.Add(30)

	val, err := al.Get(0)
	if err != nil || val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	val, err = al.Get(1)
	if err != nil || val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val, err = al.Get(2)
	if err != nil || val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
}

func TestSetRemove(t *testing.T) {
	al := NewArrayList()
	al.Add(10)
	al.Add(20)
	al.Add(30)

	err := al.Set(1, 25)
	if err != nil {
		t.Error("Expected no error when setting valid index")
	}

	val, err := al.Get(1)
	if err != nil || val != 25 {
		t.Errorf("Expected 25, got %d", val)
	}

	removedVal, err := al.Remove(0)
	if err != nil || removedVal != 10 {
		t.Errorf("Expected to remove 10, got %d", removedVal)
	}

	val, err = al.Get(0)
	if err != nil || val != 25 {
		t.Errorf("Expected first element to now be 25, got %d", val)
	}
}

func TestOutOfBounds(t *testing.T) {
	al := NewArrayList()
	al.Add(10)

	_, err := al.Get(5)
	if err == nil {
		t.Error("Expected error for out of bounds access")
	}

	err = al.Set(5, 50)
	if err == nil {
		t.Error("Expected error for out of bounds set")
	}

	_, err = al.Remove(5)
	if err == nil {
		t.Error("Expected error for out of bounds remove")
	}
}
