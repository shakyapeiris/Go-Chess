package characters

import (
	"chess-game/models"
	"errors"
	"fmt"
)

type Queen struct {
	Color        string
	CurrPosition models.Square
	Character    string
	Id           int
	Prev         *models.Square
}

func (Q *Queen) Move(target models.Square, board *models.Board) error {
	var possibleSquares = Q.GetAttackingSquares(*board)

	for _, square := range possibleSquares {
		if square[0] == target[0] && square[1] == target[1] {
			board[Q.CurrPosition[1]][Q.CurrPosition[0]] = nil
			Q.CurrPosition = target
			board[target[1]][target[0]] = Q
			return nil
		}
	}

	return errors.New("[illegal move]: cannot move the piece to requested position")
}

// GetAttackingSquares get squares piece can move/other king cannot come
func (Q *Queen) GetAttackingSquares(board models.Board) []models.Square {
	var squares []models.Square
	if board[Q.CurrPosition[1]][Q.CurrPosition[0]] == nil ||
		(board[Q.CurrPosition[1]][Q.CurrPosition[0]] != nil &&
			board[Q.CurrPosition[1]][Q.CurrPosition[0]].GetID() != Q.Id) {
		return []models.Square{}
	}

	x := Q.CurrPosition[0]
	y := Q.CurrPosition[1]
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
		squares = append(squares, models.Square{x, tY})
		if board[tY][x] != nil {
			break
		}
	}

	for tY := y - 1; tY >= 0; tY-- {
		squares = append(squares, models.Square{x, tY})
		if board[tY][x] != nil {
			break
		}
	}
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
	fmt.Println(squares)
	return squares
}

func (Q *Queen) GetPosition() models.Square {
	return Q.CurrPosition
}

func (Q *Queen) GetColor() string {
	return Q.Color
}

func (Q *Queen) GetCharacter() string {
	return Q.Character
}

func (Q *Queen) GetID() int {
	return Q.Id
}

func (Q *Queen) SetID(id int) {
	Q.Id = id
}

func (Q *Queen) GetPrev() *models.Square {
	return Q.Prev
}

func (Q *Queen) SetPrev(prev *models.Square) {
	Q.Prev = prev
}

func (Q *Queen) HardMove(sq models.Square) {
	Q.CurrPosition = sq
}
