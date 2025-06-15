package lru_cache

type Node struct {
	Key   string
	Value any
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		capacity = 1
	}

	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return &LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[string]*Node),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCache) addToHead(node *Node) {
	node.Prev = lru.head
	node.Next = lru.head.Next
	lru.head.Next.Prev = node
	lru.head.Next = node
}

func (lru *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *Node {
	lastNode := lru.tail.Prev
	lru.removeNode(lastNode)
	return lastNode
}

func (lru *LRUCache) Get(key string) (any, bool) {
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.Value, true
	}
	return nil, false
}

func (lru *LRUCache) Put(key string, value any) {
	if node, exists := lru.cache[key]; exists {
		node.Value = value
		lru.moveToHead(node)
		return
	}

	newNode := &Node{Key: key, Value: value}

	if lru.size >= lru.capacity {
		tail := lru.removeTail()
		delete(lru.cache, tail.Key)
		lru.size--
	}

	lru.addToHead(newNode)
	lru.cache[key] = newNode
	lru.size++
}

func (lru *LRUCache) Delete(key string) bool {
	if node, exists := lru.cache[key]; exists {
		lru.removeNode(node)
		delete(lru.cache, key)
		lru.size--
		return true
	}
	return false
}

func (lru *LRUCache) Has(key string) bool {
	_, exists := lru.cache[key]
	return exists
}

func (lru *LRUCache) Size() int {
	return lru.size
}

func (lru *LRUCache) Capacity() int {
	return lru.capacity
}

func (lru *LRUCache) IsEmpty() bool {
	return lru.size == 0
}

func (lru *LRUCache) IsFull() bool {
	return lru.size >= lru.capacity
}

func (lru *LRUCache) Clear() {
	lru.cache = make(map[string]*Node)
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	lru.size = 0
}

func (lru *LRUCache) Keys() []string {
	keys := make([]string, 0, lru.size)
	current := lru.head.Next
	for current != lru.tail {
		keys = append(keys, current.Key)
		current = current.Next
	}
	return keys
}

func (lru *LRUCache) Values() []any {
	values := make([]any, 0, lru.size)
	current := lru.head.Next
	for current != lru.tail {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func (lru *LRUCache) Entries() []map[string]any {
	entries := make([]map[string]any, 0, lru.size)
	current := lru.head.Next
	for current != lru.tail {
		entry := map[string]any{
			"key":   current.Key,
			"value": current.Value,
		}
		entries = append(entries, entry)
		current = current.Next
	}
	return entries
}

func (lru *LRUCache) GetMostRecentKey() (string, bool) {
	if lru.size == 0 {
		return "", false
	}
	return lru.head.Next.Key, true
}

func (lru *LRUCache) GetLeastRecentKey() (string, bool) {
	if lru.size == 0 {
		return "", false
	}
	return lru.tail.Prev.Key, true
}

func (lru *LRUCache) Peek(key string) (any, bool) {
	if node, exists := lru.cache[key]; exists {
		return node.Value, true
	}
	return nil, false
}

func (lru *LRUCache) ForEach(fn func(key string, value any)) {
	current := lru.head.Next
	for current != lru.tail {
		fn(current.Key, current.Value)
		current = current.Next
	}
}

func (lru *LRUCache) SetCapacity(newCapacity int) {
	if newCapacity <= 0 {
		newCapacity = 1
	}

	lru.capacity = newCapacity

	for lru.size > lru.capacity {
		tail := lru.removeTail()
		delete(lru.cache, tail.Key)
		lru.size--
	}
}

func Run() any {
	cache := NewLRUCache(3)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	result := make(map[string]any)
	result["initialSize"] = cache.Size()
	result["capacity"] = cache.Capacity()
	result["isFull"] = cache.IsFull()

	value1, exists1 := cache.Get("key1")
	result["getKey1"] = map[string]any{"value": value1, "exists": exists1}

	cache.Put("key4", "value4")

	result["sizeAfterEviction"] = cache.Size()

	value2, exists2 := cache.Get("key2")
	result["getKey2AfterEviction"] = map[string]any{"value": value2, "exists": exists2}

	mostRecent, hasMostRecent := cache.GetMostRecentKey()
	result["mostRecentKey"] = map[string]any{"key": mostRecent, "exists": hasMostRecent}

	leastRecent, hasLeastRecent := cache.GetLeastRecentKey()
	result["leastRecentKey"] = map[string]any{"key": leastRecent, "exists": hasLeastRecent}

	result["keys"] = cache.Keys()

	data := make(map[string]any)
	cache.ForEach(func(key string, value any) {
		data[key] = value
	})
	result["data"] = data

	return result
}
