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

func (q *Queue) Enqueue(data interface{}) {
	q.AddNodeAtEnd(data)
}

func (q *Queue) Dequeue() (interface{}, error) {
	returnVal, err := q.RemoveNodeAtFront()
	if err != nil {
		return nil, err
	}
	return returnVal, err
}

// IsEmpty returns true if queue is empty
// else returns false
func (q *Queue) IsEmpty() bool {
	if q.NumOfNodes != 0 {
		return false
	}
	return true
}
func (q *Queue) PrintQueue() {
	q.PrintLinkedList()
}