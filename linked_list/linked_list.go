package linked_list

import "fmt"

type LinkedList struct {
	First      *Node
	last       *Node
	NumOfNodes uint
}

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedListError struct {
	errorString string
}

func (e LinkedListError) Error() string {
	return e.errorString
}

func MakeLinkedList() LinkedList {
	return LinkedList{First: nil, last: nil}
}

func (ll *LinkedList) AddNodeAtFront(data interface{}) {
	oldFirst := ll.First
	ll.First = &Node{Data: data, Next: oldFirst}
	if ll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtEnd don't work without this
		ll.last = ll.First
	}
	ll.NumOfNodes++
}

func (ll *LinkedList) AddNodeAtEnd(data interface{}) {
	newLast := &Node{Data: data, Next: nil}
	if ll.First == nil {
		ll.First = newLast // if no node in link list
	} else {
		// else add new last node pointer to old last instead of nil
		ll.last.Next = newLast
	}
	ll.last = newLast
	ll.NumOfNodes++
}

func (ll *LinkedList) AddNodeAtSpecified(data interface{}, afterNode int) {
	currentNode := ll.First
	if afterNode <= 0 {
		// if insertion never happened previously
		ll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(ll.NumOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &Node{Data: data, Next: currentNode.Next}
			currentNode.Next = newNode
			ll.NumOfNodes++
			return
		}
		currentNode = currentNode.Next
	}
}

// Deletion

func (ll *LinkedList) RemoveNodeAtFront() (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	returnVal := ll.First.Data
	ll.First = ll.First.Next
	ll.NumOfNodes--
	return returnVal, nil
}

func (ll *LinkedList) RemoveNodeAtEnd() (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := ll.First
	var previousNode *Node = nil

	// iterate over linked list
	for currentNode != nil {
		if currentNode.Next == nil {
			// this is last node,
			previousNode.Next = nil
			ll.NumOfNodes--
			return currentNode.Data, nil
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return nil, nil
}

func (ll *LinkedList) RemoveNodeAtSpecified(nodeToDelete interface{}) (interface{}, error) {
	if ll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := ll.First
	var previousNode *Node = nil
	i := 1
	for currentNode != nil {
		if i == nodeToDelete {
			previousNode.Next = currentNode.Next
			ll.NumOfNodes--
			return currentNode.Data, nil
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return nil, nil
}

func (ll *LinkedList) RemoveNodeWithValue(nodeValueToDelete int) error {
	if ll.NumOfNodes == 0 {
		return &LinkedListError{"Linked List is empty!"}
	}
	currentNode := ll.First
	var previousNode *Node = nil

	for currentNode != nil {
		if currentNode.Data == nodeValueToDelete {
			previousNode.Next = currentNode.Next
			ll.NumOfNodes--
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}
	return nil
}

func (ll *LinkedList) PrintLinkedList() {
	currentNode := ll.First
	count := 1
	fmt.Println("First")
	for currentNode != nil {
		fmt.Printf(" (%d) | Node Value:%d\n", count, currentNode.Data)
		count++
		currentNode = currentNode.Next
	}
	fmt.Printf("Last | Num of Nodes: %d\n", ll.NumOfNodes)
}
