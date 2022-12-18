package board

import (
	"chess-game/models"
	"errors"
	"strconv"
	"strings"
)

var board [8][8]*models.Piece
var CurrTurn = "W"

func init() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = nil
		}
	}
}

func formatInput(mv string) (models.MoveType, error) {
	// Eg: N:g1 -> f3
	mv = strings.Trim(mv, " ")
	piece := strings.Split(mv, ":")[0]

	squares := strings.Split(strings.Split(mv, ":")[1], "->")
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

	if !(x >= 0 && x < 8 && y >= 0 && y < 8) {
		return models.Square{0, 0}, errors.New("invalid square")
	}
	return models.Square{uint(x), uint(y)}, nil
}

func Move(mv string) error {
	move, err := formatInput(mv)

	if err != nil {
		return err
	}
	var character = board[move.From[1]][move.From[0]]
	if character == nil || character.Character != move.Character || character.CurrPosition != move.From || character.Color != CurrTurn {
		return errors.New("invalid move")
	}

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
