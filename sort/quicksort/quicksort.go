package quicksort

import (
	"math/rand"
	"time"
)

var rnd = rand.New(
	rand.NewSource(
		time.Now().UnixNano(),
	),
)

// Ints sorting slice of ints
func Ints(in []int) {
	if len(in) <= 1 {
		return
	}

	if len(in) == 2 {
		if in[1] < in[0] {
			in[0], in[1] = in[1], in[0]
		}

		return
	}

	pivotI := rnd.Intn(len(in))
	pivotV := in[pivotI]

	i := 0
	j := len(in) - 1

	for i < j {
		if in[i] >= pivotV && in[j] <= pivotV {
			in[i], in[j] = in[j], in[i]

			if i == pivotI {
				pivotI = j
			} else if j == pivotI {
				pivotI = i
			}
		}

		if in[i] < pivotV {
			if i < pivotI {
				i++
			}
		}

		if in[j] >= pivotV {
			if j > pivotI {
				j--
			}
		}
	}

	Ints(in[:pivotI])
	Ints(in[pivotI+1:])
}
