package trie

import "fmt"

// Each node represents a character.
// The root node is empty ("").
// Words are inserted character by character.
// Every path from the root to a leaf represents a word.
// Nodes can indicate if they mark the end of a word.

// Autocomplete systems
// Spell checkers
// IP routing (longest prefix match)
// Dictionary/word search

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie struct, holding the root node
type Trie struct {
	root *TrieNode
}

// NewTrie creates and returns a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, ch := range word {
		if _, exists := current.children[ch]; !exists {
			current.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		current = current.children[ch]
	}
	current.isEnd = true
}

// Search returns true if the word is in the Trie
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, ch := range word {
		node, exists := current.children[ch]
		if !exists {
			return false
		}
		current = node
	}
	return current.isEnd
}

// StartsWith returns true if there is any word in the Trie with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, ch := range prefix {
		node, exists := current.children[ch]
		if !exists {
			return false
		}
		current = node
	}
	return true
}

func createTrie() {
	trie := NewTrie()

	words := []string{"go", "gopher", "golang", "game", "good"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println("Search 'golang':", trie.Search("golang"))   // true
	fmt.Println("Search 'gone':", trie.Search("gone"))       // false
	fmt.Println("Prefix 'go':", trie.StartsWith("go"))       // true
	fmt.Println("Prefix 'ga':", trie.StartsWith("ga"))       // true
	fmt.Println("Prefix 'zebra':", trie.StartsWith("zebra")) // false
}
