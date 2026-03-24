package main

import "fmt"

// Graph structure
type Graphbfs struct {
	adj map[int][]int
}

// Add edge (undirected)
func (g *Graphbfs) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

// BFS using queue
func (g *Graphbfs) BFS(start int) {
	visited := make(map[int]bool)

	// queue
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		// dequeue (first element)
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node, " ")

		// visit neighbors
		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

// func main() {
// 	g := Graphbfs{
// 		adj: make(map[int][]int),
// 	}

// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 4)

// 	fmt.Println("BFS Traversal:")
// 	g.BFS(1)
// }
