package main

import "fmt"

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return Merge(left, right)

}

// graph tress dif
func Merge(left, right []int) []int {

	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func main() {

	// arr := []int{2, 56, 7, 8, 9}

	// sortded := MergeSort(arr)

	// fmt.Print(sortded)

	g := NewGraph()

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	g.DFS(1)

}

type graph struct {
	adj map[int][]int
}

func NewGraph() *graph {
	return &graph{adj: make(map[int][]int)}
}

func (g *graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

func (g *graph) DFS(start int) {

	visited := make(map[int]bool)
	stack := []int{start}

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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
