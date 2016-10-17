package binaryheap

import "testing"

func TestInsert(t *testing.T) {
	h := NewInts()

	if h.Insert(2) != 0 {
		t.Fail()
	}

	if h.Insert(1) != 0 {
		t.Fail()
	}

	if h.Insert(3) != 2 {
		t.Fail()
	}
}

func TestGetMinEmpty(t *testing.T) {
	h := NewInts()

	val, ok := h.GetMin()

	if ok {
		t.Fail()
	}

	if val != 0 {
		t.Fail()
	}
}

func TestGetMinNotEmpty(t *testing.T) {
	h := NewInts()

	for _, key := range []int{4, 2, 1, 23, 4, 5, -92, 0} {
		h.Insert(key)
	}

	val, ok := h.GetMin()

	if !ok {
		t.Fail()
	}

	if val != -92 {
		t.Fail()
	}

	val, ok = h.GetMin()

	if !ok {
		t.Fail()
	}

	if val != -92 {
		t.Fail()
	}
}

func TestExtractMinEmpty(t *testing.T) {
	h := NewInts()

	val, ok := h.ExtractMin()

	if ok {
		t.Fail()
	}

	if val != 0 {
		t.Fail()
	}
}

func TestExtractMinNotEmpty(t *testing.T) {
	h := NewInts()

	for _, key := range []int{4, 2, 1, 23, 4, 5, -92, 0} {
		h.Insert(key)
	}

	val, ok := h.ExtractMin()

	if !ok {
		t.Fail()
	}

	if val != -92 {
		t.Fail()
	}

	val, ok = h.ExtractMin()

	if !ok {
		t.Fail()
	}

	if val != 0 {
		t.Fail()
	}
}

func TestExtractMinCorretMinSequence1(t *testing.T) {
	in := []int{4, 2, 1, 23, 4, 5, -92, 0}
	mins := []int{-92, 0, 1, 2, 4, 4, 5, 23}

	h := NewInts()

	for _, key := range in {
		h.Insert(key)
	}

	for _, min := range mins {
		val, ok := h.ExtractMin()

		if !ok {
			t.Fail()
		}

		if val != min {
			t.Fail()
		}
	}
}

func TestExtractMinCorretMinSequence2(t *testing.T) {
	in := []int{4, 5, 23}
	mins := []int{4, 5, 23}

	h := NewInts()

	for _, key := range in {
		h.Insert(key)
	}

	for _, min := range mins {
		val, ok := h.ExtractMin()

		if !ok {
			t.Fail()
		}

		if val != min {
			t.Fail()
		}
	}
}
