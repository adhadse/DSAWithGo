package queue

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.PrintQueue()
}

func TestQueue_Dequeue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	_, err := q.Dequeue()
	if err != nil {
		return
	}
	q.PrintQueue()
}

func TestQueue_IsEmpty(t *testing.T) {
	q := MakeQueue()
	if val := q.IsEmpty(); val != true {
		t.Errorf("IsEmpty() returned false when Queue was Empty")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if val := q.IsEmpty(); val != false {
		t.Errorf("IsEmpty() returned true when Queue was not Empty")
	}
}