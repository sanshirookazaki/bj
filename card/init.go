package card

type Talon interface {
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
	var initCard card
	initCard.state = cardType
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
	num int
}
