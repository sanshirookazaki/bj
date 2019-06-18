package card

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	a := []string{"a", "bb", "ccc"}
	b := []int{1, 3, 5, 7, 9}
	if !contains(a, "bb") {
		t.Fatalf("error: %v", a)
	}
	if contains(a, "c") {
		t.Fatalf("error: %v", a)
	}
	if !contains(b, 5) {
		t.Fatalf("error: %v", b)
	}
	if contains(b, 4) {
		t.Fatalf("error: %v", b)
	}
	t.Log("Pass contains func")
}

func TestRemove(t *testing.T) {
	a := [][]string{{"a", "bb", "ccc"}, {"aaa", "bb", "c"}, {"a", "a", "a", "b"}}
	target := []string{"ccc", "d", "a"}
	expect := [][]string{{"a", "bb"}, {"aaa", "bb", "c"}, {"b"}}
	for i := range a {
		result := remove(a[i], target[i])
		if !reflect.DeepEqual(expect[i], result) {
			t.Fatalf("error: expect %v but result %v", expect[i], result)
		}
	}
	t.Log("Pass remove func")
}

func TestDeal(t *testing.T) {
	cardbuilder := New()
	talon := cardbuilder.TalonPrepare().Build()
	marks, nums := talon.Deal(45)
	// check marks
	for _, v := range marks {
		if v < 0 || v > 3 {
			t.Fatalf("error: %v ", v)
		}
	}

	// check nums
	cardType := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var c int
	for _, n := range nums {
		for _, ct := range cardType {
			if n == ct {
				c++
			}
		}
	}

	if !(c == len(nums)) {
		t.Fatalf("error: deal cards %v %v", marks, nums)
	}

	t.Log("Pass Deal func")
}
