package kmp

// abcdex  abcabx
// -00000  -00012
func getNext(t string) (next []int) {
	next = make([]int, len(t))
	next[0] = -1
	i := 0
	j := -1
	for i < len(t)-1 {
		if j < 0 || t[j] == t[i] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return
}

// IndexKMP KMP字符串查找
func IndexKMP(s, t string) int {
	if len(t) > len(s) {
		return -1
	} else if t == s {
		return 0
	}

	next := getNext(t)
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if j < 0 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(t) {
		return i - len(t)
	}
	return -1
}
