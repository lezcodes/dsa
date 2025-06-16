package bloom_filter

import "testing"

func TestNewBloomFilter(t *testing.T) {
	bf := NewBloomFilter(100, 3)

	if bf.Size() != 100 {
		t.Errorf("Expected size 100, got %d", bf.Size())
	}

	if bf.HashFunctions() != 3 {
		t.Errorf("Expected 3 hash functions, got %d", bf.HashFunctions())
	}
}

func TestBasicOperations(t *testing.T) {
	bf := NewBloomFilter(1000, 3)

	items := []string{"apple", "banana", "cherry"}

	for _, item := range items {
		if bf.Contains(item) {
			t.Errorf("Item %s should not be in empty filter", item)
		}
	}

	for _, item := range items {
		bf.Add(item)
	}

	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("Item %s should be in filter after adding", item)
		}
	}
}

func TestNoFalseNegatives(t *testing.T) {
	bf := NewBloomFilter(100, 4)

	items := []string{"test1", "test2", "test3", "test4", "test5"}

	for _, item := range items {
		bf.Add(item)
	}

	for _, item := range items {
		if !bf.Contains(item) {
			t.Errorf("False negative detected for item %s - this should never happen", item)
		}
	}
}

func TestFalsePositiveRate(t *testing.T) {
	bf := NewBloomFilter(100, 3)

	addedItems := []string{"item1", "item2", "item3", "item4", "item5"}
	testItems := []string{"notadded1", "notadded2", "notadded3", "notadded4", "notadded5",
		"notadded6", "notadded7", "notadded8", "notadded9", "notadded10"}

	for _, item := range addedItems {
		bf.Add(item)
	}

	falsePositives := 0
	for _, item := range testItems {
		if bf.Contains(item) {
			falsePositives++
		}
	}

	falsePositiveRate := float64(falsePositives) / float64(len(testItems))

	if falsePositiveRate > 0.5 {
		t.Errorf("False positive rate too high: %f", falsePositiveRate)
	}
}

func TestHashFunctions(t *testing.T) {
	bf := NewBloomFilter(100, 3)

	data := []byte("test")
	hashes := bf.getHashes(data)

	if len(hashes) != 3 {
		t.Errorf("Expected 3 hashes, got %d", len(hashes))
	}

	for i, hash := range hashes {
		if hash >= bf.size {
			t.Errorf("Hash %d out of bounds: %d >= %d", i, hash, bf.size)
		}
	}

	for i := range hashes {
		for j := i + 1; j < len(hashes); j++ {
			if hashes[i] == hashes[j] {
				t.Errorf("Hash collision detected: hashes[%d] == hashes[%d] == %d", i, j, hashes[i])
			}
		}
	}
}

func TestEstimatedCount(t *testing.T) {
	bf := NewBloomFilter(1000, 3)

	if bf.EstimatedCount() != 0 {
		t.Errorf("Expected 0 estimated count for empty filter, got %d", bf.EstimatedCount())
	}

	items := []string{"a", "b", "c", "d", "e"}
	for _, item := range items {
		bf.Add(item)
	}

	estimated := bf.EstimatedCount()
	if estimated == 0 {
		t.Errorf("Expected non-zero estimated count after adding items")
	}
}

func TestMultipleHashFunctions(t *testing.T) {
	testCases := []int{1, 2, 3, 5, 7}

	for _, k := range testCases {
		bf := NewBloomFilter(100, k)

		bf.Add("test")

		if !bf.Contains("test") {
			t.Errorf("Item should be found with %d hash functions", k)
		}

		if bf.HashFunctions() != k {
			t.Errorf("Expected %d hash functions, got %d", k, bf.HashFunctions())
		}
	}
}

func TestRun(t *testing.T) {
	result := Run()
	if result == nil {
		t.Error("Expected non-nil result")
	}

	results, ok := result.([]TestResult)
	if !ok {
		t.Error("Expected result to be []TestResult")
		return
	}

	if len(results) == 0 {
		t.Error("Expected at least one result")
	}

	for i, res := range results {
		if res.FalseNegatives > 0 {
			t.Errorf("Result %d: Bloom filter should never have false negatives, got %d", i, res.FalseNegatives)
		}

		if res.TruePositives+res.FalsePositives+res.TrueNegatives+res.FalseNegatives != len(res.TestItems) {
			t.Errorf("Result %d: Test counts don't add up correctly", i)
		}

		if res.BloomFilter.Size() == 0 {
			t.Errorf("Result %d: Bloom filter size should not be 0", i)
		}

		if res.BloomFilter.HashFunctions() == 0 {
			t.Errorf("Result %d: Hash functions count should not be 0", i)
		}
	}
}

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		Run()
	}
}

func BenchmarkAdd(b *testing.B) {
	bf := NewBloomFilter(10000, 5)
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	b.ResetTimer()
	for b.Loop() {
		for _, item := range items {
			bf.Add(item)
		}
	}
}

func BenchmarkContains(b *testing.B) {
	bf := NewBloomFilter(10000, 5)
	items := []string{"item1", "item2", "item3", "item4", "item5"}

	for _, item := range items {
		bf.Add(item)
	}

	b.ResetTimer()
	for b.Loop() {
		for _, item := range items {
			bf.Contains(item)
		}
	}
}
