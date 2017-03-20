package adt

import "fmt"

// Head 堆结构
// 数组形式为index为0的元素默认为空，总index为1元素开始
// 方便定位元素
// 即如忘堆里插入1,2,3,4,5
// 数组元素为[0,1,2,3,4,5]
type Head struct {
	array []int
	isMax bool
}

// GetHead 获取一个Head实列
func GetHead(isMax bool) Head {
	return Head{array: make([]int, 1, 8), isMax: isMax}
}

// Insert 插入一个元素
func (h Head) Insert(v int) {
	if h.array == nil {
		h.array = make([]int, 1, 8)
	}

}

// insert 根据堆性质想数组中插入一个元素
func insert(array []int, value int) (nowArray []int) {
	nowArray = append(array, 0)
	insertIndex := getInsertIndex(nowArray, len(array), value)
	nowArray[insertIndex] = value
	return nowArray
}

// getInsertIndex 递归获取插入元素的数组下标
// 使用上虑策略
func getInsertIndex(array []int, nowIndex int, value int) int {
	if nowIndex < 2 {
		return 1
	}
	faIndex := getFaIndex(nowIndex)
	if array[faIndex] >= value {
		return nowIndex
	}
	array[faIndex], array[nowIndex] = array[nowIndex], array[faIndex]
	return getInsertIndex(array, faIndex, value)
}

// getSonIndex 获取一个元素的左右子元素数组下标
func getSonIndex(index int) (left, right int) {
	return 2 * index, 2*index + 1
}

// getFaIndex 获取一个元素的父级元素数组下标
func getFaIndex(index int) int {
	return index / 2
}

// Delete 删除并获取堆顶
func (h Head) Delete() (int, error) {
	if len(h.array) < 2 {
		return 0, fmt.Errorf("改堆内以没有元素")
	}
	top := h.array[1]
	h.array = reset(h.array)
	return top, nil
}

// reset 重设置一个数组
func reset(array []int) []int {
	array[1] = 0
	value := array[len(array)-1]
	nowArray := array[:len(array)-1]
	lastInsertIndex := getLastInsertIndex(nowArray, 1, value)
	nowArray[lastInsertIndex] = value
	return nowArray
}

// getLastInsertIndex 下虑
func getLastInsertIndex(array []int, nowIndex int, value int) int {
	status := noGo
	lastIndex := len(array) - 1
	leftIndex, rightIndex := getSonIndex(nowIndex)
	if lastIndex >= rightIndex {
		left, right := array[leftIndex], array[rightIndex]
		status = getStatus(left, right, value)
	} else if lastIndex == leftIndex && array[leftIndex] > value {
		status = goLeft
	}
	switch status {
	case noGo:
		return nowIndex
	case goLeft:
		array[nowIndex], array[leftIndex] = array[leftIndex], value
		return getLastInsertIndex(array, leftIndex, value)
	case goRight:
		array[nowIndex], array[rightIndex] = array[rightIndex], value
		return getLastInsertIndex(array, rightIndex, value)
	}
	panic("--")
}

const (
	noGo = iota
	goLeft
	goRight
)

// getStatus 根据三个值获取走那边获取步鄹
func getStatus(left, right, value int) (res int) {
	res = 0
	if left > value || right > value {
		res++
	} else {
		return
	}
	if right > left {
		res++
	}
	return
}
