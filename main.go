package main

import (
	"fmt"

	"github.com/sanshirookazaki/bj/card"
)

func main() {
	fmt.Println("Hello BlackJack!")
	builder := card.New()
	talon := builder.TalonPrepare().Build()
	fmt.Println(talon)
	talon.Deal()
}
