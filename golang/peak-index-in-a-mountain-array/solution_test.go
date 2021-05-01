package peakindexmountain

import "testing"

func TestPeakIndexInMountainArray(t *testing.T) {
	arr := []int{3, 6, 7, 12, 50, 25, 2, 1}
	want := 4
	got := peakIndexInMountainArray(arr)
	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
