package main

import (
	"fmt"
)

type graphs struct {
	adjList map[int][]int
}

func newgraph() *graphs {
	return &graphs{adjList: map[int][]int{}}
}

func (g *graphs) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

func (g *graphs) Print() {
	for node, neighbour := range g.adjList {
		fmt.Printf("%d -> %v", node, neighbour)
	}
}

func (g *graphs) DfSStack(start int) {

	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}

		fmt.Print(node, "")
		visited[node] = true

		for i := len(g.adjList) - 1; i >= 0; i-- {

			neigbour := g.adjList[node][i]
			if !visited[neigbour] {
				stack = append(stack, neigbour)
			}
		}

	}

}

func (g *graphs) dfsrec(node int, visited map[int]bool) {

	visited[node] = true
	fmt.Print(node, "")

	for _, neigbour := range g.adjList[node] {
		if !visited[neigbour] {

			g.dfsrec(neigbour, visited)
		}
	}

}

