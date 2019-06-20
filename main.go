package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sanshirookazaki/bj/card"
	"github.com/sanshirookazaki/bj/player"
)

func main() {
	fmt.Println("Start BlackJack!")
	cardbuilder := card.New()
	talon := cardbuilder.TalonPrepare().Build()

	playerbuilder := player.New()
	player := playerbuilder.Build()
	dealer := playerbuilder.Build()

	marks, nums := talon.Deal(2)
	for m, i := range nums {
		fmt.Println("Draw", switchMark(marks[m]), i)
		player.Draw(i)
	}

	marks, nums = talon.Deal(2)
	for _, i := range nums {
		dealer.Draw(i)
	}
	fmt.Println("手札の合計は", player.Sum())

	stdin := bufio.NewScanner(os.Stdin)
	if player.Compute() != 0 {
		fmt.Println("カードをドローする場合は h 、 引かない場合は s を入力してください")
	}

	for stdin.Scan() {
		input := stdin.Text()

		if player.Compute() <= 0 {
			break
		}

		if !(input == "h" || input == "s") {
			fmt.Println("hまたはsを入力してください")
			continue
		}

		if input == "s" {
			break
		}

		if input == "h" {
			marks, nums = talon.Deal(1)
			player.Draw(nums[0])
			fmt.Println("Draw", switchMark(marks[0]), nums[0])
			fmt.Println("手札の合計は", player.Sum())
			if player.Compute() == -1 {
				break
			}
		}

		if dealer.Sum() < 18 {
			fmt.Println("dealer draw")
			marks, nums = talon.Deal(1)
			dealer.Draw(nums[0])
		}
	}

	for dealer.Sum() < 18 {
		fmt.Println("dealer draw")
		marks, nums = talon.Deal(1)
		dealer.Draw(nums[0])
	}
	fmt.Println("dealer stop")
	fmt.Println("ディーラーの手札は", dealer.Sum())
	playerScore := player.Compute()
	dealerScore := dealer.Compute()
	fmt.Println(bout(playerScore, dealerScore))
}

func switchMark(markNum int) (mark string) {
	switch markNum {
	case 0:
		mark = "スペード"
	case 1:
		mark = "ハート"
	case 2:
		mark = "ダイヤ"
	case 3:
		mark = "クローバー"
	default:
		mark = ""
	}
	return mark
}

func bout(playerScore, dealerScore int) (result string) {
	switch {
	case playerScore == dealerScore:
		result = "引き分け"
	case playerScore == -1 && dealerScore >= 0:
		result = "ディーラーの勝ち"
	case playerScore >= 0 && dealerScore == -1:
		result = "あなたの勝ち"
	case playerScore < dealerScore:
		result = "あなたの勝ち"
	case playerScore > dealerScore:
		result = "ディーラーの勝ち"
	}
	return result
}
