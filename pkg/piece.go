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

type Piece struct {
	colour     string
	identifier int
	row, col   int
}

func (piece Piece) toString() string {
	return getPieceSymbol(piece.identifier)
}

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
		return WhiteKnight
	case 7:
		return BlackKnight
	case 8:
		return WhiteQueen
	case 9:
		return BlackQueen
	case 10:
		return WhiteKing
	case 11:
		return BlackKing
	default:
		errorString := fmt.Sprintf("Invalid identifier: %d", identifier)
		panic(errorString)
	}
}
