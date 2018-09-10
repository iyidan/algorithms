package zero_sub_array

import (
	"fmt"
)

// FindAllZeroSubArray 查找全部的和为0的子数组
// 返回的是所有子数组，区间是左闭右开 [)
// 思路： 寻找和为0的子数组规律：假设子数组[i,j]和为0
// 那么 sum(a[0]...a[i-1]) = sum(a[0]...a[j])
// 如：1 2 3 4 5 0 -5 7 8 9 其中 5 0 -5 是一个和为0的子数组
// 找出全部的：从头开始累加，将sum相等的数组下标记录下来
func FindAllZeroSubArray(a []int) [][2]int {
	if len(a) == 0 {
		return nil
	}
	if len(a) == 1 {
		if a[0] == 0 {
			return [][2]int{[2]int{0, 1}}
		}
		return nil
	}

	var ret [][2]int
	sumMap := make(map[int][]int)

	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
		sumMap[s] = append(sumMap[s], i)
		fmt.Println(i, a[i], s, sumMap)
		if len(sumMap[s]) > 1 {
			for j := 0; j < len(sumMap[s])-1; j++ {
				ret = append(ret, [2]int{sumMap[s][j] + 1, i + 1})
			}
		}

	}

	return ret
}
