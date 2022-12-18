package board

import "chess-game/characters"

type Square [2]uint

type Board *[8][8]characters.Piece

type MoveType struct {
	character string
	from      Square
	to        Square
}
