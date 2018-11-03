package datastructures

type MinHeap []int

func (h MinHeap) Add(v int) MinHeap {
	h = append(h, v)
	h.bubbleUp()
	return h
}

func (h MinHeap) Delete(v int) MinHeap {
	if len(h) == 0 {
		return h[0:]
	}

	index := h.findIndex(v)
	if index < 0 {
		return h
	}

	h[index] = h[len(h)-1]
	h = h[:len(h)-1]
	h.bubbleDown(index)
	return h
}

func (h MinHeap) Pop() MinHeap {
	if len(h) == 0 {
		return h[0:]
	}

	h[0] = h[len(h)-1]
	h = h[:len(h)-1]
	h.bubbleDown(0)
	return h
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h MinHeap) findIndex(v int) int {
	for i := 0; i < len(h); i++ {
		if h[i] == v {
			return i
		}
	}
	return -1
}

func (h MinHeap) bubbleUp() {
	current := len(h) - 1
	for h.hasParent(current) && !h.isValid(current) {
		parentIdx := h.parent(current)
		h.swap(parentIdx, current)
		current = parentIdx
	}
}

func (h MinHeap) bubbleDown(current int) {
	for h.hasLeft(current) && !h.isValid(current) {
		smallestChild := h.left(current)
		if h.hasRight(current) && h[h.right(current)] < h[smallestChild] {
			smallestChild = h.right(current)
		}
		h.swap(current, smallestChild)
		current = smallestChild
	}
}

func (h MinHeap) hasParent(i int) bool {
	return h.parent(i) >= 0
}

func (h MinHeap) hasRight(i int) bool {
	return h.right(i) < len(h)
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) hasLeft(i int) bool {
	return h.left(i) < len(h)
}

func (h MinHeap) parent(i int) int {
	return (i + 1) / 2
}

func (h MinHeap) left(i int) int {
	return (i+1)*2 - 1
}

func (h MinHeap) right(i int) int {
	return (i+1)*2 + 1 - 1
}

func (h MinHeap) isValid(i int) bool {
	if h.hasParent(i) && h[h.parent(i)] >= h[i] {
		return false
	}

	if h.hasLeft(i) && h[h.left(i)] <= h[i] {
		return false
	}

	if h.hasRight(i) && h[h.right(i)] <= h[i] {
		return false
	}

	return true
}
