package linked_list

import "fmt"

type LinkedList struct {
	first      *Node
	last       *Node
	NumOfNodes uint
}

type Node struct {
	data interface{}
	next *Node
}

type LinkedListError struct {
	errorString string
}

func (e LinkedListError) Error() string {
	return e.errorString
}

func MakeLinkedList() LinkedList {
	return LinkedList{first: nil, last: nil}
}

func (ll *LinkedList) AddNodeAtFront(data interface{}) {
	oldFirst := ll.first
	ll.first = &Node{data: data, next: oldFirst}
	if ll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtEnd don't work without this
		ll.last = ll.first
	}
	ll.NumOfNodes++
}

func (ll *LinkedList) AddNodeAtEnd(data interface{}) {
	newLast := &Node{data: data, next: nil}
	if ll.first == nil {
		ll.first = newLast // if no node in link list
	} else {
		// else add new last node pointer to old last instead of nil
		ll.last.next = newLast
	}
	ll.last = newLast
	ll.NumOfNodes++
}

func (ll *LinkedList) AddNodeAtSpecified(data interface{}, afterNode int) {
	currentNode := ll.first
	if afterNode <= 0 {
		// if insertion never happened previously
		ll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(ll.NumOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &Node{data: data, next: currentNode.next}
			currentNode.next = newNode
			ll.NumOfNodes++
			return
		}
		currentNode = currentNode.next
	}
}

// Deletion

func (ll *LinkedList) RemoveNodeAtFront() (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	returnVal := ll.first.data
	ll.first = ll.first.next
	ll.NumOfNodes--
	return returnVal, nil
}

func (ll *LinkedList) RemoveNodeAtEnd() (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := ll.first
	var previousNode *Node = nil

	// iterate over linked list
	for currentNode != nil {
		if currentNode.next == nil {
			// this is last node,
			previousNode.next = nil
			ll.NumOfNodes--
			return currentNode.data, nil
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (ll *LinkedList) RemoveNodeAtSpecified(nodeToDelete interface{}) (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := ll.first
	var previousNode *Node = nil
	i := 1
	for currentNode != nil {
		if i == nodeToDelete {
			previousNode.next = currentNode.next
			ll.NumOfNodes--
			return currentNode.data, nil
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (ll *LinkedList) RemoveNodeWithValue(nodeValueToDelete int) error {
	if ll.NumOfNodes == 0 {
		return &LinkedListError{"Linked List is empty!"}
	}
	currentNode := ll.first
	var previousNode *Node = nil

	for currentNode != nil {
		if currentNode.data == nodeValueToDelete {
			previousNode.next = currentNode.next
			ll.NumOfNodes--
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil
}

func (ll *LinkedList) PrintLinkedList() {
	currentNode := ll.first
	count := 1
	fmt.Println("First")
	for currentNode != nil {
		fmt.Printf(" (%d) | Node Value:%d\n", count, currentNode.data)
		count++
		currentNode = currentNode.next
	}
	fmt.Printf("Last | Num of Nodes: %d\n", ll.NumOfNodes)
}
