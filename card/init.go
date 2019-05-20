package card

type talonBuiler struct {
	card map[string]int // "カードの種類":残り
	num  int
}

func New() TalonBuiler {
	return &talonBuiler{}
}

type TalonBuilder interface{
	Build Talon
}