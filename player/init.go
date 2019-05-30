package player

import (
	"fmt"
	"strconv"
)

type player struct {
	hand   []string
	result int
}

type Player interface {
	Draw(string)
	Open()
	Compute()
}

type playerBuilder struct {
	hand   []string
	result int
}

func New() PlayerBuilder {
	return &playerBuilder{}
}

type PlayerBuilder interface {
	Build() Player
}

func (p *playerBuilder) Build() Player {
	return &player{
		hand:   p.hand,
		result: p.result,
	}
}

func (p *player) Draw(card string) {
	p.hand = append(p.hand, card)
}

func (p *player) Compute() {
	var result int
	for _, i := range p.hand {
		var num int
		switch i {
		case "A":
			num = 1
		case "J":
			num = 11
		case "Q":
			num = 12
		case "K":
			num = 13
		default:
			num, _ = strconv.Atoi(i)
		}
		result += num
	}
	p.result = result
}

func (p *player) Open() {
	fmt.Println(p.hand, p.result)
}
