package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	root := Constructor()
	root.Insert("apple")
	if !root.Search("apple") {
		t.Fatal()
	}
	if root.Search("app") {
		t.Fatal()
	}
	if !root.StartsWith("app") {
		t.Fatal()
	}
	root.Insert("app")
	if !root.Search("app") {
		t.Fatal()
	}
}

func Test_findMaximumXOR(t *testing.T) {
	if max := findMaximumXOR([]int{3, 10, 5, 25, 2, 8}); max != 28 {
		t.Fatal(max)
	}
}
