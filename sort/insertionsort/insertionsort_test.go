// This file should be auto generated

package insertionsort

import (
	"testing"

	"github.com/seifer/go-dsa/sort/internal/testutil"
)

func TestInts(t *testing.T) {
	input := testutil.InputInts()

	Ints(input)

	if !testutil.IsSortedInts(input) {
		t.Fail()
	}
}
