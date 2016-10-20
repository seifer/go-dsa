//go:generate go run generate.go mergesort.go
//go:generate goimports -w ../../mergesort/
//go:generate gofmt -w ../../mergesort/
package main

const PACKAGE = "mergesort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) []{{.Name}} {
    if len(in) <= 1 {
        return in
    }

    if len(in) == 2 {
        if in[0] > in[1] {
            in[0], in[1] = in[1], in[0]
        }

        return in
    }

    h := (len(in)-1)/2 + 1
    p1 := {{.FuncName}}(in[:h])
    p2 := {{.FuncName}}(in[h:])
    i1 := 0
    i2 := 0
    l1 := len(p1) - 1
    l2 := len(p2) - 1
    out := make([]{{.Name}}, 0, len(in))

    for i1 <= l1 && i2 <= l2 {
        if p1[i1] <= p2[i2] {
            out = append(out, p1[i1])
            i1++
        } else {
            out = append(out, p2[i2])
            i2++
        }
    }

    if i1 <= l1 {
        out = append(out, p1[i1:]...)
    } else if i2 <= l2 {
        out = append(out, p2[i2:]...)
    }

    return out
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
