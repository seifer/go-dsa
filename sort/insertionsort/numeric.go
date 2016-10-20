// This file should be auto generated

package insertionsort

// Ints sorting slice of int
func Ints(in []int) {
	var i, j int
	var key int

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Int8s sorting slice of int8
func Int8s(in []int8) {
	var i, j int
	var key int8

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Int16s sorting slice of int16
func Int16s(in []int16) {
	var i, j int
	var key int16

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Int32s sorting slice of int32
func Int32s(in []int32) {
	var i, j int
	var key int32

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Int64s sorting slice of int64
func Int64s(in []int64) {
	var i, j int
	var key int64

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Uints sorting slice of uint
func Uints(in []uint) {
	var i, j int
	var key uint

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) {
	var i, j int
	var key uint8

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) {
	var i, j int
	var key uint16

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) {
	var i, j int
	var key uint32

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) {
	var i, j int
	var key uint64

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Float32s sorting slice of float32
func Float32s(in []float32) {
	var i, j int
	var key float32

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Float64s sorting slice of float64
func Float64s(in []float64) {
	var i, j int
	var key float64

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}

// Strings sorting slice of string
func Strings(in []string) {
	var i, j int
	var key string

	for i = 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}
