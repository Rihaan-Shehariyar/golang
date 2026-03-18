package main

import (
	"fmt"
)

type Heap struct {
	data []int

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

func (h *Heap) ExtractMIn() int {

	if len(h.data) == 0 {
		return -1
	}

	min := h.data[0]

	last := h.data[len(h.data)-1]
	h.data[0] = last
	h.data = h.data[:len(h.data)-1]

	return min
}

func (h *Heap) HeapifyDown(index int) {

	size := len(h.data)

	for {

		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < size && h.data[left] < h.data[smallest] {
			smallest = left
		}

		if right < size && h.data[right] < h.data[smallest] {
			smallest = right
		}

		if index == smallest {
			break
		}

		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest

	}

}

func main() {

	h := Heap{}
	x := Heap{}
	h.MinInsert(10)
	h.MinInsert(5)
	h.MinInsert(30)
	h.MinInsert(2)
	h.MinInsert(1)
	x.MaxInsert(10)
	x.MaxInsert(20)
	x.MaxInsert(50)

	fmt.Println(h.data)
	fmt.Println(x.data)
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
