package main

import (
	"errors"
	"fmt"
)

type Game struct {
	board      [9]string
	player     string
	turnNumber int
}

func main() {
	var game Game
	game.player = "O"

	gameOver := false
	var winner string

	for gameOver != true {
		PrintBoard(game.board)
		move := AskForPlay()
		err := game.play(move)
		if err != nil {
			fmt.Println(err)
			continue
		}

		gameOver, winner = CheckForWinner(game.board, game.turnNumber)
	}
	PrintBoard(game.board)
	if winner == "" {
		fmt.Println("it's a draw ")
	} else {
		fmt.Printf("YaaY %s is winner ", winner)
	}
}

func CheckForWinner(b [9]string, n int) (bool, string) {

	test := false
	i := 0

	for i < 9 {
		test = b[i] == b[i+1] && b[i+1] == b[i+2] && b[i] != ""
		if !test {
			i += 3
		} else {
			return true, b[i]
		}
	}
	i = 0

	for i < 3 {
		test = b[i] == b[i+3] && b[i+3] == b[i+6] && b[i] != ""
		if !test {
			i += 1
		} else {
			return true, b[i]
		}
	}

	if b[0] == b[4] && b[4] == b[8] && b[0] != "" {
		return true, b[i]
	}

	if b[2] == b[4] && b[4] == b[6] && b[2] != "" {
		return true, b[i]
	}
	if n == 9 {
		return true, ""
	}
	return false, ""
}

func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return
	}
	game.player = "O"
}

func (game *Game) play(pos int) error {
	if game.board[pos-1] == "" {
		game.board[pos-1] = game.player
		game.SwitchPlayers()
		game.turnNumber += 1
		return nil
	}
	return errors.New("try another move")
}

func AskForPlay() int {
	var moveInt int
	fmt.Println("Enter Pos to play: ")
	fmt.Scan(&moveInt)
	return moveInt
}

func PrintBoard(b [9]string) {
	for i, v := range b {
		if v == "" {
			fmt.Printf(" ")
		} else {
			fmt.Printf(v)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")

		}

	}
}
