package card

import (
	"math/rand"
	"strconv"
	"time"
)

type Talon interface {
	Deal(int) ([]int, []string)
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
	cardType := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
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

func (t *talon) Deal(loop int) (marks []int, nums []string) {
	rand.Seed(time.Now().UnixNano())

	// deal 2 cards to player and host
	for i := 0; i < loop; {
		mark := rand.Intn(4)
		num := strconv.Itoa(rand.Intn(13) + 1) // 1 ~ 13
		switch num {
		case "1":
			num = "A"
		case "11":
			num = "J"
		case "12":
			num = "Q"
		case "13":
			num = "K"
		}
		if contains(t.cards[mark].state, num) {
			marks = append(marks, mark)
			nums = append(nums, num)
			t.cards[mark].state = remove(t.cards[mark].state, num)
			i++
		}
	}
	return marks, nums
}

func contains(s interface{}, e ...interface{}) bool {
	var c int
	switch l := s.(type) {
	case []string:
		{
			for _, si := range l {
				for _, ei := range e {
					if si == ei {
						c++
					}
				}
			}
		}
	case []int:
		for _, si := range l {
			for _, ei := range e {
				if si == ei {
					c++
				}
			}
		}
	}
	if c == len(e) {
		return true
	}

	return false
}

func remove(s []string, e string) (result []string) {
	for _, i := range s {
		if i != e {
			result = append(result, i)
		}
	}
	return result
}
