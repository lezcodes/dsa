package bloom_filter

import (
	"crypto/md5"
	"encoding/binary"
	"hash/fnv"
)

type BloomFilter struct {
	bitArray      []bool
	size          uint
	hashFunctions int
}

func NewBloomFilter(size uint, hashFunctions int) *BloomFilter {
	return &BloomFilter{
		bitArray:      make([]bool, size),
		size:          size,
		hashFunctions: hashFunctions,
	}
}

func (bf *BloomFilter) hash1(data []byte) uint {
	h := fnv.New32a()
	h.Write(data)
	return uint(h.Sum32()) % bf.size
}

func (bf *BloomFilter) hash2(data []byte) uint {
	h := md5.Sum(data)
	return uint(binary.BigEndian.Uint32(h[:])) % bf.size
}

func (bf *BloomFilter) hash3(data []byte) uint {
	hash := uint(0)
	for _, b := range data {
		hash = hash*31 + uint(b)
	}
	return hash % bf.size
}

func (bf *BloomFilter) getHashes(data []byte) []uint {
	hashes := make([]uint, bf.hashFunctions)

	if bf.hashFunctions >= 1 {
		hashes[0] = bf.hash1(data)
	}
	if bf.hashFunctions >= 2 {
		hashes[1] = bf.hash2(data)
	}
	if bf.hashFunctions >= 3 {
		hashes[2] = bf.hash3(data)
	}

	for i := 3; i < bf.hashFunctions; i++ {
		combined := hashes[0] + uint(i)*hashes[1] + uint(i*i)*hashes[2]
		hashes[i] = combined % bf.size
	}

	return hashes
}

func (bf *BloomFilter) Add(item string) {
	data := []byte(item)
	hashes := bf.getHashes(data)

	for _, hash := range hashes {
		bf.bitArray[hash] = true
	}
}

func (bf *BloomFilter) Contains(item string) bool {
	data := []byte(item)
	hashes := bf.getHashes(data)

	for _, hash := range hashes {
		if !bf.bitArray[hash] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) Size() uint {
	return bf.size
}

func (bf *BloomFilter) HashFunctions() int {
	return bf.hashFunctions
}

func (bf *BloomFilter) EstimatedCount() uint {
	setBits := uint(0)
	for _, bit := range bf.bitArray {
		if bit {
			setBits++
		}
	}

	if setBits == 0 {
		return 0
	}

	m := float64(bf.size)
	k := float64(bf.hashFunctions)
	x := float64(setBits)

	return uint(-m / k * float64(ln(1.0-x/m)))
}

func ln(x float64) float64 {
	if x <= 0 {
		return 0
	}

	result := 0.0
	term := (x - 1) / (x + 1)
	termSquared := term * term
	currentTerm := term

	for i := range 100 {
		result += currentTerm / float64(2*i+1)
		currentTerm *= termSquared
		if currentTerm < 1e-15 {
			break
		}
	}

	return 2 * result
}

type TestResult struct {
	BloomFilter    *BloomFilter
	AddedItems     []string
	TestItems      []string
	TruePositives  int
	FalsePositives int
	TrueNegatives  int
	FalseNegatives int
}

func Run() any {
	results := []TestResult{}

	testCases := []struct {
		size          uint
		hashFunctions int
		items         []string
		testItems     []string
	}{
		{
			size:          100,
			hashFunctions: 3,
			items:         []string{"apple", "banana", "cherry", "date", "elderberry"},
			testItems:     []string{"apple", "banana", "grape", "kiwi", "cherry", "mango"},
		},
		{
			size:          1000,
			hashFunctions: 5,
			items:         []string{"go", "rust", "python", "java", "javascript", "c++", "typescript"},
			testItems:     []string{"go", "rust", "scala", "kotlin", "python", "swift", "dart"},
		},
		{
			size:          50,
			hashFunctions: 2,
			items:         []string{"red", "blue", "green"},
			testItems:     []string{"red", "yellow", "blue", "purple", "orange", "green"},
		},
	}

	for _, tc := range testCases {
		bf := NewBloomFilter(tc.size, tc.hashFunctions)

		for _, item := range tc.items {
			bf.Add(item)
		}

		result := TestResult{
			BloomFilter: bf,
			AddedItems:  tc.items,
			TestItems:   tc.testItems,
		}

		addedSet := make(map[string]bool)
		for _, item := range tc.items {
			addedSet[item] = true
		}

		for _, testItem := range tc.testItems {
			contains := bf.Contains(testItem)
			actuallyAdded := addedSet[testItem]

			if contains && actuallyAdded {
				result.TruePositives++
			} else if contains && !actuallyAdded {
				result.FalsePositives++
			} else if !contains && !actuallyAdded {
				result.TrueNegatives++
			} else {
				result.FalseNegatives++
			}
		}

		results = append(results, result)
	}

	return results
}
