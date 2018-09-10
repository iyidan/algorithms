package str_permutation

func StrPermutation(str string) []string {
	if len(str) == 0 {
		return nil
	}
	if len(str) == 1 {
		return []string{str}
	}
	s0 := str[0]
	rets := StrPermutation(str[1:])
	prevLen := len(rets)
	for _, tmp := range rets {
		for i := 0; i <= len(tmp); i++ {
			tb := []byte(tmp)
			rets = append(rets, string(append(tb[:i], append([]byte{s0}, tb[i:]...)...)))
		}
	}
	return rets[prevLen:]
}
