package pkg

import (
	"fmt"
)

const (
	WhitePawn = "\u2659"
	BlackPawn = "\u265F"

	WhiteRook = "\u2656"
	BlackRook = "\u265C"

	WhiteBishop = "\u2657"
	BlackBishop = "\u265D"

	WhiteKnight = "\u2658"
	BlackKnight = "\u265E"

	WhiteQueen = "\u2655"
	BlackQueen = "\u265B"

	WhiteKing = "\u2654"
	BlackKing = "\u265A"
)

func createPiece(identifier int, row int, col int) Piece {
	var colour string

	if identifier%2 == 0 {
		colour = "white"
	} else {
		colour = "black"
	}

	return Piece{colour, identifier, row, col}
}

func getPieceSymbol(identifier int) string {
	switch identifier {
	case 0:
		return WhitePawn
	case 1:
		return BlackPawn
	case 2:
		return WhiteRook
	case 3:
		return BlackRook
	case 4:
		return WhiteBishop
	case 5:
		return BlackBishop
	case 6:
		return WhiteQueen
	case 7:
		return BlackQueen
	case 8:
		return WhiteKing
	case 9:
		return BlackKing
	default:
		errorString := fmt.Sprintf("Invalid identifier: %d", identifier)
		panic(errorString)
	}
}

func getPawnMoves(row int, col int, board Board, colour string) []string {
	var moves []string
	var firstMove bool

	oneStep := row
	twoStep := row

	switch colour {
	case "white":
		oneStep -= 1
		twoStep -= 2
		firstMove = row == boardSize-2
	case "black":
		oneStep += 1
		twoStep += 2
		firstMove = row == 1
	}
}
