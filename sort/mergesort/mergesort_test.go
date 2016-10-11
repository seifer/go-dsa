package mergesort

import (
    "testing"

    "github.com/seifer/go-dsa/sort/internal/testutil"
)

func TestInts(t *testing.T) {
    input := testutil.InputInts()
    output := Ints(input)

    if !testutil.IsSortedInts(output) {
        t.Fail()
    }
}

func BenchmarkInts(b *testing.B) {
    b.StopTimer()
    in := make([]int, 1<<10)
    for i := 0; i < len(in); i++ {
        in[i] = i ^ 0x2cc
    }
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        _ = Ints(in)
    }
}
