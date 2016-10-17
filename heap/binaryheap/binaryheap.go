package binaryheap

type BinaryHeapInts struct {
	l []int
}

func NewInts() BinaryHeapInts {
	return BinaryHeapInts{}
}

func (h *BinaryHeapInts) Insert(key int) int {
	h.l = append(h.l, key)

	if len(h.l) == 1 {
		return 0
	}

	i := len(h.l) - 1

	for j := (i - 1) / 2; j >= 0 && h.l[j] > h.l[i]; j = (i - 1) / 2 {
		h.l[j], h.l[i] = h.l[i], h.l[j]
		i = j
	}

	return i
}

func (h *BinaryHeapInts) GetMin() (int, bool) {
	if len(h.l) == 0 {
		return 0, false
	}

	return h.l[0], true
}

func (h *BinaryHeapInts) ExtractMin() (int, bool) {
	ll := len(h.l)
	li := ll - 1

	if ll == 0 {
		return 0, false
	}

	if ll == 1 {
		return h.l[0], true
	}

	min := h.l[0]

	h.l[0] = h.l[li]
	h.l = h.l[:li]
	ll = len(h.l)
	li = ll - 1

	if ll == 2 {
		if h.l[0] > h.l[1] {
			h.l[1], h.l[0] = h.l[0], h.l[1]
		}

		return min, true
	}

	for i, left, right := 0, 1, 2; i < ll && left <= li; left, right = (2*i + 1), (2*i + 2) {
		if right > li {
			right = li
		}

		if h.l[right] < h.l[left] && h.l[i] > h.l[right] {
			h.l[right], h.l[i] = h.l[i], h.l[right]
			i = right
		} else if h.l[i] > h.l[left] {
			h.l[left], h.l[i] = h.l[i], h.l[left]
			i = left
		} else {
			break
		}
	}

	return min, true
}
