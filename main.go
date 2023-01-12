package main

import (
	"chess-game/board"
	"fmt"
)

func main() {
	// Start game
	fmt.Println("Welcome to the chess game!")
	board.PrintBoard()
	for {

		// Request a move from current player
		var mv string
		if board.CurrTurn == "W" {
			fmt.Println("It's whites' turn:")
		} else {
			fmt.Println("It's blacks' turn:")
		}
		_, scanErr := fmt.Scan(&mv)

		if scanErr != nil {
			fmt.Println("Please Try Again!")
			continue
		}

		// Validate move
		err, status := board.Move(mv)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if status == "End" {
			break
		}

		// Give chance to other player
		if board.CurrTurn == "W" {
			board.CurrTurn = "B"
		} else {
			board.CurrTurn = "W"
		}
	}
}
