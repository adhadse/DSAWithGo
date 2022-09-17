// Copyright (C) 2022   Anurag Dhadse. All Rights Reserved.
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

package queue

import (
	"DSA/linked_list"
)

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
	q.AddNodeAtBack(data)
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

// IsSizeOne checks if queue has only 1 node
func (q *Queue) IsSizeOne() bool {
	return q.NumOfNodes == 1
}

func (q *Queue) PrintQueue() {
	q.PrintLinkedList()
}
