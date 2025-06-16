package hash_map

import (
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result from Run()")
	}
}

func TestNewHashMap(t *testing.T) {
	hm := NewHashMap()
	if hm == nil {
		t.Error("Expected non-nil HashMap")
	}
	if hm.Size() != 0 {
		t.Errorf("Expected empty HashMap, got size %d", hm.Size())
	}
}

func TestSetAndGet(t *testing.T) {
	hm := NewHashMap()

	hm.Set("key1", "value1")
	value, exists := hm.Get("key1")
	if !exists {
		t.Error("Expected key1 to exist")
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	hm.Set("key2", 42)
	value, exists = hm.Get("key2")
	if !exists {
		t.Error("Expected key2 to exist")
	}
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}

	hm.Set("key1", "updated_value")
	value, exists = hm.Get("key1")
	if !exists {
		t.Error("Expected key1 to still exist after update")
	}
	if value != "updated_value" {
		t.Errorf("Expected updated_value, got %v", value)
	}
	if hm.Size() != 2 {
		t.Errorf("Expected size 2 after update, got %d", hm.Size())
	}
}

func TestGetNonExistent(t *testing.T) {
	hm := NewHashMap()
	value, exists := hm.Get("nonexistent")
	if exists {
		t.Error("Expected nonexistent key to not exist")
	}
	if value != nil {
		t.Errorf("Expected nil value for nonexistent key, got %v", value)
	}
}

func TestDelete(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	if hm.Size() != 3 {
		t.Errorf("Expected size 3, got %d", hm.Size())
	}

	deleted := hm.Delete("key2")
	if !deleted {
		t.Error("Expected key2 to be deleted")
	}
	if hm.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", hm.Size())
	}

	_, exists := hm.Get("key2")
	if exists {
		t.Error("Expected key2 to not exist after deletion")
	}

	deleted = hm.Delete("nonexistent")
	if deleted {
		t.Error("Expected deletion of nonexistent key to return false")
	}
	if hm.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", hm.Size())
	}
}

func TestHas(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")

	if !hm.Has("key1") {
		t.Error("Expected key1 to exist")
	}
	if hm.Has("nonexistent") {
		t.Error("Expected nonexistent key to not exist")
	}
}

func TestKeys(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	keys := hm.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}

	expectedKeys := []string{"key1", "key2", "key3"}
	for _, expectedKey := range expectedKeys {
		if !keySet[expectedKey] {
			t.Errorf("Expected key %s to be in keys", expectedKey)
		}
	}
}
