//go:generate go run generate.go shellsort.go
//go:generate goimports -w ../../shellsort/
//go:generate gofmt -w ../../shellsort/
package main

const PACKAGE = "shellsort"
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

    // Generate Tokuda Gaps
    p := float64(1)
    l := len(in) - 1

    gaps := make([]int, 0, 8)
    gaps = append(gaps, 1)

    for {
        p = 2.25*p + 1

        if int(p+1) > l {
            break
        }

        gaps = append(gaps, int(p+1))
    }

    l = len(gaps) - 1

    for i := 0; i <= l/2; i++ {
        gaps[i], gaps[l-i] = gaps[l-i], gaps[i]
    }

    var j int
    var key {{.Name}}

    // Sorting
    for _, gap := range gaps {
        for i := gap; i < len(in); i++ {
            key = in[i]

            for j = i - gap; j >= 0 && in[j] > key; j -= gap {
                in[j+gap] = in[j]
            }

            in[j+gap] = key
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
