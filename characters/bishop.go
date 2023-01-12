package characters

import (
	"chess-game/models"
	"errors"
)

type Bishop struct {
	Color        string
	CurrPosition models.Square
	Character    string
	Id           int
	Prev         *models.Square
}

func (B *Bishop) Move(target models.Square, board *models.Board) error {
	var possibleSquares = B.GetAttackingSquares(*board)

	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[B.CurrPosition[1]][B.CurrPosition[0]] = nil
			B.CurrPosition = target
			board[target[1]][target[0]] = B
			return nil
		}
	}

	return errors.New("[illegal move]: cannot move the piece to requested position")
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (B *Bishop) GetAttackingSquares(board models.Board) []models.Square {
	var squares []models.Square
	if board[B.CurrPosition[1]][B.CurrPosition[0]] == nil ||
		(board[B.CurrPosition[1]][B.CurrPosition[0]] != nil &&
			board[B.CurrPosition[1]][B.CurrPosition[0]].GetID() != B.Id) {
		return []models.Square{}
	}

	x := B.CurrPosition[0]
	y := B.CurrPosition[1]

	tX := x - 1
	tY := y - 1
	for tX >= 0 && tY >= 0 {
		squares = append(squares, models.Square{tX, tY})
		if board[tY][tX] != nil {
			break
		}
		tX--
		tY--
	}

	tX = x + 1
	tY = y - 1
	for tX < 8 && tY >= 0 {
		squares = append(squares, models.Square{tX, tY})
		if board[tY][tX] != nil {
			break
		}
		tX++
		tY--
	}

	tX = x - 1
	tY = y + 1
	for tX >= 0 && tY < 8 {
		squares = append(squares, models.Square{tX, tY})
		if board[tY][tX] != nil {
			break
		}
		tX--
		tY++
	}

	tX = x + 1
	tY = y + 1
	for tX < 8 && tY < 8 {
		squares = append(squares, models.Square{tX, tY})
		if board[tY][tX] != nil {
			break
		}
		tX++
		tY++
	}

	return squares
}

func (B *Bishop) GetPosition() models.Square {
	return B.CurrPosition
}

func (B *Bishop) GetColor() string {
	return B.Color
}

func (B *Bishop) GetCharacter() string {
	return B.Character
}

func (B *Bishop) GetID() int {
	return B.Id
}

func (B *Bishop) SetID(id int) {
	B.Id = id
}

func (B *Bishop) GetPrev() *models.Square {
	return B.Prev
}

func (B *Bishop) SetPrev(prev *models.Square) {
	B.Prev = prev
}

func (B *Bishop) HardMove(sq models.Square) {
	B.CurrPosition = sq
}
