package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}

	extracted := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftIndex := left(index)
	rightIndex := right(index)
	childToCompare := 0

	for leftIndex <= lastIndex {
		if lastIndex == leftIndex {
			childToCompare = leftIndex // when left child is the only child
		} else if h.array[leftIndex] > h.array[rightIndex] {
			childToCompare = leftIndex // when left child is larger
		} else {
			childToCompare = rightIndex // when right child is larger
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftIndex = left(index)
			rightIndex = right(index)
		} else {
			return
		}
	}
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHead := []int{10, 20, 30, 3, 7, 2, 11, 13, 17}
	for _, v := range buildHead {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
