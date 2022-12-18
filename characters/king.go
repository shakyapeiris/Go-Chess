package characters

import "chess-game/models"

type King struct {
	models.Piece
}

func (K *King) Move(target models.Square) error {
	// left
	// left bottom
	// left top

	// right
	// right bottom
	// right top

	// bottom
	// top
	return nil
}

// CanMove Checks whether the piece can move to the entered square
func (K *King) CanMove() bool {
	return true
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
func (K *King) GetAttackingSquares() {

}
