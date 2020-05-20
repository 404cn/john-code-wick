package algorithms

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{10, 340, 34, 456, 123, 45, 57, 23}
	fmt.Printf("before sort: %v\n", arr)
	//arr = quickSort(arr)
	//arr = bubbleSort(arr)
	//arr = selectionSort(arr)
	//insertionSort(arr)
	//shellSort(arr)
	arr = mergeSort(arr)
	fmt.Printf("after sort: %v\n", arr)

}

func binarySearch(needle int, haystack []int) bool {
	left, rihgt := 0, len(haystack)-1

	for left <= right {
		mid = left + (right-left)/2
		if haystack[mid] == needle {
			return true
		} else if needle < haystack[mid] {
			right = mid - 1
		} else if needle > haystack[mid] {
			left = mid + 1
		}

	}

	return false
}

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

func shellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h; j -= h {
				if a[j] < a[j-h] {
					a[j], a[j-h] = a[j-h], a[j]
				}
			}
		}
		h = h / 3
	}
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		minIdx := i
		for j := i; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		a[i], a[minIdx] = a[minIdx], a[i]
	}

	return a
}

func bubbleSort(a []int) []int {
	var flag bool

	for i := len(a) - 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = false
			}
		}

		if !flag {
			break
		}
	}

	return a
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)

	a[right], a[pivot] = a[pivot], a[right]

	for i, _ := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}
