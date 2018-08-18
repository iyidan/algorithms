package kmp

import (
	"testing"
)

func TestIndexKMP(t *testing.T) {
	type TestCase struct {
		s, t string
		ret  int
	}

	cases := []TestCase{
		TestCase{"", "", 0},
		TestCase{"a", "a", 0},
		TestCase{"a", "aa", -1},
		TestCase{"abcde", "abcdf", -1},
		TestCase{"abcde", "ab", 0},
		TestCase{"abcde", "cd", 2},
		TestCase{"abcababcababcabxabcabx", "abcabx", 10},
	}

	for _, c := range cases {
		if ret := IndexKMP(c.s, c.t); ret != c.ret {
			t.Fatal(c.s, c.t, c.ret, ret)
		}
	}
}
