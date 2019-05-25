package card

import "fmt"

type Talon interface {
	Deal()
}

type talonBuiler struct {
	cards // "カードの種類":残り
	num   int
}

func New() TalonBuilder {
	return &talonBuiler{}
}

type TalonBuilder interface {
	Build() Talon
	TalonPrepare() TalonBuilder
}

func (tb *talonBuiler) Build() Talon {
	return &talon{
		cards: tb.cards,
		num:   tb.num,
	}
}

func (tb *talonBuiler) TalonPrepare() TalonBuilder {
	cardType := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for cardtype := range cardType {
		fmt.Println(cardtype)
	}
	return tb
}

type card struct {
	state []string
}

type cards []card

type talon struct {
	cards cards
	num   int
}

func (t *talon) Deal() {
	fmt.Println("talon!")
}
