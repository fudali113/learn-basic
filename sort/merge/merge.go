package merge

// MergeSort 归并排序
func MergeSort(array []int) []int {
	middle := len(array) / 2
	if middle < 1 {
		return array
	}
	return merge(MergeSort(array[:middle]), MergeSort(array[middle:]))
}

// merge 合并两个排好序的数组唯一组合两个元素并排好序的数组
func merge(a1 []int, a2 []int) []int {
	a1Index := 0
	a2Index := 0
	res := make([]int, 0, len(a1)+len(a2))
	for {
		if a1[a1Index] <= a2[a2Index] {
			res = append(res, a1[a1Index])
			if a1Index == len(a1)-1 {
				res = append(res, a2[a2Index:]...)
				break
			}
			a1Index++
		} else {
			res = append(res, a2[a2Index])
			if a2Index == len(a1)-1 {
				res = append(res, a1[a1Index:]...)
				break
			}
			a2Index++
		}
	}
	return res
}
