package graph

// DFS performs a depth-first search on Adjacency List
//
// Parameters
// 	- `source`: source vertex
//  - `destination`: destination vertex
//  - `asked`: a set of all nodes that are already checked
func (g *Graph) DFS(source, destination interface{}, asked ...interface{}) bool {
	if contains(asked, source) {
		return false
	}
	asked = append(asked, source)
	if source == destination {
		return true // stopping condition
	}
	for _, adjacentNode := range g.GetAdjacencyList(source) {
		if g.DFS(adjacentNode, destination, asked) {
			return true
		}
	}
	return false
}
