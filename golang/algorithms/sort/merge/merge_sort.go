package merge

func mergeSort(items []int) []int {
	var num = len(items)
	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

/*
func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	return merge(mergeSort(a[:mid]), mergeSort(a[mid:]))
}

func merge(left, right []int) []int {
	// 为了不破坏原始的切片
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		// 左侧的已经全部对比完毕并添加到了slice中
		// 右侧剩余的值添加到slice中
		// 右侧剩余值都比左侧大且右侧已经是有序的
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			// 左侧的比右侧的小,将左侧的值加入到slice中
			slice[k] = left[i]
			i++
		} else {
			// 右侧的值比左侧的小,将右侧的值加入到slice中
			slice[k] = right[j]
			j++
		}
	}

	return slice
}



// 稳定 time:O(nlogn) mem:O(n)
*/
