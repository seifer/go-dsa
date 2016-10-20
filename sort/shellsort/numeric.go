// This file should be auto generated

package shellsort

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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key int

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key int8

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key int16

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key int32

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key int64

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key uint

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key uint8

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key uint16

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key uint32

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key uint64

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key float32

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key float64

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
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

	// Generate Tokuda Gaps
	p := float64(1)
	l := len(in) - 1

	gaps := make([]int, 0, 8)
	gaps = append(gaps, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		gaps = append(gaps, int(p+1))
	}

	l = len(gaps) - 1

	for i := 0; i <= l/2; i++ {
		gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
	}

	var j int
	var key string

	// Sorting
	for _, gap := range gaps {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
}
