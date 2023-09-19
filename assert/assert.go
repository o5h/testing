package assert

import (
	"math"
	"testing"
)

func True(t *testing.T, actual interface{}, what ...any) {
	t.Helper()
	Eq(t, actual, true, what...)
}

func False(t *testing.T, actual interface{}, what ...any) {
	t.Helper()
	Eq(t, actual, false, what...)
}

func Nil(t *testing.T, actual interface{}, what ...any) {
	t.Helper()
	if actual != nil {
		logWhat(t, what...)
		t.Fatalf("Expected nil, was %v", actual)
	}
}

func NotNil(t *testing.T, actual interface{}, what ...any) {
	t.Helper()
	if actual == nil {
		logWhat(t, what...)
		t.Fatalf("Expected not nil, was nil")
	}
}

func EqUint64(t *testing.T, actual uint64, expected uint64) {
	t.Helper()
	if actual != expected {
		t.Fatalf("Expected =>'%v'<=, was =>'%v'<=", expected, actual)
	}
}

func Eq[T comparable](t *testing.T, actual T, expected T, what ...any) {
	t.Helper()
	if actual == expected {
		return
	}
	t.Fatalf("Expected =>'%v'<=, was =>'%v'<=", expected, actual)

}

func NotEq[T comparable](t *testing.T, actual T, expected T) {
	t.Helper()
	if actual != expected {
		return
	}
	t.Fatalf("Expected not `%v`, was `%v`", expected, actual)
}

func EqEpsilon[T float32 | float64](t *testing.T, actual T, expected T, epsilon float64) {
	t.Helper()
	if math.Abs(float64(actual)-float64(expected)) > epsilon {
		t.Fatalf("Expected %f, was %f (+-%f)", expected, actual, epsilon)
	}
}

func EqSlice[T comparable](t *testing.T, actual []T, expected []T) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("Slice len expected %v, was %v", len(expected), len(actual))
	}

	for i, a := range actual {
		Eq(t, a, expected[i])
	}
}

func EqSliceEpsilon[T float32 | float64](t *testing.T, actual []T, expected []T, epsilon float64) {
	t.Helper()
	if len(actual) != len(expected) {
		t.Fatalf("Slice len expected %v, was %v", len(expected), len(actual))
	}

	for i, a := range actual {
		EqEpsilon(t, a, expected[i], epsilon)
	}
}

func logWhat(t *testing.T, what ...any) {
	if len(what) > 0 {
		t.Log(what...)
	}
}
