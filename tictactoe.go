package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TicTacToe struct {
	opponent string
	gameOver bool
	p1turn   bool
	board    [3][3]string
}

func newGame(opponent string) TicTacToe {
	return TicTacToe{
		opponent: opponent,
		gameOver: false,
		p1turn:   true,
		board:    [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
	}
}

func (t *TicTacToe) play() {
	reader := bufio.NewReader(os.Stdin)
	marker := "X"
	player := "Player 1"
	if !t.p1turn {
		marker = "O"
		player = "Player 2"

		if t.opponent == "computer" {
			choice := bestMove(t)
			rows := []string{"a", "b", "c"}
			t.makeMove(rows[choice.y]+strconv.Itoa(int(choice.x)+1), marker)
			return
		}
	}

	fmt.Printf(bold("\n"+player+", please enter the coordinates to place your marker [%v] on: "), marker)

	text, _ := reader.ReadString('\n')
	move := strings.ToLower(strings.TrimSpace(text))

	if len(move) != 2 {
		fmt.Println(red("Expected a row and a column, eg. a1 or b3"))
	} else if !inArray([]interface{}{"a", "b", "c"}, string(move[0])) {
		fmt.Println(red("Row must be [a/b/c]"))
	} else if !inArray([]interface{}{"1", "2", "3"}, string(move[1])) {
		fmt.Println(red("Column must be [1/2/3]"))
	} else {
		t.makeMove(move, marker)
	}
}

func (t TicTacToe) display(highlighting []Coord, markings bool) {
	output := "   1   2   3\n"
	array := []string{"a", "b", "c"}
	spacing := "  "

	if !markings {
		output = ""
		array = []string{" ", " ", " "}
		spacing = ""
	}

	for i := range t.board {
		output += array[i] + spacing

		for j := range t.board[i] {
			if (inArray(convertCoords(highlighting), Coord{j, i})) {
				output += "\x1b[32m" + t.board[i][j] + "\x1b[34m"
			} else {
				output += t.board[i][j]
			}

			if j < 2 {
				output += " | "
			}
		}

		output += "\n"
		if i < 2 {
			output += spacing + "---|---|---\n"
		}
	}

	fmt.Println(blue(output))
}

func (t *TicTacToe) makeMove(move string, marker string) {
	y := indexOf([]string{"a", "b", "c"}, string(move[0]))
	x := indexOf([]string{"1", "2", "3"}, string(move[1]))

	if t.board[y][x] != " " {
		fmt.Println(red("That spot is already taken"))
	} else {
		t.board[y][x] = marker
		t.p1turn = !t.p1turn

		fmt.Println("")

		res := t.checkGameOver()
		if res.winner != "" {
			t.gameOver = true

			if res.winner == "tie" {
				fmt.Println(bold("The game was a tie!"))
			} else if res.winner == "X" {
				fmt.Println(bold("\x1b[32mYou win!\x1b[0m"))
			} else {
				fmt.Println(bold(red("You lose!")))
			}

			t.display(res.coords, false)
		} else {
			t.display([]Coord{}, true)
		}
	}
}

func (t TicTacToe) checkGameOver() OutputRes {
	emptySpaces := 0

	for i := 0; i < 3; i++ {
		if equals3(t.board[i][0], t.board[i][1], t.board[i][2]) && t.board[i][0] != " " {
			return OutputRes{
				winner: t.board[i][0],
				coords: []Coord{{x: 0, y: i}, {x: 1, y: i}, {x: 2, y: i}},
			}
		}

		if equals3(t.board[0][i], t.board[1][i], t.board[2][i]) && t.board[0][i] != " " {
			return OutputRes{
				winner: t.board[0][i],
				coords: []Coord{{x: i, y: 0}, {x: i, y: 1}, {x: i, y: 2}},
			}
		}

		for j := 0; j < 3; j++ {
			if t.board[i][j] == " " {
				emptySpaces++
			}
		}
	}

	if emptySpaces == 0 {
		return OutputRes{
			winner: "tie",
			coords: []Coord{},
		}
	}

	if equals3(t.board[0][0], t.board[1][1], t.board[2][2]) && t.board[0][0] != " " {
		return OutputRes{
			winner: t.board[0][0],
			coords: []Coord{{x: 0, y: 0}, {x: 1, y: 1}, {x: 2, y: 2}},
		}
	}

	if equals3(t.board[2][0], t.board[1][1], t.board[0][2]) && t.board[2][0] != " " {
		return OutputRes{
			winner: t.board[2][0],
			coords: []Coord{{x: 0, y: 2}, {x: 1, y: 1}, {x: 2, y: 0}},
		}
	}

	return OutputRes{
		winner: "",
		coords: []Coord{},
	}
}
