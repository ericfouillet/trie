package trie

/*
 * ASCII tries are a special case because the maximum size of the
 * children is already known.
 * Mulitple implementations are provided, depending on the need:
 * - ASCIITrie: simple naive approach. Use for small alphabets.
 * - ASCIIReduxTrie: uses alphabet reduction. Use for larger alphabets.
 */

import "fmt"

// invalid is used as a return value when a string could not be parsed
const invalid = -1

// NewASCIITrie creates a new trie holding strings.
func NewASCIITrie() *ASCIITrie {
	return &ASCIITrie{isRoot: true}
}

func newASCIITrieNode() *ASCIITrie {
	return &ASCIITrie{isRoot: false}
}

// Add adds a new word to an existing string trie.
func (t *ASCIITrie) Add(s string) error {
	if len(s) == 0 {
		t.isFullWord = true
		return nil
	}
	c, err := next(s)
	if err != nil {
		return err
	}
	if t.children[c] == nil {
		t.children[c] = newASCIITrieNode()
	}
	t.children[c].data = byte(c)
	return t.children[c].Add(s[1:])
}

// Contains returns true if a string trie contains the given word.
func (t *ASCIITrie) Contains(s string) bool {
	if s == "" {
		return false
	}
	c, err := next(s)
	if err != nil {
		return false
	}
	if t.children[c] == nil {
		return false
	}
	if len(s) == 1 {
		return true
	}
	return t.children[c].Contains(s[1:])
}

func next(s string) (byte, error) {
	c := s[0]
	if c > 255 {
		return byte(0), fmt.Errorf("Unexpected character in %v", s)
	}
	return c, nil
}

// NewASCIIReduxTrie creates a new trie holding strings.
func NewASCIIReduxTrie() *ASCIIReduxTrie {
	return &ASCIIReduxTrie{isRoot: true, suffixByteMask: true}
}

func newASCIIReduxTrieNode(byteMask bool) *ASCIIReduxTrie {
	return &ASCIIReduxTrie{isRoot: false, suffixByteMask: byteMask}
}

// Add adds a new word to an existing string trie.
func (t *ASCIIReduxTrie) Add(s string) error {
	if len(s) == 0 {
		t.isFullWord = true
		return nil
	}
	data, err := next4(s, t.suffixByteMask)
	if err != nil {
		return err
	}
	if t.children[data] == nil {
		t.children[data] = newASCIIReduxTrieNode(!t.suffixByteMask)
	}
	t.children[data].data = byte(data)
	if t.suffixByteMask {
		return t.children[data].Add(s[1:])
	} else {
		return t.children[data].Add(s)
	}
}

// Contains returns true if a string trie contains the given word.
func (t *ASCIIReduxTrie) Contains(s string) bool {
	if s == "" {
		return false
	}
	c, err := next4(s, t.suffixByteMask)
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
		return t.children[c].Contains(s[1:])
	} else {
		return t.children[c].Contains(s)
	}
}

// Returns the next 4 bits word in the given string
func next4(s string, suffixByteMask bool) (int, error) {
	c := s[0]
	if c > 255 {
		return invalid, fmt.Errorf("Unexpected character in %v", s)
	}
	data := int(c)
	if suffixByteMask {
		data = data & 0x0f
	} else {
		data = data >> 4 // shift to obtain a 4 bit integer
	}
	return data, nil
}
