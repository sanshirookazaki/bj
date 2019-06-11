package main

import "testing"

func TestSwitchMark(t *testing.T) {
	s := switchMark(0)
	h := switchMark(1)
	d := switchMark(2)
	c := switchMark(3)
	no := switchMark(4)

	if !(s == "スペード" && h == "ハート" && d == "ダイヤ" && c == "クローバー" && no == "") {
		t.Fatalf("error: SwitchMark %v %v %v %v %v", s, h, d, c, no)
	}
	t.Log("Pass SwitchMark func")
}

func TestBout(t *testing.T) {
	if !(bout(21, 21) == "引き分け") {
		t.Fatalf("error: Bout result %v", bout(21, 21))
	}

	t.Log("Pass Bout func")
}
