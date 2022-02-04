package graph

import "fmt"

// AdjacencyList is a collection of unordered list (here linked lists)
// to represent a finite graph.
// Each index of adjacencyList is a linked list storing nodes approachable by
// this index value with each node storing next node address, i.e.,
// [ 0 ] : 4 -> 1  (means 0 has an edge to 4 and 1)
// [ 1 ] : 0       (means 1 has an edge to 0)
// [ 2 ]
// [ 3 ]
// [ 4 ] : 0       (means 4 has an edge to 0)
type AdjacencyList struct {
	numOfVertices int
	adjacencyList []*Node
}

type Node struct {
	vertex int
	next   *Node
}

func MakeAdjacencyList(numOfVertices int) AdjacencyList {
	g := AdjacencyList{
		adjacencyList: make([]*Node, numOfVertices),
		numOfVertices: numOfVertices}
	return g
}

func (g AdjacencyList) AddEdge(src, dest int) {
	destNode := &Node{vertex: dest, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	srcNode := &Node{vertex: src, next: g.adjacencyList[dest]}
	g.adjacencyList[dest] = srcNode
}

func (g AdjacencyList) PrintAdjacencyList() {
	for i := 0; i < g.numOfVertices; i++ {
		fmt.Printf("Adjacency List of vertex %d| head", i)
		temp := g.adjacencyList[i]
		for temp != nil {
			fmt.Printf(" -> %d", temp.vertex)
			temp = temp.next
		}
		println()
	}
}
