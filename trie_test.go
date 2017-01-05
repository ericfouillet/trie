// Go trie data structure

package trie

/*
 * Tests for the Rune and ASCII trie data structure implementations.
 */

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func TestStringTrie(t *testing.T) {
	tr := NewASCIITrie()
	for i := 0; i < 10000; i++ {
		l := rand.Int() % 100
		for l == 0 {
			l = rand.Int() % 100
		}
		b := make([]byte, 0, l)
		for j := 0; j < l; j++ {
			ch := alphabet[rand.Int()%26]
			b = append(b, ch)
		}
		w := string(b)
		tr.Add(w)
		if !assert.True(t, tr.Contains(w), "The trie shoud contain %v", w) {
			return
		}
	}
}

func TestStringReduxTrie(t *testing.T) {
	tr := NewASCIIReduxTrie()
	for i := 0; i < 10000; i++ {
		l := rand.Int() % 100
		for l == 0 {
			l = rand.Int() % 100
		}
		b := make([]byte, 0, l)
		for j := 0; j < l; j++ {
			ch := alphabet[rand.Int()%26]
			b = append(b, ch)
		}
		w := string(b)
		tr.Add(w)
		if !assert.True(t, tr.Contains(w), "The trie shoud contain %v", w) {
			return
		}
	}
}

func TestRuneTrie(t *testing.T) {
	runes := [...]string{"日本語", "wordTest", "こんにちは", "読む", "パソコン"}
	tr := New(RuneGetter)
	for _, w := range runes {
		tr.Add(w)
		if !assert.True(t, tr.Contains(w), "The trie shoud contain %v", w) {
			return
		}
	}
}

func TestLinkedRuneTrie(t *testing.T) {
	runes := [...]string{"日本語", "wordTest", "こんにちは", "読む", "パソコン"}
	tr := NewLinked(RuneGetter)
	for _, w := range runes {
		tr.Add(w)
		if !assert.True(t, tr.Contains(w), "The trie shoud contain %v", w) {
			return
		}
	}
}
