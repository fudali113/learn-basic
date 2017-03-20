package adt

import (
	"log"
	"testing"
)

func Test_getInsertIndex(t *testing.T) {
	test := []int{0, 1}
	testInsertIndex := getInsertIndex(test, 1, 1)
	if testInsertIndex != 1 {
		t.Error("getInsertIndex has bug")
	}
	test = append(test, 2)
	testInsertIndex = getInsertIndex(test, 2, 2)
	if testInsertIndex != 1 {
		t.Error("getInsertIndex has bug", testInsertIndex)
	}
}

func Test_insert(t *testing.T) {
	log.Println("--------------------")
	test := []int{0}
	test = insert(test, 1)
	log.Println(test)
	test = insert(test, 5)
	log.Println(test)
	test = insert(test, 3)
	log.Println(test)
	test = insert(test, 4)
	log.Println(test)
}
