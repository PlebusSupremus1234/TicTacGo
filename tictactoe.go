package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TicTacToe struct {
	opponent string
	gameover bool
	p1turn bool
	board [3][3]string
}

func newGame(opponent string) TicTacToe {
	return TicTacToe {
		opponent: opponent,
		gameover: false,
		p1turn: true,
		board: [3][3]string { { " ", " ", " " }, { " ", " ", " " }, { " ", " ", " " } },
	}
}

func (t *TicTacToe) play() {
	reader := bufio.NewReader(os.Stdin)
	marker := "X"
	if !t.p1turn { marker = "O" }

	fmt.Printf(bold("\nPlease enter the coordinates to place your marker [%v] on: "), marker)

	text, _ := reader.ReadString('\n')
	move := strings.ToLower(strings.TrimSpace(text))

	if len(move) != 2 {
		fmt.Println(red("Expected a row and a column, eg. a1 or b3"))
	} else if !inArray([]string { "a", "b", "c" }, string(move[0])) {
		fmt.Println(red("Row must be [a/b/c]"))
	} else if !inArray([]string { "1", "2", "3" }, string(move[1])) {
		fmt.Println(red("Column must be [1/2/3]"))
	} else {
		y := indexOf([]string { "a", "b", "c" }, string(move[0]))
		x := indexOf([]string { "1", "2", "3" }, string(move[1]))

		if t.board[y][x] != " " {
			fmt.Println(red("That spot is already taken"))
		} else {
			t.board[y][x] = marker
			t.p1turn = !t.p1turn

			fmt.Println("")
			t.display()

			res := t.checkGameover()
			if res != "" {
				t.gameover = true

				fmt.Println("")
				if res == "tie" {
					fmt.Println(bold("The game was a tie!"))
				} else if res == "X" {
					fmt.Println(bold(green("You win!")))
				} else {
					fmt.Println(bold(red("You lose!")))
				}
			}
		}
	}
}

func (t TicTacToe) display() {
	rows := []string {}

	for i := range t.board {
		rows = append(rows, " " + strings.Join(t.board[i][:], " | "))
	}

	fmt.Println(blue(strings.Join(rows, "\n---|---|---\n")))
}

func (t TicTacToe) checkGameover() string {
	winner := ""
	emptyspaces := 0

	for i := 0; i < 3; i++ {
		if equals3(t.board[i][0], t.board[i][1], t.board[i][2]) && t.board[i][0] != " " {
			return t.board[i][0]
		}
		if equals3(t.board[0][i], t.board[1][i], t.board[2][i]) && t.board[0][i] != " " {
			return t.board[0][i]
		}

		for j := 0; j < 3; j++ {
			if t.board[i][j] == " " {
				emptyspaces++
			}
		}
	}

	if emptyspaces == 0 { return "tie" }

	if equals3(t.board[0][0], t.board[1][1], t.board[2][2]) && t.board[0][0] != " " {
		return t.board[0][0]
	}
	if equals3(t.board[2][0], t.board[1][1], t.board[0][2]) && t.board[2][0] != " " {
		return t.board[2][0]
	}	

	return winner
}