package tests

import (
	"reflect"
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestMaxHeap(t *testing.T) {
	mh := structures.NewMaxHeap()

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		mh.Insert(v)
	}

	if !reflect.DeepEqual(mh.Items(), []int{30, 17, 20, 13, 15, 9, 11, 5, 10, 7}) {
		t.Error("the sort of heap of inserts is wrong")
	}

	for i := 0; i < 5; i++ {
		mh.Extract()
	}

	if !reflect.DeepEqual(mh.Items(), []int{11, 10, 9, 5, 7}) {
		t.Error("the sort of heap after extract is wrong")
	}
}
