package characters

import (
	"chess-game/models"
	"errors"
)

type King struct {
	Color        string
	CurrPosition models.Square
	Character    string
	Id           int
	Prev         *models.Square
}

func (K *King) Move(target models.Square, board *models.Board) error {
	var possibleSquares = K.GetAttackingSquares(*board)

	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[K.CurrPosition[1]][K.CurrPosition[0]] = nil
			K.CurrPosition = target
			board[target[1]][target[0]] = K
			return nil
		}
	}

	return errors.New("cannot move the piece to requested position")
}

// IsChecked checks whether the king is checked
func (K *King) IsChecked(attackingSquares []models.Square) bool {
	for _, square := range attackingSquares {
		if K.CurrPosition == square {
			return true
		}
	}
	return false
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (K *King) GetAttackingSquares(_ models.Board) []models.Square {
	var squares []models.Square

	x := K.CurrPosition[0]
	y := K.CurrPosition[1]

	if x > 0 && y > 0 {
		squares = append(squares, models.Square{x - 1, y - 1})
	}
	if x > 0 && y < 7 {
		squares = append(squares, models.Square{x - 1, y + 1})
	}
	if x < 7 && y > 0 {
		squares = append(squares, models.Square{x + 1, y - 1})
	}
	if x < 7 && y < 7 {
		squares = append(squares, models.Square{x + 1, y + 1})
	}
	if x > 0 {
		squares = append(squares, models.Square{x - 1, y})
	}
	if x < 7 {
		squares = append(squares, models.Square{x + 1, y})
	}
	if y > 0 {
		squares = append(squares, models.Square{x, y - 1})
	}
	if y < 7 {
		squares = append(squares, models.Square{x, y + 1})
	}

	return squares
}

func (K *King) GetPosition() models.Square {
	return K.CurrPosition
}

func (K *King) GetColor() string {
	return K.Color
}

func (K *King) GetCharacter() string {
	return K.Character
}

func (K *King) GetID() int {
	return K.Id
}

func (K *King) SetID(id int) {
	K.Id = id
}

func (K *King) GetPrev() models.Square {
	return *K.Prev
}

func (K *King) SetPrev(prev *models.Square) {
	K.Prev = prev
}
