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
	fmt.Println("player:", player, "dealer:", dealer)

	marks, nums := talon.Deal()
	fmt.Println("get", marks, nums)
	for _, i := range nums {
		player.Draw(i)
	}

	marks, nums = talon.Deal()
	fmt.Println("get", marks, nums)
	for _, i := range nums {
		dealer.Draw(i)
	}

	player.Compute()
	player.Open()
}
