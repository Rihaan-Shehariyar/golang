package main

import "fmt"

type graph struct {
	adj map[int][]int
}

func newGraph() *graph {
	return &graph{adj: map[int][]int{}}
}
func (g *graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *graph) Print() {
	for node, neighbours := range g.adj {
		fmt.Printf("%d->%v", node, neighbours)
	}
}

type graphM struct {
	matrix [][]int
	nodes  int
}

func Newmatrix(n int) *graphM {
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return &graphM{
		matrix: matrix,
		nodes:  n,
	}

}

func (g *graphM) AddEdge(u, v int) {
	g.matrix[u][v] = 1
	g.matrix[v][u] = 1
}

func (g *graphM) Printm() {

	for i := 0; i < g.nodes; i++ {
		fmt.Print(g.matrix[i])
	}

}

func (g *graph) dfsstack(start int) {

	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := len(g.adj[node]); i >= 0; i++ {
			neigbour := g.adj[node][i]
			if !visited[neigbour] {
				stack = append(stack, neigbour)
			}
		}

	}

}

func (g *graph) dfsrecursion(node int, visited map[int]bool) {

	visited[node] = true
	fmt.Print(node, "")

	for _, neigbour := range g.adj[node] {
		if !visited[node] {
			g.dfsrecursion(neigbour, visited)
		}
	}

}

func (g *graph) dfs(start int) {
	visited := make(map[int]bool)
	g.dfsrecursion(start, visited)
}
func (g *graph) bfsqueue(start int) {

	visited := make(map[int]bool)
	visited[start] = true

	queue := []int{start}

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Println(node, "")
		for _, neiigbour := range g.adj[node] {
			if !visited[neiigbour] {
				visited[neiigbour] = true
				queue = append(queue, neiigbour)
			}
		}

	}

}
