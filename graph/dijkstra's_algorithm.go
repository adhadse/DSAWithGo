package graph

import (
	"fmt"
	"math"
)

// Dijkstra algorithm find the shortest path and lowest cost
// from one node to all other node given a source node with an amazing
// time complexity of O((V+E)log(V)) but only works if weights are non-negative.
// https://youtu.be/EFg3u_E6eHU
//
// Algorithm:
// 	- 1. Label each node with the shortest time found so far,
//		 In beginning wih 0 for source and other with '∞'.
//  - 2. Mark src as 'EXPLORED' and other as 'UNEXPLORED'
//  - 3. Loop until reached destination node:
//       - Update estimates if found smaller value than current estimate.
//       - Keep track of what previous node visited to get to this node.
//         Which nodes led to this shortest path.
//       - Choose next 'UNEXPLORED' vertex with the smallest estimate

func (g *WeightedGraph) Dijkstra(source, destination int) {
	dist := make([]int, g.numOfVertices)
	visited := make([]bool, g.numOfVertices)
	followedBy := make([]int, g.numOfVertices) // Keep tracks of what previous node led to this node
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[source] = 0
	visited[source] = true

	node := g.adjacencyList[source]
	currentlyVisited := source
	nextToVisit := node.vertex

	for source != destination {
		for node != nil {
			// Update estimates if found smaller value than current estimate.
			if dist[currentlyVisited]+node.weight < dist[node.vertex] &&
				visited[node.vertex] == false {
				dist[node.vertex] = dist[currentlyVisited] + node.weight
				followedBy[node.vertex] = currentlyVisited
			}

			// Choose next 'UNEXPLORED' vertex with the smallest estimate
			if nextToVisit == currentlyVisited && visited[node.vertex] == false ||
				dist[node.vertex] < dist[nextToVisit] &&
					nextToVisit != currentlyVisited &&
					visited[node.vertex] == false {
				nextToVisit = node.vertex
			}
			node = node.next
		}
		if nextToVisit == destination || nextToVisit == currentlyVisited {
			break
		}
		node = g.adjacencyList[nextToVisit]
		currentlyVisited = nextToVisit
		visited[nextToVisit] = true
	}
	// print result
	fmt.Printf("Shortest path distance from %d -> %d is: %d\n", source, destination, dist[destination])

	var path []int
	intermediate := destination
	for source != destination {
		if followedBy[intermediate] == source {
			break
		}
		intermediate = followedBy[intermediate]
		path = append(path, intermediate)
	}
	fmt.Printf("Path Followed: %d ->", source)
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Printf(" %d ->", path[i])
	}
	fmt.Printf(" %d\n", destination)
}
