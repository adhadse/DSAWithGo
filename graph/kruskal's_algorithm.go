package graph

import "fmt"

// A Minimum Spanning Tree is a subset of the edges
// which connect all vertices in the graph with the
// minimal total edge cost with no cycles (otherwise it is not tree)
// Time complexity: O(E logE) where E is the no. of edges
// https://youtu.be/JZBQLXgSGfs

// KruskalMinimumSpanningTree finds a MST by:
// - 1. Sort edges by ascending edge weight
// - 2. Walk through the sorted edges and look at two nodes the edge belong to,
// 	  if the nodes are already unified/belongs to same group:
//	  		we don't include this edge, (avoid creating cycle)
//    else:
// 			we include/unify the nodes to one group
// 3. Terminate when every edge has been processed OR all vertices have been unified
func (g WeightedGraph) KruskalMinimumSpanningTree() {
	var result [][]int
	var parent []int
	var rank []int
	i, e := 0, 0
	graph := g.getSortedEdgesBasedOnWeights()

	for node := 0; node < g.numOfVertices; node++ {
		parent = append(parent, node)
		rank = append(rank, 0)
	}

	for e < g.numOfVertices-1 {
		// Dammit, find the group?
		edge := graph[i]
		i++
		group1 := g.find(parent, edge[0])
		group2 := g.find(parent, edge[1])

		if group1 != group2 {
			// if the nodes do not belong to same group; UNIFY THEM!!!
			e++
			result = append(result, []int{edge[0], edge[1], edge[2]})
			parent, rank = g.union(parent, rank, group1, group2)
		}
	}
	minimumCost := 0
	fmt.Println("edges in the constructed MST")
	for _, edge := range result {
		minimumCost += edge[2]
		fmt.Printf("%d -- %d == %d\n", edge[0], edge[1], edge[2])
	}
	fmt.Printf("Minimum Spanning Tree cost: %d\n", minimumCost)
}
