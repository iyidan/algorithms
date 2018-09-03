package sort_algorithms

// 调整索引i的堆符合大顶堆性质
func shiftDown(data []int, start, end int) {
	swap := start
	for start < end {
		lChild := 2*start + 1
		if lChild <= end {
			if data[start] < data[lChild] {
				swap = lChild
			}
			if lChild+1 <= end && data[swap] < data[lChild+1] {
				swap = lChild + 1
			}
		}
		if swap == start {
			return
		}
		data[swap], data[start] = data[start], data[swap]
		start = swap
	}

}

func heapAdjust(data []int, end int) {
	start := (end - 1) / 2
	for start >= 0 {
		shiftDown(data, start, end)
		start--
	}
}

// HeapSort 堆排序 O(n*logn) 不稳定排序
func HeapSort(data []int) {
	end := len(data) - 1
	heapAdjust(data, end)
	for end > 0 {
		data[0], data[end] = data[end], data[0]
		end--
		shiftDown(data, 0, end)
	}
}
