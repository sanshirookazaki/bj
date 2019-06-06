package card

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	a := []string{"a", "bb", "ccc"}
	if !contains(a, "bb") {
		t.Fatalf("error: %v", a)
	}
	t.Log("Pass contains func")
}

func TestRemove(t *testing.T) {
	a := []string{"a", "bb", "ccc"}
	result := remove(a, "ccc")
	expect := []string{"a", "bb"}
	if !reflect.DeepEqual(result, expect) {
		t.Fatalf("error: expect %v but result %v", expect, result)
	}
	t.Log("Pass remove func")
}

func TestTalonPrepare(t *testing.T) {
	cardbuilder := New()
	talon := cardbuilder.TalonPrepare().Build()
	marks, nums := talon.Deal(1)
	if 0 > marks[0] || marks[0] > 3 {
		t.Fatalf("error: %v %v", marks, nums)
	}

	t.Log("Pass TalonPrepare func")
}
