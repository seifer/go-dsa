//go:generate go run generate.go selectionsort.go
//go:generate goimports -w ../../selectionsort/
//go:generate gofmt -w ../../selectionsort/
package main

const PACKAGE = "selectionsort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) {
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
