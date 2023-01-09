package board

import (
	"chess-game/characters"
	"chess-game/models"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var board models.Board
var CurrTurn = "W"

var whitePieces = []models.Piece{
	&characters.Rook{
		CurrPosition: models.Square{0, 0},
		Color:        "W",
		Character:    "R",
	},
	&characters.Rook{
		CurrPosition: models.Square{0, 7},
		Color:        "W",
		Character:    "R",
	},
}

var blackPieces = []models.Piece{
	&characters.Rook{
		CurrPosition: models.Square{7, 7},
		Color:        "B",
		Character:    "R",
	},
	&characters.Rook{
		CurrPosition: models.Square{7, 0},
		Color:        "B",
		Character:    "R",
	},
}

func init() {
	// Initialize board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = nil
		}
	}

	for i := 0; i < len(whitePieces); i++ {
		character := whitePieces[i]
		board[character.GetPosition()[0]][character.GetPosition()[1]] = character
	}

	for i := 0; i < len(blackPieces); i++ {
		character := blackPieces[i]
		board[character.GetPosition()[0]][character.GetPosition()[1]] = character
	}

}

func formatInput(mv string) (models.MoveType, error) {
	// Eg: N:g1 -> f3
	mv = strings.Trim(mv, " ")

	// Separate character and move
	piece := strings.Split(mv, ":")[0]
	squares := strings.Split(strings.Split(mv, ":")[1], "->")

	// Validate move
	from, fromErr := getSquare(squares[0])
	to, toErr := getSquare(squares[1])

	if fromErr != nil {
		return models.MoveType{}, fromErr
	}

	if toErr != nil {
		return models.MoveType{}, toErr
	}

	return models.MoveType{
		Character: piece,
		From:      from,
		To:        to,
	}, nil
}

func getSquare(sq string) (models.Square, error) {
	dictionary := make(map[string]int)
	dictionary["a"] = 0
	dictionary["b"] = 1
	dictionary["c"] = 2
	dictionary["d"] = 3
	dictionary["e"] = 4
	dictionary["f"] = 5
	dictionary["g"] = 6
	dictionary["h"] = 7

	x := dictionary[string(sq[0])]
	y, err := strconv.Atoi(string(sq[1]))
	if err != nil {
		return models.Square{0, 0}, err
	}
	y--
	if !(x >= 0 && x < 8 && y >= 0 && y < 8) {
		return models.Square{0, 0}, errors.New("invalid square")
	}
	return models.Square{x, y}, nil
}

func Move(mv string) error {
	move, err := formatInput(mv)

	if err != nil {
		return err
	}
	var character = board[move.From[1]][move.From[0]]
	var newSquare = board[move.To[1]][move.To[0]]

	fmt.Printf("From: %v\nTo: %v\nCharacter: %v\n", move.From, move.To, character)

	isInvalidMove := character == nil ||
		character.GetCharacter() != move.Character ||
		character.GetColor() != CurrTurn ||
		newSquare != nil && newSquare.GetColor() == character.GetColor()

	if isInvalidMove {
		return errors.New("invalid move")
	}

	moveErr := character.Move(move.To, &board)

	if moveErr != nil {
		return errors.New("invalid move")
	}

	PrintBoard()

	// check whether character can go to move.to
	// if possible update temporary board
	// check whether king is checked

	// if move is valid, make the move in the real board
	// calculate attacked squares from moved player
	// calculate squares opponents king can go to
	// if 0, and king is not checked, draw
	// if 0 and king is checked, win

	// else add board to history
	// if board is consecutively repeated three times, draw
	// else ask opponent to move

	return nil
}

func PrintBoard() {
	for _, i := range board {
		for _, j := range i {
			if j == nil {
				fmt.Print("-")
			} else {
				fmt.Print(j.GetColor() + j.GetCharacter())
			}
		}
		fmt.Println()
	}
}
