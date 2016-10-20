// This file should be auto generated

package timsort

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

func TestInts1k(t *testing.T) {
	input := testutil.InputInts1k()

	Ints(input)

	if !testutil.IsSortedInts1k(input) {
		t.Fail()
	}
}
