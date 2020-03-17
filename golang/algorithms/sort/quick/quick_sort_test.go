package quick

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	got := []int{30, 40, 60, 10, 20, 50}
	want := []int{10, 20, 30, 40, 50, 60}

	QuickSort(got, 0, len(got)-1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Quick sort failed\ngot : %v\nwant : %v", got, want)
	}
}
