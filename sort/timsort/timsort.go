// This file should be auto generated

package timsort

import "github.com/seifer/go-dsa/sort/insertionsort"

// Ints sorting slice of int
func Ints(in []int) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Ints(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []int
	var i1, i2, ci, l1, l2 int

	tmprun := make([]int, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Int8s sorting slice of int8
func Int8s(in []int8) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Int8s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []int8
	var i1, i2, ci, l1, l2 int

	tmprun := make([]int8, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Int16s sorting slice of int16
func Int16s(in []int16) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Int16s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []int16
	var i1, i2, ci, l1, l2 int

	tmprun := make([]int16, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Int32s sorting slice of int32
func Int32s(in []int32) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Int32s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []int32
	var i1, i2, ci, l1, l2 int

	tmprun := make([]int32, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Int64s sorting slice of int64
func Int64s(in []int64) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Int64s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []int64
	var i1, i2, ci, l1, l2 int

	tmprun := make([]int64, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Uints sorting slice of uint
func Uints(in []uint) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Uints(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []uint
	var i1, i2, ci, l1, l2 int

	tmprun := make([]uint, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Uint8s sorting slice of uint8
func Uint8s(in []uint8) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Uint8s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []uint8
	var i1, i2, ci, l1, l2 int

	tmprun := make([]uint8, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Uint16s sorting slice of uint16
func Uint16s(in []uint16) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Uint16s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []uint16
	var i1, i2, ci, l1, l2 int

	tmprun := make([]uint16, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Uint32s sorting slice of uint32
func Uint32s(in []uint32) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Uint32s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []uint32
	var i1, i2, ci, l1, l2 int

	tmprun := make([]uint32, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Uint64s sorting slice of uint64
func Uint64s(in []uint64) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Uint64s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []uint64
	var i1, i2, ci, l1, l2 int

	tmprun := make([]uint64, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Float32s sorting slice of float32
func Float32s(in []float32) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Float32s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []float32
	var i1, i2, ci, l1, l2 int

	tmprun := make([]float32, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Float64s sorting slice of float64
func Float64s(in []float64) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Float64s(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []float64
	var i1, i2, ci, l1, l2 int

	tmprun := make([]float64, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}

// Strings sorting slice of string
func Strings(in []string) {
	var currun, minrun,
		lbound, rbound int

	// Select min run
	length := len(in)

	for length >= 64 {
		minrun |= (length & 1)
		length >>= 1
	}

	minrun += length
	cutrun := make([]int, 0, 2*(len(in)/minrun+1))

	// Split by min run on runs
	for i := 0; i < len(in)-1; i++ {
		if currun >= minrun {
			cutrun = append(cutrun, i-currun)
			cutrun = append(cutrun, currun)
			currun = 0
			continue
		}

		if in[i] <= in[i+1] {
			currun++
			continue
		}

		lbound = i - currun
		rbound = i - currun + minrun

		if rbound > len(in) {
			rbound = len(in)
		}

		insertionsort.Strings(in[lbound:rbound])

		cutrun = append(cutrun, lbound)
		cutrun = append(cutrun, rbound-lbound)
		currun = 0

		i = rbound - 1
	}

	// Merge all of min runs
	var t, a, b []string
	var i1, i2, ci, l1, l2 int

	tmprun := make([]string, ((len(cutrun) * minrun) / 4))

	for i := 2; i <= len(cutrun)/2; i *= 2 {
		for j := 0; j < len(cutrun); j += i * 2 {
			i1, i2, ci = 0, 0, 0

			t = tmprun[:cutrun[j+1]]
			a = in[cutrun[j] : cutrun[j]+cutrun[j+1]]
			b = in[cutrun[j+i] : cutrun[j+i]+cutrun[j+i+1]]

			for i := 0; i < len(a); i++ {
				t[i] = a[i]
			}

			l1 = len(t)
			l2 = len(b)

			for i1 < l1 && i2 < l2 {
				if b[i2] < t[i1] {
					if ci < l1 {
						a[ci] = b[i2]
					} else {
						b[ci-l1] = b[i2]
					}

					ci++
					i2++
				} else {
					if ci < l1 {
						a[ci] = t[i1]
					} else {
						b[ci-l1] = t[i1]
					}

					ci++
					i1++
				}
			}

			for i1 < l1 {
				if ci < l1 {
					a[ci] = t[i1]
				} else {
					b[ci-l1] = t[i1]
				}

				ci++
				i1++
			}

			cutrun[j+1] += cutrun[j+i+1]
		}
	}
}
