//go:generate go run generate.go quicksort.go
//go:generate goimports -w ../../quicksort/
//go:generate gofmt -w ../../quicksort/
package main

const PACKAGE = "quicksort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) {
    if len(in) <= 1 {
        return
    }

    if len(in) == 2 {
        if in[1] < in[0] {
            in[0], in[1] = in[1], in[0]
        }

        return
    }

    pivot := in[rnd.Intn(len(in))]
    // TODO need to make test with corner cases.
    // In common case, median work more better than rand.
    // pivot := (in[0] + in[len(in)/2] + in[len(in)-1]) / 3

    i := 0
    j := len(in) - 1

    for i < j {
        for in[i] < pivot {
            i++
        }

        for in[j] > pivot {
            j--
        }

        in[i], in[j] = in[j], in[i]

        if in[i] == pivot {
            j--
        } else if in[j] == pivot {
            i++
        }
    }

    {{.FuncName}}(in[:i])
    {{.FuncName}}(in[i+1:])
}
`
const TEMPLATE_NUMERIC_TEST_FUNC = `
func TestInts(t *testing.T) {
    input := testutil.InputInts()

    Ints(input)

    if !testutil.IsSortedInts(input) {
        t.Fail()
    }
}
`
