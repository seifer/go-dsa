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

	pivot := in[rnd.Intn(len(in))]
	// TODO need to make test with corner cases.
	// In common case, median work more better than rand.
	// pivot := (in[0] + in[len(in)/2] + in[len(in)-1]) / 3

	i := 0
	j := len(in) - 1

	for i < j {
		for in[i] < pivot {
			i++
		}

		for in[j] > pivot {
			j--
		}

		in[i], in[j] = in[j], in[i]

		if in[i] == pivot {
			j--
		} else if in[j] == pivot {
			i++
		}
	}

	Ints(in[:i])
	Ints(in[i+1:])
}
