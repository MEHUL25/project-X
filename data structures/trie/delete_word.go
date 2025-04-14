package trie

// To delete a word:

// We traverse the Trie recursively.
// When we reach the end of the word:
// Unmark the isEnd flag.
// Then check if the current node has any children.
// If it does, we cannot delete the node.

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie struct, holding the root node
type Trie struct {
	root *TrieNode
}

// Delete removes a word from the Trie
func (t *Trie) Delete(word string) {
	t.deleteHelper(t.root, word, 0)
}

// Recursive helper for deletion
func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.isEnd {
			return false // word not found
		}
		node.isEnd = false
		return len(node.children) == 0 // delete this node if it's now a leaf
	}

	ch := rune(word[index])
	child, exists := node.children[ch]
	if !exists {
		return false // word not found
	}

	// Recursively delete the child
	shouldDeleteChild := t.deleteHelper(child, word, index+1)

	if shouldDeleteChild {
		delete(node.children, ch) // remove the mapping

		// If current node is not the end of another word and has no other children
		return !node.isEnd && len(node.children) == 0
	}

	return false
}
