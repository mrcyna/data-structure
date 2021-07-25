package structures

import "fmt"

// MaxHeap represents the Heap data structure
type MaxHeap struct {
	array []int
}

// NewMaxHeap creates a new MaxHeap
func NewMaxHeap() MaxHeap {
	return MaxHeap{}
}

// Insert inserts a new item in the MaxHeap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract extracts the root item from the MaxHeap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array len is zero")
		return -1
	}

	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func (h *MaxHeap) Items() []int {
	return h.array
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}
