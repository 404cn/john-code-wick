package inserction

func inserctionSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}
