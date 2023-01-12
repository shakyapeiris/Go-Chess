package characters

import (
	"chess-game/models"
	"errors"
)

type Pawn struct {
	Color        string
	CurrPosition models.Square
	Character    string
	Id           int
	Prev         *models.Square
}

func (P *Pawn) Move(target models.Square, board *models.Board) error {
	var possibleSquares = P.GetAttackingSquares(*board)

	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[P.CurrPosition[1]][P.CurrPosition[0]] = nil
			P.CurrPosition = target
			board[target[1]][target[0]] = P
			return nil
		}
	}

	//TODO: Enpassment
	if P.Color == "W" && P.CurrPosition[1] == 4 {

	} else if P.Color == "B" && P.CurrPosition[1] == 3 {

	}

	return errors.New("[illegal move]: cannot move the piece to requested position")
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (P *Pawn) GetAttackingSquares(board models.Board) []models.Square {
	var squares []models.Square
	if board[P.CurrPosition[1]][P.CurrPosition[0]].GetID() != P.Id {
		return []models.Square{}
	}

	x := P.CurrPosition[0]
	y := P.CurrPosition[1]

	if P.Color == "W" {
		if y+1 < 8 && board[y+1][x] == nil {
			squares = append(squares, models.Square{x, y + 1})
		}

		if y+1 < 8 && x+1 < 8 && board[y+1][x+1] != nil {
			squares = append(squares, models.Square{x + 1, y + 1})
		}

		if y+1 < 8 && x-1 >= 0 && board[y+1][x-1] != nil {
			squares = append(squares, models.Square{x - 1, y + 1})
		}
		if P.Prev == nil && board[y+2][x] == nil {
			squares = append(squares, models.Square{x, y + 2})
		}
	} else {
		if y-1 >= 0 && board[y-1][x] == nil {
			squares = append(squares, models.Square{x, y - 1})
		}

		if y-1 >= 0 && x+1 < 8 && board[y-1][x+1] != nil {
			squares = append(squares, models.Square{x + 1, y - 1})
		}

		if y-1 >= 0 && x-1 >= 0 && board[y-1][x-1] != nil {
			squares = append(squares, models.Square{x - 1, y - 1})
		}

		if P.Prev == nil && board[y-2][x] == nil {
			squares = append(squares, models.Square{x, y - 2})
		}
	}

	return squares
}

func (P *Pawn) GetPosition() models.Square {
	return P.CurrPosition
}

func (P *Pawn) GetColor() string {
	return P.Color
}

func (P *Pawn) GetCharacter() string {
	return P.Character
}

func (P *Pawn) GetID() int {
	return P.Id
}

func (P *Pawn) SetID(id int) {
	P.Id = id
}

func (P *Pawn) GetPrev() *models.Square {
	return P.Prev
}

func (P *Pawn) SetPrev(prev *models.Square) {
	P.Prev = prev
}

func (P *Pawn) HardMove(sq models.Square) {
	P.CurrPosition = sq
}
