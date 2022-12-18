package characters

import (
	"chess-game/models"
)

type Pawn struct {
	models.Piece
}

func (P *Pawn) Move(to models.Square, board *models.Board) {

}

func (P *Pawn) GetAttackingSquares(board *models.Board) {

}
