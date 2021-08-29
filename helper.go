package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bold(text string) string { return "\x1b[1m" + text + "\x1b[0m" }
func red(text string) string  { return "\x1b[31m" + text + "\x1b[0m" }
func blue(text string) string { return "\x1b[34m" + text + "\x1b[0m" }

func equals3(a, b, c interface{}) bool {
	return a == b && b == c
}

func convertCoords(coords []Coord) []interface{} {
	newArray := make([]interface{}, len(coords))
	for k, v := range coords {
		newArray[k] = v
	}
	return newArray
}

func inArray(array []interface{}, value interface{}) bool {
	for _, i := range array {
		if i == value {
			return true
		}
	}
	return false
}

func indexOf(array []string, value string) int {
	for k, v := range array {
		if value == v {
			return k
		}
	}

	return -1
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
