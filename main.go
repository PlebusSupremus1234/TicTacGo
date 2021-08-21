package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Tic Tac Toe")
	fmt.Println("---------------------")

	opponent := getOpponent()

	game := newGame(opponent)

	game.display()
}

func getOpponent() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Who would you like to play tic tac toe against? [computer/player2] ")
	text, _ := reader.ReadString('\n')
	opponent := strings.TrimSpace(text)

	if opponent == "computer" || opponent == "player2" {
		return opponent
	}

	fmt.Println("Invalid choice")
	return getOpponent()
}
