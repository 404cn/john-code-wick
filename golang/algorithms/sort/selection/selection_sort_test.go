package selection

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	got := []int{20, 50, 10, 60, 80, 30, 70, 40}
	want := []int{10, 20, 30, 40, 50, 60, 70, 80}

	selectionSort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("selection sort failed\ngot:%v\nwant:%v", got, want)
	}
}
