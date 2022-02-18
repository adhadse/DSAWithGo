package graph

// A Minimum Spanning Tree is a subset of the edges
// which connect all vertices in the graph with the
// minimal total edge cost

// KruskalMinimumSpanningTree finds a MST by:
// 1. Sort edges by ascending edge weight
// 2. Walk through the sorted edges and look at two nodes the edge belong to,
// 	  if the nodes are already unified/belongs to same group:
//	  		we don't include this edge, (avoid creating cycle)
//    else:
// 			we include it unify the nodes
// 3. Terminate when every edge has been processed OR all vertices have been unified
//func (g AdjacencyListWithWeightedNodes) KruskalMinimumSpanningTree() {
//	var result []int
//	i, e := 0, 0
//
//}
