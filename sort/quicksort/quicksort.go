// This file should be auto generated

package quicksort

// Ints sorting slice of int
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

// Int8s sorting slice of int8
func Int8s(in []int8) {
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

	Int8s(in[:i])
	Int8s(in[i+1:])
}

// Int16s sorting slice of int16
func Int16s(in []int16) {
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

	Int16s(in[:i])
	Int16s(in[i+1:])
}

// Int32s sorting slice of int32
func Int32s(in []int32) {
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

	Int32s(in[:i])
	Int32s(in[i+1:])
}

// Int64s sorting slice of int64
func Int64s(in []int64) {
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

	Int64s(in[:i])
	Int64s(in[i+1:])
}

// Uints sorting slice of uint
func Uints(in []uint) {
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

	Uints(in[:i])
	Uints(in[i+1:])
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) {
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

	Uint8s(in[:i])
	Uint8s(in[i+1:])
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) {
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

	Uint16s(in[:i])
	Uint16s(in[i+1:])
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) {
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

	Uint32s(in[:i])
	Uint32s(in[i+1:])
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) {
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

	Uint64s(in[:i])
	Uint64s(in[i+1:])
}

// Float32s sorting slice of float32
func Float32s(in []float32) {
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

	Float32s(in[:i])
	Float32s(in[i+1:])
}

// Float64s sorting slice of float64
func Float64s(in []float64) {
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

	Float64s(in[:i])
	Float64s(in[i+1:])
}

// Strings sorting slice of string
func Strings(in []string) {
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

	Strings(in[:i])
	Strings(in[i+1:])
}
