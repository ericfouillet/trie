// Go trie data structure

package trie

/*
 * Implementation of the Trie interface.
 */

import "fmt"
import "unicode/utf8"
import "container/list"

// New creates a new trie.
func New(f NodeDataExtractorFunc) *Trie {
	return &Trie{root: newNode(), DataExtractor: NodeDataExtractor(f)}
}

func newNode() *node {
	return &node{} // use zero values for all fields
}

// AddWord adds a new word to a trie.
func (t *Trie) AddWord(s string) error {
	return t.root.addWordToNode(s, t.DataExtractor)
}

// Contains returns true if a trie contains the given word.
func (t *Trie) Contains(s string) bool {
	return t.root.hasPrefix(s, t.DataExtractor)
}

func (t *node) addWordToNode(s string, e NodeDataExtractor) error {
	if len(s) == 0 {
		return nil
	}
	w, n, err := e.ExtractNextData(s)
	if err != nil {
		return fmt.Errorf("Could not decode string %v: %v", s, err)
	}
	if t.children == nil {
		t.children = make(map[interface{}]*node)
	}
	if _, ok := t.children[w]; !ok {
		t.children[w] = newNode()
	}
	t.children[w].data = w
	if len(s) == n {
		t.children[w].isFullWord = true
		return nil
	}
	return t.children[w].addWordToNode(s[n:], e)
}

func (t *node) hasPrefix(s string, e NodeDataExtractor) bool {
	if s == "" {
		return false
	}
	w, n, err := e.ExtractNextData(s)
	if err != nil {
		return false
	}
	if _, ok := t.children[w]; !ok {
		return false
	}
	if len(s) == n {
		return true
	}
	return t.children[w].hasPrefix(s[n:], e)
}

// ExtractNextRuneElement extracts the next RuneData to process from a string.
func ExtractNextRuneElement(s string) (interface{}, int, error) {
	ru, n := utf8.DecodeRuneInString(s)
	if ru == utf8.RuneError {
		return utf8.MaxRune, 0, fmt.Errorf("Could not decode rune in string %v", s)
	}
	return ru, n, nil
}

// NewLinked creates a new trie using a linked list to store children.
func NewLinked(f NodeDataExtractorFunc) *LinkedTrie {
	return &LinkedTrie{root: newLinkedNode(), DataExtractor: NodeDataExtractor(f)}
}

func newLinkedNode() *linkedNode {
	return &linkedNode{} // use zero values for all fields
}

// AddWord adds a new word to a linked trie.
func (t *LinkedTrie) AddWord(s string) error {
	return t.root.addWordToNode(s, t.DataExtractor)
}

// Contains returns true if a linked trie contains the given word.
func (t *LinkedTrie) Contains(s string) bool {
	return t.root.hasPrefix(s, t.DataExtractor)
}

func (t *linkedNode) addWordToNode(s string, e NodeDataExtractor) error {
	if len(s) == 0 {
		return nil
	}
	w, n, err := e.ExtractNextData(s)
	if err != nil {
		return fmt.Errorf("Could not decode string %v: %v", s, err)
	}
	if t.children == nil {
		t.children = list.New()
	}
	var found *linkedNode
	for e := t.children.Front(); e != nil; e = e.Next() {
		if child, ok := e.Value.(*linkedNode); ok && child.data == w {
			found = child
			break
		}
	}
	if found == nil {
		found = newLinkedNode()
		found.data = w
		t.children.PushBack(found)
	}
	if len(s) == n {
		found.isFullWord = true
		return nil
	}
	return found.addWordToNode(s[n:], e)
}

func (t *linkedNode) hasPrefix(s string, e NodeDataExtractor) bool {
	if s == "" {
		return false
	}
	w, n, err := e.ExtractNextData(s)
	if err != nil {
		return false
	}
	var found *linkedNode
	for e := t.children.Front(); e != nil; e = e.Next() {
		if child, ok := e.Value.(*linkedNode); ok && child.data == w {
			found = child
			break
		}
	}
	if found == nil {
		return false
	}
	if len(s) == n {
		return true
	}
	return found.hasPrefix(s[n:], e)
}
