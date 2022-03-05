package queue

import (
	"fmt"
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

// Priority Queue
func TestPriorityQueue_Enqueue(t *testing.T) {
	pq := MakePriorityQueue()
	pq.Enqueue(3)
	pq.Enqueue(5)
	pq.Enqueue(9)
	pq.Enqueue(10)
	pq.Enqueue(2)

	fmt.Printf("%v\n", pq.queue)
}

func TestPriorityQueue_Poll(t *testing.T) {
	pq := MakePriorityQueue()
	pq.Enqueue(3)
	pq.Enqueue(5)
	pq.Enqueue(9)
	pq.Enqueue(10)
	pq.Enqueue(2)

	fmt.Printf("%v\n", pq.queue)

	fmt.Printf("After first poll item returned: %d, %v\n", pq.Poll(), pq.queue)
	fmt.Printf("After second poll item returned: %d, %v\n", pq.Poll(), pq.queue)
}
