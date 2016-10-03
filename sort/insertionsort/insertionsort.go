package insertionsort

// Ints sorting slice of ints
func Ints(in []int) {
	var i, key int

	for j := 1; j < len(in); j++ {
		key = in[j]

		for i = j - 1; i >= 0 && in[i] > key; i-- {
			in[i+1] = in[i]
		}

		in[i+1] = key
	}
}
