package characters

import (
	"chess-game/models"
	"errors"
)

type Rook struct {
	Color        string
	CurrPosition models.Square
	Character    string
}

func (R *Rook) Move(target models.Square, board *models.Board) error {
	var possibleSquares = R.GetAttackingSquares(*board)
	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[R.CurrPosition[1]][R.CurrPosition[0]] = nil
			R.CurrPosition = target
			board[target[1]][target[0]] = R
			return nil
		}
	}

	return errors.New("cannot move the piece to requested position")
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (R *Rook) GetAttackingSquares(board models.Board) []models.Square {
	var squares []models.Square

	x := R.CurrPosition[0]
	y := R.CurrPosition[1]
	for tX := x + 1; tX < 8; tX++ {
		squares = append(squares, models.Square{tX, y})
		if board[y][tX] != nil {
			break
		}
	}
	for tX := x - 1; tX >= 0; tX-- {
		squares = append(squares, models.Square{tX, y})
		if board[y][tX] != nil {
			break
		}
	}

	for tY := y + 1; tY < 8; tY++ {
		squares = append(squares, models.Square{tY, x})
		if board[tY][x] != nil {
			break
		}
	}

	for tY := y - 1; tY >= 0; tY-- {
		squares = append(squares, models.Square{tY, x})
		if board[tY][x] != nil {
			break
		}
	}

	return squares
}

func (R *Rook) GetPosition() models.Square {
	return R.CurrPosition
}

func (R *Rook) GetColor() string {
	return R.Color
}

func (R *Rook) GetCharacter() string {
	return R.Character
}
