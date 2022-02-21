package graph

type DiGraphWithEdgeWeight struct {
	GraphWithEdgeWeight
}

func MakeDiGraphWithWeightedEdge(numOfVertices int) DiGraphWithEdgeWeight {
	g := DiGraphWithEdgeWeight{
		GraphWithEdgeWeight: MakeGraphWithWeightedEdge(numOfVertices),
	}
	return g
}

func (g *DiGraphWithEdgeWeight) AddEdge(src, dest, weight int) {
	destNode := &NodeWithWeight{vertex: dest, weight: weight, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	g.appendWeight(weight)
}

func (g *DiGraphWithEdgeWeight) Edges() [][]int {
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
