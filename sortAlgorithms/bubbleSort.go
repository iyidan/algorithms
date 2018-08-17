package sortAlgorithms

// BubbleSort 冒泡排序，稳定排序
func BubbleSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		isChange := false // 标志本轮有没有交换，没有交换就退出本轮
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}
