package main

import (
	"fmt"
	"strings"
)

type TicTacToe struct {
	opponent string
	gameover bool
	board [3][3]string
}

func newGame(opponent string) TicTacToe {
	return TicTacToe {
		opponent: opponent,
		gameover: false,
		board: [3][3]string { { " ", " ", " " }, { " ", " ", " " }, { " ", " ", " " } },
	}
}

func (t TicTacToe) display() {
	rows := []string {}

	for i := range t.board {
		rows = append(rows, " " + strings.Join(t.board[i][:], " | "))
	}

	fmt.Println(strings.Join(rows, "\n---|---|---\n"))
}