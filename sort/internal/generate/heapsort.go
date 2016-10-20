//go:generate go run generate.go heapsort.go
//go:generate goimports -w ../../heapsort/
//go:generate gofmt -w ../../heapsort/
package main

const PACKAGE = "heapsort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) []{{.Name}} {
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

    result := make([]{{.Name}}, len(in))

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
`
const TEMPLATE_NUMERIC_TEST_FUNC = `
func TestInts(t *testing.T) {
    input := testutil.InputInts()
    output := Ints(input)

    if !testutil.IsSortedInts(output) {
        t.Fail()
    }
}
`
