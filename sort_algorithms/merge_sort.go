package sort_algorithms

// max numbers
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// MergeSort 归并排序
func MergeSort(data []int) {
	if len(data) < 2 {
		return
	}
	mid := len(data) / 2
	MergeSort(data[:mid])
	MergeSort(data[mid:])
	mergeTwoSortedArr(data, mid)
}

func mergeTwoSortedArr(data []int, idx int) {
	a1 := make([]int, len(data[:idx])+1)
	a2 := make([]int, len(data[idx:])+1)
	copy(a1, data[:idx])
	copy(a2, data[idx:])
	a1[len(a1)-1] = MaxInt
	a2[len(a2)-1] = MaxInt
	i := 0
	j := 0
	for k := 0; k < len(data); k++ {
		if a1[i] <= a2[j] {
			data[k] = a1[i]
			i++
		} else {
			data[k] = a2[j]
			j++
		}
	}
}
