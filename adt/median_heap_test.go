package adt

import (
	"log"
	"testing"
)

func Test_medianHeap(t *testing.T) {
	medianHeap := GetMedianHeap()
	log.Println("median heap test :", medianHeap.GetMedian())

	medianHeap.Insert(1)
	log.Println("median heap test :", medianHeap.GetMedian())

	medianHeap.Insert(2)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(3)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(4)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(5)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(6)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(2)
	log.Println("median heap test :", medianHeap.GetMedian())
	medianHeap.Insert(5)
	log.Println("median heap test :", medianHeap.GetMedian())

}
