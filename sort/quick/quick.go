package quick

import "log"

// QuickSort 快速排序
// 是否稳定?
// 不稳定 可以试着分析下[5,1,2,2,3,4,5]
func QuickSort(array []int) {
	quickSort(array, 0, len(array))
}

// quickSort 快速排序私有递归方法
func quickSort(array []int, left, right int) {
	if right-left < 2 {
		return
	}
	medianIndex := median3(array, left, right)
	nowMedianIndex := swap(array, left, right, medianIndex)
	quickSort(array, left, nowMedianIndex)
	quickSort(array, nowMedianIndex+1, right)
}

// swap 按最开始的中位置下标将数组分为比中位置大和比中位值小的两个数组
// left right 遵循右开左闭
// return 划分后的中位值下标
func swap(array []int, left, right, medianIndex int) (nowMedianIndex int) {
	median := array[medianIndex]
	array[medianIndex], array[right-1] = array[right-1], median
	leftPoint := left
	rightPoint := right - 2
	leftStop := false
	rightStop := false
	for {
		// 如果左滑动指针停止，不做任何操作，等待右滑动指针也停止交换
		if !leftStop {
			if array[leftPoint] <= median {
				leftPoint++
			} else {
				leftStop = true
			}
		}

		if !rightStop {
			if array[rightPoint] >= median {
				rightPoint--
			} else {
				rightStop = true
			}
		}

		//当左滑动指针与右滑动指针相交后，左滑动指针位置一定处于大于等于中位值的位置
		if leftPoint > rightPoint {
			array[leftPoint], array[right-1] = array[right-1], array[leftPoint]
			nowMedianIndex = leftPoint
			return
		}

		// 如果都停止，交换位置并重启滑动
		if leftStop && rightStop {
			array[leftPoint], array[rightPoint] = array[rightPoint], array[leftPoint]
			leftStop, rightStop = false, false
		}

	}
}

// median3 三数中值分割法
// 比较数组开头，结尾和中间位置的值，放回处于中间的值得数组下标
// return 中位值数组下标
func median3(array []int, left, right int) int {
	index := (right + left - 1) / 2
	start := array[left]
	median := array[index]
	end := array[right-1]
	log.Println(start, median, end, array[left:right])
	if median > start {
		if median <= end {
			return index
		} else if end >= start {
			return right - 1
		} else {
			return left
		}
	} else {
		if median >= end {
			return index
		} else if start >= end {
			return right - 1
		} else {
			return left
		}
	}
}
