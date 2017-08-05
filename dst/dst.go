package dst

import (
	"math/rand"
	"time"
)

const (
	// MaxInt the maximum int depend on the current system
	//MaxInt = int(^uint(0) >> 1)
	MaxInt = 65535
)

// 将一个整数变为绝对值
func intabs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 将一个整数变为绝对值
func int8abs(a int8) int8 {
	if a < 0 {
		return -a
	}
	return a
}

// 计算以x为底，y为幂次方值
func intPow(x, y int) int {
	if y == 0 {
		return 1
	}
	if y > 0 {
		ret := x
		for i := 1; i < y; i++ {
			ret *= x
		}
		return ret
	}
	// y < 0 , return zero
	return 0
}

var (
	shuffleRander = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// ShuffleSliceInt 打乱一个数组
func ShuffleSliceInt(data []int) {
	for i := len(data) - 1; i > 0; i-- {
		ridx := shuffleRander.Intn(i)
		data[i], data[ridx] = data[ridx], data[i]
	}
}

// ShuffleSliceBTData 打乱
func ShuffleSliceBTData(data []BTSortDataer) {
	for i := len(data) - 1; i > 0; i-- {
		ridx := shuffleRander.Intn(i)
		data[i], data[ridx] = data[ridx], data[i]
	}
}
