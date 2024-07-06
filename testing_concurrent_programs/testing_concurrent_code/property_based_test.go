package main

/*
Define properties or invariants that must always hold true for the system, regardless of the concurrent interactions.
Tools like QuickCheck (for Haskell) and its variants for other languages can generate test cases to try to falsify these properties.
*/

import (
	"reflect"
	"testing"
	"testing/quick"
)

// TestReverseIntsProperty checks the property that reversing a slice twice should
// return the original slice.
func TestReverseIntsProperty(t *testing.T) {
	property := func(a []int) bool {
		// Reverse the slice twice
		b := ReverseInts(a)
		c := ReverseInts(b)

		// Check if the final slice is equal to the original
		return reflect.DeepEqual(a, c)
	}

	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}
