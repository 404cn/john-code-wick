package bubble

import (
	"reflect"
	"testing"
)

func TestBubble_Sort(t *testing.T) {
	got := []int{20, 40, 30, 10, 60, 50}
	want := []int{10, 20, 30, 40, 50, 60}

	bubble_sort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("bubble sort failed\ngot : %v\nwant : %v\n", got, want)
	}
}
