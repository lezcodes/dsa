package trie

import (
	"reflect"
	"testing"
)

func TestTrieInsertAndSearch(t *testing.T) {
	trie := NewTrie()

	words := []string{"apple", "app", "application", "apply"}
	for _, word := range words {
		trie.Insert(word)
	}

	if trie.Size() != 4 {
		t.Errorf("Expected size 4, got %d", trie.Size())
	}

	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Expected to find word: %s", word)
		}
	}

	if trie.Search("ap") {
		t.Error("Should not find incomplete word 'ap'")
	}

	if trie.Search("applications") {
		t.Error("Should not find non-existent word 'applications'")
	}
}

func TestTrieStartsWith(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "banana"}

	for _, word := range words {
		trie.Insert(word)
	}

	tests := []struct {
		prefix   string
		expected bool
	}{
		{"app", true},
		{"appl", true},
		{"ban", true},
		{"xyz", false},
		{"", true},
		{"application", true},
		{"applications", false},
	}

	for _, tt := range tests {
		result := trie.StartsWith(tt.prefix)
		if result != tt.expected {
			t.Errorf("StartsWith(%s) = %v, want %v", tt.prefix, result, tt.expected)
		}
	}
}

func TestTrieGetWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "appreciate", "banana"}

	for _, word := range words {
		trie.Insert(word)
	}

	appWords := trie.GetWordsWithPrefix("app")
	expected := []string{"app", "apple", "application", "apply", "appreciate"}

	if !reflect.DeepEqual(appWords, expected) {
		t.Errorf("GetWordsWithPrefix('app') = %v, want %v", appWords, expected)
	}

	banWords := trie.GetWordsWithPrefix("ban")
	expectedBan := []string{"banana"}

	if !reflect.DeepEqual(banWords, expectedBan) {
		t.Errorf("GetWordsWithPrefix('ban') = %v, want %v", banWords, expectedBan)
	}

	noWords := trie.GetWordsWithPrefix("xyz")
	if len(noWords) != 0 {
		t.Errorf("Expected empty slice for non-existent prefix, got %v", noWords)
	}
}

func TestTrieAutoComplete(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "appreciate", "approach"}

	for _, word := range words {
		trie.Insert(word)
	}

	suggestions := trie.AutoComplete("app", 3)
	if len(suggestions) != 3 {
		t.Errorf("Expected 3 suggestions, got %d", len(suggestions))
	}

	allSuggestions := trie.AutoComplete("app", 10)
	expected := []string{"app", "apple", "application", "apply", "appreciate", "approach"}

	if !reflect.DeepEqual(allSuggestions, expected) {
		t.Errorf("AutoComplete('app', 10) = %v, want %v", allSuggestions, expected)
	}
}

func TestTrieDelete(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application"}

	for _, word := range words {
		trie.Insert(word)
	}

	if !trie.Delete("app") {
		t.Error("Should successfully delete 'app'")
	}

	if trie.Search("app") {
		t.Error("Should not find 'app' after deletion")
	}

	if !trie.Search("apple") {
		t.Error("Should still find 'apple' after deleting 'app'")
	}

	if !trie.Search("application") {
		t.Error("Should still find 'application' after deleting 'app'")
	}

	if trie.Delete("nonexistent") {
		t.Error("Should not delete non-existent word")
	}

	if trie.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", trie.Size())
	}
}

func TestTrieEmptyOperations(t *testing.T) {
	trie := NewTrie()

	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}

	if trie.Size() != 0 {
		t.Error("New trie size should be 0")
	}

	if trie.Search("anything") {
		t.Error("Empty trie should not contain any words")
	}

	if !trie.StartsWith("") {
		t.Error("Empty prefix should return true")
	}

	words := trie.GetWordsWithPrefix("any")
	if len(words) != 0 {
		t.Error("Empty trie should return no words for any prefix")
	}
}

func TestTrieCaseInsensitive(t *testing.T) {
	trie := NewTrie()

	trie.Insert("Apple")
	trie.Insert("BANANA")
	trie.Insert("CaR")

	if !trie.Search("apple") {
		t.Error("Should find 'apple' (case insensitive)")
	}

	if !trie.Search("APPLE") {
		t.Error("Should find 'APPLE' (case insensitive)")
	}

	if !trie.Search("banana") {
		t.Error("Should find 'banana' (case insensitive)")
	}

	if !trie.Search("car") {
		t.Error("Should find 'car' (case insensitive)")
	}
}

func TestTrieDuplicateInsert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")
	trie.Insert("apple")
	trie.Insert("APPLE")

	if trie.Size() != 1 {
		t.Errorf("Expected size 1 after duplicate inserts, got %d", trie.Size())
	}
}

func TestTrieLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{
			name:     "Common prefix exists",
			words:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "No common prefix",
			words:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "Single word",
			words:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Empty trie",
			words:    []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range tt.words {
				trie.Insert(word)
			}

			result := trie.LongestCommonPrefix()
			if result != tt.expected {
				t.Errorf("LongestCommonPrefix() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestTrieCountWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "banana", "band"}

	for _, word := range words {
		trie.Insert(word)
	}

	if count := trie.CountWordsWithPrefix("app"); count != 4 {
		t.Errorf("Expected 4 words with prefix 'app', got %d", count)
	}

	if count := trie.CountWordsWithPrefix("ban"); count != 2 {
		t.Errorf("Expected 2 words with prefix 'ban', got %d", count)
	}

	if count := trie.CountWordsWithPrefix("xyz"); count != 0 {
		t.Errorf("Expected 0 words with prefix 'xyz', got %d", count)
	}
}

func TestTrieGetAllWords(t *testing.T) {
	trie := NewTrie()
	words := []string{"banana", "apple", "app"}

	for _, word := range words {
		trie.Insert(word)
	}

	allWords := trie.GetAllWords()
	expected := []string{"app", "apple", "banana"}

	if !reflect.DeepEqual(allWords, expected) {
		t.Errorf("GetAllWords() = %v, want %v", allWords, expected)
	}
}

func TestRun(t *testing.T) {
	result := Run()
	resultMap, ok := result.(map[string]any)
	if !ok {
		t.Fatalf("Expected map[string]any, got %T", result)
	}

	if resultMap["total_words"] != 19 {
		t.Errorf("Expected 19 total words, got %v", resultMap["total_words"])
	}

	if resultMap["search_app"] != true {
		t.Error("Expected search_app to be true")
	}

	if resultMap["search_missing"] != false {
		t.Error("Expected search_missing to be false")
	}

	autocompleteApp, ok := resultMap["autocomplete_app"].([]string)
	if !ok || len(autocompleteApp) != 5 {
		t.Errorf("Expected 5 autocomplete suggestions for 'app', got %v", autocompleteApp)
	}
}

func BenchmarkTrieInsert(b *testing.B) {
	trie := NewTrie()
	words := []string{"apple", "application", "apply", "appreciate", "approach"}

	for b.Loop() {
		for _, word := range words {
			trie.Insert(word)
		}
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	trie := NewTrie()
	words := []string{"apple", "application", "apply", "appreciate", "approach"}

	for _, word := range words {
		trie.Insert(word)
	}

	for b.Loop() {
		for _, word := range words {
			trie.Search(word)
		}
	}
}

func BenchmarkTrieAutoComplete(b *testing.B) {
	trie := NewTrie()
	words := []string{
		"apple", "application", "apply", "appreciate", "approach",
		"banana", "band", "bandana", "bandwidth",
		"cat", "car", "card", "care", "careful",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	for b.Loop() {
		trie.AutoComplete("app", 5)
	}
}
