package lru_cache

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

	if capacity, exists := resultMap["capacity"]; !exists || capacity != 3 {
		t.Errorf("Expected capacity to be 3, got %v", capacity)
	}

	if size, exists := resultMap["sizeAfterEviction"]; !exists || size != 3 {
		t.Errorf("Expected size to be 3 after eviction, got %v", size)
	}

	if isFull, exists := resultMap["isFull"]; !exists || isFull != true {
		t.Errorf("Expected cache to be full, got %v", isFull)
	}
}

func TestNewLRUCache(t *testing.T) {
	cache := NewLRUCache(5)
	if cache == nil {
		t.Error("Expected non-nil LRUCache")
	}
	if cache.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Capacity())
	}
	if cache.Size() != 0 {
		t.Errorf("Expected size 0, got %d", cache.Size())
	}
	if !cache.IsEmpty() {
		t.Error("Expected empty cache")
	}
	if cache.IsFull() {
		t.Error("Expected cache not to be full")
	}

	cache = NewLRUCache(0)
	if cache.Capacity() != 1 {
		t.Errorf("Expected minimum capacity 1 for invalid input, got %d", cache.Capacity())
	}

	cache = NewLRUCache(-5)
	if cache.Capacity() != 1 {
		t.Errorf("Expected minimum capacity 1 for negative input, got %d", cache.Capacity())
	}
}

func TestPutAndGet(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("key1", "value1")
	value, exists := cache.Get("key1")
	if !exists {
		t.Error("Expected key1 to exist")
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}

	cache.Put("key2", 42)
	value, exists = cache.Get("key2")
	if !exists {
		t.Error("Expected key2 to exist")
	}
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}

	cache.Put("key1", "updated_value")
	value, exists = cache.Get("key1")
	if !exists {
		t.Error("Expected key1 to still exist after update")
	}
	if value != "updated_value" {
		t.Errorf("Expected updated_value, got %v", value)
	}
	if cache.Size() != 2 {
		t.Errorf("Expected size 2 after update, got %d", cache.Size())
	}
}

func TestGetNonExistent(t *testing.T) {
	cache := NewLRUCache(3)
	value, exists := cache.Get("nonexistent")
	if exists {
		t.Error("Expected nonexistent key to not exist")
	}
	if value != nil {
		t.Errorf("Expected nil value for nonexistent key, got %v", value)
	}
}

func TestLRUEviction(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	if cache.Size() != 3 {
		t.Errorf("Expected size 3, got %d", cache.Size())
	}
	if !cache.IsFull() {
		t.Error("Expected cache to be full")
	}

	cache.Put("key4", "value4")

	if cache.Size() != 3 {
		t.Errorf("Expected size to remain 3 after eviction, got %d", cache.Size())
	}

	_, exists := cache.Get("key1")
	if exists {
		t.Error("Expected key1 to be evicted (least recently used)")
	}

	_, exists = cache.Get("key2")
	if !exists {
		t.Error("Expected key2 to still exist")
	}

	_, exists = cache.Get("key3")
	if !exists {
		t.Error("Expected key3 to still exist")
	}

	_, exists = cache.Get("key4")
	if !exists {
		t.Error("Expected key4 to exist (newly added)")
	}
}

func TestAccessOrderUpdate(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	cache.Get("key1")

	cache.Put("key4", "value4")

	_, exists := cache.Get("key2")
	if exists {
		t.Error("Expected key2 to be evicted (least recently used after key1 access)")
	}

	_, exists = cache.Get("key1")
	if !exists {
		t.Error("Expected key1 to still exist (recently accessed)")
	}
}

func TestDelete(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	if cache.Size() != 3 {
		t.Errorf("Expected size 3, got %d", cache.Size())
	}

	deleted := cache.Delete("key2")
	if !deleted {
		t.Error("Expected key2 to be deleted")
	}
	if cache.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", cache.Size())
	}

	_, exists := cache.Get("key2")
	if exists {
		t.Error("Expected key2 to not exist after deletion")
	}

	deleted = cache.Delete("nonexistent")
	if deleted {
		t.Error("Expected deletion of nonexistent key to return false")
	}
	if cache.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", cache.Size())
	}
}

