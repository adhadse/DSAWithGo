package linked_list

import "testing"

func TestLinkedList_AddNodeAtFront(t *testing.T) {
	ll := MakeLinkedList()
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)
	ll.PrintLinkedList()
}

func TestLinkedList_AddNodeAtEnd(t *testing.T) {
	ll := MakeLinkedList()
	ll.AddNodeAtEnd(2)
	ll.AddNodeAtEnd(1)
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
	ll.RemoveNodeAtFront() // Should print "UNDERFLOW" as linked list is empty

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtFront() // Delete node 1
	ll.RemoveNodeAtFront() // Delete node 2
	ll.RemoveNodeAtFront() // Delete node 3

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeAtEnd(t *testing.T) {
	ll := MakeLinkedList()
	ll.RemoveNodeAtEnd() // Should print "UNDERFLOW" as linked list is empty

	ll.AddNodeAtFront(4)
	ll.AddNodeAtFront(3)
	ll.AddNodeAtFront(2)
	ll.AddNodeAtFront(1)

	ll.RemoveNodeAtEnd() // Delete node 4
	ll.RemoveNodeAtEnd() // Delete node 2
	ll.RemoveNodeAtEnd() // Delete node 3

	ll.PrintLinkedList()
}

func TestLinkedList_RemoveNodeAtSpecified(t *testing.T) {
	ll := MakeLinkedList()
	ll.RemoveNodeAtSpecified(0) // Should print "UNDERFLOW" as linked list is empty

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
	ll.RemoveNodeWithValue(1) // Should print linked list is empty

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
