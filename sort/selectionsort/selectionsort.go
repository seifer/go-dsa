// This file should be auto generated

package selectionsort

// Ints sorting slice of int
func Ints(in []int) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Int8s sorting slice of int8
func Int8s(in []int8) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Int16s sorting slice of int16
func Int16s(in []int16) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Int32s sorting slice of int32
func Int32s(in []int32) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Int64s sorting slice of int64
func Int64s(in []int64) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Uints sorting slice of uint
func Uints(in []uint) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Float32s sorting slice of float32
func Float32s(in []float32) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Float64s sorting slice of float64
func Float64s(in []float64) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}

// Strings sorting slice of string
func Strings(in []string) {
	var mi int

	for i := 0; i < len(in); i++ {
		mi = i

		for j := i; j < len(in); j++ {
			if in[j] < in[mi] {
				mi = j
			}
		}

		if i != mi {
			in[i], in[mi] = in[mi], in[i]
		}
	}
}
