package pkg

const boardSize = 8

func (board Board) getSquare(coordinates string) *Square {
	col := int(coordinates[0] - 'a')
	row := ((int)(coordinates[1]-'0'))*-1 + boardSize

	if row < 0 || row >= boardSize || col < 0 || col >= boardSize {
		return nil
	}

	return &board.squares[row][col]
}

func (board Board) isValidMove(coordinates string, colour string) bool {
	square := board.getSquare(coordinates)
	if square == nil {
		return false
	}

	piece := square.piece
	return piece == nil || piece.colour != colour
}

func (board Board) isEmptyAt(coordinates string) bool {
	square := board.getSquare(coordinates)
	return square != nil && square.piece != nil
}
