package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) AddDirectEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *Graph) PrintGraph() {
	for node, neighbours := range g.adjList {
		fmt.Printf("%d->%v\n", node, neighbours)
	}
}

// func main() {
// 	g := NewGraph()

// 	// g.AddEdge(1, 2)
// 	// g.AddEdge(1, 3)
// 	// g.AddEdge(2, 4)

//  g.AddDirectEdge(1,2)
//  g.AddDirectEdge(1,3)
//  g.AddDirectEdge(1,3)
//  g.AddDirectEdge(1,3)

//  g.PrintGraph()

// }
