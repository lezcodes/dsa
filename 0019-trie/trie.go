package trie

import (
	"sort"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	word     string
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		word:     "",
	}
}

type Trie struct {
	root *TrieNode
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
		size: 0,
	}
}

func (t *Trie) Insert(word string) {
	if word == "" {
		return
	}

	word = strings.ToLower(word)
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = NewTrieNode()
		}
		current = current.children[char]
	}

	if !current.isEnd {
		current.isEnd = true
		current.word = word
		t.size++
	}
}

func (t *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	word = strings.ToLower(word)
	current := t.root

	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}

	return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	if prefix == "" {
		return true
	}

	prefix = strings.ToLower(prefix)
	current := t.root

	for _, char := range prefix {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}

	return true
}

func (t *Trie) GetWordsWithPrefix(prefix string) []string {
	prefix = strings.ToLower(prefix)
	current := t.root

	for _, char := range prefix {
		if current.children[char] == nil {
			return []string{}
		}
		current = current.children[char]
	}

	var words []string
	t.collectWords(current, &words)
	sort.Strings(words)
	return words
}

func (t *Trie) collectWords(node *TrieNode, words *[]string) {
	if node.isEnd {
		*words = append(*words, node.word)
	}

	for _, child := range node.children {
		t.collectWords(child, words)
	}
}

func (t *Trie) AutoComplete(prefix string, maxSuggestions int) []string {
	words := t.GetWordsWithPrefix(prefix)
	if len(words) > maxSuggestions {
		return words[:maxSuggestions]
	}
	return words
}

func (t *Trie) Delete(word string) bool {
	if word == "" {
		return false
	}

	word = strings.ToLower(word)

	if !t.Search(word) {
		return false
	}

	runes := []rune(word)
	t.deleteHelper(t.root, runes, 0)
	return true
}

func (t *Trie) deleteHelper(node *TrieNode, runes []rune, index int) bool {
	if index == len(runes) {
		if !node.isEnd {
			return false
		}
		node.isEnd = false
		node.word = ""
		t.size--
		return len(node.children) == 0
	}

	char := runes[index]
	child := node.children[char]
	if child == nil {
		return false
	}

	shouldDeleteChild := t.deleteHelper(child, runes, index+1)

	if shouldDeleteChild {
		delete(node.children, char)
		return !node.isEnd && len(node.children) == 0
	}

	return false
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie) GetAllWords() []string {
	var words []string
	t.collectWords(t.root, &words)
	sort.Strings(words)
	return words
}

func (t *Trie) LongestCommonPrefix() string {
	if t.size == 0 {
		return ""
	}

	var prefix strings.Builder
	current := t.root

	for len(current.children) == 1 && !current.isEnd {
		for char, child := range current.children {
			prefix.WriteRune(char)
			current = child
			break
		}
	}

	return prefix.String()
}

func (t *Trie) CountWordsWithPrefix(prefix string) int {
	return len(t.GetWordsWithPrefix(prefix))
}

func Run() any {
	trie := NewTrie()

	words := []string{
		"apple", "app", "application", "apply", "appreciate",
		"banana", "band", "bandana", "bandwidth",
		"cat", "car", "card", "care", "careful", "careless",
		"dog", "door", "down", "download",
	}

	for _, word := range words {
		trie.Insert(word)
	}

	return map[string]any{
		"total_words":           trie.Size(),
		"search_app":            trie.Search("app"),
		"search_apple":          trie.Search("apple"),
		"search_missing":        trie.Search("missing"),
		"starts_with_app":       trie.StartsWith("app"),
		"starts_with_xyz":       trie.StartsWith("xyz"),
		"autocomplete_app":      trie.AutoComplete("app", 5),
		"autocomplete_car":      trie.AutoComplete("car", 3),
		"autocomplete_ban":      trie.AutoComplete("ban", 10),
		"words_with_prefix_do":  trie.GetWordsWithPrefix("do"),
		"longest_common_prefix": trie.LongestCommonPrefix(),
		"count_app_words":       trie.CountWordsWithPrefix("app"),
		"all_words_sample":      trie.GetAllWords()[:10],
	}
}