func TestHas(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")

	if !cache.Has("key1") {
		t.Error("Expected key1 to exist")
	}
	if cache.Has("nonexistent") {
		t.Error("Expected nonexistent key to not exist")
	}
}

func TestClear(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	if cache.Size() != 3 {
		t.Errorf("Expected size 3, got %d", cache.Size())
	}

	cache.Clear()
	if cache.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", cache.Size())
	}
	if !cache.IsEmpty() {
		t.Error("Expected empty cache after clear")
	}

	_, exists := cache.Get("key1")
	if exists {
		t.Error("Expected key1 to not exist after clear")
	}
}

func TestKeys(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	keys := cache.Keys()
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	if keys[0] != "key3" || keys[1] != "key2" || keys[2] != "key1" {
		t.Errorf("Expected keys in MRU order [key3, key2, key1], got %v", keys)
	}
}

func TestValues(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	values := cache.Values()
	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	if values[0] != "value3" || values[1] != "value2" || values[2] != "value1" {
		t.Errorf("Expected values in MRU order [value3, value2, value1], got %v", values)
	}
}

func TestEntries(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	entries := cache.Entries()
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0]["key"] != "key2" || entries[0]["value"] != "value2" {
		t.Errorf("Expected first entry to be key2->value2, got %v", entries[0])
	}
	if entries[1]["key"] != "key1" || entries[1]["value"] != "value1" {
		t.Errorf("Expected second entry to be key1->value1, got %v", entries[1])
	}
}

func TestForEach(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	visited := make([]string, 0)
	cache.ForEach(func(key string, value any) {
		visited = append(visited, key)
	})

	if len(visited) != 3 {
		t.Errorf("Expected 3 entries visited, got %d", len(visited))
	}

	expectedOrder := []string{"key3", "key2", "key1"}
	for i, expectedKey := range expectedOrder {
		if visited[i] != expectedKey {
			t.Errorf("Expected key %s at position %d, got %s", expectedKey, i, visited[i])
		}
	}
}

func TestMostAndLeastRecentKeys(t *testing.T) {
	cache := NewLRUCache(3)

	mostRecent, exists := cache.GetMostRecentKey()
	if exists {
		t.Error("Expected no most recent key in empty cache")
	}

	leastRecent, exists := cache.GetLeastRecentKey()
	if exists {
		t.Error("Expected no least recent key in empty cache")
	}

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	mostRecent, exists = cache.GetMostRecentKey()
	if !exists || mostRecent != "key3" {
		t.Errorf("Expected most recent key to be key3, got %s", mostRecent)
	}

	leastRecent, exists = cache.GetLeastRecentKey()
	if !exists || leastRecent != "key1" {
		t.Errorf("Expected least recent key to be key1, got %s", leastRecent)
	}

	cache.Get("key1")

	mostRecent, exists = cache.GetMostRecentKey()
	if !exists || mostRecent != "key1" {
		t.Errorf("Expected most recent key to be key1 after access, got %s", mostRecent)
	}

	leastRecent, exists = cache.GetLeastRecentKey()
	if !exists || leastRecent != "key2" {
		t.Errorf("Expected least recent key to be key2 after key1 access, got %s", leastRecent)
	}
}

func TestPeek(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	value, exists := cache.Peek("key1")
	if !exists || value != "value1" {
		t.Errorf("Expected peek to return value1, got %v", value)
	}

	mostRecent, _ := cache.GetMostRecentKey()
	if mostRecent != "key2" {
		t.Errorf("Expected most recent key to remain key2 after peek, got %s", mostRecent)
	}

	value, exists = cache.Peek("nonexistent")
	if exists {
		t.Error("Expected peek of nonexistent key to return false")
	}
	if value != nil {
		t.Errorf("Expected nil value for nonexistent key peek, got %v", value)
	}
}

