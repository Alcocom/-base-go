package queue_test

import (
	"testing"

	"github.com/alcocom/base-go/ds/queue"
)

func TestAllFlow(t *testing.T) {
	q := queue.New[int]()

	q.EnQueue(10)
	q.EnQueue(20)
	q.EnQueue(30)

	t.Log(q) // test log - need to delete

	if q.Size() != 3 {
		t.Error()
	}

	if q.Peek() != 10 {
		t.Error()
	}

	v1_1 := q.DeQueue()
	v1_2 := q.DeQueue()

	t.Log(q) // test log - need to delete

	if q.Size() != 1 {
		t.Error()
	}

	if q.Peek() != 30 {
		t.Error()
	}

	if v1_1 != 10 {
		t.Error()
	}

	if v1_2 != 20 {
		t.Error()
	}

	q.EnQueue(40)
	q.EnQueue(50)
	t.Log(q) // test log - need to delete

	if q.Size() != 3 {
		t.Error()
	}

	if q.Peek() != 30 {
		t.Error()
	}
}
