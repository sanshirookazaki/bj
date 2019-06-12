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
	score := [][]int{{1, 1}, {0, 1}, {2, 0}, {-1, 4}, {5, -1}}

	if !(bout(score[0][0], score[0][1]) == "引き分け") {
		t.Fatalf("error: Bout result %v", bout(score[0][0], score[0][1]))
	}
	if !(bout(score[1][0], score[1][1]) == "あなたの勝ち") {
		t.Fatalf("error: Bout result %v", bout(score[1][0], score[1][1]))
	}
	if !(bout(score[2][0], score[2][1]) == "ディーラーの勝ち") {
		t.Fatalf("error: Bout result %v", bout(score[2][0], score[2][1]))
	}
	if !(bout(score[3][0], score[3][1]) == "ディーラーの勝ち") {
		t.Fatalf("error: Bout result %v", bout(score[3][0], score[3][1]))
	}
	if !(bout(score[4][0], score[4][1]) == "あなたの勝ち") {
		t.Fatalf("error: Bout result %v", bout(score[4][0], score[4][1]))
	}
	t.Log("Pass Bout func")
}
