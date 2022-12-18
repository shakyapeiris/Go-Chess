package main

import (
	"chess-game/board"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the chess game!")
	for {
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

		err := board.Move(mv)

		if err != nil {
			fmt.Println(err)
			continue
		}

		// board.CurrTurn = "B"
	}
}
