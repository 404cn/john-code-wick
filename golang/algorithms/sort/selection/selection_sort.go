package selection

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		// 最小元素的索引
		minIdx := i
		for j := i; i < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}
}
