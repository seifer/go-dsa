//go:generate go run generate.go insertionsort.go
//go:generate goimports -w ../../insertionsort/
//go:generate gofmt -w ../../insertionsort/
package main

const PACKAGE = "insertionsort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) {
    var i, j int
    var key {{.Name}}

    for i = 1; i < len(in); i++ {
        key = in[i]

        for j = i - 1; j >= 0 && in[j] > key; j-- {
            in[j+1] = in[j]
        }

        in[j+1] = key
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
