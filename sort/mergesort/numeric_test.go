// This file should be auto generated

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
