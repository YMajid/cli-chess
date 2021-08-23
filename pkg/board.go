package pkg

import (
	"strings"

	"strconv"
)

const boardSize = 8

type Board struct {
	squares [][]Square
}

func CreateBoard() *Board {
	board := new(Board)

	squares := make([][]Square, boardSize)
	for i := 0; i < boardSize; i++ {
		squares[i] = make([]Square, boardSize)
		for j := 0; j < boardSize; j++ {
			square := Square{nil, i, j}
			squares[i][j] = square
		}
	}

	board.squares = squares
	return board
}

func (board Board) toString() [][]string {
	strBoard := make([][]string, boardSize)
	for i := 0; i < boardSize; i++ {
		strBoard[i] = make([]string, boardSize)
		for j := 0; j < boardSize; j++ {
			strBoard[i][j] = board.squares[i][j].toString()
		}
	}

	return strBoard
}

func formatBoard(board [][]string) string {
	var sb strings.Builder

	for i := 0; i < boardSize; i++ {
		if i == 0 {
			sb.WriteString("   ")
		}
		c := (string)(i + 'a')
		sb.WriteString(" " + c + "  ")
	}
	sb.WriteString("\n")

	for i := boardSize - 1; i >= 0; i-- {
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString(" |")

		for j := 0; j < boardSize; j++ {
			square := board[boardSize-i-1][j]
			sb.WriteString(formatSquare(square))
		}

		sb.WriteString(" " + strconv.Itoa(i+1))
		sb.WriteString("\n")
	}

	for i := 0; i < boardSize; i++ {
		if i == 0 {
			sb.WriteString("   ")
		}
		c := (string)(i + 'a')
		sb.WriteString(" " + c + "  ")
	}

	sb.WriteString("\n")
	return sb.String()
}

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
