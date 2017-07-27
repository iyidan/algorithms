package sortAlgorithms

// QuickSort 快速排序法 不稳定排序
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	i := 1
	head := 0
	tail := len(data) - 1
	for i = 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}
