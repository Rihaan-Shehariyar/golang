package main

import "fmt"

// Graph structure
type Graph struct {
	adjList map[int][]int
}

// Initialize graph
func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

// Add edge (undirected)
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

// Print graph
func (g *Graph) Print() {
	for node, neighbors := range g.adjList {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}

func main() {
	graph := NewGraph()

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)

	graph.Print()
}
