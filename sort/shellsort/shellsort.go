package shellsort

// Ints sorting slice of ints
func Ints(in []int) {
	var j, key int

	for _, gap := range generateGapsTokuda(len(in) - 1) {
		for i := gap; i < len(in); i++ {
			key = in[i]

			for j = i - gap; j >= 0 && in[j] > key; j -= gap {
				in[j+gap] = in[j]
			}

			in[j+gap] = key
		}
	}
}

func generateGapsTokuda(l int) []int {
	if l < 2 {
		return nil
	}

	p := float64(1)
	g := make([]int, 0, 8)

	g = append(g, 1)

	for {
		p = 2.25*p + 1

		if int(p+1) > l {
			break
		}

		g = append(g, int(p+1))
	}

	l = len(g) - 1

	for i := 0; i <= l/2; i++ {
		g[i], g[l-i] = g[l-i], g[i]
	}

	return g
}
