package main

import "fmt"

type GraphMatrix struct {
	matrix [][]int
	nodes  int
}

func NewGraphMatrix(n int) *GraphMatrix {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return &GraphMatrix{
		matrix: matrix,
		nodes:  n,
	}

}

func (g *GraphMatrix) AddEdge(u, v int) {
	g.matrix[u][v] = 1
	// g.matrix[v][u] = 1
}

func (g *GraphMatrix) Print() {
	fmt.Println("Adjacency Matrix:")
	for i := 0; i < g.nodes; i++ {
		fmt.Println(g.matrix[i])
	}
}
func main() {
	g := NewGraphMatrix(4)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)

	g.Print()
}
