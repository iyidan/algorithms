package kmp

// abcdex  abcabx
// -00000  -00012
// ---------------
// 注意next(i)与next(i-1)的关系：
// abcabcccccccccabcab  a  x xxxx
//      j              i-1 i
// 设 j = next(i-1)+1 (此时  abcabc 与 abcaba)
// 1. 如果 t[i-1]与t[j]相等，那么 next(i) = next(i-1)+1 = j
// 2. 如果 t[i-1]与t[j]不等，那么需要进行回溯，回溯多少呢？
//    注意观察当前的前缀串：abcabc abcaba
//    如果j回溯到next(j)+1 也就是两个这两个前缀串的最长前缀的位置，然后再继续对比t[i-1]与t[j]是否相等
func getNext(t string) (next []int) {
	next = make([]int, len(t))
	next[0] = -1
	i := 1
	j := next[0] + 1
	for i < len(t) {
		if j < 0 || t[i-1] == t[j] {
			next[i] = j
			i++
			j++
		} else {
			j = next[j] + 1
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
