package main

import "fmt"

// import "fmt"

// type Graph struct {
// 	adjList map[int][]int
// }

// func NewGraph() *Graph {
// 	return &Graph{adjList: make(map[int][]int)}
// }

// func (g *Graph) AddEdge(u, v int) {
// 	g.adjList[u] = append(g.adjList[u], v)
// 	g.adjList[v] = append(g.adjList[v], u)
// }

// func (g *Graph) AddDirectEdge(u, v int) {
// 	g.adjList[u] = append(g.adjList[u], v)
// }

// func (g *Graph) PrintGraph() {
// 	for node, neighbours := range g.adjList {
// 		fmt.Printf("%d->%v\n", node, neighbours)
// 	}
// }

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

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) Print() {
	for node, neigbours := range g.adj {
		fmt.Printf("%d -> %v", node, neigbours)
	}
}

type graphMatrix struct {
	matrix [][]int
	nodes  int
}

func newGraphMatrix(n int) *graphMatrix {

	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return &graphMatrix{
		matrix: matrix,
		nodes:  n,
	}

}

func (g *graphMatrix) AddEdge(u, v int) {
	g.matrix[u][v] = 1
}

func (g *graphMatrix) Print() {
	for i := 0; i < g.nodes; i++ {
		fmt.Println(g.matrix[i])
	}
}

func (g *Graph) DfSStack(start int) {

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

func (g *Graph) dfsrec(node int, visited map[int]bool) {

	visited[node] = true
	fmt.Print(node, "")

	for _, neigbour := range g.adj[node] {
		if !visited[neigbour] {
			g.dfsrec(neigbour, visited)
		}
	}

}

func (g *Graph) dfs(start int) {
	visited := make(map[int]bool)
	g.dfsrec(start, visited)
}

func (g *Graph) bfs(start int) {

	visited := make(map[int]bool)

	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		fmt.Print(node, "")

		for _, neigbour := range g.adj[node] {
			if !visited[neigbour] {
				visited[neigbour] = true
				queue = append(queue, neigbour)
			}
		}

	}

}

// func main() {
// 	g := NewGraph()

// 	// g := newGraphMatrix(3)

// 	g.AddEdge(1, 2)
// 	g.AddEdge(1, 3)
// 	g.AddEdge(2, 4)

// 	g.dfs(1)

// 	// g.Print()

// }

func Partitions(arr []int, low, high int)int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] > pivot {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

 arr[i+1],arr[high] = arr[high],arr[i+1]
 
return i+1

}
