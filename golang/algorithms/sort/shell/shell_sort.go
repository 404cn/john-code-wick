package shell

func shellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// 将数组变为h有序
		for i := h; i < n; i++ {
			// 将a[i]插入到a[i-h], a[i-2*h], a[i-3*h]... 之中不过
			for j := i; j >= h; j -= h {
				if a[j] < a[j-h] {
					a[j], a[j-h] = a[j-h], a[j]
				}
			}
		}
		h = h / 3
	}
}
