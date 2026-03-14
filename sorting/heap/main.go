package main

import "fmt"

type Heap struct {
	data []int
}

func (h *Heap) MinInsert(value int) {

	h.data = append(h.data, value)
	index := len(h.data) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if h.data[index] < h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			break
		}
	}

}

func (h *Heap) MaxInsert(value int) {

	h.data = append(h.data, value)
	index := len(h.data) - 1

	for index > 0 {
		parent := (index - 1) / 2

		if h.data[index] > h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
			index = parent
		} else {
			break

		}

	}

}

func main() {

	h := MinHeap{}
	h.MinInsert(10)
	h.MinInsert(5)
	h.MinInsert(30)
	h.MinInsert(2)
	h.MinInsert(1)

	fmt.Println(h.data)
}
