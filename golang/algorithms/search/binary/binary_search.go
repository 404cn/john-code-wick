package binary

func binarySearch(needle int, haystack []int) bool {
	left, right := 0, len(haystack)-1

	for left <= right {
		// 防止溢出
		mid := left + (right-left)/2
		if haystack[mid] == needle {
			return true
		} else if needle < haystack[mid] {
			// mid已经被检查过, 所以将right赋值为mid-1
			right = mid - 1
		} else if needle > haystack[mid] {
			left = mid + 1
		}
	}

	return false
}
