^package bubble

// 冒泡排序,O(n方),稳定
func bubble_sort(a []int) {
	// 标志在一次循环中是否发生了交换
	var flag bool

	for i := len(a) - 1; i > 0; i-- {
		// 每次循环初始化
		flag = false

		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				// 如果发生交换就标记一下
				flag = true
			}
		}

		// 如果没有发生交换则已经有序,跳出循环
		if !flag {
			break
		}
	}
}
