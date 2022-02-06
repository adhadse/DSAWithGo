package queue

import "DSA/linked_list"

type Queue struct {
	// Use composition through
	// embedding to add the fields of Linked List
	// struct to Queue struct. Reuse Linked List as Queue.
	linked_list.LinkedList
}

func MakeQueue() Queue {
	return Queue{LinkedList: linked_list.MakeLinkedList()}
}

func (q *Queue) Enqueue(data int) {
	q.AddNodeAtEnd(data)
}

func (q *Queue) Dequeue() (int, error) {
	returnVal, err := q.RemoveNodeAtFront()
	if err != nil {
		return 0, err
	}
	return returnVal, err
}

func (q *Queue) PrintQueue() {
	q.PrintLinkedList()
}
