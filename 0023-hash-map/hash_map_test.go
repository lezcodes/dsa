package hash_map

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Error("Expected result to be a map")
		return
	}

	if size, exists := resultMap["sizeAfterDelete"]; !exists || size != 4 {
		t.Errorf("Expected size to be 4 after deletion, got %v", size)
	}

	if capacity, exists := resultMap["capacity"]; !exists || capacity != 16 {
		t.Errorf("Expected initial capacity to be 16, got %v", capacity)
	}

	if loadFactor, exists := resultMap["loadFactor"]; !exists || loadFactor.(float64) <= 0 {
		t.Errorf("Expected positive load factor, got %v", loadFactor)
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
	if hm.Capacity() != DefaultCapacity {
		t.Errorf("Expected capacity %d, got %d", DefaultCapacity, hm.Capacity())
	}
	if !hm.IsEmpty() {
		t.Error("Expected empty HashMap")
	}
}

func TestNewHashMapWithCapacity(t *testing.T) {
	capacity := 32
	hm := NewHashMapWithCapacity(capacity)
	if hm.Capacity() != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, hm.Capacity())
	}

	hm = NewHashMapWithCapacity(0)
	if hm.Capacity() != DefaultCapacity {
		t.Errorf("Expected default capacity %d for invalid input, got %d", DefaultCapacity, hm.Capacity())
	}

	hm = NewHashMapWithCapacity(-5)
	if hm.Capacity() != DefaultCapacity {
		t.Errorf("Expected default capacity %d for negative input, got %d", DefaultCapacity, hm.Capacity())
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

func TestClear(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	if hm.Size() != 3 {
		t.Errorf("Expected size 3, got %d", hm.Size())
	}

	hm.Clear()
	if hm.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", hm.Size())
	}
	if !hm.IsEmpty() {
		t.Error("Expected empty HashMap after clear")
	}

	_, exists := hm.Get("key1")
	if exists {
		t.Error("Expected key1 to not exist after clear")
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

func TestValues(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	values := hm.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	valueSet := make(map[string]bool)
	for _, value := range values {
		if str, ok := value.(string); ok {
			valueSet[str] = true
		}
	}

	expectedValues := []string{"value1", "value2", "value3"}
	for _, expectedValue := range expectedValues {
		if !valueSet[expectedValue] {
			t.Errorf("Expected value %s to be in values", expectedValue)
		}
	}
}

func TestEntries(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")

	entries := hm.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	entryMap := make(map[string]any)
	for _, entry := range entries {
		entryMap[entry.Key] = entry.Value
	}

	if entryMap["key1"] != "value1" {
		t.Errorf("Expected key1 -> value1, got %v", entryMap["key1"])
	}
	if entryMap["key2"] != "value2" {
		t.Errorf("Expected key2 -> value2, got %v", entryMap["key2"])
	}
}

func TestForEach(t *testing.T) {
	hm := NewHashMap()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	visited := make(map[string]any)
	hm.ForEach(func(key string, value any) {
		visited[key] = value
	})

	if len(visited) != 3 {
		t.Errorf("Expected 3 entries visited, got %d", len(visited))
	}

	expectedEntries := map[string]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	for key, expectedValue := range expectedEntries {
		if visited[key] != expectedValue {
			t.Errorf("Expected %s -> %v, got %v", key, expectedValue, visited[key])
		}
	}
}

func TestResize(t *testing.T) {
	hm := NewHashMapWithCapacity(4)
	initialCapacity := hm.Capacity()

	hm.Set("key1", "value1")
	hm.Set("key2", "value2")

	if hm.LoadFactor() >= DefaultLoadFactor {
		t.Errorf("Load factor %.2f should be less than %.2f before resize", hm.LoadFactor(), DefaultLoadFactor)
	}

	hm.Set("key3", "value3")
	hm.Set("key4", "value4")

	if hm.Capacity() <= initialCapacity {
		t.Errorf("Expected capacity to increase from %d, got %d", initialCapacity, hm.Capacity())
	}

	for i := 1; i <= 4; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		retrievedValue, exists := hm.Get(key)
		if !exists {
			t.Errorf("Expected %s to exist after resize", key)
		}
		if retrievedValue != value {
			t.Errorf("Expected %s -> %s after resize, got %v", key, value, retrievedValue)
		}
	}
}

func TestCollisionHandling(t *testing.T) {
	hm := NewHashMapWithCapacity(1)

	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		retrievedValue, exists := hm.Get(key)
		if !exists {
			t.Errorf("Expected %s to exist with collisions", key)
		}
		if retrievedValue != value {
			t.Errorf("Expected %s -> %s with collisions, got %v", key, value, retrievedValue)
		}
	}

	if hm.Size() != 3 {
		t.Errorf("Expected size 3 with collisions, got %d", hm.Size())
	}
}

