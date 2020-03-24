package binary

import "testing"

func TestBinaryTest(t *testing.T) {
	got := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7}
	want := -3

	result := binarySearch(want, got)

	if !result {
		t.Errorf("binary search failed\ngot:%v\nwant:%v", got, want)
	}
}
