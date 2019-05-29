package player

import "fmt"

type player struct {
	hand []string
}

type Player interface {
	Draw(string)
	Open()
}

type playerBuilder struct {
	hand []string
}

func New() PlayerBuilder {
	return &playerBuilder{}
}

type PlayerBuilder interface {
	Build() Player
}

func (p *playerBuilder) Build() Player {
	return &player{
		hand: p.hand,
	}
}

func (p *player) Draw(card string) {
	p.hand = append(p.hand, card)
}

func (p *player) Open() {
	fmt.Println(p.hand)
}
