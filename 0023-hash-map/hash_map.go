package hash_map

import (
	"hash/fnv"
)

type KeyValue struct {
	Key   string
	Value any
	Next  *KeyValue
}

type HashMap struct {
	buckets    []*KeyValue
	size       int
	capacity   int
	loadFactor float64
}

const (
	DefaultCapacity   = 16
	DefaultLoadFactor = 0.75
	ResizeFactor      = 2
)

func NewHashMap() *HashMap {
	return &HashMap{
		buckets:    make([]*KeyValue, DefaultCapacity),
		size:       0,
		capacity:   DefaultCapacity,
		loadFactor: DefaultLoadFactor,
	}
}

func NewHashMapWithCapacity(capacity int) *HashMap {
	if capacity < 1 {
		capacity = DefaultCapacity
	}
	return &HashMap{
		buckets:    make([]*KeyValue, capacity),
		size:       0,
		capacity:   capacity,
		loadFactor: DefaultLoadFactor,
	}
}

func (hm *HashMap) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % hm.capacity
}

func (hm *HashMap) Set(key string, value any) {
	if float64(hm.size)/float64(hm.capacity) >= hm.loadFactor {
		hm.resize()
	}

	index := hm.hash(key)
	bucket := hm.buckets[index]

	if bucket == nil {
		hm.buckets[index] = &KeyValue{Key: key, Value: value}
		hm.size++
		return
	}

	current := bucket
	for current != nil {
		if current.Key == key {
			current.Value = value
			return
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	current.Next = &KeyValue{Key: key, Value: value}
	hm.size++
}

func (hm *HashMap) Get(key string) (any, bool) {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	current := bucket
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return nil, false
}

func (hm *HashMap) Delete(key string) bool {
	index := hm.hash(key)
	bucket := hm.buckets[index]

	if bucket == nil {
		return false
	}

	if bucket.Key == key {
		hm.buckets[index] = bucket.Next
		hm.size--
		return true
	}

	current := bucket
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			hm.size--
			return true
		}
		current = current.Next
	}

	return false
}

func (hm *HashMap) Has(key string) bool {
	_, exists := hm.Get(key)
	return exists
}

func (hm *HashMap) Size() int {
	return hm.size
}

func (hm *HashMap) IsEmpty() bool {
	return hm.size == 0
}

func (hm *HashMap) Clear() {
	hm.buckets = make([]*KeyValue, hm.capacity)
	hm.size = 0
}

func (hm *HashMap) Keys() []string {
	keys := make([]string, 0, hm.size)
	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}
	return keys
}

func (hm *HashMap) Values() []any {
	values := make([]any, 0, hm.size)
	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			values = append(values, current.Value)
			current = current.Next
		}
	}
	return values
}

func (hm *HashMap) Entries() []KeyValue {
	entries := make([]KeyValue, 0, hm.size)
	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			entries = append(entries, KeyValue{Key: current.Key, Value: current.Value})
			current = current.Next
		}
	}
	return entries
}

func (hm *HashMap) LoadFactor() float64 {
	return float64(hm.size) / float64(hm.capacity)
}

func (hm *HashMap) Capacity() int {
	return hm.capacity
}

func (hm *HashMap) resize() {
	oldBuckets := hm.buckets
	oldCapacity := hm.capacity

	hm.capacity *= ResizeFactor
	hm.buckets = make([]*KeyValue, hm.capacity)
	hm.size = 0

	for _, bucket := range oldBuckets {
		current := bucket
		for current != nil {
			hm.Set(current.Key, current.Value)
			current = current.Next
		}
	}

	_ = oldCapacity
}

func (hm *HashMap) ForEach(fn func(key string, value any)) {
	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			fn(current.Key, current.Value)
			current = current.Next
		}
	}
}

func (hm *HashMap) GetBucketDistribution() []int {
	distribution := make([]int, hm.capacity)
	for i, bucket := range hm.buckets {
		count := 0
		current := bucket
		for current != nil {
			count++
			current = current.Next
		}
		distribution[i] = count
	}
	return distribution
}

func Run() any {
	hashMap := NewHashMap()

	hashMap.Set("name", "Alice")
	hashMap.Set("age", 30)
	hashMap.Set("city", "New York")
	hashMap.Set("country", "USA")
	hashMap.Set("occupation", "Engineer")

	result := make(map[string]any)
	result["size"] = hashMap.Size()
	result["capacity"] = hashMap.Capacity()
	result["loadFactor"] = hashMap.LoadFactor()

	data := make(map[string]any)
	hashMap.ForEach(func(key string, value any) {
		data[key] = value
	})
	result["data"] = data

	name, exists := hashMap.Get("name")
	result["getName"] = map[string]any{"value": name, "exists": exists}

	deleted := hashMap.Delete("age")
	result["deleteAge"] = deleted
	result["sizeAfterDelete"] = hashMap.Size()

	result["keys"] = hashMap.Keys()

	return result
}
