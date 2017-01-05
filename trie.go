// Go trie data structure

package trie

/*
 * Implementation of the Trie interface.
 */

import (
	"container/list"
	"fmt"
	"unicode/utf8"
)

// eof is used when a method is unable to extract a rune from a string
const eof rune = -1

// New creates a new trie.
func New(g DataGetter) *Trie {
	return &Trie{isRoot: true, get: g}
}

func newNode(g DataGetter) *Trie {
	return &Trie{isRoot: false, get: g}
}

// Add adds a new word to a trie.
func (t *Trie) Add(s string) error {
	if len(s) == 0 {
		return nil
	}
	w, n, err := t.get(s)
	if err != nil {
		return fmt.Errorf("Could not decode string %v: %v", s, err)
	}
	if t.children == nil {
		t.children = make(map[interface{}]*Trie)
	}
	if _, ok := t.children[w]; !ok {
		t.children[w] = newNode(t.get)
	}
	t.children[w].data = w
	if len(s) == n {
		t.children[w].isFullWord = true
		return nil
	}
	return t.children[w].Add(s[n:])
}

// Contains returns true if a trie contains the given word.
func (t *Trie) Contains(s string) bool {
	if s == "" {
		return false
	}
	w, n, err := t.get(s)
	if err != nil {
		return false
	}
	if _, ok := t.children[w]; !ok {
		return false
	}
	if len(s) == n {
		return true
	}
	return t.children[w].Contains(s[n:])
}

// RuneGetter extracts the next RuneData to process from a string.
func RuneGetter(s string) (interface{}, int, error) {
	ru, n := utf8.DecodeRuneInString(s)
	if ru == utf8.RuneError {
		return eof, 0, fmt.Errorf("Could not decode rune in string %v", s)
	}
	return ru, n, nil
}

// NewLinked creates a new trie using a linked list to store children.
func NewLinked(g DataGetter) *LinkedTrie {
	return &LinkedTrie{isRoot: true, get: g}
}

func newLinkedNode(g DataGetter) *LinkedTrie {
	return &LinkedTrie{isRoot: false, get: g}
}

// Add adds a new word to a linked trie.
func (t *LinkedTrie) Add(s string) error {
	if len(s) == 0 {
		return nil
	}
	w, n, err := t.get(s)
	if err != nil {
		return fmt.Errorf("Could not decode string %v: %v", s, err)
	}
	if t.children == nil {
		t.children = list.New()
	}
	var found *LinkedTrie
	for e := t.children.Front(); e != nil; e = e.Next() {
		if child, ok := e.Value.(*LinkedTrie); ok && child.data == w {
			found = child
			break
		}
	}
	if found == nil {
		found = newLinkedNode(t.get)
		found.data = w
		t.children.PushBack(found)
	}
	if len(s) == n {
		found.isFullWord = true
		return nil
	}
	return found.Add(s[n:])
}

// Contains returns true if a linked trie contains the given word.
func (t *LinkedTrie) Contains(s string) bool {
	if s == "" {
		return false
	}
	w, n, err := t.get(s)
	if err != nil {
		return false
	}
	var found *LinkedTrie
	for e := t.children.Front(); e != nil; e = e.Next() {
		if child, ok := e.Value.(*LinkedTrie); ok && child.data == w {
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
	return found.Contains(s[n:])
}
