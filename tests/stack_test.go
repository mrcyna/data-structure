package tests

import (
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestStack(t *testing.T) {
	q := structures.NewStack()

	if len(q.Items) != 0 {
		t.Error("stack should be zero size at the beginning")
	}
}

func TestPushOneItemToStack(t *testing.T) {
	q := structures.NewStack()
	q.Push(100)

	if len(q.Items) != 1 {
		t.Error("stack should have length of 1")
	}

	if q.Items[0] != 100 {
		t.Error("stack first item should be 100")
	}
}

func TestPushManyItemToStack(t *testing.T) {
	q := structures.NewStack()
	q.Push(100)
	q.Push(200)
	q.Push(300)

	if len(q.Items) != 3 {
		t.Error("stack should have length of 3")
	}

	if q.Items[0] != 100 {
		t.Error("stack first item should be 100")
	}

	if q.Items[1] != 200 {
		t.Error("stack first item should be 200")
	}

	if q.Items[2] != 300 {
		t.Error("stack first item should be 300")
	}
}

func TestPopOneItemFromStack(t *testing.T) {
	q := structures.NewStack()
	q.Push(100)
	res, err := q.Pop()

	if len(q.Items) != 0 {
		t.Error("stack should have length of 0")
	}

	if res != 100 {
		t.Error("pop result should be 100")
	}

	if err != nil {
		t.Error("error should be nil")
	}
}

func TestPopManyItemFromStack(t *testing.T) {
	q := structures.NewStack()
	q.Push(100)
	q.Push(200)
	q.Push(300)

	res1, err1 := q.Pop()
	res2, err2 := q.Pop()
	res3, err3 := q.Pop()

	if len(q.Items) != 0 {
		t.Error("stack should have length of 0")
	}

	if res1 != 300 {
		t.Error("pop result should be 300")
	}

	if err1 != nil {
		t.Error("error should be nil")
	}

	if res2 != 200 {
		t.Error("pop result should be 200")
	}

	if err2 != nil {
		t.Error("error should be nil")
	}

	if res3 != 100 {
		t.Error("pop result should be 100")
	}

	if err3 != nil {
		t.Error("error should be nil")
	}
}

func TestPopEmptyStackShouldThrowError(t *testing.T) {
	q := structures.NewStack()
	res, err := q.Pop()

	if res != nil {
		t.Error("result should be nil")
	}

	if err == nil {
		t.Error("error must not be nil")
	}
}
