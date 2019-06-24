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
	hand := [][]string{{"A", "A", "A", "A", "7"}, {"A", "J"}, {"K", "Q"}, {"10", "5"}}
	result := []int{0, 0, 1, 6}
	for i, cards := range hand {
		playerbuilder := New()
		player := playerbuilder.Build()
		for _, c := range cards {
			player.Draw(c)
		}
		if player.Compute() != result[i] {
			t.Fatalf("error: compute %v", player.Compute())
		}
	}
	t.Log("Pass Compute func")
}
