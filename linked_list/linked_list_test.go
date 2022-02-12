package linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList_AddNodeAtFront(t *testing.T) {
	ll := MakeLinkedList()
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)
	ll.PrintLinkedList()
}

func TestLinkedList_AddNodeAtBack(t *testing.T) {
	ll := MakeLinkedList()
	ll.AddNodeAtBack(2)
	ll.AddNodeAtBack(1)
	ll.PrintLinkedList()
}

func TestLinkedList_AddNodeAtSpecified(t *testing.T) {
	ll := MakeLinkedList()
	ll.AddNodeAtSpecified(-1, 0) // As the first node

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.AddNodeAtSpecified(32, 3) // after node with value 3
	ll.AddNodeAtSpecified(42, 6) // at end of linked list

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeAtFront(t *testing.T) {
	ll := MakeLinkedList()
	_, err := ll.RemoveNodeAtFront()
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtFront() // Delete node 1
	ll.RemoveNodeAtFront() // Delete node 2
	ll.RemoveNodeAtFront() // Delete node 3

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeAtBack(t *testing.T) {
	ll := MakeLinkedList()
	_, err := ll.RemoveNodeAtBack()
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtBack() // Delete node 4
	ll.RemoveNodeAtBack() // Delete node 3
	ll.RemoveNodeAtBack() // Delete node 2

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeAtSpecified(t *testing.T) {
	ll := MakeLinkedList()
	_, err := ll.RemoveNodeAtSpecified(0)
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtSpecified(2) // after node with value 2
	ll.RemoveNodeAtSpecified(3) // at end of linked list with value 4

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeWithValue(t *testing.T) {
	ll := MakeLinkedList()
	err := ll.RemoveNodeWithValue(1)
	if err != nil {
		fmt.Println(err) // Should print "Linked list is empty" as error
		//return         // Or you can handle your way
	}
	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeWithValue(4)
	ll.RemoveNodeWithValue(2)

	ll.PrintLinkedList()
}

// Doubly linked list test
func TestDoublyLinkedList_AddNodeAtFront(t *testing.T) {
	ll := MakeDoublyLinkedList()
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)
	ll.PrintLinkedList()
}

func TestDoublyLinkedList_AddNodeAtBack(t *testing.T) {
	ll := MakeDoublyLinkedList()
	ll.AddNodeAtBack(2)
	ll.AddNodeAtBack(1)
	ll.PrintLinkedList()
}

func TestDoublyLinkedList_AddNodeAtSpecified(t *testing.T) {
	ll := MakeDoublyLinkedList()
	ll.AddNodeAtSpecified(-1, 0) // As the first node

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.AddNodeAtSpecified(32, 3) // after node with value 3
	ll.AddNodeAtSpecified(42, 6) // at end of linked list

	ll.PrintLinkedList()
}

func TestDoublyLinkedList_RemoveNodeAtFront(t *testing.T) {
	ll := MakeDoublyLinkedList()
	_, err := ll.RemoveNodeAtFront()
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtFront() // Delete node 1
	ll.RemoveNodeAtFront() // Delete node 2
	ll.RemoveNodeAtFront() // Delete node 3

	ll.PrintLinkedList()
}

func TestDoublyLinkedList_RemoveNodeAtBack(t *testing.T) {
	ll := MakeDoublyLinkedList()
	_, err := ll.RemoveNodeAtBack()
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtBack() // Delete node 4
	ll.RemoveNodeAtBack() // Delete node 3
	ll.RemoveNodeAtBack() // Delete node 2

	ll.PrintLinkedList()
}

func TestDoublyLinkedList_RemoveNodeAtSpecified(t *testing.T) {
	ll := MakeDoublyLinkedList()
	_, err := ll.RemoveNodeAtSpecified(0)
	if err != nil {
		fmt.Println(err) // Should print "UNDERFLOW" as linked list is empty
		//return         // Or you can handle your way
	}

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtSpecified(2) // after node with value 2
	ll.RemoveNodeAtSpecified(3) // at end of linked list with value 4

	ll.PrintLinkedList()
}

func TestDoublyLinkedList_RemoveNodeWithValue(t *testing.T) {
	ll := MakeDoublyLinkedList()
	err := ll.RemoveNodeWithValue(1)
	if err != nil {
		fmt.Println(err) // Should print "Linked list is empty" as error
		//return         // Or you can handle your way
	}
	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeWithValue(4)
	ll.RemoveNodeWithValue(2)

	ll.PrintLinkedList()
}
