package linked_list

import "fmt"

type DoublyLinkedList struct {
	first      *DoublyNode
	last       *DoublyNode
	NumOfNodes uint
}

type DoublyNode struct {
	next     *DoublyNode // pointer to next node in DLL
	previous *DoublyNode // pointer to previous node in DLL
	data     interface{}
}

func MakeDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{first: nil, last: nil}
}

func (dll *DoublyLinkedList) AddNodeAtFront(data interface{}) {
	oldFirst := dll.first
	dll.first = &DoublyNode{next: oldFirst, previous: nil, data: data}
	//dll.first.next.previous = dll.first
	if dll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtBack don't work without this
		dll.last = dll.first
	} else {
		oldFirst.previous = dll.first
	}
	dll.NumOfNodes++
}

func (dll *DoublyLinkedList) AddNodeAtBack(data interface{}) {
	newLast := &DoublyNode{next: nil, previous: dll.last, data: data}
	if dll.first == nil {
		dll.first = newLast // if no node in doubly linked list
	} else {
		// else add new last node pointer to old last instead
		dll.last.next = newLast
	}
	dll.last = newLast
	dll.NumOfNodes++
}

func (dll *DoublyLinkedList) AddNodeAtSpecified(data interface{}, afterNode int) {
	currentNode := dll.first
	if afterNode <= 0 {
		// if insertion never happened previously
		dll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(dll.NumOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &DoublyNode{next: currentNode.next, previous: currentNode.previous, data: data}
			currentNode.next = newNode
			dll.NumOfNodes++
			return
		}
		currentNode = currentNode.next
	}
}

// Deletion

func (dll *DoublyLinkedList) RemoveNodeAtFront() (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	returnVal := dll.first.data
	dll.first = dll.first.next
	dll.first.previous = nil
	dll.NumOfNodes--
	return returnVal, nil
}

func (dll *DoublyLinkedList) RemoveNodeAtBack() (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	prevNodeToLast := dll.last.previous
	returnVal := dll.last.data

	dll.last = prevNodeToLast
	dll.last.next = nil
	dll.NumOfNodes--
	return returnVal, nil
}

func (dll *DoublyLinkedList) RemoveNodeAtSpecified(nodeNum int) (interface{}, error) {
	if dll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := dll.first
	var previousNode *DoublyNode = nil
	i := 1
	for currentNode != nil {
		if i == nodeNum {
			previousNode.next = currentNode.next
			currentNode.previous = nil
			dll.NumOfNodes--
			return currentNode.data, nil
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (dll *DoublyLinkedList) RemoveNodeWithValue(nodeValueToDelete int) error {
	if dll.NumOfNodes == 0 {
		return &LinkedListError{"Linked List is empty!"}
	}
	currentNode := dll.first
	var previousNode *DoublyNode = nil

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

func (dll *DoublyLinkedList) PrintLinkedList() {
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
