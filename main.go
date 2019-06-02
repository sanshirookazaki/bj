package main

import (
	"fmt"

	"github.com/sanshirookazaki/bj/card"
	"github.com/sanshirookazaki/bj/player"
)

func main() {
	fmt.Println("Hello BlackJack!")
	cardbuilder := card.New()
	talon := cardbuilder.TalonPrepare().Build()
	//fmt.Println(talon)

	playerbuilder := player.New()
	player := playerbuilder.Build()
	dealer := playerbuilder.Build()

	marks, nums := talon.Deal(2)
	for m, i := range nums {
		fmt.Println("Draw", switchMark(marks[m]), i)
		player.Draw(i)
	}

	marks, nums = talon.Deal(2)
	for _, i := range nums {
		dealer.Draw(i)
	}
	fmt.Println("player start")
	player.Open(player.Compute())

	for dealer.Compute() < 18 {
		fmt.Println("dealer draw")
		marks, nums = talon.Deal(1)
		dealer.Draw(nums[0])
	}

	fmt.Println("dealer")
	dealer.Open(dealer.Compute())
}

func switchMark(markNum int) (mark string) {
	switch markNum {
	case 0:
		mark = "スペード"
	case 1:
		mark = "ハート"
	case 2:
		mark = "ダイヤ"
	case 3:
		mark = "クローバー"
	default:
		mark = ""
	}
	return mark
}
