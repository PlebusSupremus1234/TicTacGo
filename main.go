package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("")

	opponent := getOpponent()
	game := newGame(opponent)

	fmt.Println("")
	fmt.Println(bold("Tic Tac Toe"))
	fmt.Println("-----------")
	fmt.Println("You can enter a row [a/b/c] and a column [1/2/3] to specify the position of your marker, eg. a1 or b3")
	fmt.Println("")

	game.display()

	reader := bufio.NewReader(os.Stdin)
	for {
		marker := "X"
		if !game.p1turn { marker = "O" }

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

			if game.board[y][x] != " " {
				fmt.Println(red("That spot is already taken"))
			} else {
				game.board[y][x] = marker
				game.p1turn = !game.p1turn

				fmt.Println("")
				game.display()
			}
		}
	}
}

func getOpponent() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(bold("Who would you like to play tic tac toe against? [computer/player2] "))
	text, _ := reader.ReadString('\n')
	opponent := strings.ToLower(strings.TrimSpace(text))

	if opponent == "computer" || opponent == "player2" {
		return opponent
	}

	fmt.Println(red("Invalid choice"))
	return getOpponent()
}