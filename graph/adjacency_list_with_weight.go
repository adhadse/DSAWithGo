package graph

type AdjacencyListWithWeightedNodes struct {
	numOfVertices int
	weights       []int
	adjacencyList []*NodeWithWeight
}

type NodeWithWeight struct {
	vertex int
	weight int
	next   *NodeWithWeight
}

func MakeAdjacencyListWithWeightedNodes(numOfVertices int) AdjacencyListWithWeightedNodes {
	g := AdjacencyListWithWeightedNodes{
		adjacencyList: make([]*NodeWithWeight, numOfVertices),
		numOfVertices: numOfVertices}
	return g
}

func (g AdjacencyListWithWeightedNodes) AddEdge(src, dest, weight int) {
	destNode := &NodeWithWeight{vertex: dest, weight: weight, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	srcNode := &NodeWithWeight{vertex: src, weight: weight, next: g.adjacencyList[dest]}
	g.adjacencyList[dest] = srcNode

	g.appendWeight(weight)
}

// appendWeight append unique weight to g.weights
func (g AdjacencyListWithWeightedNodes) appendWeight(weight int) {
	for i := range g.weights {
		if g.weights[i] == weight {
			return
		}
	}
	g.weights = append(g.weights, weight)
}

func (g AdjacencyListWithWeightedNodes) SortAdjacencyListBasedOnWeights() {
	sortedAdjacencyList := MakeAdjacencyListWithWeightedNodes(g.numOfVertices)

}

func (g AdjacencyListWithWeightedNodes) GetAdjacencyList(forNode int) []*NodeWithWeight {
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
//func (g AdjacencyList) Find(parent, i int) {
//	if parent[i] == i {
//		return i
//	}
//	returnf g.Find(parent, parent[i])
//}
//
//// Union merges two groups of Nodes together
//func (g AdjacencyList) Union() {
//	xRoot = g.Find(parent, x)
//	yRoot = g.Find(parent, y)
//
//	if rank[xRoot] < rank[yRoot] {
//		parent[xRoot] = yRoot
//	} else rank[xRoot] > rank[yRoot] {
//		parent[yRoot] = xRoot
//	} else {
//		parent[yRoot] = xRoot
//		rank[xRoot]++
//	}
//}
//
