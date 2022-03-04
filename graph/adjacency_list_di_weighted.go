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

// edges returns a 2 dimensional slice with all
// edges from one node to another in a graph.
//
// For Example for adjacency list (a linked list):
// A -> (B, 5) -> (C, 2)
// B ->
// C ->
//
// Returns:
// [ [A, B, 5],
//   [A, C, 2] ]
func (g *DiWeightedGraph) edges() [][]int {
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
