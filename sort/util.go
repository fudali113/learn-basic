package sort

import "log"

// CheckSort check sort is right
func CheckSort(array []int) (res bool) {
	log.Print("utils print", array)
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			res = false
			return
		}
	}
	res = true
	return
}
