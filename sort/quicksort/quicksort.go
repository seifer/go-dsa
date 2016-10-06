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

	pivotV := in[rnd.Intn(len(in))]
	// TODO need to make test with corner cases.
	// In common case, median work better than rand.
	// pivotV := (in[0] + in[len(in)/2] + in[len(in)-1]) / 3

	i := 0
	j := len(in) - 1

	for i < j {
		for in[i] < pivotV {
			i++
		}

		for in[j] > pivotV {
			j--
		}

		in[i], in[j] = in[j], in[i]

		if in[i] == pivotV {
			j--
		} else if in[j] == pivotV {
			i++
		}
	}

	Ints(in[:i])
	Ints(in[i+1:])
}
