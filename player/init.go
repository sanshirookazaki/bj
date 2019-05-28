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
	PlayerPrepare() PlayerBuilder
}

func (p *playerBuilder) Build() Player {
	return &player{
		hand: p.hand,
	}
}

func (p *playerBuilder) PlayerPrepare() PlayerBuilder {
	return p
}
