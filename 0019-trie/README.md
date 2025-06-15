# Trie (Prefix Tree)

## Description

Implementation of a Trie (prefix tree) data structure optimized for autocomplete functionality. A Trie is a tree-like data structure that stores strings in a way that allows for efficient prefix-based operations. Each node represents a character, and paths from root to leaf represent complete words.

## Features

- **Insert**: Add words to the trie
- **Search**: Check if a complete word exists
- **StartsWith**: Check if any word starts with given prefix
- **AutoComplete**: Get limited suggestions for a prefix
- **GetWordsWithPrefix**: Get all words starting with prefix
- **Delete**: Remove words from the trie
- **Case Insensitive**: Handles mixed case input
- **Utility Methods**: Size, IsEmpty, LongestCommonPrefix

## Trie Structure

```
Example Trie for words: ["app", "apple", "application"]

       root
        |
        a
        |
        p
        |
        p (end: "app")
        |
        l
       / \
      e   i
      |   |
      *   c
         / \
        a   ...
        |
        t
        |
        i
        |
        o
        |
        n (end: "application")
```

## Operations

### Insert

- Traverse character by character
- Create nodes as needed
- Mark end of word
- **Time**: O(m), **Space**: O(m) where m = word length

### Search

- Traverse character by character
- Check if path exists and ends at word
- **Time**: O(m), **Space**: O(1)

### Prefix Operations

- Navigate to prefix end
- Collect all words in subtree
- **Time**: O(p + n) where p = prefix length, n = results
- **Space**: O(n) for results

### AutoComplete

- Find prefix subtree
- Collect words with limit
- Sort results alphabetically
- **Time**: O(p + k log k) where k = suggestions
- **Space**: O(k)

## Complexity

### Time Complexity

- **Insert**: O(m) where m = word length
- **Search**: O(m) where m = word length
- **StartsWith**: O(p) where p = prefix length
- **GetWordsWithPrefix**: O(p + n) where n = number of results
- **AutoComplete**: O(p + k log k) where k = max suggestions
- **Delete**: O(m) where m = word length

### Space Complexity

- **Storage**: O(ALPHABET_SIZE × N × M) worst case
- **Average**: Much better due to shared prefixes
- **Operations**: O(m) recursion depth for word operations

## Use Cases

### Autocomplete Systems

- Search engines (Google, Bing)
- IDE code completion
- Mobile keyboard suggestions
- E-commerce product search

### Spell Checkers

- Word validation
- Suggestion generation
- Dictionary lookups

### IP Routing

- Longest prefix matching
- Network routing tables
- CIDR block management

### Text Processing

- Word frequency analysis
- Prefix-based filtering
- Natural language processing

## Advantages

- **Efficient Prefix Operations**: O(p) for prefix queries
- **Memory Efficient**: Shared prefixes save space
- **Sorted Output**: Natural alphabetical ordering
- **Flexible**: Easy to extend with additional features

## Disadvantages

- **Memory Overhead**: Pointer storage for sparse tries
- **Cache Performance**: Poor locality for large alphabets
- **Implementation Complexity**: More complex than hash tables

## Real-World Applications

### Search Engines

- Query autocompletion
- "Did you mean?" suggestions
- Related search terms

### Mobile Applications

- Contact name search
- App name filtering
- Predictive text input

### Development Tools

- Variable name completion
- API method suggestions
- File path completion

## Usage

```bash
make run NAME=0019-trie
```

## Testing

```bash
make test NAME=0019-trie
```
