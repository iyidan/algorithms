package dynamic_programming

import (
	"fmt"
)

// LCS 求两个串的最长公共子序列
// ABCBDAB
// BDCABA
// lcs BCBA
// -----------
// 问题分析：最优化问题，首先想到动态规划的方式
// 分析最优解的结构(通过假设找到了lcs，并分析与两个串之间的关系)：
// 令X<x1,x2,...,xm>，Y<y1,y2,...,yn> 为两个串，Z<z1,z2,...,zk>是他们的一个LCS
// 1. 从两个串的最后一位看起，如果 xm = yn，那么Z中的最后一个字符zk必然与xm,yn相等 zk=xm=yn
// 2. 如果 xm != yn，那我们找到的Z必然是下面几种种情况：
//    2-1：如果zk != yn，那么Z必然是<x1,x2,...,xm>与<y1,y2,...yn-1>的lcs
//    2-2：如果zk != xm，那么Z必然是<x1,x2,...,xm-1>与<y1,y2,...yn>的lcs
// 这样我们就把问题规模缩小了，并且 子问题的最优解也与父问题相关
func LCS(X, Y string) string {
	if len(X) == 0 || len(Y) == 0 {
		return ""
	}
	lcs := make([][]string, len(X))
	for m := 0; m < len(X); m++ {
		lcs[m] = make([]string, len(Y))
	}
	// 自底向上法
	for m := 0; m < len(X); m++ {
		for n := 0; n < len(Y); n++ {
			if m == 0 || n == 0 {
				if X[m] == Y[n] {
					lcs[m][n] = string(X[m])
				}
				continue
			}
			if X[m] == Y[n] { // xm = yn
				lcs[m][n] = lcs[m-1][n-1] + string(X[m])
			} else if len(lcs[m-1][n]) >= len(lcs[m][n-1]) { // xm != yn，哪个大就取哪个
				lcs[m][n] = lcs[m-1][n]
			} else {
				lcs[m][n] = lcs[m][n-1]
			}
		}
	}
	fmt.Println(X, Y, lcs)
	return lcs[len(X)-1][len(Y)-1]
}

// LCCS 最长公共子串
// 是lcs的一个变种，如果xm!=yn，那么 lccs[m][n] = lccs[m-1][n-1]
func LCCS(X, Y string) string {
	if len(X) == 0 || len(Y) == 0 {
		return ""
	}
	lccs := make([][]string, len(X))
	for m := 0; m < len(X); m++ {
		lccs[m] = make([]string, len(Y))
	}
	// 自底向上法
	maxLen := 0
	for m := 0; m < len(X); m++ {
		for n := 0; n < len(Y); n++ {
			if m == 0 || n == 0 {
				if X[m] == Y[n] {
					lccs[m][n] = string(X[m])
					if len(lccs[m][n]) > maxLen {
						maxLen = len(lccs[m][n])
					}
				}
				continue
			}
			if X[m] == Y[n] { // xm = yn
				lccs[m][n] = lccs[m-1][n-1] + string(X[m])
				if len(lccs[m][n]) > maxLen {
					maxLen = len(lccs[m][n])
				}
			} else {
				lccs[m][n] = ""
			}
		}
	}
	fmt.Printf("%s, %s, %#v, %d\n", X, Y, lccs, maxLen)
	for _, v := range lccs {
		for _, w := range v {
			if len(w) == maxLen {
				return w
			}
		}
	}
	return ""
}
