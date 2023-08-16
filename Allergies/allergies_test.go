package allergies

import (
	"reflect"
	"testing"
)

func TestAllergies(t *testing.T) {
	for _, test := range listTests {
		if actual := Allergies(test.score); !sameSliceElements(actual, test.expected) {
			t.Fatalf("FAIL: Allergies(%d): expected %#v, actual %#v", test.score, test.expected, actual)
		} else {
			t.Logf("PASS: Allergic to %#v", test.expected)
		}
	}
}

func TestAllergicTo(t *testing.T) {
	for _, test := range allergicToTests {
		actual := AllergicTo(test.input.score, test.input.allergen)
		if actual != test.expected {
			t.Fatalf("FAIL: AllergicTo(%d, %s): expected %t, actual %t", test.input.score, test.input.allergen, test.expected, actual)
		} else {
			t.Logf("PASS: AllergicTo(%d, %s) %t", test.input.score, test.input.allergen, actual)
		}

	}
}

// stringSet is a set of strings
type stringSet map[string]bool

// sameSliceElements checks if the slices have the same number of elements
// regardless of their order
func sameSliceElements(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	return reflect.DeepEqual(sliceSet(a), sliceSet(b))
}

// sliceSet creates a new stringSet from a slice of strings
func sliceSet(list []string) stringSet {
	set := make(stringSet)

	for _, elem := range list {
		set[elem] = true
	}

	return set
}
