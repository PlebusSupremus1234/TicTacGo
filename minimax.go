package main

import (
	"math"
)

func bestMove(t *TicTacToe) Coord {
	bestScore := math.Inf(-1)
	var move Coord

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.board[i][j] == " " {
				t.board[i][j] = "O"
				score := minimax(t, false)
				t.board[i][j] = " "

				if score > bestScore {
					bestScore = score
					move = Coord{j, i}
				}
			}
		}
	}

	return move
}

func minimax(t *TicTacToe, isMax bool) float64 {
	res := t.checkGameOver()
	if res.winner != "" {
		if res.winner == "X" {
			return -1
		} else if res.winner == "O" {
			return 1
		} else {
			return 0
		}
	}

	if isMax {
		bestScore := math.Inf(-1)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if t.board[i][j] == " " {
					t.board[i][j] = "O"
					score := minimax(t, false)
					t.board[i][j] = " "
					bestScore = math.Max(score, bestScore)
				}
			}
		}

		return bestScore
	} else {
		bestScore := math.Inf(1)

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if t.board[i][j] == " " {
					t.board[i][j] = "X"
					score := minimax(t, true)
					t.board[i][j] = " "
					bestScore = math.Min(score, bestScore)
				}
			}
		}
		
		return bestScore
	}
}
