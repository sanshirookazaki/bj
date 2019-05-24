package card

import "fmt"

type Talon interface{}

type talonBuiler struct {
	card // "カードの種類":残り
	num  int
}

func New() TalonBuilder {
	return &talonBuiler{}
}

type TalonBuilder interface {
	Build() Talon
	TalonPrepare() card
}

func (tb *talonBuiler) Build() Talon {
	return &talon{
		card: tb.card,
		num:  tb.num,
	}
}

func (tb *talonBuiler) TalonPrepare() (talon card) {
	cardType := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for cardtype := range cardType {
		fmt.Println(cardtype)
	}
	return
}

type card struct {
	state map[string]int
}

type cards []card

type talon struct {
	cards
	num int
}
