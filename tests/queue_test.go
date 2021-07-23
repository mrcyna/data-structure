package tests

import (
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestQueue(t *testing.T) {
	q := structures.NewQueue()

	if len(q.Items) != 0 {
		t.Error("queue should be zero size at the beginning")
	}
}

func TestEnqueueOneItemToQueue(t *testing.T) {
	q := structures.NewQueue()
	q.Enqueue(100)

	if len(q.Items) != 1 {
		t.Error("queue should have length of 1")
	}

	if q.Items[0] != 100 {
		t.Error("queue first item should be 100")
	}
}

func TestEnqueueManyItemToQueue(t *testing.T) {
	q := structures.NewQueue()
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	if len(q.Items) != 3 {
		t.Error("queue should have length of 3")
	}

	if q.Items[0] != 100 {
		t.Error("queue first item should be 100")
	}

	if q.Items[1] != 200 {
		t.Error("queue first item should be 200")
	}

	if q.Items[2] != 300 {
		t.Error("queue first item should be 300")
	}
}

func TestDequeueOneItemFromQueue(t *testing.T) {
	q := structures.NewQueue()
	q.Enqueue(100)
	res, err := q.Dequeue()

	if len(q.Items) != 0 {
		t.Error("queue should have length of 0")
	}

	if res != 100 {
		t.Error("dequeue result should be 100")
	}

	if err != nil {
		t.Error("error should be nil")
	}
}

func TestDequeueManyItemFromQueue(t *testing.T) {
	q := structures.NewQueue()
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	res1, err1 := q.Dequeue()
	res2, err2 := q.Dequeue()
	res3, err3 := q.Dequeue()

	if len(q.Items) != 0 {
		t.Error("queue should have length of 0")
	}

	if res1 != 100 {
		t.Error("dequeue result should be 100")
	}

	if err1 != nil {
		t.Error("error should be nil")
	}

	if res2 != 200 {
		t.Error("dequeue result should be 200")
	}

	if err2 != nil {
		t.Error("error should be nil")
	}

	if res3 != 300 {
		t.Error("dequeue result should be 300")
	}

	if err3 != nil {
		t.Error("error should be nil")
	}
}

func TestDequeueEmptyQueueShouldThrowError(t *testing.T) {
	q := structures.NewQueue()
	res, err := q.Dequeue()

	if res != nil {
		t.Error("result should be nil")
	}

	if err == nil {
		t.Error("error must not be nil")
	}
}
