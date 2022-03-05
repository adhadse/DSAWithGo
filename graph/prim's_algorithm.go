package graph

import "fmt"

// PrimMinimumSpanningTree finds MST by:
//
// - 1. Maintain a min Priority queue for weight edges and visited list
// - 2. Start with source node, mark it visited and add all edges from it
//      to Priority queue
// - 3. While PQ != empty & edgecount of our mst is not equal to one less than # of vertices
//      - Dequeue the next cheapest edge to travel which is not visited
//      - Mark it visited and add it to mst
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
	mst.PrintAdjacencyList()
	fmt.Printf("MST total cost: %d\n", mstCost)
	return &mst, mstCost, nil
}
