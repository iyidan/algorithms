package sortAlgorithms

// QuickSort 快速排序法 不稳定排序
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	pos := 0
	tail := len(data) - 1
	for pos < tail {
		if data[0] < data[pos+1] {
			data[pos+1], data[tail] = data[tail], data[pos+1]
			tail--
		} else {
			pos++
		}
	}

	data[0], data[pos] = data[pos], data[0]
	QuickSort(data[:pos])
	QuickSort(data[pos+1:])
}
