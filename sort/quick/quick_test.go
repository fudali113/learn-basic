package quick

import (
	"log"
	"testing"

	"github.com/fudali113/learn-basic/sort"
)

func Test_swap(t *testing.T) {
	array := []int{10, 14, 73, 25, 39, 23, 13, 27, 33, 94, 25, 59, 94, 65, 82, 45, 80, 65}
	medianIndex := swap(array, 0, len(array), 8)
	log.Println("Test_swap", medianIndex, array[medianIndex], array)

}

func Test_median3(t *testing.T) {
	array := []int{10, 14, 73, 25, 39, 23, 13, 27, 33, 94, 25, 59, 94, 65, 82, 45, 80, 65}
	medianIndex := median3(array, 0, len(array))
	log.Println("Test_median3", medianIndex, array[medianIndex])
	if medianIndex != (len(array)-1)/2 {
		t.Error("median3 has bug")
	}
}

func Test_QuickSort(t *testing.T) {
	array := []int{10, 14, 73, 25, 39, 23, 13, 27, 33, 94, 25, 59, 94, 65, 82, 45, 80, 65}
	QuickSort(array)
	if !sort.CheckSort(array) {
		t.Error("insertionSort has bug")
	}
}
