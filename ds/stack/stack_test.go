package stack_test

import (
	"testing"

	"github.com/alcocom/base-go/ds/stack"
)

func TestAllFlow(t *testing.T) {
	s := stack.New[int]()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Peek() != 30 {
		t.Error()
	}

	if s.Size() != 3 {
		t.Error()
	}

	v1 := s.Pop()

	if s.Size() != 2 {
		t.Error()
	}

	if s.Peek() != 20 {
		t.Error()
	}

	if v1 != 30 {
		t.Error()
	}

	s.Push(40)
	s.Push(50)

	if s.Size() != 4 {
		t.Error()
	}

	if s.Peek() != 50 {
		t.Error()
	}

}
