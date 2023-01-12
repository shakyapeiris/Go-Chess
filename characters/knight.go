package characters

import (
	"chess-game/models"
	"errors"
)

type Knight struct {
	Color        string
	CurrPosition models.Square
	Character    string
	Id           int
	Prev         *models.Square
}

func (N *Knight) Move(target models.Square, board *models.Board) error {
	var possibleSquares = N.GetAttackingSquares(*board)

	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[N.CurrPosition[1]][N.CurrPosition[0]] = nil
			N.CurrPosition = target
			board[target[1]][target[0]] = N
			return nil
		}
	}

	return errors.New("[illegal move]: cannot move the piece to requested position")
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (N *Knight) GetAttackingSquares(board models.Board) []models.Square {
	var squares []models.Square
	if board[N.CurrPosition[1]][N.CurrPosition[0]].GetID() != N.Id {
		return []models.Square{}
	}

	x := N.CurrPosition[0]
	y := N.CurrPosition[1]

	newSquares := [][2]int{
		{x + 2, y + 1},
		{x + 1, y + 2},
		{x - 2, y + 1},
		{x - 1, y + 2},
		{x + 2, y - 1},
		{x + 1, y - 2},
		{x - 2, y - 1},
		{x - 1, y - 2}}

	for _, tS := range newSquares {
		tX := tS[0]
		tY := tS[1]

		if tX >= 0 && tY >= 0 && tY < 8 && tX < 8 {
			squares = append(squares, tS)
		}
	}
	return squares
}

func (N *Knight) GetPosition() models.Square {
	return N.CurrPosition
}

func (N *Knight) GetColor() string {
	return N.Color
}

func (N *Knight) GetCharacter() string {
	return N.Character
}

func (N *Knight) GetID() int {
	return N.Id
}

func (N *Knight) SetID(id int) {
	N.Id = id
}

func (N *Knight) GetPrev() *models.Square {
	return N.Prev
}

func (N *Knight) SetPrev(prev *models.Square) {
	N.Prev = prev
}

func (N *Knight) HardMove(sq models.Square) {
	N.CurrPosition = sq
}
