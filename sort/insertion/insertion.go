package insertion

// InsertionSort 插入排序
func InsertionSort(array []int) {
	n := 0
	// 从1开始遍历,因为0前面没有数值
	for i := 1; i < len(array); i++ {
		// 在此处放置与在 代码1 处放置效果一样
		// 因为若 jv > iv 的话 ， j 与 j-1 调换位置， 下一次循环时 array[j] 依然是当初的 iv
		iv := array[i]
		for j := i; j >= 0 && array[j-1] >= iv; j-- {
			// 代码1  iv := array[j]

			// 如果 jv > iv , 调换两个元素之间的位置
			n++
			array[j], array[j-1] = array[j-1], iv

		}
	}
	println("----------------", n)
}
