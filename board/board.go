package board

import (
	"chess-game/characters"
	"chess-game/models"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var board models.Board
var CurrTurn = "W"
var history = make([]string, 0)

var whitePieces = []models.Piece{
	&characters.King{
		CurrPosition: models.Square{3, 0},
		Color:        "W",
		Character:    "K",
	},
	&characters.Queen{
		CurrPosition: models.Square{4, 0},
		Color:        "W",
		Character:    "Q",
	},
	&characters.Rook{
		CurrPosition: models.Square{0, 0},
		Color:        "W",
		Character:    "R",
	},
	&characters.Rook{
		CurrPosition: models.Square{7, 0},
		Color:        "W",
		Character:    "R",
	},
	&characters.Bishop{
		CurrPosition: models.Square{2, 0},
		Color:        "W",
		Character:    "B",
	},
	&characters.Bishop{
		CurrPosition: models.Square{5, 0},
		Color:        "W",
		Character:    "B",
	},
	&characters.Knight{
		CurrPosition: models.Square{1, 0},
		Color:        "W",
		Character:    "N",
	},
	&characters.Knight{
		CurrPosition: models.Square{6, 0},
		Color:        "W",
		Character:    "N",
	},
}

var blackPieces = []models.Piece{
	&characters.King{
		CurrPosition: models.Square{3, 7},
		Color:        "B",
		Character:    "K",
	},
	&characters.Queen{
		CurrPosition: models.Square{4, 7},
		Color:        "B",
		Character:    "Q",
	},
	&characters.Rook{
		CurrPosition: models.Square{7, 7},
		Color:        "B",
		Character:    "R",
	},
	&characters.Rook{
		CurrPosition: models.Square{0, 7},
		Color:        "B",
		Character:    "R",
	},
	&characters.Bishop{
		CurrPosition: models.Square{2, 7},
		Color:        "B",
		Character:    "B",
	},
	&characters.Bishop{
		CurrPosition: models.Square{5, 7},
		Color:        "B",
		Character:    "B",
	},
	&characters.Knight{
		CurrPosition: models.Square{1, 7},
		Color:        "B",
		Character:    "N",
	},
	&characters.Knight{
		CurrPosition: models.Square{6, 7},
		Color:        "B",
		Character:    "N",
	},
}

// init: Initialize game
func init() {
	// Initialize board
	for i := 0; i < 8; i++ {
		whitePieces = append(whitePieces, &characters.Pawn{
			CurrPosition: models.Square{i, 1},
			Color:        "W",
			Character:    "P",
		})
	}
	for i := 0; i < 8; i++ {
		blackPieces = append(blackPieces, &characters.Pawn{
			CurrPosition: models.Square{i, 6},
			Color:        "B",
			Character:    "P",
		})
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = nil
		}
	}
	var id = 0
	for i := 0; i < len(whitePieces); i++ {
		character := whitePieces[i]
		character.SetID(id)
		board[character.GetPosition()[1]][character.GetPosition()[0]] = character
		id++
	}

	for i := 0; i < len(blackPieces); i++ {
		character := blackPieces[i]
		character.SetID(id)
		board[character.GetPosition()[1]][character.GetPosition()[0]] = character
		id++
	}
}

