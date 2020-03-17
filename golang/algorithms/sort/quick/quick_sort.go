package quick

import "math/rand"

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

func quickSort2(a []int) []int {
	if len(a) < 2 {
		return a
	}

	// left用来标志基准值在一次排序后的位置
	// 既left左边的值都比基准值小
	left, right := 0, len(a)-1
	// 设定基准值
	pivot := rand.Int() % len(a)

	// 先将基准值放到最右侧,然后将其余值与基准值对比
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	// left 的位置左侧的值都比基准值小,将基准值与left调换位置
	// 然后继续排序两边
	a[left], a[right] = a[right], a[left]

	quickSort2(a[:left])
	// 由于基准值的位置已经确定,所以为left+1
	quickSort2(a[left+1:])
	return a
}
