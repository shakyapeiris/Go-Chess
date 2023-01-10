package models

type Piece interface {
	Move(target Square, board *Board) error
	GetCharacter() string
	GetColor() string
	GetPosition() Square
	GetAttackingSquares(board Board) []Square
	GetID() int
}