// Move : React to user inputs
func Move(mv string) (error, string) {
	move, err := formatInput(mv)

	if err != nil {
		return err, ""
	}
	var character = board[move.From[1]][move.From[0]]
	var newSquare = board[move.To[1]][move.To[0]]
	currPrev := character.GetPrev()

	fmt.Printf("From: %v\nTo: %v\nCharacter: %v\n", move.From, move.To, character)

	isInvalidMove := character == nil ||
		character.GetCharacter() != move.Character ||
		character.GetColor() != CurrTurn ||
		(newSquare != nil &&
			(newSquare.GetColor() == character.GetColor() ||
				newSquare.GetCharacter() == "K"))

	if isInvalidMove {
		return errors.New("invalid move"), ""
	}

	tempBoard := board

	isEnpassment := character.GetCharacter() == "P" &&
		(move.From[0] != move.To[0]) &&
		board[move.To[1]][move.To[0]] == nil

	moveErr := character.Move(move.To, &tempBoard)
	if moveErr != nil {
		return moveErr, ""
	}

	if isEnpassment {
		fmt.Println("Enpassing...")
		fmt.Println(tempBoard[move.From[1]][move.To[0]])
		tempBoard[move.From[1]][move.To[0]] = nil
	}
	fmt.Println("Just testing..")
	var nW, nB []models.Piece
	nW = copyCharacters(whitePieces)
	nB = copyCharacters(blackPieces)

	whiteKing := nW[0]
	blackKing := nB[0]

	var tB models.Board //new board instance

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			tB[i][j] = nil
		}
	}

	// new instances of pieces
	for _, piece := range nW {
		if tempBoard[piece.GetPosition()[1]][piece.GetPosition()[0]] != nil && tempBoard[piece.GetPosition()[1]][piece.GetPosition()[0]].GetID() == piece.GetID() {
			tB[piece.GetPosition()[1]][piece.GetPosition()[0]] = piece
		}
	}

	for _, piece := range nB {
		if tempBoard[piece.GetPosition()[1]][piece.GetPosition()[0]] != nil && tempBoard[piece.GetPosition()[1]][piece.GetPosition()[0]].GetID() == piece.GetID() {
			tB[piece.GetPosition()[1]][piece.GetPosition()[0]] = piece
		}
	}

	if CurrTurn == "B" {
		// check illegal move
		isChecked := isKingChecked(nW, blackKing, tB)
		if isChecked {
			character.HardMove(move.From)
			character.SetPrev(currPrev)
			return errors.New("[illegal move]: king is checked after the move"), ""
		}

		// Update previousSquare
		character.SetPrev(&move.From)

		// check for checkmates or stalemates
		isOpponentKingChecked := isKingChecked(nB, whiteKing, tB)
		canOppMove := canOpponentMakeMoves(nW, nB, tB)

		if isOpponentKingChecked {
			fmt.Println("Opponents king is checked...")
		}
		if !canOppMove {
			fmt.Println("Opponent cannot move")
		}
		if isOpponentKingChecked && !canOppMove {
			fmt.Println("Check Mate!")
			return nil, "End"
		}
		if !isOpponentKingChecked && !canOppMove {
			fmt.Println("Stale Mate!")
			return nil, "End"
		}
	} else {
		// check illegal move
		isChecked := isKingChecked(nB, whiteKing, tB)
		if isChecked {
			character.HardMove(move.From)
			character.SetPrev(currPrev)
			return errors.New("[illegal move]: king is checked after the move"), ""
		}

		// Update previousSquare
		character.SetPrev(&move.From)

		// check for checkmates or stalemates
		isOpponentKingChecked := isKingChecked(nW, blackKing, tB)
		canOppMove := canOpponentMakeMoves(nB, nW, tB)
		if isOpponentKingChecked {
			fmt.Println("Opponents king is checked...")
		}
		if !canOppMove {
			fmt.Println("Opponent cannot move")
		}
		if isOpponentKingChecked && !canOppMove {
			fmt.Println("Check Mate!")
			return nil, "End"
		}
		if !isOpponentKingChecked && !canOppMove {
			fmt.Println("Stale Mate!")
			return nil, "End"
		}
	}

	// update board
	board = tempBoard
	history = append(history, mv)

	if len(history) >= 6 && history[len(history)-1] == history[len(history)-5] && history[len(history)-4] == history[len(history)-6] {
		fmt.Println("Repetition Draw")
		return nil, "End"
	}

	PrintBoard()

	return nil, ""
}

// isKingChecked: Check whether king is checked
func isKingChecked(pieces []models.Piece, king models.Piece, board models.Board) bool {
	fmt.Printf("[isKingChecked]: Triggered for %v\n", king)
	for _, piece := range pieces {
		fmt.Printf("[isKingChecked]: Checking for %v\n", piece)
		for _, sq := range piece.GetAttackingSquares(board) {
			fmt.Printf("[isKingChecked]: Comparing %v\n", sq)
			if king.GetPosition()[0] == sq[0] && king.GetPosition()[1] == sq[1] {
				fmt.Printf("[isKingChecked]: King is checked by %v\n", piece)
				return true
			}
		}
	}
	return false
}

