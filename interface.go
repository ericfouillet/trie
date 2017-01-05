// Go trie data structure

package trie

import "container/list"

/*
 * Interfaces and types related to the trie data structure.
 */

// Trier represents a trie data structure.
type Trier interface {
	Add(string) error
	Contains(string) bool
}

// DataGetter is a function that extracts a trie data from a string.
// It returns the data, the number of bytes read, and an error (if any)
type DataGetter func(string) (interface{}, int, error)

// Trie is a generic implementation of a trie.
type Trie struct {
	data       interface{} // nil for the root trieNode
	isRoot     bool        // Differentiate the root node
	isFullWord bool        // words might be sub-strings of others
	children   map[interface{}]*Trie
	// This function in argument extracts the next token to process from an input data.
	// It returns the data, the number of bytes processed, and optionally
	// an error if the input string could not be processed.
	get DataGetter
}

// LinkedTrie is a generic implementation of a trie,
// using a linkedlist to store the children nodes.
type LinkedTrie struct {
	data       interface{} // nil for the root trieNode
	isFullWord bool        // words might be sub-strings of others
	isRoot     bool        // Differentiate the root node
	children   *list.List
	// The function in argument extracts the next NodeData to process from an input data.
	// It returns the data, the number of bytes processed, and optionally
	// an error if the input string could not be processed.
	get DataGetter
}

// ASCIITrie implements a trie that contains only ASCII characters.
type ASCIITrie struct {
	data       byte
	isFullWord bool            // words might be sub-strings of others
	isRoot     bool            // Differentiate the root node
	children   [256]*ASCIITrie // Size will be fixed (256 or 16, depending on the implementation)
}

// ASCIIReduxTrie implements a trie that contains only ASCII characters, and is using
// an alphabet reduction technique where strings of size n are interpreted as longer (2n)
// size strings using a reduced alphabet (16 characters).
type ASCIIReduxTrie struct {
	data           byte                // in reality only 4 bits are used
	suffixByteMask bool                // when false, the prefix byte mask 0xff00 is used to read from strings. When true, the suffix byte mask 0x00ff is used.
	isFullWord     bool                // words might be sub-strings of others
	isRoot         bool                // Differentiate the root node
	children       [16]*ASCIIReduxTrie // The alphabet contains only 16 characters
}
