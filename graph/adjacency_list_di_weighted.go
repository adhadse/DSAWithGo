package graph

type DiWeightedGraph struct {
	WeightedGraph
}

func MakeDiWeightedGraph(numOfVertices int) DiWeightedGraph {
	g := DiWeightedGraph{
		WeightedGraph: MakeWeightedGraph(numOfVertices),
	}
	return g
}

func (g *DiWeightedGraph) AddEdge(src, dest, weight int) {
	destNode := &NodeWithWeight{vertex: dest, weight: weight, next: g.adjacencyList[src]}
	g.adjacencyList[src] = destNode

	g.appendWeight(weight)
}

func (g *DiWeightedGraph) Edges() [][]int {
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
