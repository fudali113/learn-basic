package adt

import "log"

// MedianHeap 实时求解中位数堆
type MedianHeap struct {
	min *Heap
	max *Heap
}

// GetMedianHeap 获取一个中位数堆
func GetMedianHeap() MedianHeap {
	return MedianHeap{min: GetHeap(false), max: GetHeap(true)}
}

// Len 获取实际长度
func (mh MedianHeap) Len() int {
	return mh.min.Len() + mh.max.Len()
}

// Insert 添加一个元素
func (mh MedianHeap) Insert(value int) {
	if mh.min.Len() == 0 {
		mh.min.Insert(value)
		return
	}
	if mh.max.Len() == 0 {
		minValue, _ := mh.min.Get()
		if value > minValue {
			mh.min.Delete()
			mh.min.Insert(value)
			mh.max.Insert(minValue)
		} else {
			mh.max.Insert(value)
		}
		return
	}

	if mh.insertMin(value) {
		mh.min.Insert(value)
	} else {
		mh.max.Insert(value)
	}
	mh.reset()
}

// GetMedian 实时获取中位数
func (mh MedianHeap) GetMedian() (median float32) {
	log.Println("GetMedian log:", mh.min, "---", mh.max)
	minL, maxL := mh.min.Len(), mh.max.Len()
	if mh.Len() == 0 {
		return 0
	}
	gap := minL - maxL
	switch gap {
	case 1:
		minV, _ := mh.min.Get()
		median = float32(minV)
	case 0:
		minV, _ := mh.min.Get()
		maxV, _ := mh.max.Get()
		median = (float32(minV) + float32(maxV)) / 2
	case -1:
		maxV, _ := mh.max.Get()
		median = float32(maxV)
	default:
		mh.reset()
		mh.GetMedian()
	}
	return
}

func (mh MedianHeap) insertMin(value int) bool {
	minV, _ := mh.min.Get()
	maxV, _ := mh.max.Get()
	minL, maxL := mh.min.Len(), mh.max.Len()
	if value > minV {
		return true
	} else if value >= maxV {
		if minL > maxL {
			return false
		}
		return true
	}
	return false
}

func (mh MedianHeap) reset() {
	minL, maxL := mh.min.Len(), mh.max.Len()
	gap := minL - maxL
	if gap >= 2 {
		minV, _ := mh.min.Delete()
		mh.max.Insert(minV)
	} else if gap <= -2 {
		maxV, _ := mh.max.Delete()
		mh.min.Insert(maxV)
	}
}
