// Go trie data structure

package trie

import "container/list"

/*
 * Interfaces and types related to the trie data structure.
 */

// Trier represents a trie data structure.
type Trier interface {
	AddWord(string) error
	Contains(string) bool
}

type trieNoder interface {
	addWordToNode(string, NodeDataExtractor) error
	hasPrefix(string, NodeDataExtractor) bool
}

// Trie is a generic implementation of a trie.
// It requires the node data to implement the NodeData interface.
type Trie struct {
	root *node
	// The function in argument extracts the next NodeData to process from an input data.
	// It returns the data, the number of bytes processed, and optionally
	// an error if the input string could not be processed.
	DataExtractor NodeDataExtractor
}

type node struct {
	data       interface{} // nil for the root trieNode
	isFullWord bool        // words might be sub-strings of others
	children   map[interface{}]*node
}

// LinkedTrie is a generic implementation of a trie,
// using a linkedlist to store the children nodes.
// It requires the node data to implement the NodeData interface.
type LinkedTrie struct {
	root *linkedNode
	// The function in argument extracts the next NodeData to process from an input data.
	// It returns the data, the number of bytes processed, and optionally
	// an error if the input string could not be processed.
	DataExtractor NodeDataExtractor
}

type linkedNode struct {
	data       interface{} // nil for the root trieNode
	isFullWord bool        // words might be sub-strings of others
	children   *list.List
}

// NodeDataExtractor is an interface for functions that extract trie data from a string.
type NodeDataExtractor interface {
	ExtractNextData(string) (interface{}, int, error)
}

// NodeDataExtractorFunc is a function that extracts a trie data from a string.
type NodeDataExtractorFunc func(string) (interface{}, int, error)

// ExtractNextData extracts trie data from a string.
func (e NodeDataExtractorFunc) ExtractNextData(s string) (interface{}, int, error) {
	return e(s)
}

// ASCIITrie implements a trie that contains only ASCII characters.
type ASCIITrie struct {
	root *asciiTrieNode
}

type asciiTrieNode struct {
	data       byte
	isFullWord bool // words might be sub-strings of others
	children   [256]*asciiTrieNode
}

// ASCIIReduxTrie implements a trie that contains only ASCII characters, and is using
// an alphabet reduction technique where strings of size n are interpreted as longer (2n)
// size strings using a reduced alphabet (16 characters).
type ASCIIReduxTrie struct {
	root *asciiReduxTrieNode
}

type asciiReduxTrieNode struct {
	data           byte                    // in reality only 4 bits are used
	suffixByteMask bool                    // when false, the prefix byte mask 0xff00 is used to read from strings. When true, the suffix byte mask 0x00ff is used.
	isFullWord     bool                    // words might be sub-strings of others
	children       [16]*asciiReduxTrieNode // The alphabet contains only 16 characters
}