// canOpponentMakeMoves: Check whether opponent has any possible legal moves
func canOpponentMakeMoves(oppPieces []models.Piece, playerPieces []models.Piece, board models.Board) bool {
	for _, piece := range oppPieces {
		fmt.Printf("[canOpponentMakeMoves]: %v can move to\n", piece)
		fmt.Println(piece.GetAttackingSquares(board))
		for _, sq := range piece.GetAttackingSquares(board) {
			newSquare := board[sq[1]][sq[0]]
			fmt.Printf("[canOpponentMakeMoves]: %v:%v->%v\n", piece.GetCharacter(), piece.GetPosition(), sq)
			if newSquare == nil ||
				(newSquare != nil &&
					newSquare.GetColor() != piece.GetColor() &&
					newSquare.GetCharacter() != "K") {
				fmt.Println("[canOpponentMakeMoves]: Checking a valid move")
				tB := board
				curr := piece.GetPosition()
				currPrev := piece.GetPrev()

				fmt.Println("[canOpponentMakeMoves]: Moving piece...")
				x := piece.Move(sq, &tB)

				if x != nil {
					fmt.Println(x)
				}

				fmt.Println("[canOpponentMakeMoves]: Update state...")
				piece.SetPrev(&curr)

				fmt.Println("[canOpponentMakeMoves]: Checking for checks")
				isChecked := isKingChecked(playerPieces, oppPieces[0], tB)

				if !isChecked {
					fmt.Printf("%v:%v->%v is possible\n", piece.GetCharacter(), curr, sq)
					return true
				}

				fmt.Printf("[canOpponentMakeMoves]: Reverting to %v curr\n", curr)
				piece.HardMove(curr)
				piece.SetPrev(currPrev)
				fmt.Println(piece)
			}
		}
	}
	return false
}

// copyCharacters: Deep copy character objects
func copyCharacters(characterArr []models.Piece) []models.Piece {
	var nA []models.Piece
	for _, piece := range characterArr {
		var nP models.Piece
		switch piece.GetCharacter() {
		case "K":
			nP = &characters.King{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		case "Q":
			nP = &characters.Queen{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		case "R":
			nP = &characters.Rook{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		case "B":
			nP = &characters.Bishop{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		case "N":
			nP = &characters.Knight{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		case "P":
			nP = &characters.Pawn{
				CurrPosition: piece.GetPosition(),
				Color:        piece.GetColor(),
				Prev:         piece.GetPrev(),
				Character:    piece.GetCharacter(),
				Id:           piece.GetID(),
			}
			break
		}
		nA = append(nA, nP)
	}

	return nA
}

// formatInput: Format move input into required format
func formatInput(mv string) (models.MoveType, error) {
	// Eg: N:g1 -> f3
	mv = strings.Trim(mv, " ")

	// Separate character and move
	piece := strings.Split(mv, ":")[0]
	squares := strings.Split(strings.Split(mv, ":")[1], "->")

	// Validate move
	from, fromErr := getSquare(squares[0])
	to, toErr := getSquare(squares[1])

	if fromErr != nil {
		return models.MoveType{}, fromErr
	}

	if toErr != nil {
		return models.MoveType{}, toErr
	}

	return models.MoveType{
		Character: piece,
		From:      from,
		To:        to,
	}, nil
}

// getSquare: Convert square to coordinates in the board
// i.e. a6 -> {0, 5}
func getSquare(sq string) (models.Square, error) {
	dictionary := make(map[string]int)
	dictionary["h"] = 0
	dictionary["g"] = 1
	dictionary["f"] = 2
	dictionary["e"] = 3
	dictionary["d"] = 4
	dictionary["c"] = 5
	dictionary["b"] = 6
	dictionary["a"] = 7

	x := dictionary[string(sq[0])]
	y, err := strconv.Atoi(string(sq[1]))
	if err != nil {
		return models.Square{0, 0}, err
	}
	y--
	if !(x >= 0 && x < 8 && y >= 0 && y < 8) {
		return models.Square{0, 0}, errors.New("invalid square")
	}
	return models.Square{x, y}, nil
}

// PrintBoard : Print board array in to the console
func PrintBoard() {
	for _, i := range board {
		for _, j := range i {
			if j == nil {
				fmt.Print(" -- ")
			} else {
				fmt.Print(" " + j.GetColor() + j.GetCharacter() + " ")
			}
		}
		fmt.Println()
	}
}
