package player

import (
	"testing"
)

func TestSum(t *testing.T) {
	playerbuilder := New()
	player := playerbuilder.Build()
	player.Draw("A")
	player.Draw("J")
	//expect := []string{"A", "J"}
	//if !reflect.DeepEqual(player.hand == expect) {
	//	t.Fatalf("error: %v", player.hand)
	//}
	if player.Sum() != 21 {
		t.Fatalf("error: sum %v", player.Sum())
	}

	t.Log("Pass Sum func")
}

func TestCompute(t *testing.T) {
	playerbuilder := New()
	player := playerbuilder.Build()
	player.Draw("A")
	player.Draw("A")
	player.Draw("A")
	player.Draw("A")
	player.Draw("7")
	if player.Compute() != 0 {
		t.Fatalf("error: compute %v", player.Compute())
	}

	t.Log("Pass Compute func")
}
