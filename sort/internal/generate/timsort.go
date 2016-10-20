//go:generate go run generate.go timsort.go
//go:generate goimports -w ../../timsort/
//go:generate gofmt -w ../../timsort/
package main

const PACKAGE = "timsort"
const TEMPLATE = `
// {{.FuncName}} sorting slice of {{.Name}}
func {{.FuncName}}(in []{{.Name}}) {
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

        insertionsort.{{.FuncName}}(in[lbound:rbound])

        cutrun = append(cutrun, lbound)
        cutrun = append(cutrun, rbound-lbound)
        currun = 0

        i = rbound - 1
    }

    // Merge all of min runs
    var t, a, b []{{.Name}}
    var i1, i2, ci, l1, l2 int

    tmprun := make([]{{.Name}}, ((len(cutrun) * minrun) / 4))

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
`
const TEMPLATE_NUMERIC_TEST_FUNC = `
func TestInts(t *testing.T) {
    input := testutil.InputInts()

    Ints(input)

    if !testutil.IsSortedInts(input) {
        t.Fail()
    }
}

func TestInts1k(t *testing.T) {
    input := testutil.InputInts1k()

    Ints(input)

    if !testutil.IsSortedInts1k(input) {
        t.Fail()
    }
}
`
