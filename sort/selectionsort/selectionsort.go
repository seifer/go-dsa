package selectionsort

// Ints sorting slice of ints
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
