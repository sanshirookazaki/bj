package player

type player struct {
	hand []string
}

type Player interface{}

type playerBuilder struct {
	hand []string
}

func New() PlayerBuilder {
	return &playerBuilder{}
}

type PlayerBuilder interface {
	Build() Player
	Draw(string) PlayerBuilder
}

func (p *playerBuilder) Build() Player {
	return &player{
		hand: p.hand,
	}
}

func (p *playerBuilder) Draw(card string) PlayerBuilder {
	p.hand = append(p.hand, card)
	return p
}
