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

package graph

// PrimMinimumSpanningTree finds MST by:
//
// - 1. Maintain a min Priority queue for weight edges and visited list
// - 2. Start with source node, mark it visited and add all edges from it
//      to Priority queue
// - 3. While PQ != empty & edgeCount of our mst is not equal to one less than # of vertices
//      - Dequeue the next cheapest edge to travel which is not visited
//      - Mark it visited and add it to mst
//      - Add all edges from this node to PQ
// Time Complexity: O(E log(E))
// https://youtu.be/jsmMtJpPnhU
func (g *WeightedGraph) PrimMinimumSpanningTree(source int) (*WeightedGraph, int, error) {
	pq := makeEdgePriorityQueue()
	visited := make([]bool, g.numOfVertices)
	mst := MakeWeightedGraph(g.numOfVertices)
	edgeCount, mstCost := 0, 0

	visited[source] = true
	for _, edge := range g.GetEdges(source) {
		if visited[edge.to] == false {
			pq.enqueue(edge)
		}
	}

	for pq.isEmpty() == false && edgeCount != g.numOfVertices-1 {
		edge := pq.dequeue()
		if visited[edge.to] {
			continue
		}

		mst.AddEdge(edge.from, edge.to, edge.weight)
		mstCost += edge.weight
		edgeCount++

		visited[edge.to] = true
		for _, edge := range g.GetEdges(edge.to) {
			if visited[edge.to] == false {
				pq.enqueue(edge)
			}
		}
	}

	if edgeCount != g.numOfVertices-1 {
		return nil, -1, WeightedGraphError{"No Mst Exist\n"}
	}
	return &mst, mstCost, nil
}
