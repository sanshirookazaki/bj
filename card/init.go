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
}

func (tb *talonBuiler) Build() Talon {
	return &talon{
		card: tb.card,
		num:  tb.num,
	}
}

type talon struct {
	card map[string]int
	num  int
}
