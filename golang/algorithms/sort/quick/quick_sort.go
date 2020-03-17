package quick

func QuickSort(a []int, l, r int) {
	if l < r {
		i, j := l, r
		// 保存基准值
		x := a[i]
		for i < j {
			for i < j && a[j] > x {
				j--
			}
			if i < j {
				a[i] = a[j]
				i++
			}
			for i < j && a[i] < x {
				i++
			}
			if i < j {
				a[j] = a[i]
				j--
			}
		}

		a[i] = x
		QuickSort(a, l, i-1)
		QuickSort(a, i+1, r)
	}
}

// 不稳定
// 最坏O(n2), 平均复杂度(n*lgn)
