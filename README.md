[![Build Status](https://travis-ci.org/ericfouillet/trie.svg?branch=master)](https://travis-ci.org/ericfouillet/trie)

# trie

Basic trie implementations in Go.

For now, the following is implemented:

- `StringTrie`: basic trie supporting only ASCII character alphabets.
- `StringReduxTrie`: basic trie supporting only ASCII character alphabets and using an alphabet reduction technique to reduce memory usage for large dictionaries.
- `Trie`: the base trie supporting all kind of values (`interface{}`), uses a map to store children.
- `LinkedTrie`: the base trie supporting all kind of values (`interface{}`), uses a linkedlist to store children.

# Installation

`go get github.com/ericfouillet/trie`

# Development

- Install as above or clone the repository
- Install [Glide](https://glide.sh)
- Run `glide install`
- Run `go test .`
- Run `go build .` or `go install .`

# Usage

The base `Trie` requires to provide a `DataGetter` function to extract the next token to process from a string. This allows implementing tries for your own usage.

A rune extractor is provided, to extract runes from a string. Build it as follows:

```
tr := trie.New(trie.RuneGetter)
```

or, if using a `LinkedTrie`:

```
tr := trie.NewLinked(trie.RuneGetter)
```

Alternatively, provide your own `DataGetter` with the following signature:

```
// DataGetter is a function that extracts a trie data from a string.
// It returns the data, the number of bytes read, and an error (if any)
type DataGetter func(string) (interface{}, int, error)
```
