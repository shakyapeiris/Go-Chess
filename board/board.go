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
	&characters.King{
		CurrPosition: models.Square{3, 0},
		Color:        "W",
		Character:    "K",
		Id:           0,
	},
	&characters.Rook{
		CurrPosition: models.Square{0, 0},
		Color:        "W",
		Character:    "R",
		Id:           1,
	},
	&characters.Rook{
		CurrPosition: models.Square{7, 0},
		Color:        "W",
		Character:    "R",
		Id:           2,
	},
}

var blackPieces = []models.Piece{
	&characters.King{
		CurrPosition: models.Square{3, 7},
		Color:        "B",
		Character:    "K",
		Id:           3,
	},
	&characters.Rook{
		CurrPosition: models.Square{7, 7},
		Color:        "B",
		Character:    "R",
		Id:           4,
	},
	&characters.Rook{
		CurrPosition: models.Square{0, 7},
		Color:        "B",
		Character:    "R",
		Id:           5,
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
		board[character.GetPosition()[1]][character.GetPosition()[0]] = character
	}

	for i := 0; i < len(blackPieces); i++ {
		character := blackPieces[i]
		board[character.GetPosition()[1]][character.GetPosition()[0]] = character
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
	dictionary["h"] = 0
	dictionary["g"] = 1
	dictionary["f"] = 2
	dictionary["e"] = 3
	dictionary["d"] = 4
	dictionary["c"] = 5
	dictionary["b"] = 6
	dictionary["a"] = 7

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
		(newSquare != nil &&
			(newSquare.GetColor() == character.GetColor() ||
				newSquare.GetCharacter() == "K"))

	if isInvalidMove {
		return errors.New("invalid move")
	}

	tempBoard := board

	moveErr := character.Move(move.To, &tempBoard)
	if moveErr != nil {
		return errors.New("invalid move")
	}

	if CurrTurn == "B" {
		blackKing := blackPieces[0]
		for _, piece := range whitePieces {
			for _, sq := range piece.GetAttackingSquares(tempBoard) {
				if blackKing.GetPosition()[0] == sq[0] && blackKing.GetPosition()[1] == sq[1] {
					character.Move(move.From, &tempBoard)
					return errors.New("illegal move")
				}
			}
		}
	} else {
		whiteKing := whitePieces[0]
		for _, piece := range blackPieces {
			for _, sq := range piece.GetAttackingSquares(tempBoard) {
				if whiteKing.GetPosition()[0] == sq[0] && whiteKing.GetPosition()[1] == sq[1] {
					character.Move(move.From, &tempBoard)
					return errors.New("illegal move")
				}
			}
		}
	}

	board = tempBoard

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
