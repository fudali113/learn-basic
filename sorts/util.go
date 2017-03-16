package sorts

// CheckSort check sort is right
func CheckSort(array []int) (res bool) {
	for i := 1; i < len(array); i++ {
		if array[i] > array[i-1] {
			res = false
			break
		}
	}
	res = true
	return
}
