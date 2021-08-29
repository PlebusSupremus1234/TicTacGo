package main

import (
	"fmt"
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

	game.display([]Coord {}, true)

	for !game.gameOver { game.play() }
}
