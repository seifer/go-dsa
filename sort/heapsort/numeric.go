// This file should be auto generated

package heapsort

// Ints sorting slice of int
func Ints(in []int) []int {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]int, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Int8s sorting slice of int8
func Int8s(in []int8) []int8 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]int8, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Int16s sorting slice of int16
func Int16s(in []int16) []int16 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]int16, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Int32s sorting slice of int32
func Int32s(in []int32) []int32 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]int32, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Int64s sorting slice of int64
func Int64s(in []int64) []int64 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]int64, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Uints sorting slice of uint
func Uints(in []uint) []uint {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]uint, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) []uint8 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]uint8, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) []uint16 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]uint16, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) []uint32 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]uint32, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) []uint64 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]uint64, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Float32s sorting slice of float32
func Float32s(in []float32) []float32 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]float32, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Float64s sorting slice of float64
func Float64s(in []float64) []float64 {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]float64, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}

// Strings sorting slice of string
func Strings(in []string) []string {
	var left, right int

	// Make heap
	for i := len(in) / 2; i >= 0; i-- {
		for j := i; j < len(in)/2; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in) && left < len(in) {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in) && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	result := make([]string, len(in))

	for i := range in {
		result[i] = in[0]
		in[0] = in[len(in)-i-1]

		for j := 0; j < len(in)-i; {
			left = 2*j + 1
			right = 2*j + 2

			if right < len(in)-i && left < len(in)-i {
				if in[left] < in[right] && in[j] > in[left] {
					in[left], in[j] = in[j], in[left]
					j = left
				} else if in[j] > in[right] {
					in[right], in[j] = in[j], in[right]
					j = right
				} else {
					break
				}
			} else if left < len(in)-i && in[j] > in[left] {
				in[left], in[j] = in[j], in[left]
				j = left
			} else {
				break
			}
		}
	}

	return result
}
