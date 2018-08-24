package trie

func findMaximumXOR(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	root := NewBitTrie()
	for _, num := range nums {
		root.Add(num)
	}
	max := 0
	for _, num := range nums {
		max = swapMax(max, root.XOR(num))
	}
	return max
}

type BitTrie struct {
	childrens map[uint8]*BitTrie
}

func NewBitTrie() *BitTrie {
	return &BitTrie{childrens: make(map[uint8]*BitTrie)}
}

func (t *BitTrie) Add(num int) {
	for i := 31; i >= 0; i-- {
		bit := uint8(num>>uint(i)) & 1
		child, ok := t.childrens[bit]
		if !ok {
			child = NewBitTrie()
			t.childrens[bit] = child
		}
		t = child
	}
}

func (t *BitTrie) XOR(num int) int {
	xor := 0
	for i := 31; i >= 0; i-- {
		bit := uint8(num>>uint(i)) & 1
		xor = xor << 1
		child, ok := t.childrens[1&^bit]
		if ok {
			xor++
		} else {
			child = t.childrens[bit]
		}
		t = child
	}
	return xor
}

func swapMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
