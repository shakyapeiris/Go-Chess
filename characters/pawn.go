package characters

import "chess-game/board"

type Pawn struct {
	Piece
}

func (P *Pawn) Move(to board.Square, board *board.Board) {

}

func (P *Pawn) GetAttackingSquares(board *board.Board) {

}