func TestDeleteWithCollisions(t *testing.T) {
	hm := NewHashMapWithCapacity(1)

	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	deleted := hm.Delete("key2")
	if !deleted {
		t.Error("Expected key2 to be deleted with collisions")
	}

	_, exists := hm.Get("key2")
	if exists {
		t.Error("Expected key2 to not exist after deletion with collisions")
	}

	value1, exists1 := hm.Get("key1")
	value3, exists3 := hm.Get("key3")
	if !exists1 || !exists3 {
		t.Error("Expected key1 and key3 to still exist after deleting key2 with collisions")
	}
	if value1 != "value1" || value3 != "value3" {
		t.Error("Expected correct values for remaining keys after deletion with collisions")
	}

	if hm.Size() != 2 {
		t.Errorf("Expected size 2 after deletion with collisions, got %d", hm.Size())
	}
}

func TestEmptyHashMapOperations(t *testing.T) {
	hm := NewHashMap()

	if !hm.IsEmpty() {
		t.Error("Expected new HashMap to be empty")
	}

	keys := hm.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected empty keys slice, got %d keys", len(keys))
	}

	values := hm.Values()
	if len(values) != 0 {
		t.Errorf("Expected empty values slice, got %d values", len(values))
	}

	entries := hm.Entries()
	if len(entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(entries))
	}

	hm.ForEach(func(key string, value any) {
		t.Error("ForEach should not be called on empty HashMap")
	})

	if hm.LoadFactor() != 0 {
		t.Errorf("Expected load factor 0 for empty HashMap, got %f", hm.LoadFactor())
	}
}

func TestLargeDataset(t *testing.T) {
	hm := NewHashMap()
	numItems := 1000

	for i := range numItems {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hm.Set(key, value)
	}

	if hm.Size() != numItems {
		t.Errorf("Expected size %d, got %d", numItems, hm.Size())
	}

	for i := range numItems {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		value, exists := hm.Get(key)
		if !exists {
			t.Errorf("Expected %s to exist", key)
		}
		if value != expectedValue {
			t.Errorf("Expected %s -> %s, got %v", key, expectedValue, value)
		}
	}

	for i := 0; i < numItems; i += 2 {
		key := fmt.Sprintf("key%d", i)
		hm.Delete(key)
	}

	expectedSize := numItems / 2
	if hm.Size() != expectedSize {
		t.Errorf("Expected size %d after deletions, got %d", expectedSize, hm.Size())
	}
}

func BenchmarkHashMapSet(b *testing.B) {
	hm := NewHashMap()
	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		hm.Set(key, b.N)
	}
}

func BenchmarkHashMapGet(b *testing.B) {
	hm := NewHashMap()
	for i := range 1000 {
		hm.Set(strconv.Itoa(i), i)
	}

	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		hm.Get(key)
	}
}

func BenchmarkHashMapDelete(b *testing.B) {
	hm := NewHashMap()
	for i := 0; b.Loop(); i++ {
		hm.Set(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		hm.Delete(key)
	}
}

func BenchmarkHashMapSetWithCollisions(b *testing.B) {
	hm := NewHashMapWithCapacity(1)
	for b.Loop() {
		key := strconv.Itoa(b.N % 100)
		hm.Set(key, b.N)
	}
}

func BenchmarkHashMapGetWithCollisions(b *testing.B) {
	hm := NewHashMapWithCapacity(1)
	for i := range 100 {
		hm.Set(strconv.Itoa(i), i)
	}

	for b.Loop() {
		key := strconv.Itoa(b.N % 100)
		hm.Get(key)
	}
}

func BenchmarkHashMapResize(b *testing.B) {
	for b.Loop() {
		hm := NewHashMapWithCapacity(4)
		for i := range 20 {
			hm.Set(strconv.Itoa(i), i)
		}
	}
}

func BenchmarkHashMapKeys(b *testing.B) {
	hm := NewHashMap()
	for i := range 1000 {
		hm.Set(strconv.Itoa(i), i)
	}

	for b.Loop() {
		hm.Keys()
	}
}

func BenchmarkHashMapValues(b *testing.B) {
	hm := NewHashMap()
	for i := range 1000 {
		hm.Set(strconv.Itoa(i), i)
	}

	for b.Loop() {
		hm.Values()
	}
}

func BenchmarkHashMapForEach(b *testing.B) {
	hm := NewHashMap()
	for i := range 1000 {
		hm.Set(strconv.Itoa(i), i)
	}

	for b.Loop() {
		hm.ForEach(func(key string, value any) {})
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
