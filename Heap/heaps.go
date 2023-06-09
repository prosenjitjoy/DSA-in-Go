package main

// import "fmt"

// // MaxHeap struct has a slice that holds the array
// type MaxHeap struct {
// 	array []int
// }

// // Insert adds an element to the heap

// func (h *MaxHeap) Insert(key int) {
// 	h.array = append(h.array, key)
// 	h.maxHeapifyUp(len(h.array) - 1)
// }

// // Extract returns the largest key, and removes it from the heap

// func (h *MaxHeap) Extract() int {
// 	// when the array is empty
// 	if len(h.array) == 0 {
// 		fmt.Println("cannot extract because array length is 0")
// 		return -1
// 	}

// 	extracted := h.array[0]
// 	lastIndex := len(h.array) - 1
// 	h.array[0] = h.array[lastIndex]
// 	h.array = h.array[:lastIndex]
// 	h.maxHeapifyDown(0)
// 	return extracted
// }

// // maxHeapifyDown will heapify top to bottom
// func (h *MaxHeap) maxHeapifyDown(index int) {
// 	lastIndex := len(h.array) - 1
// 	leftIndex := left(index)
// 	rightIndex := right(index)
// 	childToCompare := 0

// 	// loop while index has at least one child
// 	for leftIndex <= lastIndex {
// 		if leftIndex == lastIndex { // when left child is the only child
// 			childToCompare = leftIndex
// 		} else if h.array[leftIndex] > h.array[rightIndex] { // when left child is larger
// 			childToCompare = leftIndex
// 		} else { // when right child is larger
// 			childToCompare = rightIndex
// 		}

// 		// compare array value of current index to larger child and swap if smaller
// 		if h.array[index] < h.array[childToCompare] {
// 			h.swap(index, childToCompare)
// 			index = childToCompare
// 			leftIndex = left(index)
// 			rightIndex = right(index)
// 		} else {
// 			return
// 		}
// 	}
// }

// // maxHeapifyUp will heapify from bottom top
// func (h *MaxHeap) maxHeapifyUp(index int) {
// 	for h.array[parent(index)] < h.array[index] {
// 		h.swap(parent(index), index)
// 		index = parent(index)
// 	}
// }

// // get the parent index
// func parent(i int) int {
// 	return (i - 1) / 2
// }

// // get the left child index
// func left(i int) int {
// 	return 2*i + 1
// }

// // get the right child index
// func right(i int) int {
// 	return 2*i + 2
// }

// // swap keys in the array
// func (h *MaxHeap) swap(i1, i2 int) {
// 	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
// }

// func main() {
// 	m := &MaxHeap{}
// 	fmt.Println(m)
// 	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

// 	for _, v := range buildHeap {
// 		m.Insert(v)
// 		fmt.Println(m)
// 	}

// 	for i := 0; i < 5; i++ {
// 		m.Extract()
// 		fmt.Println(m)
// 	}
// }
