package insertionsort

// Ints sorting slice of ints
func Ints(in []int) {
	var j, key int

	for i := 1; i < len(in); i++ {
		key = in[i]

		for j = i - 1; j >= 0 && in[j] > key; j-- {
			in[j+1] = in[j]
		}

		in[j+1] = key
	}
}
