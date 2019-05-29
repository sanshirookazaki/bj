package card

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Talon interface {
	Deal()
}

type talonBuiler struct {
	cards // "カードの種類":残り
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
	}
}

func (tb *talonBuiler) TalonPrepare() TalonBuilder {
	cardType := []string{"A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var initCard card
	initCard.state = cardType
	// cards[0] spade
	// cards[1] heart
	// cards[2] diamond
	// cards[3] clover
	for i := 0; i < 4; i++ {
		tb.cards = append(tb.cards, initCard)
	}
	return tb
}

type card struct {
	state []string
}

type cards []card

type talon struct {
	cards
}

func (t *talon) Deal() {
	rand.Seed(time.Now().UnixNano())
	// deal 2 cards to player and host
	for i := 0; i < 2; i++ {
		mark := rand.Intn(4)
		num := strconv.Itoa(rand.Intn(14))
		if contains(t.cards[mark].state, num) {
			fmt.Println("Deal", num)
		}
	}

}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
