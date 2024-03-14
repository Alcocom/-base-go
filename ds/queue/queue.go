package queue

import (
	"bytes"
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	datastorage *list.List
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		datastorage: list.New(),
	}
}

func (q *Queue[T]) EnQueue(data T) {
	q.datastorage.PushBack(data) // Rear
}

func (q *Queue[T]) DeQueue() T {
	e := q.datastorage.Front() // Head
	q.datastorage.Remove(e)
	return e.Value.(T)
}

func (q *Queue[T]) Peek() T {
	return q.datastorage.Front().Value.(T)
}

func (q *Queue[T]) Size() int {
	return q.datastorage.Len()
}

func (q *Queue[T]) String() string {
	var buf bytes.Buffer // buf := bytes.Buffer{}

	buf.WriteString("[ *")
	for e := q.datastorage.Front(); e != nil; e = e.Next() {

		v := e.Value.(T)
		buf.WriteString(fmt.Sprintf("%v", v))

		if e.Next() != nil {
			buf.WriteString(" -> ")
		}
	}
	buf.WriteString(" ]")

	return buf.String()
}

// TODO
/*
Clear()
isEmpty()
Contains()
*/
