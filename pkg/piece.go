package piece

import (
	"fmt"
)

type Piece struct {
	colour     string
	identifier int
	row, col   int
}

const (
	WhitePawn string = "\u2659"
	BlackPawn string = "\u265F"

	WhiteRook string = "\u2656"
	BlackRook string = "\u265C"

	WhiteBishop string = "\u2657"
	BlackBishop string = "\u265D"

	WhiteKnight string = "\u2658"
	BlackKnight string = "\u265E"

	WhiteQueen string = "\u2655"
	BlackQueen string = "\u265B"

	WhiteKing string = "\u2654"
	BlackKing string = "\u265A"
)

func CreatePiece(identifier int, row int, col int) Piece {
	var colour string

	if identifier%2 == 0 {
		colour = "white"
	} else {
		colour = "black"
	}

	return Piece{colour, identifier, row, col}
}

func GetPieceSymbol(identifier int) string {
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
