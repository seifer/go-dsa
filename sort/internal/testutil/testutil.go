package testutil

import (
	"reflect"
	"sort"
)

var ints1k [1000]int
var ints1ksorted [1000]int

func init() {
	for i := 0; i < 1000; i++ {
		ints1k[i] = i ^ 0x2cc
		ints1ksorted[i] = ints1k[i]
	}

	sort.Ints(ints1ksorted[:])
}

func InputInts() []int {
	return []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
}

func InputInts1k() []int {
	return ints1k[:]
}

func IsSortedInts(ints []int) bool {
	return reflect.DeepEqual([]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}, ints)
}

func IsSortedInts1k(ints []int) bool {
	return reflect.DeepEqual(ints, ints1ksorted[:])
}
