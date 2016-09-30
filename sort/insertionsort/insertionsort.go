package insertionsort

// Ints sorting slice of ints
func Ints(in []int) {
	var i, key int

	for j := 1; j < len(in); j++ {
		i = j - 1
		key = in[j]

		for i >= 0 && in[i] > key {
			in[i+1] = in[i]
			i = i - 1
		}

		in[i+1] = key
	}
}
