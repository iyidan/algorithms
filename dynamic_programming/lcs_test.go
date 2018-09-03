package dynamic_programming

import (
	"testing"
)

func TestLCS(t *testing.T) {
	if lcs := LCS("", ""); lcs != "" {
		t.Fatal(lcs)
	}
	if lcs := LCS("ABCDE", "FGH"); lcs != "" {
		t.Fatal(lcs)
	}
	if lcs := LCS("BDCABA", "ABCBDAB"); lcs != "BDAB" {
		t.Fatal(lcs)
	}
	if lcs := LCS("ABCBDABDDD", "BDCABAEEE"); lcs != "BCBA" {
		t.Fatal(lcs)
	}
}

func TestLCCS(t *testing.T) {
	if lccs := LCCS("", ""); lccs != "" {
		t.Fatal(lccs)
	}
	if lccs := LCCS("ABCDE", "FGH"); lccs != "" {
		t.Fatal(lccs)
	}
	if lccs := LCCS("acbac", "acaccbabb"); lccs != "cba" {
		t.Fatal(lccs)
	}
}
