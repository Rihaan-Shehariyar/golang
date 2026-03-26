package main

import "fmt"

type GraphBfs struct {
	adj map[int][]int
}

func (g *GraphBfs) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *GraphBfs) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Println(node, "")

		for _, neigbours := range g.adj[node] {
			if !visited[neigbours] {
				visited[neigbours] = true
				queue = append(queue, neigbours)
			}
		}

	}

}

// func main(){

//  g := GraphBfs{
//  adj: map[int][]int{},
// }

// g.AddEdge(1,3)
// g.AddEdge(1,2)
// g.AddEdge(2,4)

// g.BFS(1)

// }
