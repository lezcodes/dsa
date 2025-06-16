package hash_map

type KeyValue struct {
	Key   string
	Value any
	Next  *KeyValue
}

type HashMap struct {
	buckets map[string]any
}

const (
	DefaultCapacity   = 16
	DefaultLoadFactor = 0.75
	ResizeFactor      = 2
)

func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make(map[string]any),
	}
}

func (hm *HashMap) Set(key string, value any) {
	hm.buckets[key] = value
}

func (hm *HashMap) Get(key string) (any, bool) {
	value, exists := hm.buckets[key]
	return value, exists
}

func (hm *HashMap) Delete(key string) bool {
	if _, exists := hm.buckets[key]; exists {
		delete(hm.buckets, key)
		return true
	}
	return false
}

func (hm *HashMap) Has(key string) bool {
	_, exists := hm.buckets[key]
	return exists
}

func (hm *HashMap) Size() int {
	return len(hm.buckets)
}

func (hm *HashMap) Keys() []string {
	keys := make([]string, 0, len(hm.buckets))
	for key := range hm.buckets {
		keys = append(keys, key)
	}
	return keys
}

func Run() any {
	hm := NewHashMap()

	hm.Set("name", "Alice")
	hm.Set("age", 30)
	hm.Set("city", "New York")

	name, _ := hm.Get("name")
	exists := hm.Has("age")
	hm.Delete("city")
	keys := hm.Keys()

	return map[string]any{
		"name":   name,
		"exists": exists,
		"size":   hm.Size(),
		"keys":   keys,
	}
}
