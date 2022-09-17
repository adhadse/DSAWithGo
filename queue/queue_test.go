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
	"fmt"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.PrintQueue()
}

func TestQueue_Dequeue(t *testing.T) {
	q := MakeQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	_, err := q.Dequeue()
	if err != nil {
		return
	}
	q.PrintQueue()
}

func TestQueue_IsEmpty(t *testing.T) {
	q := MakeQueue()
	if val := q.IsEmpty(); val != true {
		t.Errorf("IsEmpty() returned false when Queue was Empty")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if val := q.IsEmpty(); val != false {
		t.Errorf("IsEmpty() returned true when Queue was not Empty")
	}
}

// Priority Queue
func TestPriorityQueue_Enqueue(t *testing.T) {
	pq := MakePriorityQueue()
	pq.Enqueue(3)
	pq.Enqueue(5)
	pq.Enqueue(9)
	pq.Enqueue(10)
	pq.Enqueue(2)

	fmt.Printf("%v\n", pq.queue)
}

func TestPriorityQueue_Poll(t *testing.T) {
	pq := MakePriorityQueue()
	pq.Enqueue(3)
	pq.Enqueue(5)
	pq.Enqueue(9)
	pq.Enqueue(10)
	pq.Enqueue(2)

	fmt.Printf("%v\n", pq.queue)

	fmt.Printf("After first dequeue item returned: %d, %v\n", pq.Dequeue(), pq.queue)
	fmt.Printf("After second dequeue item returned: %d, %v\n", pq.Dequeue(), pq.queue)
}
