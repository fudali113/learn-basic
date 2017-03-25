package heap

import "github.com/fudali113/learn-basic/adt"

// HeapSort 堆排序
func HeapSort(array []int) {
	arrayLen := len(array)
	minHeap := adt.GetHeap(false)
	minHeap.Insert(array...)
	for i := 0; i < arrayLen; i++ {
		array[i], _ = minHeap.Delete()
	}
}
