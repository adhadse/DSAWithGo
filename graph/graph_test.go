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

func TestAdjacencyList(t *testing.T) {
	graph := MakeAdjacencyList(5)
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
func TestAdjacencyList_GetAdjacencyList(t *testing.T) {
	graph := MakeAdjacencyList(5)
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

func TestAdjacencyList_BFS(t *testing.T) {
	graph := MakeAdjacencyList(5)
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
