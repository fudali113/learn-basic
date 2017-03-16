package insertion

import (
	"sorts"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	array := []int{10, 14, 73, 25, 23, 13, 27, 94, 33, 39, 25, 59, 94, 65, 82, 45, 100}
	sort(array)
	if !sorts.CheckSort(array) {
		t.Error("insertionSort has bug")
	}
}
