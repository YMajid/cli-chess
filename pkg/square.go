package pkg

type Square struct {
	piece    *Piece
	row, col int
}

func (square Square) hasPiece() bool {
	return square.piece != nil
}

func (square *Square) setPiece(piece *Piece) {
	square.piece = piece
}

func (square Square) getPiece() *Piece {
	return square.piece
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
