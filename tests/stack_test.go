package tests

import (
	"testing"

	"github.com/mrcyna/data-structures/structures"
)

func TestStack(t *testing.T) {
	s := structures.NewStack()

	if len(s.Items) != 0 {
		t.Error("stack should be zero size at the beginning")
	}
}

func TestPushOneItemToStack(t *testing.T) {
	s := structures.NewStack()
	s.Push(100)

	if len(s.Items) != 1 {
		t.Error("stack should have length of 1")
	}

	if s.LastIndex() != 0 {
		t.Error("last index of stack should be 0")
	}

	if s.Items[0] != 100 {
		t.Error("stack first item should be 100")
	}
}

func TestPushManyItemToStack(t *testing.T) {
	s := structures.NewStack()
	s.Push(100)
	s.Push(200)
	s.Push(300)

	if len(s.Items) != 3 {
		t.Error("stack should have length of 3")
	}

	if s.Items[0] != 100 {
		t.Error("stack first item should be 100")
	}

	if s.Items[1] != 200 {
		t.Error("stack first item should be 200")
	}

	if s.Items[2] != 300 {
		t.Error("stack first item should be 300")
	}

	if s.LastIndex() != 2 {
		t.Error("last index of stack should be 2")
	}
}

func TestPopOneItemFromStack(t *testing.T) {
	s := structures.NewStack()
	s.Push(100)
	res, err := s.Pop()

	if len(s.Items) != 0 {
		t.Error("stack should have length of 0")
	}

	if res != 100 {
		t.Error("pop result should be 100")
	}

	if err != nil {
		t.Error("error should be nil")
	}

	if s.LastIndex() != -1 {
		t.Error("last index of stack should be -1")
	}
}

func TestPopManyItemFromStack(t *testing.T) {
	s := structures.NewStack()
	s.Push(100)
	s.Push(200)
	s.Push(300)

	res1, err1 := s.Pop()
	res2, err2 := s.Pop()
	res3, err3 := s.Pop()

	if len(s.Items) != 0 {
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

	if s.LastIndex() != -1 {
		t.Error("last index of stack should be -1")
	}
}

func TestPopEmptyStackShouldThrowError(t *testing.T) {
	s := structures.NewStack()
	res, err := s.Pop()

	if res != nil {
		t.Error("result should be nil")
	}

	if err == nil {
		t.Error("error must not be nil")
	}

	if s.LastIndex() != -1 {
		t.Error("last index of stack should be -1")
	}
}
