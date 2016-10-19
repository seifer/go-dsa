package timsort

import (
	"reflect"
	"testing"

	"github.com/seifer/go-dsa/sort/internal/testutil"
)

func TestInts(t *testing.T) {
	input := testutil.InputInts()

	Ints(input)

	if !testutil.IsSortedInts(input) {
		t.Fail()
	}
}

func TestInts1k(t *testing.T) {
	input := testutil.InputInts1k()

	Ints(input)

	if !testutil.IsSortedInts1k(input) {
		t.Fail()
	}
}

func TestMerge1(t *testing.T) {
	a := []int{1, 5, 10, 3, 20}
	b := []int{1, 3, 5, 10, 20}

	merge(a[:3], a[3:], make([]int, 3))

	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func TestMerge2(t *testing.T) {
	a := []int{1, 5, 6, 7, 20}
	b := []int{1, 5, 6, 7, 20}

	merge(a[:3], a[3:], make([]int, 3))

	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}

func TestMerge3(t *testing.T) {
	a := []int{5, 7, 10, 12, 1, 2}
	b := []int{1, 2, 5, 7, 10, 12}

	merge(a[:4], a[4:], make([]int, 4))

	if !reflect.DeepEqual(a, b) {
		t.Fail()
	}
}
