package adt

import (
	"log"
	"testing"
)

func Test_getInsertIndex(t *testing.T) {
	test := []int{0, 1}
	testInsertIndex := getInsertIndex(test, 1, 1, true)
	if testInsertIndex != 1 {
		t.Error("getInsertIndex has bug")
	}
	test = append(test, 2)
	testInsertIndex = getInsertIndex(test, 2, 2, true)
	if testInsertIndex != 1 {
		t.Error("getInsertIndex has bug", testInsertIndex)
	}
}

func Test_insert(t *testing.T) {
	log.Println("--------------------")
	test := []int{0}
	test = insert(test, 1, true)
	log.Println(test)
	test = insert(test, 5, true)
	log.Println(test)
	test = insert(test, 3, true)
	log.Println(test)
	test = insert(test, 4, true)
	log.Println(test)
}

func Test_reset(t *testing.T) {
	test := []int{0, 5, 4, 3, 1}
	test, _, _ = reset(test, true)
	log.Println(test)
	test, _, _ = reset(test, true)
	log.Println(test)
	test, _, _ = reset(test, true)
	log.Println(test)
}

func Test_HeadInsert(t *testing.T) {
	h := GetHead(true)
	h.Insert(1)
	h.Insert(2)
	h.Insert(5)
	h.Insert(4)
	log.Println("max heap", h)
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())

	h = GetHead(false)
	h.Insert(1)
	h.Insert(2)
	h.Insert(5)
	h.Insert(4)
	log.Println("min heap", h)
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
	log.Println(h.Delete())
}
