package trie

type Trie struct {
	childrens map[rune]*Trie
	isWord    bool
}

/** Initialize your data structure here. */
func Constructor() *Trie {
	return &Trie{childrens: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		child, ok := this.childrens[c]
		if !ok {
			child = &Trie{childrens: make(map[rune]*Trie)}
			this.childrens[c] = child
		}
		this = child
	}
	this.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		child, ok := this.childrens[c]
		if !ok {
			return false
		}
		this = child
	}
	return this.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		child, ok := this.childrens[c]
		if !ok {
			return false
		}
		this = child
	}
	return this.isWord || len(this.childrens) > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
