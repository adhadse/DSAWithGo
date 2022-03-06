package graph

import (
	"fmt"
	"testing"
)

func TestAdjacencyMatrix(t *testing.T) {
	graph := MakeAdjacencyMatrix(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.PrintMatrix()
}

func TestGraph(t *testing.T) {
	graph := MakeGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	graph.PrintAdjacencyList()
}

// Should Print:
// Adjacency List of vertex 1| head -> 4 -> 3 -> 2 -> 0
func TestGraph_GetAdjacencyList(t *testing.T) {
	graph := MakeGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	adjacencyList := graph.GetAdjacencyList(1)

	fmt.Printf("Adjacency List of vertex %d| head", 1)
	for _, node := range adjacencyList {
		fmt.Printf(" -> %d", node.vertex)
	}
	println()
}

func TestGraph_BFS(t *testing.T) {
	graph := MakeGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	if found, err := graph.BFS(0, 3); found != true && err == nil {
		t.Errorf("BFS returned false for source:0 destination: 3")
	} else {
		// err != nil; some error deal with it
		return
	}
	if found, err := graph.BFS(0, 2); found != true && err == nil {
		t.Errorf("BFS returned false for source:0 destination: 2")
	} else {
		// err != nil; some error deal with it
		return
	}
	if found, err := graph.BFS(0, 5); found != false && err == nil {
		t.Errorf("BFS returned true for source:0 destination: 3")
	} else {
		// err != nil; some error deal with it
		return
	}
}

func TestGraph_DFS(t *testing.T) {
	graph := MakeGraph(5)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)

	if val := graph.DFS(1, 3, []int{}); val != true {
		t.Error("DFS returned false for source: 0, destination: 3")
	}
}

func TestWeightedGraph_KruskalMinimumSpanningTree(t *testing.T) {
	graph := MakeWeightedGraph(4)
	graph.AddEdge(0, 1, 10)
	graph.AddEdge(0, 2, 6)
	graph.AddEdge(0, 3, 5)
	graph.AddEdge(1, 3, 15)
	graph.AddEdge(2, 3, 4)

	graph.KruskalMinimumSpanningTree()
}

func TestDiWeightedGraph_BellmanFord(t *testing.T) {
	g := MakeDiWeightedGraph(5)

	g.AddEdge(0, 1, -1)
	g.AddEdge(0, 2, 4)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 3, 2)
	g.AddEdge(1, 4, 2)
	g.AddEdge(3, 2, 5)
	g.AddEdge(3, 1, 1)
	g.AddEdge(4, 3, -3)

	g.BellmanFord(0)
}

func TestWeightedGraph_Dijkstra(t *testing.T) {
	g := MakeWeightedGraph(7)

	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 2, 3)

	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 3)
	g.AddEdge(1, 4, 6)
	g.AddEdge(1, 6, 5)

	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 5, 4)

	g.AddEdge(3, 4, 2)
	g.AddEdge(5, 4, 1)
	g.AddEdge(6, 4, 2)

	g.DijkstraShortestPath(0, 4)
}

// https://youtu.be/jsmMtJpPnhU?t=117
func TestWeightedGraph_PrimMinimumSpanningTree(t *testing.T) {
	g := MakeWeightedGraph(6)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 2)
	g.AddEdge(0, 2, 5)
	g.AddEdge(0, 4, 3)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 6)
	g.AddEdge(3, 4, 4)
	g.AddEdge(3, 5, 8)

	mst, mstCost, err := g.PrimMinimumSpanningTree(0)
	if err != nil {
		print(err.Error())
		return
	}
	mst.PrintAdjacencyList()
	fmt.Printf("MST total cost: %d\n", mstCost)
}