func TestSetCapacity(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")
	cache.Put("key4", "value4")

	if cache.Size() != 4 {
		t.Errorf("Expected size 4, got %d", cache.Size())
	}

	cache.SetCapacity(2)
	if cache.Capacity() != 2 {
		t.Errorf("Expected capacity 2, got %d", cache.Capacity())
	}
	if cache.Size() != 2 {
		t.Errorf("Expected size 2 after capacity reduction, got %d", cache.Size())
	}

	_, exists := cache.Get("key1")
	if exists {
		t.Error("Expected key1 to be evicted after capacity reduction")
	}
	_, exists = cache.Get("key2")
	if exists {
		t.Error("Expected key2 to be evicted after capacity reduction")
	}
	_, exists = cache.Get("key3")
	if !exists {
		t.Error("Expected key3 to remain after capacity reduction")
	}
	_, exists = cache.Get("key4")
	if !exists {
		t.Error("Expected key4 to remain after capacity reduction")
	}

	cache.SetCapacity(0)
	if cache.Capacity() != 1 {
		t.Errorf("Expected minimum capacity 1 for invalid input, got %d", cache.Capacity())
	}
}

func TestEmptyCacheOperations(t *testing.T) {
	cache := NewLRUCache(3)

	if !cache.IsEmpty() {
		t.Error("Expected new cache to be empty")
	}

	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected empty keys slice, got %d keys", len(keys))
	}

	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected empty values slice, got %d values", len(values))
	}

	entries := cache.Entries()
	if len(entries) != 0 {
		t.Errorf("Expected empty entries slice, got %d entries", len(entries))
	}

	cache.ForEach(func(key string, value any) {
		t.Error("ForEach should not be called on empty cache")
	})
}

func TestLargeCapacity(t *testing.T) {
	cache := NewLRUCache(1000)
	numItems := 500

	for i := range numItems {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		cache.Put(key, value)
	}

	if cache.Size() != numItems {
		t.Errorf("Expected size %d, got %d", numItems, cache.Size())
	}

	for i := range numItems {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		value, exists := cache.Get(key)
		if !exists {
			t.Errorf("Expected %s to exist", key)
		}
		if value != expectedValue {
			t.Errorf("Expected %s -> %s, got %v", key, expectedValue, value)
		}
	}
}

func TestEvictionOrder(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put("A", 1)
	cache.Put("B", 2)
	cache.Put("C", 3)

	cache.Get("A")
	cache.Get("B")

	cache.Put("D", 4)

	_, exists := cache.Get("C")
	if exists {
		t.Error("Expected C to be evicted (least recently used)")
	}

	keys := cache.Keys()
	expectedOrder := []string{"D", "B", "A"}
	for i, expectedKey := range expectedOrder {
		if keys[i] != expectedKey {
			t.Errorf("Expected key %s at position %d, got %s", expectedKey, i, keys[i])
		}
	}
}

func BenchmarkLRUCachePut(b *testing.B) {
	cache := NewLRUCache(1000)
	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		cache.Put(key, b.N)
	}
}

func BenchmarkLRUCacheGet(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := range 1000 {
		cache.Put(strconv.Itoa(i), i)
	}

	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		cache.Get(key)
	}
}

func BenchmarkLRUCacheDelete(b *testing.B) {
	cache := NewLRUCache(b.N)
	for i := range b.N {
		cache.Put(strconv.Itoa(i), i)
	}

	b.ResetTimer()
	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		cache.Delete(key)
	}
}

func BenchmarkLRUCacheEviction(b *testing.B) {
	cache := NewLRUCache(100)
	for b.Loop() {
		key := strconv.Itoa(b.N)
		cache.Put(key, b.N)
	}
}

func BenchmarkLRUCachePeek(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := range 1000 {
		cache.Put(strconv.Itoa(i), i)
	}

	for b.Loop() {
		key := strconv.Itoa(b.N % 1000)
		cache.Peek(key)
	}
}

func BenchmarkLRUCacheKeys(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := range 1000 {
		cache.Put(strconv.Itoa(i), i)
	}

	for b.Loop() {
		cache.Keys()
	}
}

func BenchmarkLRUCacheForEach(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := range 1000 {
		cache.Put(strconv.Itoa(i), i)
	}

	for b.Loop() {
		cache.ForEach(func(key string, value any) {})
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}
