package linked_list

import "fmt"

type LinkedList struct {
	first      *Node
	last       *Node
	numOfNodes uint
}

type Node struct {
	data int
	next *Node
}

func MakeLinkedList() LinkedList {
	return LinkedList{first: nil, last: nil}
}

func (ll *LinkedList) AddNodeAtFront(data int) {
	oldFirst := ll.first
	ll.first = &Node{data: data, next: oldFirst}
	if ll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtEnd don't work without this
		ll.last = ll.first
	}
	ll.numOfNodes++
}

func (ll *LinkedList) AddNodeAtEnd(data int) {
	newLast := &Node{data: data, next: nil}
	if ll.first == nil {
		ll.first = newLast // if no node in link list
	} else {
		// else add new last node pointer to old last instead of nil
		ll.last.next = newLast
	}
	ll.last = newLast
	ll.numOfNodes++
}

func (ll *LinkedList) AddNodeAtSpecified(data, afterNode int) {
	currentNode := ll.first
	if afterNode <= 0 {
		// if insertion never happened previously
		ll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(ll.numOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &Node{data: data, next: currentNode.next}
			currentNode.next = newNode
			ll.numOfNodes++
			return
		}
		currentNode = currentNode.next
	}
}

// Deletion

func (ll *LinkedList) RemoveNodeAtFront() {
	if ll.numOfNodes == 0 {
		fmt.Println("UNDERFLOW")
		return
	}
	ll.first = ll.first.next
	ll.numOfNodes--
}

func (ll *LinkedList) RemoveNodeAtEnd() {
	if ll.numOfNodes == 0 {
		fmt.Println("UNDERFLOW")
		return
	}
	currentNode := ll.first
	var previousNode *Node = nil

	// iterate over linked list
	for currentNode != nil {
		if currentNode.next == nil {
			// this is last node,
			previousNode.next = nil
			ll.numOfNodes--
			return
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func (ll *LinkedList) RemoveNodeAtSpecified(nodeToDelete int) {
	if ll.numOfNodes == 0 {
		fmt.Println("UNDERFLOW")
		return
	}
	currentNode := ll.first
	var previousNode *Node = nil
	i := 1
	for currentNode != nil {
		if i == nodeToDelete {
			previousNode.next = currentNode.next
			ll.numOfNodes--
			return
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func (ll *LinkedList) RemoveNodeWithValue(nodeValueToDelete int) {
	if ll.numOfNodes == 0 {
		fmt.Println("Linked List is already empty!")
	}
	currentNode := ll.first
	var previousNode *Node = nil

	for currentNode != nil {
		if currentNode.data == nodeValueToDelete {
			previousNode.next = currentNode.next
			ll.numOfNodes--
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
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
	fmt.Printf("Last | Num of Nodes: %d\n", ll.numOfNodes)
}
