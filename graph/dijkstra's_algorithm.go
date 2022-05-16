// Copyright 2022 The  Authors. All Rights Reserved.
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

package graph

import (
	"fmt"
	"math"
)

// DijkstraShortestPath find the shortest path and lowest cost
// from one node to all other node given a source node with an amazing
// time complexity of O((V+E)log(V)) but only works if weights are non-negative.
// https://youtu.be/EFg3u_E6eHU
//
// Algorithm:
// 	- 1. Label each node with the shortest path found so far,
//		 In beginning wih 0 for source and other with '∞'.
//  - 2. Mark src as 'EXPLORED' and other as 'UNEXPLORED'
//  - 3. Loop until reached destination node:
//       - Update estimates if found smaller value than current estimate.
//       - Keep track of what previous node visited to get to this node.
//         Which nodes led to this shortest path.
//       - Choose next 'UNEXPLORED' vertex with the smallest estimate
func (g *WeightedGraph) DijkstraShortestPath(source, destination int) {
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
			// First of OR condition checks when you visit nextToVisit
			// currentlyVisited becomes nextToVisit; that mean nextToVisit is not known now
			if nextToVisit == currentlyVisited && visited[node.vertex] == false ||
				dist[node.vertex] < dist[nextToVisit] &&
					nextToVisit != currentlyVisited &&
					visited[node.vertex] == false {
				nextToVisit = node.vertex
			}
			node = node.next
		}
		// Second condition of OR applies when we are not able to choose next 'UNEXPLORED'
		// node; i.e., we reached a node from where all nodes are already visited
		if nextToVisit == destination || nextToVisit == currentlyVisited {
			break
		}
		node = g.adjacencyList[nextToVisit]
		currentlyVisited = nextToVisit
		visited[nextToVisit] = true
	}

	// print result
	fmt.Printf("Shortest path distance from %d -> %d is: %d\n", source, destination, dist[destination])

	// Path finding using followedBy can be done by reverse iterating starting
	// from destination by indexing it and then keep indexing the intermediate
	// value until reached source: src(4) dest(0) => 4->3->2->0
	// index: 0 1 2 3 4 5 6
	// value: 0 0 0 2 3 2 1
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
