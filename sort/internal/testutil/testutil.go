package testutil

import "reflect"

func InputInts() []int {
	return []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
}

func IsSortedInts(ints []int) bool {
	return reflect.DeepEqual([]int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}, ints)
}
