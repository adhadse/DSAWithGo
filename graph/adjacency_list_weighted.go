package graph

import (
	"DSA/sorting"
)

type WeightedGraph struct {
	numOfVertices int
	weights       []int
	adjacencyList []*NodeWithWeight
}

type NodeWithWeight struct {
	vertex int
	weight int
	next   *NodeWithWeight
}

func MakeWeightedGraph(numOfVertices int) WeightedGraph {
	g := WeightedGraph{
		adjacencyList: make([]*NodeWithWeight, numOfVertices),
		numOfVertices: numOfVertices}
	return g
}

func (g *WeightedGraph) AddEdge(src, dest, weight int) {
	destNode := &NodeWithWeight{vertex: dest, weight: weight, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	srcNode := &NodeWithWeight{vertex: src, weight: weight, next: g.adjacencyList[dest]}
	g.adjacencyList[dest] = srcNode

	g.appendWeight(weight)
}

// appendWeight append unique weight to g.weights
func (g *WeightedGraph) appendWeight(weight int) {
	for i := range g.weights {
		if g.weights[i] == weight {
			return
		}
	}
	g.weights = append(g.weights, weight)
}

// GetSortedEdgesBasedOnWeights get a 2 dimensional slice with all
// possible combination (repeating edges) of edges from one node to another sorted in order of weights.
//
// For Example for adjacency list (a linked list):
// A -> (B, 5) -> (C, 2)
// B -> (A, 5)
// C -> (A, 2)
//
// Returns:
// [ [A, C, 2],
//   [C, A, 2],
//	 [B, A, 5],
//   [A, B, 5] ]
func (g *WeightedGraph) GetSortedEdgesBasedOnWeights() [][]int {
	var sortedAdjacencyList [][]int
	g.weights = sorting.InsertionSort(g.weights)
	for _, weight := range g.weights {
		for forNode := 0; forNode < g.numOfVertices; forNode++ {
			temp := g.adjacencyList[forNode]
			for temp != nil {
				if temp.weight == weight {
					sortedAdjacencyList = append(sortedAdjacencyList, []int{forNode, temp.vertex, temp.weight})
				}
				temp = temp.next
			}
		}
	}
	return sortedAdjacencyList
}

// GetAdjacencyList returns a slice containing all nodes
// adjacent to forNode
func (g *WeightedGraph) GetAdjacencyList(forNode int) []*NodeWithWeight {
	var adjacencyList []*NodeWithWeight
	temp := g.adjacencyList[forNode]
	for temp != nil {
		adjacencyList = append(adjacencyList, temp)
		temp = temp.next
	}
	return adjacencyList
}

// Union Find is a data structure that keeps tracks of
// elements which are split into one or more disjoint sets.
// It's two primary operations are Union and Find.
// https://youtu.be/ibjEGG7ylHk

// Find finds the group the Node belongs to
func (g *WeightedGraph) Find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return g.Find(parent, parent[i])
}

// Union merges two groups of Nodes together
func (g *WeightedGraph) Union(parent, rank []int, group1, group2 int) ([]int, []int) {
	xRoot := g.Find(parent, group1)
	yRoot := g.Find(parent, group2)

	if rank[xRoot] < rank[yRoot] {
		parent[xRoot] = yRoot
	} else if rank[xRoot] > rank[yRoot] {
		parent[yRoot] = xRoot
	} else {
		parent[yRoot] = xRoot
		rank[xRoot]++
	}
	return parent, rank
}

// getDistance returns the distance (edge weight) from
// one node to another
func (g *WeightedGraph) getDistance(from, to int) int {
	temp := g.adjacencyList[from]
	for temp != nil {
		if temp.vertex == to {
			return temp.weight
		}
		temp = temp.next
	}
	return 0
}
