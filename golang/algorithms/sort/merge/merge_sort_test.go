package merge

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	got := []int{80, 30, 60, 40, 20, 10, 50, 70}
	want := []int{10, 20, 30, 40, 50, 60, 70, 80}

	got = mergeSort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Merge sort failed\ngot:%v\nwant:%v\n", got, want)
	}
}
