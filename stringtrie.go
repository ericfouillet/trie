package trie

/*
 * ASCII tries are a special case because the maximum size of the
 * children is already known.
 * Mulitple implementations are provided, depending on the need:
 * - ASCIITrie: simple naive approach. Good for small alphabets.
 * - ASCIIReduxTrie: uses alphabet reduction. Good for larger alphabets.
 */

import "fmt"

// NewASCIITrie creates a new trie holding strings.
func NewASCIITrie() *ASCIITrie {
	return &ASCIITrie{root: newASCIITrieNode()}
}

func newASCIITrieNode() *asciiTrieNode {
	return &asciiTrieNode{} // use zero values for both fields
}

// AddWord adds a new word to an existing string trie.
func (t *ASCIITrie) AddWord(s string) error {
	return t.root.addWordToNode(s)
}

// Contains returns true if a string trie contains the given word.
func (t *ASCIITrie) Contains(s string) bool {
	return t.root.hasPrefix(s)
}

func (t *asciiTrieNode) extractNextData(s string) (byte, error) {
	c := s[0]
	if c > 255 {
		return byte(0), fmt.Errorf("Unexpected character in %v", s)
	}
	return c, nil
}

func (t *asciiTrieNode) addWordToNode(s string) error {
	if len(s) == 0 {
		t.isFullWord = true
		return nil
	}
	c, err := t.extractNextData(s)
	if err != nil {
		return err
	}
	if t.children[c] == nil {
		t.children[c] = newASCIITrieNode()
	}
	t.children[c].data = byte(c)
	return t.children[c].addWordToNode(s[1:])
}

func (t *asciiTrieNode) hasPrefix(s string) bool {
	if s == "" {
		return false
	}
	c, err := t.extractNextData(s)
	if err != nil {
		return false
	}
	if t.children[c] == nil {
		return false
	}
	if len(s) == 1 {
		return true
	}
	return t.children[c].hasPrefix(s[1:])
}

// NewASCIIReduxTrie creates a new trie holding strings.
func NewASCIIReduxTrie() *ASCIIReduxTrie {
	return &ASCIIReduxTrie{root: newASCIIReduxTrieNode(false)}
}

func newASCIIReduxTrieNode(suffixByteMask bool) *asciiReduxTrieNode {
	return &asciiReduxTrieNode{suffixByteMask: suffixByteMask} // use zero values for other fields
}

// AddWord adds a new word to an existing string trie.
func (t *ASCIIReduxTrie) AddWord(s string) error {
	return t.root.addWordToNode(s)
}

// Contains returns true if a string trie contains the given word.
func (t *ASCIIReduxTrie) Contains(s string) bool {
	return t.root.hasPrefix(s)
}

func (t *asciiReduxTrieNode) extractNextData(s string) (int, error) {
	c := s[0]
	if c > 255 {
		return -1, fmt.Errorf("Unexpected character in %v", s)
	}
	data := int(c)
	if t.suffixByteMask {
		data = data & 0x0f
	} else {
		data = data >> 4 // shift to obtain a 4 bit integer
	}
	return data, nil
}

func (t *asciiReduxTrieNode) addWordToNode(s string) error {
	if len(s) == 0 {
		t.isFullWord = true
		return nil
	}
	data, err := t.extractNextData(s)
	if err != nil {
		return err
	}
	if t.children[data] == nil {
		t.children[data] = newASCIIReduxTrieNode(!t.suffixByteMask)
	}
	t.children[data].data = byte(data)
	if t.suffixByteMask {
		return t.children[data].addWordToNode(s[1:])
	} else {
		return t.children[data].addWordToNode(s)
	}
}

func (t *asciiReduxTrieNode) hasPrefix(s string) bool {
	if s == "" {
		return false
	}
	c, err := t.extractNextData(s)
	if err != nil {
		return false
	}
	if t.children[c] == nil {
		return false
	}
	if len(s) == 1 {
		return true
	}
	if t.suffixByteMask {
		return t.children[c].hasPrefix(s[1:])
	} else {
		return t.children[c].hasPrefix(s)
	}
}
