package heapsort

// Ints sorting slice of ints
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
