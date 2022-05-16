// Copyright 2022 The DSAWithGo Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linked_list

import "fmt"

type CircularLinkedList struct {
	first      *Node
	last       *Node
	NumOfNodes uint
}

func MakeCircularLinkedList() CircularLinkedList {
	return CircularLinkedList{first: nil, last: nil}
}

func (cll *CircularLinkedList) AddNodeAtFront(data interface{}) {
	oldFirst := cll.first
	cll.first = &Node{next: oldFirst, data: data}
	if cll.last == nil {
		// When last pointer is null
		// and only inserting at front; last is null
		// AddNodeAtBack don't work without this
		cll.last = cll.first
	} else {
		cll.last.next = cll.first
	}
	cll.NumOfNodes++
}

func (cll *CircularLinkedList) AddNodeAtBack(data interface{}) {
	newLast := &Node{next: cll.first, data: data}
	if cll.first == nil {
		cll.first = newLast // if no node in doubly linked list
	} else {
		// else add new last node pointer to old last instead
		cll.last.next = newLast
	}
	cll.last = newLast
	cll.NumOfNodes++
}

func (cll *CircularLinkedList) AddNodeAtSpecified(data interface{}, afterNode int) {
	currentNode := cll.first
	if afterNode <= 0 {
		// if insertion never happened previously
		cll.AddNodeAtFront(data)
		return
	}
	for i := 1; i <= int(cll.NumOfNodes) && currentNode != nil; i++ {
		if i == afterNode {
			newNode := &Node{next: currentNode.next, data: data}
			currentNode.next = newNode
			cll.NumOfNodes++
			return
		}
		currentNode = currentNode.next
	}
}

// Deletion

func (cll *CircularLinkedList) RemoveNodeAtFront() (interface{}, error) {
	if cll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	returnVal := cll.first.data
	cll.first = cll.first.next
	cll.last.next = cll.first
	cll.NumOfNodes--
	return returnVal, nil
}

func (cll *CircularLinkedList) RemoveNodeAtBack() (interface{}, error) {
	if cll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := cll.first
	previousNode := cll.last
	returnVal := cll.last.data

	for currentNode != cll.last {
		previousNode = currentNode
		currentNode = currentNode.next
	}
	cll.last = previousNode
	cll.last.next = cll.first
	cll.NumOfNodes--
	return returnVal, nil
}

func (cll *CircularLinkedList) RemoveNodeAtSpecified(nodeNum int) (interface{}, error) {
	if cll.NumOfNodes == 0 {
		return nil, &LinkedListError{"UNDERFLOW"}
	}
	currentNode := cll.first
	var previousNode *Node = nil
	i := 1
	for currentNode != nil {
		if i == nodeNum {
			previousNode.next = currentNode.next
			cll.NumOfNodes--
			return currentNode.data, nil
		}
		i += 1
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil, nil
}

func (cll *CircularLinkedList) RemoveNodeWithValue(nodeValueToDelete int) error {
	if cll.NumOfNodes == 0 {
		return &LinkedListError{"Linked List is empty!"}
	}
	currentNode := cll.first
	var previousNode *Node = nil

	for currentNode.next != cll.first {
		if currentNode.data == nodeValueToDelete && currentNode != cll.last {
			previousNode.next = currentNode.next
			if currentNode == cll.last {
				cll.last = previousNode
			}
			cll.NumOfNodes--
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return nil
}

func (cll *CircularLinkedList) PrintLinkedList() {
	currentNode := cll.first
	count := 1
	fmt.Println("first")
	for {
		fmt.Printf(" (%d) | Node Value:%d\n", count, currentNode.data)
		count++
		currentNode = currentNode.next
		// check if other nodes exist or there is only one item
		if currentNode == cll.first || currentNode == nil {
			break
		}
	}
	fmt.Printf("Last | Num of Nodes: %d\n", cll.NumOfNodes)
}
