package main

import "fmt"

type GraphDfs struct {
	adj map[int][]int
}

func (g *GraphDfs) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *GraphDfs) DfSStack(start int) {
	visited := make(map[int]bool)

	stack := []int{start}

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true
			fmt.Print(node, "")

			for i := len(g.adj[node]) - 1; i >= 0; i-- {
				neighbour := g.adj[node][i]
				if !visited[neighbour] {
					stack = append(stack, neighbour)
				}
			}
		}

	}

}

// func main() {

// 	g := GraphDfs{
// 		adj: make(map[int][]int),
// 	}

// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 4)

// 	g.DfSStack(1)

// }
