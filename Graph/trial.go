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

func (g *graphs) bfs(start int) {

	visited := make(map[int]bool)
	visited[start] = true
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Print(node, "")

		for _, neigbour := range g.adjList[node] {
			if !visited[neigbour] {
				queue = append(queue, neigbour)
				visited[neigbour] = true
			}
		}

	}

}

func QuickSorts(arr []int, low, high int) {

	for low < high {
		pivotIndex := part(arr, low, high)

		QuickSorts(arr, low, pivotIndex-1)
		QuickSorts(arr, pivotIndex+1, high)

	}

}

func part(arr []int, low, high int) int {

	pivotIndex := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivotIndex {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1

}

func mergeSorts(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSorts(arr[:mid])
	right := mergeSorts(arr[mid:])

	return merges(left, right)

}

func merges(left, right []int) []int {

	result := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			result = append(result, left[i:]...)
			i++
		} else {
			result = append(result, right[j:]...)
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}

func HeapSorts(arr []int) {

	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapifys(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapifys(arr, i, 0)
	}

}

func heapifys(arr []int, n, i int) {
	largest := i
	left := 2*n + 1
	right := 2*n + 2

	if n < left && arr[left] > arr[largest] {
		largest = left
	}
	if n < right && arr[right] > arr[largest] {
		largest = right
	}

	if i != largest {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifys(arr, n, largest)
	}

}
