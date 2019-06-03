package player

import (
	"strconv"
)

type player struct {
	hand  []string
	sum   int
	score int
}

type Player interface {
	Draw(string)
	Sum() int
	Compute() int
}

type playerBuilder struct {
	hand  []string
	sum   int
	score int
}

func New() PlayerBuilder {
	return &playerBuilder{}
}

type PlayerBuilder interface {
	Build() Player
}

func (p *playerBuilder) Build() Player {
	return &player{
		hand:  p.hand,
		sum:   p.sum,
		score: p.score,
	}
}

func (p *player) Draw(card string) {
	p.hand = append(p.hand, card)
}

func (p *player) Sum() (sum int) {
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
		sum += num
	}
	p.sum = sum
	return sum
}

func (p *player) Compute() (score int) {
	sum := p.Sum()
	score = 21 - sum
	if score < 0 {
		score = -1
	}
	return score
}
