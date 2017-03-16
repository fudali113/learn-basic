package quick

import (
	"testing"

	"github.com/fudali113/learn-basic/sort"
)

func Test_quickSort(t *testing.T) {
	array := []int{10, 14, 73, 25, 23, 13, 27, 94, 33, 39, 25, 59, 94, 65, 82, 45, 100}
	quickSort(array)
	if !sort.CheckSort(array) {
		t.Error("insertionSort has bug")
	}
}
