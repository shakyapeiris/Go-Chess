package models

type Square [2]int

type Board [8][8]Piece

type MoveType struct {
	Character string
	From      Square
	To        Square
}
