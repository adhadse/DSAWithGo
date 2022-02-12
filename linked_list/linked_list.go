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

func (dll *LinkedList) AddNodeAtFront(data interface{}) {
	oldFirst := dll.first
	dll.first = &Node{data: data, next: oldFirst}
	if dll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtBack don't work without this
		dll.last = dll.first
	}
	dll.NumOfNodes++
}

func (dll *LinkedList) AddNodeAtBack(data interface{}) {
	newLast := &Node{data: data, next: nil}
	if dll.first == nil {
		dll.first = newLast // if no node in link list
	} else {
		// else add new last node pointer to old last instead of nil
		dll.last.next = newLast
	}
	dll.last = newLast
	dll.NumOfNodes++
}

func (dll *LinkedList) AddNodeAtSpecified(data interface{}, afterNode int) {
	currentNode := dll.first
	if afterNode <= 0 {
		// if insertion never happened previously
		dll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(dll.NumOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &Node{data: data, next: currentNode.next}
			currentNode.next = newNode
			dll.NumOfNodes++
			return
		}
		currentNode = currentNode.next
	}
}

// Deletion

func (dll *LinkedList) RemoveNodeAtFront() (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	returnVal := dll.first.data
	dll.first = dll.first.next
	dll.NumOfNodes--
	return returnVal, nil
}

func (dll *LinkedList) RemoveNodeAtBack() (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := dll.first
	var previousNode *Node = nil

	// iterate over linked list
	for currentNode != nil {
		if currentNode.next == nil {
			// this is last node,
			previousNode.next = nil
			dll.NumOfNodes--
			return currentNode.data, nil
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (dll *LinkedList) RemoveNodeAtSpecified(nodeNum int) (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := dll.first
	var previousNode *Node = nil
	i := 1
	for currentNode != nil {
		if i == nodeNum {
			previousNode.next = currentNode.next
			dll.NumOfNodes--
			return currentNode.data, nil
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (dll *LinkedList) RemoveNodeWithValue(nodeValueToDelete int) error {
	if dll.NumOfNodes == 0 {
		return &LinkedListError{"Linked List is empty!"}
	}
	currentNode := dll.first
	var previousNode *Node = nil

	for currentNode != nil {
		if currentNode.data == nodeValueToDelete {
			previousNode.next = currentNode.next
			dll.NumOfNodes--
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil
}

func (dll *LinkedList) PrintLinkedList() {
	currentNode := dll.first
	count := 1
	fmt.Println("first")
	for currentNode != nil {
		fmt.Printf(" (%d) | Node Value:%d\n", count, currentNode.data)
		count++
		currentNode = currentNode.next
	}
	fmt.Printf("Last | Num of Nodes: %d\n", dll.NumOfNodes)
}
