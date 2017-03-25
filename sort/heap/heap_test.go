package heap

import (
	"testing"

	"github.com/fudali113/learn-basic/sort"
)

func Test_HeapSort(t *testing.T) {
	array := []int{10, 14, 73, 25, 39, 23, 13, 27, 33, 94, 25, 59, 94, 65, 82, 45, 80, 65}
	HeapSort(array)
	if !sort.CheckSort(array) {
		t.Error("insertionSort has bug")
	}
}
