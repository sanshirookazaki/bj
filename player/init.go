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
	Open(int)
	Compute() int
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

func (p *player) Compute() (result int) {
	for _, i := range p.hand {
		var num int
		switch i {
		case "A":
			num = 1
		case "J":
			num = 10
		case "Q":
			num = 10
		case "K":
			num = 10
		default:
			num, _ = strconv.Atoi(i)
		}
		result += num
	}
	p.result = result
	return result
}

func (p *player) Open(result int) {
	fmt.Println(result)
}
