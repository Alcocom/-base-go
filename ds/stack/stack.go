package stack

import (
	"bytes"
	"container/list"
	"fmt"
)

type Stack[T any] struct {
	datastorage *list.List
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		datastorage: list.New(),
	}
}

func (s *Stack[T]) Push(data T) {
	s.datastorage.PushBack(data)
}

func (s *Stack[T]) Pop() T {
	e := s.datastorage.Back()
	s.datastorage.Remove(e)
	return e.Value.(T)
}

func (s *Stack[T]) Peek() T {
	return s.datastorage.Back().Value.(T)
}

func (s *Stack[T]) Size() int {
	return s.datastorage.Len()
}

func (s *Stack[T]) String() string {
	var buf bytes.Buffer // buf := bytes.Buffer{}

	buf.WriteString("[ ")
	for e := s.datastorage.Front(); e != nil; e = e.Next() {

		v := e.Value.(T)
		buf.WriteString(fmt.Sprintf("%v", v))

		if e.Next() != nil {
			buf.WriteString(" -> ")
		}
	}
	buf.WriteString("* ]")

	return buf.String()
}

// TODO
/*
Clear()
isEmpty()
Contains()
*/
