package pkg

type Square struct {
	piece    *Piece
	row, col int
}

func (square Square) toString() string {
	if square.piece == nil {
		return ""
	}

	return square.piece.toString()
}

func formatSquare(square string) string {
	if len(square) == 0 {
		return " - |"
	}

	return " " + square + " |"
}
