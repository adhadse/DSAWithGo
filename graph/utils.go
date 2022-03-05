package graph

// contains check if slice contains element
// or not
func contains(slice []interface{}, element interface{}) bool {
	for _, val := range slice {
		if val == element {
			return true
		}
	}
	return false
}

// EdgePriorityQueue for Prim's algorithm
type EdgePriorityQueue struct {
	queue []Edge
}

func makeEdgePriorityQueue() EdgePriorityQueue {
	return EdgePriorityQueue{}
}

func (pq *EdgePriorityQueue) enqueue(edge Edge) {
	if len(pq.queue) == 0 {
		pq.queue = append(pq.queue, edge)
	} else {
		pq.queue = append(pq.queue, edge)
		for root := len(pq.queue)/2 - 1; root >= 0; root-- {
			pq.heapify(root, len(pq.queue))
		}
	}
}

func (pq *EdgePriorityQueue) dequeue() Edge {
	minPriorityItem := pq.queue[0]
	pq.queue = append(pq.queue[:0], pq.queue[1:]...)
	for root := len(pq.queue)/2 - 1; root >= 0; root-- {
		pq.heapify(root, len(pq.queue))
	}
	return minPriorityItem
}

func (pq *EdgePriorityQueue) heapify(root, end int) {
	left := 2*root + 1  // left child index of root node
	right := 2*root + 2 // right child index of root node
	min := root         // initialize smallest as root

	// If left child of root exists and is greater than root/min node
	if left < end && pq.queue[left].weight < pq.queue[min].weight {
		min = left
	}
	// If right child of root exists and is greater than root/min node
	if right < end && pq.queue[right].weight < pq.queue[min].weight {
		min = right
	}
	// if min is left or right child, exchange root with min
	if min != root {
		pq.queue[root], pq.queue[min] = pq.queue[min], pq.queue[root]
		pq.heapify(min, end)
	}
}

func (pq *EdgePriorityQueue) isEmpty() bool {
	if len(pq.queue) <= 0 {
		return true
	}
	return false
}
