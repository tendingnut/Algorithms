package sorts_test

import (
	"testing"

	"github.com/tendingnut/algorithms/sorts"
)

func TestSorted(t *testing.T) {
	sortedInc := []int{0, 3, 3, 4, 5, 5, 5, 5, 6, 21, 25}
	sortedDec := []int{432, 325, 323, 323, 2, 2, 2, 2, 1, 1, 0}
	sortedSame := []int{3, 3, 3, 3, 3}
	sortedEmpty := []int{}
	unsorted := []int{3, 5, 1, 2, 3, 3, 12, 2, 45, 2}
	expected := []bool{true, true, true, true, false}
	got := []bool{}
	got = append(got, sorts.Sorted(sortedInc))
	got = append(got, sorts.Sorted(sortedDec))
	got = append(got, sorts.Sorted(sortedSame))
	got = append(got, sorts.Sorted(sortedEmpty))
	got = append(got, sorts.Sorted(unsorted))
	for i := 0; i < len(got); i++ {
		if expected[i] != got[i] {
			t.Errorf("expected: %v. got: %v on test %v.",
				expected[i], got[i], i+1)
		}
	}
}
