// This file should be auto generated

package mergesort

// Ints sorting slice of int
func Ints(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Ints(in[:h])
	p2 := Ints(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]int, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Int8s sorting slice of int8
func Int8s(in []int8) []int8 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Int8s(in[:h])
	p2 := Int8s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]int8, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Int16s sorting slice of int16
func Int16s(in []int16) []int16 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Int16s(in[:h])
	p2 := Int16s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]int16, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Int32s sorting slice of int32
func Int32s(in []int32) []int32 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Int32s(in[:h])
	p2 := Int32s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]int32, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Int64s sorting slice of int64
func Int64s(in []int64) []int64 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Int64s(in[:h])
	p2 := Int64s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]int64, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Uints sorting slice of uint
func Uints(in []uint) []uint {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Uints(in[:h])
	p2 := Uints(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]uint, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) []uint8 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Uint8s(in[:h])
	p2 := Uint8s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]uint8, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) []uint16 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Uint16s(in[:h])
	p2 := Uint16s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]uint16, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) []uint32 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Uint32s(in[:h])
	p2 := Uint32s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]uint32, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) []uint64 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Uint64s(in[:h])
	p2 := Uint64s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]uint64, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Float32s sorting slice of float32
func Float32s(in []float32) []float32 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Float32s(in[:h])
	p2 := Float32s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]float32, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Float64s sorting slice of float64
func Float64s(in []float64) []float64 {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Float64s(in[:h])
	p2 := Float64s(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]float64, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}

// Strings sorting slice of string
func Strings(in []string) []string {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[1] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	h := (len(in)-1)/2 + 1
	p1 := Strings(in[:h])
	p2 := Strings(in[h:])
	i1 := 0
	i2 := 0
	l1 := len(p1) - 1
	l2 := len(p2) - 1
	out := make([]string, 0, len(in))

	for i1 <= l1 && i2 <= l2 {
		if p1[i1] <= p2[i2] {
			out = append(out, p1[i1])
			i1++
		} else {
			out = append(out, p2[i2])
			i2++
		}
	}

	if i1 <= l1 {
		out = append(out, p1[i1:]...)
	} else if i2 <= l2 {
		out = append(out, p2[i2:]...)
	}

	return out
}
