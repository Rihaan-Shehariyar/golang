package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

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
