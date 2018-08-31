package dynamic_programming

import (
	"testing"
)

func TestCutRod(t *testing.T) {
	p := []int{
		0,
		1, 5, 8, 9, 10, 17, 17, 20, 24, 30,
	}

	cases := []int{
		0,
		1, 5, 8, 10, 13, 17, 18, 22, 25, 30,
	}

	for n := 0; n < len(cases); n++ {
		if q := memoizedCutRod(p, n); q != cases[n] {
			t.Fatal("memoizedCutRod:", n, q)
		}
		if q := bottomUpCutRod(p, n); q != cases[n] {
			t.Fatal("bottomUpCutRod:", n, q)
		}
	}
}
