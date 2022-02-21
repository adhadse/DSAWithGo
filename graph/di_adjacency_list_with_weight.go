package graph

type DiAdjacencyListWithWeightedNodes struct {
	AdjacencyListWithWeightedNodes
}

func MakeDiGraphWithWeightedEdge(numOfVertices int) DiAdjacencyListWithWeightedNodes {
	g := DiAdjacencyListWithWeightedNodes{
		AdjacencyListWithWeightedNodes: MakeGraphWithWeightedEdge(numOfVertices),
	}
	return g
}

func (g *DiAdjacencyListWithWeightedNodes) AddEdge(src, dest, weight int) {
	destNode := &NodeWithWeight{vertex: dest, weight: weight, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	g.appendWeight(weight)
}

func (g *DiAdjacencyListWithWeightedNodes) Edges() [][]int {
	var edges [][]int
	for forNode := 0; forNode < g.numOfVertices; forNode++ {
		temp := g.adjacencyList[forNode]
		for temp != nil {
			edges = append(edges, []int{forNode, temp.vertex, temp.weight})
			temp = temp.next
		}
	}
	return edges
}
