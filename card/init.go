package card

type Talon interface{}

type talonBuiler struct {
	card map[string]int // "カードの種類":残り
	num  int
}

func New() TalonBuilder {
	return &talonBuiler{}
}

type TalonBuilder interface {
	Build() Talon
	TalonPrepare() map[string]int
}

func (tb *talonBuiler) Build() Talon {
	return &talon{
		card: tb.card,
		num:  tb.num,
	}
}

func (tb *talonBuiler) TalonPrepare() (talon map[string]int) {
	cardType = ["A", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"]
}

type talon struct {
	card map[string]int
	num  int
}
