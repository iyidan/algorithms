package sortAlgorithms

// InsertionSort 插入排序 稳定排序
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i - 1; j >= 0 && data[j+1] < data[j]; j-- {
			data[j+1], data[j] = data[j], data[j+1]
		}
	}
}
