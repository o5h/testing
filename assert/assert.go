package assert

import (
	"math"
	"reflect"
	"testing"
)

func True(t *testing.T, description string, actual interface{}) {
	t.Helper()
	Eq(t, description, actual, true)
}

func False(t *testing.T, description string, actual interface{}) {
	t.Helper()
	Eq(t, description, actual, false)
}

func Nil(t *testing.T, description string, actual interface{}) {
	t.Helper()
	if actual != nil {
		t.Fatalf("Expected nil, was %v", actual)
	}
}

func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual == nil {
		t.Fatalf("Expected not nil, was nil")
	}
}

func Eq(t *testing.T, description string, actual interface{}, expected interface{}) {
	t.Helper()
	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(expected)
	if actualType != expectedType {
		t.Fatalf("%v. Expected =>'%T'<=, was =>'%T'<=", description, expectedType, actualType)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v. Expected =>'%v'<=, was =>'%v'<=", description, expected, actual)
	}
}

func NotEq(t *testing.T, description string, actual interface{}, expected interface{}) {
	t.Helper()
	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(expected)
	if actualType != expectedType {
		t.Fatalf("%v. Expected %T, was %T", description, expectedType, actualType)
	}
	if reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v. Expected not %v, was %v", description, expected, actual)
	}
}

func EqEpsilon(t *testing.T, actual interface{}, expected interface{}, epsilon float64) {
	t.Helper()
	act := asFloat64(t, actual)
	exp := asFloat64(t, expected)
	if math.Abs(act-exp) > epsilon {
		t.Fatalf("Expected %f, was %f (+-%f)", exp, act, epsilon)
	}
}

func EqSlice(t *testing.T, actual interface{}, expected interface{}, epsilon float64) {
	t.Helper()
	var act reflect.Value
	var exp reflect.Value

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Slice:
		act = reflect.ValueOf(actual)
	default:
		t.Fatalf("Slice is expected, was %T", actual)
	}

	switch reflect.TypeOf(expected).Kind() {
	case reflect.Slice:
		exp = reflect.ValueOf(expected)
	default:
		t.Fatalf("Slice is expected, was %T", expected)
	}

	if exp.Len() != act.Len() {
		t.Fatalf("Slice len expected %v, was %v", exp.Len(), act.Len())
	}

	for i := 0; i < exp.Len(); i++ {
		expValue := exp.Index(i).Interface()
		actValue := act.Index(i).Interface()
		EqEpsilon(t, actValue, expValue, epsilon)
	}
}

func asFloat64(t *testing.T, i interface{}) float64 {
	switch v := i.(type) {
	case uint:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	}
	t.Fatalf("Unsupported type %T for %v", i, i)
	return 0
}
