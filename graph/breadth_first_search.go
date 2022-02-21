package graph

import (
	"DSA/queue"
)

func (g Graph) BFS(source, destination interface{}) (bool, error) {
	// Queue to ask next node for their adjacency
	nextNodeQueue := queue.MakeQueue()
	// Hashmap to tell which nodes are visited
	// if they are visited, we'll not check its adjacency
	visited := make(map[interface{}]interface{})

	nextNodeQueue.Enqueue(source)
	visited[source] = true

	for nextNodeQueue.IsEmpty() == false {
		currentNode, err := nextNodeQueue.Dequeue()
		if err != nil {
			return false, err
		}
		if currentNode == destination {
			return true, nil
		}
		visited[currentNode] = true

		// Add all nodes which are adjacent to currenNode
		// BUT are not visited yet
		for _, nextNode := range g.GetAdjacencyList(currentNode) {
			if _, found := visited[nextNode]; found != true {
				nextNodeQueue.Enqueue(nextNode.vertex)
			}
		}
	}
	return false, nil
}
