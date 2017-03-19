package merge

import (
	"log"
	"testing"
)

// func Test_MergeSort(t *testing.T) {
// 	array := []int{10, 14, 73, 25, 39, 23, 13, 27, 33, 94, 25, 59, 94, 65, 82, 45, 80, 65}
// 	array = MergeSort(array)
// 	if !sort.CheckSort(array) {
// 		t.Error("mergeSort has bug")
// 	}
// }

func Test_merge(t *testing.T) {
	a1 := []int{4, 25, 33}
	a2 := []int{17, 18, 31}
	res := merge(a1, a2)
	log.Println(res)
}
